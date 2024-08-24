package main

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

type (
	CacheApi struct {
		Cache        *cache.Cache
		authorizeKey string
	}

	CacheInterface interface {
		HandleCacheApi(c *fiber.Ctx) error
	}

	Temp struct {
	}
)

func newCacheApi(cache *cache.Cache, authorizeKey string) CacheInterface {
	return &CacheApi{
		Cache:        cache,
		authorizeKey: authorizeKey,
	}
}

func (ca *CacheApi) HandleCacheApi(c *fiber.Ctx) error {

	cacheData, exist := ca.Cache.Get(API_CACHE_KEY)
	if exist {
		if data, ok := cacheData.([]Project); ok {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"msg":  "ok",
				"data": data,
			})
		}
	}

	res, err := handleCache(ca.Cache, ca.authorizeKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "ok",
		"data": res,
	})
}

func handleCache(c *cache.Cache, authorizeKey string) ([]Project, error) {
	res, err := GetAllProjects(authorizeKey)
	if err != nil {
		slog.Error("failed to handle cache", "error", err)
		return []Project{}, err
	}
	c.Set(API_CACHE_KEY, res, cache.NoExpiration)
	return res, nil
}
