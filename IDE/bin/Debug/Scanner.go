package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"unicode"
)

type token struct {
	token_type string
	lexema     string
	linea      int16
}

var ncol int16 = 0
var nline int16 = 1
var buffer string
var tamArchivo int = 0
var ContadorChar int = 0
var bandSalto bool = false
var Tokens []token

func Scanner() {

	var Token token
	TOKENS := make(map[int16]string)
	TOKENS[1] = "TKN_NUM"
	TOKENS[2] = "TKN_ASIGN"
	TOKENS[3] = "TKN_LPAREN"
	TOKENS[4] = "TKN_RPAREN"
	TOKENS[5] = "TKN_MINUS"
	TOKENS[6] = "TKN_ADD"
	TOKENS[7] = "TKN_IDEN"
	TOKENS[8] = "TKN_ERR"
	TOKENS[9] = "TKN_IF"
	TOKENS[10] = "TKN_INT"
	TOKENS[11] = "TKN_EOF"
	TOKENS[12] = "TKN_SEMICOLOM" //Token para el punto y coma
	TOKENS[13] = "TKN_FLOAT"
	TOKENS[14] = "TKN_STRING"
	TOKENS[15] = "TKN_DOUBLE"
	TOKENS[16] = "TKN_PUNFLOT"
	TOKENS[17] = "TKN_FOR"
	TOKENS[18] = "TKN_CASE"
	TOKENS[19] = "TKN_SWITCH"
	TOKENS[20] = "TKN_WHILE"
	TOKENS[21] = "TKN_LLLAVE"
	TOKENS[22] = "TKN_RLLAVE"
	TOKENS[23] = "TKN_LCOR"
	TOKENS[24] = "TKN_RCOR"
	TOKENS[25] = "TKN_FLOTANTE"
	TOKENS[26] = "TKN_ELSE"
	TOKENS[27] = "TKN_RETURN"
	TOKENS[28] = "TKN_MAIN"
	TOKENS[29] = "TKN_BOOL"
	TOKENS[30] = "TKN_FI"
	TOKENS[31] = "TKN_UNTIL"
	TOKENS[32] = "TKN_READ"
	TOKENS[33] = "TKN_WRITE"
	TOKENS[34] = "TKN_NOT"
	TOKENS[35] = "TKN_OR"
	TOKENS[36] = "TKN_AND"
	TOKENS[37] = "TKN_MUL"
	TOKENS[38] = "TKN_DIV"
	TOKENS[39] = "TKN_POT"
	TOKENS[40] = "TKN_MENOR"
	TOKENS[41] = "TKN_MENOR_IGUAL"
	TOKENS[42] = "TKN_MAYOR"
	TOKENS[43] = "TKN_MAYOR_IGUAL"
	TOKENS[44] = "TKN_COMP_IGUALDAD"
	TOKENS[45] = "TKN_NEGACION"
	TOKENS[46] = "TKN_DIFERENTE"
	TOKENS[47] = "TKN_COMA"
	TOKENS[48] = "TKN_COMENT_LINE"
	TOKENS[49] = "TKN_SALTO"
	TOKENS[50] = "TKN_TRUE"
	TOKENS[51] = "TKN_FALSE"
	TOKENS[52] = "TKN_THEN"
	TOKENS[53] = "TKN_DO"
	TOKENS[54] = "TKN_PROGRAM"

	band := false
	Token = getToken(TOKENS)
	for {
		if Token.token_type != "TKN_SALTO" {
			Tokens = append(Tokens, Token)
			fmt.Printf("(%s, %s, Linea: %d) \n", Token.token_type, Token.lexema, Token.linea)
		}
		f, err := os.OpenFile("salida.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		} else {
			if Token.token_type != "TKN_SALTO" {
				fmt.Fprintf(f, "%s %s %d\n", Token.token_type, Token.lexema, Token.linea)
			}

			f.Close()

		}
		if Token.token_type == "TKN_EOF" {
			band = true
			break
		}

		if band == true {
			break
		} else {
			Token = getToken(TOKENS)
		}
	}
	fmt.Println()
	fmt.Println(nline, " Lineas Analizadas ")
	fmt.Println()
	fmt.Println("Parte del analizador Sintactico")
	//Analyze()
}

