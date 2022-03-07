parser grammar Nparser;

options{
    tokenVocab = Nlexer;
    language = Go;
}


@header{
    import "github.com/colegno/arraylist"
    import "Back/analizador/Ast"
    import "Back/analizador/expresiones"
    import "Back/analizador/instrucciones"
    import "Back/analizador/exp_ins"
    import "Back/analizador/transferencia"
    import "Back/analizador/bucles"
}

/* 
@members{
}
*/


inicio returns[*arraylist.List lista] 
            : instrucciones	
            {
                $lista = $instrucciones.list
            }
;


bloque returns[*arraylist.List list] 
            : LLAVE_IZQ instrucciones LLAVE_DER	
            {
                $list = $instrucciones.list
            }
;

instrucciones returns [*arraylist.List list]
			@init{
				$list =  arraylist.New()
			}
			: e += instruccion*  {
				listInt := localctx.(*InstruccionesContext).GetE()
					for _, e := range listInt {
						$list.Add(e.GetEx())
					}
			}
;

instruccion returns[interface{} ex] 
			:expresion 		    	{$ex = $expresion.ex}	
            |declaracion PUNTOCOMA  {$ex = $declaracion.ex}	
            |asignacion PUNTOCOMA   {$ex = $asignacion.ex}
            |control_if             {$ex = $control_if.ex}	 
            |control_match          {$ex = $control_match.ex}   
            |control_loop           {$ex = $control_loop.ex}
            |control_while          {$ex = $control_while.ex}
            |ibreak PUNTOCOMA       {$ex = $ibreak.ex}             
            |icontinue PUNTOCOMA    {$ex = $icontinue.ex} 
            |ireturn PUNTOCOMA      {$ex = $ireturn.ex} 
            |printNormal PUNTOCOMA  {$ex = $printNormal.ex} 
            |printFormato PUNTOCOMA {$ex = $printFormato.ex} 

;

/* 
instrucciones returns[*arraylist.List list]  : instruccion;


instruccion returns[*arraylist.List list] : expresion;
*/


declaracion returns[Ast.Instruccion ex]
    : LET ID IGUAL expresion
        {
            fila := $LET.line
            columna := $LET.pos
            $ex = instrucciones.NewDeclaracion($ID.text,Ast.INDEFINIDO,
            false,false,Ast.VOID,$expresion.ex,fila,columna)
        }
    | LET ID IGUAL control_expresion
        {
            fila := $LET.line
            columna := $LET.pos
            $ex = instrucciones.NewDeclaracion($ID.text,Ast.INDEFINIDO,
            false,false,Ast.VOID,$control_expresion.ex,fila,columna)
        }
    | LET ID DOSPUNTOS tipo_dato IGUAL expresion
        {
            fila := $LET.line
            columna := $LET.pos
            $ex = instrucciones.NewDeclaracion($ID.text,$tipo_dato.ex,
            false,false,Ast.VOID,$expresion.ex,fila,columna)            
        }
    | LET ID DOSPUNTOS tipo_dato IGUAL control_expresion
        {
            fila := $LET.line
            columna := $LET.pos
            $ex = instrucciones.NewDeclaracion($ID.text,$tipo_dato.ex,
            false,false,Ast.VOID,$control_expresion.ex,fila,columna)            
        }
    | LET MUT ID IGUAL expresion
        {
            fila := $LET.line
            columna := $LET.pos
            $ex = instrucciones.NewDeclaracion($ID.text,Ast.INDEFINIDO,
                true,false,Ast.VOID,$expresion.ex,fila,columna)                 
        }
    | LET MUT ID IGUAL control_expresion
    {
        fila := $ID.line
        columna := $ID.pos
        $ex = instrucciones.NewDeclaracion($ID.text,Ast.INDEFINIDO,
            true,false,Ast.VOID,$control_expresion.ex,fila,columna)  
    }
    | LET MUT ID DOSPUNTOS tipo_dato
        {        
            fila := $LET.line
            columna := $LET.pos
            valor := expresiones.NewPrimitivo(nil, Ast.NULL,fila,columna)
            $ex = instrucciones.NewDeclaracion($ID.text,$tipo_dato.ex,
            true,false,Ast.VOID,valor,fila,columna)                
        }
    | LET MUT ID DOSPUNTOS tipo_dato IGUAL expresion
        {
            fila := $LET.line
            columna := $LET.pos
            $ex = instrucciones.NewDeclaracion($ID.text,$tipo_dato.ex,
            true,false,Ast.VOID,$expresion.ex,fila,columna)               
        }
    | LET MUT ID DOSPUNTOS tipo_dato IGUAL control_expresion
        {
            fila := $LET.line
            columna := $LET.pos
            $ex = instrucciones.NewDeclaracion($ID.text,$tipo_dato.ex,
            true,false,Ast.VOID,$control_expresion.ex,fila,columna)               
        }
    //| STRUCT ID_CAMEL LLAVE_IZQ atributos LLAVE_DER
