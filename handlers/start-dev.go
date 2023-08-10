package handlers

import (
	"log"
	"net/http"
	hwi "github.com/3AM-Developer/dae/hw-interface" // Import path based on your previous code
) 

type StartDev struct {
	l *log.Logger
}

func NewStartDev(l *log.Logger) *StartDev {
	return &StartDev{l}
}

func (s *StartDev) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	deviceID := r.URL.Query().Get("id")
	if deviceID == "" {
		http.Error(rw, "Device ID is required", http.StatusBadRequest)
		return
	}

	device := queryDatabase(deviceID)

	hwi.StartDev(device)

	rw.WriteHeader(http.StatusOK)
}

// Fake database query function
func queryDatabase(deviceID string) hwi.Device {
	// In a real-world scenario, replace this with an actual database lookup
	// Here, we're just creating a dummy device for the sake of demonstration
	return hwi.Device{
		Name:  "Device_" + deviceID,
		IP:    "192.168.1." + deviceID, // Fake IP based on deviceID
		State: "Unknown",
	}
}