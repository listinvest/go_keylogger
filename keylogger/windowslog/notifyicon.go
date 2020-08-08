package windowslog

import (
	"crypto/rand"
	"errors"
	"unsafe"

	"github.com/hallazzang/go-windows-programming/pkg/win"
	"golang.org/x/sys/windows"
)

// https://github.com/hallazzang/go-windows-programming/blob/master/example/gui/notifyicon/notifyicon.go

const notifyIconMsg = win.WM_APP + 1

var errShellNotifyIcon = errors.New("Shell_NotifyIcon error")

// NotifyIcon holds the data for something
type NotifyIcon struct {
	hwnd uintptr
	guid win.GUID
}

// NewGUID creates a guid for the icon, probably
func NewGUID() win.GUID {
	var buf [16]byte
	rand.Read(buf[:])
	return *(*win.GUID)(unsafe.Pointer(&buf[0]))
}

// NewNotifyIcon returns the notifyIcon for the keylogger
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
func (ni *NotifyIcon) NewData() *win.NOTIFYICONDATA {
	var nid win.NOTIFYICONDATA
	nid.CbSize = uint32(unsafe.Sizeof(nid))
	nid.UFlags = win.NIF_GUID
	nid.HWnd = ni.hwnd
	nid.GuidItem = ni.guid
	return &nid
}

// SetIcon comment to shut up the linter
func (ni *NotifyIcon) SetIcon(hIcon uintptr) error {
	data := ni.NewData()
	data.UFlags |= win.NIF_ICON
	data.HIcon = hIcon
	if win.Shell_NotifyIcon(win.NIM_MODIFY, data) == win.FALSE {
		return errShellNotifyIcon
	}
	return nil
}

// Dispose comment to shut up the linter
func (ni *NotifyIcon) Dispose() {
	win.Shell_NotifyIcon(win.NIM_DELETE, ni.NewData())
}

// SetTooltip comment to shut up the linter
func (ni *NotifyIcon) SetTooltip(tooltip string) error {
	data := ni.NewData()
	data.UFlags |= win.NIF_TIP
	copy(data.SzTip[:], windows.StringToUTF16(tooltip))
	if win.Shell_NotifyIcon(win.NIM_MODIFY, data) == win.FALSE {
		return errShellNotifyIcon
	}
	return nil
}

// ShowNotificationWithIcon comment to shut up the linter
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