;


asignacion returns[Ast.Instruccion ex]
    : ID IGUAL expresion
    {
        fila := $ID.line
        columna := $ID.pos
        $ex = instrucciones.NewAsignacion($ID.text,$expresion.ex,fila,columna)
    }
    | ID IGUAL control_expresion
    {
        fila := $ID.line
        columna := $ID.pos
        $ex = instrucciones.NewAsignacion($ID.text,$control_expresion.ex,fila,columna)
    }
;


/* 
atributos returns[*arraylist.List list]
@init{ list := arrayList.New()}
    : atributos atributo
    | atributo
;

atributo returns[simbolos.Atributo ex]
    : ID DOSPUNTOS tipo_dato
;
*/

expresion returns[Ast.Expresion ex] 
    :   op=(RESTA|NOT) op_izq= expresion
        {
            fila := $op.line
            columna := $op.pos
            $ex = expresiones.NewOperation($op_izq.ex,$op.text,nil,true,fila,columna)
        }
    |   op_izq=expresion op=(MULTIPLICACION|DIVISION|MODULO) op_der=expresion
        {
            fila := $op.line
            columna := $op.pos
            $ex = expresiones.NewOperation($op_izq.ex,$op.text,$op_der.ex,false,fila,columna)
        }    
    
    |   op_izq=expresion op=(SUMA|RESTA) op_der=expresion
        {
            fila := $op.line
            columna := $op.pos
            $ex = expresiones.NewOperation($op_izq.ex,$op.text,$op_der.ex,false,fila,columna)
        }
    |   op_izq=expresion op=(MAYOR_I|MAYOR|MENOR_I|MENOR|IGUALDAD|DISTINTO) op_der=expresion
        {
            fila := $op.line
            columna := $op.pos
            $ex = expresiones.NewOperation($op_izq.ex,$op.text,$op_der.ex,false,fila,columna)
        }
    |   op_izq=expresion op=(IGUALDAD|DISTINTO) op_der=expresion
        {
            fila := $op.line
            columna := $op.pos
            $ex = expresiones.NewOperation($op_izq.ex,$op.text,$op_der.ex,false,fila,columna)
        }
    |   op_izq=expresion AND op_der=expresion
        {
            fila := $AND.line
            columna := $AND.pos
            $ex = expresiones.NewOperation($op_izq.ex,$AND.text,$op_der.ex,false,fila,columna)
        }
    |   op_izq=expresion OR op_der=expresion
        {
            fila := $OR.line
            columna := $OR.pos
            $ex = expresiones.NewOperation($op_izq.ex,$OR.text,$op_der.ex,false,fila,columna)
        }
    |   PAR_IZQ expresion PAR_DER
        {
            $ex = $expresion.ex
        }
    |   ID		
        {
            id := $ID.text
            fila := $DECIMAL.line
            columna := $DECIMAL.pos
            $ex = expresiones.NewIdentificador (id,Ast.IDENTIFICADOR,fila,columna)
        }
    |   TRUE        
        {
            valor := true
            fila := $DECIMAL.line
            columna := $DECIMAL.pos
            $ex = expresiones.NewPrimitivo(valor, Ast.BOOLEAN,fila,columna)
        }         
    |   FALSE
        {
            valor := false
            fila := $DECIMAL.line
            columna := $DECIMAL.pos
            $ex = expresiones.NewPrimitivo(valor, Ast.BOOLEAN,fila,columna)
        }     
    |   CARACTER
        {
            valor := $CARACTER.text
            valor = valor[1:len(valor)-1]
            fila := $DECIMAL.line
            columna := $DECIMAL.pos
            $ex = expresiones.NewPrimitivo(valor, Ast.CHAR,fila,columna)
        }   
    |   DECIMAL
        {
            valor,err := strconv.ParseFloat($DECIMAL.text,64)
            if err!= nil{
                fmt.Println("Hay un error en el get número")
                fmt.Println(err)
            }
            fila := $DECIMAL.line
            columna := $DECIMAL.pos
            $ex = expresiones.NewPrimitivo(valor, Ast.F64,fila,columna)
        }   
    |   NUMERO
        {
            valor,err := strconv.Atoi($NUMERO.text)
            if err!= nil{
                fmt.Println("Hay un error en el get número")
                fmt.Println(err)
            }
            fila := $NUMERO.line
            columna := $NUMERO.pos
            $ex = expresiones.NewPrimitivo(valor, Ast.I64,fila,columna)
        }   
    |   CADENA
        {
            fila := $CADENA.line
            columna := $CADENA.pos
            valor := $CADENA.text
            valor = valor[1:len(valor)-1]
            $ex = expresiones.NewPrimitivo(valor, Ast.STR,fila,columna)
        }             
