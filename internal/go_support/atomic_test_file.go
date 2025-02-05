package go_support

import (
	"log"
	"path/filepath"
)

func CreateAtomicTest(directory, packageName, fileName, functionPrefix string, resultCalculator func(int, int) int, count int) {
	addBuilder := NewTestBuilder(packageName)

	for i := 1; i <= count; i++ {
		addBuilder.AppendTestFunction(functionPrefix, i, 17_000, resultCalculator(17_000, i))
	}

	err := addBuilder.WriteFile(filepath.Join(directory, fileName))
	if err != nil {
		log.Fatalf("Unable to create %s/%s: %v\n", packageName, fileName, err)
	}
}
