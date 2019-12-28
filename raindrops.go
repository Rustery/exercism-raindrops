// Exercism side exercise about Raindrops
package raindrops

import (
	"math"
	"strconv"
)

// Convert converts a number into a string that contains raindrop sounds corresponding to certain potential factors.
func Convert(i int) (result string) {

	if math.Mod(float64(i), 3) == 0 {
		result += "Pling"
	}

	if math.Mod(float64(i), 5) == 0 {
		result += "Plang"
	}

	if math.Mod(float64(i), 7) == 0 {
		result += "Plong"
	}

	if result == "" {
		result = strconv.Itoa(i)
	}

	return
}
