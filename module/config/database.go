package config

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
