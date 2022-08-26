package main

import (
	"context"
	"fmt"

	"github.com/cyp0633/joycon-terminal/devices"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// ConnectSerial sets up connection
func (a *App) ConnectSerial(path string) string {
	// get available ports
	_ = a.GetAvailablePorts()
	// connect port
	err := devices.ConnectSerial(path)
	if err != nil {
		return "连接串口失败：" + err.Error()
	}
	// test connection
	err = devices.TestConnection()
	if err != nil {
		return "连接测试失败：" + err.Error()
	}
	return "success"
}

func (a *App) GetAvailablePorts() string {
	list, err := devices.GetPortList()
	if err != nil {
		return "获取串口列表失败：" + err.Error()
	}
	return fmt.Sprintf("可用端口列表：%v", list)
}

func (a *App) StartListen() {
	devices.EnableReadMtx.Lock()
	devices.EnableRead = true
	devices.EnableReadMtx.Unlock()
	devices.RealtimeRead()
}

func (a *App) StopListen() {
	devices.EnableReadMtx.Lock()
	devices.EnableRead = false
	devices.EnableReadMtx.Unlock()
}

func (a *App) Disconnect() {
	a.StopListen()
	devices.Conn.Close()
}
