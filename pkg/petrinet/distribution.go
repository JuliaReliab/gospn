package petrinet

type Distribution struct {
	name   string
	params []float64
	cdf    func(float64) float64
}

func NewDistribution(name string, params []float64) *Distribution {
	return &Distribution{
		name:   name,
		params: params,
	}
}

func (d *Distribution) SetCDF(f func(float64) float64) {
	d.cdf = f
}
