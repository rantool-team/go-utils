package path

import (
	"github.com/rantool-team/go-utils/file"
	"github.com/rantool-team/go-utils/folder"
)

func (p Path) ThisPathExists() bool {
	if p.IsFile() {
		return file.DoesThisFileExists(p.Path)
	}

	return folder.DoesThisFolderExists(p.Path)
}
