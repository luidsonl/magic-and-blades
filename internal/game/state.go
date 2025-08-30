package game

// State representa o estado atual do jogo
type State struct {
	Running bool
}

// NewState cria um novo estado do jogo
func NewState() *State {
	return &State{
		Running: true,
	}
}
