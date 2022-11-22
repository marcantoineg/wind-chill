package windchill

type speedUnit rune

const (
	Kph speedUnit = 'K'
	Mph speedUnit = 'M'

	mileToKm = 1.609344
)

type speed struct {
	value float64
	unit  speedUnit
}

func (s speed) Kph() speed {
	if s.unit == Kph || s.unit < 0 {
		return s
	}
	return speed{s.value * mileToKm, Kph}
}

func (s speed) Mph() speed {
	if s.unit == Mph || s.unit < 0 {
		return s
	}
	return speed{s.value / mileToKm, Mph}
}
