package mean
// Calculate the average of a given slice
func Compute(x []float64) float64 {
  if len(x) < 1 {
    return float64(0)
  }
  sum := float64(0)
  for _, i := range x {
    sum += i
  }
  return sum / float64(len(x))
}
