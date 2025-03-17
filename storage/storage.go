package storage

import (
	"UrlShortenerGoLang/types"
	"context"
)

type Storage interface {
	CreateUrl(ctx context.Context, url *types.Url) (*types.Url, error)
	GetUrlByCode(ctx context.Context, code string) (*types.Url, error)
	GetUrlByOriginalUrl(ctx context.Context, originalUrl string) (*types.Url, error)
	GetAllUrls(ctx context.Context) ([]types.Url, error)
}
