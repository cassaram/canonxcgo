package canonxcgo

// File contains system information commands / values

import (
	"net"
	"strconv"
	"time"
)

func (d *Device) GetSystemActions() []string {
	val := d.getValue("s.action.list")
	return parseListCSV(val)
}

func (d *Device) DoSystemAction(action string) {
	d.setValue("s.action", action)
}

func (d *Device) GetSystemStatus() string {
	return d.getValue("s.system.status")
}

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

func (d *Device) GetSystemMaxImageSize() string {
	return d.getValue("s.hardware.maxsize")
}

func (d *Device) GetSystemFirmwareVersion() string {
	return d.getValue("s.firmware")
}

func (d *Device) GetSystemProtocolVersion() string {
	return d.getValue("s.protocol")
}

func (d *Device) GetSystemMaxResolution() string {
	return d.getValue("s.maxsize")
}

func (d *Device) GetSystemPowerSource() string {
	return d.getValue("s.power.source")
}

func (d *Device) GetSystemPowerVolts() float64 {
	valStr := d.getValue("s.power.volt")
	val, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		// Handles nil string values
		return 0
	}
	return val
}

// Time left on battery
// equal to time.ParseDuration("0m") if N/A
func (d *Device) GetSystemPowerTime() time.Duration {
	valStr := d.getValue("s.power.minute")
	val, err := time.ParseDuration(valStr + "m")
	if err != nil {
		val, _ = time.ParseDuration("0m")
	}
	return val
}

// Battery percentage from 0 to 100
// 0 if N/A
func (d *Device) GetSystemPowerPercent() int {
	valStr := d.getValue("s.power.percent")
	val, err := strconv.ParseInt(valStr, 10, 64)
	if err != nil {
		return 0
	}
	return int(val)
}
