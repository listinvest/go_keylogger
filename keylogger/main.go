package main

import (
	"fmt"
	"sync"
	"time"
	"unsafe"

	"golang.org/x/sys/windows"

	"github.com/hallazzang/go-windows-programming/pkg/win"
	"github.com/reinoj/go_keylogger/keylogger/keyboard"
)

func getNotifyIconData() *win.NOTIFYICONDATA {
	desktopHWnd, _, _ := keyboard.GetDesktopWindow.Call()
	var nid win.NOTIFYICONDATA
	nid.CbSize = uint32(unsafe.Sizeof(nid))
	nid.UFlags = win.NIF_GUID
	nid.HWnd = desktopHWnd
	nid.GuidItem = keyboard.NewGUID()
	return &nid
}

func main() {
	// sleep at the beginning to get rid of some initial erroneous data
	time.Sleep(2 * time.Second)

	nid := getNotifyIconData()
	icon, _, _ := keyboard.LoadImageW.Call(win.NULL, uintptr(unsafe.Pointer(windows.StringToUTF16Ptr("icon32.ico"))), win.IMAGE_ICON, 0, 0, win.LR_DEFAULTSIZE)
	nid.HIcon = icon
	nid.UFlags |= win.NIF_ICON
	success, _, _ := keyboard.ShellNotifyIconW.Call(uintptr(win.NIM_ADD), uintptr(unsafe.Pointer(&nid)))
	fmt.Printf("%#X\n", success)

	var wg sync.WaitGroup

	// stores active keys from most recent cycle
	var CurrentlyActiveKeys [keyboard.NumKeys]bool

	var currentCount []uint64
	var resetCount bool = true
	var lastUpdate time.Time = time.Now()

	for c := 0; c < 500; c++ {
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
	_, _, _ = keyboard.DestroyIcon.Call(icon)
	_, _, _ = keyboard.ShellNotifyIconW.Call(win.NIM_DELETE, uintptr(unsafe.Pointer(&nid)))
	fmt.Printf("All updates finished.\n")
}
