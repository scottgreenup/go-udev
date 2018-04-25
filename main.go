package main

import (
    "fmt"
    "strings"
)

// #cgo LDFLAGS: -ludev
// #include <libudev.h>
// #include <stdlib.h>
import "C"

func printKeyboards(manager *Manager) {
    enumerate := manager.NewEnumerate()
    enumerate.AddMatchSubsystem("usb")
    enumerate.AddMatchAttribute("interface", "Keyboard")

    for _, entry := range enumerate.GetList() {
        dev := manager.GetDeviceFromSystemPath(entry.Name)
        dev.Print()
        fmt.Println(strings.Repeat("-", 80))

        //C.udev_device_unref(dev)
    }
}


func main() {
    manager := NewManager()
    //printKeyboards(manager)

    enumerate := manager.NewEnumerate()
    enumerate.AddMatchSubsystem("drm")
    for _, entry := range enumerate.GetList() {
        dev := manager.GetDeviceFromSystemPath(entry.Name)
        dev.Print()
        fmt.Println(strings.Repeat("-", 80))
    }
}