;

tipo_dato returns[Ast.TipoDato ex]
    :   BOOL    {$ex = Ast.BOOLEAN}
    |   CHAR    {$ex = Ast.CHAR}
    |   I64     {$ex = Ast.I64}
    |   F64     {$ex = Ast.F64}
    |   STR     {$ex = Ast.STR}
    |   STRING  {$ex = Ast.STRING}
    |   USIZE   {$ex = Ast.USIZE}
;





control_if returns[Ast.Instruccion ex]
	:IF expresion bloqueIf = bloque
	{
		fila:= $IF.line
		columna:= $IF.pos
		columna++
		lista_null := arraylist.New()
		$ex = exp_ins.NewIF($expresion.ex,$bloqueIf.list,lista_null,Ast.IF,fila,columna,false)
	}	
	|IF expresion bloqueIf = bloque ELSE bloqueElse = bloque
	  
	{
		fila:= $IF.line
		columna:= $IF.pos
		columna++
        lista_entonces := arraylist.New()
        lista_null := arraylist.New()
        Else := exp_ins.NewIF ($expresion.ex,$bloqueElse.list,lista_null,Ast.ELSE,fila,columna,false)
		lista_entonces.Add(Else)
		$ex = exp_ins.NewIF($expresion.ex,$bloqueIf.list,lista_entonces,Ast.IF,fila,columna,false)	
	}
	|IF expresion bloqueIf = bloque bloque_else_if
	{
		fila:= $IF.line
		columna:= $IF.pos
		columna++
		lista_entonces := $bloque_else_if.list
		$ex = exp_ins.NewIF($expresion.ex,$bloqueIf.list,lista_entonces,Ast.IF,fila,columna,false)		
	}
	|IF expresion bloqueIf = bloque bloque_else_if ELSE bloqueElse = bloque
	{
		fila:= $IF.line
		columna:= $IF.pos
		columna++
        lista_null := arraylist.New()
        Else := exp_ins.NewIF ($expresion.ex,$bloqueElse.list,lista_null,Ast.ELSE,fila,columna,false)
		lista_entonces := $bloque_else_if.list
        lista_entonces.Add(Else)
		$ex = exp_ins.NewIF($expresion.ex,$bloqueIf.list,lista_entonces,Ast.IF,fila,columna,false)		
	}
;


bloque_else_if returns [*arraylist.List list]
@init{ $list = arraylist.New()}
: lista += else_if+ {
                    listas := localctx.(*Bloque_else_ifContext).GetLista()
                    for _, e := range listas {
                        $list.Add(e.GetEx())
                    }
    }
;


else_if returns [Ast.Instruccion ex]
    : ELSE IF expresion bloque
    {
        fila:= $ELSE.line
		columna:= $ELSE.pos
		columna++
        lista_null := arraylist.New()
        $ex = exp_ins.NewIF($expresion.ex,$bloque.list,lista_null,Ast.ELSEIF,fila,columna,false)	
    }
;





