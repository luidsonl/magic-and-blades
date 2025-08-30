package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/luidsonl/magic-and-blades/internal/engine"
	"github.com/luidsonl/magic-and-blades/internal/game"
	"github.com/veandco/go-sdl2/sdl"
)

func init() {
	// OpenGL requires operations to be executed on the main thread
	runtime.LockOSThread()
}

func main() {
	fmt.Println("Starting Magic and Blades...")

	// Initialize SDL
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf("Failed to initialize SDL: %v", err)
	}
	defer sdl.Quit()

	// Game configuration
	config := game.Config{
		WindowTitle:  "Magic and Blades",
		WindowWidth:  800,
		WindowHeight: 600,
		Fullscreen:   false,
		Language:     "", // Auto-detect language
	}

	// Initialize game engine
	gameEngine, err := engine.NewEngine(config)
	if err != nil {
		log.Fatalf("Failed to create engine: %v", err)
	}
	defer gameEngine.Destroy()

	// Main game loop
	gameEngine.Run()
}
