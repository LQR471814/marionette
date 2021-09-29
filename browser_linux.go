package marionette

import (
	"os/exec"
	"regexp"
)

var defaultBrowserMap map[string]string = map[string]string{
	"firefox": "firefox",
	"chrome":  "google-chrome",
}

func GetBrowserPath() (string, error) {
	defaultBrowser, err := exec.Command("xdg-settings", "get", "default-web-browser").Output()
	if err != nil {
		return "", err
	}

	browser := ""
	for key := range defaultBrowserMap {
		matched, err := regexp.Match(key, defaultBrowser)
		if err != nil {
			return "", err
		}

		if matched {
			browser = defaultBrowserMap[key]
			break
		}
	}

	path, err := exec.Command("which", browser).Output()
	return string(path), err
}

func OpenBrowser(args ...string) error {
	path, err := GetBrowserPath()
	if err != nil {
		return err
	}

	return exec.Command(path, args...).Run()
}
