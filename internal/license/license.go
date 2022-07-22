package license

import (
	"strconv"
	"strings"
	"text/template"
)

type Kind uint8

const (
	None Kind = iota
	MIT
)

func GenerateMITLicense(year int, name string) (string, error) {
	tmpl, err := template.New("MIT License").Parse(mitTemplate)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	err = tmpl.Execute(&sb, struct {
		Year string
		Name string
	}{
		Year: strconv.Itoa(year),
		Name: name,
	})
	if err != nil {
		return "", err
	}

	return sb.String(), nil
}
