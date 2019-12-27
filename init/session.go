package init

type session struct {
	Provider string `yaml:"Provider"`
	Path     string `yaml:"Path"`
	Domain   string `yaml:"Domain"`
	Secret   string `yaml:"Secret"`
	MaxAge   int    `yaml:"MaxAge"`
	HttpOnly bool   `yaml:"HttpOnly"`
}
