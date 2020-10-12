package impl

import (
	"MyGo/entity"
	"MyGo/entity/dto"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type UserServiceImpl struct {}

func (u UserServiceImpl) GetUser(db *gorm.DB) {
	findUsers(db)
}

func (u UserServiceImpl) Register(db *gorm.DB) int {
	user1Hobbies := datatypes.JSON(`{"coding": "JAVA", "movie": "JAV"}`)
	user2Hobbies := datatypes.JSON(`{"sport": ["Soccer", "Gym"]}`)

	user1 := createUser("ABC", "abc@kms-technology.com", 99, user1Hobbies)
	user2 := createUser("DEF", "def@kms-technology.com", 100, nil)
	user3 := createUser("GHI", "ghi@kms-technology.com", 101, nil)
	persistUser(db, user1)
	persistUser(db, user2)
	persistUser(db, user3)

	updateUserAge(db, user1)
	updateUserHobbies(db, user2Hobbies)
	updateUserBySpecified(db, user3)

	deleteUsers(db, user1)
	return 1
}

func createUser(name string, email string, age uint8, hobbies datatypes.JSON) *entity.User {
	user := entity.User{Name: name, Email: &email, Age: age, Birthday: time.Now(), Hobbies: hobbies}

	return &user
}

func persistUser(db *gorm.DB, user *entity.User) *gorm.DB {
	return db.Create(&user)
}

func updateUserAge(db *gorm.DB, user *entity.User)  {
	user.Age = 200
	db.Save(user)
}

func updateUserHobbies(db *gorm.DB, hobbies datatypes.JSON) {
	db.Model(&entity.User{}).
		Where("email = ?", "def@kms-technology.com").
		Update("hobbies", hobbies)
}

func updateUserBySpecified(db *gorm.DB, user *entity.User) {
	db.Model(&user).
		Select("name", "Birthday").
		Omit("Email").
		Where("email = ?", user.Email).
		Updates(map[string]interface{} {"name": "GHI - Gnopus", "age": 18, "actived": false})
}

func findUsers(db *gorm.DB)  {
	var result []dto.Result

	db.Raw("SELECT id, name, age FROM users WHERE age > 100").Scan(&result)

	fmt.Print(result)
}

func deleteUsers(db *gorm.DB, user1 *entity.User) {
	user1.ID = 1
	db.Delete(&user1)
	db.Delete(&entity.User{}, "age = 100")
	db.Model(&entity.User{}).Where("email = ?", "ghi@kms-technology.com").Delete(&entity.User{})
}
