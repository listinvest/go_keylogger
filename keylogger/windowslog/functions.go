package windowslog

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
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

var currentFile string = "key_totals.csv"
var oldFile string = "old_key_totals.csv"

// PeekMessage uses the PeekMessageW proc to listen for certain events
func PeekMessage() {
	var msg win.MSG
	result, _, _ := PeekMessageW.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0, 0)
	if result != 0 {
		win.TranslateMessage(&msg)
		win.DispatchMessage(&msg)
	}
}

func csvRead() [NumKeys]uint64 {
	var totals [NumKeys]uint64
	file, err := os.OpenFile(currentFile, os.O_RDWR, 0644)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatal("Unable to open csv file: ", err)
		} else {
			return totals
		}
	}

	i := 0
	csvReader := csv.NewReader(file)
	for {
		val, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Problem reading csv file: ", err)
		}
		num, err := strconv.ParseUint(val[1], 10, 64)
		if err != nil {
			log.Fatal("Problem parsing number in csv file: ", err)
		}
		totals[i] = num
		i++
	}

	file.Close()
	return totals
}

func addTotals(totals, current [NumKeys]uint64) [NumKeys]uint64 {
	var newTotal [NumKeys]uint64
	for i := 0; i < NumKeys; i++ {
		newTotal[i] = totals[i] + current[i]
	}
	return newTotal
}

func csvWrite(newTotals [NumKeys]uint64) {
	if err := os.Remove(oldFile); err != nil {
		if !os.IsNotExist(err) {
			log.Fatal("Unable to remove old csv file: ", err)
		}
	}

	if err := os.Rename(currentFile, oldFile); err != nil {
		if !os.IsNotExist(err) {
			log.Fatal("Unable to rename file: ", err)
		}
	}

	file, err := os.OpenFile(currentFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Error opening file for writing: ", err)
	}
	csvWriter := csv.NewWriter(file)

	for i := 0; i < NumKeys; i++ {
		record := []string{VirtualKeyCodes[Keys[i]], strconv.FormatUint(newTotals[i], 10)}
		if err = csvWriter.Write(record); err != nil {
			log.Fatal("Error writing to csv file: ", err)
		}
	}
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		log.Fatal("Error after flushing writer: ", err)
	}
	file.Close()
}

// UpdateCounts opens the count file, updates the count for each key, and closes it
func UpdateCounts(wg *sync.WaitGroup, current [NumKeys]uint64) {
	totals := csvRead()
	newTotals := addTotals(totals, current)
	csvWrite(newTotals)
	wg.Done()
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

// ResetCount resets the uint64 array to 0s
func ResetCount(currentCount *[NumKeys]uint64) {
	for i := 0; i < NumKeys; i++ {
		currentCount[i] = 0
	}
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
