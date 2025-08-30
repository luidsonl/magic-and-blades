package engine

import (
	"fmt"

	"github.com/luidsonl/magic-and-blades/internal/game"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/veandco/go-sdl2/sdl"
)

// Engine representa o motor principal do jogo
type Engine struct {
	window  *sdl.Window
	context sdl.GLContext
	config  game.Config
	state   *game.State
}

// NewEngine cria uma nova instância do motor do jogo
func NewEngine(config game.Config) (*Engine, error) {
	// Configurar atributos do OpenGL antes de criar a janela
	if err := sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 3); err != nil {
		return nil, fmt.Errorf("falha ao definir versão major do OpenGL: %v", err)
	}
	if err := sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 3); err != nil {
		return nil, fmt.Errorf("falha ao definir versão minor do OpenGL: %v", err)
	}
	if err := sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE); err != nil {
		return nil, fmt.Errorf("falha ao definir perfil do OpenGL: %v", err)
	}
	if err := sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1); err != nil {
		return nil, fmt.Errorf("falha ao habilitar double buffering: %v", err)
	}

	// Criar janela
	window, err := sdl.CreateWindow(
		config.WindowTitle,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		config.WindowWidth, config.WindowHeight,
		sdl.WINDOW_OPENGL|sdl.WINDOW_SHOWN,
	)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar janela: %v", err)
	}

	// Criar contexto OpenGL
	context, err := window.GLCreateContext()
	if err != nil {
		window.Destroy()
		return nil, fmt.Errorf("falha ao criar contexto OpenGL: %v", err)
	}

	// Inicializar estado do jogo
	state := game.NewState()

	engine := &Engine{
		window:  window,
		context: context,
		config:  config,
		state:   state,
	}

	// Inicializar OpenGL
	if err := engine.initOpenGL(); err != nil {
		engine.Destroy()
		return nil, err
	}

	return engine, nil
}

// Run inicia o loop principal do jogo
func (e *Engine) Run() {
	defer e.Destroy()

	// Loop principal do jogo
	for e.state.Running {
		// Processar eventos
		e.processEvents()

		// Renderizar cena
		e.render()

		// Trocar buffers
		e.window.GLSwap()
	}
}

// processEvents processa todos os eventos pendentes
func (e *Engine) processEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			e.state.Running = false
		case *sdl.KeyboardEvent:
			keyEvent := event.(*sdl.KeyboardEvent)
			if keyEvent.Keysym.Sym == sdl.K_ESCAPE && keyEvent.State == sdl.PRESSED {
				e.state.Running = false
			}
		}
	}
}

// render renderiza a cena do jogo
func (e *Engine) render() {
	// Limpar buffer de cor
	gl.ClearColor(0.2, 0.3, 0.3, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

// initOpenGL inicializa o OpenGL
func (e *Engine) initOpenGL() error {
	// Inicializar OpenGL
	if err := gl.Init(); err != nil {
		return fmt.Errorf("falha ao inicializar OpenGL: %v", err)
	}

	// Exibir informações da versão do OpenGL
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Printf("OpenGL versão: %s\n", version)

	return nil
}

// Destroy libera recursos da engine
func (e *Engine) Destroy() {
	if e.context != nil {
		sdl.GLDeleteContext(e.context)
	}
	if e.window != nil {
		e.window.Destroy()
	}
}
