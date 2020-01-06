package email

type SMTPConfig struct {
	Host         string  `required:"true" yaml:"host"`
	Port         int     `mapstructure:"post" yaml:"post"`
	Username     string  `mapstructure:"username" yaml:"username"`
	Password     string  `mapstructure:"password" yaml:"password"`
	From         Address `mapstructure:"from" yaml:"from"`
	TemplatePath string  `split_words:"true" required:"true" yaml:"templatePath"`
}
