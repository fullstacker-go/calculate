package main

import (
	"fmt"

	"github.com/fullstacker-go/calculate/calculation"
)

func main() {
	result := calculation.Add(23, 33)
	fmt.Println(result)
}
