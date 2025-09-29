package errorslist

import (
	goerror "github.com/rantool-team/go-error"
	"github.com/rantool-team/go-error/language"
)

var ErrorsList = errorsList{
	File: fileErrorsList{
		ReadFile: func(path string) goerror.Error {
			return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
				Portuguese: "Erro ao ler arquivo " + path,
				English:    "Error on Read file " + path,
			}, language.MessageSet{})
		},
		ReadFileInfo: func(path string) goerror.Error {
			return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
				Portuguese: "Erro ao ler informações do arquivo " + path,
				English:    "Error on Read file " + path + " infos",
			}, language.MessageSet{})
		},
		WriteFile: func(path string) goerror.Error {
			return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
				Portuguese: "Erro ao escrever no arquivo " + path,
				English:    "Error on Write file " + path,
			}, language.MessageSet{})
		},
		AppendInFile: func(path string) goerror.Error {
			return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
				Portuguese: "Erro ao adicionar linha no arquivo " + path,
				English:    "Error on Append Line in file " + path,
			}, language.MessageSet{})
		},
		RemoveFile: func(path string) goerror.Error {
			return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
				Portuguese: "Erro ao remover o arquivo: " + path,
				English:    "Error on Remove file " + path,
			}, language.MessageSet{})
		},
		CopyErrors: copyFile{
			ErrorInOpenFileOrigin: func(path string) goerror.Error {
				return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
					Portuguese: "Erro ao abrir arquivo de origem: " + path,
					English:    "Error in open origin file " + path,
				}, language.MessageSet{})
			},
			ErrorInOpenDestinationFile: func(path string) goerror.Error {
				return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
					Portuguese: "Erro ao abrir arquivo de destino: " + path,
					English:    "Error in open destination file " + path,
				}, language.MessageSet{})
			},
			ErrorInCopyContent: func(src string, dst string) goerror.Error {
				return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
					Portuguese: "Erro ao copiar o conteúdo do arquivo: " + src + " para o arquivo: " + dst,
					English:    "Error in copy content from the file " + src + " to the file " + dst,
				}, language.MessageSet{})
			},
			ErrorInSyncFile: func(path string) goerror.Error {
				return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
					Portuguese: "Erro ao sincronizar arquivo: " + path,
					English:    "Error in sync the file " + path,
				}, language.MessageSet{})
			},
		},
		ErrorInOpenFile: func(path string) goerror.Error {
			return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
				Portuguese: "Erro ao abrir o arquivo: " + path,
				English:    "Error in open file: " + path,
			}, language.MessageSet{})
		},
		RenameFile: func(src string, dst string) goerror.Error {
			return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
				Portuguese: "Erro ao renomear o arquivo " + src + " para o nome " + dst,
				English:    "Error in reanme the file " + src + " for the name " + dst,
			}, language.MessageSet{})
		},
		GetAbs: func(path string) goerror.Error {
			return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
				Portuguese: "Erro ao obter caminho absoluto do arquivo: " + path,
				English:    "Error in get abs file from: " + path,
			}, language.MessageSet{})
		},
	},
	FolderFile: folderFileErrorsList{
		ErrorOnGetInfo: func(path string) goerror.Error {
			return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
				Portuguese: "Erro ao obter informações do caminho: " + path,
				English:    "Error in get path from: " + path,
			}, language.MessageSet{})
		},
		ItMustBeAFolder: func(path string) goerror.Error {
			return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
				Portuguese: "Essa operação só pode ser realizada com pastas, e o arquivo " + path + " Não é uma",
				English:    "This operation only can be possible with folders, and the the file " + path + " it's not",
			}, language.MessageSet{})
		},
		ItMustBEAFile: func(path string) goerror.Error {
			return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
				Portuguese: "Essa operação só pode ser realizada com arquivos, e a pasta " + path + " Não é uma",
				English:    "This operation only can be possible with files, and the the folder " + path + " it's not",
			}, language.MessageSet{})
		},
	},
	JsonErrors: jsonErrors{
		ErrorOnUmMarshall: func(content string) goerror.Error {
			return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
				Portuguese: "O conteúdo a seguir não pode ser transformado em JSON\nContent:\n" + content,
				English:    "The folowing content can't be transformed in JSON\nContent:\n" + content,
			}, language.MessageSet{})
		},
	},
	Folder: folderErrorsList{

		ErrorOnGetAllPaths: func(path string) goerror.Error {
			return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
				Portuguese: "Erro ao ler os arquivos da pasta: " + path,
				English:    "Error on Read the file from the folder " + path,
			}, language.MessageSet{})
		},
		ErrorOnCreateFolder: func(path string) goerror.Error {
			return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
				Portuguese: "Erro ao criar a pasta " + path,
				English:    "Error on create folder " + path,
			}, language.MessageSet{})
		},
		ErrorOnRemoveFolder: func(path string) goerror.Error {
			return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
				Portuguese: "Erro ao remover a pasta: " + path,
				English:    "Error on Remove folder " + path,
			}, language.MessageSet{})
		},
		ErrorOnRenameFolder: func(src string, dst string) goerror.Error {
			return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
				Portuguese: "Erro ao renomear a pasta: " + src + " para o nome " + dst,
				English:    "Error in rename the folder: " + src + " for the name " + dst,
			}, language.MessageSet{})
		},
		CopyFolder: copyFolder{
			ErrorOnGetFileInfo: func(path string) goerror.Error {
				return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
					Portuguese: "Erro ao arquivo de origem: " + path,
					English:    "Error in open origin file " + path,
				}, language.MessageSet{})
			},
			ErrorOnCopyFolder: func(src string, dst string) goerror.Error {
				return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
					Portuguese: "Erro ao copiar a pasta " + src + " para o destino " + dst,
					English:    "Error in copy the folder " + src + " for the destination " + dst,
				}, language.MessageSet{})
			},
			ErrorOnGetRelativePath: func(src string, dst string) goerror.Error {
				return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
					Portuguese: "Erro de obter o caminho relativo de " + src + " com " + dst,
					English:    "Error in get the relative folder " + src + " to " + dst,
				}, language.MessageSet{})
			},
			ErrorOnCreateFolder: func(path string) goerror.Error {
				return goerror.CreateErrorInMultiplesLanguage(language.MessageSet{
					Portuguese: "Erro ao criar a pasta: " + path,
					English:    "Error in create folder " + path,
				}, language.MessageSet{})
			},
		},
	},
}
