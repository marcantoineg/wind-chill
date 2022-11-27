package windchill

import (
	"fmt"
	"testing"

	"github.com/marcantoineg/windchill/windchill/speed"
	temp "github.com/marcantoineg/windchill/windchill/temperature"
	"github.com/stretchr/testify/assert"
)

const testNameTemplate = "air temp (°%s): %f | wind speed (%s): %f => expects (°%s): %f"

type testrun struct {
	givenTemp  temp.Temperature
	givenSpeed speed.Speed
	expected   temp.Temperature
}

func TestGetWindChillIndex(t *testing.T) {
	testruns := []testrun{
		// 10 km/h
		// givenTemp, givenSpeed, expectedTemp
		{temp.C(10), speed.Kph(10), temp.C(9)},
		{temp.C(5), speed.Kph(10), temp.C(3)},
		{temp.C(0), speed.Kph(10), temp.C(-3)},
		{temp.C(-5), speed.Kph(10), temp.C(-9)},
		{temp.C(-10), speed.Kph(10), temp.C(-15)},
		{temp.C(-15), speed.Kph(10), temp.C(-21)},
		{temp.C(-20), speed.Kph(10), temp.C(-27)},
		{temp.C(-25), speed.Kph(10), temp.C(-33)},
		{temp.C(-30), speed.Kph(10), temp.C(-39)},
		{temp.C(-35), speed.Kph(10), temp.C(-45)},
		{temp.C(-40), speed.Kph(10), temp.C(-51)},
		{temp.C(-45), speed.Kph(10), temp.C(-57)},
		{temp.C(-50), speed.Kph(10), temp.C(-63)},

		// 15 km/h
		// givenTemp, givenSpeed, expectedTemp
		{temp.C(10), speed.Kph(15), temp.C(8)},
		{temp.C(5), speed.Kph(15), temp.C(2)},
		{temp.C(0), speed.Kph(15), temp.C(-4)},
		{temp.C(-5), speed.Kph(15), temp.C(-11)},
		{temp.C(-10), speed.Kph(15), temp.C(-17)},
		{temp.C(-15), speed.Kph(15), temp.C(-23)},
		{temp.C(-20), speed.Kph(15), temp.C(-29)},
		{temp.C(-25), speed.Kph(15), temp.C(-35)},
		{temp.C(-30), speed.Kph(15), temp.C(-41)},
		{temp.C(-35), speed.Kph(15), temp.C(-48)},
		{temp.C(-40), speed.Kph(15), temp.C(-54)},
		{temp.C(-45), speed.Kph(15), temp.C(-60)},
		{temp.C(-50), speed.Kph(15), temp.C(-66)},

		// 20 km/h
		// givenTemp, givenSpeed, expectedTemp
		{temp.C(10), speed.Kph(20), temp.C(7)},
		{temp.C(5), speed.Kph(20), temp.C(1)},
		{temp.C(0), speed.Kph(20), temp.C(-5)},
		{temp.C(-5), speed.Kph(20), temp.C(-12)},
		{temp.C(-10), speed.Kph(20), temp.C(-18)},
		{temp.C(-15), speed.Kph(20), temp.C(-24)},
		{temp.C(-20), speed.Kph(20), temp.C(-31)},
		{temp.C(-25), speed.Kph(20), temp.C(-37)},
		{temp.C(-30), speed.Kph(20), temp.C(-43)},
		{temp.C(-35), speed.Kph(20), temp.C(-49)},
		{temp.C(-40), speed.Kph(20), temp.C(-56)},
		{temp.C(-45), speed.Kph(20), temp.C(-62)},
		{temp.C(-50), speed.Kph(20), temp.C(-68)},

		// 25 km/h
		// givenTemp, givenSpeed, expectedTemp
		{temp.C(10), speed.Kph(25), temp.C(7)},
		{temp.C(5), speed.Kph(25), temp.C(1)},
		{temp.C(0), speed.Kph(25), temp.C(-6)},
		{temp.C(-5), speed.Kph(25), temp.C(-12)},
		{temp.C(-10), speed.Kph(25), temp.C(-19)},
		{temp.C(-15), speed.Kph(25), temp.C(-25)},
		{temp.C(-20), speed.Kph(25), temp.C(-32)},
		{temp.C(-25), speed.Kph(25), temp.C(-38)},
		{temp.C(-30), speed.Kph(25), temp.C(-45)},
		{temp.C(-35), speed.Kph(25), temp.C(-51)},
		{temp.C(-40), speed.Kph(25), temp.C(-57)},
		{temp.C(-45), speed.Kph(25), temp.C(-64)},
		{temp.C(-50), speed.Kph(25), temp.C(-70)},

		// 30 km/h
		// givenTemp, givenSpeed, expectedTemp
		{temp.C(10), speed.Kph(30), temp.C(7)},
		{temp.C(5), speed.Kph(30), temp.C(0)},
		{temp.C(0), speed.Kph(30), temp.C(-7)},
		{temp.C(-5), speed.Kph(30), temp.C(-13)},
		{temp.C(-10), speed.Kph(30), temp.C(-19)},
		{temp.C(-15), speed.Kph(30), temp.C(-26)},
		{temp.C(-20), speed.Kph(30), temp.C(-33)},
		{temp.C(-25), speed.Kph(30), temp.C(-39)},
		{temp.C(-30), speed.Kph(30), temp.C(-46)},
		{temp.C(-35), speed.Kph(30), temp.C(-52)},
		{temp.C(-40), speed.Kph(30), temp.C(-59)},
		{temp.C(-45), speed.Kph(30), temp.C(-65)},
		{temp.C(-50), speed.Kph(30), temp.C(-72)},

		// 35 km/h
		// givenTemp, givenSpeed, expectedTemp
		{temp.C(10), speed.Kph(35), temp.C(6)},
		{temp.C(5), speed.Kph(35), temp.C(0)},
		{temp.C(0), speed.Kph(35), temp.C(-7)},
		{temp.C(-5), speed.Kph(35), temp.C(-13)},
		{temp.C(-10), speed.Kph(35), temp.C(-20)},
		{temp.C(-15), speed.Kph(35), temp.C(-27)},
		{temp.C(-20), speed.Kph(35), temp.C(-33)},
		{temp.C(-25), speed.Kph(35), temp.C(-40)},
		{temp.C(-30), speed.Kph(35), temp.C(-47)},
		{temp.C(-35), speed.Kph(35), temp.C(-53)},
		{temp.C(-40), speed.Kph(35), temp.C(-60)},
		{temp.C(-45), speed.Kph(35), temp.C(-66)},
		{temp.C(-50), speed.Kph(35), temp.C(-73)},

		// 40 km/h
		// givenTemp, givenSpeed, expectedTemp
		{temp.C(10), speed.Kph(40), temp.C(6)},
		{temp.C(5), speed.Kph(40), temp.C(-1)},
		{temp.C(0), speed.Kph(40), temp.C(-7)},
		{temp.C(-5), speed.Kph(40), temp.C(-14)},
		{temp.C(-10), speed.Kph(40), temp.C(-21)},
		{temp.C(-15), speed.Kph(40), temp.C(-27)},
		{temp.C(-20), speed.Kph(40), temp.C(-34)},
		{temp.C(-25), speed.Kph(40), temp.C(-41)},
		{temp.C(-30), speed.Kph(40), temp.C(-48)},
		{temp.C(-35), speed.Kph(40), temp.C(-54)},
		{temp.C(-40), speed.Kph(40), temp.C(-61)},
		{temp.C(-45), speed.Kph(40), temp.C(-68)},
		{temp.C(-50), speed.Kph(40), temp.C(-74)},

		// 45 km/h
		// givenTemp, givenSpeed, expectedTemp
		{temp.C(10), speed.Kph(45), temp.C(6)},
		{temp.C(5), speed.Kph(45), temp.C(-1)},
		{temp.C(0), speed.Kph(45), temp.C(-8)},
		{temp.C(-5), speed.Kph(45), temp.C(-15)},
		{temp.C(-10), speed.Kph(45), temp.C(-21)},
		{temp.C(-15), speed.Kph(45), temp.C(-28)},
		{temp.C(-20), speed.Kph(45), temp.C(-35)},
		{temp.C(-25), speed.Kph(45), temp.C(-42)},
		{temp.C(-30), speed.Kph(45), temp.C(-48)},
		{temp.C(-35), speed.Kph(45), temp.C(-55)},
		{temp.C(-40), speed.Kph(45), temp.C(-62)},
		{temp.C(-45), speed.Kph(45), temp.C(-69)},
		{temp.C(-50), speed.Kph(45), temp.C(-75)},

		// 50 km/h
		// givenTemp, givenSpeed, expectedTemp
		{temp.C(10), speed.Kph(50), temp.C(6)},
		{temp.C(5), speed.Kph(50), temp.C(-1)},
		{temp.C(0), speed.Kph(50), temp.C(-8)},
		{temp.C(-5), speed.Kph(50), temp.C(-15)},
		{temp.C(-10), speed.Kph(50), temp.C(-22)},
		{temp.C(-15), speed.Kph(50), temp.C(-29)},
		{temp.C(-20), speed.Kph(50), temp.C(-35)},
		{temp.C(-25), speed.Kph(50), temp.C(-42)},
		{temp.C(-30), speed.Kph(50), temp.C(-49)},
		{temp.C(-35), speed.Kph(50), temp.C(-56)},
		{temp.C(-40), speed.Kph(50), temp.C(-63)},
		{temp.C(-45), speed.Kph(50), temp.C(-70)},
		{temp.C(-50), speed.Kph(50), temp.C(-76)},

		// 55 km/h
		// givenTemp, givenSpeed, expectedTemp
		{temp.C(10), speed.Kph(55), temp.C(5)},
		{temp.C(5), speed.Kph(55), temp.C(-2)},
		{temp.C(0), speed.Kph(55), temp.C(-9)},
		{temp.C(-5), speed.Kph(55), temp.C(-15)},
		{temp.C(-10), speed.Kph(55), temp.C(-22)},
		{temp.C(-15), speed.Kph(55), temp.C(-29)},
		{temp.C(-20), speed.Kph(55), temp.C(-36)},
		{temp.C(-25), speed.Kph(55), temp.C(-43)},
		{temp.C(-30), speed.Kph(55), temp.C(-50)},
		{temp.C(-35), speed.Kph(55), temp.C(-57)},
		{temp.C(-40), speed.Kph(55), temp.C(-63)},
		{temp.C(-45), speed.Kph(55), temp.C(-70)},
		{temp.C(-50), speed.Kph(55), temp.C(-77)},

		// 60 km/h
		// givenTemp, givenSpeed, expectedTemp
		{temp.C(10), speed.Kph(60), temp.C(5)},
		{temp.C(5), speed.Kph(60), temp.C(-2)},
		{temp.C(0), speed.Kph(60), temp.C(-9)},
		{temp.C(-5), speed.Kph(60), temp.C(-16)},
		{temp.C(-10), speed.Kph(60), temp.C(-23)},
		{temp.C(-15), speed.Kph(60), temp.C(-30)},
		{temp.C(-20), speed.Kph(60), temp.C(-37)},
		{temp.C(-25), speed.Kph(60), temp.C(-43)},
		{temp.C(-30), speed.Kph(60), temp.C(-50)},
		{temp.C(-35), speed.Kph(60), temp.C(-57)},
		{temp.C(-40), speed.Kph(60), temp.C(-64)},
		{temp.C(-45), speed.Kph(60), temp.C(-71)},
		{temp.C(-50), speed.Kph(60), temp.C(-78)},

		// 70 km/h
		// givenTemp, givenSpeed, expectedTemp
		{temp.C(10), speed.Kph(70), temp.C(5)},
		{temp.C(5), speed.Kph(70), temp.C(-2)},
		{temp.C(0), speed.Kph(70), temp.C(-9)},
		{temp.C(-5), speed.Kph(70), temp.C(-16)},
		{temp.C(-10), speed.Kph(70), temp.C(-23)},
		{temp.C(-15), speed.Kph(70), temp.C(-30)},
		{temp.C(-20), speed.Kph(70), temp.C(-37)},
		{temp.C(-25), speed.Kph(70), temp.C(-44)},
		{temp.C(-30), speed.Kph(70), temp.C(-51)},
		{temp.C(-35), speed.Kph(70), temp.C(-59)},
		{temp.C(-40), speed.Kph(70), temp.C(-66)},
		{temp.C(-45), speed.Kph(70), temp.C(-73)},
		{temp.C(-50), speed.Kph(70), temp.C(-80)},

		// 80 km/h
		// givenTemp, givenSpeed, expectedTemp
		{temp.C(10), speed.Kph(80), temp.C(4)},
		{temp.C(5), speed.Kph(80), temp.C(-3)},
		{temp.C(0), speed.Kph(80), temp.C(-10)},
		{temp.C(-5), speed.Kph(80), temp.C(-17)},
		{temp.C(-10), speed.Kph(80), temp.C(-24)},
		{temp.C(-15), speed.Kph(80), temp.C(-31)},
		{temp.C(-20), speed.Kph(80), temp.C(-38)},
		{temp.C(-25), speed.Kph(80), temp.C(-45)},
		{temp.C(-30), speed.Kph(80), temp.C(-52)},
		{temp.C(-35), speed.Kph(80), temp.C(-60)},
		{temp.C(-40), speed.Kph(80), temp.C(-67)},
		{temp.C(-45), speed.Kph(80), temp.C(-74)},
		{temp.C(-50), speed.Kph(80), temp.C(-81)},
	}

	for _, tr := range testruns {
		t.Run(
			fmt.Sprintf(testNameTemplate, tr.givenTemp.Unit, tr.givenTemp.Value, tr.givenSpeed.Unit, tr.givenSpeed.Value, tr.expected.Unit, tr.expected.Value),
			func(tt *testing.T) {
				actual := GetWindChillIndex(tr.givenTemp, tr.givenSpeed)
				assert.Equal(tt, tr.expected.Unit, actual.Unit)
				assert.InDelta(tt, tr.expected.Value, actual.Value, 1.0)
			},
		)
	}
}

