# MyGram API (Final Project GO-FGA)

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)


## About <a name = "about"></a>

This is API for post a photo, comment on a photo, and see other people

## Getting Started <a name = "getting_started"></a>

To run this api, you need to clone this repository
```
git clone https://github.com/syahrilmaulayahya/final-project-go-fga
```
and then run the program using 
```
go run main.go
```
### Prerequisites

Before run the program, you need to install this :

Gin-Gonic
```
go get github.com/gin-gonic/gin
```
Godotenv
```
go get github.com/joho/godotenv
```
Gorm
```
go get gorm.io/gorm
```
Postgres driver
```
go get gorm.io/driver/postgres
```
Google UUID
```
go get github.com/google/uuid
```
Kataras JWT
```
go get github.com/kataras/jwt
```
Golang Bcrypt
```
go get golang.org/x/crypto/bcrypt
```

## Usage <a name = "usage"></a>

To see all available API, please see this postman collection

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/4018ce66626ab62f9c04?action=collection%2Fimport#?env%5BNew%20Environment%5D=W3sia2V5Ijoiand0IiwidmFsdWUiOiJleUpoYkdjaU9pSklVekkxTmlJc0luUjVjQ0k2SWtwWFZDSjkuZXlKcFpDSTZNaXdpWlhod0lqbzFNak15T0RnME5qRXdmUS5ROGNrTjIzejNjb2laMTZJd1JfSFptS0xnZGlyTzR5cHFQMTl4Z3g5akpnIiwiZW5hYmxlZCI6dHJ1ZSwic2Vzc2lvblZhbHVlIjoiZXlKaGJHY2lPaUpJVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SnFkR2tpT2lJeE56a3daak5pWWkwMlpESmpMVFF4Tm1VdFlURTBOaTB5WlRRNU5qSTBNalkxWXpVaUxDSnpkV0lpT2pJc0ltbHpjeUk2SW0xNVozSmhiUzVqYjIwaUxDSmguLi4iLCJzZXNzaW9uSW5kZXgiOjB9LHsia2V5Ijoiand0YWRtaW4iLCJ2YWx1ZSI6ImV5SmhiR2NpT2lKSVV6STFOaUlzSW5SNWNDSTZJa3BYVkNKOS5leUpwWkNJNk1Td2laWGh3SWpveE5qTXpOall3T1RBMWZRLmlFUTF2UGNhSlc0aXhVOEFJdzN4OElJV3g2djhmRTM1Mk9kaDMwZFdsb1UiLCJlbmFibGVkIjp0cnVlLCJzZXNzaW9uVmFsdWUiOiJleUpoYkdjaU9pSklVekkxTmlJc0luUjVjQ0k2SWtwWFZDSjkuZXlKcFpDSTZNU3dpWlhod0lqb3hOak16TmpZd09UQTFmUS5pRVExdlBjYUpXNGl4VThBSXczeDhJSVd4NnY4ZkUzNTJPZGgzMGRXbG9VIiwic2Vzc2lvbkluZGV4IjoxfSx7ImtleSI6Imhvc3QiLCJ2YWx1ZSI6IjE4LjIxNi4yNTMuMzMiLCJlbmFibGVkIjp0cnVlLCJzZXNzaW9uVmFsdWUiOiIxOC4yMTYuMjUzLjMzIiwic2Vzc2lvbkluZGV4IjoyfV0=)
