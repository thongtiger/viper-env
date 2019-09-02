package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		HOST string
		PORT int
	}
	Mongodb struct {
		Host     string
		Database string
		Username string
		Password string
	}
}

func main() {
	cf := &Config{}
	cf.GetEnv()
	fmt.Println(cf)
}

// GetEnv read config from environment
func (c *Config) GetEnv() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // path to look for the config file in
	viper.SetConfigType("json")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

	// app
	if val, ok := os.LookupEnv("BASE_URL"); ok {
		c.App.HOST = strings.TrimSpace(val)
	} else {
		c.App.HOST = strings.TrimSpace(viper.GetString("app.host"))
	}

	// port
	if val, ok := os.LookupEnv("PORT"); ok {
		if Num, err := strconv.Atoi(val); err == nil {
			c.App.PORT = Num
		}
	} else {
		c.App.PORT = viper.GetInt("app.port")
	}

	// mongo host
	if val, ok := os.LookupEnv("MONGO_HOST"); ok {
		c.Mongodb.Host = strings.TrimSpace(val)
	} else {
		c.Mongodb.Host = viper.GetString("app.datastore.mongodb.host")
	}

	// mongo dbname
	if val, ok := os.LookupEnv("MONGO_DB"); ok {
		c.Mongodb.Database = strings.TrimSpace(val)
	} else {
		c.Mongodb.Database = viper.GetString("app.datastore.mongodb.dbname")

	}
	// mongo username
	if val, ok := os.LookupEnv("MONGO_USER"); ok {
		c.Mongodb.Username = strings.TrimSpace(val)
	} else {
		c.Mongodb.Username = viper.GetString("app.datastore.mongodb.username")

	}
	// mongo password
	if val, ok := os.LookupEnv("MONGO_PASSWORD"); ok {
		c.Mongodb.Password = strings.TrimSpace(val)
	} else {

		c.Mongodb.Password = viper.GetString("app.datastore.mongodb.password")
	}
}
