// package speed defines a type `Speed` to represent & convert speed units.
package speed

type speedUnit rune

const (
	Kph speedUnit = 'K'
	Mph speedUnit = 'M'

	mileToKm = 1.609344
)

// Speed reprensents the numerical value of a speed and it's unit.
type Speed struct {
	value float64
	unit  speedUnit
}

// Kph returns a new speed converted to km/h.
func (s Speed) Kph() Speed {
	if s.unit == Kph || s.unit < 0 {
		return s
	}
	return Speed{s.value * mileToKm, Kph}
}

// Mph returns a new speed converted to mph.
func (s Speed) Mph() Speed {
	if s.unit == Mph || s.unit < 0 {
		return s
	}
	return Speed{s.value / mileToKm, Mph}
}
