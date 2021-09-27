package config

import (
	"github.com/spf13/viper"
	"os"
)

func Init(cfgName string) {
	path,err := os.Getwd()
	if err != nil {
		panic("get path,err:"+err.Error())
	}
	cfgPath := path+"/app/config/"+cfgName
	viper.SetConfigFile(cfgPath)
	if err := viper.ReadInConfig();err != nil{
		panic("read config err,err:"+err.Error())
	}
}

func Get(key string)interface{}  {
	return viper.Get(key)
}

func GetInt(key string)int  {
	return viper.GetInt(key)
}

func GetFloat(key string)float64  {
	return viper.GetFloat64(key)
}

func GetString(key string)string  {
	return viper.GetString(key)
}

func GetBool(key string)bool  {
	return viper.GetBool(key)
}