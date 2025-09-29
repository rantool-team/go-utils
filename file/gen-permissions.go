package file

import (
	"os"
)

type FilePermissions struct {
	OwnerRead    bool
	OwnerWrite   bool
	OwnerExec    bool
	GroupRead    bool
	GroupWrite   bool
	GroupExec    bool
	OtherRead    bool
	OtherWrite   bool
	OtherExec    bool
	IsDir        bool
	IsRegular    bool
	IsSymlink    bool
	IsSocket     bool
	IsNamedPipe  bool
	IsDevice     bool
	IsCharDevice bool
}

func (fp FilePermissions) ToFileMode() os.FileMode {
	var mode os.FileMode

	if fp.OwnerRead {
		mode |= 0400
	}
	if fp.OwnerWrite {
		mode |= 0200
	}
	if fp.OwnerExec {
		mode |= 0100
	}
	if fp.GroupRead {
		mode |= 0040
	}
	if fp.GroupWrite {
		mode |= 0020
	}
	if fp.GroupExec {
		mode |= 0010
	}
	if fp.OtherRead {
		mode |= 0004
	}
	if fp.OtherWrite {
		mode |= 0002
	}
	if fp.OtherExec {
		mode |= 0001
	}

	if fp.IsDir {
		mode |= os.ModeDir
	}
	if fp.IsSymlink {
		mode |= os.ModeSymlink
	}
	if fp.IsSocket {
		mode |= os.ModeSocket
	}
	if fp.IsNamedPipe {
		mode |= os.ModeNamedPipe
	}
	if fp.IsDevice {
		mode |= os.ModeDevice
	}
	if fp.IsCharDevice {
		mode |= os.ModeCharDevice
	}

	return mode
}
