package parser

import (
	"errors"
	"fmt"
	"log"
	"math"
)

type ASTInt int32
type ASTFloat float64
type ASTBool bool
type ASTString string

// The structure of ASTDist
type ASTDist struct {
	name   *ASTValue
	params []*ASTValue
}

func (a ASTDist) GetName() *ASTValue {
	return a.name
}

func (a ASTDist) GetParams() []*ASTValue {
	return a.params
}

// type ASTPlace *petrinet.Place

// Operations for ASTValue
// - uplus: unaryy plus
// - uminus: unarry minus
// - not
// - plus
// - minus
// - mul
// - idiv
// - div
// - and
// - or
// - eq
// - neq
// - lt
// - lte
// - gt
// - gte
// - ite (if-then-else)

func (a ASTInt) uplus() (*ASTValue, error) {
	return MakeValue(a), nil
}

func (a ASTFloat) uplus() (*ASTValue, error) {
	return MakeValue(a), nil
}

func (a ASTString) uplus() (*ASTValue, error) {
	return MakeValue(fmt.Sprintf("(+%s)", a)), nil
}

//

func (a ASTInt) uminus() (*ASTValue, error) {
	return MakeValue(-a), nil
}

func (a ASTFloat) uminus() (*ASTValue, error) {
	return MakeValue(-a), nil
}

func (a ASTString) uminus() (*ASTValue, error) {
	return MakeValue(fmt.Sprintf("(-%s)", a)), nil
}

//

func (a ASTBool) not() (*ASTValue, error) {
	return MakeValue(!a), nil
}

func (a ASTString) not() (*ASTValue, error) {
	return MakeValue(fmt.Sprintf("(!%s)", a)), nil
}

///

func (a ASTInt) plus(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(a + v), nil
	case ASTFloat:
		return MakeValue(ASTFloat(a) + v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%d + %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to plus")
	}
}

func (a ASTFloat) plus(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(a + ASTFloat(v)), nil
	case ASTFloat:
		return MakeValue(a + v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%e + %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to plus")
	}
}

func (a ASTString) plus(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(fmt.Sprintf("(%s + %d)", a, v)), nil
	case ASTFloat:
		return MakeValue(fmt.Sprintf("(%s + %e)", a, v)), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%s + %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to plus")
	}
}

///

func (a ASTInt) minus(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(a + v), nil
	case ASTFloat:
		return MakeValue(ASTFloat(a) - v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%d - %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to minus")
	}
}

func (a ASTFloat) minus(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(a - ASTFloat(v)), nil
	case ASTFloat:
		return MakeValue(a - v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%e - %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to minus")
	}
}

func (a ASTString) minus(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(fmt.Sprintf("(%s - %d)", a, v)), nil
	case ASTFloat:
		return MakeValue(fmt.Sprintf("(%s - %e)", a, v)), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%s - %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to minus")
	}
}

///

func (a ASTInt) mul(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(a * v), nil
	case ASTFloat:
		return MakeValue(ASTFloat(a) * v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%d * %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to multiple")
	}
}

func (a ASTFloat) mul(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(a * ASTFloat(v)), nil
	case ASTFloat:
		return MakeValue(a * v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%e * %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to mul")
	}
}

func (a ASTString) mul(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(fmt.Sprintf("(%s * %d)", a, v)), nil
	case ASTFloat:
		return MakeValue(fmt.Sprintf("(%s * %e)", a, v)), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%s * %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to mul")
	}
}

///

func (a ASTInt) idiv(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(a / v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%d / %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to idiv")
	}
}

func (a ASTString) idiv(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(fmt.Sprintf("(%s / %d)", a, v)), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%s / %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to idiv")
	}
}

///

func (a ASTInt) div(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(float64(a) / float64(v)), nil
	case ASTFloat:
		return MakeValue(ASTFloat(a) / v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%d / %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to div")
	}
}

func (a ASTFloat) div(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(a / ASTFloat(v)), nil
	case ASTFloat:
		return MakeValue(a / v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%e / %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to div")
	}
}

func (a ASTString) div(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(fmt.Sprintf("(%s / %d)", a, v)), nil
	case ASTFloat:
		return MakeValue(fmt.Sprintf("(%s / %e)", a, v)), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%s / %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to div")
	}
}

///

func (a ASTBool) and(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTBool:
		return MakeValue(a && v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%t && %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to and")
	}
}

