package keyboard

import "sort"

// NumKeys is the number of virtual keycodes being checked
const NumKeys = 103

// VirtualKeyCodes is a map from the int code to a string name of the code
var VirtualKeyCodes = map[int]string{
	0x01: "Left Mouse",
	0x02: "Right Mouse",
	0x04: "Middle Mouse",
	0x08: "Backspace",
	0x09: "Tab",
	0x0D: "Enter",
	0x12: "Alt",
	0x13: "Pause",
	0x14: "Caps Lock",
	0x1B: "Escape",
	0x20: "Spacebar",
	0x21: "Page Up",
	0x22: "Page Down",
	0x23: "End",
	0x24: "Home",
	0x25: "Left Arrow",
	0x26: "Up Arrow",
	0x27: "Right Arrow",
	0x28: "Down Arrow",
	0x2C: "Print Screen",
	0x2D: "Insert",
	0x2E: "Delete",
	0x30: "0",
	0x31: "1",
	0x32: "2",
	0x33: "3",
	0x34: "4",
	0x35: "5",
	0x36: "6",
	0x37: "7",
	0x38: "8",
	0x39: "9",
	0x41: "A",
	0x42: "B",
	0x43: "C",
	0x44: "D",
	0x45: "E",
	0x46: "F",
	0x47: "G",
	0x48: "H",
	0x49: "I",
	0x4A: "J",
	0x4B: "K",
	0x4C: "L",
	0x4D: "M",
	0x4E: "N",
	0x4F: "O",
	0x50: "P",
	0x51: "Q",
	0x52: "R",
	0x53: "S",
	0x54: "T",
	0x55: "U",
	0x56: "V",
	0x57: "W",
	0x58: "X",
	0x59: "Y",
	0x5A: "Z",
	0x5B: "Left Windows Key",
	0x5C: "Right Windows Key",
	0x5D: "Menu Key",
	0x60: "Numpad 0",
	0x61: "Numpad 1",
	0x62: "Numpad 2",
	0x63: "Numpad 3",
	0x64: "Numpad 4",
	0x65: "Numpad 5",
	0x66: "Numpad 6",
	0x67: "Numpad 7",
	0x68: "Numpad 8",
	0x69: "Numpad 9",
	0x6A: "Multiply",
	0x6B: "Add",
	0x6D: "Subtract",
	0x6E: "Decimal",
	0x6F: "Divide",
	0x70: "F1",
	0x71: "F2",
	0x72: "F3",
	0x73: "F4",
	0x74: "F5",
	0x75: "F6",
	0x76: "F7",
	0x77: "F8",
	0x78: "F9",
	0x79: "F10",
	0x7A: "F11",
	0x7B: "F12",
	0x90: "Num Lock",
	0x91: "Scroll Lock",
	0xA0: "Left Shift",
	0xA1: "Right Shift",
	0xA2: "Left Control",
	0xA3: "Right Control",
	0xA4: "Left Alt",
	0xA5: "Right Alt",
	0xBA: "Semi-Colon/Colon",
	0xBB: "Equal/Plus",
	0xBC: "Comma/Left Angle Bracket",
	0xBD: "Dash/Underscore",
	0xBE: "Period/Right Angle Bracket",
	0xBF: "Forward Slash/Question Mark",
	0xC0: "Backtick/Tilde",
	0xDB: "Left Bracket/Left Brace",
	0xDC: "Vertical Bar/Back Slash",
	0xDD: "Right Bracket/ Right Brace",
	0xDE: "Single Quote/Double Quote",
}

// IsKeyActive stores whether the key is actively being pressed or not
var IsKeyActive [NumKeys]bool

// TimesPressed counts the number of times the key at each index is pressed
var TimesPressed [NumKeys]uint64

// KeyIndex is a map that takes in the hex and returns the index of that hex in the array
var KeyIndex = make(map[int]int)

// Keys is an int slice of the relevant keys
var Keys []int

func init() {
	// appends the key to the slice and then sorts it
	for key := range VirtualKeyCodes {
		Keys = append(Keys, key)
	}
	sort.Ints(Keys)

	// adds the key and index to the [int]int map
	for i, key := range Keys {
		KeyIndex[key] = i
	}
}
