package storage

import (
	"UrlShortenerGoLang/types"
	"context"
)

type UrlRepository struct {
	urls []types.Url
}

func NewUrlRepository() *UrlRepository {
	return &UrlRepository{
		urls: make([]types.Url, 0),
	}
}

func (u *UrlRepository) CreateUrl(_ context.Context, url *types.Url) (*types.Url, error) {
	u.urls = append(u.urls, *url)
	return url, nil
}

func (u *UrlRepository) GetUrlByCode(_ context.Context, code string) (*types.Url, error) {
	for i := range u.urls {
		if u.urls[i].Code == code {
			u.urls[i].VisitCount++
			return &u.urls[i], nil
		}
	}
	return nil, nil
}

func (u *UrlRepository) GetUrlByOriginalUrl(_ context.Context, originalUrl string) (*types.Url, error) {
	for i := range u.urls {
		if u.urls[i].OriginalUrl == originalUrl {
			return &u.urls[i], nil
		}
	}
	return nil, nil
}

func (u *UrlRepository) GetAllUrls(ctx context.Context) ([]types.Url, error) {
	return u.urls, nil
}
