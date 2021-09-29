package marionette

import (
	"fmt"
	"os/exec"
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
	id, err := exec.Command("xdg-settings", "get", "default-web-browser").Output()
	if err != nil {
		return UNDEFINED, err
	}

	browser := filterBrowserIDMap(string(id), idMap)
	if browser == UNDEFINED {
		return UNDEFINED, &UnknownBrowserType{}
	}

	return browser, nil
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
