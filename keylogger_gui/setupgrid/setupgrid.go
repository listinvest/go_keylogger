package setupgrid

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

// SetupKeyboard creates and attaches the buttons to the grid
func SetupKeyboard(grid *gtk.Grid) {
	// --- 1ST ROW ---
	escape, err := gtk.ToggleButtonNewWithLabel("Esc")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	f1, err := gtk.ToggleButtonNewWithLabel("F1")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	f2, err := gtk.ToggleButtonNewWithLabel("F2")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	f3, err := gtk.ToggleButtonNewWithLabel("F3")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	f4, err := gtk.ToggleButtonNewWithLabel("F4")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	f5, err := gtk.ToggleButtonNewWithLabel("F5")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	f6, err := gtk.ToggleButtonNewWithLabel("F6")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	f7, err := gtk.ToggleButtonNewWithLabel("F7")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	f8, err := gtk.ToggleButtonNewWithLabel("F8")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	f9, err := gtk.ToggleButtonNewWithLabel("F9")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	f10, err := gtk.ToggleButtonNewWithLabel("F10")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	f11, err := gtk.ToggleButtonNewWithLabel("F11")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	f12, err := gtk.ToggleButtonNewWithLabel("F12")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	pritnScreen, err := gtk.ToggleButtonNewWithLabel("Print\nScreen")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	scrollLock, err := gtk.ToggleButtonNewWithLabel("Scroll\nLock")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	pauseKey, err := gtk.ToggleButtonNewWithLabel("Pause")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	// --- 2ND ROW ---
	tilde, err := gtk.ToggleButtonNewWithLabel("~\n`")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	one, err := gtk.ToggleButtonNewWithLabel("!\n1")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	two, err := gtk.ToggleButtonNewWithLabel("@\n2")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	three, err := gtk.ToggleButtonNewWithLabel("#\n3")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	four, err := gtk.ToggleButtonNewWithLabel("$\n4")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	five, err := gtk.ToggleButtonNewWithLabel("%\n5")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	six, err := gtk.ToggleButtonNewWithLabel("^\n6")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	seven, err := gtk.ToggleButtonNewWithLabel("&\n7")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	eight, err := gtk.ToggleButtonNewWithLabel("*\n8")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	nine, err := gtk.ToggleButtonNewWithLabel("(\n9")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	zero, err := gtk.ToggleButtonNewWithLabel(")\n0")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	dash, err := gtk.ToggleButtonNewWithLabel("_\n-")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	equals, err := gtk.ToggleButtonNewWithLabel("+\n=")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	backspace, err := gtk.ToggleButtonNewWithLabel("Backspace")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	insert, err := gtk.ToggleButtonNewWithLabel("Insert")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	home, err := gtk.ToggleButtonNewWithLabel("Home")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	pageUp, err := gtk.ToggleButtonNewWithLabel("Page\nUp")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	numLock, err := gtk.ToggleButtonNewWithLabel("Num\nLock")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	divide, err := gtk.ToggleButtonNewWithLabel("/")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	multiply, err := gtk.ToggleButtonNewWithLabel("*")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	subtract, err := gtk.ToggleButtonNewWithLabel("-")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	empty, err := gtk.LabelNew("")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	// --- 3RD ROW ---

	// --- 4TH ROW ---

	// --- 5TH ROW ---

	// --- 6TH ROW ---

	// 1ST ROW
	grid.Attach(escape, 0, 0, 3, 3)
	grid.Attach(f1, 5, 0, 3, 3)
	grid.Attach(f2, 8, 0, 3, 3)
	grid.Attach(f3, 11, 0, 3, 3)
	grid.Attach(f4, 14, 0, 3, 3)
	grid.Attach(f5, 19, 0, 3, 3)
	grid.Attach(f6, 22, 0, 3, 3)
	grid.Attach(f7, 25, 0, 3, 3)
	grid.Attach(f8, 28, 0, 3, 3)
	grid.Attach(f9, 33, 0, 3, 3)
	grid.Attach(f10, 36, 0, 3, 3)
	grid.Attach(f11, 39, 0, 3, 3)
	grid.Attach(f12, 42, 0, 3, 3)
	grid.Attach(pritnScreen, 46, 0, 3, 3)
	grid.Attach(scrollLock, 49, 0, 3, 3)
	grid.Attach(pauseKey, 52, 0, 3, 3)
	// 2ND ROW
	grid.Attach(tilde, 0, 4, 3, 3)
	grid.Attach(one, 3, 4, 3, 3)
	grid.Attach(two, 6, 4, 3, 3)
	grid.Attach(three, 9, 4, 3, 3)
	grid.Attach(four, 12, 4, 3, 3)
	grid.Attach(five, 15, 4, 3, 3)
	grid.Attach(six, 18, 4, 3, 3)
	grid.Attach(seven, 21, 4, 3, 3)
	grid.Attach(eight, 24, 4, 3, 3)
	grid.Attach(nine, 27, 4, 3, 3)
	grid.Attach(zero, 30, 4, 3, 3)
	grid.Attach(dash, 33, 4, 3, 3)
	grid.Attach(equals, 36, 4, 3, 3)
	grid.Attach(backspace, 39, 4, 6, 3)
	grid.Attach(insert, 46, 4, 3, 3)
	grid.Attach(home, 49, 4, 3, 3)
	grid.Attach(pageUp, 52, 4, 3, 3)
	grid.Attach(numLock, 56, 4, 3, 3)
	grid.Attach(divide, 59, 4, 3, 3)
	grid.Attach(multiply, 62, 4, 3, 3)
	grid.Attach(subtract, 65, 4, 3, 3)
	// 3RD ROW

	// 4TH ROW

	// 5TH ROW

	// 6TH ROW

	// EMPTY SPACE
	grid.Attach(empty, 45, 0, 1, 1)
	grid.Attach(empty, 0, 3, 1, 1)
}
