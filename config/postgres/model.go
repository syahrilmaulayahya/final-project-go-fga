package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host         string
	Port         string
	DatabaseName string
	User         string
	Password     string
}

type PostgresClient interface {
	GetClient() *gorm.DB
}

type PostgresClientImpl struct {
	cln    *gorm.DB
	config Config
}

func NewPostgresConnection(config Config) PostgresClient {
	connectionString := fmt.Sprintf(`
	host=%s
		port=%s
		user=%s
		password=%s
		dbname=%s`,
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DatabaseName,
	)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &PostgresClientImpl{cln: db, config: config}
}

func (p *PostgresClientImpl) GetClient() *gorm.DB {
	return p.cln
}
