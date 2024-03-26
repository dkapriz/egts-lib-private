package egts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	srDispatcherIdentityPkgBytesV1 = []byte{0x01, 0x00, 0x00, 0x0b, 0x00, 0x0f, 0x00, 0x01, 0x00,
		0x01, 0x06, 0x08, 0x00, 0x00, 0x00, 0x98, 0x01, 0x01, 0x05, 0x05, 0x00, 0x00, 0x47, 0x00,
		0x00, 0x00, 0x51, 0x9d}

	testDispatcherIdentityPkgV1 = Package{
		ProtocolVersion:  1,
		SecurityKeyID:    0,
		Prefix:           "00",
		Route:            "0",
		EncryptionAlg:    "00",
		Compression:      "0",
		Priority:         "00",
		HeaderLength:     11,
		HeaderEncoding:   0,
		FrameDataLength:  15,
		PacketIdentifier: 1,
		PacketType:       PtAppdataPacket,
		HeaderCheckSum:   6,
		ServicesFrameData: &ServiceDataSet{
			{
				RecordLength:             0x08,
				SourceServiceOnDevice:    "1",
				RecipientServiceOnDevice: "0",
				Group:                    "0",
				RecordProcessingPriority: "11",
				TimeFieldExists:          "0",
				EventIDFieldExists:       "0",
				ObjectIDFieldExists:      "0",
				SourceServiceType:        0x01,
				RecipientServiceType:     0x01,
				RecordDataSet: RecordDataSet{
					{
						SubrecordType:   0x05,
						SubrecordLength: 0x05,
						SubrecordData: &SrDispatcherIdentity{
							DispatcherType: 0,
							DispatcherID:   71,
						},
					},
				},
			},
		},
		ServicesFrameDataCheckSum: 40273,
	}

	srDispatcherIdentityPkgBytesV2 = []byte{0x2, 0x0, 0x0, 0xb, 0x0, 0x19, 0x0, 0x1, 0x0, 0x1, 0x2c, 0x12, 0x0, 0x0,
		0x0, 0x98, 0x1, 0x1, 0x5, 0xf, 0x0, 0x0, 0x47, 0x0, 0x0, 0x0, 0xda, 0xdb, 0x31, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x30, 0x32, 0xdb, 0x17}

	testDispatcherIdentityPkgV2 = Package{
		ProtocolVersion:  2,
		SecurityKeyID:    0,
		Prefix:           "00",
		Route:            "0",
		EncryptionAlg:    "00",
		Compression:      "0",
		Priority:         "00",
		HeaderLength:     11,
		HeaderEncoding:   0,
		FrameDataLength:  25,
		PacketIdentifier: 1,
		PacketType:       PtAppdataPacket,
		HeaderCheckSum:   44,
		ServicesFrameData: &ServiceDataSet{
			{
				RecordLength:             0x12,
				SourceServiceOnDevice:    "1",
				RecipientServiceOnDevice: "0",
				Group:                    "0",
				RecordProcessingPriority: "11",
				TimeFieldExists:          "0",
				EventIDFieldExists:       "0",
				ObjectIDFieldExists:      "0",
				SourceServiceType:        0x01,
				RecipientServiceType:     0x01,
				RecordDataSet: RecordDataSet{
					{
						SubrecordType:   0x05,
						SubrecordLength: 0x0F,
						SubrecordData: &SrDispatcherIdentity{
							DispatcherType:                     0,
							DispatcherID:                       71,
							TerminalIdentifier:                 3267546,
							ServiceSupportLevelProtocolVersion: "02",
						},
					},
				},
			},
		},
		ServicesFrameDataCheckSum: 6107,
	}

	testEgtsSrDispatcherIdentityBinV2 = []byte{0x0, 0x47, 0x0, 0x0, 0x0, 0xda, 0xdb, 0x31, 0x0, 0x0, 0x0, 0x0, 0x0,
		0x30, 0x32}
	testEgtsSrDispatcherIdentityV2 = SrDispatcherIdentity{
		DispatcherType:                     0,
		DispatcherID:                       71,
		TerminalIdentifier:                 3267546,
		ServiceSupportLevelProtocolVersion: "02",
	}
)

func TestEgtsSrDispatcherIdentity_Encode_Version_2(t *testing.T) {
	sti, err := testEgtsSrDispatcherIdentityV2.Encode(ProtocolVersionV2)

	if assert.NoError(t, err) {
		assert.Equal(t, sti, testEgtsSrDispatcherIdentityBinV2)
	}
}

func TestEgtsSrDispatcherIdentity_Decode_Version_2(t *testing.T) {
	srDispatcherIdent := SrDispatcherIdentity{}

	if assert.NoError(t, srDispatcherIdent.Decode(testEgtsSrDispatcherIdentityBinV2, ProtocolVersionV2)) {
		assert.Equal(t, srDispatcherIdent, testEgtsSrDispatcherIdentityV2)
	}
}

func TestEgtsSrDispatcherIdentityPkg_Encode_Version_1(t *testing.T) {
	dispatcherIdentity, err := testDispatcherIdentityPkgV1.Encode()
	if assert.NoError(t, err) {
		assert.Equal(t, dispatcherIdentity, srDispatcherIdentityPkgBytesV1)
	}
}

func TestEgtsSrDispatcherIdentityPkg_Decode_Version_1(t *testing.T) {
	pkg := Package{}

	if _, err := pkg.Decode(srDispatcherIdentityPkgBytesV1); assert.NoError(t, err) {
		assert.Equal(t, pkg, testDispatcherIdentityPkgV1)
	}
}

func TestEgtsSrDispatcherIdentityPkg_Encode_Version_2(t *testing.T) {
	dispatcherIdentity, err := testDispatcherIdentityPkgV2.Encode()
	if assert.NoError(t, err) {
		assert.Equal(t, dispatcherIdentity, srDispatcherIdentityPkgBytesV2)
	}
}

func TestEgtsSrDispatcherIdentityPkg_Decode_Version_2(t *testing.T) {
	pkg := Package{}

	if _, err := pkg.Decode(srDispatcherIdentityPkgBytesV2); assert.NoError(t, err) {
		assert.Equal(t, pkg, testDispatcherIdentityPkgV2)
	}
}
