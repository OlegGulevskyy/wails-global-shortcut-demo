package main

import (
	"context"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
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

	// register hotkey on the app startup
	// if you try to register it anywhere earlier - the app will hang on compile step
	mainthread.Init(a.RegisterHotKey)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// just a wrapper to have access to App functions
// not necessary if you don't plan to do anything with your App on shortcut use
func (a *App) RegisterHotKey() {
	registerHotkey(a)
}

func registerHotkey(a *App) {
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyS)
	err := hk.Register()
	if err != nil {
		return
	}

	// you have 2 events available - Keyup and Keydown
	// you can either or neither, or both
	fmt.Printf("hotkey: %v is registered\n", hk)
	<-hk.Keydown()
	// do anything you want on Key down event
	fmt.Printf("hotkey: %v is down\n", hk)

	<-hk.Keyup()
	// do anything you want on Key up event
	fmt.Printf("hotkey: %v is up\n", hk)

	runtime.EventsEmit(a.ctx, "Backend:GlobalHotkeyEvent", time.Now().String())

	hk.Unregister()
	fmt.Printf("hotkey: %v is unregistered\n", hk)

	// reattach listener
	registerHotkey(a)
}
