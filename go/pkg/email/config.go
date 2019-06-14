package email

type SMTPConfig struct {
	Host         string  `yaml:"host"`
	Port         int     `yaml:"post"`
	Username     string  `yaml:"username"`
	Password     string  `yaml:"password"`
	From         Address `yaml:"From"`
	TemplatePath string  `yaml:"templatePath"`
}
