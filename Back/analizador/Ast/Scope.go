package Ast

import (
	"strings"

	"github.com/colegno/arraylist"
)

type Scope struct {
	Nombre               string
	prev                 *Scope
	tablaSimbolos        map[string]interface{}
	tablaModulos         map[string]interface{}
	tablaFunciones       map[string]interface{}
	tablaStructs         map[string]interface{}
	tablaSimbolosReporte *arraylist.List
	Errores              *arraylist.List
	Consola              string
	Global               bool
}

func NewScope(name string, prev *Scope) Scope {

	nuevo := Scope{Nombre: name, prev: prev}
	nuevo.Errores = arraylist.New()
	nuevo.tablaSimbolos = make(map[string]interface{})
	nuevo.tablaModulos = make(map[string]interface{})
	nuevo.tablaFunciones = make(map[string]interface{})
	nuevo.tablaStructs = make(map[string]interface{})
	nuevo.tablaSimbolosReporte = arraylist.New()
	nuevo.Global = false
	return nuevo
}

func (scope *Scope) Add(simbolo Simbolo) {
	var global *Scope = scope
	id := strings.ToUpper(simbolo.Identificador)
	//Agregar el símbolo al scope actual
	scope.tablaSimbolos[id] = simbolo

	//Recuperar el scope global
	for scope_actual := scope; scope_actual.prev != nil; scope_actual = scope_actual.prev {
		global = scope_actual
	}
	//Crear el símbolo para la tabla de reporte de símbolos
	simboloReporte := simbolo.NewSimboloReporte(scope)
	global.tablaSimbolosReporte.Add(simboloReporte)
}

func (scope *Scope) Exist(ident string) bool {
	id := strings.ToUpper(ident)
	for scope_actual := scope; scope_actual != nil; scope_actual = scope_actual.prev {
		for key, _ := range scope_actual.tablaSimbolos {
			if key == id {
				return true
			}
		}
	}
	return false
}

func (scope *Scope) Exist_fms(ident string) Simbolo {
	//Primero conseguir el scope global
	var scope_global *Scope
	var scope_encontrado *Scope
	var retorno TipoRetornado
	id := strings.ToUpper(ident)
	if scope.prev != nil {
		for scope_global = scope; scope_global.prev != nil; scope_global = scope_global.prev {
			//Buscando el scope global
		}
	} else {
		scope_global = scope
	}

	//Verificar que la fms exista y si puede ser accedida
	retorno = buscarMap(id, MODULO, scope_global, scope)
	if retorno.Tipo != NULL {
		scope_encontrado = retorno.Valor.(Simbolo).Entorno
		if scope_global != scope_encontrado {
			//Los scopes no son iguales, verificar que la función sea pública
			if scope_encontrado != scope {
				if !retorno.Valor.(Simbolo).Publico {
					return NewSimbolo("", nil, -1, -1, ERROR_ACCESO_PRIVADO, false, false)
				}
			}

		}
		return retorno.Valor.(Simbolo)
	}
	retorno = buscarMap(id, STRUCT, scope_global, scope)
	if retorno.Tipo != NULL {
		scope_encontrado = retorno.Valor.(Simbolo).Entorno
		if scope_global != scope_encontrado {
			//Los scopes no son iguales, verificar que el struct sea público
			if scope_encontrado != scope {
				if !retorno.Valor.(Simbolo).Publico {
					return NewSimbolo("", nil, -1, -1, ERROR_ACCESO_PRIVADO, false, false)
				}
			}
		}
		return retorno.Valor.(Simbolo)
	}
	retorno = buscarMap(id, FUNCION, scope_global, scope)
	if retorno.Tipo != NULL {
		scope_encontrado = retorno.Valor.(Simbolo).Entorno
		if scope_global != scope_encontrado {
			//Los scopes no son iguales, verificar que el módulo sea pública
			if scope_encontrado != scope {
				if !retorno.Valor.(Simbolo).Publico {
					return NewSimbolo("", nil, -1, -1, ERROR_ACCESO_PRIVADO, false, false)
				}
			}
		}
		return retorno.Valor.(Simbolo)
	}
	return NewSimbolo("", nil, -1, -1, ERROR_NO_EXISTE, false, false)
}

func (scope *Scope) Exist_fms_declaracion(ident string) Simbolo {
	//Primero conseguir el scope global
	var scope_global *Scope
	var retorno TipoRetornado
	id := strings.ToUpper(ident)
	if scope.prev != nil {
		for scope_global = scope; scope_global.prev != nil; scope_global = scope_global.prev {
			//Buscando el scope global
		}
	} else {
		scope_global = scope
	}

	//Verificar que la fms exista y si puede ser accedida
	retorno = buscarMap(id, MODULO, scope_global, scope)
	if retorno.Tipo != NULL {

		return retorno.Valor.(Simbolo)
	}
	retorno = buscarMap(id, STRUCT, scope_global, scope)
	if retorno.Tipo != NULL {

		return retorno.Valor.(Simbolo)
	}
	retorno = buscarMap(id, FUNCION, scope_global, scope)
	if retorno.Tipo != NULL {
		return retorno.Valor.(Simbolo)
	}
	return NewSimbolo("", nil, -1, -1, ERROR_NO_EXISTE, false, false)
}

