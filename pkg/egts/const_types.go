package egts

// SrRecordResponseType код типа под записи EGTS_SR_RECORD_RESPONSE
const SrRecordResponseType = 0

// SrTermIdentityType код типа под записи EGTS_SR_TERM_IDENTITY
const SrTermIdentityType = 1

// SrModuleDataType код типа под записи EGTS_SR_MODULE_DATA
const SrModuleDataType = 2

// SrResultCodeType код типа под записи EGTS_SR_RESULT_CODE
const SrResultCodeType = 9

// SrAuthInfoType код типа под записи EGTS_SR_AUTH_INFO
const SrAuthInfoType = 7

// SrEgtsPlusDataType код типа под записи EGTS_SR_EGTS_PLUS_DATA
const SrEgtsPlusDataType = 15

// SrPosDataType код типа под записи EGTS_SR_POS_DATA
const SrPosDataType = 16

// SrExtPosDataType код типа под записи EGTS_SR_EXT_POS_DATA
const SrExtPosDataType = 17

// SrAdSensorsDataType код типа под записи EGTS_SR_AD_SENSORS_DATA
const SrAdSensorsDataType = 18

// SrCountersDataType код типа под записи EGTS_SR_COUNTERS_DATA
const SrCountersDataType = 19

// SrType20 в зависимости от длины может содержать секцию EGTS_SR_STATE_DATA (если длина 5 байт) или EGTS_SR_ACCEL_DATA
const SrType20 = 20

// SrStateDataType код типа под записи EGTS_SR_STATE_DATA
const SrStateDataType = 21

// SrLoopInDataType код типа под записи EGTS_SR_TERM_IDENTITY_TYPE
const SrLoopInDataType = 22

// SrAbsDigSensDataType код типа под записи EGTS_SR_ABS_DIG_SENS_DATA
const SrAbsDigSensDataType = 23

// SrAbsAnSensDataType код типа под записи EGTS_SR_ABS_AN_SENS_DATA
const SrAbsAnSensDataType = 24

// SrAbsCntrDataType код типа под записи EGTS_SR_ABS_CNTR_DATA
const SrAbsCntrDataType = 25

// SrAbsLoopInDataType код типа под записи EGTS_SR_ABS_LOOPIN_DATA
const SrAbsLoopInDataType = 26

// SrLiquidLevelSensorType код код типа под записи EGTS_SR_LIQUID_LEVEL_SENSOR
const SrLiquidLevelSensorType = 27

// SrPassengersCountersType код типа под записи EGTS_SR_PASSENGERS_COUNTERS
const SrPassengersCountersType = 28

// PtAppdataPacket код типа пакета PT_APP_DATA
const PtAppdataPacket = 1

// PtResponsePacket код типа пакета PT_RESPONSE
const PtResponsePacket = 0

// AuthService тип сервиса AUTH_SERVICE
const AuthService = 1

// TeleDataService тип сервиса TELEDATA_SERVICE
const TeleDataService = 2

// SrDispatcherIdentityType код типа под записи EGTS_SR_DISPATCHER_IDENTITY
const SrDispatcherIdentityType = 5

const ProtocolVersionV1 byte = 1
const ProtocolVersionV2 byte = 2
