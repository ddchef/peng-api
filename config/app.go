package config

type App struct {
	Env  string `mapstructure:"env" json:"env" yaml:"env"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
	Name string `mapstructure:"name" json:"name" yaml:"name"`
	Url  string `mapstructure:"url" json:"url" yaml:"url"`
}
