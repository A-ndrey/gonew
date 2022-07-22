package builder

import (
	"github.com/A-ndrey/gonew/internal/file"
	"github.com/A-ndrey/gonew/internal/git"
	"github.com/A-ndrey/gonew/internal/license"
	"os"
	"time"
)

type Builder struct {
	projectName string
	authorName  string
	gitLink     string
	gitignore   []string
	gomod       goMod
	readMe      bool
	license     license.Kind
}

type goMod struct {
	moduleName string
	goVersion  string
}

func (b *Builder) Build() error {
	chain := []func() error{
		b.gitCloneOrCreateDir,
		b.genProjectStruct,
		b.addGoModFile,
		b.addGitignoreFile,
		b.addReadMeFile,
		b.addLicenseFile,
	}

	for _, f := range chain {
		err := f()
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Builder) gitCloneOrCreateDir() error {
	var err error
	if b.gitLink != "" {
		err = git.Clone(b.gitLink)
	} else {
		err = os.Mkdir(b.projectName, os.ModePerm)
	}

	if err != nil {
		return err
	}

	err = os.Chdir(b.projectName)
	if err != nil {
		return err

	}

	return nil
}

func (b *Builder) genProjectStruct() error {
	const (
		cmdDir      = "cmd"
		internalDir = "internal"
		back        = ".."
	)

	err := os.Mkdir(cmdDir, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.Mkdir(internalDir, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.Chdir(cmdDir)
	if err != nil {
		return err
	}

	err = file.CreateMain()
	if err != nil {
		return err
	}

	err = os.Chdir(back)
	if err != nil {
		return err
	}

	return nil
}

func (b *Builder) addGoModFile() error {
	var zeroVal goMod
	if b.gomod == zeroVal {
		return nil
	}

	err := file.CreateGoMod(b.gomod.moduleName, b.gomod.goVersion)
	if err != nil {
		return err
	}

	return nil
}

func (b *Builder) addGitignoreFile() error {
	if len(b.gitignore) == 0 {
		return nil
	}

	const fileName = ".gitignore"
	content, err := git.DownloadGitignore(b.gitignore)
	if err != nil {
		return err
	}

	err = file.Create(fileName, content)
	if err != nil {
		return err
	}

	return nil
}

func (b *Builder) addReadMeFile() error {
	err := file.CreateReadMe(b.projectName)
	if err != nil {
		return err
	}

	return nil
}

func (b *Builder) addLicenseFile() error {
	if b.license == license.None {
		return nil
	}

	var licenseContent string
	var err error

	switch b.license {
	case license.MIT:
		licenseContent, err = license.GenerateMITLicense(time.Now().Year(), b.authorName)
	}
	if err != nil {
		return err
	}

	err = file.Create("LICENSE", licenseContent)
	if err != nil {
		return err
	}

	return nil
}
