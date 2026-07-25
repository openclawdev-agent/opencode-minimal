package theme

import "github.com/charmbracelet/lipgloss"

// AdaptiveColor represents a color that adapts to the current theme.
type AdaptiveColor = lipgloss.AdaptiveColor

// Theme defines the interface for theme colors.
type Theme interface {
	Primary() AdaptiveColor
	Secondary() AdaptiveColor
	Accent() AdaptiveColor
	Error() AdaptiveColor
	Warning() AdaptiveColor
	Success() AdaptiveColor
	Info() AdaptiveColor

	Text() AdaptiveColor
	TextMuted() AdaptiveColor
	TextEmphasized() AdaptiveColor

	Background() AdaptiveColor
	BackgroundSecondary() AdaptiveColor
	BackgroundDarker() AdaptiveColor

	BorderNormal() AdaptiveColor
	BorderFocused() AdaptiveColor
	BorderDim() AdaptiveColor

	DiffAdded() AdaptiveColor
	DiffRemoved() AdaptiveColor
	DiffContext() AdaptiveColor
	DiffHunkHeader() AdaptiveColor
	DiffHighlightAdded() AdaptiveColor
	DiffHighlightRemoved() AdaptiveColor
	DiffAddedBg() AdaptiveColor
	DiffRemovedBg() AdaptiveColor
	DiffContextBg() AdaptiveColor
	DiffLineNumber() AdaptiveColor
	DiffAddedLineNumberBg() AdaptiveColor
	DiffRemovedLineNumberBg() AdaptiveColor

	MarkdownText() AdaptiveColor
	MarkdownHeading() AdaptiveColor
	MarkdownLink() AdaptiveColor
	MarkdownLinkText() AdaptiveColor
	MarkdownCode() AdaptiveColor
	MarkdownBlockQuote() AdaptiveColor
	MarkdownEmph() AdaptiveColor
	MarkdownStrong() AdaptiveColor
	MarkdownHorizontalRule() AdaptiveColor
	MarkdownListItem() AdaptiveColor
	MarkdownListEnumeration() AdaptiveColor
	MarkdownImage() AdaptiveColor
	MarkdownImageText() AdaptiveColor
	MarkdownCodeBlock() AdaptiveColor

	SyntaxComment() AdaptiveColor
	SyntaxKeyword() AdaptiveColor
	SyntaxFunction() AdaptiveColor
	SyntaxVariable() AdaptiveColor
	SyntaxString() AdaptiveColor
	SyntaxNumber() AdaptiveColor
	SyntaxType() AdaptiveColor
	SyntaxOperator() AdaptiveColor
	SyntaxPunctuation() AdaptiveColor
}

// defaultTheme implements Theme with all methods returning empty colors.
type defaultTheme struct{}

func (defaultTheme) Primary() AdaptiveColor                  { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) Secondary() AdaptiveColor                { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) Accent() AdaptiveColor                   { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) Error() AdaptiveColor                    { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) Warning() AdaptiveColor                  { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) Success() AdaptiveColor                  { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) Info() AdaptiveColor                     { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) Text() AdaptiveColor                     { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) TextMuted() AdaptiveColor                { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) TextEmphasized() AdaptiveColor           { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) Background() AdaptiveColor               { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) BackgroundSecondary() AdaptiveColor      { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) BackgroundDarker() AdaptiveColor         { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) BorderNormal() AdaptiveColor             { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) BorderFocused() AdaptiveColor            { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) BorderDim() AdaptiveColor                { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) DiffAdded() AdaptiveColor                { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) DiffRemoved() AdaptiveColor              { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) DiffContext() AdaptiveColor              { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) DiffHunkHeader() AdaptiveColor           { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) DiffHighlightAdded() AdaptiveColor       { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) DiffHighlightRemoved() AdaptiveColor     { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) DiffAddedBg() AdaptiveColor              { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) DiffRemovedBg() AdaptiveColor            { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) DiffContextBg() AdaptiveColor            { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) DiffLineNumber() AdaptiveColor           { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) DiffAddedLineNumberBg() AdaptiveColor    { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) DiffRemovedLineNumberBg() AdaptiveColor  { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) MarkdownText() AdaptiveColor             { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) MarkdownHeading() AdaptiveColor          { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) MarkdownLink() AdaptiveColor             { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) MarkdownLinkText() AdaptiveColor         { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) MarkdownCode() AdaptiveColor             { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) MarkdownBlockQuote() AdaptiveColor       { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) MarkdownEmph() AdaptiveColor             { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) MarkdownStrong() AdaptiveColor           { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) MarkdownHorizontalRule() AdaptiveColor   { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) MarkdownListItem() AdaptiveColor         { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) MarkdownListEnumeration() AdaptiveColor  { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) MarkdownImage() AdaptiveColor            { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) MarkdownImageText() AdaptiveColor        { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) MarkdownCodeBlock() AdaptiveColor        { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) SyntaxComment() AdaptiveColor            { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) SyntaxKeyword() AdaptiveColor            { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) SyntaxFunction() AdaptiveColor           { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) SyntaxVariable() AdaptiveColor           { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) SyntaxString() AdaptiveColor             { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) SyntaxNumber() AdaptiveColor             { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) SyntaxType() AdaptiveColor               { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) SyntaxOperator() AdaptiveColor           { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }
func (defaultTheme) SyntaxPunctuation() AdaptiveColor        { return lipgloss.AdaptiveColor{Light: "", Dark: ""} }

// CurrentTheme returns the current theme.
func CurrentTheme() Theme { return &defaultTheme{} }

// Init initializes the theme system with the given name.
func Init(name string) {}

// SetTheme sets the current theme by name.
func SetTheme(name string) error { return nil }

// GetTheme returns the current theme.
func GetTheme() Theme { return CurrentTheme() }
