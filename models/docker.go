package models

type DockerCompose struct {
	Services struct {
		App struct {
			Image        string   `yaml:"image" json:"image"`
			Network_mode string   `yaml:"network_mode,omitempty" json:"network_mode,omitempty"`
			Restart      string   `yaml:"restart,omitempty" json:"restart,omitempty"`
			Ports        []string `yaml:"ports,omitempty" json:"ports,omitempty" `
			Environment  struct {
				API_TOKEN             string `yaml:"API_TOKEN"`
				APP_KEY               string `yaml:"APP_KEY" json:"APP_KEY,omitempty"`
				APP_URL               string `yaml:"APP_URL"`
				AWS_ACCESS_KEY_ID     string `yaml:"AWS_ACCESS_KEY_ID"`
				AWS_BUCKET            string `yaml:"AWS_BUCKET"`
				AWS_DEFAULT_REGION    string `yaml:"AWS_DEFAULT_REGION"`
				AWS_SECRET_ACCESS_KEY string `yaml:"AWS_SECRET_ACCESS_KEY" json:"AWS_SECRET_ACCESS_KEY,omitempty"`
				DB_DATABASE           string `yaml:"DB_DATABASE"`
				DB_HOST               string `yaml:"DB_HOST"`
				DB_PASSWORD           string `yaml:"DB_PASSWORD" json:"DB_PASSWORD,omitempty"`
				DB_USERNAME           string `yaml:"DB_USERNAME"`
				FILESYSTEM_DISK       string `yaml:"FILESYSTEM_DISK"`
				MAIL_ENCRYPTION       string `yaml:"MAIL_ENCRYPTION"`
				MAIL_FROM_ADDRESS     string `yaml:"MAIL_FROM_ADDRESS"`
				MAIL_FROM_NAME        string `yaml:"MAIL_FROM_NAME"`
				MAIL_HOST             string `yaml:"MAIL_HOST"`
				MAIL_PASSWORD         string `yaml:"MAIL_PASSWORD" json:"MAIL_PASSWORD,omitempty"`
				MAIL_PORT             string `yaml:"MAIL_PORT"`
				MAIL_USERNAME         string `yaml:"MAIL_USERNAME"`
				PHP_POOL_NAME         string `yaml:"PHP_POOL_NAME"`
			} `yaml:"environment,omitempty" json:"environment,omitempty"`
		} `yaml:"app" json:"app"`
	} `yaml:"services" json:"services"`
}
