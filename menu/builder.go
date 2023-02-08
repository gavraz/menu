package menu

// Builder builds a menu and submenu hierarchies.
//
// The order of the resulting menu is equivalent to the build order.
// Labels can be reused and will result in a new menu entry.
type Builder struct {
	menu    *Menu
	handler *Handler
}

// NewBuilder instantiates a menu builder that can be controlled by the given menu Handler
func NewBuilder(h *Handler) *Builder {
	return &Builder{
		menu:    &Menu{},
		handler: h,
	}
}

func (mb *Builder) add(label string, action action) {
	mb.menu.actions = append(mb.menu.actions, action)
	mb.menu.labels = append(mb.menu.labels, label)
}

// WithOption adds an option to the current menu with given label and action
func (mb *Builder) WithOption(label string, action func()) *Builder {
	mb.add(label, func() *Menu {
		action()
		return mb.menu
	})

	return mb
}

// WithGoBack adds a go-back option to the current menu with the given label
func (mb *Builder) WithGoBack(label string) *Builder {
	mb.add(label, func() *Menu {
		mb.handler.GoBack()
		return mb.menu
	})

	return mb
}

// WithSubMenu adds a submenu option to the current menu with the given label
func (mb *Builder) WithSubMenu(label string, submenu *Menu) *Builder {
	mb.add(label, func() *Menu {
		return submenu
	})

	return mb
}

// Build returns the menu
func (mb *Builder) Build() *Menu {
	return mb.menu
}
