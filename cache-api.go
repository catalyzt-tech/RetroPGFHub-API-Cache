package main

import (
	"encoding/json"
	"log/slog"

	"github.com/akrylysov/pogreb"
	"github.com/gofiber/fiber/v2"
)

type (
	CacheApi struct {
		Cache        *pogreb.DB
		authorizeKey string
	}

	CacheInterface interface {
		HandleCacheApi(c *fiber.Ctx) error
	}
)

// newCacheApi creates a new instance of CacheApi with the given cache and authorize key
func newCacheApi(cache *pogreb.DB, authorizeKey string) CacheInterface {
	return &CacheApi{
		Cache:        cache,
		authorizeKey: authorizeKey,
	}
}

// HandleCacheApi handles the cache API request. It checks if the cache is available,
// if so, it returns the cached data. If not, it calls the handleCache function to
// retrieve the data, store it in the cache, and return it
func (ca *CacheApi) HandleCacheApi(c *fiber.Ctx) error {
	cacheData, err := ca.Cache.Get([]byte(API_CACHE_KEY))
	if err == nil && len(cacheData) > 0 {
		var data []Project
		if err := json.Unmarshal(cacheData, &data); err == nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"msg":  "ok",
				"data": data,
			})
		}
	}

	res, err := handleCache(ca.Cache, ca.authorizeKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "ok",
		"data": res,
	})
}

// handleCache retrieves all projects using the given authorize key,
// marshals the result into JSON, stores it in the cache, and returns the result
func handleCache(c *pogreb.DB, authorizeKey string) ([]Project, error) {
	res, err := GetAllProjects(authorizeKey)
	if err != nil {
		slog.Error("failed to handle cache", "error", err)
		return []Project{}, err
	}

	jsonData, err := json.Marshal(res)
	if err != nil {
		slog.Error("failed to marshal projects", "error", err)
		return []Project{}, err
	}

	err = c.Put([]byte(API_CACHE_KEY), jsonData)
	if err != nil {
		slog.Error("failed to store in cache", "error", err)
		return []Project{}, err
	}

	return res, nil
}
