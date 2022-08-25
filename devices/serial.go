package devices

import (
	"log"
	"runtime"
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
	log.Printf("Connected to port %s\n", path)
	return err
}

// read fetches data from the serial port
// Return: size of the data read, error
func read() ([]byte, int) {
	buf := make([]byte, 128)
	num, err := Conn.Read(buf)
	if err != nil {
		log.Println(err)
		return nil, 0
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
func TestConnection() {
	if runtime.GOOS == "windows" {
		ConnectSerial("COM3")
	} else {
		ConnectSerial("/dev/ttyUSB0")
	}
	write([]byte{0x55, 0xaa, 0xff, 0x01, 0x00, 0x00, 0x00, 0x00})
	buf, num := read()
	log.Printf("Test data received: %v, %d bytes in total\n", buf, num)
}
