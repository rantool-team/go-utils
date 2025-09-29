package path

import (
	errorslist "github.com/pedro-makoski/go-utils/errors-list"
	"github.com/pedro-makoski/go-utils/file"
	"github.com/pedro-makoski/go-utils/folder"

	goerror "github.com/rantool-team/go-error"
)

func (p Path) Read() (string, goerror.Error) {
	if p.IsFile() {
		return p.read()
	}

	return "", errorslist.ErrorsList.FolderFile.ItMustBEAFile(p.Path)
}

func (p Path) read() (string, goerror.Error) {
	return file.ReadFile(p.Path)
}

func (p Path) GetInsidePaths() ([]string, goerror.Error) {
	if p.IsFolder() {
		return folder.GetAllPaths(p.Path)
	}

	return nil, errorslist.ErrorsList.FolderFile.ItMustBeAFolder(p.Path)
}

func (p Path) CachedGetInsidePaths() ([]string, goerror.Error) {
	if p.IsFolder() {
		return folder.CachedGetAllPaths(p.Path)
	}

	return nil, errorslist.ErrorsList.FolderFile.ItMustBeAFolder(p.Path)
}

func (p Path) CachedGetInsidePathsWithExt(ext string) ([]string, goerror.Error) {
	if p.IsFolder() {
		return folder.CachedGetAllPathsWithExt(ext, p.Path)
	}

	return nil, errorslist.ErrorsList.FolderFile.ItMustBeAFolder(p.Path)
}

func (p Path) CachedGetAllFolders() ([]string, goerror.Error) {
	if p.IsFolder() {
		return folder.CachedGetAllFolders(p.Path)
	}

	return nil, errorslist.ErrorsList.FolderFile.ItMustBeAFolder(p.Path)
}

func (p Path) CachedGetAllFiles() ([]string, goerror.Error) {
	if p.IsFolder() {
		return folder.CachedGetAllFiles(p.Path)
	}

	return nil, errorslist.ErrorsList.FolderFile.ItMustBeAFolder(p.Path)
}
