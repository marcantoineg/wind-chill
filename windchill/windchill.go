package windchill

import (
	"math"

	"github.com/marcantoineg/windchill/windchill/speed"
	temp "github.com/marcantoineg/windchill/windchill/temperature"
)

// GetWindChillIndex returns the windchill index for a given air temperature and wind speed.
// The Unit of the new returned Temperature will match `t`'s Unit.
// for more information see https://journals.ametsoc.org/view/journals/bams/86/10/bams-86-10-1453.xml?tab_body=pdf
func GetWindChillIndex(t temp.Temperature, s speed.Speed) temp.Temperature {
	t_a := t.Celsius().Value
	v_pow := math.Pow(s.Kph().Value, 0.16) // v^0.16

	// T_wc = 13.12 + (0.6215 * T_a) - (11.37 * (v^0.16)) + (0.3965 * T_a * (v^0.16))
	// where `T_wc` is the winchill index in °C, `T_a` is the air temperature in °C and `v` is the wind speed in km/h at 10m.
	wcFactorInC := temp.Temperature{
		Value: (13.12 + (0.6215 * t_a) - (11.37 * (v_pow)) + (0.3965 * t_a * v_pow)),
		Unit:  temp.Celsius,
	}

	switch t.Unit {
	case temp.Celsius:
		return wcFactorInC
	case temp.Fahrenheit:
		return wcFactorInC.Fahrenheit()
	case temp.Kelvin:
		return wcFactorInC.Kelvin()
	}

	return temp.Temperature{}
}
