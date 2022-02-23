package main

type Builder struct {
	menu    *main.Menu
	handler *main.Handler
}

func NewBuilder(h *main.Handler) *Builder {
	return &Builder{
		menu:    &main.Menu{},
		handler: h,
	}
}

func (mb *Builder) init() {
	if mb.menu == nil {
		mb.menu = &main.Menu{}
	}
}

func (mb *Builder) add(label string, action main.action) {
	mb.init()

	mb.menu.actions = append(mb.menu.actions, action)
	mb.menu.labels = append(mb.menu.labels, label)
}

func (mb *Builder) WithOption(label string, action func()) *Builder {
	mb.add(label, func() *main.Menu {
		action()
		return mb.menu
	})

	return mb
}

func (mb *Builder) WithGoBack(label string) *Builder {
	mb.add(label, func() *main.Menu {
		mb.handler.GoBack()
		return mb.menu
	})

	return mb
}

func (mb *Builder) WithSubMenu(label string, submenu *main.Menu) *Builder {
	mb.add(label, func() *main.Menu {
		return submenu
	})

	return mb
}

func (mb *Builder) Build() *main.Menu {
	return mb.menu
}
