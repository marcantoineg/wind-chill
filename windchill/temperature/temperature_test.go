package temperature

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestC(t *testing.T) {
	test(t, []testrun{
		{given: C(0.), expected: Temperature{0., Celsius}},
		{given: C(10.5), expected: Temperature{10.5, Celsius}},
		{given: C(-10.), expected: Temperature{-10., Celsius}},
	}, func(tr testrun) {
		assert.Equal(tr.t, tr.expected, tr.given)
	})
}

func TestF(t *testing.T) {
	test(t, []testrun{
		{given: F(0.), expected: Temperature{0., Fahrenheit}},
		{given: F(10.5), expected: Temperature{10.5, Fahrenheit}},
		{given: F(-10.), expected: Temperature{-10., Fahrenheit}},
	}, func(tr testrun) {
		assert.Equal(tr.t, tr.expected, tr.given)
	})
}

func TestK(t *testing.T) {
	test(t, []testrun{
		{given: K(0.), expected: Temperature{0., Kelvin}},
		{given: K(10.5), expected: Temperature{10.5, Kelvin}},
		{given: K(-10.), expected: Temperature{-10., Kelvin}},
	}, func(tr testrun) {
		assert.Equal(tr.t, tr.expected, tr.given)
	})
}

func TestCelsius(t *testing.T) {
	test(t, []testrun{
		// °C to °C
		{given: C(0), expected: C(0)},
		{given: C(0.5), expected: C(0.5)},
		{given: C(100), expected: C(100)},
		{given: C(-12), expected: C(-12)},
		// °F to °C
		{given: F(-103), expected: C(-75)},
		{given: F(0), expected: C(-17.7778)},
		{given: F(32), expected: C(0)},
		{given: F(212), expected: C(100)},
		// °K to °C
		{given: K(-10), expected: C(-283.15)},
		{given: K(0), expected: C(-273.15)},
		{given: K(273.15), expected: C(0)},
		{given: K(373.15), expected: C(100)},
	}, func(tr testrun) {
		actual := tr.given.Celsius()
		assert.InDelta(tr.t, tr.expected.Value, actual.Value, 0.0001)
		assert.Equal(tr.t, tr.expected.Unit, actual.Unit)
	})
}

func TestFahrenheit(t *testing.T) {
	test(t, []testrun{
		// °C to °F
		{given: C(-75), expected: F(-103)},
		{given: C(-17.7778), expected: F(0)},
		{given: C(0), expected: F(32)},
		{given: C(100), expected: F(212)},
		// °F to °F
		{given: F(0), expected: F(0)},
		{given: F(0.5), expected: F(0.5)},
		{given: F(100), expected: F(100)},
		{given: F(-12), expected: F(-12)},
		// °K to °F
		{given: K(-10), expected: F(-477.67)},
		{given: K(0), expected: F(-459.67)},
		{given: K(255.372), expected: F(0)},
		{given: K(273.15), expected: F(32)},
	}, func(tr testrun) {
		actual := tr.given.Fahrenheit()
		assert.InDelta(tr.t, tr.expected.Value, actual.Value, 0.001)
		assert.Equal(tr.t, tr.expected.Unit, actual.Unit)
	})
}

func TestKelvin(t *testing.T) {
	test(t, []testrun{
		// °C to °K
		{given: C(-283.15), expected: K(-10)},
		{given: C(-273.15), expected: K(0)},
		{given: C(0), expected: K(273.15)},
		{given: C(100), expected: K(373.15)},
		// °F to °K
		{given: F(-477.67), expected: K(-10)},
		{given: F(-459.67), expected: K(0)},
		{given: F(0), expected: K(255.372)},
		{given: F(32), expected: K(273.15)},
		// °K to °K
		{given: K(-10), expected: K(-10)},
		{given: K(0), expected: K(0)},
		{given: K(273.15), expected: K(273.15)},
		{given: K(373.15), expected: K(373.15)},
	}, func(tr testrun) {
		actual := tr.given.Kelvin()
		assert.InDelta(tr.t, tr.expected.Value, actual.Value, 0.001)
		assert.Equal(tr.t, tr.expected.Unit, actual.Unit)
	})
}

type testrun struct {
	given    Temperature
	expected Temperature
	t        *testing.T
}

func test(tt *testing.T, testruns []testrun, testblock func(tr testrun)) {
	for _, tr := range testruns {
		tt.Run(fmt.Sprintf("°%c to °%c", tr.given.Unit, tr.expected.Unit), func(t *testing.T) {
			tr.t = t
			testblock(tr)
		})
	}
}
