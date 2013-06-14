package main

import "fmt"
import "mean"
import "stdev"

func main() {
  x := make([]float64, 10000)
  for i, _ := range x {
    x[i] = float64(i)
  }

  //fmt.Println("List", x)
  fmt.Println("Avg:", mean.Compute(x))
  fmt.Println("Stdev:", stdev.Compute(x))
}
