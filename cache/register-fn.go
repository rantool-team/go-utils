package cache

type RegisterConfigs struct {
	Params     []any
	Retornos   []any
	NameUnique string
}

func Register(configs RegisterConfigs) {
	AllCache.generateComandInCacheIfNecessari(configs.NameUnique)
	AllCache[configs.NameUnique] = generateParamsReturns(configs.Params, configs.Retornos, AllCache[configs.NameUnique])
}

func generateParamsReturns(params []any, returns []any, previousName cacheForComand) cacheForComand {
	new := previousName
	parametroName := generateNomeDeParametros(params)

	new[parametroName] = cacheUnique{
		Retornos: returns,
	}

	return new
}
