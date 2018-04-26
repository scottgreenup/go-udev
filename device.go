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

func (d Device) DevType() string {
    return C.GoString(C.udev_device_get_devtype(d.ptr))
}

func (d Device) Kernel() string {
    return C.GoString(C.udev_device_get_sysname(d.ptr))
}

func (d Device) DevNode() string {
    return C.GoString(C.udev_device_get_devnode(d.ptr))
}

func (d Device) SysPath() string {
    return C.GoString(C.udev_device_get_syspath(d.ptr))
}

func (d Device) SysNum() string {
    return C.GoString(C.udev_device_get_sysnum(d.ptr))
}

func (d Device) IsInitialized() bool {
	return C.udev_device_get_is_initialized(d.ptr) == 1
}

func (d Device) DevLinks() ListEntryArray{
	return NewListEntryArray(C.udev_device_get_devlinks_list_entry(d.ptr))
}

func (d Device) Properties() ListEntryArray{
	return NewListEntryArray(C.udev_device_get_properties_list_entry(d.ptr))
}

func (d Device) Tags() ListEntryArray{
	return NewListEntryArray(C.udev_device_get_tags_list_entry(d.ptr))
}

func (d Device) Attributes() ListEntryArray{
	return NewListEntryArray(C.udev_device_get_sysattr_list_entry(d.ptr))
}

/*
TODO
const char *udev_device_get_property_value(struct udev_device *udev_device, const char *key);
const char *udev_device_get_driver(struct udev_device *udev_device);
dev_t udev_device_get_devnum(struct udev_device *udev_device);
const char *udev_device_get_action(struct udev_device *udev_device);
unsigned long long int udev_device_get_seqnum(struct udev_device *udev_device);
unsigned long long int udev_device_get_usec_since_initialized(struct udev_device *udev_device);
const char *udev_device_get_sysattr_value(struct udev_device *udev_device, const char *sysattr);
int udev_device_set_sysattr_value(struct udev_device *udev_device, const char *sysattr, char *value);
int udev_device_has_tag(struct udev_device *udev_device, const char *tag);
*/

func (d Device) GetAttribute(name string) string {
    cName := C.CString(name)
    defer C.free(unsafe.Pointer(cName))
    return C.GoString(C.udev_device_get_sysattr_value(d.ptr, cName))
}

func (d Device) GetAttributes() ListEntryMap {
    entries := NewListEntryArray(C.udev_device_get_sysattr_list_entry(d.ptr))
    attributes := ListEntryMap{}
    for _, entry := range entries {
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

