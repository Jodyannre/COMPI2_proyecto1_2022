package simbolos

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"Back/analizador/expresiones"
	"strconv"
)

type AccesoStruct struct {
	NombreStruct   interface{} //Id del struct
	NombreAtributo interface{} //Id del atributo
	Tipo           Ast.TipoDato
	Fila           int
	Columna        int
}

func (a AccesoStruct) GetStruct(scope *Ast.Scope) Ast.TipoRetornado {
	return a.NombreStruct.(Ast.Expresion).GetValue(scope)
}

func (a AccesoStruct) GetNombreAtributo() string {
	return a.NombreAtributo.(expresiones.Identificador).Valor
}

func NewAccesoStruct(nombre interface{}, atributo interface{}, fila, columna int) AccesoStruct {
	nA := AccesoStruct{
		NombreStruct:   nombre,
		NombreAtributo: atributo,
		Fila:           fila,
		Columna:        columna,
		Tipo:           Ast.ACCESO_STRUCT,
	}
	return nA
}

func (a AccesoStruct) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	//Verificar que el nombre sea un id válido
	var structInstancia StructInstancia
	var simboloAtributo Ast.Simbolo
	var existeAtributo bool
	var nombreAtributo string
	var tipoParticular Ast.TipoDato
	//Para verificar si es un struct
	valor := a.NombreStruct.(Ast.Expresion).GetValue(scope)
	//Si es error devolverlo
	if valor.Tipo == Ast.ERROR {
		return valor
	}

	if valor.Tipo != Ast.STRUCT {
		fila := a.Fila
		columna := a.Columna
		msg := "Semantic error, STRUCT expected, found " + Ast.ValorTipoDato[valor.Tipo] +
			". -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	_, tipoParticular = a.NombreStruct.(Ast.Abstracto).GetTipo()

	if tipoParticular != Ast.IDENTIFICADOR {
		//Verificar que sea un struct
		if valor.Tipo != Ast.STRUCT {
			//Error, se esperaba un identificador
			fila := a.Fila
			columna := a.Columna
			msg := "Semantic error, STRUCT expected, found " + Ast.ValorTipoDato[tipoParticular] +
				". -- Line: " + strconv.Itoa(fila) +
				" Column: " + strconv.Itoa(columna)
			nError := errores.NewError(fila, columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			nError.Ambito = scope.GetTipoScope()
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}
	}
	//Recuperar el struct desde el scope
	structInstancia = valor.Valor.(StructInstancia)

	//Verificar que el valor del símbolo sea un struct
	if structInstancia.Tipo != Ast.STRUCT {
		fila := a.NombreStruct.(Ast.Abstracto).GetFila()
		columna := a.NombreStruct.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, STRUCT expected, found " + Ast.ValorTipoDato[structInstancia.Tipo] +
			". -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Verificar que el nombre del atributo sea un id
	_, tipoParticular = a.NombreAtributo.(Ast.Abstracto).GetTipo()

	if tipoParticular != Ast.IDENTIFICADOR {
		fila := a.NombreAtributo.(Ast.Abstracto).GetFila()
		columna := a.NombreAtributo.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, IDENTIFICADOR expected, found " + Ast.ValorTipoDato[tipoParticular] +
			". -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	//Recuperar el string del nombre del atributo
	nombreAtributo = a.NombreAtributo.(expresiones.Identificador).Valor

	//Recuperar la instancia
	//structInstancia = simboloStruct.Valor.(Ast.TipoRetornado).Valor.(StructInstancia)
	//Verificar que el elemento existe
	existeAtributo = structInstancia.Entorno.Exist(nombreAtributo)

	if !existeAtributo {
		//Error, ese campo no existe
		fila := a.NombreAtributo.(Ast.Abstracto).GetFila()
		columna := a.NombreAtributo.(Ast.Abstracto).GetColumna()
		msg := "Semantic error, field \"" + nombreAtributo +
			"\" doesn't exist. -- Line: " + strconv.Itoa(fila) +
			" Column: " + strconv.Itoa(columna)
		nError := errores.NewError(fila, columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		nError.Ambito = scope.GetTipoScope()
		scope.Errores.Add(nError)
		scope.Consola += msg + "\n"
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}

	//Sí existe, entonces recuperar el símbolo del atributo
	simboloAtributo = structInstancia.Entorno.GetSimbolo(nombreAtributo)

	//Verificar que el atributo sea público
	/*
		if !simboloAtributo.Publico {
			//Error, no se puede acceder a ese campo
			fila := a.Fila
			columna := a.Columna
			msg := "Semantic error, field \"" + nombreAtributo + "\" is not public." +
				" -- Line: " + strconv.Itoa(fila) +
				" Column: " + strconv.Itoa(columna)
			nError := errores.NewError(fila, columna, msg)
			nError.Tipo = Ast.ERROR_SEMANTICO
			scope.Errores.Add(nError)
			scope.Consola += msg + "\n"
			return Ast.TipoRetornado{
				Tipo:  Ast.ERROR,
				Valor: nError,
			}
		}
	*/
	return simboloAtributo.Valor.(Ast.TipoRetornado)
}

func (v AccesoStruct) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v AccesoStruct) GetFila() int {
	return v.Fila
}
func (v AccesoStruct) GetColumna() int {
	return v.Columna
}

func (a AccesoStruct) GetIdentificador() string {
	return a.NombreStruct.(expresiones.Identificador).Valor
}
