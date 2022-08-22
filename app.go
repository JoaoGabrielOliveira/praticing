package main

import (
	"context"
	"fmt"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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

func (a *App) SaveNote(text string) {
	selection, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title:           "Salvando documento",
		DefaultFilename: ".txt",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Texto (*.txt",
				Pattern:     "*.txt",
			}, {
				DisplayName: "Todos (*.*",
				Pattern:     "*.*",
			},
		},
	})

	if err != nil {
		panic(err)
	}

	if len(selection) == 0 {
		runtime.LogInfo(a.ctx, "Operation was canceled")
		return
	}

	d := []byte(text)
	os.WriteFile(selection, d, 0644)
	runtime.LogInfo(a.ctx, "A file as saved: "+selection)
}
