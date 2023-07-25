package database

import (
	"section11.ORM-CodeStructure/config"
	// "project/config"
	//"project/models"
)

func GetUsers() (interface{}, error) {
	var users []models.users
	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}
