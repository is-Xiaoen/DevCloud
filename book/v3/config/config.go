package config

import (
	"fmt"
	"sync"

	"122.51.31.227/go-course/go18/book/v3/models"
	"github.com/infraboard/mcube/v2/tools/pretty"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Default() *Config {
	return &Config{
		Application: &application{
			Host: "127.0.0.1",
			Port: 8080,
		},
		MySQL: &mySQL{
			Host:     "127.0.0.1",
			Port:     3306,
			DB:       "test",
			Username: "root",
			Password: "123456",
			Debug:    true,
		},
	}
}

// 这歌对象就是程序配置
// yaml, toml
type Config struct {
	Application *application `toml:"app" yaml:"app" json:"app"`
	MySQL       *mySQL       `toml:"mysql" yaml:"mysql" json:"mysql"`
}

func (c *Config) String() string {
	return pretty.ToJSON(c)
}

// 应用服务
type application struct {
	Host string `toml:"host" yaml:"host" json:"host" env:"HOST"`
	Port int    `toml:"port" yaml:"port" json:"port" env:"PORT"`
}

// db对象也是一个单列模式
type mySQL struct {
	Host     string `json:"host" yaml:"host" toml:"host" env:"DATASOURCE_HOST"`
	Port     int    `json:"port" yaml:"port" toml:"port" env:"DATASOURCE_PORT"`
	DB       string `json:"database" yaml:"database" toml:"database" env:"DATASOURCE_DB"`
	Username string `json:"username" yaml:"username" toml:"username" env:"DATASOURCE_USERNAME"`
	Password string `json:"password" yaml:"password" toml:"password" env:"DATASOURCE_PASSWORD"`
	Debug    bool   `json:"debug" yaml:"debug" toml:"debug" env:"DATASOURCE_DEBUG"`

	// gorm db对象, 只需要有1个,不运行重复生成
	db *gorm.DB
	// 互斥锁
	lock sync.Mutex
}

func (m *mySQL) GetDB() *gorm.DB {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.db == nil {
		// 初始化数据库
		// dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			m.Username,
			m.Password,
			m.Host,
			m.Port,
			m.DB,
		)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		db.AutoMigrate(&models.Book{}) // 自动迁移
		m.db = db
	}

	return m.db
}
