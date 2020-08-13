package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
	"github.com/reinoj/go_keylogger/keylogger_gui/setupgrid"
)

func main() {
	setupgrid.GetImage()
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
	// 1765 x 395 window size

	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}

	grid.SetMarginTop(10)
	grid.SetMarginBottom(10)
	grid.SetMarginStart(10)
	grid.SetMarginEnd(10)
	grid.SetRowSpacing(5)
	grid.SetColumnSpacing(5)
	grid.SetColumnHomogeneous(true)
	grid.SetRowHomogeneous(true)

	setupgrid.SetupKeyboard(grid)

	img, err := gtk.ImageNewFromFile("heatmap.png")
	if err != nil {
		log.Fatal("Unable to load image:", err)
	}

	winGrid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}

	winGrid.Attach(img, 0, 0, 1, 1)
	winGrid.Attach(grid, 0, 0, 1, 1)
	win.Add(winGrid)
	win.ShowAll()

	width, height := win.GetSize()
	fmt.Printf("width: %d\theight: %d\n", width, height)

	gtk.Main()
}
