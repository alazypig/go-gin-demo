package controller

import (
	"log"
	"net/http"
	"strconv"

	"gilab.com/pragmaticreviews/golang-gin-poc/entity"
	"gilab.com/pragmaticreviews/golang-gin-poc/service"
	"gilab.com/pragmaticreviews/golang-gin-poc/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

// Delete implements VideoController.
func (c *controller) Delete(ctx *gin.Context) error {
	var video entity.Video

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)

	if err != nil {
		log.Println("Invalid ID", err)
		return err
	}

	video.ID = id

	c.service.Delete(video)

	return nil
}

// Update implements VideoController.
func (c *controller) Update(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)

	if err != nil {
		log.Println("Invalid ID", err)
		return err
	}

	video.ID = id

	if err != nil {
		return err
	}
	c.service.Update(video)
	return nil
}

// ShowAll implements VideoController.
func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "List of Videos",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}

// FindAll implements VideoController.
func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

// Save implements VideoController.
func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)

	return &controller{service: service}
}
