package folder

import (
	"os"
	"path/filepath"

	"github.com/pedro-makoski/go-utils/cache"
	errorslist "github.com/pedro-makoski/go-utils/errors-list"
	"github.com/pedro-makoski/go-utils/file"

	goerror "github.com/rantool-team/go-error"
)

const GET_ALL_PATHS_REGISTER_NAME = "folder.get-all-paths"
const GET_ALL_PATHS_WITH_EXT_REGISTER_NAME = "folder.get-all-paths-with-ext-register-name"
const GET_ALL_FOLDERS_REGISTER_NAME = "folder.get-all-folders"
const GET_ALL_FILES_REGISTER_NAME = "folder.get-all-files"

func GetAllPaths(root string) ([]string, goerror.Error) {
	return GetAllPathsFromCondition(func(path string) bool {
		return true
	}, root)
}

func GetAllPathsWithExt(ext string, root string) ([]string, goerror.Error) {
	return GetAllPathsFromCondition(func(path string) bool {
		return filepath.Ext(path) == ext || filepath.Ext(path) == ""
	}, root)
}

func GetAllFolders(root string) ([]string, goerror.Error) {
	return GetAllPathsFromCondition(func(path string) bool {
		return IsFolder(path)
	}, root)
}

func GetAllFiles(root string) ([]string, goerror.Error) {
	return GetAllPathsFromCondition(func(path string) bool {
		return file.IsFile(path)
	}, root)
}

func GetAllPathsFromCondition(conditionFn func(path string) bool, root string) ([]string, goerror.Error) {
	var paths []string

	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, errorslist.ErrorsList.Folder.ErrorOnGetAllPaths(root)
	}

	for _, entry := range entries {
		path := filepath.Join(root, entry.Name())
		if conditionFn(path) {
			paths = append(paths, path)
		}
	}

	return paths, goerror.CreateBlankError()
}

func CachedGetAllPaths(path string) ([]string, goerror.Error) {
	var content []string
	var err goerror.Error

	hasOnCache := cache.AssignWithCacheAndReturnIfExists(GET_ALL_PATHS_REGISTER_NAME, []any{path}, func(retornos []any) {
		content = retornos[0].([]string)
		err = retornos[1].(goerror.Error)
	})

	if hasOnCache {
		return content, err
	}

	content, err = GetAllPaths(path)

	cache.Register(cache.RegisterConfigs{
		Params:     []any{path},
		Retornos:   []any{content, err},
		NameUnique: GET_ALL_PATHS_REGISTER_NAME,
	})

	return content, err
}

func CachedGetAllPathsWithExt(ext string, path string) ([]string, goerror.Error) {
	var content []string
	var err goerror.Error

	hasOnCache := cache.AssignWithCacheAndReturnIfExists(GET_ALL_PATHS_WITH_EXT_REGISTER_NAME, []any{path}, func(retornos []any) {
		content = retornos[0].([]string)
		err = retornos[1].(goerror.Error)
	})

	if hasOnCache {
		return content, err
	}

	content, err = GetAllPathsWithExt(ext, path)

	cache.Register(cache.RegisterConfigs{
		Params:     []any{path},
		Retornos:   []any{content, err},
		NameUnique: GET_ALL_PATHS_REGISTER_NAME,
	})

	return content, err
}

func CachedGetAllFolders(root string) ([]string, goerror.Error) {
	var content []string
	var err goerror.Error

	hasOnCache := cache.AssignWithCacheAndReturnIfExists(GET_ALL_FOLDERS_REGISTER_NAME, []any{root}, func(retornos []any) {
		content = retornos[0].([]string)
		err = retornos[1].(goerror.Error)
	})

	if hasOnCache {
		return content, err
	}

	content, err = GetAllFolders(root)

	cache.Register(cache.RegisterConfigs{
		Params:     []any{root},
		Retornos:   []any{content, err},
		NameUnique: GET_ALL_PATHS_REGISTER_NAME,
	})

	return content, err
}

func CachedGetAllFiles(root string) ([]string, goerror.Error) {
	var content []string
	var err goerror.Error

	hasOnCache := cache.AssignWithCacheAndReturnIfExists(GET_ALL_FILES_REGISTER_NAME, []any{root}, func(retornos []any) {
		content = retornos[0].([]string)
		err = retornos[1].(goerror.Error)
	})

	if hasOnCache {
		return content, err
	}

	content, err = GetAllFiles(root)

	cache.Register(cache.RegisterConfigs{
		Params:     []any{root},
		Retornos:   []any{content, err},
		NameUnique: GET_ALL_PATHS_REGISTER_NAME,
	})

	return content, err
}
