package service

import (
	"github.com/crypto-pulse/news/internal/integration/crypto_panic"
	"github.com/crypto-pulse/news/internal/model"
)

func GetNews() ([]*model.News, error) {
	return crypto_panic.FetchNews()
}
