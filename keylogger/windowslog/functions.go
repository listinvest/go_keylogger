package windowslog

import (
	"fmt"
	"log"
	"sync"
	"syscall"
	"unsafe"

	"github.com/hallazzang/go-windows-programming/pkg/win"
	"golang.org/x/sys/windows"
)

var (
	user32 = syscall.NewLazyDLL("user32.dll")
	// CreateWindowExW creates window
	CreateWindowExW = user32.NewProc("CreateWindowExW")
	// GetAsyncKeyState gets state of specified key
	GetAsyncKeyState = user32.NewProc("GetAsyncKeyState")
	// GetDesktopWindow gets hwnd to desktop window
	GetDesktopWindow = user32.NewProc("GetDesktopWindow")
	// PeekMessageW non blocking version of GetMessage
	PeekMessageW = user32.NewProc("PeekMessageW")
)

// PeekMessage uses the PeekMessageW proc to listen for certain events
func PeekMessage() {
	var msg win.MSG
	result, _, _ := PeekMessageW.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0, 0)
	if result != 0 {
		win.TranslateMessage(&msg)
		win.DispatchMessage(&msg)
	}
}

// UpdateCounts opens the count file, updates the count for each key, and closes it
func UpdateCounts(wg *sync.WaitGroup, curCounts []uint64) {
	//
	defer wg.Done()
	for i, count := range curCounts {
		fmt.Printf("%d: %d\n", i, count)
	}
}

// WndProc comment to shut up the linter
func WndProc(hWnd uintptr, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case notifyIconMsg:
		switch nmsg := win.LOWORD(uint32(lParam)); nmsg {
		case win.WM_LBUTTONDOWN:
			iconMenu()
		}
	case win.WM_DESTROY:
		win.PostQuitMessage(0)
	default:
		return win.DefWindowProc(hWnd, msg, wParam, lParam)
	}
	return 0
}

func iconMenu() {
	log.Println("Right clicked icon")
}

// https://github.com/hallazzang/go-windows-programming/blob/master/example/gui/notifyicon/main.go

// CreateMainWindow creates the handle for the window, its never shown, but is needed for the icon to be in the notification area, and registers the function that handles the
func CreateMainWindow() (uintptr, error) {
	hInstance := win.GetModuleHandle(nil)

	wndClass := windows.StringToUTF16Ptr("MyWindow")

	var wcex win.WNDCLASSEX
	wcex.CbSize = uint32(unsafe.Sizeof(wcex))
	wcex.Style = win.CS_HREDRAW | win.CS_VREDRAW
	wcex.LpfnWndProc = windows.NewCallback(WndProc)
	wcex.HInstance = hInstance
	wcex.HCursor = win.LoadCursor(0, win.MAKEINTRESOURCE(win.IDC_ARROW))
	wcex.HbrBackground = win.COLOR_WINDOW + 1
	wcex.LpszClassName = wndClass
	atom := win.RegisterClassEx(&wcex)
	if atom == 0 {
		return 0, win.GetLastError()
	}

	hWnd := win.CreateWindowEx(0, wndClass, windows.StringToUTF16Ptr("Go Keylogger"), win.WS_OVERLAPPEDWINDOW, win.CW_USEDEFAULT, win.CW_USEDEFAULT, 1, 1, 0, 0, hInstance, nil)
	if hWnd == win.NULL {
		log.Fatal(win.GetLastError())
	}
	win.ShowWindow(hWnd, win.SW_HIDE)

	return hWnd, nil
}

// LoadIconFromFile comment to shut up the linter
func LoadIconFromFile(name string) (uintptr, error) {
	hIcon := win.LoadImage(
		win.NULL,
		windows.StringToUTF16Ptr(name),
		win.IMAGE_ICON,
		0, 0,
		win.LR_DEFAULTSIZE|win.LR_LOADFROMFILE)
	if hIcon == win.NULL {
		return 0, win.GetLastError()
	}

	return hIcon, nil
}
