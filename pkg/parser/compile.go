package parser

// Compiles the expression language into typed Go closures, so that a rate, guard,
// multiplicity or reward costs one chain of direct calls instead of a walk over the
// AST.
//
// What the walk costs, and why this exists: EvalWithMark boxes every intermediate
// value into an *ASTValue (allocation per node), looks up each variable in the
// ASTEnv by string, and looks up each `#place` with net.GetPlace by string. None
// of that depends on the marking -- the environment is fixed once parsing is done
// -- so all of it can happen once, here.
//
// Anything outside the supported subset returns ok=false and the caller keeps the
// interpreted closure, so behaviour is unchanged for what this cannot compile.

import (
	"math"

	"github.com/okamumu/gospn/pkg/petrinet"
)

// CompileStats counts how much of a net compiled rather than falling back. Without
// it, a change that quietly stops compiling anything looks exactly like a change
// that compiles everything and does not help.
type compileStats struct {
	RateCompiled, RateFallback     int
	GuardCompiled, GuardFallback   int
	MultiCompiled, MultiFallback   int
	RewardCompiled, RewardFallback int
}

var CompileStats compileStats

// CheckCompiled makes every compiled closure also run the interpreter and panic on
// the first disagreement, which is how the two implementations are kept honest --
// see TestCompiledAgreesWithInterpreter. It roughly doubles the work, so it is off
// outside that test.
var CheckCompiled bool

func CompileStatsReset() { CompileStats = compileStats{} }

// The static type of a compiled subexpression.
type ctype int

const (
	cInt ctype = iota
	cFloat
	cBool
)

// A compiled expression: exactly one of the three closures is set, per typ.
// constant marks a subexpression that does not read the marking, which powf needs
// to decide statically whether its result is an integer.
type compiled struct {
	typ      ctype
	constant bool
	i        func([]petrinet.MarkInt) ASTInt
	f        func([]petrinet.MarkInt) float64
	b        func([]petrinet.MarkInt) bool
}

// asFloat views a compiled int or float subexpression as a float64 closure,
// matching the evaluator's promotion of int to float in mixed arithmetic.
func (c compiled) asFloat() func([]petrinet.MarkInt) float64 {
	if c.typ == cFloat {
		return c.f
	}
	ci := c.i
	return func(m []petrinet.MarkInt) float64 { return float64(ci(m)) }
}

// A compiler carries the state that inlining needs: the closure already built for
// each variable, and which variables are currently being compiled.
//
// Both matter. Inlining a variable substitutes its whole subtree, so without the
// cache a chain of definitions -- `enall` referring to `enall.1`..`enall.3`, each
// of those to more -- duplicates work exponentially; the first version of this
// spike bounded only the recursion depth and was killed for running out of
// memory. inProgress catches a variable that refers to itself, which the
// interpreter would loop on forever.
type compiler struct {
	net        *petrinet.Net
	env        ASTEnv
	cache      map[string]compiled
	cacheOK    map[string]bool
	inProgress map[string]bool
}

func newCompiler(net *petrinet.Net, env ASTEnv) *compiler {
	return &compiler{
		net:        net,
		env:        env,
		cache:      make(map[string]compiled),
		cacheOK:    make(map[string]bool),
		inProgress: make(map[string]bool),
	}
}

