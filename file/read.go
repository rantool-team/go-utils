package file

import (
	"os"

	"github.com/pedro-makoski/go-utils/cache"
	errorslist "github.com/pedro-makoski/go-utils/errors-list"

	goerror "github.com/rantool-team/go-error"
)

const READ_FILE_NAME_UNIQUE = "file.readFile"
const READ_FILE_INFO_NAME_UNIQUE = "file.readFileInfo"

func ReadFile(path string) (string, goerror.Error) {
	file, err := os.ReadFile(path)

	if err != nil {
		return "", errorslist.ErrorsList.File.ReadFile(path)
	}

	return string(file), goerror.CreateBlankError()
}

func CacheReadFile(path string) (string, goerror.Error) {
	var content string
	var err goerror.Error

	hasOnCache := cache.AssignWithCacheAndReturnIfExists(READ_FILE_NAME_UNIQUE, []any{path}, func(retornos []any) {
		content = retornos[0].(string)
		err = retornos[1].(goerror.Error)
	})

	if hasOnCache {
		return content, err
	}

	content, err = ReadFile(path)

	cache.Register(cache.RegisterConfigs{
		Params:     []any{path},
		Retornos:   []any{content, err},
		NameUnique: READ_FILE_NAME_UNIQUE,
	})

	return content, err
}

func ReadFileInfo(path string) (*FileInfo, goerror.Error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, errorslist.ErrorsList.File.ReadFileInfo(path)
	}

	content, errorFull := ReadFile(path)
	if err != nil {
		return nil, errorFull
	}

	return &FileInfo{
		Name:    info.Name(),
		Size:    info.Size(),
		Mode:    info.Mode(),
		ModTime: info.ModTime(),
		IsDir:   info.IsDir(),
		Content: content,
	}, goerror.CreateBlankError()
}
