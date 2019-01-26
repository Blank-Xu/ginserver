package config

type config struct {
	RunMode    string      `yaml:"RunMode"`
	AppName    string      `yaml:"AppName"`
	AssetsFile string      `yaml:"AssetsFile"`
	ViewFile   string      `yaml:"ViewFile"`
	RbacFile   string      `yaml:"RbacFile"`
	Server     *server     `yaml:"Server"`
	Fix        *fix        `yaml:"Fix"`
	Log        *log        `yaml:"Log"`
	DataBase   []*DataBase `yaml:"DataBase"`
	Session    *session    `yaml:"Session"`
	Redis      *redis      `yaml:"Redis"`
	Lang       *lang       `yaml:"Lang"`
}

type server struct {
	HttpPort     string `yaml:"HttpPort"`
	ReadTimeout  int    `yaml:"ReadTimeout"`
	WriteTimeout int    `yaml:"WriteTimeout"`
}

type fix struct {
	TimeZone *struct {
		Name   string `yaml:"Name"`
		Offset int    `yaml:"Offset"`
	} `yaml:"TimeZone"`
}

type log struct {
	Path     string `yaml:"Path"`
	FileName string `yaml:"FileName"`
	Level    uint32 `yaml:"Level"`
}

type session struct {
	Provider string `yaml:"Provider"`
	Path     string `yaml:"Path"`
	Domain   string `yaml:"Domain"`
	Secret   string `yaml:"Secret"`
	MaxAge   int    `yaml:"MaxAge"`
	HttpOnly bool   `yaml:"HttpOnly"`
}

type redis struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Password string `yaml:"Password"`
}

type lang struct {
	Default string   `yaml:"Default"`
	Lang    []string `yaml:"Lang"`
}
