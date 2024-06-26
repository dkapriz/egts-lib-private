package egts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	extPosDataBytes      = []byte{0x0E, 0x32, 0x00, 0x00, 0x00, 0x0C}
	testEgtsSrExtPosData = SrExtPosData{
		NavigationSystemFieldExists:   "0",
		SatellitesFieldExists:         "1",
		PdopFieldExists:               "1",
		HdopFieldExists:               "1",
		VdopFieldExists:               "0",
		HorizontalDilutionOfPrecision: 50,
		PositionDilutionOfPrecision:   0,
		Satellites:                    12,
	}
)

func TestEgtsSrExtPosData_Encode(t *testing.T) {
	posDataBytes, err := testEgtsSrExtPosData.Encode(ProtocolVersionV1)
	if assert.NoError(t, err) {
		assert.Equal(t, posDataBytes, extPosDataBytes)
	}
}

func TestEgtsSrExtPosData_Decode(t *testing.T) {
	extPosData := SrExtPosData{}
	if assert.NoError(t, extPosData.Decode(extPosDataBytes, ProtocolVersionV1)) {
		assert.Equal(t, extPosData, testEgtsSrExtPosData)
	}
}

// Проверяем, что рекорд сет работает правильно с данным типом под записи
func TestEgtsSrExtPosDataRs(t *testing.T) {
	extPosDataRDBytes := append([]byte{0x11, 0x06, 0x00}, extPosDataBytes...)
	extPosDataRD := RecordDataSet{
		RecordData{
			SubrecordType:   SrExtPosDataType,
			SubrecordLength: 6,
			SubrecordData:   &testEgtsSrExtPosData,
		},
	}
	testStruct := RecordDataSet{}

	testBytes, err := extPosDataRD.Encode(ProtocolVersionV1)
	if assert.NoError(t, err) {
		assert.Equal(t, testBytes, extPosDataRDBytes)

		if assert.NoError(t, testStruct.Decode(extPosDataRDBytes, ProtocolVersionV1)) {
			assert.Equal(t, extPosDataRD, testStruct)
		}
	}
}
