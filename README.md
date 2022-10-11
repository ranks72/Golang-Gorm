# Golang-Gorm
Disini saya belajar membuat rest api menggunakan golang dengan database Postgres

##Tools
- Golang
- Postgres
- VsCode

##First Step

open your terminal

run this code :
```sh
go mod init golang-gorm
```

or run this code:

```sh
go mod init (name folder)
```

##Second Step

```sh
 go get -u github.com/gin-gonic/gin
```

##Third Step
```sh
 go get -u gorm.io/gorm
```

#Fourth Step

Open the file db.go in folder database
search this code

```sh
//db.Debug().AutoMigrate(models.Order{}, models.Item{})
```
uncomment this code because this code make table in your database

And fill your database user and password