func (a ASTString) and(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTBool:
		return MakeValue(fmt.Sprintf("(%s && %t)", a, v)), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%s && %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to and")
	}
}

//

func (a ASTBool) or(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTBool:
		return MakeValue(a || v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%t || %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to or")
	}
}

func (a ASTString) or(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTBool:
		return MakeValue(fmt.Sprintf("(%s || %t)", a, v)), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%s || %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to or")
	}
}

//

func (a ASTBool) eq(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTBool:
		return MakeValue(a == v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%t == %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to eq")
	}
}

func (a ASTInt) eq(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(a == v), nil
	case ASTFloat:
		return MakeValue(ASTFloat(a) == v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%d == %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to eq")
	}
}

func (a ASTFloat) eq(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(a == ASTFloat(v)), nil
	case ASTFloat:
		return MakeValue(a == v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%e == %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to eq")
	}
}

func (a ASTString) eq(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTBool:
		return MakeValue(fmt.Sprintf("(%s == %t)", a, v)), nil
	case ASTInt:
		return MakeValue(fmt.Sprintf("(%s == %d)", a, v)), nil
	case ASTFloat:
		return MakeValue(fmt.Sprintf("(%s == %e)", a, v)), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%s == %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to eq")
	}
}

///

func (a ASTBool) neq(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTBool:
		return MakeValue(a != v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%t != %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to neq")
	}
}

func (a ASTInt) neq(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(a != v), nil
	case ASTFloat:
		return MakeValue(ASTFloat(a) == v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%d != %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to neq")
	}
}

func (a ASTFloat) neq(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(a != ASTFloat(v)), nil
	case ASTFloat:
		return MakeValue(a != v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%e != %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to neq")
	}
}

func (a ASTString) neq(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTBool:
		return MakeValue(fmt.Sprintf("(%s != %t)", a, v)), nil
	case ASTInt:
		return MakeValue(fmt.Sprintf("(%s != %d)", a, v)), nil
	case ASTFloat:
		return MakeValue(fmt.Sprintf("(%s != %e)", a, v)), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%s != %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to neq")
	}
}

///

func (a ASTInt) lt(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(a < v), nil
	case ASTFloat:
		return MakeValue(ASTFloat(a) < v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%d < %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to lt")
	}
}

func (a ASTFloat) lt(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(a < ASTFloat(v)), nil
	case ASTFloat:
		return MakeValue(a < v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%e < %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to lt")
	}
}

func (a ASTString) lt(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(fmt.Sprintf("(%s < %d)", a, v)), nil
	case ASTFloat:
		return MakeValue(fmt.Sprintf("(%s < %e)", a, v)), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%s < %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to lt")
	}
}

///

func (a ASTInt) lte(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(a <= v), nil
	case ASTFloat:
		return MakeValue(ASTFloat(a) <= v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%d <= %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to lte")
	}
}

func (a ASTFloat) lte(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(a <= ASTFloat(v)), nil
	case ASTFloat:
		return MakeValue(a <= v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%e <= %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to lte")
	}
}

func (a ASTString) lte(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(fmt.Sprintf("(%s <= %d)", a, v)), nil
	case ASTFloat:
		return MakeValue(fmt.Sprintf("(%s <= %e)", a, v)), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%s <= %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to lte")
	}
}

///

func (a ASTInt) gt(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(a > v), nil
	case ASTFloat:
		return MakeValue(ASTFloat(a) > v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%d > %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to gt")
	}
}

func (a ASTFloat) gt(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(a > ASTFloat(v)), nil
	case ASTFloat:
		return MakeValue(a > v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%e > %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to gt")
	}
}

func (a ASTString) gt(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(fmt.Sprintf("(%s > %d)", a, v)), nil
	case ASTFloat:
		return MakeValue(fmt.Sprintf("(%s > %e)", a, v)), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%s > %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to gt")
	}
}

///

func (a ASTInt) gte(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(a >= v), nil
	case ASTFloat:
		return MakeValue(ASTFloat(a) >= v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%d >= %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to gte")
	}
}

func (a ASTFloat) gte(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(a >= ASTFloat(v)), nil
	case ASTFloat:
		return MakeValue(a >= v), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%e >= %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to gte")
	}
}

