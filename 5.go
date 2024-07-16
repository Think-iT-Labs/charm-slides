var (
	SOME_TEXT = `
# Hello
- [https://google.com](Google)
> This is a quote
`
)

// In the view function
// dark is a "stylePath", I'm not sure what it is
out, err := glamour.Render(INTRO, "dark")
