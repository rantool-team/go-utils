package file

import (
	"os"

	errorslist "github.com/rantool-team/go-utils/errors-list"

	goerror "github.com/rantool-team/go-error"
)

func WriteFile(path string, newContent string, permission os.FileMode) goerror.Error {
	err := os.WriteFile(path, []byte(newContent), permission)

	if err != nil {
		return errorslist.ErrorsList.File.WriteFile(path)
	}

	return goerror.CreateBlankError()
}

func AppendInFile(path string, newLine string) goerror.Error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errorslist.ErrorsList.File.AppendInFile(path)
	}
	defer f.Close()

	if _, err := f.WriteString(newLine + "\n"); err != nil {
		return errorslist.ErrorsList.File.AppendInFile(path)
	}

	return goerror.CreateBlankError()
}
