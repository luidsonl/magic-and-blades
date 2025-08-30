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
	// OpenGL requer que as operações sejam executadas na thread principal
	runtime.LockOSThread()
}

func main() {
	fmt.Println("Iniciando Magic and Blades...")

	// Inicializar SDL
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf("Falha ao inicializar SDL: %v", err)
	}
	defer sdl.Quit()

	// Configurações do jogo
	config := game.Config{
		WindowTitle:  "Magic and Blades - Hello World",
		WindowWidth:  800,
		WindowHeight: 600,
		Fullscreen:   false,
	}

	// Inicializar motor do jogo
	gameEngine, err := engine.NewEngine(config)
	if err != nil {
		log.Fatalf("Falha ao criar engine: %v", err)
	}
	defer gameEngine.Destroy()

	// Loop principal do jogo
	gameEngine.Run()
}
