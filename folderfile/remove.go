package folderfile

import (
	"os"

	errorslist "github.com/rantool-team/go-utils/errors-list"
	"github.com/rantool-team/go-utils/file"
	"github.com/rantool-team/go-utils/folder"

	goerror "github.com/rantool-team/go-error"
)

func Remove(path string) goerror.Error {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return goerror.CreateBlankError()
	}

	if err != nil {
		return errorslist.ErrorsList.FolderFile.ErrorOnGetInfo(path)
	}

	if info.IsDir() {
		return folder.RemoveFolder(path)
	}

	return file.RemoveFile(path)
}
