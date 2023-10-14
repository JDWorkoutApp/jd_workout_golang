package main

import (
	"gorm.io/gorm"
	"jd_workout_golang/lib/database"
)

type Equip struct {
	gorm.Model
	UserId  int    `gorm:"type:int;uniqueIndex:idx_user_equip"`
	Name    string `gorm:"type:varchar(255);uniqueIndex:idx_user_equip"`
	Weights string `gorm:"type:json"`
	Note    string `gorm:"type:varchar(255)"`
}

func UpCreateEquipTable() error {
	err := database.Connection.Migrator().CreateTable(&Equip{})
	if err != nil {
		return err
	}
	
	return nil
}

func DownCreateEquipTable() error {
	err := database.Connection.Migrator().DropTable(&Equip{})
	if err != nil {
		return err
	}

	return nil
}
