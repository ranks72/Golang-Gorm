# Golang-Gorm
At this Github, i learn about make rest-api using gorm in golang

## Tools
- Golang
- Postgres
- VsCode
- Postman for testing

## First Step

open your terminal

run this code :
```sh
go mod init golang-gorm
```

or run this code:

```sh
go mod init (name folder)
```

## Second Step

```sh
 go get -u github.com/gin-gonic/gin
```

## Third Step
```sh
 go get -u gorm.io/gorm
```

## Fourth Step

Open the file db.go in folder database
search this code

```sh
//db.Debug().AutoMigrate(models.Order{}, models.Item{})
```
uncomment this code because this code make table in your database

And fill your database user and password

## Testing

For the test in my rest-api, you can use this postman

```sh
https://www.getpostman.com/collections/8fe8d958fc9f29b65906
```
