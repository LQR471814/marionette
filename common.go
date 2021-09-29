package marionette

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
