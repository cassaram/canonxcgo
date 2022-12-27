package canonxcgo

// File contains light amount correction information commands / values

import (
	"strconv"

	bsutil "github.com/cassaram/broadcastutils"
)

func (d *Device) GetCameraExposureMode(ID int) string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".exp"
	response := d.getValue(key)
	return response
}

func (d *Device) SetCameraExposureMode(ID int, mode string) {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".exp"
	d.setValue(key, mode)
}

func (d *Device) GetCameraExposureModeList(ID int) []string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".exp.list"
	response := d.getValue(key)
	values := parseListCSV(response)
	return values
}

func (d *Device) GetCameraAEBrightness(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".ae.brightness"
	response := d.getValue(key)
	brightness, _ := strconv.Atoi(response)
	return brightness
}

func (d *Device) GetCameraAEBrightnessMin(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".ae.brightness.min"
	response := d.getValue(key)
	brightness, _ := strconv.Atoi(response)
	return brightness
}

func (d *Device) GetCameraAEBrightnessMax(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".ae.brightness.max"
	response := d.getValue(key)
	brightness, _ := strconv.Atoi(response)
	return brightness
}

func (d *Device) GetCameraAEBrightnessList(ID int) []int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".ae.brightness.list"
	response := d.getValue(key)
	list := parseListCSV(response)
	result := make([]int, len(list))
	for id, val := range list {
		brightness, _ := strconv.Atoi(val)
		result[id] = brightness
	}
	return result
}

func (d *Device) GetCameraAEMeteringMode(ID int) string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".ae.photometry"
	response := d.getValue(key)
	return response
}

func (d *Device) GetCameraAEMeteringModeList(ID int) []string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".ae.photometry"
	response := d.getValue(key)
	list := parseListCSV(response)
	return list
}

func (d *Device) GetCameraAEGainLimitMax(ID int) bsutil.Decibel {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".ae.gainlimit.max"
	response := d.getValue(key)
	valueInt, _ := strconv.Atoi(response)
	valueDb := bsutil.IntToDecibel(valueInt)
	return valueDb
}

func (d *Device) SetCameraAEGainLimitMax(ID int, gain bsutil.Decibel) {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".ae.gainlimit.max"
	gainInt := bsutil.DecibelToInt(gain)
	gainString := strconv.Itoa(gainInt)
	d.setValue(key, gainString)
}

func (d *Device) GetCameraAEGainLimitMaxMin(ID int) bsutil.Decibel {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".ae.gainlimit.max.min"
	response := d.getValue(key)
	valueInt, _ := strconv.Atoi(response)
	valueDb := bsutil.IntToDecibel(valueInt)
	return valueDb
}

func (d *Device) GetCameraAEGainLimitMaxMax(ID int) bsutil.Decibel {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".ae.gainlimit.max.max"
	response := d.getValue(key)
	valueInt, _ := strconv.Atoi(response)
	valueDb := bsutil.IntToDecibel(valueInt)
	return valueDb
}

func (d *Device) GetCameraAEFlickerReduction(ID int) string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".ae.flickerreduct"
	response := d.getValue(key)
	return response
}

func (d *Device) SetCameraAEFlickerReduction(ID int, setting string) {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".ae.flickerreduct"
	d.setValue(key, setting)
}

func (d *Device) GetCameraAEFlickerReductionList(ID int) []string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".ae.flickerreduct.list"
	response := d.getValue(key)
	list := parseListCSV(response)
	return list
}

func (d *Device) GetCameraAEResponse(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".ae.resp"
	response := d.getValue(key)
	value, _ := strconv.Atoi(response)
	return value
}

func (d *Device) SetCameraAEResponse(ID int, rate int) {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".ae.resp"
	valueString := strconv.Itoa(rate)
	d.setValue(key, valueString)
}

func (d *Device) GetCameraAEResponseMin(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".ae.resp.min"
	response := d.getValue(key)
	value, _ := strconv.Atoi(response)
	return value
}

func (d *Device) GetCameraAEResponseMax(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".ae.resp.max"
	response := d.getValue(key)
	value, _ := strconv.Atoi(response)
	return value
}

func (d *Device) GetCameraMEShutter(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.shutter"
	response := d.getValue(key)
	value, _ := strconv.Atoi(response)
	return value
}

func (d *Device) SetCameraMEShutter(ID int, shutter int) {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.shutter"
	valStr := strconv.Itoa(shutter)
	d.setValue(key, valStr)
}

func (d *Device) GetCameraMEShutterList(ID int) []int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.shutter.list"
	response := d.getValue(key)
	list := parseListCSV(response)
	listInt := make([]int, len(list))
	for idx, val := range list {
		valInt, _ := strconv.Atoi(val)
		listInt[idx] = valInt
	}
	return listInt
}

