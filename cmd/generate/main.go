package main

import (
	"github.com/initialcapacity/atomic-arithmetic/internal/go_support"
	"log"
)

func main() {
	operationsDir, err := go_support.CreatePackageDirectory("operations")
	if err != nil {
		log.Fatalf("Unable to create pkg/operations directory: %v\n", err)
	}

	maximumNumber := 10_000

	go_support.CreateAtomicFile(operationsDir, "operations", "add.go", "Add", "+", maximumNumber)
	go_support.CreateAtomicTest(operationsDir, "operations", "add_test.go", "Add", func(i, j int) int { return i + j }, maximumNumber)

	go_support.CreateAtomicFile(operationsDir, "operations", "subtract.go", "Subtract", "-", maximumNumber)
	go_support.CreateAtomicTest(operationsDir, "operations", "subtract_test.go", "Subtract", func(i, j int) int { return i - j }, maximumNumber)

	go_support.CreateAtomicFile(operationsDir, "operations", "multiply.go", "MultiplyBy", "*", maximumNumber)
	go_support.CreateAtomicTest(operationsDir, "operations", "multiply_test.go", "MultiplyBy", func(i, j int) int { return i * j }, maximumNumber)

	go_support.CreateAtomicFile(operationsDir, "operations", "divide.go", "DivideBy", "/", maximumNumber)
	go_support.CreateAtomicTest(operationsDir, "operations", "divide_test.go", "DivideBy", func(i, j int) int { return i / j }, maximumNumber)

	go_support.CreateAtomicFile(operationsDir, "operations", "modulo.go", "Mod", "%", maximumNumber)
	go_support.CreateAtomicTest(operationsDir, "operations", "modulo_test.go", "Mod", func(i, j int) int { return i % j }, maximumNumber)

	log.Println("Success!")
}
