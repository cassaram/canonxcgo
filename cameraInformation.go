package canonxcgo

// File contains camera information commands / values

import "strconv"

// Should always return "1"
func (d *Device) GetCameraNumber() int {
	key := "c"
	response := d.getValue(key)
	number, _ := strconv.Atoi(response)
	return number
}

func (d *Device) SetCameraNumber(number int) {
	key := "c"
	valueString := strconv.Itoa(number)
	d.setValue(key, valueString)
}

// Should always return "1"
func (d *Device) GetCameraCount() int {
	key := "c.count"
	response := d.getValue(key)
	number, _ := strconv.Atoi(response)
	return number
}

func (d *Device) GetCameraType(ID int) string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".type"
	response := d.getValue(key)
	return response
}

func (d *Device) GetCameraStatus(ID int) string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".status"
	response := d.getValue(key)
	return response
}

func (d *Device) GetCameraCaptureMode(ID int) string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".shooting"
	response := d.getValue(key)
	return response
}

func (d *Device) SetCameraCaptureMode(ID int, mode string) {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".shooting"
	d.setValue(key, mode)
}

func (d *Device) GetCameraCaptureModeList(ID int) []string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".shooting.list"
	response := d.getValue(key)
	values := parseListCSV(response)
	return values
}

func (d *Device) GetCameraPlatformStatus(ID int) string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".platform.status"
	response := d.getValue(key)
	return response
}

func (d *Device) GetCameraPlatformError(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".platform.error"
	response := d.getValue(key)
	errorCode, _ := strconv.Atoi(response)
	return errorCode
}

func (d *Device) GetCameraName(ID int) string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".name.utf8"
	response := d.getValue(key)
	return response
}
