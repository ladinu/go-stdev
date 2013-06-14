package stdev
import (
  "math"
  "mean"
)
// Return the standarad deviation of a slice
func Compute(x []float64) float64 {
  avg := mean.Compute(x)
  sum := float64(0)

  for _, i := range x {
    sum += math.Pow(i - avg, 2.0)
  }

  sum = sum / float64(len(x) - 1)
  return math.Sqrt(sum)
}
