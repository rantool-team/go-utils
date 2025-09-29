package folderfile

import (
	"path/filepath"

	"github.com/pedro-makoski/go-utils/folder"

	goerror "github.com/rantool-team/go-error"
)

func ensureDestinationDirExists(destinationPath string) goerror.Error {
	destDir := filepath.Dir(destinationPath)
	if !folder.DoesThisFolderExists(destDir) {
		err := folder.CreateFolder(destDir)
		if err.HasError() {
			return err
		}
	}
	return goerror.CreateBlankError()
}
