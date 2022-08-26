package devices

const ProtocolHead = 0x01
const DeviceSpecialProtocol = 0x00

// key selection bit
const (
	KeyTestConnection = 0x00
)

// key action bit
const (
	KeyActionRequireTest  = 0x00
	KeyActionConnectionOk = 0x01
	KeyActionPress        = 0x01
	KeyActionRelease      = 0x02
)
