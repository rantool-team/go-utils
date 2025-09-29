package cache

func AssignWithCacheAndReturnIfExists(comandUnique string, params []any, assignFn func(retornos []any)) bool {
	comand, ok := AllCache[comandUnique]
	if !ok {
		return false
	}

	parametros := generateNomeDeParametros(params)
	retornos, ok := comand[parametros]
	if !ok {
		return false
	}

	assignFn(retornos.Retornos)
	return true
}
