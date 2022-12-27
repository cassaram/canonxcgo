package canonxcgo

import "strconv"

// File contains video information commands / values

func (d *Device) GetVideoStreamParameters() []string {
	key := "v.list"
	response := d.getValue(key)
	list := parseListCSV(response)
	return list
}

// Unit is in kbs
func (d *Device) GetH264TargetBitrate() string {
	key := "v.h264.cbr"
	response := d.getValue(key)
	return response
}

func (d *Device) GetVideoStreamMaxSize() string {
	return d.getValue("w.maxsize")
}

func (d *Device) GetVideoStreamCount() int {
	countString := d.getValue("w.count")
	count, _ := strconv.Atoi(countString)
	return count
}

func (d *Device) GetVideoStreamStatus(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "w." + IDString + ".status"
	statusString := d.getValue(key)
	status, _ := strconv.Atoi(statusString)
	return status
}

// Retuns codec of video stream
func (d *Device) GetVideoStreamType(ID int) string {
	IDString := strconv.Itoa(ID)
	key := "w." + IDString + ".type"
	streamType := d.getValue(key)
	return streamType
}

func (d *Device) GetVideoStreamProfile(ID int) string {
	IDString := strconv.Itoa(ID)
	key := "w." + IDString + ".type.profile"
	streamProfile := d.getValue(key)
	return streamProfile
}

func (d *Device) GetVideoStreamKind(ID int) string {
	IDString := strconv.Itoa(ID)
	key := "w." + IDString + ".kind"
	streamKind := d.getValue(key)
	return streamKind
}

func (d *Device) GetVideoStreamSize(ID int) string {
	IDString := strconv.Itoa(ID)
	key := "w." + IDString + ".size"
	streamSize := d.getValue(key)
	return streamSize
}

func (d *Device) GetVideoStreamQuality(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "w." + IDString + ".quality"
	response := d.getValue(key)
	streamQuality, _ := strconv.Atoi(response)
	return streamQuality
}

func (d *Device) GetVideoStreamBitrate(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "w." + IDString + ".cbr"
	response := d.getValue(key)
	bitrate, _ := strconv.Atoi(response)
	return bitrate
}

func (d *Device) GetVideoStreamFramerate(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "w." + IDString + ".frate"
	response := d.getValue(key)
	framerate, _ := strconv.Atoi(response)
	return framerate
}

func (d *Device) SetVideoStreamFramerate(ID int, framerate int) {
	IDString := strconv.Itoa(ID)
	valueString := strconv.Itoa(framerate)
	key := "w." + IDString + ".frate"
	d.setValue(key, valueString)
}

func (d *Device) GetVideoStreamFramerateMin(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "w." + IDString + ".frate.min"
	response := d.getValue(key)
	framerate, _ := strconv.Atoi(response)
	return framerate
}

func (d *Device) GetVideoStreamFramerateMax(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "w." + IDString + ".frate.max"
	response := d.getValue(key)
	framerate, _ := strconv.Atoi(response)
	return framerate
}

func (d *Device) GetVideoStreamCrop(ID int) string {
	IDString := strconv.Itoa(ID)
	key := "w." + IDString + ".crop"
	response := d.getValue(key)
	return response
}
