package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/Craftbec/Shortener_link/config"
	"github.com/Craftbec/Shortener_link/internal/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	db *gorm.DB
}

type links struct {
	ShortLink    string `gorm:"type:char(10);primaryKey"`
	OriginalLink string `gorm:"type:text;unique;not_null"`
}

func NewDB(conf *config.Config) (*DB, error) {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		conf.POSTGRES.Host, conf.POSTGRES.User, conf.POSTGRES.Password, conf.POSTGRES.Dbname, conf.POSTGRES.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&links{})
	if err != nil {
		return nil, err
	}
	return &DB{db: db}, nil
}

func (s *DB) Get(ctx context.Context, short string) (string, error) {
	tmp := links{ShortLink: short}
	if result := s.db.WithContext(ctx).First(&tmp); result.RowsAffected == 0 {
		return "", errors.NotFound
	}
	return tmp.OriginalLink, nil
}

func (s *DB) Post(ctx context.Context, original string, short string) error {
	tmp := links{OriginalLink: original, ShortLink: short}
	result := s.db.WithContext(ctx).Create(&tmp)
	return result.Error
}

func (s *DB) CheckPost(ctx context.Context, original string) (string, error) {
	var res links
	tmp := links{OriginalLink: original}
	if result := s.db.WithContext(ctx).First(&res, &tmp); result.RowsAffected == 0 {
		return "", errors.NotFound
	}
	return res.ShortLink, nil
}

func (s *DB) GracefulStopDB() {
	links, _ := s.db.DB()
	err := links.Close()
	if err != nil {
		log.Fatal(err)
	}
}
