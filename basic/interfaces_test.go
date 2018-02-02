package basic

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMeasure(t *testing.T) {
	r := rect{width:3, height:4}
	c := circle{radius:2}

	rArea, rPerim := Measure(r)
	cArea, cPerim := Measure(c)

	assert.Equal(t, rArea, float64(12))
	assert.Equal(t, rPerim, float64(14))

	assert.Equal(t, cArea, float64(12.566370614359172))
	assert.Equal(t, cPerim, float64(12.566370614359172))
}