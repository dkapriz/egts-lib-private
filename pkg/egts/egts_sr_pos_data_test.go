package egts

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	testEgtsSrPosDataBytesV1 = []byte{0x55, 0x91, 0x02, 0x10, 0x6F, 0x1C, 0x05, 0x9E, 0x7A, 0xB5, 0x3C, 0x35,
		0x01, 0xD0, 0x87, 0x2C, 0x01, 0x00, 0x00, 0x00, 0x00}
	testEgtsSrPosDataV1 = SrPosData{
		NavigationTime:      time.Date(2018, time.July, 6, 20, 8, 53, 0, time.UTC),
		Latitude:            55.55389399769574,
		Longitude:           37.43236696287812,
		ALTE:                "0",
		LOHS:                "0",
		LAHS:                "0",
		MV:                  "0",
		BB:                  "0",
		CS:                  "0",
		FIX:                 "0",
		VLD:                 "1",
		DirectionHighestBit: 1,
		AltitudeSign:        0,
		Speed:               200,
		Direction:           172,
		Odometer:            1,
		DigitalInputs:       0,
		Source:              0,
	}
)

var (
	testEgtsSrPosDataBytesV2 = []byte{0x55, 0x91, 0x2, 0x10, 0x6f, 0x1c, 0x5, 0x9e, 0x7a, 0xb5, 0x3c, 0x35, 0x1,
		0xd0, 0x87, 0x2c, 0x1, 0x0, 0x0, 0x0, 0x0, 0xd0, 0x11, 0x0, 0x30, 0x5d, 0x0, 0x0, 0x1, 0xc8, 0x1c}
	testEgtsSrPosDataV2 = SrPosData{
		NavigationTime:      time.Date(2018, time.July, 6, 20, 8, 53, 0, time.UTC),
		Latitude:            55.55389399769574,
		Longitude:           37.43236696287812,
		ALTE:                "0",
		LOHS:                "0",
		LAHS:                "0",
		MV:                  "0",
		BB:                  "0",
		CS:                  "0",
		FIX:                 "0",
		VLD:                 "1",
		DirectionHighestBit: 1,
		AltitudeSign:        0,
		Speed:               200,
		Direction:           172,
		Odometer:            1,
		DigitalInputs:       0,
		Source:              0,
		NetworkIdentifier:   4560,  //Protocol version 2
		LocalAreaCode:       23856, //Protocol version 2
		CellIdentifier:      456,   //Protocol version 2
		SignalStrength:      28,    //Protocol version 2
	}
)

func TestEgtsSrPosData_Encode_Protocol_V1(t *testing.T) {
	posDataBytes, err := testEgtsSrPosDataV1.Encode(ProtocolVersionV1)

	if assert.NoError(t, err) {
		assert.Equal(t, posDataBytes, testEgtsSrPosDataBytesV1)
	}
}

func TestEgtsSrPosData_Decode_Protocol_V1(t *testing.T) {
	posData := SrPosData{}

	if assert.NoError(t, posData.Decode(testEgtsSrPosDataBytesV1, ProtocolVersionV1)) {
		assert.Equal(t, posData, testEgtsSrPosDataV1)
	}
}

func TestEgtsSrPosData_Encode_Protocol_V2(t *testing.T) {
	posDataBytes, err := testEgtsSrPosDataV2.Encode(ProtocolVersionV2)

	if assert.NoError(t, err) {
		assert.Equal(t, posDataBytes, testEgtsSrPosDataBytesV2)
	}
}

func TestEgtsSrPosData_Decode_Protocol_V2(t *testing.T) {
	posData := SrPosData{}

	if assert.NoError(t, posData.Decode(testEgtsSrPosDataBytesV2, ProtocolVersionV2)) {
		assert.Equal(t, posData, testEgtsSrPosDataV2)
	}
}
