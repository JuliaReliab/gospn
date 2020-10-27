package parser

import (
	"../petrinet"
)

// Get an integer from ASTExpr. When AST cannot be converted to integer,
// the return value is x.
func getInt(expr ASTExpr, env ASTEnv, x ASTInt) (ASTInt, error) {
	if result, err := expr.Eval(env); err == nil {
		if val, err := result.GetInt(); err == nil {
			return val, nil
		} else {
			return x, err
		}
	} else {
		return x, err
	}
}

// Get a float value from ASTExpr. When AST cannot be converted to a float value,
// the return value is x.
func getFloat(expr ASTExpr, env ASTEnv, x float64) (float64, error) {
	if result, err := expr.Eval(env); err == nil {
		if val, err := result.GetInt(); err == nil {
			return float64(val), nil
		} else if val, err := result.GetFloat(); err == nil {
			return float64(val), nil
		} else {
			return x, err
		}
	} else {
		return x, err
	}
}

// Get a string from ASTExpr. When AST cannot be converted to string,
// the return value is str.
func getString(expr ASTExpr, env ASTEnv, str string) (string, error) {
	if result, err := expr.Eval(env); err == nil {
		if val, err := result.GetString(); err == nil {
			return string(val), nil
		} else {
			return str, err
		}
	} else {
		return str, err
	}
}

// Create a closure for guard function
func createGuardFunc(expr ASTExpr, net *petrinet.Net, env ASTEnv) func([]petrinet.MarkInt) bool {
	return func(mark []petrinet.MarkInt) bool {
		if result, err := expr.EvalWithMark(net, mark, env); err == nil {
			if val, err := result.GetBool(); err == nil {
				return bool(val)
			} else {
				logger.Panic(err)
			}
		} else {
			logger.Panic(err)
		}
		return false
	}
}

// Create a closure for weight and rate functions
func createRateFunc(expr ASTExpr, net *petrinet.Net, env ASTEnv) func([]petrinet.MarkInt) float64 {
	return func(mark []petrinet.MarkInt) float64 {
		if result, err := expr.EvalWithMark(net, mark, env); err == nil {
			if val, err := result.GetInt(); err == nil {
				return float64(val)
			} else if val, err := result.GetFloat(); err == nil {
				return float64(val)
			} else {
				logger.Panic(err)
			}
		} else {
			logger.Panic(err)
		}
		return 0.0
	}
}

// Create a closure for multi function
func createMultiFunc(expr ASTExpr, net *petrinet.Net, env ASTEnv) func([]petrinet.MarkInt) petrinet.MarkInt {
	return func(mark []petrinet.MarkInt) petrinet.MarkInt {
		if astval, err := expr.EvalWithMark(net, mark, env); err == nil {
			if val, err := astval.GetInt(); err == nil {
				return petrinet.MarkInt(val)
			} else {
				logger.Panic(err)
			}
		} else {
			logger.Panic(err)
		}
		return 0
	}
}

// Create a closure for reward function
func createRewardFunc(expr ASTExpr, net *petrinet.Net, env ASTEnv) func([]petrinet.MarkInt) float64 {
	return func(mark []petrinet.MarkInt) float64 {
		if result, err := expr.EvalWithMark(net, mark, env); err == nil {
			if val, err := result.GetInt(); err == nil {
				return float64(val)
			} else if val, err := result.GetFloat(); err == nil {
				return float64(val)
			} else {
				logger.Panic(err)
			}
		} else {
			logger.Panic(err)
		}
		return 0.0
	}
}

// Create au update function
func createUpdateFunc(expr ASTExpr, net *petrinet.Net, env ASTEnv) func([]petrinet.MarkInt) []petrinet.MarkInt {
	return func(mark []petrinet.MarkInt) []petrinet.MarkInt {
		if _, err := expr.EvalWithMark(net, mark, env); err == nil {
			return mark
		} else {
			logger.Panic(err)
		}
		return []petrinet.MarkInt{}
	}
}

