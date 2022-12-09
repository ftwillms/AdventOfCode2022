package other

import (
	"fmt"
	"math"
)

func WhatIsImage() {
	test := float64(5 - 1)
	test2 := float64(3)
	// s copy sign copies the sign from the sign value to the f value
	fmt.Println(math.Copysign(test2, test))
}
