package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 指针, 传递的是值类型

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure: "name"`
	Mode         string `mapstructure: "mode"`
	Version      string `mapstructure: "version"`
	Port         int    `mapstructure: "port"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure: "level"`
	Filename   string `mapstructure: "filename"`
	MaxSize    int    `mapstructure: "max_size"`
	MaxAge     int    `mapstructure: "max_age"`
	MaxBackups int    `mapstructure: "max_backups"`
}

type MysqlConfig struct {
	Host         string `mapstructure: "host"`
	User         string `mapstructure: "user"`
	Password     string `mapstructure: "password"`
	DbName       string `mapstructure: "dbname"`
	Port         int    `mapstructure: "port"`
	MaxOpenConns int    `mapstructure: "max_open_conns"`
	MaxIdleConns int    `mapstructure: "max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure: "host"`
	Password string `mapstructure: "password"`
	Port     int    `mapstructure: "port"`
	DB       int    `mapstructure: "db"`
	PoolSize int    `mapstructure: "conn_pool"`
}

func Init() (err error) {
	viper.SetConfigFile("config.yaml")
	// viper.SetConfigName("config") // 配置文件名称(无扩展名)
	// viper.SetConfigType("yaml")   // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath(".") // 查找配置文件所在的路径

	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		fmt.Printf("viper.ReadInitConfig() failed, err:%v\n", err)
		return
	}

	// 把读取到的配置信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了!")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return
}
