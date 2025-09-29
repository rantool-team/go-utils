package path

import (
	"path/filepath"
	"runtime"
	"strings"
)

func isWindowsRoot(path string) bool {
	return len(path) == 3 && path[1:3] == ":\\"
}

func isUNCBase(path string) bool {
	if !strings.HasPrefix(path, `\\`) {
		return false
	}
	parts := strings.SplitN(path, `\`, 4)
	return len(parts) <= 3
}

func (p Path) HasParentPath() bool {
	clean := filepath.Clean(p.Path)

	if runtime.GOOS == "windows" {
		clean = strings.ReplaceAll(clean, "/", `\`)
		if isWindowsRoot(clean) || isUNCBase(clean) {
			return false
		}
	} else {
		if clean == "/" {
			return false
		}
	}

	return filepath.Dir(clean) != clean
}
