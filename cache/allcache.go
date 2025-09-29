package cache

type CacheCompose map[string]cacheForComand

type cacheForComand map[string]cacheUnique

type cacheUnique struct {
	Retornos []any
}

var AllCache = make(CacheCompose)