func createPlace(net *petrinet.Net, initmark map[string]petrinet.MarkInt,
	label string, node *PNNode, env ASTEnv) *petrinet.Place {
	max := petrinet.MarkInt(255) // default value
	for attr, optval := range node.options {
		switch attr {
		case "init":
			switch expr := optval.(type) {
			case int:
				initmark[label] = petrinet.MarkInt(expr)
			case ASTExpr:
				if val, err := getInt(expr, env, 0); err == nil {
					initmark[label] = petrinet.MarkInt(val)
				} else {
					logger.Panic(err)
				}
			default:
				logger.Panicf("type mismatch in %s of %s", attr, label)
			}
		case "max":
			switch expr := optval.(type) {
			case int:
				max = petrinet.MarkInt(expr)
			case ASTExpr:
				if val, err := getInt(expr, env, 0); err == nil {
					max = petrinet.MarkInt(val)
				} else {
					logger.Panic(err)
				}
			default:
				logger.Panicf("type mismatch in %s of %s", attr, label)
			}
		default:
		}
	}
	p := net.NewPlace(label, max)
	env[label] = p
	logger.Print("Create place: label ", label, " max ", max)
	return p
}

func createImmTrans(net *petrinet.Net, label string, node *PNNode, env ASTEnv) *petrinet.ImmTrans {
	// default values
	priority := 0
	vanishable := true
	weight := 1.0
	var guardexpr ASTExpr
	var weightexpr ASTExpr
	var updateexpr ASTExpr
	for attr, optval := range node.options {
		switch attr {
		case "weight":
			switch expr := optval.(type) {
			case int:
				weight = float64(expr)
			case float64:
				weight = expr
			case ASTExpr:
				if astval, err := expr.Eval(env); err == nil {
					if val, err := astval.GetInt(); err == nil {
						weight = float64(val)
					} else if val, err := astval.GetFloat(); err == nil {
						weight = float64(val)
					} else {
						weightexpr = expr
					}
				} else {
					logger.Panic(err)
				}
			default:
				logger.Panicf("type mismatch in %s of %s", attr, label)
			}
		case "priority":
			switch expr := optval.(type) {
			case int:
				priority = expr
			case ASTExpr:
				if astval, err := expr.Eval(env); err == nil {
					if val, err := astval.GetInt(); err == nil {
						priority = int(val)
					} else {
						logger.Panic(err)
					}
				} else {
					logger.Panic(err)
				}
			default:
				logger.Panicf("type mismatch in %s of %s", attr, label)
			}
		case "vanishable":
			switch expr := optval.(type) {
			case bool:
				vanishable = expr
			case ASTExpr:
				if astval, err := expr.Eval(env); err == nil {
					if val, err := astval.GetBool(); err == nil {
						vanishable = bool(val)
					} else {
						logger.Panic(err)
					}
				} else {
					logger.Panic(err)
				}
			default:
				logger.Panicf("type mismatch in %s of %s", attr, label)
			}
		case "guard":
			if expr, ok := optval.(ASTExpr); ok {
				guardexpr = expr
			}
		case "update":
			if expr, ok := optval.(ASTExpr); ok {
				updateexpr = expr
			}
		default:
		}
	}
	tr := net.NewImmTrans(label, priority, vanishable, weight)
	if guardexpr != nil {
		fstr, err := getString(guardexpr, env, "guard of "+label)
		if err != nil {
			logger.Panic("Fail to create guardfunc of ", label, " ", guardexpr)
		}
		net.SetGuard(tr, fstr, createGuardFunc(guardexpr, net, env))
	}
	if weightexpr != nil {
		net.SetWeightRate(tr, createRateFunc(weightexpr, net, env))
	}
	if updateexpr != nil {
		fstr, err := getString(updateexpr, env, "update of "+label)
		if err != nil {
			logger.Panic("Fail to create updatefunc of ", label, " ", updateexpr)
		}
		net.SetUpdate(tr, fstr, createUpdateFunc(updateexpr, net, env))
	}
	env[label] = tr
	logger.Print("Create immtrans: label ", label, " priority ", priority, " vanishable ", vanishable, " weight ", weight, " guard ", guardexpr)
	return tr
}