//Funcion para retornar los Tokens al analizador Sintactico
func GetTokens() token {
	index := pos
	pos++
	return Tokens[index]
}

func getToken(Tokens map[int16]string) token {
	var c byte
	estado := "START"
	var Token token

	index := 0

	for estado != "TERMINADO" {
		switch estado {
		case "START":
			c = getChar()
			for {
				if isDelimit(c) {
					c = getChar()
				} else {
					break
				}
			}
			if unicode.IsLetter(rune(c)) {
				estado = "IDENTIFICADOR"
				index++
				Token.lexema = Token.lexema + string(c)
			} else if c == '\r' {
				Token.token_type = Tokens[49]
				Token.linea = nline
				estado = "TERMINADO"
				index++
			} else if unicode.IsNumber(rune(c)) {
				estado = "NUMERO"
				index++
				Token.lexema = Token.lexema + string(c)
			} else if c == ',' {
				Token.token_type = Tokens[47]
				Token.linea = nline
				estado = "TERMINADO"
				index++
				Token.lexema = Token.lexema + string(c)
			} else if c == '{' {
				Token.token_type = Tokens[21]
				Token.linea = nline
				estado = "TERMINADO"
				index++
				Token.lexema = Token.lexema + string(c)
			} else if c == '}' {
				Token.token_type = Tokens[22]
				Token.linea = nline
				estado = "TERMINADO"
				index++
				Token.lexema = Token.lexema + string(c)
			} else if c == '[' {
				Token.token_type = Tokens[23]
				Token.linea = nline
				estado = "TERMINADO"
				index++
				Token.lexema = Token.lexema + string(c)
			} else if c == ']' {
				Token.token_type = Tokens[24]
				Token.linea = nline
				estado = "TERMINADO"
				index++
				Token.lexema = Token.lexema + string(c)
			} else if c == '(' {
				Token.token_type = Tokens[3] //Lparen
				Token.linea = nline
				estado = "TERMINADO"
				index++
				Token.lexema = Token.lexema + string(c)
			} else if c == ')' {
				Token.token_type = Tokens[4]
				Token.linea = nline
				estado = "TERMINADO"
				index++
				Token.lexema = Token.lexema + string(c)
			} else if c == ';' {
				Token.token_type = Tokens[12]
				Token.linea = nline
				estado = "TERMINADO"
				index++
				Token.lexema = Token.lexema + string(c)
			} else if c == '=' {
				estado = "COMPARAIGUALDAD"
				index++
				Token.lexema = Token.lexema + string(c)
			} else if c == '+' {
				Token.token_type = Tokens[6]
				Token.linea = nline
				estado = "TERMINADO"
				index++
				Token.lexema = Token.lexema + string(c)
			} else if c == '-' {
				Token.token_type = Tokens[5]
				Token.linea = nline
				estado = "TERMINADO"
				index++
				Token.lexema = Token.lexema + string(c)
			} else if c == '*' {
				Token.token_type = Tokens[37]
				Token.linea = nline
				estado = "TERMINADO"
				index++
				Token.lexema = Token.lexema + string(c)
			} else if c == '/' {
				Token.token_type = Tokens[38]
				Token.linea = nline
				estado = "TERMINADO"
				index++
				Token.lexema = Token.lexema + string(c)
			} else if c == '^' {
				Token.token_type = Tokens[39]
				Token.linea = nline
				estado = "TERMINADO"
				index++
				Token.lexema = Token.lexema + string(c)
			} else if c == '<' {
				estado = "MENORIGUAL"
				index++
				Token.lexema = Token.lexema + string(c)
			} else if c == '>' {
				estado = "MAYORIGUAL"
				index++
				Token.lexema = Token.lexema + string(c)
			} else if c == '!' {
				estado = "DIFERENTE"
				index++
				Token.lexema = Token.lexema + string(c)
			} else if c == '~' {
				Token.token_type = Tokens[11]
				Token.linea = nline
				estado = "TERMINADO"
			} else {
				//token[8] = Tokens[8]
				estado = "xxx"
			}
			break
		case "IDENTIFICADOR":
			c = getChar()
			index++

			if !(unicode.IsLetter(rune(c)) || unicode.IsDigit(rune(c)) || c == '_') {
				Token.token_type = Tokens[7]
				Token.linea = nline
				estado = "TERMINADO"
				ungetchar()
				index--
				//lexema[index] = ' '
				Token = BuscarPalReservada(Token.lexema, Token, Tokens)
			} else {
				Token.lexema = Token.lexema + string(c)
			}
			break
		case "COMENTARIO":
			c = getChar()
			index++
			if c != '/' {
				if c == '*' {
					estado = "COMENTARIOBLOQUE"
					Token.lexema = Token.lexema + string(c)
				} else {
					Token.token_type = Tokens[38]
					Token.linea = nline
					estado = "TERMINADO"
					ungetchar()
					index--
				}
			} else if c == '/' {
				estado = "COMENTARIOLINE"
				Token.lexema = Token.lexema + string(c)
			}
			break
		case "COMENTARIOLINE":
			c = getChar()
			index++
			if ncol != 0 {
				Token.lexema = Token.lexema + string(c)
			} else {
				Token.token_type = Tokens[48]
				Token.linea = nline
				estado = "TERMINADO"
				ungetchar()
				index++
			}
			break
		case "NUMERO":
			c = getChar()
			index++

			if !unicode.IsNumber(rune(c)) && c != '.' {
				Token.token_type = Tokens[1]
				Token.linea = nline
				estado = "TERMINADO"
				ungetchar()
				index--
				//lexema[pos] = ' '
			} else if c == '.' {
				estado = "FLOTANTE"
				index++
				Token.lexema = Token.lexema + string(c)
			} else {
				Token.lexema = Token.lexema + string(c)
			}
			break
		case "FLOTANTE":
			c = getChar()
			index++
			if !unicode.IsNumber(rune(c)) {
				Token.token_type = Tokens[25]
				Token.linea = nline
				estado = "TERMINADO"
				ungetchar()
				index--
				//lexema[pos] = ' '
			} else {
				Token.lexema = Token.lexema + string(c)
			}
			break
		case "MENORIGUAL":
			c = getChar()
			index++
			if c != '=' {
				Token.token_type = Tokens[40]
				estado = "TERMINADO"
				ungetchar()
				index--
			} else if c == '=' {
				Token.token_type = Tokens[41]
				Token.linea = nline
				Token.lexema = Token.lexema + string(c)
				estado = "TERMINADO"
			}
			break
		case "MAYORIGUAL":
			c = getChar()
			index++
			if c != '=' {
				Token.token_type = Tokens[42]
				Token.linea = nline
				estado = "TERMINADO"
				ungetchar()
				index--
			} else if c == '=' {
				Token.token_type = Tokens[43]
				Token.linea = nline
				Token.lexema = Token.lexema + string(c)
				estado = "TERMINADO"
			}
			break
		case "COMPARAIGUALDAD":
			c = getChar()
			index++
			if c != '=' {
				Token.token_type = Tokens[2]
				Token.linea = nline
				estado = "TERMINADO"
				ungetchar()
				index--
			} else if c == '=' {
				Token.token_type = Tokens[44]
				Token.linea = nline
				Token.lexema = Token.lexema + string(c)
				estado = "TERMINADO"
			}
			break
		case "DIFERENTE":
			c = getChar()
			index++
			if c != '=' {
				Token.token_type = Tokens[45]
				Token.linea = nline
				estado = "TERMINADO"
				ungetchar()
				index--
			} else if c == '=' {
				Token.token_type = Tokens[46]
				Token.linea = nline
				Token.lexema = Token.lexema + string(c)
				estado = "TERMINADO"
			}
			break
		default:
			Token.token_type = Tokens[8]
			Token.linea = nline
			estado = "TERMINADO"
			index++
			Token.lexema = Token.lexema + string(c)
			//break
		}
	}
	if Token.token_type == "TKN_ERR" {
		f, err := os.OpenFile("errores.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		} else {
			fmt.Fprintf(f, "Error Linea %d:%d, el caracter '%c' no es valido \n", nline, ncol, c)
			f.Close()
		}

		fmt.Printf("Error Linea %d:%d, el caracter '%q' no es valido \n", nline, ncol, c)
		return Token
	} else {
		if Token.token_type != "TKN_SALTO" {
			fmt.Printf("lexema: %s\n", Token.lexema)
		}
		return Token
	}

}

