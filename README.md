# Menu
A simple, yet flexible, menu with submenu.

## Example

### Import

```go
import "github.com/gavraz/menu/menu"

```

### Usage
```go
h := menu.NewHandler()

mainMenu := menu.NewBuilder(h).
  WithOption("Option A", func() {
  fmt.Println("You chose: Option A")
}).Build()

h.SetMenu(mainMenu)
```

To build a menu with submenu and a go-back option see main.go.

