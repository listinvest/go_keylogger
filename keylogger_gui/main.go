package main

import (
	"log"

	"github.com/reinoj/go_keylogger/keylogger_gui/setupgrid"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Keylogger GUI")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}

	setupgrid.SetupKeyboard(grid)

	grid.SetRowSpacing(5)
	grid.SetColumnSpacing(5)

	win.Add(grid)
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}
