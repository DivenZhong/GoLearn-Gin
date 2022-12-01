/**
 * @Description mongoDb配置信息
 **/
package config

// MySQL信息
type mongoDb struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Enable   bool   `yaml:"enable"`
}
