package main

import (
	"fmt"
	"math"

	"github.com/qmeehd/tutorial_go/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(4.2))
	fmt.Println(math.Ceil(6.9))
	fmt.Println(math.Sqrt(42.69))
	fmt.Println(strutil.Reverse("olleh"))
}