control_if_exp returns[Ast.Instruccion ex]
	:IF expresion bloqueIf = bloque
	{
		fila:= $IF.line
		columna:= $IF.pos
		columna++
		lista_null := arraylist.New()
		$ex = exp_ins.NewIF($expresion.ex,$bloqueIf.list,lista_null,Ast.IF_EXPRESION,fila,columna,true)
	}	
	|IF expresion bloqueIf = bloque ELSE bloqueElse = bloque
	  
	{
		fila:= $IF.line
		columna:= $IF.pos
		columna++
        lista_entonces := arraylist.New()
        lista_null := arraylist.New()
        Else := exp_ins.NewIF ($expresion.ex,$bloqueElse.list,lista_null,Ast.ELSE_EXPRESION,fila,columna,true)
		lista_entonces.Add(Else)
		$ex = exp_ins.NewIF($expresion.ex,$bloqueIf.list,lista_entonces,Ast.IF_EXPRESION,fila,columna,true)	
	}
	|IF expresion bloqueIf = bloque bloque_else_if_exp
	{
		fila:= $IF.line
		columna:= $IF.pos
		columna++
		lista_entonces := $bloque_else_if_exp.list
		$ex = exp_ins.NewIF($expresion.ex,$bloqueIf.list,lista_entonces,Ast.IF_EXPRESION,fila,columna,true)		
	}
	|IF expresion bloqueIf = bloque bloque_else_if_exp ELSE bloqueElse = bloque
	{
		fila:= $IF.line
		columna:= $IF.pos
		columna++
        lista_null := arraylist.New()
        Else := exp_ins.NewIF ($expresion.ex,$bloqueElse.list,lista_null,Ast.ELSE_EXPRESION,fila,columna,true)
		lista_entonces := $bloque_else_if_exp.list
        lista_entonces.Add(Else)
		$ex = exp_ins.NewIF($expresion.ex,$bloqueIf.list,lista_entonces,Ast.IF_EXPRESION,fila,columna,true)		
	}
;


bloque_else_if_exp returns [*arraylist.List list]
@init{ $list = arraylist.New()}
: lista += else_if_exp+ {
                    listas := localctx.(*Bloque_else_if_expContext).GetLista()
                    for _, e := range listas {
                        $list.Add(e.GetEx())
                    }
    }
;


else_if_exp returns [Ast.Instruccion ex]
    : ELSE IF expresion bloque
    {
        fila:= $ELSE.line
		columna:= $ELSE.pos
		columna++
        lista_null := arraylist.New()
        $ex = exp_ins.NewIF($expresion.ex,$bloque.list,lista_null,Ast.ELSEIF_EXPRESION,fila,columna,true)	
    }
;


control_expresion returns [Ast.Instruccion ex]
    : control_if_exp {$ex = $control_if_exp.ex}
    | control_match_exp {$ex = $control_match_exp.ex}
    | control_loop_exp {$ex = $control_loop_exp.ex}
;


control_match returns[Ast.Instruccion ex]
    : MATCH expresion LLAVE_IZQ control_case LLAVE_DER
    {
        fila := $MATCH.line
        columna := $MATCH.line -1
        $ex = exp_ins.NewMatch($expresion.ex,$control_case.list,Ast.MATCH,fila,columna)
    }
;


control_case returns[*arraylist.List list]
@init{$list = arraylist.New()}
    : lista += cases+ 
        {
            listas := localctx.(*Control_caseContext).GetLista()
            for _, e := range listas {
                $list.Add(e.GetEx())
        }
    }
    
;



cases returns[Ast.Instruccion ex]
    : case_match CASE bloque COMA
    {

        fila := $CASE.line
        columna := $CASE.line -1
        //Verificar si lo que vienen es un default
        listaTemp := $case_match.list
        _, tipo := listaTemp.GetValue(0).(Ast.Abstracto).GetTipo()
        if tipo == Ast.DEFAULT{
            $ex = exp_ins.NewCase($case_match.list,$bloque.list, Ast.CASE,fila,columna,true)
        }else{
            $ex = exp_ins.NewCase($case_match.list,$bloque.list, Ast.CASE,fila,columna,false)
        }      
    }
;


case_match returns[*arraylist.List list]
@init{$list = arraylist.New()}
    : lista_cases = case_match O expresion 
        {
            $lista_cases.list.Add($expresion.ex)
            $list = $lista_cases.list
        }
    | expresion 
        {
            $list.Add($expresion.ex)
        }
    | DEFAULT
        {
            fila := $DEFAULT.line
            columna := $DEFAULT.pos
            expresion := expresiones.NewPrimitivo(false, Ast.DEFAULT, fila, columna)
            $list.Add(expresion)
        }
;





control_match_exp returns[Ast.Instruccion ex]
    : MATCH expresion LLAVE_IZQ control_case_exp LLAVE_DER
    {
        fila := $MATCH.line
        columna := $MATCH.line -1
        $ex = exp_ins.NewMatch($expresion.ex,$control_case_exp.list,Ast.MATCH_EXPRESION,fila,columna)
    }
;