// compileExpr compiles expr, or reports ok=false if it is outside the subset.
func (c *compiler) compileExpr(expr ASTExpr) (compiled, bool) {
	net, env := c.net, c.env
	switch e := expr.(type) {
	case *ASTValue:
		switch v := e.val.(type) {
		case ASTInt:
			k := v
			return compiled{typ: cInt, constant: true, i: func([]petrinet.MarkInt) ASTInt { return k }}, true
		case ASTFloat:
			k := float64(v)
			return compiled{typ: cFloat, constant: true, f: func([]petrinet.MarkInt) float64 { return k }}, true
		case ASTBool:
			k := bool(v)
			return compiled{typ: cBool, constant: true, b: func([]petrinet.MarkInt) bool { return k }}, true
		}
		return compiled{}, false

	case *ASTVar:
		// The environment is fixed by the time closures run, so a variable can be
		// inlined instead of looked up by string on every evaluation.
		if got, seen := c.cache[e.label]; seen {
			return got, c.cacheOK[e.label]
		}
		if c.inProgress[e.label] {
			return compiled{}, false // cyclic definition
		}
		value, ok := env[e.label]
		if !ok {
			return compiled{}, false
		}
		inner, ok := value.(ASTExpr)
		if !ok {
			return compiled{}, false
		}
		c.inProgress[e.label] = true
		got, gotOK := c.compileExpr(inner)
		delete(c.inProgress, e.label)
		c.cache[e.label], c.cacheOK[e.label] = got, gotOK
		return got, gotOK

	case *ASTNToken:
		place, ok := net.GetPlace(e.label)
		if !ok {
			return compiled{}, false
		}
		// The place is resolved here -- that is the string map lookup this whole
		// exercise is about removing -- but its *index* is read at evaluation time.
		// Net.Finalize assigns the indices, and it runs after every closure has been
		// built, so an index captured now would be the pre-Finalize placeholder 0 for
		// every place. That reads the wrong token count without failing anywhere
		// obvious: it took a differential run against the interpreter to see it.
		return compiled{typ: cInt, i: func(m []petrinet.MarkInt) ASTInt {
			return ASTInt(m[place.GetIndex()])
		}}, true

	case *ASTEnableCond:
		tr, ok := net.GetTrans(e.label)
		if !ok {
			return compiled{}, false
		}
		return compiled{typ: cBool, b: func(m []petrinet.MarkInt) bool {
			return tr.IsEnabled(net, m) == petrinet.ENABLE
		}}, true

	case *ASTTriOp:
		return c.compileTri(e)

	case *ASTUnaryOp:
		return c.compileUnary(e)

	case *ASTBinOp:
		return c.compileBinary(e)
	}
	return compiled{}, false
}

func (c *compiler) compileUnary(e *ASTUnaryOp) (compiled, bool) {
	x, ok := c.compileExpr(e.elem)
	if !ok {
		return compiled{}, false
	}
	switch e.op {
	case "uminus":
		if x.typ == cInt {
			xi := x.i
			return compiled{typ: cInt, i: func(m []petrinet.MarkInt) ASTInt { return -xi(m) }}, true
		}
		if x.typ == cFloat {
			xf := x.f
			return compiled{typ: cFloat, f: func(m []petrinet.MarkInt) float64 { return -xf(m) }}, true
		}
	case "uplus":
		if x.typ != cBool {
			return x, true
		}
	case "not":
		if x.typ == cBool {
			xb := x.b
			return compiled{typ: cBool, b: func(m []petrinet.MarkInt) bool { return !xb(m) }}, true
		}
	}
	return compiled{}, false
}

