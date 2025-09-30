package json

import (
	"encoding/json"
	"os"

	errorslist "github.com/pedro-makoski/go-utils/errors-list"
	"github.com/pedro-makoski/go-utils/file"
	goerror "github.com/rantool-team/go-error"
)

func WriteJSONFile[Format any](path string, object Format, prefix string, indent string, permissionFile os.FileMode) goerror.Error {
	text, err := GetStringOnJSONFile(object, prefix, indent)

	if err.HasError() {
		return err
	}

	return file.WriteFile(path, text, permissionFile)
}

func GetStringOnJSONFile[Format any](object Format, prefix string, indent string) (string, goerror.Error) {
	text, err := json.MarshalIndent(object, prefix, indent)

	if err != nil {
		return "", errorslist.ErrorsList.JsonErrors.ErronOnMarshal(object)
	}

	return string(text), goerror.CreateBlankError()
}