func (a ASTString) gte(b interface{}) (*ASTValue, error) {
	switch v := b.(type) {
	case ASTInt:
		return MakeValue(fmt.Sprintf("(%s >= %d)", a, v)), nil
	case ASTFloat:
		return MakeValue(fmt.Sprintf("(%s >= %e)", a, v)), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("(%s >= %s)", a, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to gte")
	}
}

///

func (a ASTBool) ite(b interface{}, c interface{}) (*ASTValue, error) {
	if a {
		return MakeValue(b), nil
	} else {
		return MakeValue(c), nil
	}
}

func (a ASTString) ite(b interface{}, c interface{}) (*ASTValue, error) {
	var s1, s2 string
	switch v1 := b.(type) {
	case ASTInt:
		s1 = fmt.Sprintf("%d", v1)
	case ASTFloat:
		s1 = fmt.Sprintf("%e", v1)
	case ASTString:
		s1 = fmt.Sprintf("%s", v1)
	default:
		return MakeValue(nil), errors.New("fail to ite")
	}
	switch v2 := c.(type) {
	case ASTInt:
		s2 = fmt.Sprintf("%d", v2)
	case ASTFloat:
		s2 = fmt.Sprintf("%e", v2)
	case ASTString:
		s2 = fmt.Sprintf("%s", v2)
	default:
		return MakeValue(nil), errors.New("fail to ite")
	}
	return MakeValue(fmt.Sprintf("(%s ? %s : %s)", a, s1, s2)), nil
}

// The structure to represent the AST value. The value can take ASTInt, ASTBool, ASTFloat, ASTString.
type ASTValue struct {
	val interface{}
}

// The function to create the ASTValue
func MakeValue(val interface{}) *ASTValue {
	switch v := val.(type) {
	case ASTValue:
		return &v
	case *ASTValue:
		return v
	case bool:
		return &ASTValue{
			val: ASTBool(v),
		}
	case int:
		return &ASTValue{
			val: ASTInt(int32(v)),
		}
	case int32:
		return &ASTValue{
			val: ASTInt(v),
		}
	case float64:
		return &ASTValue{
			val: ASTFloat(v),
		}
	case string:
		return &ASTValue{
			val: ASTString(v),
		}
	case ASTBool:
		return &ASTValue{
			val: v,
		}
	case ASTInt:
		return &ASTValue{
			val: v,
		}
	case ASTFloat:
		return &ASTValue{
			val: v,
		}
	case ASTString:
		return &ASTValue{
			val: v,
		}
	case ASTDist:
		return &ASTValue{
			val: v,
		}
	case nil:
		return &ASTValue{
			val: v,
		}
	default:
		log.Print("value is not a common type ", val)
		log.Panicf("value is not a common type %T", val)
		return &ASTValue{
			val: val,
		}
	}
}

// The function to get value
func (a *ASTValue) GetBool() (ASTBool, error) {
	switch v := a.val.(type) {
	case ASTBool:
		return v, nil
	default:
		// log.Print(v)
		return false, fmt.Errorf("The value is not bool %T", v)
	}
}

func (a *ASTValue) GetInt() (ASTInt, error) {
	switch v := a.val.(type) {
	case ASTInt:
		return v, nil
	default:
		// log.Print(v)
		return 0, fmt.Errorf("The value is not int32 %T", v)
	}
}

func (a *ASTValue) GetFloat() (ASTFloat, error) {
	switch v := a.val.(type) {
	case ASTFloat:
		return v, nil
	default:
		// log.Print(v)
		return 0, fmt.Errorf("The value is not float64 %T", v)
	}
}

func (a *ASTValue) GetString() (ASTString, error) {
	switch v := a.val.(type) {
	case ASTString:
		return v, nil
	default:
		// log.Print(v)
		return "", fmt.Errorf("The value is not string %T", v)
	}
}

func (a *ASTValue) GetDist() (ASTDist, error) {
	switch v := a.val.(type) {
	case ASTDist:
		return v, nil
	default:
		// log.Print(v)
		return ASTDist{}, fmt.Errorf("The value is not dist %T", v)
	}
}

func uplus(x *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTInt:
		return v.uplus()
	case ASTFloat:
		return v.uplus()
	case ASTString:
		return v.uplus()
	default:
		return MakeValue(nil), errors.New("fail to uplus")
	}
}

func uminus(x *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTInt:
		return v.uminus()
	case ASTFloat:
		return v.uminus()
	case ASTString:
		return v.uminus()
	default:
		return MakeValue(nil), errors.New("fail to uminus")
	}
}

