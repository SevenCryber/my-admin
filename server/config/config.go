package config

type Server struct {
	JWT       JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap       Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis     Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	RedisList []Redis `mapstructure:"redis-list" json:"redis-list" yaml:"redis-list"`
	Email     Email   `mapstructure:"email" json:"email" yaml:"email"`
	System    System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha   Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// gorm
	Mysql  Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	// oss
	Local Local `mapstructure:"local" json:"local" yaml:"local"`

	Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`

	DiskList []DiskList `mapstructure:"disk-list" json:"disk-list" yaml:"disk-list"`

	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
