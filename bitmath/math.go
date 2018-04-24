package bitmath

import (
	"math"
)

// TwoPi 2 pi
const TwoPi = math.Pi * 2

// HalfPi pi / 2
const HalfPi = math.Pi / 2

// Norm returns a normalized value in a min/max range.
func Norm(value float64, min float64, max float64) float64 {
	return (value - min) / (max - min)
}

// Lerp is linear interpolation within a min/max range.
func Lerp(t float64, min float64, max float64) float64 {
	return min + (max-min)*t
}

// MapTo maps a value within one min/max range to a value within another range.
func MapTo(srcValue float64, srcMin float64, srcMax float64, dstMin float64, dstMax float64) float64 {
	norm := Norm(srcValue, srcMin, srcMax)
	return Lerp(norm, dstMin, dstMax)
}

// Wrap wraps a value around so it remains between min and max.
func Wrap(value float64, min float64, max float64) float64 {
	r := max - min
	return min + math.Mod((math.Mod(value-min, r)+r), r)
}

// Clamp enforces a value does not go beyond a min/max range.
func Clamp(value float64, min float64, max float64) float64 {
	// let min and max be reversed and still work.
	realMin := min
	realMax := max
	if min > max {
		realMin = max
		realMax = min
	}
	result := value
	if value < realMin {
		result = realMin
	}
	if value > realMax {
		result = realMax
	}
	return result
}

// SinRange returns the sin of an angle mapped to a min/max range.
func SinRange(angle float64, min float64, max float64) float64 {
	return MapTo(math.Sin(angle), -1, 1, min, max)
}

// CosRange returns the cos of an angle mapped to a min/max range.
func CosRange(angle float64, min float64, max float64) float64 {
	return MapTo(math.Cos(angle), -1, 1, min, max)
}

// Equalish returns whether the two values are approximately equal.
func Equalish(a float64, b float64, delta float64) bool {
	return math.Abs(a-b) <= delta
}
