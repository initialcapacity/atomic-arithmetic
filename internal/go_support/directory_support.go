package go_support

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreatePackageDirectory(packageName string) (string, error) {
	pkgDir := filepath.Join(".", "pkg")
	err := os.RemoveAll(pkgDir)
	if err != nil {
		return "", fmt.Errorf("failed to remove pkg directory: %v", err)
	}
	operationsDir := filepath.Join(pkgDir, packageName)
	err = os.MkdirAll(operationsDir, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create %s directory: %v\n", packageName, err)
	}
	return operationsDir, nil
}
