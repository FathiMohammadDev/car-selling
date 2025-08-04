package migrations

import (
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/constants"
	"github.com/FathiMohammadDev/car-selling/data/db"
	"github.com/FathiMohammadDev/car-selling/data/models"
	"github.com/FathiMohammadDev/car-selling/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.GetConfig())

func Up1() {
	database := db.GetDb()

	CreateTables(database)
	createDefaultInformation(database)
	createCountry(database)
}

func CreateTables(database *gorm.DB) {
	tables := []interface{}{}

	file := models.File{}
	country := models.Country{}
	city := models.City{}
	user := models.User{}
	role := models.Role{}
	userRole := models.UserRole{}

	tables = addNewTable(database, file, tables)
	tables = addNewTable(database, country, tables)
	tables = addNewTable(database, city, tables)
	tables = addNewTable(database, user, tables)
	tables = addNewTable(database, role, tables)
	tables = addNewTable(database, userRole, tables)

	database.Migrator().CreateTable(tables...)
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)
}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func createDefaultInformation(database *gorm.DB) {
	adminRole := models.Role{Name: constants.AdminRoleName}
	createRoleIfNotExist(database, &adminRole)
	defaultRole := models.Role{Name: constants.DefaultRoleName}
	createRoleIfNotExist(database, &defaultRole)
	u := models.User{Username: constants.DefaultUserName, FirstName: "Test", LastName: "Test",
		MobileNumber: "09111112222", Email: "admin@admin.com"}
	pass := "12345678"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)

	createAdminUserIfNotExists(database, &u, adminRole.Id)
}

func createRoleIfNotExist(database *gorm.DB, r *models.Role) {
	exist := 0
	database.Model(&models.Role{}).
		Select("1").
		Where("name= ?", r.Name).
		First(&exist)

	if exist == 0 {
		database.Create(&r)
	}
}

func createAdminUserIfNotExists(database *gorm.DB, u *models.User, roleId int) {
	exist := 0
	database.Model(&models.User{}).
		Select("1").
		Where("name= ?", u.Username).
		First(&exist)

	if exist == 0 {
		database.Create(u)
		ur := models.UserRole{UserId: u.Id, RoleId: roleId}
		database.Create(&ur)
	}
}

func createCountry(database *gorm.DB) {
	count := 0
	database.
		Model(&models.Country{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {
		database.Create(&models.Country{Name: "Iran", Cities: []models.City{
			{Name: "Tehran"},
			{Name: "Isfahan"},
			{Name: "Shiraz"},
			{Name: "Chalus"},
			{Name: "Ahwaz"},
		}})
		database.Create(&models.Country{Name: "USA", Cities: []models.City{
			{Name: "New York"},
			{Name: "Washington"},
		}})
		database.Create(&models.Country{Name: "Germany", Cities: []models.City{
			{Name: "Berlin"},
			{Name: "Munich"},
		}})
		database.Create(&models.Country{Name: "China", Cities: []models.City{
			{Name: "Beijing"},
			{Name: "Shanghai"},
		}})
		database.Create(&models.Country{Name: "Italy", Cities: []models.City{
			{Name: "Roma"},
			{Name: "Turin"},
		}})
		database.Create(&models.Country{Name: "France", Cities: []models.City{
			{Name: "Paris"},
			{Name: "Lyon"},
		}})
		database.Create(&models.Country{Name: "Japan", Cities: []models.City{
			{Name: "Tokyo"},
			{Name: "Kyoto"},
		}})
		database.Create(&models.Country{Name: "South Korea", Cities: []models.City{
			{Name: "Seoul"},
			{Name: "Ulsan"},
		}})
	}
}
