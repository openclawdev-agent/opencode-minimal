package theme

// 精简版 stub：移除UI主题依赖
// 原版依赖 charmbracelet/lipgloss，精简版不需要

type AdaptiveColor = string

// Theme interface stub
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

// SetTheme stub - 精简版不需要主题切换
func SetTheme(name string) error {
	return nil
}

// GetTheme stub
func GetTheme() Theme {
	return nil
}
