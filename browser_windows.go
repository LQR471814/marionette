package marionette

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"golang.org/x/sys/windows/registry"
)

var idMap map[string]BrowserType = map[string]BrowserType{
	"Chrome":  CHROME,
	"Firefox": FIREFOX,
	"Edge":    EDGE,
}

func getDefaultBrowserProgID() (string, error) {
	browserKey, err := registry.OpenKey(
		registry.CURRENT_USER,
		`SOFTWARE\Microsoft\Windows\Shell\Associations\UrlAssociations\http\UserChoice`,
		registry.QUERY_VALUE,
	)

	if err != nil {
		return "", err
	}

	progID, _, err := browserKey.GetStringValue("ProgId")

	return progID, err
}

func DefaultBrowser() (BrowserType, error) {
	browser, err := getDefaultBrowserProgID()
	if err != nil {
		return -1, err
	}

	for key := range idMap {
		matches, err := regexp.MatchString(key, browser)
		if err != nil {
			return -1, err
		}

		if matches {
			return idMap[key], nil
		}
	}

	return -1, &UnknownBrowserType{}
}

func GetBrowserPath() (string, error) {
	progID, err := getDefaultBrowserProgID()
	if err != nil {
		return "", err
	}

	pathKey, err := registry.OpenKey(
		registry.LOCAL_MACHINE,
		fmt.Sprintf(`SOFTWARE\Classes\%v\shell\open\command`, progID),
		registry.QUERY_VALUE,
	)
	if err != nil {
		return "", err
	}

	pathValue, _, err := pathKey.GetStringValue("")
	if err != nil {
		return "", err
	}

	re, err := regexp.Compile(`".*"`)
	if err != nil {
		return "", err
	}

	path := string(re.Find([]byte(pathValue)))
	path = strings.ReplaceAll(path, "\"", "")

	return path, err
}

func OpenBrowser(args ...string) error {
	path, err := GetBrowserPath()
	if err != nil {
		return err
	}

	return exec.Command(path, args...).Run()
}
