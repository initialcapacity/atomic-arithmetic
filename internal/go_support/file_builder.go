package go_support

import (
	"fmt"
	"os"
	"strings"
)

type FileBuilder struct {
	stringBuilder strings.Builder
}

func NewFileBuilder(packageName string) *FileBuilder {
	b := &FileBuilder{}
	b.stringBuilder.WriteString(fmt.Sprintf(`package %s
`, packageName))
	return b
}

func (b *FileBuilder) AppendFunction(functionPrefix, operator string, n int) {
	b.stringBuilder.WriteString(fmt.Sprintf(`
func %s%d(operand int) int {
	return operand %s %d
}
`, functionPrefix, n, operator, n))
}

func (b *FileBuilder) WriteFile(path string) error {
	return os.WriteFile(path, []byte(b.stringBuilder.String()), 0644)
}
