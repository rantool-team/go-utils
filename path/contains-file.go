package path

import (
	"path/filepath"

	"github.com/pedro-makoski/go-utils/array"

	goerror "github.com/rantool-team/go-error"
)

func (p Path) ContainsFile(fileName string) (bool, goerror.Error) {
	paths, err := p.CachedGetInsidePathsWithExt(filepath.Ext(fileName))
	if err.HasError() {
		return false, err
	}

	return array.Contains(paths, fileName), goerror.CreateBlankError()
}

func (p Path) ContainsFileName(fileName string) (bool, goerror.Error) {
	paths, err := p.CachedGetInsidePathsWithExt(filepath.Ext(fileName))
	if err.HasError() {
		return false, err
	}

	newPaths := GenerateFileNameListFromAbsPaths(paths)

	return array.Contains(newPaths, fileName), goerror.CreateBlankError()
}

func (p Path) ContainsFolder(folderName string) (bool, goerror.Error) {
	paths, err := p.CachedGetAllFolders()
	if err.HasError() {
		return false, err
	}

	newPaths := GenerateFileNameListFromAbsPaths(paths)

	return array.Contains(newPaths, folderName), goerror.CreateBlankError()
}

func GenerateFileNameListFromAbsPaths(paths []string) []string {
	newPaths := []string{}
	for _, path := range paths {
		newPaths = append(newPaths, filepath.Base(path))
	}

	return newPaths
}
