package json

import (
	"encoding/json"

	errorslist "github.com/pedro-makoski/go-utils/errors-list"
	"github.com/pedro-makoski/go-utils/file"

	goerror "github.com/rantool-team/go-error"
)

func GetJsonFile[format any](path string, obj *format) goerror.Error {
	content, err := file.ReadFile(path)
	if err.HasError() {
		return err
	}

	return Convert(content, obj)
}

func Convert[format any](content string, obj *format) goerror.Error {
	err := json.Unmarshal([]byte(content), &obj)

	if err != nil {
		return errorslist.ErrorsList.JsonErrors.ErrorOnUmMarshall(content)
	}

	return goerror.CreateBlankError()
}