func buscarMap(id string, tipo TipoDato, scope *Scope, local *Scope) TipoRetornado {
	var encontrado = false
	var simbolo Simbolo
	switch tipo {
	case FUNCION:
		for key, value := range scope.tablaFunciones {
			if key == id {
				simbolo = value.(Simbolo)
				encontrado = true
				break
			}
		}

	case MODULO:
		for key, value := range scope.tablaModulos {
			if key == id {
				//Verificar que sea publica
				simbolo = value.(Simbolo)
				encontrado = true
				break
			}
		}

	case STRUCT:
		for key, value := range scope.tablaStructs {
			if key == id {
				//Verificar que sea publica
				simbolo = value.(Simbolo)
				encontrado = true
				break
			}
		}
	}
	if encontrado {
		return TipoRetornado{
			Tipo:  SIMBOLO,
			Valor: simbolo,
		}
	}
	return TipoRetornado{
		Valor: true,
		Tipo:  NULL,
	}
}

/*fms = funcion modulo struct*/
func (scope *Scope) Addfms(simbolo Simbolo) {
	var scope_global *Scope = scope
	id := strings.ToUpper(simbolo.Identificador)
	//Recuperar el scope global
	if scope.prev != nil {
		for scope_global = scope; scope_global.prev != nil; scope_global = scope_global.prev {
			//Buscando el scope global
		}
	} else {
		scope_global = scope
	}
	switch simbolo.Tipo {
	case FUNCION:
		scope_global.tablaFunciones[id] = simbolo
	case MODULO:
		scope_global.tablaModulos[id] = simbolo
	case STRUCT_TEMPLATE:
		scope_global.tablaStructs[id] = simbolo
	}
}

func (scope *Scope) AddfmsParticular(simbolo Simbolo) {
	scope_global := scope
	id := strings.ToUpper(simbolo.Identificador)

	switch simbolo.Tipo {
	case FUNCION:
		scope_global.tablaFunciones[id] = simbolo
	case MODULO:
		scope_global.tablaModulos[id] = simbolo
	case STRUCT_TEMPLATE:
		scope_global.tablaStructs[id] = simbolo
	}
}

/*
func (scope *Scope) Exist_acceso(lista *arraylist.List) bool {
	var global *Scope
	var validador bool
	var simbolo_actual Simbolo
	//Recuperar el global
	for scope_actual := scope; scope_actual != nil; scope_actual = scope_actual.prev {
		if scope_actual.prev == nil{
			global = scope_actual
		}
	}
	for i:= 0; i< lista.Len();i++{
		//Obtener los padres
		var actual string = lista.GetValue(i).(string)
		for key, simbolo := range scope.tablaSimbolos {
			if key == actual {
				validador = true
				simbolo_actual = simbolo.(Simbolo)
				break
			}
		}
		if validador{
			actual = simbolo.
		}
	}



	return false
}
*/

func (scope *Scope) Exist_actual(ident string) bool {
	id := strings.ToUpper(ident)
	for key, _ := range scope.tablaSimbolos {
		if key == id {
			return true
		}
	}
	return false
}

func (scope *Scope) UpdateSimbolo(ident string, valorNuevo Simbolo) {
	id := strings.ToUpper(ident)
	for scope_actual := scope; scope_actual != nil; scope_actual = scope_actual.prev {

		for key, _ := range scope_actual.tablaSimbolos {
			if key == id {
				scope_actual.tablaSimbolos[key] = valorNuevo
				return
			}
		}
	}

}

func (scope *Scope) GetSimbolo(ident string) Simbolo {
	id := strings.ToUpper(ident)

	for scope_actual := scope; scope_actual != nil; scope_actual = scope_actual.prev {
		for key, simboloRetorno := range scope_actual.tablaSimbolos {
			if key == id {
				nsimbolo := simboloRetorno.(Simbolo)
				return nsimbolo
			}
		}
	}
	var simboloNull Simbolo
	return simboloNull
}

func (scope *Scope) GetSimboloReferencia(ident string) Simbolo {
	id := strings.ToUpper(ident)

	for scope_actual := scope; scope_actual != nil; scope_actual = scope_actual.prev {
		for key, simboloRetorno := range scope_actual.tablaSimbolos {
			if key == id && scope_actual != scope {
				nsimbolo := simboloRetorno.(Simbolo)
				return nsimbolo
			}
		}
	}
	var simboloNull Simbolo
	return simboloNull
}

func (scope *Scope) UpdateReferencias() TipoRetornado {
	var valor TipoRetornado
	var scopeReferencia *Scope
	var simboloReferencia Simbolo
	for _, simboloGuardado := range scope.tablaSimbolos {
		if simboloGuardado.(Simbolo).Referencia {
			valor = simboloGuardado.(Simbolo).Valor.(TipoRetornado)
			simboloReferencia = *simboloGuardado.(Simbolo).Referencia_puntero
			scopeReferencia = simboloReferencia.Entorno
			simboloReferencia.Valor = valor
			//Get el símbolo
			scopeReferencia.UpdateSimbolo(simboloReferencia.Identificador, simboloReferencia)
			break
		}
	}
	return TipoRetornado{Valor: true, Tipo: EJECUTADO}
}

func (s *Scope) UpdateScopeGlobal() {
	//Primero actualizar todas los valores por referencia

	//Obtener el scope global
	var scope_global *Scope
	if s.prev != nil {
		for scope_global = s; scope_global.prev != nil; scope_global = scope_global.prev {
			//Buscando el scope global
		}
	} else {
		scope_global = s
	}
	if s != scope_global {
		scope_global.Consola += s.Consola
		for i := 0; i < s.Errores.Len(); i++ {
			elemento := s.Errores.GetValue(i)
			scope_global.Errores.Add(elemento)
		}
	}
}