func createExpTrans(net *petrinet.Net, label string, node *PNNode, env ASTEnv) *petrinet.ExpTrans {
	// default values
	priority := 0
	vanishable := true
	rate := 1.0
	var guardexpr ASTExpr
	var rateexpr ASTExpr
	var updateexpr ASTExpr
	for attr, optval := range node.options {
		switch attr {
		case "rate":
			switch expr := optval.(type) {
			case int:
				rate = float64(expr)
			case float64:
				rate = expr
			case ASTExpr:
				if astval, err := expr.Eval(env); err == nil {
					if val, err := astval.GetInt(); err == nil {
						rate = float64(val)
					} else if val, err := astval.GetFloat(); err == nil {
						rate = float64(val)
					} else {
						rateexpr = expr
					}
				} else {
					logger.Panic(err)
				}
			default:
				logger.Panicf("type mismatch in %s of %s", attr, label)
			}
		case "priority":
			switch expr := optval.(type) {
			case int:
				priority = expr
			case ASTExpr:
				if astval, err := expr.Eval(env); err == nil {
					if val, err := astval.GetInt(); err == nil {
						priority = int(val)
					} else {
						logger.Panic(err)
					}
				} else {
					logger.Panic(err)
				}
			default:
				logger.Panicf("type mismatch in %s of %s", attr, label)
			}
		case "vanishable":
			switch expr := optval.(type) {
			case bool:
				vanishable = expr
			case ASTExpr:
				if astval, err := expr.Eval(env); err == nil {
					if val, err := astval.GetBool(); err == nil {
						vanishable = bool(val)
					} else {
						logger.Panic(err)
					}
				} else {
					logger.Panic(err)
				}
			default:
				logger.Panicf("type mismatch in %s of %s", attr, label)
			}
		case "guard":
			if expr, ok := optval.(ASTExpr); ok {
				guardexpr = expr
			}
		case "update":
			if expr, ok := optval.(ASTExpr); ok {
				updateexpr = expr
			}
		default:
		}
	}
	tr := net.NewExpTrans(label, priority, vanishable, rate)
	if guardexpr != nil {
		fstr, err := getString(guardexpr, env, "guard of "+label)
		if err != nil {
			logger.Panic("Fail to create guardfunc of ", label, " ", guardexpr)
		}
		net.SetGuard(tr, fstr, createGuardFunc(guardexpr, net, env))
	}
	if rateexpr != nil {
		net.SetWeightRate(tr, createRateFunc(rateexpr, net, env))
	}
	if updateexpr != nil {
		fstr, err := getString(updateexpr, env, "update of "+label)
		if err != nil {
			logger.Panic("Fail to create updatefunc of ", label, " ", updateexpr)
		}
		net.SetUpdate(tr, fstr, createUpdateFunc(updateexpr, net, env))
	}
	env[label] = tr
	logger.Print("Create exptrans: label ", label, " priority ", priority, " vanishable ", vanishable, " rate ", rate, " guard ", guardexpr, " rate ", rateexpr)
	return tr
}

