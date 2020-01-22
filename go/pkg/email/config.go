package email

type SMTPConfig struct {
	Host     string  `required:"true" yaml:"host"`
	Port     int     `mapstructure:"port" yaml:"port"`
	Username string  `mapstructure:"username" yaml:"username"`
	Password string  `mapstructure:"password" yaml:"password"`
	From     Address `mapstructure:"from" yaml:"from"`
}
