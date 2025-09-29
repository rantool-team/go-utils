package path

import (
	"github.com/pedro-makoski/go-utils/folderfile"

	goerror "github.com/rantool-team/go-error"
)

func (p *Path) RenameOrMove(newPath string) goerror.Error {
	defer func() {
		p.Path = newPath
	}()

	return folderfile.RenameOrMove(p.Path, newPath)
}
