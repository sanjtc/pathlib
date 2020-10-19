package module

import (
	"os"
	"strings"

	"github.com/pantskun/pathlib/path"
)

// GetModulePath
// get current module path.
func GetModulePath(moduleName string) string {
	fp, _ := os.Getwd()
	fp = path.ConvertBackslashToSlash(fp)
	fp = strings.SplitAfter(fp, moduleName)[0]

	return fp
}
