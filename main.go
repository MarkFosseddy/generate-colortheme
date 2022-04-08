package main

import (
	"fmt"
	"math"
	"log"
	"strconv"
)

func calcHue(r, g, b uint8) float64 {
	R := float64(r) / 255
	G := float64(g) / 255
	B := float64(b) / 255

	max := max3(R, G, B)
	min := min3(R, G, B)

	if max - min == 0 {
		return 0
	}

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

func hexToRgb(hex string) (uint8, uint8, uint8) {
	hex = hex[1:]

	if len(hex) != 6 {
		log.Fatal("Invalid hex")
	}

	r, err := strconv.ParseInt(hex[0:2], 16, 0)
	if err != nil {
		log.Fatal("Invalid hex")
	}

	g, err := strconv.ParseInt(hex[2:4], 16, 0)
	if err != nil {
		log.Fatal("Invalid hex")
	}

	b, err := strconv.ParseInt(hex[4:6], 16, 0)
	if err != nil {
		log.Fatal("Invalid hex")
	}

	return uint8(r), uint8(g), uint8(b)
}

func main() {
	r, g, b := hexToRgb("#161821")
	hue := calcHue(r, g, b)

	fmt.Println(r, g, b, hue)
}
