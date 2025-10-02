package file

import (
	"os"

	errorslist "github.com/rantool-team/go-utils/errors-list"

	goerror "github.com/rantool-team/go-error"
)

func RenameFile(oldPath, newPath string) goerror.Error {
	err := os.Rename(oldPath, newPath)

	if err != nil {
		return errorslist.ErrorsList.File.RenameFile(oldPath, newPath)
	}

	return goerror.CreateBlankError()
}
