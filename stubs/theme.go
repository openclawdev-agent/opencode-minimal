package theme

// AdaptiveColor represents a color that adapts to the current theme.
type AdaptiveColor = string

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

// defaultTheme implements Theme with all methods returning empty strings.
type defaultTheme struct{}

func (defaultTheme) Primary() AdaptiveColor                  { return "" }
func (defaultTheme) Secondary() AdaptiveColor                { return "" }
func (defaultTheme) Accent() AdaptiveColor                   { return "" }
func (defaultTheme) Error() AdaptiveColor                    { return "" }
func (defaultTheme) Warning() AdaptiveColor                  { return "" }
func (defaultTheme) Success() AdaptiveColor                  { return "" }
func (defaultTheme) Info() AdaptiveColor                     { return "" }
func (defaultTheme) Text() AdaptiveColor                     { return "" }
func (defaultTheme) TextMuted() AdaptiveColor                { return "" }
func (defaultTheme) TextEmphasized() AdaptiveColor           { return "" }
func (defaultTheme) Background() AdaptiveColor               { return "" }
func (defaultTheme) BackgroundSecondary() AdaptiveColor      { return "" }
func (defaultTheme) BackgroundDarker() AdaptiveColor         { return "" }
func (defaultTheme) BorderNormal() AdaptiveColor             { return "" }
func (defaultTheme) BorderFocused() AdaptiveColor            { return "" }
func (defaultTheme) BorderDim() AdaptiveColor                { return "" }
func (defaultTheme) DiffAdded() AdaptiveColor                { return "" }
func (defaultTheme) DiffRemoved() AdaptiveColor              { return "" }
func (defaultTheme) DiffContext() AdaptiveColor              { return "" }
func (defaultTheme) DiffHunkHeader() AdaptiveColor           { return "" }
func (defaultTheme) DiffHighlightAdded() AdaptiveColor       { return "" }
func (defaultTheme) DiffHighlightRemoved() AdaptiveColor     { return "" }
func (defaultTheme) DiffAddedBg() AdaptiveColor              { return "" }
func (defaultTheme) DiffRemovedBg() AdaptiveColor            { return "" }
func (defaultTheme) DiffContextBg() AdaptiveColor            { return "" }
func (defaultTheme) DiffLineNumber() AdaptiveColor           { return "" }
func (defaultTheme) DiffAddedLineNumberBg() AdaptiveColor    { return "" }
func (defaultTheme) DiffRemovedLineNumberBg() AdaptiveColor  { return "" }
func (defaultTheme) MarkdownText() AdaptiveColor             { return "" }
func (defaultTheme) MarkdownHeading() AdaptiveColor          { return "" }
func (defaultTheme) MarkdownLink() AdaptiveColor             { return "" }
func (defaultTheme) MarkdownLinkText() AdaptiveColor         { return "" }
func (defaultTheme) MarkdownCode() AdaptiveColor             { return "" }
func (defaultTheme) MarkdownBlockQuote() AdaptiveColor       { return "" }
func (defaultTheme) MarkdownEmph() AdaptiveColor             { return "" }
func (defaultTheme) MarkdownStrong() AdaptiveColor           { return "" }
func (defaultTheme) MarkdownHorizontalRule() AdaptiveColor   { return "" }
func (defaultTheme) MarkdownListItem() AdaptiveColor         { return "" }
func (defaultTheme) MarkdownListEnumeration() AdaptiveColor  { return "" }
func (defaultTheme) MarkdownImage() AdaptiveColor            { return "" }
func (defaultTheme) MarkdownImageText() AdaptiveColor        { return "" }
func (defaultTheme) MarkdownCodeBlock() AdaptiveColor        { return "" }
func (defaultTheme) SyntaxComment() AdaptiveColor            { return "" }
func (defaultTheme) SyntaxKeyword() AdaptiveColor            { return "" }
func (defaultTheme) SyntaxFunction() AdaptiveColor           { return "" }
func (defaultTheme) SyntaxVariable() AdaptiveColor           { return "" }
func (defaultTheme) SyntaxString() AdaptiveColor             { return "" }
func (defaultTheme) SyntaxNumber() AdaptiveColor             { return "" }
func (defaultTheme) SyntaxType() AdaptiveColor               { return "" }
func (defaultTheme) SyntaxOperator() AdaptiveColor           { return "" }
func (defaultTheme) SyntaxPunctuation() AdaptiveColor        { return "" }

// CurrentTheme is the package-level theme variable.
var CurrentTheme Theme = defaultTheme{}

// Init initializes the theme system with the given name.
func Init(name string) {}

// SetTheme sets the current theme by name.
func SetTheme(name string) error { return nil }

// GetTheme returns the current theme.
func GetTheme() Theme { return CurrentTheme }
