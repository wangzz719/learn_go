package main

import (
    "fmt"
    "golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
    var pic = make([][]uint8, dx)
    for x:=0; x < dx; x++{
        var picy = make([]uint8, dy)
        for y:=0; y < dy; y++ {
            picy[y] = uint8((x+y)/2)
        }
        pic[x] = picy
    }
    return pic
}

func main() {
	pic.Show(Pic)
}
