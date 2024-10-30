package main

import (
	"embed"
	_ "embed"
	"github.com/wailsapp/wails/v3/pkg/application"
	"log"
	"main/core"
)

//go:embed frontend/build
var assets embed.FS

func main() {

	testStruct := core.TestStruct{
		Name: "Hello from wails with Sveltekit !",
	}

	app := application.New(application.Options{
		Name:        "Dogu",
		Description: "A demo of using raw HTML & CSS",
		Services: []application.Service{
			application.NewService(&testStruct),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})
	app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: "Main Window",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
	})

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
