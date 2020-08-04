package main

import (
	"fmt"
	"log"
	"sync"
	"time"
	"unsafe"

	"github.com/hallazzang/go-windows-programming/pkg/win"
	"github.com/reinoj/go_keylogger/keylogger/keyboard"
	"golang.org/x/sys/windows"
)

// func clickHandle() {
// 	for {
// 		fmt.Printf("got here\n")
// 		var msg win.MSG
// 		for win.GetMessage(&msg, 0, 0, 0) != 0 {
// 			win.TranslateMessage(&msg)
// 			win.DispatchMessage(&msg)
// 		}
// 	}
// }

func main() {
	// sleep at the beginning to get rid of some initial erroneous data
	time.Sleep(1 * time.Second)

	hInstance := win.GetModuleHandle(nil)

	wndClass := windows.StringToUTF16Ptr("MyWindow")

	var wcex win.WNDCLASSEX
	wcex.CbSize = uint32(unsafe.Sizeof(wcex))
	wcex.Style = win.CS_HREDRAW | win.CS_VREDRAW
	wcex.LpfnWndProc = windows.NewCallback(keyboard.WndProc)
	wcex.HInstance = hInstance
	wcex.HCursor = win.LoadCursor(0, win.MAKEINTRESOURCE(win.IDC_ARROW))
	wcex.HbrBackground = win.COLOR_WINDOW + 1
	wcex.LpszClassName = wndClass
	atom := win.RegisterClassEx(&wcex)
	if atom == 0 {
		log.Println(win.GetLastError())
	}

	hWnd := win.CreateWindowEx(0, wndClass, windows.StringToUTF16Ptr("Go Keylogger"), win.WS_OVERLAPPEDWINDOW, win.CW_USEDEFAULT, win.CW_USEDEFAULT, 100, 100, 0, 0, hInstance, nil)
	if hWnd == win.NULL {
		log.Fatal(win.GetLastError())
	}
	win.ShowWindow(hWnd, win.SW_HIDE)

	nid, err := keyboard.NewNotifyIcon(hWnd)
	if err != nil {
		log.Println(err)
	}
	defer nid.Dispose()

	hIcon, err := keyboard.LoadIconFromFile("icon.ico")
	if err != nil {
		log.Println(err)
	}
	defer win.DestroyIcon(hIcon)

	nid.SetIcon(hIcon)
	nid.SetTooltip("Go Keylogger")

	var wg sync.WaitGroup

	// time.Sleep(10 * time.Second)
	// stores active keys from most recent cycle
	var CurrentlyActiveKeys [keyboard.NumKeys]bool

	var currentCount []uint64
	var resetCount bool = true
	var lastUpdate time.Time = time.Now()

	// go clickHandle()

	for c := 0; c < 500; c++ {
		// need to move this message loop somehow so it doesn't block the rest of the function
		var msg win.MSG
		result := win.GetMessage(&msg, 0, 0, 0)
		if result != 0 {
			win.TranslateMessage(&msg)
			win.DispatchMessage(&msg)
		}
		// for {
		// start := time.Now()

		if resetCount {
			currentCount = make([]uint64, keyboard.NumKeys)
			resetCount = false
			lastUpdate = time.Now()
		}
		for _, key := range keyboard.Keys {
			// ret is a short where if the most or least significant bits are set then the key is being pressed or was unpressed since the last check
			ret, _, _ := keyboard.GetAsyncKeyState.Call(uintptr(key))
			if uint16(ret) != 0 {
				CurrentlyActiveKeys[keyboard.KeyIndex[key]] = true
			} else {
				CurrentlyActiveKeys[keyboard.KeyIndex[key]] = false
			}
		}
		// go through each key's bool value and set/unset the main bool array and increment the times pressed counter
		for i := 0; i < keyboard.NumKeys; i++ {
			if CurrentlyActiveKeys[i] && !keyboard.IsKeyActive[i] {
				keyboard.IsKeyActive[i] = true
			} else if !CurrentlyActiveKeys[i] && keyboard.IsKeyActive[i] {
				keyboard.IsKeyActive[i] = false
				currentCount[i]++
			}
		}
		// 1 minute in nanoseconds
		if time.Since(lastUpdate) >= 60_000_000_000 {
			fmt.Println("One minute passed")
			wg.Add(1)
			go keyboard.UpdateCounts(&wg, currentCount)
			resetCount = true
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Printf("Waiting for all updates to finish...\n")
	wg.Wait()
	// _, _, _ = keyboard.DestroyIcon.Call(icon)
	// _, _, _ = keyboard.ShellNotifyIconW.Call(win.NIM_DELETE, uintptr(unsafe.Pointer(&nid)))
	fmt.Printf("All updates finished.\n")
}
