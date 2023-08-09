## Android Play Store

[Link](https://play.google.com/store/apps/details?id=com.govel.workout&hl=zh-TW)

## iOS App Store

[//]: # (emoji angry)

Can't afford the fee. :hankey: :hankey: :hankey:

## Online API document

### Dev (private)

### Prod

http://www.govel-workout.com/swagger/index.html

## Sentry

http://www.govel-workout.com:6002/

## Run docker

* `docker-compose build`
* `docker-compose up -d`

## Hot reload

use command `air`

Have issue with windows + docker. 

[Issue and solution](https://github.com/cosmtrek/air/issues/190)

## Windows hot reload (Docker)

Solution that replace air.

[Install](https://github.com/benblamey/when_changed/releases/) this, 
and put in `bat\when_changed.exe`

1. run command in container: `./bash/main.bash`
2. run command in windows (project root folder): `bat\when_changed.exe  ./**.go  bat\reload.bat`
    > Don't use git bash if it not work as you think.
3. `main.go` should reload when ever you change any go file.

**Dev routine**

* `docker-compose up -d`
* In app container: `./bash/main.bash`
  * This command will run `go run cmd/main.go` and restart if session get killed.
* In windows cmd run `bat\when_changed.exe  ./**.go  bat\reload.bat`
  * `when_changed.exe` download from gitHub mentioned above.
  * `./**.go` listen all go file change.
  * run `bat\reload.bat` when ever file change.
* `bat\reload.bat` will execute `bash\reload.bash` in container.
* `bash\reload.bash` will kill `go run cmd/main.go` and run `go run cmd/main.go` again through `./bash/main.bash` we run before.

### test api Hot reload

run `air -c .air_test.toml`

## Run test api

run `go run cmd/test-api.go` in the container.

## Migrate

### create new migration

`migrate create -ext sql -dir database/migrations create_users_table`

### run migrate
- up
`migrate -database "mysql://${username}:${password}@tcp${MYSQL_URL}/${DB_NAME}" -path db/migrations up`


- down
`migrate -database "mysql://${username}:${password}@tcp${MYSQL_URL}/${DB_NAME}" -path db/migrations down ${BATCH}`

## Api document

### Generate document

`swag init -g cmd/main.go`

### Check in browser

http://localhost:8080/swagger/index.html