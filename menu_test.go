package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler_PrevNextChoice(t *testing.T) {
	h := NewHandler()
	m := NewBuilder(h).
		WithOption("1", func() {}).
		WithOption("2", func() {}).
		WithOption("3", func() {}).
		Build()
	h.SetMenu(m)

	assert.Equal(t, 0, h.currentChoice)

	h.PrevChoice()
	assert.Equal(t, 0, h.currentChoice)

	h.NextChoice()
	assert.Equal(t, 1, h.currentChoice)

	h.NextChoice()
	assert.Equal(t, 2, h.currentChoice)

	h.NextChoice()
	assert.Equal(t, 2, h.currentChoice)

	h.PrevChoice()
	assert.Equal(t, 1, h.currentChoice)
}

func TestHandler_Choose(t *testing.T) {
	var a1, a2 bool

	h := NewHandler()
	m := NewBuilder(h).
		WithOption("1", func() { a1 = true }).
		WithOption("2", func() { a2 = true }).
		Build()
	h.SetMenu(m)

	h.Choose()
	assert.True(t, a1)

	h.NextChoice()
	h.Choose()
	assert.True(t, a2)
}

func TestHandler_GoBack(t *testing.T) {
	h := NewHandler()
	sub := NewBuilder(h).Build()
	m := NewBuilder(h).
		WithSubMenu("2", sub).
		WithGoBack("1").
		Build()
	h.SetMenu(m)

	// no go-back should have no effect
	before := h.current
	h.NextChoice()
	h.Choose()
	assert.Equal(t, 1, h.currentChoice)
	assert.Equal(t, before, h.current)

	// choose the submenu
	h.PrevChoice()
	h.Choose()
	assert.Equal(t, 0, h.currentChoice)
	assert.Equal(t, sub, h.current)

	// go-back
	h.GoBack()
	assert.Equal(t, 0, h.currentChoice)
	assert.Equal(t, m, h.current)
}
