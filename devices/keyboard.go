package devices

import "github.com/go-vgo/robotgo"

func keytap(key string) {
	robotgo.KeyTap(key)
	robotgo.KeyDown(key)
	robotgo.KeyUp(key)
}