func (c *compiler) compileBinary(e *ASTBinOp) (compiled, bool) {
	l, ok := c.compileExpr(e.left)
	if !ok {
		return compiled{}, false
	}
	r, ok := c.compileExpr(e.right)
	if !ok {
		return compiled{}, false
	}
	bothInt := l.typ == cInt && r.typ == cInt
	bothBool := l.typ == cBool && r.typ == cBool
	numeric := l.typ != cBool && r.typ != cBool

	switch e.op {
	case "plus", "minus", "mul":
		if !numeric {
			return compiled{}, false
		}
		if bothInt {
			li, ri := l.i, r.i
			switch e.op {
			case "plus":
				return compiled{typ: cInt, i: func(m []petrinet.MarkInt) ASTInt { return li(m) + ri(m) }}, true
			case "minus":
				return compiled{typ: cInt, i: func(m []petrinet.MarkInt) ASTInt { return li(m) - ri(m) }}, true
			default:
				return compiled{typ: cInt, i: func(m []petrinet.MarkInt) ASTInt { return li(m) * ri(m) }}, true
			}
		}
		lf, rf := l.asFloat(), r.asFloat()
		switch e.op {
		case "plus":
			return compiled{typ: cFloat, f: func(m []petrinet.MarkInt) float64 { return lf(m) + rf(m) }}, true
		case "minus":
			return compiled{typ: cFloat, f: func(m []petrinet.MarkInt) float64 { return lf(m) - rf(m) }}, true
		default:
			return compiled{typ: cFloat, f: func(m []petrinet.MarkInt) float64 { return lf(m) * rf(m) }}, true
		}

	case "div":
		// Note int / int yields a float here, matching ASTInt.div; `idiv` is the
		// integer division.
		if !numeric {
			return compiled{}, false
		}
		lf, rf := l.asFloat(), r.asFloat()
		return compiled{typ: cFloat, f: func(m []petrinet.MarkInt) float64 { return lf(m) / rf(m) }}, true

	case "idiv":
		if !bothInt {
			return compiled{}, false
		}
		li, ri := l.i, r.i
		return compiled{typ: cInt, i: func(m []petrinet.MarkInt) ASTInt { return li(m) / ri(m) }}, true

	case "and", "or":
		if !bothBool {
			return compiled{}, false
		}
		lb, rb := l.b, r.b
		if e.op == "and" {
			return compiled{typ: cBool, b: func(m []petrinet.MarkInt) bool { return lb(m) && rb(m) }}, true
		}
		return compiled{typ: cBool, b: func(m []petrinet.MarkInt) bool { return lb(m) || rb(m) }}, true

	case "eq", "neq":
		if bothBool {
			lb, rb := l.b, r.b
			if e.op == "eq" {
				return compiled{typ: cBool, b: func(m []petrinet.MarkInt) bool { return lb(m) == rb(m) }}, true
			}
			return compiled{typ: cBool, b: func(m []petrinet.MarkInt) bool { return lb(m) != rb(m) }}, true
		}
		fallthrough
	case "lt", "lte", "gt", "gte":
		if !numeric {
			return compiled{}, false
		}
		if bothInt {
			li, ri := l.i, r.i
			switch e.op {
			case "eq":
				return compiled{typ: cBool, b: func(m []petrinet.MarkInt) bool { return li(m) == ri(m) }}, true
			case "neq":
				return compiled{typ: cBool, b: func(m []petrinet.MarkInt) bool { return li(m) != ri(m) }}, true
			case "lt":
				return compiled{typ: cBool, b: func(m []petrinet.MarkInt) bool { return li(m) < ri(m) }}, true
			case "lte":
				return compiled{typ: cBool, b: func(m []petrinet.MarkInt) bool { return li(m) <= ri(m) }}, true
			case "gt":
				return compiled{typ: cBool, b: func(m []petrinet.MarkInt) bool { return li(m) > ri(m) }}, true
			default:
				return compiled{typ: cBool, b: func(m []petrinet.MarkInt) bool { return li(m) >= ri(m) }}, true
			}
		}
		lf, rf := l.asFloat(), r.asFloat()
		switch e.op {
		case "eq":
			return compiled{typ: cBool, b: func(m []petrinet.MarkInt) bool { return lf(m) == rf(m) }}, true
		case "neq":
			return compiled{typ: cBool, b: func(m []petrinet.MarkInt) bool { return lf(m) != rf(m) }}, true
		case "lt":
			return compiled{typ: cBool, b: func(m []petrinet.MarkInt) bool { return lf(m) < rf(m) }}, true
		case "lte":
			return compiled{typ: cBool, b: func(m []petrinet.MarkInt) bool { return lf(m) <= rf(m) }}, true
		case "gt":
			return compiled{typ: cBool, b: func(m []petrinet.MarkInt) bool { return lf(m) > rf(m) }}, true
		default:
			return compiled{typ: cBool, b: func(m []petrinet.MarkInt) bool { return lf(m) >= rf(m) }}, true
		}

	case "powf":
		if !numeric {
			return compiled{}, false
		}
		if bothInt {
			// ASTInt.powf uses the integer powi for a non-negative exponent and
			// math.Pow otherwise, so the result type depends on the exponent's sign
			// and is only static when the exponent is a constant.
			k, isConst := constInt(r)
			if !isConst {
				return compiled{}, false
			}
			li := l.i
			if k >= 0 {
				return compiled{typ: cInt, i: func(m []petrinet.MarkInt) ASTInt {
					return ASTInt(powi(int32(li(m)), int32(k)))
				}}, true
			}
			return compiled{typ: cFloat, f: func(m []petrinet.MarkInt) float64 {
				return math.Pow(float64(li(m)), float64(k))
			}}, true
		}
		lf, rf := l.asFloat(), r.asFloat()
		return compiled{typ: cFloat, f: func(m []petrinet.MarkInt) float64 { return math.Pow(lf(m), rf(m)) }}, true

	case "max", "min":
		if !numeric {
			return compiled{}, false
		}
		if bothInt {
			li, ri := l.i, r.i
			if e.op == "max" {
				return compiled{typ: cInt, i: func(m []petrinet.MarkInt) ASTInt {
					if a, b := li(m), ri(m); a >= b {
						return a
					} else {
						return b
					}
				}}, true
			}
			return compiled{typ: cInt, i: func(m []petrinet.MarkInt) ASTInt {
				if a, b := li(m), ri(m); a <= b {
					return a
				} else {
					return b
				}
			}}, true
		}
		lf, rf := l.asFloat(), r.asFloat()
		if e.op == "max" {
			return compiled{typ: cFloat, f: func(m []petrinet.MarkInt) float64 { return math.Max(lf(m), rf(m)) }}, true
		}
		return compiled{typ: cFloat, f: func(m []petrinet.MarkInt) float64 { return math.Min(lf(m), rf(m)) }}, true
	}
	return compiled{}, false
}

