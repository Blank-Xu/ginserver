package config

type config struct {
	Server   *Server     `yaml:"Server"`
	Fix      *Fix        `yaml:"Fix"`
	Log      *Log        `yaml:"Log"`
	DataBase []*DataBase `yaml:"DataBase"`
	Session  *Session    `yaml:"Session"`
	Redis    *Redis      `yaml:"Redis"`
	Lang     *Lang       `yaml:"Lang"`
}

type Server struct {
	RunMode      string `yaml:"RunMode"`
	AppName      string `yaml:"AppName"`
	HttpPort     string `yaml:"HttpPort"`
	ReadTimeout  int    `yaml:"ReadTimeout"`
	WriteTimeout int    `yaml:"WriteTimeout"`
}

type Fix struct {
	TimeZone *struct {
		Name   string `yaml:"Name"`
		Offset int    `yaml:"Offset"`
	} `yaml:"TimeZone"`
}

type Log struct {
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

type Session struct {
	Name     string `yaml:"Name"`
	Expired  int    `yaml:"Expired"`
	Provider string `yaml:"Provider"`
	Sign     string `yaml:"Sign"`
}

type Redis struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Password string `yaml:"Password"`
}

type Lang struct {
	Default string   `yaml:"Default"`
	Lang    []string `yaml:"Lang"`
}
