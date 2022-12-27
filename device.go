package canonxcgo

import (
	"io"
	"net/http"
	"strings"
)

type Device struct {
	IPAddress string
}

// Function to set a device value via HTTP
func (d *Device) setValue(key string, val string) {
	url := "http://" + d.IPAddress + "/-wvhttp-01-/control.cgi?" + key + "=" + val

	_, err := http.Post(url, "text/plain", strings.NewReader(""))
	if err != nil {
		return
	}
}

// Function to get a device value via HTTP
func (d *Device) getValue(key string) string {
	url := "http://" + d.IPAddress + "/-wvhttp-01-/info.cgi?item=" + key

	resp, err := http.Get(url)
	if err != nil {
		return ""
	}

	body, _ := io.ReadAll(resp.Body)
	respMap := readResponse(string(body))

	val := respMap[key]
	return val
}
