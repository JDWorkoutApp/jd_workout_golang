package main

import (
	"gorm.io/gorm"
	"jd_workout_golang/lib/database"
)

type Record struct {
	gorm.Model
	UserId  uint  `json:"userId"`
	EquipId uint `json:"equip_id"`
	Name string `json:"name"`
	Weight float32 `json:"weight" gorm:"default:null"`
	Reps uint `json:"reps" gorm:"default:null"`
	Note string `json:"note" gorm:"default:null"`
}

func UpCreateRecordTable() error {
	err := database.Connection.Migrator().CreateTable(&Record{})
	if err != nil {
		return err
	}

	return nil
}

func DownCreateRecordTable() error {
	err := database.Connection.Migrator().DropTable(&Record{})
	if err != nil {
		return err
	}

	return nil
}