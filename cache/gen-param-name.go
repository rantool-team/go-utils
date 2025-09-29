package cache

import "fmt"

func generateNomeDeParametros(parametros []any) string {
	res := ""
	for _, param := range parametros {
		res += generateLiteralParametroForNomeDeParametros(param)
	}

	return res
}

func generateLiteralParametroForNomeDeParametros(param any) string {
	return fmt.Sprintf("'%q'", param)
}
