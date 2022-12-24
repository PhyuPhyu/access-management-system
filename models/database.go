package models

import (
	"access-management-system/config"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

// Connect database
func ConnectDB(config *config.Config) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true&loc=Local", config.DBUsername, config.DBPassword, config.DBHost, config.DBName)

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}

	fmt.Println("Connected Successfully to the Database")
}

// Migrate tables
func MigrateTables() {
	Db.Migrator().HasTable(&Role{})
	Db.Migrator().DropTable(&Role{})

	Db.Migrator().HasTable(&Admin{})
	Db.Migrator().DropTable(&Admin{})

	Db.AutoMigrate(&Role{}, &Admin{})

	MigrateUser()
}

// Migrate user table
func MigrateUser() {
	Db.Migrator().HasTable(&User{})
	Db.Migrator().DropTable(&User{})

	Db.AutoMigrate(&User{})
}

// Seed data
func SeedData() {
	roles := []Role{
		{
			Name: "admin",
		},
		{
			Name: "staff",
		},
		{
			Name: "dev",
		},
	}

	// Create roles
	for _, role := range roles {
		Db.Create(&role)
	}

	AdminMocker()

	UserMocker(3)
}

func AdminMocker() []Admin {
	var admins []Admin
	for i := 1; i <= 3; i++ {
		admin := Admin{
			Name:     fmt.Sprintf("admin%v", i),
			Password: "password",
			RoleId:   uint(i),
		}
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("Password hash error: ", err)
		}

		admin.Password = string(passwordHash)

		Db.Create(&admin)
		admins = append(admins, admin)
	}
	return admins
}

func UserMocker(n int) []User {
	var users []User
	for i := 1; i <= n; i++ {
		user := User{
			Name:     fmt.Sprintf("user%v", i),
			Email:    fmt.Sprintf("user%v@gmail.com", i),
			Password: "password",
		}
		GenerateAndSetPasswordHash(&user)
		Db.Create(&user)
		users = append(users, user)
	}
	return users
}