func (d *Device) GetCameraMEShutterMode(ID int) string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.shutter.mode"
	response := d.getValue(key)
	return response
}

func (d *Device) GetCameraMEShutterModeList(ID int) []string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.shutter.mode"
	response := d.getValue(key)
	list := parseListCSV(response)
	return list
}

func (d *Device) GetCameraMEClearscanShutter(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.clearscan"
	response := d.getValue(key)
	valueInt, _ := strconv.Atoi(response)
	return valueInt
}

func (d *Device) SetCameraMEClearscanShutter(ID int, shutter int) {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.clearscan"
	valueInt := strconv.Itoa(shutter)
	d.setValue(key, valueInt)
}

func (d *Device) GetCameraMEClearscanShutterMin(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.clearscan.min"
	response := d.getValue(key)
	valueInt, _ := strconv.Atoi(response)
	return valueInt
}

func (d *Device) GetCameraMEClearscanShutterMax(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.clearscan.max"
	response := d.getValue(key)
	valueInt, _ := strconv.Atoi(response)
	return valueInt
}

// WARNING: abstract value
//
// Recommended to use Device.GetCameraMEFStop for FStop values
func (d *Device) GetCameraMEIris(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.iris"
	response := d.getValue(key)
	valueInt, _ := strconv.Atoi(response)
	return valueInt
}

// WARNING: abstract value
//
// Recommended to use Device.SetCameraMEFStop for FStop values
func (d *Device) SetCameraMEIris(ID int, iris int) {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.iris"
	valueInt := strconv.Itoa(iris)
	d.setValue(key, valueInt)
}

// WARNING: abstract value
func (d *Device) GetCameraMEIrisMin(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.iris.min"
	response := d.getValue(key)
	valueInt, _ := strconv.Atoi(response)
	return valueInt
}

// WARNING: abstract value
func (d *Device) GetCameraMEIrisMax(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.iris.max"
	response := d.getValue(key)
	valueInt, _ := strconv.Atoi(response)
	return valueInt
}

func (d *Device) GetCameraMEFStop(ID int) bsutil.FStop {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.diaphragm"
	response := d.getValue(key)
	valueInt, _ := strconv.Atoi(response)
	fs := bsutil.IntToFStop(valueInt)
	return fs
}

func (d *Device) SetCameraMEFStop(ID int, fs bsutil.FStop) {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.diaphragm"
	valueInt := bsutil.FStopToInt(fs)
	valueString := strconv.Itoa(valueInt)
	d.setValue(key, valueString)
}

func (d *Device) GetCameraMEFStopList(ID int) []bsutil.FStop {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.diaphragm"
	response := d.getValue(key)
	list := parseListCSV(response)
	listResult := make([]bsutil.FStop, len(list))
	for idx, val := range list {
		valInt, _ := strconv.Atoi(val)
		valFStop := bsutil.IntToFStop(valInt)
		listResult[idx] = valFStop
	}
	return listResult
}

func (d *Device) GetCameraMEIrisMode(ID int) string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.diaphragm.mode"
	response := d.getValue(key)
	return response
}

func (d *Device) SetCameraMEIrisMode(ID int, mode string) {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.diaphragm.mode"
	d.setValue(key, mode)
}

func (d *Device) GetCameraMEIrisModeList(ID int) []string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.diaphragm.mode.list"
	response := d.getValue(key)
	list := parseListCSV(response)
	return list
}

func (d *Device) GetCameraMEGain(ID int) bsutil.Decibel {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.gain"
	response := d.getValue(key)
	gainInt, _ := strconv.Atoi(response)
	gain := bsutil.IntToDecibel(gainInt)
	return gain
}

func (d *Device) SetCameraMEGain(ID int, gain bsutil.Decibel) {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.gain"
	gainInt := bsutil.DecibelToInt(gain)
	gainString := strconv.Itoa(gainInt)
	d.setValue(key, gainString)
}

func (d *Device) GetCameraMEGainMin(ID int) bsutil.Decibel {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.gain.min"
	response := d.getValue(key)
	gainInt, _ := strconv.Atoi(response)
	gain := bsutil.IntToDecibel(gainInt)
	return gain
}

func (d *Device) GetCameraMEGainMax(ID int) bsutil.Decibel {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.gain.max"
	response := d.getValue(key)
	gainInt, _ := strconv.Atoi(response)
	gain := bsutil.IntToDecibel(gainInt)
	return gain
}

func (d *Device) GetCameraMEGainMode(ID int) string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.gain.mode"
	response := d.getValue(key)
	return response
}

func (d *Device) SetCameraMEGainMode(ID int, mode string) {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.gain.mode"
	d.setValue(key, mode)
}

