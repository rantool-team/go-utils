package folderfile

import (
	"os"

	errorslist "github.com/pedro-makoski/go-utils/errors-list"
	"github.com/pedro-makoski/go-utils/file"
	"github.com/pedro-makoski/go-utils/folder"

	goerror "github.com/rantool-team/go-error"
)

func Copy(srcPath, dstPath string) goerror.Error {
	info, osErr := os.Stat(srcPath)
	if osErr != nil {
		return errorslist.ErrorsList.FolderFile.ErrorOnGetInfo(srcPath)
	}

	err := ensureDestinationDirExists(dstPath)
	if err.HasError() {
		return err
	}

	if info.IsDir() {
		return folder.CopyFolder(srcPath, dstPath)
	}

	return file.CopyFile(srcPath, dstPath)
}
