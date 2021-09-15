## marionette

A tiny package for manipulating the default browser cross platform.

### Usage

```go
package main

import (
    "fmt"
    "github.com/LQR471814/marionette"
)

func main() {
    path, err := marionette.GetBrowserPath()
    if err != nil {
        panic(err)
    }

    fmt.Println(path)

    marionette.OpenBrowser(`--app=https://google.com`, `--guest`)
}
```
