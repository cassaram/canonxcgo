package canonxcgo

// File contains system information commands / values

import (
	"net"
	"time"
)

func (d *Device) GetSystemStartTime() time.Time {
	val := d.getValue("s.epoch")
	const template string = "Mon, 2 Jan 2006 15:04:05 -0700"
	t, _ := time.Parse(template, val)
	return t
}

func (d *Device) GetSystemModelName() string {
	return d.getValue("s.hardware")
}

func (d *Device) GetSystemHardwareID() string {
	return d.getValue("s.hardware.id")
}

func (d *Device) GetSystemMACAddress() net.HardwareAddr {
	valStr := d.getValue("s.hardware.address")
	// valStr is a flat string (000085000000)

	// Inject dashes into string so it becomes a parsable address
	var valParsable string
	for idx, char := range valStr {
		if idx%2 == 0 && idx != 0 {
			valParsable += string('-')
		}
		valParsable += string(char)
	}

	// Parse and return
	address, _ := net.ParseMAC(valParsable)
	return address
}

func (d *Device) GetSystemFirmwareVersion() string {
	return d.getValue("s.firmware")
}

func (d *Device) GetSystemProtocolVersion() string {
	return d.getValue("s.protocol")
}
