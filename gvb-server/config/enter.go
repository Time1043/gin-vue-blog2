package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	Upload   Upload   `yaml:"upload"`
	SiteInfo SiteInfo `yaml:"site_info"`

	QQ    QQ    `yaml:"qq"`
	QiNiu QiNiu `yaml:"qi_niu"`
	Jwt   Jwt   `yaml:"jwt"`
	Email Email `yaml:"email"`
}
