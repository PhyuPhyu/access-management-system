package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name     string
	Password string
	RoleId   uint
}

// Get admin by name
func GetAdmin(Admin *Admin, name string) (err error) {
	err = Db.Where("name = ?", name).First(&Admin).Error
	if err != nil {
		return err
	}
	return nil
}
