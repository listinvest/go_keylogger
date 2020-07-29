package main

import (
	"fmt"
	"sync"
	"syscall"
	"time"

	"github.com/reinoj/go_keylogger/keylogger/keyboard"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	getAsyncKeyState = user32.NewProc("GetAsyncKeyState")
)

func main() {
	// sleep at the beginning to get rid of some initial erroneous data
	time.Sleep(2 * time.Second)

	var wg sync.WaitGroup

	// stores active keys from most recent cycle
	var CurrentlyActiveKeys [keyboard.NumKeys]bool

	var currentCount []uint64
	var resetCount bool = true
	var lastUpdate time.Time = time.Now()

	// for c := 0; c < 500; c++ {
	for {
		// start := time.Now()

		if resetCount {
			currentCount = make([]uint64, keyboard.NumKeys)
			resetCount = false
		}
		for _, key := range keyboard.Keys {
			// ret is a short where if the most or least significant bits are set then the key is being pressed or was unpressed since the last check
			ret, _, _ := getAsyncKeyState.Call(uintptr(key))
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
	fmt.Printf("All updates finished.\n")
}
