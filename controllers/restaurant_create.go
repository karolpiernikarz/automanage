package controllers

import (
	"encoding/base64"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	isd "github.com/jbenet/go-is-domain"
	"github.com/karolpiernikarz/automanage/helpers"
	"github.com/karolpiernikarz/automanage/models"
	"github.com/sethvargo/go-password/password"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func RestaurantCreate(c *gin.Context) {
	restaurantid := c.PostForm("id")
	restauranttoken := c.PostForm("token")
	restaurantdomain := c.PostForm("domain")
	restaurantname := c.PostForm("name")
	dbname := viper.GetString("db.prefix") + restaurantid

	// check if id query is set
	if restaurantid == "" {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusBadRequest, Error: []string{"id can't be empty"}})
		log.WithFields(log.Fields{"url": c.Request.URL, "client_ip": c.ClientIP()}).Warn("id can't be empty")
		return
	}

	//check if id start with dot
	if strings.Contains(restaurantid, ".") {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusBadRequest, Error: []string{"restaurant id contains ."}})
		log.WithFields(log.Fields{"restaurantid": restaurantid, "url": c.Request.URL, "client_ip": c.ClientIP()}).Warn("restaurant id contains .")
		return
	}

	// check if token query is set
	if restauranttoken == "" {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusBadRequest, Error: []string{"token can't be empty"}})
		log.WithFields(log.Fields{"restaurantid": restaurantid, "url": c.Request.URL, "client_ip": c.ClientIP()}).Warn("token can't be empty")
		return
	}

	// check if domain query is set
	if restaurantdomain == "" {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusBadRequest, Error: []string{"domain can't be empty"}})
		log.WithFields(log.Fields{"restaurantid": restaurantid, "url": c.Request.URL, "client_ip": c.ClientIP()}).Warn("domain can't be empty")
		return
	}

	// check if restaurantname query is set
	if restaurantname == "" {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusBadRequest, Error: []string{"restaurantname can't be empty"}})
		log.WithFields(log.Fields{"restaurantid": restaurantid, "url": c.Request.URL, "client_ip": c.ClientIP()}).Warn("restaurantname can't be empty")
		return
	}
	//check if domain variables is correct
	url, err := url.Parse(restaurantdomain)
	domain := ""
	if url.Scheme != "" {
		domain = url.Hostname()
	} else {
		domain = url.String()
	}
	hostParts := strings.Split(domain, ".")
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusInternalServerError, Error: []string{err.Error()}})
		log.WithFields(log.Fields{"domain": restaurantdomain, "restaurantid": restaurantid, "url": c.Request.URL, "client_ip": c.ClientIP()}).Error(err.Error())
		return
	}
	if !isd.IsDomain(domain) {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusBadRequest, Error: []string{"not a domain"}})
		log.WithFields(log.Fields{"domain": restaurantdomain, "restaurantid": restaurantid, "url": c.Request.URL, "client_ip": c.ClientIP()}).Warn("not a domain")
		return
	}

	// check if restaurant is already created
	if _, err := os.Stat(filepath.Join(viper.GetString("app.workdir"), restaurantid)); !os.IsNotExist(err) {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusBadRequest, Error: []string{"already exist"}})
		log.WithFields(log.Fields{"restaurantid": restaurantid, "url": c.Request.URL, "client_ip": c.ClientIP()}).Error("restaurant already exist")
		return
	}

	// check if domain already exist
	if helpers.IsDomainExist(domain) {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusBadRequest, Error: []string{"domain exist in Caddyfile"}})
		return
	}

	//check if database user already exist
	result, err := helpers.CheckDatabaseUser(dbname)
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusInternalServerError, Error: []string{err.Error()}})
		return
	}
	if result {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusInternalServerError, Error: []string{"username alredy exist"}})
		log.WithFields(log.Fields{"restaurantid": restaurantid, "url": c.Request.URL, "client_ip": c.ClientIP()}).Error("username alredy exist")
		return
	}

	//create a random passwords to use with database,appkey and bucketname
	dbpassword, err := password.Generate(32, 10, 0, false, true)
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusInternalServerError, Error: []string{err.Error()}})
		log.WithFields(log.Fields{"restaurantid": restaurantid, "url": c.Request.URL, "client_ip": c.ClientIP()}).Error(err.Error())
		return
	}
	laravelkey, err := password.Generate(32, 10, 0, false, true)
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusInternalServerError, Error: []string{err.Error()}})
		log.WithFields(log.Fields{"restaurantid": restaurantid, "url": c.Request.URL, "client_ip": c.ClientIP()}).Error(err.Error())
		return
	}
	uid := uuid.New()
	bucketname := uid.String()
	//find a free port to use
	free_port, err := helpers.GetFreePort()
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusInternalServerError, Error: []string{err.Error()}})
		return
	}

	//create aws bucket,user,policy
	err = helpers.CreateS3Bucket(bucketname, restaurantid, domain)
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusInternalServerError, Error: []string{err.Error()}})
		return
	}
	err = helpers.CreateS3BucketPolicy(bucketname)
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusInternalServerError, Error: []string{err.Error()}})
		return
	}
	err = helpers.CreateIAMUser(bucketname)
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusInternalServerError, Error: []string{err.Error()}})
		return
	}
	err = helpers.CreateIAMPolicy(bucketname, restaurantid, domain, string(hostParts[0])+"@"+viper.GetString("aws.smtpdomain"))
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusInternalServerError, Error: []string{err.Error()}})
		return
	}
	err = helpers.AttachIAMPolicy(bucketname)
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusInternalServerError, Error: []string{err.Error()}})
		return
	}
	accesskey, secretkey, err := helpers.CreateAccessKey(bucketname)
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusInternalServerError, Error: []string{err.Error()}})
		return
	}
	smtppassword, err := helpers.GenerateSmtpCredentials(secretkey, viper.GetString("aws.region"))
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusInternalServerError, Error: []string{err.Error()}})
		return
	}

	//create the database,user and give permisson
	err = helpers.CreateDatabase(dbname)
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusInternalServerError, Error: []string{err.Error()}})
		return
	}
	err = helpers.CreateDatabaseUser(dbname, dbpassword)
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusInternalServerError, Error: []string{err.Error()}})
		return
	}
	err = helpers.GivePermToUser(dbname, dbname)
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusInternalServerError, Error: []string{err.Error()}})
		return
	}

	// import the database
	_, err = helpers.ImportDatabase(dbname, dbpassword)
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusInternalServerError, Error: []string{err.Error()}})
		return
	}

	var dcfile models.DockerCompose
	dcfile.Services.App.Image = viper.GetString("app.image")
	dcfile.Services.App.Network_mode = "bridge"
	dcfile.Services.App.Restart = "unless-stopped"
	dcfile.Services.App.Ports = []string{
		"127.0.0.1:" + strconv.Itoa(free_port) + ":80",
	}
	dcfile.Services.App.Environment.PHP_POOL_NAME = restaurantid
	dcfile.Services.App.Environment.API_TOKEN = restauranttoken
	dcfile.Services.App.Environment.APP_KEY = "base64:" + base64.StdEncoding.EncodeToString([]byte(laravelkey))
	dcfile.Services.App.Environment.APP_URL = "https://" + domain
	dcfile.Services.App.Environment.DB_HOST = viper.GetString("db.host")
	dcfile.Services.App.Environment.DB_DATABASE = dbname
	dcfile.Services.App.Environment.DB_USERNAME = dbname
	dcfile.Services.App.Environment.DB_PASSWORD = dbpassword
	dcfile.Services.App.Environment.FILESYSTEM_DISK = "s3"
	dcfile.Services.App.Environment.MAIL_HOST = "email-smtp." + viper.GetString("aws.region") + ".amazonaws.com"
	dcfile.Services.App.Environment.MAIL_PORT = viper.GetString("aws.mailport")
	dcfile.Services.App.Environment.MAIL_USERNAME = "\"" + accesskey + "\""
	dcfile.Services.App.Environment.MAIL_PASSWORD = "\"" + smtppassword + "\""
	dcfile.Services.App.Environment.MAIL_ENCRYPTION = viper.GetString("aws.mailencryption")
	dcfile.Services.App.Environment.MAIL_FROM_ADDRESS = string(hostParts[0]) + "@" + viper.GetString("aws.smtpdomain")
	dcfile.Services.App.Environment.MAIL_FROM_NAME = "\"" + restaurantname + " (" + viper.GetString("app.company") + ")" + "\""
	dcfile.Services.App.Environment.AWS_ACCESS_KEY_ID = "\"" + accesskey + "\""
	dcfile.Services.App.Environment.AWS_SECRET_ACCESS_KEY = "\"" + secretkey + "\""
	dcfile.Services.App.Environment.AWS_DEFAULT_REGION = viper.GetString("aws.region")
	dcfile.Services.App.Environment.AWS_BUCKET = bucketname

	// create the restaurant
	err = os.Mkdir(filepath.Join(viper.GetString("app.workdir"), restaurantid), os.ModePerm)
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusInternalServerError, Error: []string{err.Error()}})
		log.WithFields(log.Fields{"restaurantid": restaurantid, "url": c.Request.URL, "client_ip": c.ClientIP()}).Error(err.Error())
		return
	}

	file, _ := yaml.Marshal(&dcfile)
	err = os.WriteFile(viper.GetString("app.workdir")+"/"+restaurantid+"/docker-compose.yaml", file, 0660)
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusInternalServerError, Error: []string{err.Error()}})
		log.WithFields(log.Fields{"restaurantid": restaurantid, "url": c.Request.URL, "client_ip": c.ClientIP()}).Error(err.Error())
		return
	}

	err = helpers.AppendCaddyFile(domain, strconv.Itoa(free_port), restaurantid)
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusInternalServerError, Error: []string{err.Error()}})
		return
	}

	command := "docker compose -f " + viper.GetString("app.workdir") + "/" + restaurantid + "/docker-compose.yaml up -d"
	cmd := exec.Command("sh", "-c", command)
	err = cmd.Run()
	if err != nil {
		helpers.SendResponse(c, helpers.Response{Status: http.StatusInternalServerError, Error: []string{err.Error()}})
		log.WithFields(log.Fields{"restaurantid": restaurantid, "url": c.Request.URL, "client_ip": c.ClientIP()}).Error(err.Error())
		return
	}

	helpers.SendResponse(c, helpers.Response{
		Status:  http.StatusOK,
		Message: []string{"Ok"},
		Error:   []string{},
	})
}