control_case_exp returns[*arraylist.List list]
@init{$list = arraylist.New()}
    : lista += cases_exp+ 
        {
            listas := localctx.(*Control_case_expContext).GetLista()
            for _, e := range listas {
                $list.Add(e.GetEx())
        }
    }
    
;



cases_exp returns[Ast.Instruccion ex]
    : case_match_exp CASE bloque COMA
    {
        fila := $CASE.line
        columna := $CASE.line -1
        //Verificar si lo que vienen es un default
        listaTemp := $case_match_exp.list
        _, tipo := listaTemp.GetValue(0).(Ast.Abstracto).GetTipo()
        if tipo == Ast.DEFAULT{
            $ex = exp_ins.NewCase($case_match_exp.list,$bloque.list, Ast.CASE_EXPRESION,fila,columna,true)
        }else{
            $ex = exp_ins.NewCase($case_match_exp.list,$bloque.list, Ast.CASE_EXPRESION,fila,columna,false)
        }      
    }
;


case_match_exp returns[*arraylist.List list]
@init{$list = arraylist.New()}
    : lista_cases = case_match_exp O expresion 
        {
            $lista_cases.list.Add($expresion.ex)
            $list = $lista_cases.list
        }
    | expresion 
        {
            $list.Add($expresion.ex)
        }
    | DEFAULT
        {
            fila := $DEFAULT.line
            columna := $DEFAULT.pos
            expresion := expresiones.NewPrimitivo(false, Ast.DEFAULT, fila, columna)
            $list.Add(expresion)
        }
;


ireturn returns[Ast.Instruccion ex]
    : RETURN 
        {
            fila := $RETURN.line
            columna := $RETURN.pos
            $ex = transferencia.NewReturn(Ast.RETURN,nil,fila,columna)
        }
    | RETURN expresion
        {
            fila := $RETURN.line
            columna := $RETURN.pos
            $ex = transferencia.NewReturn(Ast.RETURN_EXPRESION,$expresion.ex,fila,columna)
        }
;

ibreak returns[Ast.Instruccion ex]
    : BREAK 
        {
            fila := $BREAK.line
            columna := $BREAK.pos
            $ex = transferencia.NewBreak(Ast.BREAK,nil,fila,columna)            
        }
    | BREAK expresion
        {
            fila := $BREAK.line
            columna := $BREAK.pos
            $ex = transferencia.NewBreak(Ast.BREAK_EXPRESION,$expresion.ex,fila,columna)                   
        }
;

icontinue returns[Ast.Instruccion ex]
    : CONTINUE 
        {
            fila := $CONTINUE.line
            columna := $CONTINUE.pos
            $ex = transferencia.NewContinue(fila,columna)                 
        }
;

control_loop returns[Ast.Instruccion ex]
    : LOOP bloque 
    {
        fila := $LOOP.line
        columna := $LOOP.pos
        $ex = bucles.NewLoop(Ast.LOOP,$bloque.list,fila,columna)
    }
;


control_loop_exp returns[Ast.Instruccion ex]
    : LOOP bloque 
    {
        fila := $LOOP.line
        columna := $LOOP.pos
        $ex = bucles.NewLoop(Ast.LOOP_EXPRESION,$bloque.list,fila,columna)
    }
;

printNormal returns[Ast.Instruccion ex]
    : PRINT PAR_IZQ expresion PAR_DER
        {
            fila := $PRINT.line
            columna := $PRINT.pos
            $ex = instrucciones.NewPrint($expresion.ex,Ast.PRINT,fila,columna)            
        }
;

printFormato returns[Ast.Instruccion ex]
    : PRINT PAR_IZQ CADENA COMA expresiones=elementosPrint PAR_DER
        {
            fila := $PRINT.line
            columna := $PRINT.pos
            valor := $CADENA.text
            valor = valor[1:len(valor)-1]
            $ex = instrucciones.NewPrintF($expresiones.list,valor,Ast.PRINTF,fila,columna)       
        }
;

elementosPrint returns[*arraylist.List list]
@init{$list = arraylist.New()}
    : lista_elementos = elementosPrint COMA expresion 
        {
            $lista_elementos.list.Add($expresion.ex)
            $list = $lista_elementos.list
        }
    | expresion 
        {
            $list.Add($expresion.ex)
        }
;


control_while returns[Ast.Instruccion ex]
    : WHILE expresion bloque 
        {
            fila := $WHILE.line
            columna := $WHILE.pos
            $ex = bucles.NewWhile(Ast.WHILE,$expresion.ex,$bloque.list,fila,columna)
        }
;
