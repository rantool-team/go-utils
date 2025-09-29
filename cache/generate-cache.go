package cache

func (c CacheCompose) generateComandInCacheIfNecessari(comandName string) {
	if _, ok := c[comandName]; !ok {
		c[comandName] = make(cacheForComand)
	}
}
