package service

import (
	"fmt"
	"url_shotener/models"
	"url_shotener/repository"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type URLService interface {
	ShortenURL(originalURL string) (*models.Url, error)
	GetOriginalURL(code string) (*models.Url, error)
}

type urlService struct {
	repo repository.URLRepository
}

func NewUrlService(repo repository.URLRepository) URLService {
	return &urlService{repo: repo}
}

func (s *urlService) ShortenURL(originalURL string) (*models.Url, error) {
	url := models.Url{OriginalURL: originalURL}
	err := s.repo.Create(&url)
	if err != nil {
		return nil, err
	}
	url.Code = toBase62(int(url.ID))

	err = s.repo.UpdateCode(&url)
	if err != nil {
		return nil, fmt.Errorf("failed to update code: %w", err)
	}
	return &url, nil

}

func (s *urlService) GetOriginalURL(code string) (*models.Url, error) {
	url, err := s.repo.FindByCode(code)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func toBase62(num int) string {
	if num == 0 {
		return string(charset[0])
	}

	var result []byte
	base := len(charset)

	for num > 0 {
		remainder := num % base
		result = append(result, charset[remainder])
		num = num / base
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
