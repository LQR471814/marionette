package marionette

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func getDefaultBrowserKey() (registry.Key, error) {
	browserKey, err := registry.OpenKey(
		registry.CURRENT_USER,
		`SOFTWARE\Microsoft\Windows\Shell\Associations\UrlAssociations\http\UserChoice`,
		registry.QUERY_VALUE,
	)

	return browserKey, err
}

func GetBrowserPath() (string, error) {
	browserKey, err := getDefaultBrowserKey()
	if err != nil {
		return "", err
	}

	progID, _, err := browserKey.GetStringValue("ProgId")
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

	cmd := exec.Command(path, args...)
	cmd.Run()

	return nil
}
