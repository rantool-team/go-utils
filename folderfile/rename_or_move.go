package folderfile

import (
	"os"

	errorslist "github.com/pedro-makoski/go-utils/errors-list"
	"github.com/pedro-makoski/go-utils/file"
	"github.com/pedro-makoski/go-utils/folder"

	goerror "github.com/rantool-team/go-error"
)

func RenameOrMove(oldPath, newPath string) goerror.Error {
	err := ensureDestinationDirExists(newPath)
	if err.HasError() {
		return err
	}

	info, osErr := os.Stat(oldPath)
	if osErr != nil {
		return errorslist.ErrorsList.FolderFile.ErrorOnGetInfo(oldPath)
	}

	if info.IsDir() {
		return folder.RenameFolder(oldPath, newPath)
	}

	return file.RenameFile(oldPath, newPath)
}
