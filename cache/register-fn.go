package cache

type RegisterConfigs struct {
	Params     []any
	Retornos   []any
	NameUnique string
}

func Register(configs RegisterConfigs) {
	AllCache.generateComandInCacheIfNecessari(configs.NameUnique)
	parametroName := generateNomeDeParametros(configs.Params)

	AllCache[configs.NameUnique][parametroName] = cacheUnique{
		Retornos: configs.Retornos,
	}
}