func not(x *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTBool:
		return v.not()
	case ASTString:
		return v.not()
	default:
		return MakeValue(nil), errors.New("fail to not")
	}
}

func plus(x, y *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTInt:
		return v.plus(y.val)
	case ASTFloat:
		return v.plus(y.val)
	case ASTString:
		return v.plus(y.val)
	default:
		return MakeValue(nil), errors.New("fail to plus")
	}
}

func minus(x, y *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTInt:
		return v.minus(y.val)
	case ASTFloat:
		return v.minus(y.val)
	case ASTString:
		return v.minus(y.val)
	default:
		return MakeValue(nil), errors.New("fail to minus")
	}
}

func mul(x, y *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTInt:
		return v.mul(y.val)
	case ASTFloat:
		return v.mul(y.val)
	case ASTString:
		return v.mul(y.val)
	default:
		return MakeValue(nil), errors.New("fail to mul")
	}
}

func idiv(x, y *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTInt:
		return v.idiv(y.val)
	case ASTString:
		return v.idiv(y.val)
	default:
		return MakeValue(nil), errors.New("fail to idiv")
	}
}

func div(x, y *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTInt:
		return v.div(y.val)
	case ASTFloat:
		return v.div(y.val)
	case ASTString:
		return v.div(y.val)
	default:
		return MakeValue(nil), errors.New("fail to div")
	}
}

func and(x, y *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTBool:
		return v.and(y.val)
	case ASTString:
		return v.and(y.val)
	default:
		return MakeValue(nil), errors.New("fail to and")
	}
}

func or(x, y *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTBool:
		return v.or(y.val)
	case ASTString:
		return v.or(y.val)
	default:
		return MakeValue(nil), errors.New("fail to or")
	}
}

func eq(x, y *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTBool:
		return v.eq(y.val)
	case ASTInt:
		return v.eq(y.val)
	case ASTFloat:
		return v.eq(y.val)
	case ASTString:
		return v.eq(y.val)
	default:
		return MakeValue(nil), fmt.Errorf("fail to eq x:%s", x)
	}
}

func neq(x, y *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTBool:
		return v.neq(y.val)
	case ASTInt:
		return v.neq(y.val)
	case ASTFloat:
		return v.neq(y.val)
	case ASTString:
		return v.neq(y.val)
	default:
		return MakeValue(nil), errors.New("fail to neq")
	}
}

func lt(x, y *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTInt:
		return v.lt(y.val)
	case ASTFloat:
		return v.lt(y.val)
	case ASTString:
		return v.lt(y.val)
	default:
		return MakeValue(nil), errors.New("fail to lt")
	}
}

func lte(x, y *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTInt:
		return v.lte(y.val)
	case ASTFloat:
		return v.lte(y.val)
	case ASTString:
		return v.lte(y.val)
	default:
		return MakeValue(nil), errors.New("fail to lte")
	}
}

func gt(x, y *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTInt:
		return v.gt(y.val)
	case ASTFloat:
		return v.gt(y.val)
	case ASTString:
		return v.gt(y.val)
	default:
		return MakeValue(nil), errors.New("fail to gt")
	}
}

func gte(x, y *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTInt:
		return v.gte(y.val)
	case ASTFloat:
		return v.gte(y.val)
	case ASTString:
		return v.gte(y.val)
	default:
		return MakeValue(nil), errors.New("fail to gte")
	}
}

func ite(x, y, z *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTBool:
		return v.ite(y.val, z.val)
	case ASTString:
		return v.ite(y.val, z.val)
	default:
		return MakeValue(nil), errors.New("fail to ite")
	}
}

//

func expf(x *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTInt:
		return MakeValue(math.Exp(float64(v))), nil
	case ASTFloat:
		return MakeValue(math.Exp(float64(v))), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("exp(%s)", v)), nil
	default:
		return MakeValue(nil), errors.New("fail to exp")
	}
}

func logf(x *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTInt:
		return MakeValue(math.Log(float64(v))), nil
	case ASTFloat:
		return MakeValue(math.Log(float64(v))), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("log(%s)", v)), nil
	default:
		return MakeValue(nil), errors.New("fail to log")
	}
}

func sqrtf(x *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTInt:
		return MakeValue(math.Sqrt(float64(v))), nil
	case ASTFloat:
		return MakeValue(math.Sqrt(float64(v))), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("sqrt(%s)", v)), nil
	default:
		return MakeValue(nil), errors.New("fail to sqrt")
	}
}

///

