// package speed defines a type `Speed` to represent & convert speed units.
package speed

import "math"

type speedUnit string

const (
	KPH speedUnit = "Km/h"
	MPH speedUnit = "Mph"

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

// Kph returns a new, always positive speed converted to kph.
func (s Speed) Kph() Speed {
	s.Value = math.Abs(s.Value)
	if s.Unit == KPH {
		return s
	}
	return Speed{s.Value * mileToKm, KPH}
}

// Mph returns a new, always positive speed converted to mph.
func (s Speed) Mph() Speed {
	s.Value = math.Abs(s.Value)
	if s.Unit == MPH {
		return s
	}
	return Speed{s.Value / mileToKm, MPH}
}
