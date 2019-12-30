package email

type SMTPConfig struct {
	Host         string  `mapstructure:"host" yaml:"host"`
	Port         int     `mapstructure:"post" yaml:"post"`
	Username     string  `mapstructure:"username" yaml:"username"`
	Password     string  `mapstructure:"password" yaml:"password"`
	From         Address `mapstructure:"from" yaml:"from"`
	TemplatePath string  `mapstructure:"templatePath" yaml:"templatePath"`
}
