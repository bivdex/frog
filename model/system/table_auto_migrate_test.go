package system

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestMigrateDbG(t *testing.T) {
	host := "198.2.235.111"
	port := "3306"
	dbNa := "ericg"
	dbUn := "ericg"
	dbPw := "ericg"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUn, dbPw, host, port, dbNa)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&TMainAddr{})
}

func TestMigrateDbBoost(t *testing.T) {
	host := "198.2.235.111"
	port := "3306"
	dbNa := "boost"
	dbUn := "boost"
	dbPw := "boost"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUn, dbPw, host, port, dbNa)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&TAddressActivity{})
}
