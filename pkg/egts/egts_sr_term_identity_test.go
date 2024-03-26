package egts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testEgtsSrTermIdentityBinV1 = []byte{0xB0, 0x09, 0x02, 0x00, 0x10}
	testEgtsSrTermIdentityV1    = SrTermIdentity{
		TerminalIdentifier: 133552,
		MNE:                "0",
		BSE:                "0",
		NIDE:               "0",
		SSRA:               "1",
		LNGCE:              "0",
		IMSIE:              "0",
		IMEIE:              "0",
		HDIDE:              "0",
	}
	testEgtsSrTermIdentityPkgBinV1 = []byte{0x01, 0x00, 0x03, 0x0B, 0x00, 0x13, 0x00, 0x86, 0x00, 0x01, 0xB6, 0x08, 0x00,
		0x5F, 0x00, 0x99, 0x02, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01, 0x05, 0x00, 0xB0, 0x09, 0x02, 0x00, 0x10, 0x0D, 0xCE}
	testEgtsSrTermIdentityPkgV1 = Package{
		ProtocolVersion:  1,
		SecurityKeyID:    0,
		Prefix:           "00",
		Route:            "0",
		EncryptionAlg:    "00",
		Compression:      "0",
		Priority:         "11",
		HeaderLength:     11,
		HeaderEncoding:   0,
		FrameDataLength:  19,
		PacketIdentifier: 134,
		PacketType:       PtAppdataPacket,
		HeaderCheckSum:   182,
		ServicesFrameData: &ServiceDataSet{
			ServiceDataRecord{
				RecordLength:             8,
				RecordNumber:             95,
				SourceServiceOnDevice:    "1",
				RecipientServiceOnDevice: "0",
				Group:                    "0",
				RecordProcessingPriority: "11",
				TimeFieldExists:          "0",
				EventIDFieldExists:       "0",
				ObjectIDFieldExists:      "1",
				ObjectIdentifier:         2,
				SourceServiceType:        AuthService,
				RecipientServiceType:     AuthService,
				RecordDataSet: RecordDataSet{
					RecordData{
						SubrecordType:   SrTermIdentityType,
						SubrecordLength: 5,
						SubrecordData:   &testEgtsSrTermIdentityV1,
					},
				},
			},
		},
		ServicesFrameDataCheckSum: 52749,
	}

	testEgtsSrTermIdentityBinV2 = []byte{0x30, 0x0, 0x45, 0xf6, 0x4f, 0x0, 0x0, 0x0, 0x10, 0x30, 0x32}
	testEgtsSrTermIdentityV2    = SrTermIdentity{
		TerminalIdentifier:                 343434133552,
		MNE:                                "0",
		BSE:                                "0",
		NIDE:                               "0",
		SSRA:                               "1",
		LNGCE:                              "0",
		IMSIE:                              "0",
		IMEIE:                              "0",
		HDIDE:                              "0",
		ServiceSupportLevelProtocolVersion: "02",
	}
)

func TestEgtsSrTermIdentity_Encode_Version1(t *testing.T) {
	sti, err := testEgtsSrTermIdentityV1.Encode(ProtocolVersionV1)

	if assert.NoError(t, err) {
		assert.Equal(t, sti, testEgtsSrTermIdentityBinV1)
	}
}

func TestEgtsSrTermIdentity_Decode_Version1(t *testing.T) {
	srTermIdent := SrTermIdentity{}

	if assert.NoError(t, srTermIdent.Decode(testEgtsSrTermIdentityBinV1, ProtocolVersionV1)) {
		assert.Equal(t, srTermIdent, testEgtsSrTermIdentityV1)
	}
}

func TestEgtsSrTermIdentity_Encode_Version2(t *testing.T) {
	sti, err := testEgtsSrTermIdentityV2.Encode(ProtocolVersionV2)

	if assert.NoError(t, err) {
		assert.Equal(t, sti, testEgtsSrTermIdentityBinV2)
	}
}

func TestEgtsSrTermIdentity_Decode_Version2(t *testing.T) {
	srTermIdent := SrTermIdentity{}

	if assert.NoError(t, srTermIdent.Decode(testEgtsSrTermIdentityBinV2, ProtocolVersionV2)) {
		assert.Equal(t, srTermIdent, testEgtsSrTermIdentityV2)
	}
}

func TestEgtsSrTermIdentityPkg_Encode_Version1(t *testing.T) {
	pkg, err := testEgtsSrTermIdentityPkgV1.Encode()

	if assert.NoError(t, err) {
		assert.Equal(t, pkg, testEgtsSrTermIdentityPkgBinV1)
	}
}

func TestEgtsSrTermIdentityPkg_Decode_Version1(t *testing.T) {
	srTermIdentPkg := Package{}

	_, err := srTermIdentPkg.Decode(testEgtsSrTermIdentityPkgBinV1)
	if assert.NoError(t, err) {
		assert.Equal(t, srTermIdentPkg, testEgtsSrTermIdentityPkgV1)
	}
}
