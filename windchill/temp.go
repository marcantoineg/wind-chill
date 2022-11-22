package windchill

type tempUnit rune

const (
	Celsius    tempUnit = 'C'
	Fahrenheit tempUnit = 'F'
	Kelvin     tempUnit = 'K'
)

type temp struct {
	value float64
	unit  tempUnit
}

func (t temp) Celsius() temp {
	switch t.unit {
	case Fahrenheit:
		return temp{0.0, Celsius} // todo: F to C

	case Kelvin:
		return temp{0.0, Celsius} // todo F to K

	default:
		return t
	}
}

func (t temp) Fahrenheit() temp {
	switch t.unit {
	case Celsius:
		return temp{0.0, Celsius} // todo: C to F

	case Kelvin:
		return temp{0.0, Celsius} // todo C to K

	default:
		return t
	}
}

func (t temp) Kelvin() temp {
	switch t.unit {
	case Celsius:
		return temp{0.0, Celsius} // todo: C to K

	case Fahrenheit:
		return temp{0.0, Celsius} // todo: F to K

	default:
		return t
	}
}
