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

// key map
var keymap = [][]string{
	{}, // left blank intentionally
	{"", "q", "e", "r", "w", "s", "a", "d", "enter"},            // device 0x01 (master)
	{"", "a", "b", "c", "up", "down", "left", "right", "enter"}, // device 0x02
	{"", "", "", "", "", "", "", "", ""},                        // device 0x03
	{"", "", "", "", "", "", "", "", ""},                        // device 0x04
	{"", "", "", "", "", "", "", "", ""},                        // device 0x05
	{"", "", "", "", "", "", "", "", ""},                        // device 0x06
	{"", "", "", "", "", "", "", "", ""},                        // device 0x07
	{"", "", "", "", "", "", "", "", ""},                        // device 0x08
}

var keymap_preset map[string]([][]string) = map[string]([][]string){
	"ForzaHorizon5": {
		{}, // left blank intentionally
		{"", "q", "e", "r", "w", "s", "a", "d", "enter"},            // device 0x01 (master)
		{"", "a", "b", "c", "up", "down", "left", "right", "enter"}, // device 0x02
		{"", "", "", "", "", "", "", "", ""},                        // device 0x03
		{"", "", "", "", "", "", "", "", ""},                        // device 0x04
		{"", "", "", "", "", "", "", "", ""},                        // device 0x05
		{"", "", "", "", "", "", "", "", ""},                        // device 0x06
		{"", "", "", "", "", "", "", "", ""},                        // device 0x07
		{"", "", "", "", "", "", "", "", ""},                        // device 0x08
	},

	"GenshinImpact": {
		{}, // left blank intentionally
		{"", "e", "q", "space", "w", "s", "a", "d", "lclick"}, // device 0x01 (master)
		{"", "1", "2", "3", "f", "", "", "", "rclick"},        // device 0x02
		{"", "", "", "", "", "", "", "", ""},                  // device 0x03
		{"", "", "", "", "", "", "", "", ""},                  // device 0x04
		{"", "", "", "", "", "", "", "", ""},                  // device 0x05
		{"", "", "", "", "", "", "", "", ""},                  // device 0x06
		{"", "", "", "", "", "", "", "", ""},                  // device 0x07
		{"", "", "", "", "", "", "", "", ""},                  // device 0x08
	},
	"ForestIceFire": {},
}
