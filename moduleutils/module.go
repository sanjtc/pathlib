package moduleutils

import (
	"os"
	"strings"

	"github.com/pantskun/pathlib/pathutils"
)

// GetModulePath
// get current module path.
func GetModulePath(moduleName string) string {
	fp, _ := os.Getwd()
	fp = pathutils.ConvertBackslashToSlash(fp)
	fp = strings.SplitAfter(fp, moduleName)[0]

	return fp
}
