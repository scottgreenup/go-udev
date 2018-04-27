package udev

// #include <libudev.h>
// #include <stdlib.h>
import "C"

import (
    "fmt"
    "unsafe"
)

type Manager struct {
    ptr *C.struct_udev
}

func NewManager() *Manager {
    udev := C.udev_new()

    if udev == nil {
        panic(fmt.Errorf("udev_new() failed"))
    }

    return &Manager{
        ptr: udev,
    }
}

func (m *Manager) Free() {
    C.udev_unref(m.ptr)
}

func (m *Manager) NewEnumerate() *Enumerate {
    ptr := C.udev_enumerate_new(m.ptr)
    if ptr == nil {
        return nil
    }
    return &Enumerate{
        ptr: ptr,
    }
}

func (m *Manager) GetDeviceFromSystemPath(path string) *Device {
    cPath := C.CString(path)
    defer C.free(unsafe.Pointer(cPath))

    return NewDevice(C.udev_device_new_from_syspath(m.ptr, cPath))
}

func (m *Manager) GetMonitorFromNetlink(name string) *Monitor {
    cName := C.CString(name)
    defer C.free(unsafe.Pointer(cName))
	mon := C.udev_monitor_new_from_netlink(m.ptr, cName)

	if mon == nil {
		return nil
	}

	return NewMonitor(mon)
}
