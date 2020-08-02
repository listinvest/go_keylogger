package keyboard

import (
	"crypto/rand"
	"syscall"
	"unsafe"

	"github.com/hallazzang/go-windows-programming/pkg/win"
)

var (
	user32  = syscall.NewLazyDLL("user32.dll")
	shell32 = syscall.NewLazyDLL("Shell32.dll")
	// DestroyIcon destroys icon and frees memory
	DestroyIcon = user32.NewProc("DestroyIcon")
	// GetAsyncKeyState gets state of specified key
	GetAsyncKeyState = user32.NewProc("GetAsyncKeyState")
	// GetDesktopWindow gets hwnd to desktop window
	GetDesktopWindow = user32.NewProc("GetDesktopWindow")
	// LoadImageW loads icon
	LoadImageW = user32.NewProc("LoadImageW")
	// ShellNotifyIconW adds icon to status area
	ShellNotifyIconW = shell32.NewProc("Shell_NotifyIconW")
)

// NewGUID creates a guid for the icon, probably
// https://github.com/hallazzang/go-windows-programming/blob/master/example/gui/notifyicon/notifyicon.go
func NewGUID() win.GUID {
	var buf [16]byte
	rand.Read(buf[:])
	return *(*win.GUID)(unsafe.Pointer(&buf[0]))
}
