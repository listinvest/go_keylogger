package main

import (
	"fmt"
	"syscall"
	"time"

	"keylogger/keylogger/keyboard"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	getAsyncKeyState = user32.NewProc("GetAsyncKeyState")
)

func main() {
	// sleep at the beginning to get rid of some initial erroneous data
	time.Sleep(2 * time.Second)

	// stores active keys from most recent cycle
	var CurrentlyActiveKeys [103]bool
	for c := 0; c < 500; c++ {
		// start := time.Now()
		for _, key := range keyboard.Keys {
			ret, _, _ := getAsyncKeyState.Call(uintptr(key))
			if uint16(ret) != 0 {
				CurrentlyActiveKeys[keyboard.KeyIndex[key]] = true
			} else {
				CurrentlyActiveKeys[keyboard.KeyIndex[key]] = false
			}
		}

		for i := 0; i < keyboard.NumKeys; i++ {
			if CurrentlyActiveKeys[i] && !keyboard.IsKeyActive[i] {
				keyboard.IsKeyActive[i] = true
			} else if !CurrentlyActiveKeys[i] && keyboard.IsKeyActive[i] {
				keyboard.IsKeyActive[i] = false
				keyboard.TimesPressed[i]++
			}
		}
		// total := time.Since(start)
		// fmt.Println(total)

		// doesn't need to be running constantly
		time.Sleep(10 * time.Millisecond)
		// break
	}

	for i, tp := range keyboard.TimesPressed {
		fmt.Printf("%#x: %d\n", keyboard.Keys[i], tp)
	}
}

// package main

// import (
// 	"github.com/gotk3/gotk3/gtk"
// 	"log"
// )

// func main() {
// 	// Initialize GTK without parsing any command line arguments.
// 	gtk.Init(nil)

// 	// Create a new toplevel window, set its title, and connect it to the
// 	// "destroy" signal to exit the GTK main loop when it is destroyed.
// 	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
// 	if err != nil {
// 		log.Fatal("Unable to create window:", err)
// 	}
// 	win.SetTitle("Simple Example")
// 	win.Connect("destroy", func() {
// 		gtk.MainQuit()
// 	})

// 	// Create a new label widget to show in the window.
// 	l, err := gtk.LabelNew("Hello, gotk3!")
// 	if err != nil {
// 		log.Fatal("Unable to create label:", err)
// 	}

// 	// Add the label to the window.
// 	win.Add(l)

// 	// Set the default window size.
// 	win.SetDefaultSize(800, 600)

// 	// Recursively show all widgets contained in this window.
// 	win.ShowAll()

// 	// Begin executing the GTK main loop.  This blocks until
// 	// gtk.MainQuit() is run.
// 	gtk.Main()
// }
