package udev

// #include <libudev.h>
// #include <stdlib.h>
import "C"

import (
    "unsafe"
)

type ListEntry struct {
    Name string
    Value string
}

type ListEntryArray []ListEntry
type ListEntryMap   map[string]string

func NewListEntryArray(ptr *C.struct_udev_list_entry) ListEntryArray {
    le := ListEntryArray{}
    for ptr != nil {
        le = append(le, ListEntry{
            Name: C.GoString(C.udev_list_entry_get_name(ptr)),
            Value: C.GoString(C.udev_list_entry_get_value(ptr)),
        })

        ptr = C.udev_list_entry_get_next(ptr)
    }
    return le
}

func (le ListEntryArray) Map() ListEntryMap {
    m := ListEntryMap{}
    for _, entry := range le {
        m[entry.Name] = entry.Value
    }
    return m

}

type Enumerate struct {
    ptr *C.struct_udev_enumerate
}

func (e *Enumerate) Free() {
    C.udev_enumerate_unref(e.ptr)
}

func (e *Enumerate) AddMatchSubsystem(subsystemType string) {
    cSubsystemType := C.CString(subsystemType)
    defer C.free(unsafe.Pointer(cSubsystemType))
    C.udev_enumerate_add_match_subsystem(e.ptr, cSubsystemType)
}

func (e *Enumerate) AddMatchAttribute(name string, value string) {
    cName := C.CString(name)
    defer C.free(unsafe.Pointer(cName))
    cValue := C.CString(value)
    defer C.free(unsafe.Pointer(cValue))
    C.udev_enumerate_add_match_sysattr(e.ptr, cName, cValue)
}

func (e *Enumerate) GetList() ListEntryArray {
    C.udev_enumerate_scan_devices(e.ptr)
    return NewListEntryArray(C.udev_enumerate_get_list_entry(e.ptr))
}

