// package speed defines a type `Speed` to represent & convert speed units.
package speed

type speedUnit rune

const (
	KPH speedUnit = 'K'
	MPH speedUnit = 'M'

	mileToKm = 1.609344
)

// Speed reprensents the numerical value of a speed and it's unit.
type Speed struct {
	Value float64
	Unit  speedUnit
}

// Kph returns a new KPH speed with its Value set to the given float.
func Kph(v float64) Speed {
	return Speed{v, KPH}
}

// Kph returns a new MPH speed with its Value set to the given float.
func Mph(v float64) Speed {
	return Speed{v, MPH}
}

// KPH returns a new speed converted to km/h.
func (s Speed) Kph() Speed {
	if s.Unit == KPH || s.Unit < 0 {
		return s
	}
	return Speed{s.Value * mileToKm, KPH}
}

// MPH returns a new speed converted to mph.
func (s Speed) Mph() Speed {
	if s.Unit == MPH || s.Unit < 0 {
		return s
	}
	return Speed{s.Value / mileToKm, MPH}
}
