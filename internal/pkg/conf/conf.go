package conf

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var conf = &Config{}

func Conf() *Config {
	return conf
}

func Init(config string) error {
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	if err := v.Unmarshal(conf); err != nil {
		return err
	}

	return nil

}

type Config struct {
	DBZet Mysql `mapstructure:"db_zet" json:"db_zet" yaml:"db_zet"`
}

type Mysql struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Port         string `mapstructure:"port" json:"port" yaml:"port"`
	DBName       string `mapstructure:"db_name" json:"db_name" yaml:"db_name"`
	User         string `mapstructure:"user" json:"user" yaml:"user"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`
	MaxLifeTime  int    `mapstructure:"max_life_time" json:"max_life_time" yaml:"max_life_time"`
	MaxIdleTime  int    `mapstructure:"max_idle_time" json:"max_idle_time" yaml:"max_idle_time"`
}

func (m Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.DBName + "?" + m.Config
}

func (m Mysql) GormConfig() *gorm.Config {
	return &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}
