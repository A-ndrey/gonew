package git

import (
	"io"
	"net/http"
	"strings"
	"time"
)

const apiURL = "https://www.toptal.com/developers/gitignore/api/"

var client = http.Client{
	Timeout: 30 * time.Second,
}

func DownloadGitignore(options []string) (string, error) {
	resp, err := client.Get(apiURL + strings.Join(options, ","))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
