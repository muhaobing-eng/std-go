package database

import "fmt"

type Option struct {
	IP       string
	Port     int
	User     string
	Password string
	DB       string
	Other    string
}

func (opt Option) GetURL() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?%s", opt.User, opt.Password, opt.IP, opt.Port, opt.DB, opt.Other)
}

func GetDefaultOption() *Option {
	return &Option{
		IP:       "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "",
		Other:    "charset=utf8&parseTime=True&loc=Local",
	}
}
