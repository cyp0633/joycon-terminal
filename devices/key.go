package devices

import (
	"log"

	"github.com/go-vgo/robotgo"
)

func pressKey(key string) {
	switch key {
	case "":
		return
	case "lclick":
		robotgo.Click("left", false)
	case "rclick":
		robotgo.Click("right", false)
	default:
		robotgo.KeyDown(key)
	}
	log.Printf("Key %s pressed", key)

}

func releaseKey(key string) {

	switch key {
	case "":
		return
	case "lclick":
	case "rclick":
	default:
		robotgo.KeyUp(key)
	}
	log.Printf("Key %s released", key)
}

func SetKey(device, key int, target string) {
	log.Printf("Set key %d of device %d to %s", key, device, target)
	keymap[device][key] = target
}

func UsePreset(preset string) {
	log.Printf("Use preset %s", preset)
	keymap = keymap_preset[preset]
}
