package game

// GameState represents the overall game state
type GameState struct {
	Running      bool
	CurrentScene string
	// Add other game state variables as needed
}

// NewState creates a new game state
func NewState() *GameState {
	return &GameState{
		Running:      true,
		CurrentScene: "menu",
	}
}
