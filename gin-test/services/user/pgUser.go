package user

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/luviz/gin-test/datasources"
	"github.com/luviz/gin-test/models"
	"gorm.io/gorm"
)

type PGUserConfig struct {
	PG datasources.PostgresDsnConfig
}

type PGUser struct {
	db     gorm.DB
	ctx    context.Context
	user   gorm.Interface[UserPGStore]
	logger logr.Logger
}

type UserPGStore struct {
	gorm.Model
	Id   string
	Name string
	Mail string
}

func (user UserPGStore) TDO() models.User {
	return models.User{
		Id:   user.Id,
		Name: user.Name,
		Mail: user.Mail,
	}
}

func InitPGUser(Config PGUserConfig, logger logr.Logger) (*PGUser, error) {
	db, err := datasources.InitGORM(Config.PG.DSN())
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&UserPGStore{})
	ormUser := gorm.G[UserPGStore](db)

	return &PGUser{
		ctx:    context.Background(),
		db:     *db,
		user:   ormUser,
		logger: logger.WithName("pg-user-service"),
	}, nil
}

func (s PGUser) Get(id string) (models.User, bool) {
	user, err := s.user.Where("Id = ?", id).Find(s.ctx)
	if err != nil {
		return models.User{}, false
	}
	return user[0].TDO(), true
}

func (s PGUser) List() []models.User {
	users, err := s.user.Find(s.ctx)
	if err != nil {
		return []models.User{}
	}
	dto := []models.User{}
	for _, v := range users {
		dto = append(dto, v.TDO())
	}
	return dto
}

func (s PGUser) Create(dto models.User) error {
	return s.user.Create(s.ctx, &UserPGStore{
		Id:   dto.Id,
		Name: dto.Name,
		Mail: dto.Mail,
	})
}
