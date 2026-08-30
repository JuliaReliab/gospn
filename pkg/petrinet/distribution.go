package petrinet

import (
	"fmt"
	"log"
	"math"
)

type RandomNumberGenerator interface {
	Float64() float64 // return [0,1) uniform random number
}

// DistributionInterface is a firing-time distribution of a GEN transition.
//
// String renders it in the definition language -- `det(2)`, `unif(1,3)`, `expdist(0.5)`
// -- so that a result file can say which distribution governs a regeneration group. The
// block matrices alone cannot: a general block is a 0/1 jump matrix, unchanged by the
// distribution, so a file without this is not enough to solve the process it describes.
type DistributionInterface interface {
	Float64(RandomNumberGenerator) float64
	String() string
}

func NewDistribution(dist string, params ...float64) DistributionInterface {
	switch dist {
	case "constant":
		return &ConstantDist{
			x: params[0],
		}
	case "uniform":
		return &UniformDist{
			min: params[0],
			max: params[1],
		}
	case "exponential":
		return &ExpDist{
			rate: params[0],
		}
	default:
		log.Panicf("Distribution %s is not implemented", dist)
		return nil
	}
}

type ConstantDist struct {
	x float64
}

func (d *ConstantDist) Float64(_ RandomNumberGenerator) float64 {
	return d.x
}

func (d *ConstantDist) String() string { return fmt.Sprintf("det(%v)", d.x) }

type UniformDist struct {
	min float64
	max float64
}

func (d *UniformDist) Float64(rng RandomNumberGenerator) float64 {
	return (d.max-d.min)*rng.Float64() + d.min
}

func (d *UniformDist) String() string { return fmt.Sprintf("unif(%v,%v)", d.min, d.max) }

type ExpDist struct {
	rate float64
}

func (d *ExpDist) Float64(rng RandomNumberGenerator) float64 {
	return -1 / d.rate * math.Log(rng.Float64())
}

func (d *ExpDist) String() string { return fmt.Sprintf("expdist(%v)", d.rate) }
