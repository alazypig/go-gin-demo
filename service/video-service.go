package service

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/entity"
	"gilab.com/pragmaticreviews/golang-gin-poc/repository"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	Update(entity.Video) entity.Video
	Delete(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videoRepository repository.VideoRepository
}

// Delete implements VideoService.
func (service *videoService) Delete(video entity.Video) entity.Video {
	service.videoRepository.Delete(video)

	return video
}

// Update implements VideoService.
func (service *videoService) Update(video entity.Video) entity.Video {
	service.videoRepository.Update(&video)

	return video
}

// FindAll implements VideoService.
func (service *videoService) FindAll() []entity.Video {
	return service.videoRepository.FindAll()
}

// Save implements VideoService.
func (service *videoService) Save(video entity.Video) entity.Video {
	service.videoRepository.Save(&video)

	return video
}

func New(repo repository.VideoRepository) VideoService {
	return &videoService{
		videoRepository: repo,
	}
}
