package main

import (
	"fmt"
	"github.com/initialcapacity/atomic-arithmetic/pkg/operations"
)

func main() {
	fmt.Println("\nAdd 365 to 17")
	fmt.Println(operations.Add365(17))

	fmt.Println("\nSubtract 714 from 893")
	fmt.Println(operations.Subtract714(893))

	fmt.Println("\nMultiple 87 by 199")
	fmt.Println(operations.MultiplyBy199(87))

	fmt.Println("\nDivide 744 by 342")
	fmt.Println(operations.DivideBy342(744))

	fmt.Println("\nCalculate 654 modulo 17")
	fmt.Println(operations.Mod17(654))
}