func createGenTrans(net *petrinet.Net, label string, node *PNNode, env ASTEnv) *petrinet.GenTrans {
	// default values
	priority := 0
	vanishable := true
	dist := petrinet.NewDistribution("constant", 1)
	policy := petrinet.GenTransPolicyPRD
	var guardexpr ASTExpr
	var updateexpr ASTExpr
	for attr, optval := range node.options {
		switch attr {
		case "dist":
			switch expr := optval.(type) {
			case ASTDist:
				name, err := expr.GetName().GetString()
				if err != nil {
					logger.Panic(err)
				}
				params := make([]float64, 0)
				for _, v := range expr.GetParams() {
					var parameter float64
					if y, err := v.GetInt(); err == nil {
						parameter = float64(y)
					} else if y, err := v.GetFloat(); err == nil {
						parameter = float64(y)
					} else {
						logger.Panic(err)
					}
					params = append(params, float64(parameter))
				}
				dist = petrinet.NewDistribution(string(name), params...)
			case ASTExpr:
				if astval, err := expr.Eval(env); err == nil {
					if val, err := astval.GetDist(); err == nil {
						name, err := val.GetName().GetString()
						if err != nil {
							logger.Panic(err)
						}
						params := make([]float64, 0)
						for _, v := range val.GetParams() {
							var parameter float64
							if y, err := v.GetInt(); err == nil {
								parameter = float64(y)
							} else if y, err := v.GetFloat(); err == nil {
								parameter = float64(y)
							} else {
								logger.Panic(err)
							}
							params = append(params, float64(parameter))
						}
						dist = petrinet.NewDistribution(string(name), params...)
					} else {
						logger.Panic(err)
					}
				} else {
					logger.Panic(err)
				}
			default:
				logger.Panicf("type mismatch in %s of %s", attr, label)
			}
		case "policy":
			var pol string
			switch expr := optval.(type) {
			case string:
				pol = expr
			case ASTExpr:
				if astval, err := expr.Eval(env); err == nil {
					if val, err := astval.GetString(); err == nil {
						pol = string(val)
					} else {
						logger.Panic(err)
					}
				} else {
					logger.Panic(err)
				}
			default:
				logger.Panicf("type mismatch in %s of %s", attr, label)
			}
			switch pol {
			case "prd":
				policy = petrinet.GenTransPolicyPRD
			case "prs":
				policy = petrinet.GenTransPolicyPRS
			case "pri":
				policy = petrinet.GenTransPolicyPRI
			default:
				logger.Panicf("Policy %s is undefined", pol)
			}
		case "priority":
			switch expr := optval.(type) {
			case int:
				priority = expr
			case ASTExpr:
				if astval, err := expr.Eval(env); err == nil {
					if val, err := astval.GetInt(); err == nil {
						priority = int(val)
					} else {
						logger.Panic(err)
					}
				} else {
					logger.Panic(err)
				}
			default:
				logger.Panicf("type mismatch in %s of %s", attr, label)
			}
		case "vanishable":
			switch expr := optval.(type) {
			case bool:
				vanishable = expr
			case ASTExpr:
				if astval, err := expr.Eval(env); err == nil {
					if val, err := astval.GetBool(); err == nil {
						vanishable = bool(val)
					} else {
						logger.Panic(err)
					}
				} else {
					logger.Panic(err)
				}
			default:
				logger.Panicf("type mismatch in %s of %s", attr, label)
			}
		case "guard":
			if expr, ok := optval.(ASTExpr); ok {
				guardexpr = expr
			}
		case "update":
			if expr, ok := optval.(ASTExpr); ok {
				updateexpr = expr
			}
		default:
		}
	}
	tr := net.NewGenTrans(label, priority, vanishable, dist, policy)
	if guardexpr != nil {
		fstr, err := getString(guardexpr, env, "guard of "+label)
		if err != nil {
			logger.Panic("Fail to create guardfunc of ", label, " ", guardexpr)
		}
		net.SetGuard(tr, fstr, createGuardFunc(guardexpr, net, env))
	}
	if updateexpr != nil {
		fstr, err := getString(updateexpr, env, "update of "+label)
		if err != nil {
			logger.Panic("Fail to create updatefunc of ", label, " ", updateexpr)
		}
		net.SetUpdate(tr, fstr, createUpdateFunc(updateexpr, net, env))
	}
	env[label] = tr
	logger.Print("Create gentrans: label ", label, " priority ", priority, " vanishable ", vanishable, " dist ", dist, " policy ", policy, " guard ", guardexpr)
	return tr
}

func createArc(net *petrinet.Net, node *PNNode, env ASTEnv) {
	// default values
	var src, dest string
	multi := petrinet.MarkInt(1)
	var multiexprstring string
	var multiexpr ASTExpr
	for attr, optval := range node.options {
		switch attr {
		case "src":
			switch val := optval.(type) {
			case string:
				src = val
			default:
				logger.Panicf("type mismatch in %s of %s", attr, "Arc")
			}
		case "dest":
			switch val := optval.(type) {
			case string:
				dest = val
			default:
				logger.Panicf("type mismatch in %s of %s", attr, "Arc")
			}
		case "multi":
			switch expr := optval.(type) {
			case ASTExpr:
				if astval, err := expr.Eval(env); err == nil {
					if val, err := astval.GetInt(); err == nil {
						multi = petrinet.MarkInt(val)
					} else if val, err := astval.GetString(); err == nil {
						multiexprstring = string(val)
						multiexpr = expr
					}
				} else {
					logger.Panic(err)
				}
			case int:
				multi = petrinet.MarkInt(expr)
			default:
				logger.Panicf("type mismatch in %s of %s", attr, "Arc")
			}
		default:
		}
	}
	if place, ok := net.GetPlace(src); ok {
		if tr, ok := net.GetTrans(dest); ok {
			tr := net.NewInArc(place, tr, multi)
			if multiexpr != nil {
				net.SetInArcMulti(tr, multiexprstring, createMultiFunc(multiexpr, net, env))
			}
			logger.Print("Create inarc: src", src, " dest ", dest, " multi ", multi, " multiexpr ", multiexpr)
		} else {
			logger.Panicf("%s is a place but %s is not a trans", src, dest)
		}
	} else if tr, ok := net.GetTrans(src); ok {
		if place, ok := net.GetPlace(dest); ok {
			tr := net.NewOutArc(tr, place, multi)
			if multiexpr != nil {
				net.SetOutArcMulti(tr, multiexprstring, createMultiFunc(multiexpr, net, env))
			}
			logger.Print("Create outarc: src", src, " dest ", dest, " multi ", multi, " multiexpr ", multiexpr)
		} else {
			logger.Panicf("%s is a trans but %s is not a place", src, dest)
		}
	} else {
		logger.Panicf("%s is neither in place and trans.", src)
	}
}

