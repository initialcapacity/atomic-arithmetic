package go_support

import (
	"log"
	"path/filepath"
)

func CreateAtomicFile(directory, packageName, fileName, functionPrefix, operator string, count int) {
	builder := NewFileBuilder(packageName)

	for i := 1; i <= count; i++ {
		builder.AppendFunction(functionPrefix, operator, i)
	}

	err := builder.WriteFile(filepath.Join(directory, fileName))
	if err != nil {
		log.Fatalf("Unable to create %s/%s: %v\n", packageName, fileName, err)
	}
}
