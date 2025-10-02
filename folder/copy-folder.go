package folder

import (
	"os"
	"path/filepath"

	errorslist "github.com/rantool-team/go-utils/errors-list"
	"github.com/rantool-team/go-utils/file"

	goerror "github.com/rantool-team/go-error"
)

func CopyFolder(srcRoot, dstRoot string) goerror.Error {
	allPaths, err := GetAllPathsRecursive(srcRoot)
	if err.HasError() {
		return errorslist.ErrorsList.Folder.ErrorOnGetAllPaths(srcRoot)
	}

	for _, path := range allPaths {
		if err := processPath(path, srcRoot, dstRoot); err.HasError() {
			return err
		}
	}

	return goerror.CreateBlankError()
}

func processPath(srcPath, srcRoot, dstRoot string) goerror.Error {
	relativePath, err := filepath.Rel(srcRoot, srcPath)
	if err != nil {
		return errorslist.ErrorsList.Folder.CopyFolder.ErrorOnGetRelativePath(srcRoot, dstRoot)
	}
	dstPath := filepath.Join(dstRoot, relativePath)

	srcInfo, err := os.Stat(srcPath)
	if err != nil {
		return errorslist.ErrorsList.Folder.CopyFolder.ErrorOnGetFileInfo(srcPath)
	}

	if srcInfo.IsDir() {
		return createDestinationDirectory(dstPath, srcInfo.Mode())
	}

	return copySourceFile(srcPath, dstPath)
}

func createDestinationDirectory(path string, mode os.FileMode) goerror.Error {
	if err := os.MkdirAll(path, mode); err != nil {
		return errorslist.ErrorsList.Folder.CopyFolder.ErrorOnCreateFolder(path)
	}
	return goerror.CreateBlankError()
}

func copySourceFile(src, dst string) goerror.Error {
	err := file.CopyFile(src, dst)
	if err.HasError() {
		return err
	}

	return goerror.CreateBlankError()
}
