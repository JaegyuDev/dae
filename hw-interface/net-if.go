package hw_interface

import "fmt"

type DeviceState int

const (
	Unknown DeviceState = iota
	Offline
	Online
	Error
)

type App struct {
	Name       string //Human readable name
	AppID      string //ID for records
	IsDeployed bool   //Whether the app is deployed to a device
	Dev        Device //Device the app is deployed (nullable)
	FLoc       string //Location in the filesystem (must be within /opt/das/) (non-nullable, and persistant)
}

type Device struct {
	Name  string //Human readable name
	DevID string //ID for records
	IP    string //netaddr for device (non-nullable, and persistant)
	State string //DeviceState
	Apps  []App  //Apps deployed here
}

func StartDev(dev Device) {
	fmt.Printf("Device %s was started!\n", dev.Name)
}
