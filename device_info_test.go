package adb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDeviceShort(t *testing.T) {
	dev, err := parseDeviceShort("192.168.56.101:5555	device\n")
	assert.NoError(t, err)
	assert.Equal(t, &DeviceInfo{
		Serial: "192.168.56.101:5555",
		State:  StateOnline,
	}, dev)
}

func TestParseDeviceLong(t *testing.T) {
	dev, err := parseDeviceLong("SERIAL    device product:PRODUCT model:MODEL device:DEVICE\n")
	assert.NoError(t, err)
	assert.Equal(t, &DeviceInfo{
		Serial:     "SERIAL",
		State:      StateOnline,
		Product:    "PRODUCT",
		Model:      "MODEL",
		DeviceInfo: "DEVICE",
	}, dev)
}

func TestParseDeviceLongUsb(t *testing.T) {
	dev, err := parseDeviceLong("SERIAL    device usb:1234 product:PRODUCT model:MODEL device:DEVICE \n")
	assert.NoError(t, err)
	assert.Equal(t, &DeviceInfo{
		Serial:     "SERIAL",
		State:      StateOnline,
		Product:    "PRODUCT",
		Model:      "MODEL",
		DeviceInfo: "DEVICE",
		Usb:        "1234",
	}, dev)
}

func TestParseDeviceNotAuthorized(t *testing.T) {
	dev, err := parseDeviceLong("SERIAL         unauthorized usb:1234 transport_id:1\n")
	assert.NoError(t, err)
	assert.Equal(t, &DeviceInfo{
		Serial: "SERIAL",
		State:  StateUnauthorized,
		Usb:    "1234",
	}, dev)
}
