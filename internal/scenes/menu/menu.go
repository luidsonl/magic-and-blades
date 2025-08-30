package menu

import (
	"log"

	"github.com/luidsonl/magic-and-blades/internal/game"
	"github.com/luidsonl/magic-and-blades/internal/i18n"
	"github.com/veandco/go-sdl2/sdl"
)

// MenuScene represents the main menu scene
type MenuScene struct {
	translator     i18n.Translator
	config         *game.Config
	renderer       *sdl.Renderer
	font           *sdl.Texture
	menuItems      []string
	settingsItems  []string
	currentMenu    string
	selectedIndex  int
	resolutionOpts []string
	languageOpts   []string
}

// NewMenuScene creates a new menu scene
func NewMenuScene(translator i18n.Translator, config *game.Config, renderer *sdl.Renderer) *MenuScene {
	menu := &MenuScene{
		translator:    translator,
		config:        config,
		renderer:      renderer,
		currentMenu:   "main",
		selectedIndex: 0,
	}

	// Initialize menu items (will be translated in UpdateMenuText)
	menu.menuItems = []string{
		"button.play",
		"button.options",
		"button.quit",
	}

	menu.settingsItems = []string{
		"settings.language",
		"settings.resolution",
		"settings.back",
	}

	// Available options
	menu.resolutionOpts = []string{
		"800x600",
		"1024x768",
		"1280x720",
		"1366x768",
		"1920x1080",
	}

	menu.languageOpts = []string{
		"English",
		"Português",
		"Español",
		"Français",
	}

	return menu
}

// Update processes menu logic
func (m *MenuScene) Update() {
	// Menu logic will be handled in ProcessEvent
}

// Render draws the menu
func (m *MenuScene) Render() {
	m.renderer.SetDrawColor(0, 0, 0, 255)
	m.renderer.Clear()

	// Draw background
	m.renderer.SetDrawColor(30, 30, 50, 255)
	m.renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: m.config.WindowWidth, H: m.config.WindowHeight})

	// Draw title
	title := m.translator.Translate(i18n.TitleMainMenu)
	m.drawText(title, m.config.WindowWidth/2, 100, 48, true)

	// Draw menu items
	var items []string
	if m.currentMenu == "main" {
		items = m.menuItems
	} else if m.currentMenu == "settings" {
		items = m.settingsItems
	} else if m.currentMenu == "language" {
		items = m.languageOpts
	} else if m.currentMenu == "resolution" {
		items = m.resolutionOpts
	}

	for i, itemKey := range items {
		var text string
		if m.currentMenu == "main" || m.currentMenu == "settings" {
			text = m.translator.Translate(itemKey)
		} else {
			text = itemKey
		}

		yPos := int32(200 + i*60) // Convert to int32

		if i == m.selectedIndex {
			// Gold for selected item
			m.drawText("> "+text, m.config.WindowWidth/2-20, yPos, 32, false)
		} else {
			m.drawText(text, m.config.WindowWidth/2, yPos, 32, false)
		}
	}

	m.renderer.Present()
}

// ProcessEvent handles input events
func (m *MenuScene) ProcessEvent(event sdl.Event) bool {
	switch e := event.(type) {
	case *sdl.KeyboardEvent:
		if e.State == sdl.PRESSED {
			switch e.Keysym.Sym {
			case sdl.K_UP:
				m.moveSelection(-1)
				return true
			case sdl.K_DOWN:
				m.moveSelection(1)
				return true
			case sdl.K_RETURN, sdl.K_SPACE:
				m.selectItem()
				return true
			case sdl.K_ESCAPE:
				if m.currentMenu != "main" {
					m.currentMenu = "main"
					m.selectedIndex = 0
					return true
				}
			}
		}
	}
	return false
}

// moveSelection changes the selected menu item
func (m *MenuScene) moveSelection(direction int) {
	var maxItems int
	switch m.currentMenu {
	case "main":
		maxItems = len(m.menuItems)
	case "settings":
		maxItems = len(m.settingsItems)
	case "language":
		maxItems = len(m.languageOpts)
	case "resolution":
		maxItems = len(m.resolutionOpts)
	}

	m.selectedIndex = (m.selectedIndex + direction + maxItems) % maxItems
}

// selectItem handles menu item selection
func (m *MenuScene) selectItem() {
	switch m.currentMenu {
	case "main":
		switch m.selectedIndex {
		case 0: // Play
			log.Println("Starting game...")
			// Transition to gameplay scene
		case 1: // Options
			m.currentMenu = "settings"
			m.selectedIndex = 0
		case 2: // Quit
			log.Println("Quitting game...")
			// Signal to quit the game
		}
	case "settings":
		switch m.selectedIndex {
		case 0: // Language
			m.currentMenu = "language"
			m.selectedIndex = 0
		case 1: // Resolution
			m.currentMenu = "resolution"
			m.selectedIndex = 0
		case 2: // Back
			m.currentMenu = "main"
			m.selectedIndex = 0
		}
	case "language":
		switch m.selectedIndex {
		case 0: // English
			m.translator.SetLanguage("en")
			m.config.Language = "en"
		case 1: // Portuguese
			m.translator.SetLanguage("pt")
			m.config.Language = "pt"
		case 2: // Spanish
			// Add support for more languages as needed
		case 3: // French
			// Add support for more languages as needed
		}
		m.currentMenu = "settings"
	case "resolution":
		switch m.selectedIndex {
		case 0: // 800x600
			m.config.WindowWidth = 800
			m.config.WindowHeight = 600
		case 1: // 1024x768
			m.config.WindowWidth = 1024
			m.config.WindowHeight = 768
		case 2: // 1280x720
			m.config.WindowWidth = 1280
			m.config.WindowHeight = 720
		case 3: // 1366x768
			m.config.WindowWidth = 1366
			m.config.WindowHeight = 768
		case 4: // 1920x1080
			m.config.WindowWidth = 1920
			m.config.WindowHeight = 1080
		}
		// In a real implementation, you would recreate the window here
		log.Printf("Resolution changed to: %dx%d", m.config.WindowWidth, m.config.WindowHeight)
		m.currentMenu = "settings"
	}
}

// drawText draws text on the screen (simplified implementation)
func (m *MenuScene) drawText(text string, x, y int32, size int32, centered bool) {
	// This is a simplified text rendering function
	// In a real implementation, you would use a proper font rendering library

	// For now, we'll just log the text that would be displayed
	log.Printf("Would draw text: %s at (%d, %d)", text, x, y)

	// Placeholder: actual SDL text rendering would go here
	// You would typically use SDL_ttf for proper text rendering
}

// Cleanup releases resources
func (m *MenuScene) Cleanup() {
	if m.font != nil {
		m.font.Destroy()
	}
}
