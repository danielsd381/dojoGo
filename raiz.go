//Daniel Sierra D'

package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  //codigo que implementa la raiz
  //
   z := float64(x/2)
   contador := 0


     for delta := z; math.Abs(delta) > 1e-6; {
          zi := z
          z = z - ((z*z)-x)/(2*z)
          delta = z - zi
          contador++

     }
fmt.Println(contador)
return z
}

 
func main() {
  fmt.Println(Sqrt(10))
  fmt.Println(math.Sqrt(10))
}
