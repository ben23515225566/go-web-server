package database

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type User struct {
	Username string
	Password string
}

type Article struct {
	Id        int
	Title     string
	Content   string
	Created_at time.Time
	Updated_at time.Time
}

var (
	db *gorm.DB
)

// Connect to database
func init() {
	vp := viper.New()
	vp.AddConfigPath("./database/config")
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	err := vp.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
	cfg := &config{}
	cfg.Host = vp.GetString("app.Host")
	cfg.Port = vp.GetInt("app.Port")
	cfg.User = vp.GetString("app.User")
	cfg.Password = vp.GetString("app.Password")
	cfg.Database = vp.GetString("app.Database")

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Database,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Database connection succeed.")
}

func InsertToUsers(username string, password string) error {
	new_user := &User{Username: username, Password: password}
	result := db.Create(new_user)
	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func InsertToArticles(title string, content string) (error){
	now := time.Now()
	new_article := &Article{Title: title, Content: content, Created_at: now, Updated_at: now}
	result := db.Create(new_article)
	if result.Error != nil{
		return result.Error
	}else{
		return nil
	}
}

func QueryAllArticles() ([]Article, error) {
	var articles []Article
	result := db.Find(&articles)
	return articles, result.Error
}

func QueryOneArticle(id int)(*Article, error){
	article := Article{}
	result := db.Where("Id = ?", id).Find(&article)
	if result.Error != nil{
		return nil, result.Error
	}else{
		return &article, nil
	}
}

func DeleteArticle(id int)(error){
	deletedArticle := Article{Id: id}
	result := db.Delete(&deletedArticle)
	if result.Error != nil{
		return result.Error
	}else{
		return nil
	}
}

func CloseDatabase() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Panic(err)
	}
	log.Println("Database close.")
	sqlDB.Close()
}
