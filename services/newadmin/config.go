package newAdmin

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

var (
	// Config is the configuration for the newAdmin service
	Config config
)

func InitConfig() error {
	c := config{}
	c.Url = viper.GetString("newadmin.url")
	c.Email = viper.GetString("newadmin.email")
	c.Password = viper.GetString("newadmin.password")
	err := getToken(&c)
	if err != nil {
		return err
	}
	Config = c
	fmt.Println(Config)
	return nil
}

func getToken(c *config) error {
	// use email and password to get token
	body := fmt.Sprintf(`{"email":"%s","password":"%s"}`, c.Email, c.Password)

	req, err := http.NewRequest("POST", c.Url+"/api/admin/login", bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(req.Header.Get("Authorization"))
	c.Token = req.Header.Get("Authorization")
	return nil
}

type config struct {
	Url      string `json:"url"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type Client struct {
	c      *http.Client
	apiKey string
}

func setDefaultHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", `application/json`)
	req.Header.Set("Authorization", Config.Token)
}
