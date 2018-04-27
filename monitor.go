package udev

// #include <libudev.h>
// #include <stdlib.h>
import "C"

import (
    "unsafe"
)

type Monitor struct {
    ptr *C.struct_udev_monitor
}

func NewMonitor(ptr *C.struct_udev_monitor) *Monitor {
    if ptr == nil {
        return nil
    }
    return &Monitor{
        ptr: ptr,
    }
}

func (m *Monitor) Free() {
    C.udev_monitor_unref(m.ptr)
}

func (m *Monitor) GetManager() *Manager {
	return &Manager{
		ptr: C.udev_monitor_get_udev(m.ptr),
	}
}

func (m *Monitor) Enable() bool {
	return C.udev_monitor_enable_receiving(m.ptr) >= 0
}

func (m *Monitor) SetReceiveBufferSize(size int) bool {
	return C.udev_monitor_set_receive_buffer_size(m.ptr, _Ctype_int(size)) >= 0
}

func (m *Monitor) GetFileDesc() int {
	return int(C.udev_monitor_get_fd(m.ptr))
}

func (m *Monitor) ReceiveDevice() *Device {
	return NewDevice(C.udev_monitor_receive_device(m.ptr))
}

func (m *Monitor) AddFilterToMatchSubsystemDevType(subsystem string, devType string) bool {
	cSubsystem := C.CString(subsystem)
	defer C.free(unsafe.Pointer(cSubsystem))
	cDevType := C.CString(devType)
	defer C.free(unsafe.Pointer(cDevType))
	return C.udev_monitor_filter_add_match_subsystem_devtype(m.ptr, cSubsystem, cDevType) >= 0
}

func (m *Monitor) AddFilterMatchTag(tag string) bool {
	cTag := C.CString(tag)
	defer C.free(unsafe.Pointer(cTag))
	return C.udev_monitor_filter_add_match_tag(m.ptr, cTag) >= 0
}

func (m *Monitor) UpdateFilter() bool {
	return C.udev_monitor_filter_update(m.ptr) >= 0
}

func (m *Monitor) RemoveFilter() bool {
	return C.udev_monitor_filter_remove(m.ptr) >= 0
}
