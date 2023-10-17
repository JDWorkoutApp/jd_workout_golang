package main

import (
	migrate "github.com/govel-golang-migration/govel-golang-migration"
	"github.com/joho/godotenv"
	"jd_workout_golang/lib/database"
	"jd_workout_golang/lib/file"
	"os"
)

func main() {
	path := file.AccessFromRootDir(".env")
	if err := godotenv.Load(path); err != nil {
		panic(err)
	}

	dsn := os.Getenv("DB_TEST_DSN")
	database.InitDatabase(dsn)
	database.Connection.Exec("DROP DATABASE IF EXISTS jd_workout_test;")
	database.Connection.Exec("CREATE DATABASE jd_workout_test;")
	database.Connection.Exec("USE jd_workout_test;")
	migrate.Install(dsn)

	migrationPath := file.AccessFromRootDir("database")
	migrate.Migrate(dsn, migrationPath, false)
}
