package windchill

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMph(t *testing.T) {
	tests := []struct {
		name    string
		given   speed
		expects speed
	}{
		{name: "0 kph to mph", given: speed{0, Kph}, expects: speed{0, Mph}},
		{name: "1 kph to mph", given: speed{1, Kph}, expects: speed{0.621371, Mph}},
		{name: "30 kph to mph", given: speed{30, Kph}, expects: speed{18.6411, Mph}},
		{name: "50 kph to mph", given: speed{50, Kph}, expects: speed{31.0686, Mph}},
		{name: "100 kph to mph", given: speed{100, Kph}, expects: speed{62.1371, Mph}},
		{name: "250 kph to mph", given: speed{250, Kph}, expects: speed{155.343, Mph}},
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
		given   speed
		expects speed
	}{
		{name: "0 mph to kph", given: speed{0, Mph}, expects: speed{0, Kph}},
		{name: "1 mph to kph", given: speed{1, Mph}, expects: speed{1.60934, Kph}},
		{name: "25 mph to kph", given: speed{25, Mph}, expects: speed{40.2336, Kph}},
		{name: "55 mph to kph", given: speed{55, Mph}, expects: speed{88.5139, Kph}},
		{name: "70 mph to kph", given: speed{70, Mph}, expects: speed{112.654, Kph}},
		{name: "150 mph to kph", given: speed{150, Mph}, expects: speed{241.402, Kph}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.given.Kph()
			assert.InDelta(t, actual.value, tt.expects.value, 0.001)
			assert.Equal(t, actual.unit, tt.expects.unit)
		})
	}
}
