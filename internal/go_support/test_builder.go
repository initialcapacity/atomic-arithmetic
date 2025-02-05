package go_support

import (
	"fmt"
	"os"
	"strings"
)

type TestBuilder struct {
	stringBuilder strings.Builder
}

func NewTestBuilder(packageName string) *TestBuilder {
	b := &TestBuilder{}
	b.stringBuilder.WriteString(fmt.Sprintf(`package %s_test

import (
	"github.com/initialcapacity/atomic-arithmetic/pkg/%s"
	"testing"
)
`, packageName, packageName))
	return b
}

func (b *TestBuilder) AppendTestFunction(functionPrefix string, n, input, result int) {
	b.stringBuilder.WriteString(fmt.Sprintf(`
func Test%s%d(t *testing.T) {
	result := operations.%s%d(%d)
	
	if result != %d {
		t.Errorf("Expected %d but got %%d", result)
	}
}
`, functionPrefix, n, functionPrefix, n, input, result, result))
}

func (b *TestBuilder) WriteFile(path string) error {
	return os.WriteFile(path, []byte(b.stringBuilder.String()), 0644)
}
