package path

import (
	errorslist "github.com/pedro-makoski/go-utils/errors-list"
	"github.com/pedro-makoski/go-utils/file"

	goerror "github.com/rantool-team/go-error"
)

func (p Path) WriteInFile(content string, permission file.FilePermissions) goerror.Error {
	if p.IsFile() {
		p.writeInFile(content, permission)
	}

	return errorslist.ErrorsList.FolderFile.ItMustBEAFile(p.Path)
}

func (p Path) writeInFile(content string, permission file.FilePermissions) goerror.Error {
	return file.WriteFile(p.Path, content, permission.ToFileMode())
}
