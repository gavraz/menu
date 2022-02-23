package menu

type Builder struct {
	menu    *Menu
	handler *Handler
}

func NewBuilder(h *Handler) *Builder {
	return &Builder{
		menu:    &Menu{},
		handler: h,
	}
}

func (mb *Builder) init() {
	if mb.menu == nil {
		mb.menu = &Menu{}
	}
}

func (mb *Builder) add(label string, action action) {
	mb.init()

	mb.menu.actions = append(mb.menu.actions, action)
	mb.menu.labels = append(mb.menu.labels, label)
}

func (mb *Builder) WithOption(label string, action func()) *Builder {
	mb.add(label, func() *Menu {
		action()
		return mb.menu
	})

	return mb
}

func (mb *Builder) WithGoBack(label string) *Builder {
	mb.add(label, func() *Menu {
		mb.handler.GoBack()
		return mb.menu
	})

	return mb
}

func (mb *Builder) WithSubMenu(label string, submenu *Menu) *Builder {
	mb.add(label, func() *Menu {
		return submenu
	})

	return mb
}

func (mb *Builder) Build() *Menu {
	return mb.menu
}
