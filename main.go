/*
Copyright © 2024 Rob Duarte <me@robduarte.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// gradescale
// takes a maximum number of points and produces a letter grade scale
// the new scale is based on the FSU grading scale but is proportional to the new max
// while the original points are fractional, the new scale are rounded integers

type grade struct {
	Letter string
	Min    float64
}

var scale = []grade{
	{"A", 93.5},
	{"A-", 89.5},
	{"B+", 86.5},
	{"B", 82.5},
	{"B-", 79.5},
	{"C+", 76.5},
	{"C", 72.5},
	{"C-", 69.5},
	{"D+", 66.5},
	{"D", 62.5},
	{"D-", 59.5},
	{"F", 0},
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: gradescale MAXPOINTS")
		os.Exit(1)
	}

	max, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Argument must be a whole number")
		os.Exit(1)
	}

	fmt.Println("| Letter | Low | High |")
	fmt.Println("| :----- | :-- | :--- |")

	for i, g := range scale {

		thismin := math.Round((float64(max) * g.Min) / 100)

		if i == 0 {
			fmt.Printf("| %s | ≥ %.0f | ≤ %.0f |\n", g.Letter, thismin, max)
			continue
		}

		thismax := math.Round((float64(max) * scale[i-1].Min) / 100)

		fmt.Printf("| %s | ≥ %.0f | < %.0f |\n", g.Letter, thismin, thismax)
	}
}
