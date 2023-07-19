package main

import (
	"boom-gorl/pkg/settings"
	"boom-gorl/pkg/render"
	"boom-gorl/pkg/scenes"
	"log"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func log_init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	log_init()

	// PRE-INIT
	settings_path := "settings.json"
    err := settings.LoadSettings(settings_path)
	if err != nil {
		settings.FallbackSettings()
		InfoLogger.Printf("Failed to load settings from '%s', using fallback!",
			settings_path)
	}

	// INITIALIZATION
	rl.InitWindow(
		int32(settings.CurrentSettings().ScreenWidth),
		int32(settings.CurrentSettings().ScreenHeight), "boom-gorl window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(int32(settings.CurrentSettings().TargetFps))

	rd := render.RenderState{
		RenderResolution: rl.NewVector2(
			float32(settings.CurrentSettings().RenderWidth),
			float32(settings.CurrentSettings().RenderHeight)),
	}

	dev_scene := scenes.DevScene{}
	dev_scene.Init()
	defer dev_scene.Deinit()

	// GAME LOOP
	for !rl.WindowShouldClose() {
		rl.ClearBackground(rl.Black)
        rd.BeginCustomRender()
        rl.BeginDrawing()
            rl.ClearBackground(rl.DarkGreen)

            // Draw GUI
            //dev_scene.DrawGUI()

            dev_scene.Draw()
            rl.DrawFPS(10, 10);
            rl.DrawGrid(10, 1.0);


        rl.EndDrawing()
        rd.EndCustomRender()
	}
}
