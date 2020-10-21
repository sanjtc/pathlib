package pathlib

import (
	"os"
	"strings"
)

// GetModulePath
// get current module path.
func GetModulePath(moduleName string) string {
	fp, _ := os.Getwd()
	fp = ConvertBackslashToSlash(fp)
	fp = strings.SplitAfter(fp, moduleName)[0]

	return fp
}
