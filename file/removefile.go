package file

import (
	"os"

	errorslist "github.com/pedro-makoski/go-utils/errors-list"

	goerror "github.com/rantool-team/go-error"
)

func RemoveFile(path string) goerror.Error {
	err := os.Remove(path)
	if err != nil {
		return errorslist.ErrorsList.File.RemoveFile(path)
	}
	return goerror.CreateBlankError()
}