func TestGetWindChillIndex_Units(t *testing.T) {
	testruns := []testrun{
		// kph
		// givenTemp, givenSpeed, expectedTemp
		{temp.C(-15), speed.Kph(35), temp.C(-27)},
		{temp.C(-15).Fahrenheit(), speed.Kph(35), temp.C(-27).Fahrenheit()},
		{temp.C(-15).Kelvin(), speed.Kph(35), temp.C(-27).Kelvin()},

		// mph
		// givenTemp, givenSpeed, expectedTemp
		{temp.C(-15), speed.Kph(35).Mph(), temp.C(-27)},
		{temp.C(-15).Fahrenheit(), speed.Kph(35).Mph(), temp.C(-27).Fahrenheit()},
		{temp.C(-15).Kelvin(), speed.Kph(35).Mph(), temp.C(-27).Kelvin()},
	}

	for _, tr := range testruns {
		t.Run(
			fmt.Sprintf(testNameTemplate, tr.givenTemp.Unit, tr.givenTemp.Value, tr.givenSpeed.Unit, tr.givenSpeed.Value, tr.expected.Unit, tr.expected.Value),
			func(tt *testing.T) {
				actual := GetWindChillIndex(tr.givenTemp, tr.givenSpeed)
				assert.Equal(tt, tr.expected.Unit, actual.Unit)
				assert.InDelta(tt, tr.expected.Value, actual.Value, 1.0)
			},
		)
	}
}
