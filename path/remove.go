package path

import (
	"github.com/pedro-makoski/go-utils/folderfile"

	goerror "github.com/rantool-team/go-error"
)

func (p Path) Remove() goerror.Error {
	return folderfile.Remove(p.Path)
}
