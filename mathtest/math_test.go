package mathtest

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	area := 37
	t.Logf("%f", math.Sqrt(float64(area)))
	start := int(math.Sqrt(float64(area)))
	for start >= 0 {
		if area%start == 0 {
			t.Logf("%d, %d", start, area/start)
			break
		}
		start -= 1
	}
}
