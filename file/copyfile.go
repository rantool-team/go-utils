package file

import (
	"io"
	"os"

	errorslist "github.com/pedro-makoski/go-utils/errors-list"

	goerror "github.com/rantool-team/go-error"
)

func CopyFile(src, dst string) goerror.Error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return errorslist.ErrorsList.File.CopyErrors.ErrorInOpenFileOrigin(src)
	}

	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return errorslist.ErrorsList.File.CopyErrors.ErrorInOpenDestinationFile(dst)
	}

	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return errorslist.ErrorsList.File.CopyErrors.ErrorInCopyContent(src, dst)
	}

	err = destFile.Sync()
	if err != nil {
		return errorslist.ErrorsList.File.CopyErrors.ErrorInSyncFile(dst)
	}

	return goerror.CreateBlankError()
}
