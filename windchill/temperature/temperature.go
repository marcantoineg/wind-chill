// package temperature defines a type `Temperature` to represent & convert temperature units.
package temperature

type TemperatureUnit rune

const (
	Celsius    TemperatureUnit = 'C'
	Fahrenheit TemperatureUnit = 'F'
	Kelvin     TemperatureUnit = 'K'
)

// Temperature is a type used to describe the numerical value of a Temperature and it's unit.
type Temperature struct {
	value float64
	unit  TemperatureUnit
}

// C returns a new Celsius Temperature with its value set to the given float.
func C(f float64) Temperature {
	return Temperature{f, Celsius}
}

// F returns a new Fahrenheit Temperature with its value set to the given float.
func F(f float64) Temperature {
	return Temperature{f, Fahrenheit}
}

// K returns a new Kelvin Temperature with its value set to the given float.
func K(f float64) Temperature {
	return Temperature{f, Kelvin}
}

// Celsius returns a new Temperature converted to Celsius.
func (t Temperature) Celsius() Temperature {
	switch t.unit {
	case Fahrenheit:
		return Temperature{
			// °C = (°F − 32) * 5/9
			(t.value - 32.) * (5. / 9.),
			Celsius,
		}

	case Kelvin:
		return Temperature{
			// °C = °K - 273.15
			t.value - 273.15,
			Celsius,
		}

	default:
		return t
	}
}

// Fahrenheit returns a new Temperature converted to Fahrenheit.
func (t Temperature) Fahrenheit() Temperature {
	switch t.unit {
	case Celsius:
		return Temperature{
			// °F = (°C * 9/5) + 32
			(t.value * (9. / 5.)) + 32.,
			Fahrenheit,
		}

	case Kelvin:
		return Temperature{
			// °F = (°K − 273.15) * 9/5 + 32
			(t.value-273.15)*(9./5.) + 32,
			Fahrenheit,
		}

	default:
		return t
	}
}

// Kelvin returns a new Temperature converted to Kelvin.
func (t Temperature) Kelvin() Temperature {
	switch t.unit {
	case Celsius:
		return Temperature{
			// °K = °C + 273.15
			t.value + 273.15,
			Kelvin,
		}

	case Fahrenheit:
		return Temperature{
			// °K = 5/9 * (°F + 459.67)
			(5. / 9.) * (t.value + 459.67),
			Kelvin,
		}

	default:
		return t
	}
}