// constInt reports the value of a compiled int subexpression when it does not depend
// on the marking. Only constants are recognised, which is all powf needs.
func constInt(c compiled) (ASTInt, bool) {
	if c.typ != cInt || c.i == nil {
		return 0, false
	}
	if !c.constant {
		return 0, false
	}
	return c.i(nil), true
}

// compileTri compiles `ifelse(cond, a, b)`. ASTBool.ite returns whichever branch was
// taken, so the interpreter's result type is whatever that branch happened to be;
// a compiled closure needs one static type, so the branches must agree, or be int
// and float and promote.
func (c *compiler) compileTri(e *ASTTriOp) (compiled, bool) {
	if e.op != "ite" {
		return compiled{}, false
	}
	cond, ok := c.compileExpr(e.first)
	if !ok || cond.typ != cBool {
		return compiled{}, false
	}
	l, ok := c.compileExpr(e.second)
	if !ok {
		return compiled{}, false
	}
	r, ok := c.compileExpr(e.third)
	if !ok {
		return compiled{}, false
	}
	cb := cond.b
	switch {
	case l.typ == cBool && r.typ == cBool:
		lb, rb := l.b, r.b
		return compiled{typ: cBool, b: func(m []petrinet.MarkInt) bool {
			if cb(m) {
				return lb(m)
			}
			return rb(m)
		}}, true
	case l.typ == cInt && r.typ == cInt:
		li, ri := l.i, r.i
		return compiled{typ: cInt, i: func(m []petrinet.MarkInt) ASTInt {
			if cb(m) {
				return li(m)
			}
			return ri(m)
		}}, true
	case l.typ != cBool && r.typ != cBool:
		lf, rf := l.asFloat(), r.asFloat()
		return compiled{typ: cFloat, f: func(m []petrinet.MarkInt) float64 {
			if cb(m) {
				return lf(m)
			}
			return rf(m)
		}}, true
	}
	return compiled{}, false
}

// compileRate compiles a numeric expression to the shape createRateFunc needs.
func compileRate(expr ASTExpr, net *petrinet.Net, env ASTEnv) (func([]petrinet.MarkInt) float64, bool) {
	r, ok := newCompiler(net, env).compileExpr(expr)
	if !ok || r.typ == cBool {
		return nil, false
	}
	return r.asFloat(), true
}

// compileMulti compiles an integer expression to the shape createMultiFunc needs.
func compileMulti(expr ASTExpr, net *petrinet.Net, env ASTEnv) (func([]petrinet.MarkInt) petrinet.MarkInt, bool) {
	r, ok := newCompiler(net, env).compileExpr(expr)
	if !ok || r.typ != cInt {
		return nil, false
	}
	ri := r.i
	return func(m []petrinet.MarkInt) petrinet.MarkInt { return petrinet.MarkInt(ri(m)) }, true
}

// compileGuard compiles a boolean expression to the shape createGuardFunc needs.
func compileGuard(expr ASTExpr, net *petrinet.Net, env ASTEnv) (func([]petrinet.MarkInt) bool, bool) {
	r, ok := newCompiler(net, env).compileExpr(expr)
	if !ok || r.typ != cBool {
		return nil, false
	}
	return r.b, true
}
