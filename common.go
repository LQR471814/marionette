package marionette

import "regexp"

func filterBrowserIDMap(target string, comp map[string]BrowserType) BrowserType {
	for key := range idMap {
		matches, err := regexp.MatchString(key, target)
		if err != nil {
			return UNDEFINED
		}

		if matches {
			return idMap[key]
		}
	}

	return UNDEFINED
}

type BrowserType int

const (
	UNDEFINED BrowserType = iota
	CHROME
	EDGE
	FIREFOX
)

type UnknownBrowserType struct{}

func (*UnknownBrowserType) Error() string {
	return "Default browser is not recognized!"
}
