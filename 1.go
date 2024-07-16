import tea "github.com/charmbracelet/bubbletea"

// Model is the Bubble Tea model for the application.
type Model struct{}

// Implement the Model interface.
// ...

func newModel() Model {
	return Model{
		// ... initialize the model with whatever it needs
	}
}

func main() {
	p := tea.NewProgram(newModel())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
