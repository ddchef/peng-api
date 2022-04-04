package config

type Captcha struct {
	Len         int   `mapstructure:"captcha_len" json:"captcha_len" yaml:"captcha_len"`
	CaptchaTtl  int64 `mapstructure:"captcha_ttl" json:"captcha_ttl" yaml:"captcha_ttl"`
	Width       int   `mapstructure:"captcha_width" json:"captcha_width" yaml:"captcha_width"`
	Height      int   `mapstructure:"captcha_height" json:"captcha_height" yaml:"captcha_height"`
	GracePeriod int   `mapstructure:"captcha_grace_period" json:"captcha_grace_period" yaml:"captcha_grace_period"`
}
