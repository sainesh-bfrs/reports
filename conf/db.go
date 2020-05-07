package conf

import "github.com/spf13/viper"

// Get ...
func Get(key string) string {
	viper.ReadInConfig()
	return viper.Get(key).(string)
}
