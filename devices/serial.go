package devices

import (
	"errors"
	"log"
	"time"

	"github.com/tarm/serial"
)

var conf = &serial.Config{
	Baud:        4800,
	ReadTimeout: time.Second,
}

var Conn *serial.Port

// ConnectSerial connects to the serial port
func ConnectSerial(path string) error {
	conf.Name = path
	conn, err := serial.OpenPort(conf)
	Conn = conn
	log.Printf("Connected to port %s", path)
	if err != nil {
		log.Print(err.Error())
	}
	return err
}

// read fetches data from the serial port
// Return: size of the data read, error
func read() ([]byte, int) {
	buf := make([]byte, 128)
	num, err := Conn.Read(buf)
	if err != nil {
		log.Println(err)
		return nil, -1
	}
	return buf, num
}

// write writes data to the serial port
func write(buffer []byte) int {
	num, err := Conn.Write(buffer)
	if err != nil {
		log.Println(err)
		return -1
	}
	return num
}

// TestConnection test read and write function
func TestConnection() error {
	write([]byte{ProtocolHead, DeviceSpecialProtocol, KeyTestConnection, KeyActionRequireTest, 0x00, 0x00, 0x00, 0x00})
	buf, num := read()
	log.Printf("Test data received: %v, %d bytes in total\n", buf, num)
	if num == -1 {
		return errors.New("读取失败，请查看控制台日志")
	}
	if num != 8 || (buf[0] != ProtocolHead || buf[1] != DeviceSpecialProtocol || buf[2] != KeyTestConnection || buf[3] != KeyActionConnectionOk) {
		return errors.New("握手失败 - 是否连接了正确的设备？")
	}
	return nil
}
