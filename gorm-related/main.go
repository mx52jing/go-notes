package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
}

type Teacher struct {
	ID     int    `gorm:"primaryKey;type:int unsigned"`
	Name   string `gorm:"type:varchar(32)"`
	Age    uint8
	Gender uint8
}

func (t *Teacher) TableName() string {
	return "teacher"
}

var DB *gorm.DB

// 获取数据库配置
func loadDbConfig(configFile string) (*DatabaseConfig, error) {
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	var config DatabaseConfig
	err = viper.UnmarshalKey("mysql.db1", &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return &config, nil
}

// 生产连接数据库链接
func genDbUrl(config *DatabaseConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Username, config.Password, config.Host, config.Port, config.DBName)
}

// 连接数据库
func connectToDb(dbUrl string) (*gorm.DB, error) {
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold:              time.Second,   // Slow SQL threshold
	// 		LogLevel:                   logger.Silent, // Log level
	// 		IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
	// 		ParameterizedQueries:      true,           // Don't include params in the SQL log
	// 		Colorful:                  false,          // Disable color
	// 	},
	// )
	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		// Logger: newLogger,
		// NamingStrategy: schema.NamingStrategy{
		// TablePrefix: "lyb_", // 生成的表名前缀，User结构体生成的表名为lyb_users
		// SingularTable: true, // 使用单数表名，设置为true后，User结构体生成的表名为user，而不是users
		// NoLowerCase: true, // 跳过转换蛇形命名
		// NameReplacer: strings.NewReplacer("User", "Cid"), // 名称替换，执行gorm名称替换策略之前，将User替换为Cid
		// },
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL database: %w", err)
	}
	return db, nil
}

type User struct {
	ID       uint
	UserName string //`gorm:"size:12;not null;comment:姓名"`
	Age      int    //`gorm:"size:3;comment:年龄"`
	Email    string // `gorm:"size:32;comment:邮箱"`
}

// TableName 会将 User 的表名重写为 `customUser`
// func (User) TableName() string {
// 	return "customUser"
// }

// func genUserTable() {
// 	DB.AutoMigrate(&User{})
// }

func GetDB() *gorm.DB {
	return DB
}

func init() {
	// 获取数据库配置
	dbConfig, err := loadDbConfig("config/db.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	// 根据数据库配置获取url
	dbUrl := genDbUrl(dbConfig)
	// 根据生产的url连接数据库
	db, connectErr := connectToDb(dbUrl)
	if connectErr != nil {
		log.Fatalf("Failed to connect to MySQL database: %v", connectErr)
	}
	fmt.Println("Successfully connected to MySQL database!")
	//err = db.AutoMigrate(&Teacher{})
	//if err != nil {
	//	return
	//}
	DB = db
}