func createHArc(net *petrinet.Net, node *PNNode, env ASTEnv) {
	// default values
	var src, dest string
	multi := petrinet.MarkInt(1)
	var multiexpr ASTExpr
	var multiexprstring string
	for attr, optval := range node.options {
		switch attr {
		case "src":
			switch val := optval.(type) {
			case string:
				src = val
			default:
				logger.Panicf("type mismatch in %s of %s", attr, "Arc")
			}
		case "dest":
			switch val := optval.(type) {
			case string:
				dest = val
			default:
				logger.Panicf("type mismatch in %s of %s", attr, "Arc")
			}
		case "multi":
			switch expr := optval.(type) {
			case ASTExpr:
				if astval, err := expr.Eval(env); err == nil {
					if val, err := astval.GetInt(); err == nil {
						multi = petrinet.MarkInt(val)
					} else if val, err := astval.GetString(); err == nil {
						multiexprstring = string(val)
						multiexpr = astval
					}
				} else {
					logger.Panic(err)
				}
			case int:
				multi = petrinet.MarkInt(expr)
			default:
				logger.Panicf("type mismatch in %s of %s", attr, "Arc")
			}
		default:
		}
	}
	if place, ok := net.GetPlace(src); ok {
		if tr, ok := net.GetTrans(dest); ok {
			tr := net.NewInhibitArc(place, tr, multi)
			if multiexpr != nil {
				net.SetInArcMulti(tr, multiexprstring, createMultiFunc(multiexpr, net, env))
			}
		} else {
			logger.Panicf("%s is a place but %s is not a trans", src, dest)
		}
		logger.Print("Create harc: src", src, " dest ", dest, " multi ", multi, " multiexpr ", multiexpr)
	} else {
		logger.Panicf("%s is not a place", src)
	}
}

func createReward(net *petrinet.Net, label string, node *PNNode, env ASTEnv) {
	var f ASTExpr
	for attr, optval := range node.options {
		switch attr {
		case "formula":
			switch expr := optval.(type) {
			case int:
				f = MakeValue(expr)
			case float64:
				f = MakeValue(expr)
			case ASTExpr:
				f = expr
			default:
				logger.Panicf("type mismatch in %s of %s", attr, "Reward")
			}
		default:
		}
	}
	net.SetReward(label, createRewardFunc(f, net, env))
	logger.Print("Create reward: label ", label, " formula ", f)
}

func makeNet(labels []string, env ASTEnv) (*petrinet.Net, []petrinet.MarkInt) {
	logger.Print("Start compile ...")
	net := petrinet.NewNet()
	initmark := make(map[string]petrinet.MarkInt)

	for _, label := range labels {
		value := env[label]
		if node, ok := value.(*PNNode); ok {
			switch node.options["type"] {
			case "place":
				createPlace(net, initmark, label, node, env)
			case "imm":
				createImmTrans(net, label, node, env)
			case "exp":
				createExpTrans(net, label, node, env)
			case "gen":
				createGenTrans(net, label, node, env)
			default:
			}
		}
	}

	for _, label := range labels {
		value := env[label]
		if node, ok := value.(*PNNode); ok {
			switch node.options["type"] {
			case "reward":
				createReward(net, label, node, env)
			default:
			}
		}
	}

	for _, label := range labels {
		value := env[label]
		if node, ok := value.(*PNNode); ok {
			switch node.options["type"] {
			case "arc":
				createArc(net, node, env)
			case "harc":
				createHArc(net, node, env)
			default:
			}
		}
	}

	net.Finalize()

	m := net.MakeMark(initmark)

	logger.Print("Done: compile")
	return net, m
}
