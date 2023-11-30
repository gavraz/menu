package main

type action func() *Menu

// Menu represents a menu/submenu
type Menu struct {
	actions []action
	labels  []string
}

// Handler controls the state of the menu
type Handler struct {
	prev          []*Menu
	current       *Menu
	currentChoice int
}

// NewHandler returns a menu handler with no menu attached
//
// A menu handler should be used to manipulate a menu only after setting the menu using Handler.SetMenu
func NewHandler() *Handler {
	return &Handler{}
}

// SetMenu sets the current menu of the menu handler
func (h *Handler) SetMenu(menu *Menu) {
	h.current = menu
}

// PrevChoice moves to the previous choice if applicable
func (h *Handler) PrevChoice() {
	if h.currentChoice <= 0 {
		return
	}

	h.currentChoice--
}

// NextChoice moves to the next choice if applicable
func (h *Handler) NextChoice() {
	if h.currentChoice >= len(h.current.actions)-1 {
		return
	}

	h.currentChoice++
}

func (h *Handler) push(next *Menu) {
	h.prev = append(h.prev, h.current)
	h.current = next

	h.currentChoice = 0
}

// Choose invokes the action attached to the current choice
func (h *Handler) Choose() {
	act := h.current.actions[h.currentChoice]

	nextMenu := act()
	if nextMenu == nil {
		return
	}

	h.push(nextMenu)
}

func (h *Handler) pop() {
	last := len(h.prev) - 1
	h.current = h.prev[last]
	h.prev[last] = nil
	h.prev = h.prev[:last]

	h.currentChoice = 0
}

// GoBack goes to the previous menu if applicable
func (h *Handler) GoBack() {
	if h.prev == nil || len(h.prev) == 0 {
		return
	}

	h.pop()
}

// Choices returns the labels of the current menu
func (h *Handler) Choices() []string {
	c := make([]string, len(h.current.labels))
	copy(c, h.current.labels[:])
	return c
}

// CurrentChoice returns the current zero-based choice index
func (h *Handler) CurrentChoice() int {
	return h.currentChoice
}
