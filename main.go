package main

import (
	"fmt"
	"math"
)

func calcHue(r, g, b int) float64 {
	R := float64(r) / 255
	G := float64(g) / 255
	B := float64(b) / 255

	max := max3(R, G, B)
	min := min3(R, G, B)

	var hue float64
	if R == max {
		hue = (G - B) / (max - min)
	} else if G == max {
		hue = 2.0 + (B - R) / (max - min)
	} else {
		hue = 4.0 + (R - G) / (max - min)
	}

	hue *= 60

	if hue < 0 {
		hue += 360
	}
	
	return math.Round(hue * 10) / 10
}

func max3(a, b, c float64) float64 {
	return math.Max(math.Max(a, b), c)
}

func min3(a, b, c float64) float64 {
	return math.Min(math.Min(a, b), c)
}

func main() {
	fmt.Println("hello, world")
}
