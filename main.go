package main

import (
	"fmt"
	"math"
	"log"
	"strconv"
)

type rgb struct { r, g, b uint8 }
type hsl struct { h, s, l float64 }

func max3(a, b, c float64) float64 {
	return math.Max(math.Max(a, b), c)
}

func min3(a, b, c float64) float64 {
	return math.Min(math.Min(a, b), c)
}

func (color rgb) toHsl() hsl {
	r := float64(color.r) / 255
	g := float64(color.g) / 255
	b := float64(color.b) / 255

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


	h = math.Round(h * 10) / 10
	l = math.Round(l * 1000) / 10
	s = math.Round(s * 1000) / 10

	return hsl{h, s, l}
}

func (color rgb) toHex() string {
	return fmt.Sprintf("#%X%X%X", color.r, color.g, color.b)
}

func (color hsl) toRgb() rgb {
	h := color.h
	s := color.s / 100
	l := color.l / 100
	a := s * math.Min(l, 1 - l)

	f := func(n float64) uint8 {
		k := math.Mod(n + h / 30, 12)
		val := l - a * math.Max(-1, min3(k - 3, 9 - k, 1))
		val = math.Round(val * 255)
		return uint8(val)
	}

	return rgb{f(0), f(8), f(4)}
}

func hexToRgb(hex string) rgb {
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

	return rgb{r, g, b}
}

func main() {
	hsl := hexToRgb("#161821").toHsl()
	for i := 0; i < 6; i += 1 {
		hsl.l += 4
		rgb := hsl.toRgb()
		fmt.Println("https://colorhexa.com/" + rgb.toHex()[1:])
	}
}
