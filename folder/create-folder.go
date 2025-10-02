package folder

import (
	"os"

	errorslist "github.com/rantool-team/go-utils/errors-list"

	goerror "github.com/rantool-team/go-error"
)

func CreateFolder(path string) goerror.Error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return errorslist.ErrorsList.Folder.ErrorOnCreateFolder(path)
	}
	return goerror.CreateBlankError()
}
