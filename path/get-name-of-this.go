package path

import "path/filepath"

func (p Path) GetNameOfThis() string {
	name := filepath.Base(p.Path)
	return name
}
