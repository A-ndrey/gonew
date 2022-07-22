package git

import (
	"fmt"
	"os/exec"
	"strings"
)

const (
	prefixSSH   = "git@"
	prefixHTTP  = "http://"
	prefixHTTPS = "https://"
	suffixGIT   = ".git"
)

func ProjectName(link string) (string, bool) {
	if !strings.HasSuffix(link, suffixGIT) {
		return "", false
	}

	sub := strings.Split(link, "/")
	if len(sub) == 0 {
		return "", false
	}

	projName := strings.TrimSuffix(sub[len(sub)-1], suffixGIT)

	return projName, true
}

func ModuleName(link string) string {
	if strings.HasPrefix(link, prefixSSH) {
		sshLink := strings.TrimPrefix(link, prefixSSH)
		sshLink = strings.TrimSuffix(sshLink, suffixGIT)
		return strings.Replace(sshLink, ":", "/", 1)
	}

	if strings.HasPrefix(link, prefixHTTPS) {
		httpsLink := strings.TrimPrefix(link, prefixHTTPS)
		httpsLink = strings.TrimSuffix(httpsLink, suffixGIT)
		return httpsLink
	}

	if strings.HasPrefix(link, prefixHTTP) {
		httpLink := strings.TrimPrefix(link, prefixHTTPS)
		httpLink = strings.TrimSuffix(httpLink, suffixGIT)
		return httpLink
	}

	return ""
}

func Clone(link string) error {
	cmd := exec.Command("git", "clone", link)
	fmt.Println("Cloning from git ", link)
	return cmd.Run()
}
