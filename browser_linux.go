package marionette

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var browserWhichQuery map[BrowserType]string = map[BrowserType]string{
	CHROME:  "google-chrome",
	FIREFOX: "firefox",
}

var idMap map[string]BrowserType = map[string]BrowserType{
	"chrome":  CHROME,
	"firefox": FIREFOX,
}

func DefaultBrowser() (BrowserType, error) {
	defaultBrowser, err := exec.Command("xdg-settings", "get", "default-web-browser").Output()
	if err != nil {
		return UNDEFINED, err
	}

	for key := range idMap {
		matches, err := regexp.MatchString(key, string(defaultBrowser))
		if err != nil {
			return UNDEFINED, err
		}

		if matches {
			return idMap[key], nil
		}
	}

	return UNDEFINED, &UnknownBrowserType{}
}

func GetBrowserPath() (string, error) {
	browser, err := DefaultBrowser()
	if err != nil {
		return "", err
	}

	path, err := exec.Command("which", browserWhichQuery[browser]).Output()
	pathStr := strings.TrimSpace(string(path))
	return pathStr, err
}

func OpenBrowser(args ...string) error {
	path, err := GetBrowserPath()
	if err != nil {
		return err
	}

	fmt.Println(args)

	return exec.Command(path, args...).Run()
}
