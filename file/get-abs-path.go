package file

import (
	"path/filepath"

	errorslist "github.com/rantool-team/go-utils/errors-list"

	goerror "github.com/rantool-team/go-error"
)

func GetAbs(pathRelative string) (string, goerror.Error) {
	pathAbs, err := filepath.Abs(pathRelative)

	if err != nil {
		return "", errorslist.ErrorsList.File.GetAbs(pathRelative)
	}

	return pathAbs, goerror.CreateBlankError()
}
