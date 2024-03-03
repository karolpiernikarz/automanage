package helpers

import (
	"bufio"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func AppendCaddyFile(domainname string, port string, restaurantid string) error {
	// caddyfile
	caddyfile := []string{
		"",
		"# Restaurantid: " + restaurantid,
		domainname + ", www." + domainname + " {",
		"	encode gzip zstd",
		"	import httpsimports",
		"	import globalimports",
		"	@www host www." + domainname,
		"	handle @www {",
		"		redir https://" + domainname + "{uri} permanent",
		"	}",
		"	import log " + domainname,
		"	reverse_proxy 127.0.0.1:" + port,
		"}",
		"http://" + domainname + ", http://www." + domainname + " {",
		"	encode gzip zstd",
		"	import orderbox",
		"	import httpimports",
		"	import globalimports",
		"	@www host www." + domainname,
		"	handle @www {",
		"		redir https://" + domainname + "{uri} permanent",
		"	}",
		"	reverse_proxy 127.0.0.1:" + port,
		"}",
		"",
	}
	// convert slice to string
	str := strings.Join(caddyfile, "\n")
	// convert string to byte
	sliceByte := []byte(str)
	// open file
	file, _ := os.OpenFile(viper.GetString("caddy.file"),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0660)
	defer file.Close()
	// write to file
	if _, err := file.Write(sliceByte); err != nil {
		return err
	}
	return nil
}

func IsDomainExist(domainname string) bool {
	// open file
	f, err := os.Open(viper.GetString("caddy.file"))
	if err != nil {
		return false
	}
	// read file
	scanner := bufio.NewScanner(f)
	// check if domainname exist
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "https://"+domainname) {
			return true
		}
	}
	defer f.Close()
	return false
}

// ChangeDomain changes domainname in caddyfile
func ChangeDomain(before string, after string) (err error) {
	caddyFile := viper.GetString("caddy.file")
	// read file
	file, err := os.ReadFile(caddyFile)
	if err != nil {
		return err
	}
	// change before to after
	newFile := strings.Replace(string(file), before, after, -1)
	// write to file
	if err := os.WriteFile(caddyFile, []byte(newFile), 0); err != nil {
		return err
	}
	return nil
}
