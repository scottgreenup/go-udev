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
        dev.Free()
    }
    enumerate.Free()
}


func main() {
	manager := NewManager()
	defer manager.Free()

    //printKeyboards(manager)

    enumerate := manager.NewEnumerate()
    enumerate.AddMatchSubsystem("drm")
    for _, entry := range enumerate.GetList() {
        dev := manager.GetDeviceFromSystemPath(entry.Name)
        dev.Print()
        fmt.Println(strings.Repeat("-", 80))
    }

    enumerate.Free()
}
