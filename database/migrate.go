package main

import (
	"flag"
	"fmt"
	migrate "github.com/govel-golang-migration/govel-golang-migration"
	"jd_workout_golang/lib/database"
	"jd_workout_golang/lib/file"
	"os"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()

	method := flag.Arg(0)

	file.LoadConfigAndEnv()

	database.InitDatabase()

	dispatch(method)
}

var migrationPath = "database"

func dispatch(method string) {
	fmt.Printf("dispatch %s method \n", method)
	
	rebuild, _ := strconv.ParseBool(os.Getenv("REBUILD"))

	switch method {
	case "status":
		fmt.Println("status ...")

		migrate.Status(os.Getenv("DB_HOST"), migrationPath)
	case "install":
		fmt.Println("install ...")


		migrate.Install(os.Getenv("DB_HOST"))
	case "make":
		fmt.Println("make")

		fileName := strings.Trim(flag.Arg(1), "")

		if fileName == "" {
			fmt.Println("file name is required")

			return
		}

		migrate.Make(fileName, migrationPath)
	case "migrate":
		fmt.Println("migrate")

		migrate.Migrate(os.Getenv("DB_HOST"), migrationPath, rebuild)
	case "rollback":
		fmt.Println("rollback")

		stage, err := strconv.Atoi(strings.Trim(flag.Arg(1), ""))

		if err != nil {
			fmt.Println("stage is required and must be numeric")

			return
		}

		migrate.Rollback(stage, os.Getenv("DB_HOST"), migrationPath, rebuild)
	default:
		fmt.Println("method not support")
	}
}
