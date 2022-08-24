package devices

import (
	"time"

	"github.com/tarm/serial"
)

var conf = &serial.Config{
	Baud:        4800,
	ReadTimeout: time.Second,
}

var Conn *serial.Port

func ConnectSerial(path string) error {
	conf.Name = path
	conn, err := serial.OpenPort(conf)
	Conn = conn
	return err
}
