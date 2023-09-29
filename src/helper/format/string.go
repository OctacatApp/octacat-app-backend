package format

import (
	"bytes"
	"html/template"
)

func String[D any](str string, data D) (string, error) {
	tmpl, err := template.New("tmpl").Parse(str)
	if err != nil {
		return "", err
	}

	buff := new(bytes.Buffer)
	if err := tmpl.Execute(buff, data); err != nil {
		return "", err
	}

	return buff.String(), nil
}
