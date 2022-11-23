package speed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMph(t *testing.T) {
	tests := []struct {
		name    string
		given   Speed
		expects Speed
	}{
		{name: "0 kph to mph", given: Speed{0, Kph}, expects: Speed{0, Mph}},
		{name: "1 kph to mph", given: Speed{1, Kph}, expects: Speed{0.621371, Mph}},
		{name: "30 kph to mph", given: Speed{30, Kph}, expects: Speed{18.6411, Mph}},
		{name: "50 kph to mph", given: Speed{50, Kph}, expects: Speed{31.0686, Mph}},
		{name: "100 kph to mph", given: Speed{100, Kph}, expects: Speed{62.1371, Mph}},
		{name: "250 kph to mph", given: Speed{250, Kph}, expects: Speed{155.343, Mph}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.given.Mph()
			assert.InDelta(t, actual.value, tt.expects.value, 0.001)
			assert.Equal(t, actual.unit, tt.expects.unit)
		})
	}
}

func TestKph(t *testing.T) {
	tests := []struct {
		name    string
		given   Speed
		expects Speed
	}{
		{name: "0 mph to kph", given: Speed{0, Mph}, expects: Speed{0, Kph}},
		{name: "1 mph to kph", given: Speed{1, Mph}, expects: Speed{1.60934, Kph}},
		{name: "25 mph to kph", given: Speed{25, Mph}, expects: Speed{40.2336, Kph}},
		{name: "55 mph to kph", given: Speed{55, Mph}, expects: Speed{88.5139, Kph}},
		{name: "70 mph to kph", given: Speed{70, Mph}, expects: Speed{112.654, Kph}},
		{name: "150 mph to kph", given: Speed{150, Mph}, expects: Speed{241.402, Kph}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.given.Kph()
			assert.InDelta(t, actual.value, tt.expects.value, 0.001)
			assert.Equal(t, actual.unit, tt.expects.unit)
		})
	}
}
