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

func (a *App) ConnectSerial(path string) string {
	err := devices.ConnectSerial(path)
	if err != nil {
		return err.Error()
	}
	return "success"
}
