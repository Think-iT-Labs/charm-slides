type Model interface {
	// What to start the program with.
	Init() Cmd

	// What happens when something happens (like a key press).
	Update(Msg) (Model, Cmd)

	// How to render the view.
	View() string
}
