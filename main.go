package main

import (
	"fmt"
	"math"
	"log"
	"strconv"
)

type RGB [3]uint8
type HSL [3]float64

func rgbToHsl(c RGB) HSL {
	r := float64(c[0]) / 255
	g := float64(c[1]) / 255
	b := float64(c[2]) / 255

	max := max3(r, g, b)
	min := min3(r, g, b)

	l := (max + min) / 2

	var s float64
	if l == 0 || l == 1 {
		s = 0
	} else {
		s = (max - min) / (1 - math.Abs(2 * l - 1))
	}

	var h float64
	switch {
	case max - min == 0:
		h = 0
	case r == max:
		h = math.Mod((g - b) / (max - min), 6)
	case g == max:
		h = 2.0 + (b - r) / (max - min)
	case b == max:
		h = 4.0 + (r - g) / (max - min)
	}

	h *= 60

	if h < 0 {
		h += 360
	}

	h = math.Floor(h * 10) / 10
	l = math.Floor(l * 1000) / 10
	s = math.Floor(s * 1000) / 10

	return HSL{h, s, l}
}

func max3(a, b, c float64) float64 {
	return math.Max(math.Max(a, b), c)
}

func min3(a, b, c float64) float64 {
	return math.Min(math.Min(a, b), c)
}

func hexToRgb(hex string) RGB {
	hex = hex[1:]

	if len(hex) != 6 {
		log.Fatal("Invalid hex")
	}

	toDecimal := func(val string) uint8 {
		v, err := strconv.ParseInt(val, 16, 0)
		if err != nil {
			log.Fatal("Invalid hex")
		}

		return uint8(v)
	}

	r := toDecimal(hex[0:2])
	g := toDecimal(hex[2:4])
	b := toDecimal(hex[4:6])

	return RGB{r,g,b}
}

func main() {
	fmt.Println(rgbToHsl(hexToRgb("#FFFFFF")))
	fmt.Println(rgbToHsl(hexToRgb("#808080")))
	fmt.Println(rgbToHsl(hexToRgb("#000000")))
	fmt.Println(rgbToHsl(hexToRgb("#FF0000")))
	fmt.Println(rgbToHsl(hexToRgb("#BFBF00")))
	fmt.Println(rgbToHsl(hexToRgb("#008000")))
	fmt.Println(rgbToHsl(hexToRgb("#80FFFF")))
	fmt.Println(rgbToHsl(hexToRgb("#8080FF")))
	fmt.Println(rgbToHsl(hexToRgb("#BF40BF")))
	fmt.Println(rgbToHsl(hexToRgb("#A0A424")))
	fmt.Println(rgbToHsl(hexToRgb("#362698")))
	fmt.Println(rgbToHsl(hexToRgb("#7E7EB8")))
}
