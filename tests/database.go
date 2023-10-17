package tests

import (
	"fmt"
	"github.com/joho/godotenv"
	"jd_workout_golang/lib/database"
	"jd_workout_golang/lib/file"
	"os"
	"os/exec"
)

func InitDatabase() {
	cmd := exec.Command("go", "run", "databaseRefresh.go")
	cmd.Dir = file.AccessFromRootDir("tests/cmd")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("%v: %s\n", err, out)
		panic(err)
	}

	path := file.AccessFromRootDir(".env")
	if err := godotenv.Load(path); err != nil {
		panic(err)
	}

	dsn := os.Getenv("DB_TEST_DSN")
	database.InitDatabase(dsn)
}
