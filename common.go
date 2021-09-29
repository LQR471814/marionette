package marionette

type BrowserType int

const (
	CHROME BrowserType = iota
	EDGE
	FIREFOX
)

type UnknownBrowserType struct{}

func (*UnknownBrowserType) Error() string {
	return "Default browser is not recognized!"
}
