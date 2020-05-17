package utils

import (
	"github.com/spf13/viper"
)

// Config ...
var Config *Configuration

// Configuration ...
type Configuration struct {
	Service       ServiceConfiguration
	DB            DBConfiguration
	MongoDBConfig MongoDBConfiguration
	DebugMode     bool
}

// ServiceConfiguration ...
type ServiceConfiguration struct {
	Port string
}

// DBConfiguration ...
type DBConfiguration struct {
	Host     string
	DBType   string
	DBName   string
	User     string
	Password string
	Port     string
	LogMode  bool
}

// MongoDBConfiguration ...
type MongoDBConfiguration struct {
	Host                  string
	DBName                string
	EntityCollection      string
	CompanyCollection     string
	InsInvestorCollection string
	LPCollection          string
	PersonCollection      string
	FundCollection        string
	DealCollection        string
	VerticalCollection    string
}

// NewConfiguration ...
func NewConfiguration(configName string, configPaths []string) error {
	v := viper.New()
	v.SetConfigName(configName)
	v.AutomaticEnv()
	for _, configPath := range configPaths {
		v.AddConfigPath(configPath)
	}

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	err := v.Unmarshal(&Config)
	if err != nil {
		return err
	}

	return nil
}