func BuscarPalReservada(lexema string, tokenR token, Tokens map[int16]string) token {
	PAL_RESERVADAS := make(map[int16]string)
	PAL_RESERVADAS[9] = "if"
	PAL_RESERVADAS[10] = "int"
	PAL_RESERVADAS[13] = "float"
	PAL_RESERVADAS[14] = "string"
	PAL_RESERVADAS[15] = "double"
	PAL_RESERVADAS[17] = "for"
	PAL_RESERVADAS[18] = "case"
	PAL_RESERVADAS[19] = "switch"
	PAL_RESERVADAS[20] = "while"
	PAL_RESERVADAS[26] = "else"
	PAL_RESERVADAS[27] = "return"
	PAL_RESERVADAS[28] = "main"
	PAL_RESERVADAS[29] = "bool"
	PAL_RESERVADAS[30] = "fi"
	PAL_RESERVADAS[31] = "until"
	PAL_RESERVADAS[32] = "read"
	PAL_RESERVADAS[33] = "write"
	PAL_RESERVADAS[34] = "NOT"
	PAL_RESERVADAS[35] = "OR"
	PAL_RESERVADAS[36] = "AND"
	PAL_RESERVADAS[50] = "true"
	PAL_RESERVADAS[51] = "false"
	PAL_RESERVADAS[31] = "until"
	PAL_RESERVADAS[52] = "then"
	PAL_RESERVADAS[53] = "do"
	PAL_RESERVADAS[54] = "program"

	var Token token
	cont := 0
	letra := 0
	for i := 0; i < len(lexema); i++ {
		if unicode.IsLetter(rune(lexema[i])) {
			letra++
		}
	}

	//fmt.Println("caracteres: ", letra, " palReser: ", len(PAL_RESERVADAS[10]))
	for key, value := range PAL_RESERVADAS {
		cont = 0
		if letra == len(value) {
			for i := 0; i < len(value); i++ {
				if value[i] == lexema[i] {
					cont++
				} else {
					//cont--
				}

			}
		}
		if cont == len(value) {
			Token.token_type = Tokens[key]
			Token.lexema = PAL_RESERVADAS[key]
			Token.linea = nline
			return Token
			break
		}
	}
	return tokenR
}

func ungetchar() {
	ncol--
	ContadorChar--
}

//Ya funciona de manera Correcta
func getChar() byte {
	//var n int16 = 0
	if ContadorChar == 0 {
		archivo := flag.String("archivo", "", "El nombre del archivo")
		flag.Parse()
		content, err := ioutil.ReadFile(*archivo)
		if err == io.EOF {
			fmt.Println("Error fin de archivo")
		}
		if err != nil {
			log.Fatal(err)
		} else {
			tamArchivo = len(string(content))
			buffer = string(content)
		}
	}
	pos := ContadorChar
	ContadorChar++
	ncol++

	if ContadorChar <= tamArchivo {
		return buffer[pos]
	} else {
		return '~'
	}

}

//Funcion Para saber si es un delimitador ya funciona de manera correcta
func isDelimit(letra byte) bool {
	DELIMITADORES := [3]byte{'\n', '\t', ' '}
	band := false
	for i := 0; i < len(DELIMITADORES); i++ {
		if letra == DELIMITADORES[i] {
			band = true
			if letra == '\n' {
				nline++
				ncol = 0

			}
		}
	}
	return band

}
