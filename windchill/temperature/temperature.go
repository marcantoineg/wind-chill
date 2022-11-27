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
	Value float64
	Unit  TemperatureUnit
}

// C returns a new Celsius Temperature with its Value set to the given float.
func C(f float64) Temperature {
	return Temperature{f, Celsius}
}

// F returns a new Fahrenheit Temperature with its Value set to the given float.
func F(f float64) Temperature {
	return Temperature{f, Fahrenheit}
}

// K returns a new Kelvin Temperature with its Value set to the given float.
func K(f float64) Temperature {
	return Temperature{f, Kelvin}
}

type converter struct {
	convertTo TemperatureUnit
	fromC     func(v float64) float64
	fromF     func(v float64) float64
	fromK     func(v float64) float64
}

// Celsius returns a new Temperature converted to Celsius.
func (t Temperature) Celsius() Temperature {
	return t.convert(
		converter{
			convertTo: Celsius,
			fromC: func(v float64) float64 {
				return t.Value
			},
			fromF: func(v float64) float64 {
				// °C = (°F − 32) * 5/9
				return (t.Value - 32.) * (5. / 9.)
			},
			fromK: func(v float64) float64 {
				// °C = °K - 273.15
				return t.Value - 273.15
			},
		},
	)
}

// Fahrenheit returns a new Temperature converted to Fahrenheit.
func (t Temperature) Fahrenheit() Temperature {
	return t.convert(
		converter{
			convertTo: Fahrenheit,
			fromC: func(v float64) float64 {
				// °F = (°C * 9/5) + 32
				return (t.Value * (9. / 5.)) + 32.
			},
			fromF: func(v float64) float64 {
				return t.Value
			},
			fromK: func(v float64) float64 {
				// °F = (°K − 273.15) * 9/5 + 32
				return (t.Value-273.15)*(9./5.) + 32
			},
		},
	)
}

// Kelvin returns a new Temperature converted to Kelvin.
func (t Temperature) Kelvin() Temperature {
	return t.convert(
		converter{
			convertTo: Kelvin,
			fromC: func(v float64) float64 {
				// °K = °C + 273.15
				return t.Value + 273.15
			},
			fromF: func(v float64) float64 {
				// °K = 5/9 * (°F + 459.67)
				return (5. / 9.) * (t.Value + 459.67)
			},
			fromK: func(v float64) float64 {
				return t.Value
			},
		},
	)
}

func (t Temperature) convert(c converter) Temperature {
	var convertedValue float64
	switch t.Unit {
	case Celsius:
		convertedValue = c.fromC(t.Value)
	case Fahrenheit:
		convertedValue = c.fromF(t.Value)
	case Kelvin:
		convertedValue = c.fromK(t.Value)
	default:
		panic("unit not supported")
	}

	return Temperature{convertedValue, c.convertTo}
}
