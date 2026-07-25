package protocol

type DocumentUri string

func (uri DocumentUri) Path() string { return string(uri) }

type URI = string

type DiagnosticSeverity uint32

const (
	SeverityError       DiagnosticSeverity = 1
	SeverityWarning     DiagnosticSeverity = 2
	SeverityInformation DiagnosticSeverity = 3
	SeverityHint        DiagnosticSeverity = 4
)

type DiagnosticTag uint32

const (
	Unnecessary DiagnosticTag = 1
	Deprecated  DiagnosticTag = 2
)

type Position struct {
	Line      uint32 `json:"line"`
	Character uint32 `json:"character"`
}

type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

type Diagnostic struct {
	Range    Range              `json:"range"`
	Severity DiagnosticSeverity `json:"severity,omitempty"`
	Code     interface{}        `json:"code,omitempty"`
	Source   string             `json:"source,omitempty"`
	Message  string             `json:"message"`
	Tags     []DiagnosticTag    `json:"tags,omitempty"`
}

type PublishDiagnosticsParams struct {
	URI         DocumentUri  `json:"uri"`
	Version     int32        `json:"version,omitempty"`
	Diagnostics []Diagnostic `json:"diagnostics"`
}

type InitializeResult struct{}
