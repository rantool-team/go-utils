package path

import (
	"github.com/pedro-makoski/go-utils/folderfile"

	goerror "github.com/rantool-team/go-error"
)

func (p Path) Copy(destinoPath string) goerror.Error {
	return folderfile.Copy(p.Path, destinoPath)
}
