package generator

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"path/filepath"
	"talkspace/utils/constant"
)

func GenerateEmailTemplate(fileTemplate string, data interface{}) (string, error) {
	templatePath, err := filepath.Abs(fmt.Sprintf("talkspace/utils/helper/email/template/%s", fileTemplate))
	if err != nil {
		return "", errors.New(constant.ERROR_TEMPLATE_FILE)
	}

	emailTemplate, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	var templateBuffer bytes.Buffer
	if err := emailTemplate.Execute(&templateBuffer, data); err != nil {
		return "", err
	}

	return templateBuffer.String(), nil
}
