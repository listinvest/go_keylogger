package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/hallazzang/go-windows-programming/pkg/win"
	"github.com/reinoj/go_keylogger/keylogger/windowslog"
)

func main() {
	// sleep at the beginning to get rid of some initial erroneous data
	time.Sleep(1 * time.Second)

	hWnd, err := windowslog.CreateMainWindow()
	if err != nil {
		log.Fatal("Problem creating window: ", err)
	}

	nid, err := windowslog.NewNotifyIcon(hWnd)
	if err != nil {
		log.Fatal("Problem in NewNotifyIcon: ", err)
	}
	defer nid.Dispose()

	hIcon, err := windowslog.LoadIconFromFile("icon.ico")
	if err != nil {
		log.Fatal("Problem loading icon: ", err)
	}
	defer win.DestroyIcon(hIcon)

	nid.SetIcon(hIcon)
	nid.SetTooltip("Go Keylogger")

	var wg sync.WaitGroup

	// time.Sleep(10 * time.Second)
	// stores active keys from most recent cycle
	var CurrentlyActiveKeys [windowslog.NumKeys]bool

	var currentCount [windowslog.NumKeys]uint64
	resetCount := true
	lastUpdate := time.Now()

	TESTCOUNT := 0

	for {
		// for c := 0; c < 500; c++ {
		// start := time.Now()
		windowslog.PeekMessage()

		if resetCount {
			// currentCount = make([]uint64, windowslog.NumKeys)
			windowslog.ResetCount(&currentCount)
			resetCount = false
			lastUpdate = time.Now()
		}
		for _, key := range windowslog.Keys {
			// ret is a short where if the most or least significant bits are set then the key is being pressed or was unpressed since the last check
			ret, _, _ := windowslog.GetAsyncKeyState.Call(uintptr(key))
			if uint16(ret) != 0 {
				CurrentlyActiveKeys[windowslog.KeyIndex[key]] = true
			} else {
				CurrentlyActiveKeys[windowslog.KeyIndex[key]] = false
			}
		}
		// go through each key's bool value and set/unset the main bool array and increment the times pressed counter
		for i := 0; i < windowslog.NumKeys; i++ {
			if CurrentlyActiveKeys[i] && !windowslog.IsKeyActive[i] {
				windowslog.IsKeyActive[i] = true
			} else if !CurrentlyActiveKeys[i] && windowslog.IsKeyActive[i] {
				windowslog.IsKeyActive[i] = false
				currentCount[i]++
			}
		}
		// 1 minute in nanoseconds
		// if time.Since(lastUpdate) >= 60_000_000_000 {
		if time.Since(lastUpdate) >= 15_000_000_000 {
			wg.Add(1)
			go windowslog.UpdateCounts(&wg, currentCount)
			resetCount = true
			TESTCOUNT++
			if TESTCOUNT > 2 {
				break
			}
		}
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Printf("Waiting for all updates to finish...\n")
	wg.Wait()
	fmt.Printf("All updates finished.\n")
}
