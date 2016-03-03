package main
import (
    "fmt"
    "math"
)
func Sqrt( x float64) float64{
    var z float64 = 1
    for math.Abs(z*z -x) > 1e-15{
        z = z - (z * z - x) / (2 * z)
    }
    return z
}

func main(){
    fmt.Println(Sqrt(2))
}
