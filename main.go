//main package declaration

package main

import (
	"fmt"

	"github.com/fullstacker-go/calculate/calculation" //importing custom package calculation from its path
)

func main() {
	result := calculation.Add(23, 33)
	fmt.Println(result)
}
