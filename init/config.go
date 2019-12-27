package init

type config struct {
	RunMode     string      `yaml:"RunMode"`
	AppName     string      `yaml:"AppName"`
	StaticDir   string      `yaml:"StaticDir"`
	TemplateDir string      `yaml:"TemplateDir"`
	Casbin      *Casbin     `yaml:"casbin"`
	// HttpServer  *HttpServer `yaml:"http_server"`
	// Fix         *Fix        `yaml:"Fix"`
	// Log         *Log        `yaml:"Log"`
	// DataBase    []*Database `yaml:"DataBase"`
	// Session     *session    `yaml:"Session"`
	// Redis       *redis      `yaml:"Redis"`
}
