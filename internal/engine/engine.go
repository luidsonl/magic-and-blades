package engine

import (
	"fmt"
	"log"

	"github.com/luidsonl/magic-and-blades/internal/game"
	"github.com/luidsonl/magic-and-blades/internal/i18n"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/veandco/go-sdl2/sdl"
)

// Engine represents the main game engine
type Engine struct {
	window     *sdl.Window
	context    sdl.GLContext
	config     game.Config
	state      *game.GameState
	translator i18n.Translator
}

// NewEngine creates a new instance of the game engine
func NewEngine(config game.Config) (*Engine, error) {
	// Configure OpenGL attributes before creating the window
	if err := sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 3); err != nil {
		return nil, fmt.Errorf("failed to set OpenGL major version: %v", err)
	}
	if err := sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 3); err != nil {
		return nil, fmt.Errorf("failed to set OpenGL minor version: %v", err)
	}
	if err := sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE); err != nil {
		return nil, fmt.Errorf("failed to set OpenGL profile: %v", err)
	}
	if err := sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1); err != nil {
		return nil, fmt.Errorf("failed to enable double buffering: %v", err)
	}

	// Create window
	window, err := sdl.CreateWindow(
		config.WindowTitle,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		config.WindowWidth, config.WindowHeight,
		sdl.WINDOW_OPENGL|sdl.WINDOW_SHOWN,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create window: %v", err)
	}

	// Create OpenGL context
	context, err := window.GLCreateContext()
	if err != nil {
		window.Destroy()
		return nil, fmt.Errorf("failed to create OpenGL context: %v", err)
	}

	// Initialize game state
	state := game.NewState()

	// Initialize internationalization system
	var translator i18n.Translator
	var initErr error

	if config.Language != "" {
		// Use specified language
		translator, initErr = i18n.NewWithLanguage(config.Language)
	} else {
		// Auto-detect system language
		translator, initErr = i18n.New()
	}

	if initErr != nil {
		log.Printf("Warning: Failed to initialize translator: %v", initErr)
		// Create fallback translator
		translator = &i18n.FallbackTranslator{}
	}

	engine := &Engine{
		window:     window,
		context:    context,
		config:     config,
		state:      state,
		translator: translator,
	}

	// Initialize OpenGL
	if err := engine.initOpenGL(); err != nil {
		engine.Destroy()
		return nil, err
	}

	// Log successful initialization
	log.Printf("Engine initialized successfully")
	log.Printf("OpenGL context created")
	log.Printf("Language set to: %s", translator.GetLanguage())
	log.Printf("Initial scene: %s", state.CurrentScene)

	return engine, nil
}

// Run starts the main game loop
func (e *Engine) Run() {
	defer e.Destroy()

	// Log game start
	log.Printf("Starting game loop")

	// Main game loop
	for e.state.Running {
		// Process events
		e.processEvents()

		// Update game logic based on current scene
		e.update()

		// Render scene
		e.render()

		// Swap buffers
		e.window.GLSwap()

		// Cap frame rate
		sdl.Delay(16) // ~60 FPS
	}

	log.Printf("Game loop ended")
}

// processEvents processes all pending events
func (e *Engine) processEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch evt := event.(type) {
		case *sdl.QuitEvent:
			e.state.Running = false
			log.Printf("Quit event received")
		case *sdl.KeyboardEvent:
			if evt.State == sdl.PRESSED {
				switch evt.Keysym.Sym {
				case sdl.K_ESCAPE:
					// Handle ESC based on current scene
					if e.state.CurrentScene == "menu" {
						e.state.Running = false
						log.Printf("Escape key pressed, exiting game")
					} else {
						// Return to menu from other scenes
						e.state.CurrentScene = "menu"
						log.Printf("Returning to menu")
					}
				}
			}
		case *sdl.WindowEvent:
			if evt.Event == sdl.WINDOWEVENT_RESIZED {
				log.Printf("Window resized: %dx%d", evt.Data1, evt.Data2)
				// Update viewport size
				gl.Viewport(0, 0, evt.Data1, evt.Data2)
			}
		}
	}
}

// update handles game logic updates based on current scene
func (e *Engine) update() {
	// Scene-specific update logic would go here
	// For now, we just log the current scene
	switch e.state.CurrentScene {
	case "menu":
		// Menu scene logic
	case "gameplay":
		// Gameplay logic
	case "pause":
		// Pause menu logic
	}
}

// render renders the game scene based on current state
func (e *Engine) render() {
	// Clear color and depth buffers
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	// Scene-specific rendering
	switch e.state.CurrentScene {
	case "menu":
		e.renderMenu()
	case "gameplay":
		e.renderGameplay()
	case "pause":
		e.renderPauseMenu()
	default:
		// Default background
		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)
	}
}

// renderMenu renders the main menu scene
func (e *Engine) renderMenu() {
	gl.ClearColor(0.1, 0.1, 0.2, 1.0) // Dark blue background
	gl.Clear(gl.COLOR_BUFFER_BIT)

	// Menu rendering logic will be implemented here
	// For now, just clear with a different color
}

// renderGameplay renders the gameplay scene
func (e *Engine) renderGameplay() {
	gl.ClearColor(0.3, 0.5, 0.3, 1.0) // Green background for gameplay
	gl.Clear(gl.COLOR_BUFFER_BIT)

	// Gameplay rendering logic will be implemented here
}

// renderPauseMenu renders the pause menu
func (e *Engine) renderPauseMenu() {
	gl.ClearColor(0.2, 0.2, 0.2, 0.8) // Semi-transparent dark background
	gl.Clear(gl.COLOR_BUFFER_BIT)

	// Pause menu rendering logic will be implemented here
}

// initOpenGL initializes OpenGL
func (e *Engine) initOpenGL() error {
	// Initialize OpenGL
	if err := gl.Init(); err != nil {
		return fmt.Errorf("failed to initialize OpenGL: %v", err)
	}

	// Display OpenGL version information
	version := gl.GoStr(gl.GetString(gl.VERSION))
	vendor := gl.GoStr(gl.GetString(gl.VENDOR))
	renderer := gl.GoStr(gl.GetString(gl.RENDERER))

	log.Printf("OpenGL version: %s", version)
	log.Printf("OpenGL vendor: %s", vendor)
	log.Printf("OpenGL renderer: %s", renderer)

	// Set basic OpenGL settings
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	// Set viewport size
	gl.Viewport(0, 0, e.config.WindowWidth, e.config.WindowHeight)

	// Set default clear color
	gl.ClearColor(0.2, 0.3, 0.3, 1.0)

	return nil
}

// Destroy releases engine resources
func (e *Engine) Destroy() {
	log.Printf("Destroying engine resources")

	if e.context != nil {
		sdl.GLDeleteContext(e.context)
		log.Printf("OpenGL context destroyed")
	}

	if e.window != nil {
		e.window.Destroy()
		log.Printf("Window destroyed")
	}

	log.Printf("Engine cleanup completed")
}

// GetTranslator returns the translator instance
func (e *Engine) GetTranslator() i18n.Translator {
	return e.translator
}

// GetConfig returns the game configuration
func (e *Engine) GetConfig() game.Config {
	return e.config
}

// GetState returns the game state
func (e *Engine) GetState() *game.GameState {
	return e.state
}

// SetScene changes the current scene
func (e *Engine) SetScene(sceneName string) {
	e.state.CurrentScene = sceneName
	log.Printf("Scene changed to: %s", sceneName)
}
