package devices

const ProtocolHead = 0x01
const DeviceSpecialProtocol = 0x00

// key selection bit
const (
	KeyTestConnection = 0x00

	Key1     = 0x01
	Key2     = 0x02
	Key3     = 0x03
	KeyUp    = 0x04
	KeyDown  = 0x05
	KeyLeft  = 0x06
	KeyRight = 0x07
	KeyOk    = 0x08
)

// key action bit
const (
	KeyActionRequireTest  = 0x00
	KeyActionConnectionOk = 0x01

	KeyActionPress   = 0x01
	KeyActionRelease = 0x02
)

var TestConnectionSend = []byte{ProtocolHead, DeviceSpecialProtocol, KeyTestConnection, KeyActionRequireTest, 0x00, 0x00, 0x00, 0x00}
