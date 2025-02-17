package tool

import "testing"

func TestSlidingWindow(t *testing.T) {

	window := NewSlidingWindow(3, 2)
	hotkey := window.AddCount(1)
	t.Logf("Is hotkey?: %v", hotkey)
}
