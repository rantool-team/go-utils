package folder

import (
	"os"

	errorslist "github.com/rantool-team/go-utils/errors-list"

	goerror "github.com/rantool-team/go-error"
)

func RemoveFolder(path string) goerror.Error {
	err := os.RemoveAll(path)
	if err != nil {
		return errorslist.ErrorsList.Folder.ErrorOnRemoveFolder(path)
	}
	return goerror.CreateBlankError()
}
