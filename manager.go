package main

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
