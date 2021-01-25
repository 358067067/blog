package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db     string
	DbHost string
	DbPort string
	DbUser string
	DbPwd  string
	DbName string

	AK         string
	SK         string
	Bucket     string
	FileServer string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查配置文件路径：", err)
	}
	loadServer(file)
	loadData(file)
	LoadSevenCow(file)
}

func loadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8080")
	JwtKey = file.Section("server").Key("JwtKey").MustString("734jks824")

}

func loadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPwd = file.Section("database").Key("DbPwd").MustString("root")
	DbName = file.Section("database").Key("DbName").MustString("test")
}

func LoadSevenCow(file *ini.File) {
	AK = file.Section("sevencow").Key("AK").String()
	SK = file.Section("sevencow").Key("SK").String()
	Bucket = file.Section("sevencow").Key("Bucket").String()
	FileServer = file.Section("sevencow").Key("FileServer").String()
}
