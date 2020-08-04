package keyboard

import (
	"crypto/rand"
	"errors"
	"log"
	"syscall"
	"unsafe"

	"github.com/hallazzang/go-windows-programming/pkg/win"
	"golang.org/x/sys/windows"
)

var (
	user32 = syscall.NewLazyDLL("user32.dll")
	// GetAsyncKeyState gets state of specified key
	GetAsyncKeyState = user32.NewProc("GetAsyncKeyState")
	// GetDesktopWindow gets hwnd to desktop window
	GetDesktopWindow = user32.NewProc("GetDesktopWindow")
	// CreateWindowExW creates window
	CreateWindowExW = user32.NewProc("CreateWindowExW")
)

const notifyIconMsg = win.WM_APP + 1

var errShellNotifyIcon = errors.New("Shell_NotifyIcon error")

// NotifyIcon holds the data for something
type NotifyIcon struct {
	hwnd uintptr
	guid win.GUID
}

// NewGUID creates a guid for the icon, probably
// https://github.com/hallazzang/go-windows-programming/blob/master/example/gui/notifyicon/notifyicon.go
func NewGUID() win.GUID {
	var buf [16]byte
	rand.Read(buf[:])
	return *(*win.GUID)(unsafe.Pointer(&buf[0]))
}

// NewNotifyIcon returns the notifyIcon for the keylogger
// https://github.com/hallazzang/go-windows-programming/blob/master/example/gui/notifyicon/notifyicon.go
func NewNotifyIcon(hwnd uintptr) (*NotifyIcon, error) {
	ni := &NotifyIcon{
		hwnd: hwnd,
		guid: NewGUID(),
	}
	data := ni.NewData()
	data.UFlags |= win.NIF_MESSAGE
	data.UCallbackMessage = notifyIconMsg
	if win.Shell_NotifyIcon(win.NIM_ADD, data) == win.FALSE {
		return nil, errShellNotifyIcon
	}
	return ni, nil
}

// NewData returns the NOTIFYICONDATA of the keylogger
// https://github.com/hallazzang/go-windows-programming/blob/master/example/gui/notifyicon/notifyicon.go
func (ni *NotifyIcon) NewData() *win.NOTIFYICONDATA {
	var nid win.NOTIFYICONDATA
	nid.CbSize = uint32(unsafe.Sizeof(nid))
	nid.UFlags = win.NIF_GUID
	nid.HWnd = ni.hwnd
	nid.GuidItem = ni.guid
	return &nid
}

func (ni *NotifyIcon) SetIcon(hIcon uintptr) error {
	data := ni.NewData()
	data.UFlags |= win.NIF_ICON
	data.HIcon = hIcon
	if win.Shell_NotifyIcon(win.NIM_MODIFY, data) == win.FALSE {
		return errShellNotifyIcon
	}
	return nil
}

func (ni *NotifyIcon) Dispose() {
	win.Shell_NotifyIcon(win.NIM_DELETE, ni.NewData())
}

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

func (ni *NotifyIcon) SetTooltip(tooltip string) error {
	data := ni.NewData()
	data.UFlags |= win.NIF_TIP
	copy(data.SzTip[:], windows.StringToUTF16(tooltip))
	if win.Shell_NotifyIcon(win.NIM_MODIFY, data) == win.FALSE {
		return errShellNotifyIcon
	}
	return nil
}

func (ni *NotifyIcon) ShowNotificationWithIcon(title, text string, hIcon uintptr) error {
	data := ni.NewData()
	data.UFlags |= win.NIF_INFO
	copy(data.SzInfoTitle[:], windows.StringToUTF16(title))
	copy(data.SzInfo[:], windows.StringToUTF16(text))
	data.DwInfoFlags = win.NIIF_USER | win.NIIF_LARGE_ICON
	if win.Shell_NotifyIcon(win.NIM_MODIFY, data) == win.FALSE {
		return errShellNotifyIcon
	}
	return nil
}

func iconMenu() {
	log.Println("Right clicked icon")
}

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
