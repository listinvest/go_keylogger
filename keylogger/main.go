package main

import (
	"fmt"
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

	// stores active keys from most recent cycle
	var CurrentlyActiveKeys [103]bool
	for c := 0; c < 500; c++ {
		// start := time.Now()
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
				keyboard.TimesPressed[i]++
			}
		}

		// doesn't need to be running constantly
		time.Sleep(10 * time.Millisecond)

		// total := time.Since(start)
		// fmt.Println(total)
		// break

		// // test code
		// for i := 0; i < 256; i++ {
		// 	ret, _, _ := getAsyncKeyState.Call(uintptr(i))
		// 	if uint16(ret) != 0 {
		// 		fmt.Printf("%#X\n", i)
		// 	}
		// }
		// time.Sleep(10 * time.Millisecond)
	}

	for i, tp := range keyboard.TimesPressed {
		fmt.Printf("%s: %d\n", keyboard.VirtualKeyCodes[keyboard.Keys[i]], tp)
	}
}
