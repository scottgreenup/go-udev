package main

// #include <libudev.h>
// #include <stdlib.h>
import "C"

import (
    "fmt"
    "unsafe"
)

type Device struct {
    ptr *C.struct_udev_device
}

func NewDevice(ptr *C.struct_udev_device) *Device {
    if ptr == nil {
        return nil
    }
    return &Device{
        ptr: ptr,
    }
}

func (d *Device) Free() {
    C.udev_device_unref(d.ptr)
    C.free(unsafe.Pointer(d.ptr))
}

func (d Device) Parent() *Device {
    return NewDevice(C.udev_device_get_parent(d.ptr))
}

func (d Device) DevPath() string {
    return C.GoString(C.udev_device_get_devpath(d.ptr))
}

func (d Device) Subsystem() string {
    return C.GoString(C.udev_device_get_subsystem(d.ptr))
}

func (d Device) Kernel() string {
    return C.GoString(C.udev_device_get_sysname(d.ptr))
}

func (d Device) DevNode() string {
    return C.GoString(C.udev_device_get_devnode(d.ptr))
}

func (d Device) GetAttribute(name string) string {
    cName := C.CString(name)
    defer C.free(unsafe.Pointer(cName))
    return C.GoString(C.udev_device_get_sysattr_value(d.ptr, cName))
}

func (d Device) GetAttributes() (map[string]string) {
    attributesList := C.udev_device_get_sysattr_list_entry(d.ptr)
    listEntries := iterateListEntry(attributesList)

    attributes := map[string]string{}

    for _, entry := range listEntries {
        attributes[entry.Name] = d.GetAttribute(entry.Name)
    }

    return attributes
}

func (d Device) Print() {
    fmt.Printf("DEVPATH: %s\n", d.DevPath())
    fmt.Printf("SUBSYSTEM: %s\n", d.Subsystem())
    fmt.Printf("SYSNAME/KERNEL: %s\n", d.Kernel())
    fmt.Printf("DEVNODE: %s\n", d.DevNode())

    for name, value := range d.GetAttributes() {
        fmt.Printf("  ATTR{%s}: %s\n", name, value)
    }

    /*
    parent := d.Parent()
    for parent != nil {
        fmt.Printf("\n||\n")
        parent.Print()
        parent = parent.Parent()
    }
    */
}

