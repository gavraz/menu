package menu

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
	state         bool
}

// NewHandler returns a menu handler with no menu attached
//
// A menu handler should be used only after setting the menu using Handler.SetMenu
func NewHandler() *Handler {
	return &Handler{}
}

// SetMenu sets the current menu of the menu handler
func (h *Handler) SetMenu(menu *Menu) {
	h.current = menu
}

// PrevChoice moves to the previous choice if applicable
func (h *Handler) PrevChoice() {
	if h.currentChoice == 0 {
		return
	}

	h.currentChoice--
}

// NextChoice moves to the next choice if applicable
func (h *Handler) NextChoice() {
	if h.currentChoice == len(h.current.actions)-1 {
		return
	}

	h.currentChoice++
}

func (h *Handler) nextMenu(next *Menu) {
	h.prev = append(h.prev, h.current)
	h.current = next
	h.currentChoice = 0
}

func (h *Handler) stateChanged(prev bool) bool {
	return prev != h.state
}

// Choose invokes the action attached to the current choice
func (h *Handler) Choose() {
	stateStamp := h.state
	act := h.current.actions[h.currentChoice]
	next := act()
	if h.current != next && !h.stateChanged(stateStamp) {
		h.nextMenu(next)
	}
}

func (h *Handler) changeState() {
	h.state = !h.state
}

// GoBack goes to the previous menu if applicable
func (h *Handler) GoBack() {
	if h.prev == nil || len(h.prev) == 0 {
		return
	}

	h.changeState()
	last := len(h.prev) - 1
	h.current = h.prev[last]
	h.prev[last] = nil
	h.prev = h.prev[:last]
	h.currentChoice = 0
}

// Choices returns a copy of the labels of the current menu
func (h *Handler) Choices() []string {
	c := make([]string, len(h.current.labels))
	copy(c, h.current.labels[:])
	return c
}

// CurrentChoice returns the current zero-based choice index
func (h *Handler) CurrentChoice() int {
	return h.currentChoice
}
