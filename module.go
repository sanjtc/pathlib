package pathlib

import (
	"os"
	"strings"

	"github.com/pantskun/pathlib"
)

// GetModulePath
// get current module path.
func GetModulePath(moduleName string) string {
	fp, _ := os.Getwd()
	fp = pathlib.ConvertBackslashToSlash(fp)
	fp = strings.SplitAfter(fp, moduleName)[0]

	return fp
}
