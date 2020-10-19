package petrinet

import (
	"log"
	"math"
)

type RandomNumberGenerator interface {
	Float64() float64 // return [0,1) uniform random number
}

type DistributionInterface interface {
	Float64(RandomNumberGenerator) float64
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

type UniformDist struct {
	min float64
	max float64
}

func (d *UniformDist) Float64(rng RandomNumberGenerator) float64 {
	return (d.max-d.min)*rng.Float64() + d.min
}

type ExpDist struct {
	rate float64
}

func (d *ExpDist) Float64(rng RandomNumberGenerator) float64 {
	return -1 / d.rate * math.Log(rng.Float64())
}
