package config

type config struct {
	RunMode  string      `yaml:"RunMode"`
	AppName  string      `yaml:"AppName"`
	Server   *server     `yaml:"Server"`
	Fix      *fix        `yaml:"Fix"`
	Log      *log        `yaml:"Log"`
	DataBase []*DataBase `yaml:"DataBase"`
	Session  *session    `yaml:"Session"`
	Redis    *redis      `yaml:"Redis"`
	Lang     *lang       `yaml:"Lang"`
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

type DataBase struct {
	DriverName      string `yaml:"DriverName"`
	DataBase        string `yaml:"DataBase"`
	Host            string `yaml:"Host"`
	Port            string `yaml:"Port"`
	Username        string `yaml:"Username"`
	Password        string `yaml:"Password"`
	Charset         string `yaml:"Charset"`
	LogLevel        int    `yaml:"LogLevel"`
	ConnMaxLifetime int    `yaml:"MaxLifetime"`
	MaxIdleConns    int    `yaml:"MaxIdleConns"`
	MaxOpenConns    int    `yaml:"MaxOpenConns"`
	ShowSql         bool   `yaml:"ShowSql"`
	ShowExecTime    bool   `yaml:"ShowExecTime"`
	Connect         bool   `yaml:"Connect"`
}

type session struct {
	Provider string `yaml:"Provider"`
	Path     string `yaml:"Path"`
	Domain   string `yaml:"Domain"`
	Secure   string `yaml:"Secure"`
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
