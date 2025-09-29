package path

import (
	"github.com/pedro-makoski/go-utils/file"
	"github.com/pedro-makoski/go-utils/folder"
)

func (p Path) ThisPathExists() bool {
	if p.IsFile() {
		return file.DoesThisFileExists(p.Path)
	}

	return folder.DoesThisFolderExists(p.Path)
}
