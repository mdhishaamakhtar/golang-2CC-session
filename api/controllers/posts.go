package controllers

import (
	"errors"
	"github.com/mdhishaamakhtar/learnGo/pkg/database"
	"github.com/mdhishaamakhtar/learnGo/pkg/models"
)

func AddPosts(title, description string) (*models.Posts, error) {
	db := database.GetDB()
	post := models.Posts{
		Title:       title,
		Description: description,
	}
	err := db.Create(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func FetchPost(id string) (*models.Posts, error) {
	db := database.GetDB()
	var post models.Posts
	err := db.Model(models.Posts{}).Where("id=?", id).Find(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func UpdatePost(id, title, description string) (*models.Posts, error) {
	db := database.GetDB()
	var post models.Posts
	err := db.Model(&models.Posts{}).Where("id=?", id).Find(&post).Error
	if err != nil {
		return nil, err
	}
	if post.ID == "" {
		return nil, errors.New("no such post found")
	}
	post.Title = title
	post.Description = description
	err = db.Save(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func DeletePost(id string) error {
	db := database.GetDB()
	err := db.Where("id=?", id).Delete(&models.Posts{}).Error
	if err != nil {
		return err
	}
	return nil
}
