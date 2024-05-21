package egts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	srAdAccelDataBytes = []byte{0x2, 0xb6, 0x91, 0x2b, 0x66, 0x29, 0x9, 0xee, 0xf3, 0x19, 0x2c, 0x5, 0xbb, 0x29, 0x9,
		0xfb, 0x43, 0x9, 0x25, 0xc, 0xad}
	testEgtsSrAccelData = SrAccelData{
		StructuresAmount: 2,
		AbsoluteTime:     1714131382,
		AccelDataStruct: []AccelDataStruct{
			{
				RelativeTime:    2345,
				XAxisAccelValue: -4365,
				YAxisAccelValue: 6444,
				ZAxisAccelValue: 1467,
			},
			{
				RelativeTime:    2345,
				XAxisAccelValue: -1213,
				YAxisAccelValue: 2341,
				ZAxisAccelValue: 3245,
			},
		},
	}
)

func TestEgtsSrAccelData_Encode(t *testing.T) {
	accelDataBytes, err := testEgtsSrAccelData.Encode(ProtocolVersionV1)
	if assert.NoError(t, err) {
		assert.Equal(t, accelDataBytes, srAdAccelDataBytes)
	}
}

func TestEgtsSrAccelData_Decode(t *testing.T) {
	accelDataBytes := SrAccelData{}

	if assert.NoError(t, accelDataBytes.Decode(srAdAccelDataBytes, ProtocolVersionV1)) {
		assert.Equal(t, accelDataBytes, testEgtsSrAccelData)
	}
}

// Проверяем, что рекордсет работает правильно с данным типом под записи
func TestEgtsSrAccelDataRs(t *testing.T) {
	adAccelDataBytes := append([]byte{0x14, 0x15, 0x00}, srAdAccelDataBytes...)
	adAccelData := RecordDataSet{
		RecordData{
			SubrecordType:   SrType20,
			SubrecordLength: testEgtsSrAccelData.Length(ProtocolVersionV1),
			SubrecordData:   &testEgtsSrAccelData,
		},
	}
	testStruct := RecordDataSet{}

	testBytes, err := adAccelData.Encode(ProtocolVersionV1)
	if assert.NoError(t, err) {
		assert.Equal(t, testBytes, adAccelDataBytes)

		if assert.NoError(t, testStruct.Decode(adAccelDataBytes, ProtocolVersionV1)) {
			assert.Equal(t, adAccelData, testStruct)
		}
	}
}
