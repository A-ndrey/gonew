package file

import (
	"fmt"
	"os"
)

func Create(name, content string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func CreateMain() error {
	const fileName = "main.go"
	const content = `package main

func main() {

}

`

	return Create(fileName, content)
}

func CreateReadMe(projectName string) error {
	const fileName = "README.md"
	content := fmt.Sprintf("# %s", projectName)

	return Create(fileName, content)
}

func CreateGoMod(module, version string) error {
	const fileName = "go.mod"
	const template = `module %s

go %s

`
	content := fmt.Sprintf(template, module, version)

	return Create(fileName, content)
}
