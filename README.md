# Menu
A simple, yet flexible, menu with submenu.

### Example
```go
package main

import (
	"fmt"
	"github.com/gavraz/menu"
	"os"
)

func buildMenuHandler() *menu.Handler {
	h := menu.NewHandler()
	subMenu := menu.NewBuilder(h).
		WithOption("Inner Option 1", func() {
			fmt.Println("Inner 1")
		}).
		WithOption("Inner Option 2", func() {
			fmt.Println("Inner 2")
		}).
		WithGoBack("Go Back").
		Build()

	mainMenu := menu.NewBuilder(h).
		WithOption("Start", func() {
			fmt.Println("Starting...")
		}).
		WithSubMenu("Settings", subMenu).
		WithOption("Quit", func() {
			fmt.Println("quitting")
			os.Exit(0)
		}).
		Build()

	h.SetMenu(mainMenu)

	return h
}

func main() {
	h := buildMenuHandler()
	h.Choose()

	// choose submenu
	h.NextChoice()
	h.Choose()
	// choose inner 1
	h.Choose()
	h.GoBack()

	// choose to quit
	h.NextChoice()
	h.NextChoice()
	h.Choose()
}
```
