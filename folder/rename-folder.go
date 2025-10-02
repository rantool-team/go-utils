package folder

import (
	"os"

	errorslist "github.com/rantool-team/go-utils/errors-list"

	goerror "github.com/rantool-team/go-error"
)

func RenameFolder(oldPath, newPath string) goerror.Error {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		return errorslist.ErrorsList.Folder.ErrorOnRenameFolder(oldPath, newPath)
	}
	return goerror.CreateBlankError()
}
