package errorslist

import (
	goerror "github.com/rantool-team/go-error"
)

type errorsList struct {
	File       fileErrorsList
	Folder     folderErrorsList
	FolderFile folderFileErrorsList
	JsonErrors jsonErrors
}

type jsonErrors struct {
	ErrorOnUmMarshall func(string) goerror.Error
	ErronOnMarshal    func(any) goerror.Error
}

type folderErrorsList struct {
	ErrorOnCreateFolder func(string) goerror.Error
	ErrorOnRenameFolder func(string, string) goerror.Error
	ErrorOnRemoveFolder func(string) goerror.Error
	ErrorOnGetAllPaths  func(string) goerror.Error
	CopyFolder          copyFolder
}

type fileErrorsList struct {
	ReadFile        func(string) goerror.Error
	ReadFileInfo    func(string) goerror.Error
	WriteFile       func(string) goerror.Error
	ErrorInOpenFile func(string) goerror.Error
	AppendInFile    func(string) goerror.Error
	RemoveFile      func(string) goerror.Error
	RenameFile      func(string, string) goerror.Error
	GetAbs          func(string) goerror.Error
	CopyErrors      copyFile
}

type folderFileErrorsList struct {
	ErrorOnGetInfo  func(string) goerror.Error
	ItMustBeAFolder func(string) goerror.Error
	ItMustBEAFile   func(string) goerror.Error
}

type copyFile struct {
	ErrorInOpenFileOrigin      func(string) goerror.Error
	ErrorInOpenDestinationFile func(string) goerror.Error
	ErrorInCopyContent         func(string, string) goerror.Error
	ErrorInSyncFile            func(string) goerror.Error
}

type copyFolder struct {
	ErrorOnCopyFolder      func(string, string) goerror.Error
	ErrorOnGetRelativePath func(string, string) goerror.Error
	ErrorOnGetFileInfo     func(string) goerror.Error
	ErrorOnCreateFolder    func(string) goerror.Error
}
