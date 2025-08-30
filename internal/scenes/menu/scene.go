package menu

import (
	"github.com/luidsonl/magic-and-blades/internal/game"
	"github.com/luidsonl/magic-and-blades/internal/i18n"
	"github.com/veandco/go-sdl2/sdl"
)

// Scene interface for menu scenes
type Scene interface {
	Update()
	Render()
	ProcessEvent(event sdl.Event) bool
	Cleanup()
}

// SceneType represents different scene types
type SceneType int

const (
	SceneMainMenu SceneType = iota
	SceneSettings
	SceneGameplay
	ScenePause
)

// SceneManager manages scene transitions
type SceneManager struct {
	currentScene Scene
	sceneType    SceneType
	translator   i18n.Translator
	config       *game.Config
	renderer     *sdl.Renderer
}

// NewSceneManager creates a new scene manager
func NewSceneManager(translator i18n.Translator, config *game.Config, renderer *sdl.Renderer) *SceneManager {
	return &SceneManager{
		translator: translator,
		config:     config,
		renderer:   renderer,
	}
}

// SwitchTo changes the current scene
func (sm *SceneManager) SwitchTo(sceneType SceneType) {
	if sm.currentScene != nil {
		sm.currentScene.Cleanup()
	}

	switch sceneType {
	case SceneMainMenu:
		sm.currentScene = NewMenuScene(sm.translator, sm.config, sm.renderer)
		sm.sceneType = SceneMainMenu
	case SceneSettings:
		// You can create a dedicated settings scene if needed
		sm.sceneType = SceneSettings
	case SceneGameplay:
		// Initialize gameplay scene
		sm.sceneType = SceneGameplay
	case ScenePause:
		// Initialize pause menu
		sm.sceneType = ScenePause
	}
}

// GetCurrentScene returns the current scene
func (sm *SceneManager) GetCurrentScene() Scene {
	return sm.currentScene
}

// GetSceneType returns the current scene type
func (sm *SceneManager) GetSceneType() SceneType {
	return sm.sceneType
}

// Update updates the current scene
func (sm *SceneManager) Update() {
	if sm.currentScene != nil {
		sm.currentScene.Update()
	}
}

// Render renders the current scene
func (sm *SceneManager) Render() {
	if sm.currentScene != nil {
		sm.currentScene.Render()
	}
}

// ProcessEvent processes events for the current scene
func (sm *SceneManager) ProcessEvent(event sdl.Event) bool {
	if sm.currentScene != nil {
		return sm.currentScene.ProcessEvent(event)
	}
	return false
}
