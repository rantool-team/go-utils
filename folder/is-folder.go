package folder

import "github.com/pedro-makoski/go-utils/file"

func IsFolder(path string) bool {
	return !file.IsFile(path)
}
