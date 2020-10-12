package main

import (
	"MyGo/config"
	"MyGo/entity"
	"MyGo/service"
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config.InitDIContainer()
	config.BindingDependencies()

	db := initSqlLiteDb()
	//initPostgresDb()
	//initPostgresDbExisting()
	//initVariables()

	migrateError := db.AutoMigrate(
		&entity.User{},
		&entity.Address{},
		&entity.Tag{},
		&entity.Blog{})

	if migrateError != nil {
		panic("Cannot create Users table")
	}

	truncateDB(db)

	createAddresses(db)
	createBlogs(db)
	findBlogs(db)
	invokeUserService(db)
}

func initSqlLiteDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		panic("Something went wrong in db connection")
	}

	return db
}

func createBlogs(db *gorm.DB) {
	blog1:= &entity.Blog{
		Author: entity.Author{ Name: "Hieu Dang" },
		Author2: entity.Author{ Name: "Hieu Duy" },
		Tag: []*entity.Tag{ {Locale: "EN"}, {Locale: "VN"} },
	}

	blog2:= &entity.Blog{
		Author: entity.Author{ Name: "Hieu Elsevier" },
		Author2: entity.Author{ Name: "Hieu Media" },
		Tag: []*entity.Tag{ {Locale: "FR"}, {Locale: "VN"} },
		LocaleTags: []*entity.Tag{ {Locale: "FR"}, {Locale: "VN"} },
	}

	db.Create(&blog1)
	db.Create(&blog2)
}

func createAddresses(db *gorm.DB) {
	db.Model(&entity.Address{}).Create([]map[string]interface{} {
		{"Street": "123 Cộng Hòa", "City": "Hồ Chí Minh", "ZipCode": "008428"},
		{"Street": "1600 John F Kennedy", "City": "Philly", "ZipCode": "19103 2398", "State": "Pennsylvania"},
	})
}

func findBlogs(db *gorm.DB)  {
	var blogResults []entity.Blog
	db.Preload("Tag").Find(&blogResults)
	fmt.Println(blogResults)
}

func truncateDB(db *gorm.DB) {
	db.Exec("DELETE FROM users")
	db.Exec("DELETE FROM addresses")
	db.Exec("DELETE FROM blog_tags")
	db.Exec("DELETE FROM locale_blog_tags")
	db.Exec("DELETE FROM blogs")
	db.Exec("DELETE FROM tags")
}

func myPostgresDbInit() {
	dsn := "user=postgres password=postgres dbname=reccontent port=5433 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Something went wrong in db connection")
	}

	fmt.Print(db.DB())
}

func initPostgresDbExisting() {
	dataSourceName := "user=postgres password=postgres dbname=reccontent port=5433 sslmode=disable TimeZone=Asia/Jakarta"
	sqlDB, err1 := sql.Open("postgres", dataSourceName)
	//// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	//sqlDB.SetMaxIdleConns(10)
	//// SetMaxOpenConns sets the maximum number of open connections to the database.
	//sqlDB.SetMaxOpenConns(100)
	//// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	//sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Print(err1)

	gormDB, err2 := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err2 != nil {
		panic("Something went wrong in db connection")
	}

	fmt.Print(gormDB.DB())
}

func initPostgresDbExisting2() {
	dataSourceName := "postgres:postgres@tcp(127.0.0.1:5433)/ocs?charset=utf8&parseTime=True"
	sqlDB, err1 := sql.Open("postgres", dataSourceName)
	//// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	//sqlDB.SetMaxIdleConns(10)
	//// SetMaxOpenConns sets the maximum number of open connections to the database.
	//sqlDB.SetMaxOpenConns(100)
	//// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	//sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Print(err1)

	gormDB, err2 := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err2 != nil {
		panic("Something went wrong in db connection")
	}

	fmt.Print(gormDB.DB())
}

func initVariables()  {
	var s1 string
	var s2 *string
	s1 = "123 and "
	s2 = &s1
	fmt.Print(s1, *s2)
}

func invokeUserService(db *gorm.DB) {
	var userService service.UserService
	config.Container.Make(&userService)

	userService.Register(db)
	userService.GetUser(db)
}