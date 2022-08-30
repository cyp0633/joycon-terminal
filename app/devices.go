package app

import (
	"log"
	"strconv"

	"github.com/cyp0633/joycon-terminal/devices"
)

func GetOnlineDevices() []SerialOptions {
	var serialOptionsList = SerialOptions{
		Type:  "group",
		Label: "可用设备编号",
		Key:   "devices",
	}
	list := devices.GetDevices()
	log.Printf("list: %v", list)
	for _, device := range list {
		var children = SerialChoice{
			Label:    "设备" + strconv.Itoa(device),
			Value:    strconv.Itoa(device),
			Disabled: false,
		}
		serialOptionsList.Children = append(serialOptionsList.Children, children)
	}
	if len(serialOptionsList.Children) == 0 {
		var children = SerialChoice{
			Label:    "无可用设备",
			Value:    "",
			Disabled: true,
		}
		serialOptionsList.Children = append(serialOptionsList.Children, children)
	}
	return []SerialOptions{serialOptionsList}
}
