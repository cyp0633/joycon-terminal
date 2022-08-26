package devices

import (
	"errors"
	"log"
	"time"

	"go.bug.st/serial"
)

var conf = &serial.Mode{
	BaudRate: 4800,
	DataBits: 8,
	Parity:   serial.NoParity,
	StopBits: serial.OneStopBit,
}

var Conn serial.Port

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

func read() ([]byte, int) {
	time.Sleep(time.Millisecond * 500) // wait 500 ms to read into system buffer
	buf := make([]byte, 128)
	num, err := Conn.Read(buf)
	log.Printf("Data received: %v, %d bytes in total\n", buf[:num], num)
	if err != nil {
		log.Println(err)
		return nil, -1
	}
	return buf, num
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
	buf, num := read()
	if num == -1 {
		err = errors.New("读取失败，请查看控制台日志")
	} else if num != 8 || buf[0] != ProtocolHead || buf[1] != DeviceSpecialProtocol || buf[2] != KeyTestConnection || buf[3] != KeyActionConnectionOk {
		err = errors.New("握手失败 - 是否连接了正确的设备？")
	}
	if err != nil {
		Conn.Close()
	}
	return err
}
