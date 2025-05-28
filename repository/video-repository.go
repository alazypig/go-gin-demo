package repository

import (
	"errors"

	"gilab.com/pragmaticreviews/golang-gin-poc/entity"
	"gilab.com/pragmaticreviews/golang-gin-poc/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go.uber.org/zap"
)

type VideoRepository interface {
	Save(video *entity.Video)
	Update(video *entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

// 处理Person逻辑
func (db *database) handlePersonAssociation(tx *gorm.DB, video *entity.Video) error {
	var existedPerson entity.Person

	if err := tx.Where("first_name = ? and last_name = ?", video.Author.FirstName, video.Author.LastName).First(&existedPerson).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newPerson := entity.Person{
				FirstName: video.Author.FirstName,
				LastName:  video.Author.LastName,
				Age:       video.Author.Age,
				Email:     video.Author.Email,
			}

			if err := tx.Create(&newPerson).Error; err != nil {
				return err
			}

			video.PersonID = newPerson.ID
		} else {
			return err
		}
	} else {
		video.PersonID = existedPerson.ID
	}

	video.Author = nil // 清除Author字段，避免循环引用

	return nil
}

// Delete implements VideoRepository.
func (db *database) Delete(video entity.Video) {
	db.connection.Delete(&video)
}

// FindAll implements VideoRepository.
func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Set("gorm:auto_preload", true).Find(&videos)

	return videos
}

// Save implements VideoRepository.
func (db *database) Save(video *entity.Video) {
	tx := db.connection.Begin()

	if tx.Error != nil {
		logger.Error("Failed to start transaction:", zap.Error(tx.Error))
		return
	}

	if err := db.handlePersonAssociation(tx, video); err != nil {
		logger.Error("Error handling person association:", zap.Error(err))
		tx.Rollback()
	}

	if err := tx.Create(&video).Error; err != nil {
		logger.Error("Error creating video:", zap.Error(err))
		tx.Rollback()
		return
	}

	tx.Commit()
}

// Update implements VideoRepository.
func (db *database) Update(video *entity.Video) {
	tx := db.connection.Begin()

	if tx.Error != nil {
		logger.Error("Failed to start transaction:", zap.Error(tx.Error))
		return
	}

	if err := db.handlePersonAssociation(tx, video); err != nil {
		logger.Error("Error handling person association:", zap.Error(err))
		tx.Rollback()
	}

	if err := tx.Save(&video).Error; err != nil {
		logger.Error("Error updating video:", zap.Error(err))
		tx.Rollback()
		return
	}

	tx.Commit()
}

func (db *database) CloseDB() {
	err := db.connection.Close()

	if err != nil {
		panic("failed to close database")
	}
}

func NewVideoRepository() VideoRepository {
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entity.Video{}, &entity.Person{})

	return &database{
		connection: db,
	}
}
