package builder

import (
	"bufio"
	"fmt"
	"github.com/A-ndrey/gonew/file/license"
	"github.com/A-ndrey/gonew/git"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func CreateFromUserChoices() *Builder {
	var b Builder
	b.askProject()
	b.askGoMod()
	b.askGitignore()
	b.askReadMe()
	b.askLicense()

	return &b
}

func (b *Builder) askProject() {
	for {
		userInput := inputStr("Enter project name or link to git project", "")
		if projName, ok := git.ProjectName(userInput); ok {
			b.projectName = projName
			b.gitLink = userInput
		} else {
			b.projectName = userInput
		}
		if b.projectName != "" {
			return
		}
		fmt.Println("Project name cannot be empty")
	}
}

func (b *Builder) askGoMod() {
	if !inputBool("Add go.mod file", true) {
		return
	}
	if b.gitLink != "" {
		b.gomod.moduleName = git.ModuleName(b.gitLink)
	}
	if b.gomod.moduleName == "" {
		b.gomod.moduleName = b.projectName
	}
	validVersionRE := regexp.MustCompile(`^1.\d{2}$`)
	for {
		b.gomod.goVersion = inputStr("Specify version of Go", defaultGoVersion())
		if validVersionRE.MatchString(b.gomod.goVersion) {
			return
		}
		fmt.Println("Go version is incorrect")
	}
}

func (b *Builder) askGitignore() {
	if !inputBool("Add .gitignore file", false) {
		return
	}
	b.gitignore = append(b.gitignore, "go")
	ide := inputStr("Specify your IDE (e.g. GoLand/VSCode/Vim)", "")
	if ide != "" {
		b.gitignore = append(b.gitignore, ide)
	}
}

func (b *Builder) askReadMe() {
	b.readMe = inputBool("Add README.md file", false)
}

func (b *Builder) askLicense() {
	if !inputBool("Add LICENSE file (currently available only MIT)", true) {
		return
	}
	b.license = license.MIT
	b.authorName = inputStr("Enter your name for license", "")
}

func inputStr(message, def string) string {
	const (
		template            = ">>> %s: "
		templateWithDefault = ">>> %s (default %s): "
	)

	if def == "" {
		fmt.Printf(template, message)
	} else {
		fmt.Printf(templateWithDefault, message, def)
	}

	scanner.Scan()
	ans := scanner.Text()
	ans = strings.TrimSpace(ans)
	if ans == "" {
		return def
	}
	return ans
}

func inputBool(message string, def bool) bool {
	defStr := "N"
	if def {
		defStr = "Y"
	}
	ans := inputStr(fmt.Sprint(message, " [Y/N]"), defStr)
	ans = strings.ToLower(ans)
	switch ans {
	case "y", "yes":
		return true
	default:
		return false
	}
}

func defaultGoVersion() (ver string) {
	ver = "1.18"
	verRE := regexp.MustCompile(`go(1.\d{2})`)
	out, err := exec.Command("go", "version").Output()
	if err != nil {
		return
	}

	found := verRE.FindSubmatch(out)
	if len(found) < 2 {
		return
	}

	ver = string(found[1])
	return
}
