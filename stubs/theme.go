package theme

type Theme struct{}

func (Theme) Get() Theme       { return Theme{} }
func Init(_ string)            {}
func SetTheme(_ string) error  { return nil }
