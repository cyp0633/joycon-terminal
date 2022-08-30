package devices

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/go-vgo/robotgo"
	"go.bug.st/serial"
)

var conf = &serial.Mode{
	BaudRate: 4800,
	DataBits: 8,
	Parity:   serial.NoParity,
	StopBits: serial.OneStopBit,
}

var Conn serial.Port

var EnableRead bool
var EnableReadMtx sync.Mutex

var onlineStatus [9]int8

// GetPortList returns a list of available serial ports
func GetPortList() ([]string, error) {
	list, err := serial.GetPortsList()
	log.Printf("Available ports: %v", list)
	if err != nil {
		log.Printf("Get port list failed: %v", err)
	}
	return list, err
}

// ConnectSerial connects to the serial port
func ConnectSerial(path string) error {
	conn, err := serial.Open(path, conf)
	Conn = conn
	log.Printf("Connected to port %s", path)
	if err != nil {
		log.Print(err.Error())
	}
	return err
}

func read(sleepTime time.Duration) ([]byte, int) {
	time.Sleep(sleepTime) // wait 500 ms to read into system buffer
	buf := make([]byte, 8)
	num, err := Conn.Read(buf)
	log.Printf("Data received: %v, %d bytes in total\n", buf[:num], num)
	if err != nil {
		log.Println(err)
		return nil, -1
	}
	return buf, num
}

func read8bits(sleepTime time.Duration) ([]byte, int) {
	var num, numTemp int
	var buf, bufTemp []byte
	for num = 0; num < 8; {
		bufTemp, numTemp = read(sleepTime)
		if numTemp == -1 {
			continue
		}
		buf = append(buf, bufTemp[:numTemp]...)
		num += numTemp
	}
	log.Printf("Merged 8-bit message: %v", buf)
	return buf[:num], num
}

// write writes data to the serial port
func write(buffer []byte) int {
	num, err := Conn.Write(buffer)
	log.Printf("Data sent: %v, %d bytes in total\n", buffer[:num], num)
	if err != nil {
		log.Println(err)
		return -1
	}
	return num
}

// TestConnection test read and write function
func TestConnection() error {
	log.Printf("Test connection")
	Conn.ResetInputBuffer()
	var err error
	write(TestConnectionSend)
	buf, num := read8bits(500 * time.Millisecond)
	if num == -1 {
		err = errors.New("读取失败，请查看控制台日志")
	} else if num != 8 ||
		buf[0] != ProtocolHead ||
		buf[1] != DeviceSpecialProtocol ||
		buf[2] != KeyTestConnection ||
		buf[3] != KeyActionConnectionOk {
		err = errors.New("握手失败 - 是否连接了正确的设备？")
	}
	if err != nil {
		Conn.Close()
	} else {
		onlineStatus[1] = ((int8)(buf[4]) >> 0) & 0x01
		onlineStatus[2] = ((int8)(buf[4]) >> 1) & 0x01
		onlineStatus[3] = ((int8)(buf[4]) >> 2) & 0x01
		onlineStatus[4] = ((int8)(buf[4]) >> 3) & 0x01
		onlineStatus[5] = ((int8)(buf[4]) >> 4) & 0x01
		onlineStatus[6] = ((int8)(buf[4]) >> 5) & 0x01
		onlineStatus[7] = ((int8)(buf[4]) >> 6) & 0x01
		onlineStatus[8] = ((int8)(buf[4]) >> 7) & 0x01
		log.Printf("Online devices: %v", onlineStatus)
	}
	return err
}

func RealtimeRead() {
	log.Printf("Start listening")
	for {
		EnableReadMtx.Lock()
		if !EnableRead {
			EnableReadMtx.Unlock()
			return
		}
		EnableReadMtx.Unlock()
		buf, num := read8bits(50 * time.Millisecond)
		if num <= 0 || buf[0] != ProtocolHead {
			continue
		}
		if buf[1] == DeviceSpecialProtocol { // special protocol

		} else { // device key
			device := (int)(buf[1])
			key := (int)(buf[2])
			action := (int)(buf[3])
			switch action {
			case KeyActionPress:
				robotgo.KeyDown(keymap[device][key])
				log.Printf("Key %s pressed", keymap[device][key])
			case KeyActionRelease:
				robotgo.KeyUp(keymap[device][key])
				log.Printf("Key %s released", keymap[device][key])
			}
		}
	}
}

// GetDevices 获取在线设备列表
func GetDevices() []int {
	var devices []int
	for i := 1; i < 9; i++ {
		if onlineStatus[i] == 1 {
			devices = append(devices, i)
		}
	}
	return devices
}
