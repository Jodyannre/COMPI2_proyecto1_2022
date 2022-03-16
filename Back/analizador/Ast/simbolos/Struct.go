package simbolos

import (
	"Back/analizador/Ast"
	"Back/analizador/errores"
	"strconv"

	"github.com/colegno/arraylist"
)

type StructInstancia struct {
	Plantilla   Ast.TipoRetornado
	Tipo        Ast.TipoDato
	Mutable     bool
	Entorno     *Ast.Scope
	AtributosIn *arraylist.List
	Fila        int
	Columna     int
}

func NewStructInstancia(plantilla Ast.TipoRetornado, atributos *arraylist.List, mutable bool, fila, columna int) StructInstancia {
	//Variables para la validación de tipos
	nS := StructInstancia{
		Plantilla:   plantilla,
		Tipo:        Ast.STRUCT,
		Mutable:     mutable,
		AtributosIn: atributos,
		Fila:        fila,
		Columna:     columna,
	}
	return nS
}

func (s StructInstancia) GetValue(scope *Ast.Scope) Ast.TipoRetornado {
	var plantilla StructTemplate
	if s.Plantilla.Tipo != Ast.STRUCT {
		//Error, porque la plantilla no es struct
		fila := s.Fila
		columna := s.Columna
		msg := "Semantic error, an STRUCT was expected." +
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

	simboloPlantilla := scope.Exist_fms(s.Plantilla.Valor.(string))
	newScope := Ast.NewScope(s.Plantilla.Valor.(string), scope)

	//Verificar que la plantilla exista o que no haya algún tipo de error
	if simboloPlantilla.Tipo == Ast.ERROR_NO_EXISTE {
		fila := s.Fila
		columna := s.Columna
		msg := "Semantic error, \"" + s.Plantilla.Valor.(string) + "\" doesn't exist." +
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

	if simboloPlantilla.Tipo == Ast.ERROR_ACCESO_PRIVADO {
		fila := s.Fila
		columna := s.Columna
		msg := "Semantic error, \"" + s.Plantilla.Valor.(string) + "\" is private to this scope." +
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

	//Verificar que sea un tipo struct
	if simboloPlantilla.Valor.(Ast.TipoRetornado).Tipo != Ast.STRUCT {
		fila := s.Fila
		columna := s.Columna
		msg := "Semantic error, \"" + s.Plantilla.Valor.(string) + "\" isn't a STRUCT." +
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
	//Recuperar el struct
	plantilla = simboloPlantilla.Valor.(Ast.TipoRetornado).Valor.(StructTemplate)

	if plantilla.AtributosIn.Len() != s.AtributosIn.Len() {
		//Error la cantidad de atributos no concuerda
		fila := s.Fila
		columna := s.Columna
		msg := "Semantic error, expected (" + strconv.Itoa(plantilla.AtributosIn.Len()) + ") values, found (" +
			strconv.Itoa(s.AtributosIn.Len()) + ") values." +
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

	//Verificar si los atributos existen y si son del mismo tipo
	for i := 0; i < s.AtributosIn.Len(); i++ {
		atributoActual := s.AtributosIn.GetValue(i).(Atributo)
		_, ok := plantilla.Atributos[atributoActual.Nombre]
		//Verificar que exista el atributo
		if !ok {
			//Error, ese nombre de atributo no existe
			//struct `Personaje` has no field named `vida`
			fila := atributoActual.Fila
			columna := atributoActual.Columna
			msg := "Semantic error, STRUCT " + plantilla.Nombre + " has no field named \"" + atributoActual.Nombre + "\"" +
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
		//Verificar si los tipos son correctos
		//Variables para validar el tipo

		//Get el valor del atributo
		valorAtt := atributoActual.GetValue(scope)
		if valorAtt.Tipo == Ast.ERROR {
			return valorAtt
		}
		attActual := valorAtt.Valor.(Atributo)
		if valorAtt.Tipo == Ast.ERROR {
			return valorAtt
		}

		attPlantilla := plantilla.Atributos[atributoActual.Nombre]
		validadorTipo := CompararTipos(attActual.TipoAtributo.(Ast.TipoRetornado),
			attPlantilla.TipoAtributo.(Ast.TipoRetornado))
		if !validadorTipo {
			return GetmsjError(validadorTipo, attActual, attPlantilla, scope)
		}
		//Todo bien ,entonces crear el struct
		nuevoSimbolo := Ast.NewSimbolo(atributoActual.Nombre, attActual.Valor, atributoActual.Fila, atributoActual.Columna, attActual.TipoAtributo.(Ast.TipoRetornado).Tipo, atributoActual.Mutable, atributoActual.Publico)

		newScope.Add(nuevoSimbolo)
	}
	//Agregar el scope al struct y finalizar retornando el scope
	s.Entorno = &newScope
	return Ast.TipoRetornado{
		Tipo:  Ast.STRUCT,
		Valor: s,
	}
}

func GetmsjError(validadorTipo bool, atributo, template Atributo, scope *Ast.Scope) Ast.TipoRetornado {
	fila := atributo.Fila
	columna := atributo.Columna
	msg := ""
	msg = "Semantic error,  can't assign " + Tipo_String(atributo.TipoAtributo.(Ast.TipoRetornado)) +
		" to field named \"" + template.Nombre + "\" " + Tipo_String(template.TipoAtributo.(Ast.TipoRetornado)) +
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

func (v StructInstancia) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.EXPRESION, v.Tipo
}

func (v StructInstancia) GetFila() int {
	return v.Fila
}
func (v StructInstancia) GetColumna() int {
	return v.Columna
}

func (s StructInstancia) GetPlantilla() string {
	return s.Plantilla.Valor.(string)
}
