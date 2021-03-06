package client

import (
	"encoding/hex"
	"github.com/globalpokecache/POGOProtos-go"
	"math/rand"
	"time"
)

func randInt(l int) int {
	var s1 = rand.NewSource(time.Now().UnixNano())
	var r1 = rand.New(s1)
	return r1.Intn(l)
}

func randFloat() float64 {
	var s1 = rand.NewSource(time.Now().UnixNano())
	var r1 = rand.New(s1)
	return r1.Float64()
}

func getTimestamp(t time.Time) uint64 {
	return uint64(t.UnixNano() / int64(time.Millisecond))
}

func buildLocationFixes(requestEnvelope *protos.RequestEnvelope, timestampSinceStart uint64) (fixes []*protos.Signature_LocationFix) {
	if len(requestEnvelope.Requests) == 0 {
		return
	}

	providerCount := 4 + randInt(6)
	for i := 0; i < providerCount; i++ {
		var timestampSnapshot = timestampSinceStart + uint64(150*(i+1)+randInt(250*(i+1)-150*(i+1)))
		if timestampSnapshot >= timestampSinceStart {
			if len(fixes) != 0 {
				break
			}

			timestampSnapshot = timestampSinceStart - uint64(20+randInt(30))
			if timestampSnapshot < 0 {
				timestampSnapshot = 0
			}
		}

		fixes = append(fixes, &protos.Signature_LocationFix{
			TimestampSnapshot:  timestampSnapshot,
			Longitude:          float32(requestEnvelope.Longitude + 10 + (randFloat() * 100)),
			Latitude:           float32(requestEnvelope.Latitude + 10 + (randFloat() * 100)),
			HorizontalAccuracy: float32(5.0 + (randFloat() * 20.0)),
			VerticalAccuracy:   float32(5.0 + (randFloat() * 20.0)),
			Altitude:           float32(10.0 + (randFloat() * 20.0)),
			Provider:           "fused",
			ProviderStatus:     3,
			LocationType:       uint64(1),
			Course:             1,
		})
	}
	return
}

var (
	Devices = [][]string{
		{"iPad3,1", "iPad", "J1AP"},
		{"iPad3,2", "iPad", "J2AP"},
		{"iPad3,3", "iPad", "J2AAP"},
		{"iPad3,4", "iPad", "P101AP"},
		{"iPad3,5", "iPad", "P102AP"},
		{"iPad3,6", "iPad", "P103AP"},
		{"iPad4,1", "iPad", "J71AP"},
		{"iPad4,2", "iPad", "J72AP"},
		{"iPad4,3", "iPad", "J73AP"},
		{"iPad4,4", "iPad", "J85AP"},
		{"iPad4,5", "iPad", "J86AP"},
		{"iPad4,6", "iPad", "J87AP"},
		{"iPad4,7", "iPad", "J85mAP"},
		{"iPad4,8", "iPad", "J86mAP"},
		{"iPad4,9", "iPad", "J87mAP"},
		{"iPad5,1", "iPad", "J96AP"},
		{"iPad5,2", "iPad", "J97AP"},
		{"iPad5,3", "iPad", "J81AP"},
		{"iPad5,4", "iPad", "J82AP"},
		{"iPad6,7", "iPad", "J98aAP"},
		{"iPad6,8", "iPad", "J99aAP"},
		{"iPhone5,1", "iPhone", "N41AP"},
		{"iPhone5,2", "iPhone", "N42AP"},
		{"iPhone5,3", "iPhone", "N48AP"},
		{"iPhone5,4", "iPhone", "N49AP"},
		{"iPhone6,1", "iPhone", "N51AP"},
		{"iPhone6,2", "iPhone", "N53AP"},
		{"iPhone7,1", "iPhone", "N56AP"},
		{"iPhone7,2", "iPhone", "N61AP"},
		{"iPhone8,1", "iPhone", "N71AP"},
	}

	OsVersions = []string{
		"8.1.1", "8.1.2", "8.1.3", "8.2", "8.3",
		"8.4", "8.4.1", "9.0", "9.0.1", "9.0.2",
		"9.1", "9.2", "9.2.1", "9.3", "9.3.1",
		"9.3.2", "9.3.3", "9.3.4",
	}
)

func GetRandomDevice() *protos.Signature_DeviceInfo {
	var device = Devices[randInt(len(Devices))]
	var firmwareType = OsVersions[randInt(len(OsVersions))]

	shash := make([]byte, 16)
	rand.Read(shash)
	deviceID := hex.EncodeToString(shash)

	return &protos.Signature_DeviceInfo{
		DeviceId:             deviceID,
		DeviceBrand:          "Apple",
		DeviceModelBoot:      device[0],
		DeviceModel:          device[1],
		HardwareModel:        device[2],
		HardwareManufacturer: "Apple",
		FirmwareBrand:        "iPhone OS",
		FirmwareType:         firmwareType,
	}
}
