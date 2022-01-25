package marionette

import (
	"fmt"
	"os/exec"
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
	id, err := getDefaultBrowserProgID()
	if err != nil {
		return UNDEFINED, err
	}

	browser := filterBrowserIDMap(id, idMap)
	if browser == UNDEFINED {
		return UNDEFINED, &UnknownBrowserType{}
	}

	return browser, nil
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

	path := strings.Split(pathValue, "\"")[1]
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
