package main

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

func iterateListEntry(ptr *C.struct_udev_list_entry) []ListEntry {

    le := []ListEntry{}

    for ptr != nil {
        le = append(le, ListEntry{
            Name: C.GoString(C.udev_list_entry_get_name(ptr)),
            Value: C.GoString(C.udev_list_entry_get_value(ptr)),
        })

        ptr = C.udev_list_entry_get_next(ptr)
    }

    return le
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

func (e *Enumerate) GetList() []ListEntry {
    C.udev_enumerate_scan_devices(e.ptr)

    list_entry := C.udev_enumerate_get_list_entry(e.ptr)

    if list_entry == nil {
        return []ListEntry{}
    }

    return iterateListEntry(list_entry)
}