func powf(x, y *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTInt:
		return v.powf(y.val)
	case ASTFloat:
		return v.powf(y.val)
	case ASTString:
		return v.powf(y.val)
	default:
		return MakeValue(nil), errors.New("fail to pow")
	}
}

func (x ASTInt) powf(y interface{}) (*ASTValue, error) {
	switch v := y.(type) {
	case ASTInt:
		return MakeValue(powi(int32(x), int32(v))), nil
	case ASTFloat:
		return MakeValue(math.Pow(float64(x), float64(v))), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("pow(%d, %s)", x, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to pow")
	}
}

func (x ASTFloat) powf(y interface{}) (*ASTValue, error) {
	switch v := y.(type) {
	case ASTInt:
		return MakeValue(math.Pow(float64(x), float64(v))), nil
	case ASTFloat:
		return MakeValue(math.Pow(float64(x), float64(v))), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("pow(%e, %s)", x, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to pow")
	}
}

func (x ASTString) powf(y interface{}) (*ASTValue, error) {
	switch v := y.(type) {
	case ASTInt:
		return MakeValue(fmt.Sprintf("pow(%s, %d)", x, v)), nil
	case ASTFloat:
		return MakeValue(fmt.Sprintf("pow(%s, %e", x, v)), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("pow(%s, %s)", x, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to pow")
	}
}

func powi(x, n int32) int32 {
	ans := int32(1)
	for n != 0 {
		if n%2 == 1 {
			ans *= x
		}
		x *= x
		n >>= 1
	}
	return ans
}

///

func max(x, y *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTInt:
		return v.max(y.val)
	case ASTFloat:
		return v.max(y.val)
	case ASTString:
		return v.max(y.val)
	default:
		return MakeValue(nil), errors.New("fail to max")
	}
}

func (x ASTInt) max(y interface{}) (*ASTValue, error) {
	switch v := y.(type) {
	case ASTInt:
		if x >= v {
			return MakeValue(x), nil
		} else {
			return MakeValue(v), nil
		}
	case ASTFloat:
		return MakeValue(math.Max(float64(x), float64(v))), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("max(%d, %s)", x, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to max")
	}
}

func (x ASTFloat) max(y interface{}) (*ASTValue, error) {
	switch v := y.(type) {
	case ASTInt:
		return MakeValue(math.Max(float64(x), float64(v))), nil
	case ASTFloat:
		return MakeValue(math.Max(float64(x), float64(v))), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("max(%e, %s)", x, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to max")
	}
}

func (x ASTString) max(y interface{}) (*ASTValue, error) {
	switch v := y.(type) {
	case ASTInt:
		return MakeValue(fmt.Sprintf("max(%s, %d)", x, v)), nil
	case ASTFloat:
		return MakeValue(fmt.Sprintf("max(%s, %e", x, v)), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("max(%s, %s)", x, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to max")
	}
}

///

func min(x, y *ASTValue) (*ASTValue, error) {
	switch v := x.val.(type) {
	case ASTInt:
		return v.min(y.val)
	case ASTFloat:
		return v.min(y.val)
	case ASTString:
		return v.min(y.val)
	default:
		return MakeValue(nil), errors.New("fail to min")
	}
}

func (x ASTInt) min(y interface{}) (*ASTValue, error) {
	switch v := y.(type) {
	case ASTInt:
		if x <= v {
			return MakeValue(x), nil
		} else {
			return MakeValue(v), nil
		}
	case ASTFloat:
		return MakeValue(math.Min(float64(x), float64(v))), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("min(%d, %s)", x, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to min")
	}
}

func (x ASTFloat) min(y interface{}) (*ASTValue, error) {
	switch v := y.(type) {
	case ASTInt:
		return MakeValue(math.Min(float64(x), float64(v))), nil
	case ASTFloat:
		return MakeValue(math.Min(float64(x), float64(v))), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("min(%e, %s)", x, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to min")
	}
}

func (x ASTString) min(y interface{}) (*ASTValue, error) {
	switch v := y.(type) {
	case ASTInt:
		return MakeValue(fmt.Sprintf("min(%s, %d)", x, v)), nil
	case ASTFloat:
		return MakeValue(fmt.Sprintf("min(%s, %e", x, v)), nil
	case ASTString:
		return MakeValue(fmt.Sprintf("min(%s, %s)", x, v)), nil
	default:
		return MakeValue(nil), errors.New("fail to min")
	}
}
