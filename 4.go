var (
	textStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4"))
)

// In the view function
s := ""
// after some stuff..
s += textStyle.Render("HELLO")
