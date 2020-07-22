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

	// --- 3RD ROW ---
	tab, err := gtk.ToggleButtonNewWithLabel("Tab")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	q, err := gtk.ToggleButtonNewWithLabel("Q")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	w, err := gtk.ToggleButtonNewWithLabel("W")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	e, err := gtk.ToggleButtonNewWithLabel("E")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	r, err := gtk.ToggleButtonNewWithLabel("R")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	t, err := gtk.ToggleButtonNewWithLabel("T")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	y, err := gtk.ToggleButtonNewWithLabel("Y")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	u, err := gtk.ToggleButtonNewWithLabel("U")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	i, err := gtk.ToggleButtonNewWithLabel("I")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	o, err := gtk.ToggleButtonNewWithLabel("O")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	p, err := gtk.ToggleButtonNewWithLabel("P")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	leftBracket, err := gtk.ToggleButtonNewWithLabel("{\n[")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	rightBracket, err := gtk.ToggleButtonNewWithLabel("}\n]")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	backslash, err := gtk.ToggleButtonNewWithLabel("|\n\\")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	deleteKey, err := gtk.ToggleButtonNewWithLabel("Del")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	end, err := gtk.ToggleButtonNewWithLabel("End")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	pageDown, err := gtk.ToggleButtonNewWithLabel("Page\nDown")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	num7, err := gtk.ToggleButtonNewWithLabel("7")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	num8, err := gtk.ToggleButtonNewWithLabel("8")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	num9, err := gtk.ToggleButtonNewWithLabel("9")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	plus, err := gtk.ToggleButtonNewWithLabel("+")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	// --- 4TH ROW ---
	capsLock, err := gtk.ToggleButtonNewWithLabel("Caps Lock")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	a, err := gtk.ToggleButtonNewWithLabel("A")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	s, err := gtk.ToggleButtonNewWithLabel("S")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	d, err := gtk.ToggleButtonNewWithLabel("D")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	f, err := gtk.ToggleButtonNewWithLabel("F")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	g, err := gtk.ToggleButtonNewWithLabel("G")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	h, err := gtk.ToggleButtonNewWithLabel("H")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	j, err := gtk.ToggleButtonNewWithLabel("J")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	k, err := gtk.ToggleButtonNewWithLabel("K")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	l, err := gtk.ToggleButtonNewWithLabel("L")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	colon, err := gtk.ToggleButtonNewWithLabel(":\n;")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	quote, err := gtk.ToggleButtonNewWithLabel("\"\n'")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	enter, err := gtk.ToggleButtonNewWithLabel("Enter")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	num4, err := gtk.ToggleButtonNewWithLabel("4")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	num5, err := gtk.ToggleButtonNewWithLabel("5")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	num6, err := gtk.ToggleButtonNewWithLabel("6")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	// --- 5TH ROW ---
	leftShift, err := gtk.ToggleButtonNewWithLabel("Shift")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	z, err := gtk.ToggleButtonNewWithLabel("Z")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	x, err := gtk.ToggleButtonNewWithLabel("X")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	c, err := gtk.ToggleButtonNewWithLabel("C")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	v, err := gtk.ToggleButtonNewWithLabel("V")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	b, err := gtk.ToggleButtonNewWithLabel("B")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	n, err := gtk.ToggleButtonNewWithLabel("N")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	m, err := gtk.ToggleButtonNewWithLabel("M")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	comma, err := gtk.ToggleButtonNewWithLabel("<\n,")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	period, err := gtk.ToggleButtonNewWithLabel(">\n.")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	forwardSlash, err := gtk.ToggleButtonNewWithLabel("?\n/")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	rightShift, err := gtk.ToggleButtonNewWithLabel("Shift")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	upArrow, err := gtk.ToggleButtonNewWithLabel("🠕")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	num1, err := gtk.ToggleButtonNewWithLabel("1")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	num2, err := gtk.ToggleButtonNewWithLabel("2")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	num3, err := gtk.ToggleButtonNewWithLabel("3")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	numEnter, err := gtk.ToggleButtonNewWithLabel("Enter")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	// --- 6TH ROW ---
	leftControl, err := gtk.ToggleButtonNewWithLabel("Ctrl")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	leftWindows, err := gtk.ToggleButtonNewWithLabel("Win")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	leftAlt, err := gtk.ToggleButtonNewWithLabel("Alt")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	spacebar, err := gtk.ToggleButtonNewWithLabel("Spacebar")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	rightAlt, err := gtk.ToggleButtonNewWithLabel("Alt")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	rightWindows, err := gtk.ToggleButtonNewWithLabel("Win")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	menuKey, err := gtk.ToggleButtonNewWithLabel("⊟")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	rightControl, err := gtk.ToggleButtonNewWithLabel("Ctrl")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	leftArrow, err := gtk.ToggleButtonNewWithLabel("🠔")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	downArrow, err := gtk.ToggleButtonNewWithLabel("🠗")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	rightArrow, err := gtk.ToggleButtonNewWithLabel("🠖")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	num0, err := gtk.ToggleButtonNewWithLabel("0")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	decimal, err := gtk.ToggleButtonNewWithLabel(".")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

	// --- EMPTY SPACE ---
	empty0, err := gtk.LabelNew("\t")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}
	empty1, err := gtk.LabelNew("\t")
	if err != nil {
		log.Fatal("Unable to make button:", err)
	}

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
	grid.Attach(pritnScreen, 47, 0, 3, 3)
	grid.Attach(scrollLock, 50, 0, 3, 3)
	grid.Attach(pauseKey, 53, 0, 3, 3)
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
	grid.Attach(insert, 47, 4, 3, 3)
	grid.Attach(home, 50, 4, 3, 3)
	grid.Attach(pageUp, 53, 4, 3, 3)
	grid.Attach(numLock, 58, 4, 3, 3)
	grid.Attach(divide, 61, 4, 3, 3)
	grid.Attach(multiply, 64, 4, 3, 3)
	grid.Attach(subtract, 67, 4, 3, 3)
	// 3RD ROW
	grid.Attach(tab, 0, 7, 5, 3)
	grid.Attach(q, 5, 7, 3, 3)
	grid.Attach(w, 8, 7, 3, 3)
	grid.Attach(e, 11, 7, 3, 3)
	grid.Attach(r, 14, 7, 3, 3)
	grid.Attach(t, 17, 7, 3, 3)
	grid.Attach(y, 20, 7, 3, 3)
	grid.Attach(u, 23, 7, 3, 3)
	grid.Attach(i, 26, 7, 3, 3)
	grid.Attach(o, 29, 7, 3, 3)
	grid.Attach(p, 32, 7, 3, 3)
	grid.Attach(leftBracket, 35, 7, 3, 3)
	grid.Attach(rightBracket, 38, 7, 3, 3)
	grid.Attach(backslash, 41, 7, 4, 3)
	grid.Attach(deleteKey, 47, 7, 3, 3)
	grid.Attach(end, 50, 7, 3, 3)
	grid.Attach(pageDown, 53, 7, 3, 3)
	grid.Attach(num7, 58, 7, 3, 3)
	grid.Attach(num8, 61, 7, 3, 3)
	grid.Attach(num9, 64, 7, 3, 3)
	grid.Attach(plus, 67, 7, 3, 6)
	// 4TH ROW
	grid.Attach(capsLock, 0, 10, 6, 3)
	grid.Attach(a, 6, 10, 3, 3)
	grid.Attach(s, 9, 10, 3, 3)
	grid.Attach(d, 12, 10, 3, 3)
	grid.Attach(f, 15, 10, 3, 3)
	grid.Attach(g, 18, 10, 3, 3)
	grid.Attach(h, 21, 10, 3, 3)
	grid.Attach(j, 24, 10, 3, 3)
	grid.Attach(k, 27, 10, 3, 3)
	grid.Attach(l, 30, 10, 3, 3)
	grid.Attach(colon, 33, 10, 3, 3)
	grid.Attach(quote, 36, 10, 3, 3)
	grid.Attach(enter, 39, 10, 6, 3)
	grid.Attach(num4, 58, 10, 3, 3)
	grid.Attach(num5, 61, 10, 3, 3)
	grid.Attach(num6, 64, 10, 3, 3)
	// 5TH ROW
	grid.Attach(leftShift, 0, 13, 7, 3)
	grid.Attach(z, 7, 13, 3, 3)
	grid.Attach(x, 10, 13, 3, 3)
	grid.Attach(c, 13, 13, 3, 3)
	grid.Attach(v, 16, 13, 3, 3)
	grid.Attach(b, 19, 13, 3, 3)
	grid.Attach(n, 22, 13, 3, 3)
	grid.Attach(m, 25, 13, 3, 3)
	grid.Attach(comma, 28, 13, 3, 3)
	grid.Attach(period, 31, 13, 3, 3)
	grid.Attach(forwardSlash, 34, 13, 3, 3)
	grid.Attach(rightShift, 37, 13, 8, 3)
	grid.Attach(upArrow, 50, 13, 3, 3)
	grid.Attach(num1, 58, 13, 3, 3)
	grid.Attach(num2, 61, 13, 3, 3)
	grid.Attach(num3, 64, 13, 3, 3)
	grid.Attach(numEnter, 67, 13, 3, 6)
	// 6TH ROW
	grid.Attach(leftControl, 0, 16, 4, 3)
	grid.Attach(leftWindows, 4, 16, 4, 3)
	grid.Attach(leftAlt, 8, 16, 4, 3)
	grid.Attach(spacebar, 12, 16, 17, 3)
	grid.Attach(rightAlt, 29, 16, 4, 3)
	grid.Attach(rightWindows, 33, 16, 4, 3)
	grid.Attach(menuKey, 37, 16, 4, 3)
	grid.Attach(rightControl, 41, 16, 4, 3)
	grid.Attach(leftArrow, 47, 16, 3, 3)
	grid.Attach(downArrow, 50, 16, 3, 3)
	grid.Attach(rightArrow, 53, 16, 3, 3)
	grid.Attach(num0, 58, 16, 6, 3)
	grid.Attach(decimal, 64, 16, 3, 3)
	// EMPTY SPACE
	grid.Attach(empty0, 45, 3, 2, 1)
	grid.Attach(empty1, 56, 3, 2, 1)
}
