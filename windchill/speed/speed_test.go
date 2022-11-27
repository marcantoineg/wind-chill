package speed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMPH(t *testing.T) {
	tests := []struct {
		name    string
		given   Speed
		expects Speed
	}{
		{name: "-30 kph to mph", given: Speed{-30, KPH}, expects: Speed{18.6411, MPH}},
		{name: "0 kph to mph", given: Speed{0, KPH}, expects: Speed{0, MPH}},
		{name: "1 kph to mph", given: Speed{1, KPH}, expects: Speed{0.621371, MPH}},
		{name: "30 kph to mph", given: Speed{30, KPH}, expects: Speed{18.6411, MPH}},
		{name: "50 kph to mph", given: Speed{50, KPH}, expects: Speed{31.0686, MPH}},
		{name: "100 kph to mph", given: Speed{100, KPH}, expects: Speed{62.1371, MPH}},
		{name: "250 kph to mph", given: Speed{250, KPH}, expects: Speed{155.343, MPH}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.given.Mph()
			assert.InDelta(t, tt.expects.Value, actual.Value, 0.001)
			assert.Equal(t, tt.expects.Unit, actual.Unit)
		})
	}
}

func TestKPH(t *testing.T) {
	tests := []struct {
		name    string
		given   Speed
		expects Speed
	}{
		{name: "-25 mph to kph", given: Speed{-25, MPH}, expects: Speed{40.2336, KPH}},
		{name: "0 mph to kph", given: Speed{0, MPH}, expects: Speed{0, KPH}},
		{name: "1 mph to kph", given: Speed{1, MPH}, expects: Speed{1.60934, KPH}},
		{name: "25 mph to kph", given: Speed{25, MPH}, expects: Speed{40.2336, KPH}},
		{name: "55 mph to kph", given: Speed{55, MPH}, expects: Speed{88.5139, KPH}},
		{name: "70 mph to kph", given: Speed{70, MPH}, expects: Speed{112.654, KPH}},
		{name: "150 mph to kph", given: Speed{150, MPH}, expects: Speed{241.402, KPH}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.given.Kph()
			assert.InDelta(t, tt.expects.Value, actual.Value, 0.001)
			assert.Equal(t, tt.expects.Unit, actual.Unit)
		})
	}
}