func (d *Device) GetCameraMEGainModeList(ID int) []string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.gain.mode.list"
	response := d.getValue(key)
	list := parseListCSV(response)
	return list
}

func (d *Device) GetCameraMEGainLimitMax(ID int) bsutil.Decibel {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.gainlimit.max"
	response := d.getValue(key)
	gainInt, _ := strconv.Atoi(response)
	gain := bsutil.IntToDecibel(gainInt)
	return gain
}

func (d *Device) SetCameraMEGainLimitMax(ID int, gain bsutil.Decibel) {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.gainlimit.max"
	gainInt := bsutil.DecibelToInt(gain)
	gainString := strconv.Itoa(gainInt)
	d.setValue(key, gainString)
}

func (d *Device) GetCameraMEGainLimitMaxMin(ID int) bsutil.Decibel {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.gainlimit.max.min"
	response := d.getValue(key)
	gainInt, _ := strconv.Atoi(response)
	gain := bsutil.IntToDecibel(gainInt)
	return gain
}

func (d *Device) GetCameraMEGainLimitMaxMax(ID int) bsutil.Decibel {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.gainlimit.max.max"
	response := d.getValue(key)
	gainInt, _ := strconv.Atoi(response)
	gain := bsutil.IntToDecibel(gainInt)
	return gain
}

func (d *Device) GetCameraMEBrightness(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.brightness"
	response := d.getValue(key)
	valueInt, _ := strconv.Atoi(response)
	return valueInt
}

func (d *Device) SetCameraMEBrightness(ID int, brightness int) {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.brightness"
	valueString := strconv.Itoa(brightness)
	d.setValue(key, valueString)
}

func (d *Device) GetCameraMEBrightnessMin(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.brightness.min"
	response := d.getValue(key)
	valueInt, _ := strconv.Atoi(response)
	return valueInt
}

func (d *Device) GetCameraMEBrightnessMax(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.brightness.max"
	response := d.getValue(key)
	valueInt, _ := strconv.Atoi(response)
	return valueInt
}

func (d *Device) GetCameraMEMeteringMode(ID int) string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.photometry"
	response := d.getValue(key)
	return response
}

func (d *Device) SetCameraMEMeteringMode(ID int, mode string) {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.photometry"
	d.setValue(key, mode)
}

func (d *Device) GetCameraMEMeteringModeList(ID int) []string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.photometry.list"
	response := d.getValue(key)
	list := parseListCSV(response)
	return list
}

func (d *Device) GetCameraMEResponse(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.resp"
	response := d.getValue(key)
	respInt, _ := strconv.Atoi(response)
	return respInt
}

func (d *Device) SetCameraMEResponse(ID int, response int) {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.resp"
	valStr := strconv.Itoa(response)
	d.setValue(key, valStr)
}

func (d *Device) GetCameraMEResponseMin(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.resp.min"
	response := d.getValue(key)
	respInt, _ := strconv.Atoi(response)
	return respInt
}

func (d *Device) GetCameraMEResponseMax(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.resp.max"
	response := d.getValue(key)
	respInt, _ := strconv.Atoi(response)
	return respInt
}

func (d *Device) GetCameraMEFlickerReductionMode(ID int) string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.flickerreduct"
	response := d.getValue(key)
	return response
}

func (d *Device) SetCameraMEFlickerReductionMode(ID int, mode string) {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.flickerreduct"
	d.setValue(key, mode)
}

func (d *Device) GetCameraMEFlickerReductionModeList(ID int) []string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".me.flickerreduct"
	response := d.getValue(key)
	list := parseListCSV(response)
	return list
}

func (d *Device) GetCameraNDMode(ID int) string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".nd.mode"
	response := d.getValue(key)
	return response
}

func (d *Device) SetCameraNDMode(ID int, mode string) {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".nd.mode"
	d.setValue(key, mode)
}

func (d *Device) GetCameraNDModeList(ID int) []string {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".nd.mode.list"
	response := d.getValue(key)
	list := parseListCSV(response)
	return list
}

func (d *Device) GetCameraNDFilter(ID int) int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".nd.filter"
	response := d.getValue(key)
	filter, _ := strconv.Atoi(response)
	return filter
}

func (d *Device) SetCameraNDFilter(ID int, filter int) {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".nd.filter"
	valueString := strconv.Itoa(filter)
	d.setValue(key, valueString)
}

func (d *Device) GetCameraNDFilterList(ID int) []int {
	IDString := strconv.Itoa(ID)
	key := "c." + IDString + ".nd.filter"
	response := d.getValue(key)
	list := parseListCSV(response)
	listResult := make([]int, len(list))
	for idx, item := range list {
		itemInt, _ := strconv.Atoi(item)
		listResult[idx] = itemInt
	}
	return listResult
}
