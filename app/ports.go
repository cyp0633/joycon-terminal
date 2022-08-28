package app

import "github.com/cyp0633/joycon-terminal/devices"

type SerialChoice struct {
	Label    string `json:"label"` // use json binding to pass to frontend
	Value    string `json:"value"`
	Disabled bool   `json:"disabled"`
}

type SerialOptions struct {
	Type     string         `json:"type"`
	Label    string         `json:"label"`
	Key      string         `json:"key"`
	Children []SerialChoice `json:"children"`
}

func GetSerialPortsList() []SerialOptions {
	var serialOptionsList = SerialOptions{
		Type:  "group",
		Label: "串口选择",
		Key:   "serial",
	}
	list, _ := devices.GetPortList()
	for _, port := range list {
		var children = SerialChoice{
			Label:    port,
			Value:    port,
			Disabled: false,
		}
		serialOptionsList.Children = append(serialOptionsList.Children, children)
	}
	if len(serialOptionsList.Children) == 0 {
		var children = SerialChoice{
			Label:    "无可用串口",
			Value:    "",
			Disabled: true,
		}
		serialOptionsList.Children = append(serialOptionsList.Children, children)
	}
	return []SerialOptions{serialOptionsList}
}
