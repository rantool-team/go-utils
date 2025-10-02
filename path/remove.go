package path

import (
	"github.com/rantool-team/go-utils/folderfile"

	goerror "github.com/rantool-team/go-error"
)

func (p Path) Remove() goerror.Error {
	return folderfile.Remove(p.Path)
}
