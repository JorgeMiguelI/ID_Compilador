package main

import (
	"fmt"
	"os"
	"strconv"
)

var TipoNodo = [4]string{"Stmtk", "Expk", "Deck", "Programk"}
var TipoDatoDec = []string{"int", "float", "bool"}
var TipoSententencia = [6]string{"Ifk", "RepeatK", "Assignk", "ReadK", "WriteK", "doUntilk"}
var TipoExpresion = [5]string{"Opk", "ConstK", "Idk", "OpLogicok", "Boolk"}
var pos int = 0
var bandif = false
var bandwhile = false
var bandasign = false

type Node struct {
	child         [3]*Node
	sibling       *Node
	TipoNodo      string
	TipoSentencia string
	TipoExpresion string
	TipoDato      string
	op            string
	val           int
	name          string
}

var Token token
var line int
var ERROR bool = false
var ERRORSIMICOLOM bool = false
var TokenError token
var bandsentenciaif bool = false

//Analizador Sintactico
func Analyze() *Node {
	var t *Node
	Token = GetTokens()
	t = program()
	if Token.token_type != "TKN_EOF" {
		TokenError = Token
		syntaxError("")
	}
	return t
}

func syntaxError(msg string) {
	f, err := os.OpenFile("erroresSintactico.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	} else {
		fmt.Fprintf(f, "%s \n", msg)
	}
	fmt.Println(msg)
	//ERROR = true

}

func newExpNode(tipoExpresion string) *Node {
	var t *Node = new(Node)
	for i := 0; i < 3; i++ {
		t.child[i] = nil
	}
	t.sibling = nil
	t.TipoNodo = TipoNodo[1]
	t.TipoExpresion = tipoExpresion
	t.val = 0
	t.name = ""
	t.op = ""
	return t
}
func newProgram(program string) *Node {
	var t *Node = new(Node)
	for i := 0; i < 3; i++ {
		t.child[i] = nil
	}
	t.sibling = nil
	t.TipoNodo = TipoNodo[3]
	t.TipoExpresion = ""
	t.val = 0
	t.name = program
	t.op = ""
	return t
}
func newDecNode(tipoDato string) *Node {
	var t *Node = new(Node)
	for i := 0; i < 3; i++ {
		t.child[i] = nil
	}
	t.sibling = nil
	t.TipoNodo = TipoNodo[2]
	t.TipoExpresion = ""
	t.TipoDato = tipoDato
	t.val = 0
	t.name = ""
	t.op = ""
	return t
}
func newStmtNode(tipoSentencia string) *Node {
	var t *Node = new(Node)
	for i := 0; i < 3; i++ {
		t.child[i] = nil
	}
	t.sibling = nil
	t.TipoNodo = TipoNodo[0]
	t.TipoSentencia = tipoSentencia
	t.val = 0
	t.name = ""
	t.op = ""
	return t
}

func match(expectedToken string) {
	//fmt.Println("Evaluacion: ", Token.token_type, "vs", expectedToken)
	if Token.token_type == expectedToken {
		Token = GetTokens()
	} else {
		/*if Token.token_type == "TKN_SALTO" || Token.token_type == "TKN_EOF" {
			ERRORSIMICOLOM = true
			fmt.Println("Error Sintactico en la linea ", Token.linea, " Se esperaba ;")
		} else {*/
		//	ERROR = true
		if ERROR == true {
			ERROR = false
			TokenError = Token
			syntaxError("")
		}

		//}
		//
	}
}

func program() *Node {
	var t *Node = nil
	if Token.token_type == "TKN_PROGRAM" {
		t = newProgram("program")
		match("TKN_PROGRAM")
	} else {
		t = newProgram("")
		TokenError = Token
		msg := fmt.Sprint("Error en la linea ", TokenError.linea, " Se esperaba 'program'")
		syntaxError(msg)
		match(Token.token_type)
	}
	if Token.token_type == "TKN_LLLAVE" {
		match("TKN_LLLAVE")
	} else {
		TokenError = Token
		msg := fmt.Sprint("Error en la linea ", TokenError.linea, " Se esperaba '{'")
		syntaxError(msg)
		//match(Token.token_type)
	}
	if t != nil {
		t.child[0] = lista_declaracion()
		t.child[1] = lista_sentencias()
		if Token.token_type != "TKN_RLLAVE" {
			TokenError = Token
			TokenError = Token
			msg := fmt.Sprint("Error en la linea ", TokenError.linea, " Se esperaba '}'")
			syntaxError(msg)
			for {
				if Token.token_type != "TKN_EOF" {
					//fmt.Println(Token.token_type)
					match(Token.token_type)
				} else {
					break
				}
			}
		} else {
			match("TKN_RLLAVE")
		}

	}

	return t

}
func lista_declaracion() *Node {
	var t *Node = declaracion()
	var p *Node = t
	for {
		if Token.token_type != "TKN_EOF" {
			var q *Node
			//fmt.Println(Token.token_type)
			if Token.token_type == "TKN_SEMICOLOM" {
				match("TKN_SEMICOLOM")
			} else {
				TokenError = Token
				msg := fmt.Sprint("Error en la linea ", TokenError.linea-1, ", Se esperaba ';'")
				syntaxError(msg)
				if t != nil {
					/*var temp *Node = t
					temp = temp.sibling
					for {
						if temp.sibling != nil {
							temp = temp.sibling
						} else {
							fmt.Printf("entre %s \n", temp.child[0].name)
							temp.sibling = nil
							temp = nil

							break
						}
					}*/
					//fmt.Printf("%s \n", p.child[0].name)
					t.sibling = nil

				}

			}
			//fmt.Println(Token.token_type)
			if Token.token_type == "TKN_EOF" || Token.token_type == "TKN_RLLAVE" || Token.lexema == "if" || Token.token_type != "TKN_INT" && Token.token_type != "TKN_FLOAT" && Token.token_type != "TKN_BOOL" {
				//fmt.Println("ENTRE")
				break
			}
			q = declaracion()
			if q != nil {
				if t == nil {
					p = q
					t = p
				} else {
					p.sibling = q
					p = q

				}
			}
		} else {
			break
		}
	}
	return t
}

func declaracion() *Node {
	var t *Node = nil
	if Token.lexema == "int" || Token.lexema == "float" || Token.lexema == "bool" {
		t = newDecNode(Token.lexema)
		t.TipoDato = Token.lexema
		match(Token.token_type)
		t.child[0] = lista_id()
		if t.child[0] == nil {
			t = nil
		}
	} else {
		TokenError = Token
		msg := fmt.Sprint("Error en la linea ", TokenError.linea, ", Token ", TokenError.lexema, " no valido, Se esperaba un tipo de dato valido")
		syntaxError(msg)
		for {
			if Token.token_type != "TKN_SEMICOLOM" {
				//fmt.Println(Token.token_type)
				match(Token.token_type)
			} else {
				break
			}
		}
	}

	return t

}
func lista_id() *Node {
	var t *Node = nil
	if Token.token_type == "TKN_IDEN" {
		t = newExpNode(TipoExpresion[2])
		t.name = Token.lexema
		var p *Node = t
		match("TKN_IDEN")
		for {
			if Token.token_type == "TKN_COMA" {
				match("TKN_COMA")
				if Token.token_type != "TKN_IDEN" {
					TokenError = Token
					msg := fmt.Sprint("Error en la linea ", TokenError.linea, " Token ", TokenError.lexema, " no valido, Se esperaba un identificador")
					syntaxError(msg)
					for {
						if Token.token_type != "TKN_SEMICOLOM" {
							//fmt.Println(Token.token_type)
							match(Token.token_type)
						} else {
							break
						}
					}
					t = nil
					break
				} else {
					var q *Node = newExpNode(TipoExpresion[2])
					q.name = Token.lexema
					p.sibling = q
					p = q
					match(Token.token_type)
				}
				//t.sibling=
			} else if Token.token_type != "TKN_SEMICOLOM" && Token.lexema != "if" && Token.lexema != "int" && Token.lexema != "float" && Token.lexema != "bool" {
				TokenError = Token
				msg := fmt.Sprint("Error en la linea ", TokenError.linea, " Se esperaba ','")
				syntaxError(msg)
				for {
					if Token.token_type != "TKN_SEMICOLOM" && Token.lexema != "int" && Token.lexema != "float" && Token.lexema != "bool" && Token.lexema != "if" && Token.lexema != "write" && Token.lexema != "read" {
						//fmt.Println(Token.token_type)
						match(Token.token_type)
					} else {
						/*if Token.lexema == "int" || Token.lexema == "float" || Token.lexema == "bool" {
							//Token.token_type = "TKN_SEMICOLOM"
							//pos = pos - 2
						}*/
						break
					}
				}
				t = nil
				p = nil
				break
			} else if Token.lexema == "int" || Token.lexema == "float" || Token.lexema == "bool" {
				t = nil
				p = nil
				break
			} else {
				break
			}
		}
	} else {
		TokenError = Token
		msg := fmt.Sprint("Error en la linea ", TokenError.linea, " Token ", TokenError.lexema, " no valido, Se esperaba un identificador")
		syntaxError(msg)
		for {
			if Token.token_type != "TKN_SEMICOLOM" {
				//fmt.Println(Token.token_type)
				match(Token.token_type)
			} else {
				break
			}
		}

	}

	return t
}

func lista_sentencias() *Node {
	var t *Node = sentencia()
	var p *Node = t
	for {
		//fmt.Println("Token: ", Token.token_type)
		if Token.token_type != "TKN_EOF" || Token.token_type != "TKN_ELSE" || Token.token_type != "TKN_FI" || Token.token_type != "TKN_UNTIL" {
			var q *Node
			match("TKN_SEMICOLOM")
			if Token.token_type == "TKN_EOF" || Token.token_type == "TKN_RLLAVE" || Token.token_type == "TKN_FI" && Token.lexema != "read" && Token.lexema != "write" {
				//fmt.Printf("ENTRE con token %s \n", Token.lexema)
				break
			}
			//fmt.Printf("Entre con %s, voy en la linea %d \n", Token.lexema, Token.linea)
			q = sentencia()
			if q != nil {
				if t == nil {
					p = q
					t = p

				} else {
					p.sibling = q
					p = q
				}
			}
		} else {
			break
		}
	}
	return t
}

func sentencia() *Node {
	var t *Node = nil
	switch Token.token_type {
	case "TKN_READ":
		//bandif = false
		if bandif == true {
			bandif = false
			t = sent_read()
			bandif = true
		} else {
			t = sent_read()
		}

		break
	case "TKN_WRITE":
		//bandif = false
		if bandif == true {
			bandif = false
			t = sent_write()
			bandif = true
		} else {
			t = sent_write()
		}

		break
	case "TKN_IDEN":
		//bandif = false
		t = asignacion()
		break
	case "TKN_IF":
		bandif = true
		t = sent_If()
		break
	case "TKN_WHILE":
		if bandif == true {
			bandif = false
			bandwhile = true
			t = sent_while()
			bandif = true
		} else {
			bandwhile = true
			t = sent_while()
		}

		break
	case "TKN_DO":
		//bandif = false
		t = sent_doUntil()
		break
	default:
		//bandif = false
		TokenError = Token
		msg := fmt.Sprint("Error en la linea ", TokenError.linea, " Token ", TokenError.lexema, " no valido, Se esperaba una sentencia valida")
		syntaxError(msg)
		for {
			if Token.token_type != "TKN_SEMICOLOM" {
				//fmt.Println(Token.token_type)
				match(Token.token_type)
			} else {
				break
			}
		}
	}
	return t
}
func sent_doUntil() *Node {
	var t *Node = newStmtNode(TipoSententencia[5])
	match("TKN_DO")
	if t != nil {
		t.child[0] = bloque()
	}
	if Token.token_type == "TKN_UNTIL" {
		match("TKN_UNTIL")
		if Token.token_type == "TKN_LPAREN" {
			match("TKN_LPAREN")
		} else {
			TokenError = Token
			msg := fmt.Sprint("Error en la linea ", TokenError.linea, ", Se esperaba '('")
			syntaxError(msg)
		}
		if t != nil {
			t.child[1] = b_expresion()
			/*if t.child[1] == nil {
				t = nil
			}*/
		}
		if Token.token_type == "TKN_RPAREN" {
			match("TKN_RPAREN")
		} else {
			TokenError = Token
			msg := fmt.Sprint("Error en la linea ", TokenError.linea, ", Se esperaba ')'")
			syntaxError(msg)
			for {
				if Token.token_type != "TKN_SEMICOLOM" {
					//fmt.Println(Token.token_type)
					match(Token.token_type)
				} else {
					break
				}
			}
		}

	} else {
		TokenError = Token
		msg := fmt.Sprint("Error en la linea ", TokenError.linea, ", Se esperaba 'until'")
		syntaxError(msg)
		for {
			if Token.token_type != "TKN_SEMICOLOM" {
				//fmt.Println(Token.token_type)
				match(Token.token_type)
			} else {
				break
			}
		}
		t = nil
	}
	return t
}

func sent_while() *Node {
	var t *Node = newStmtNode(TipoSententencia[1])
	match("TKN_WHILE")
	if Token.token_type == "TKN_LPAREN" {
		match("TKN_LPAREN")
	} else {
		TokenError = Token
		msg := fmt.Sprint("Error en la linea ", TokenError.linea, ", Se esperaba '('")
		syntaxError(msg)
		t = nil
		for {
			if Token.lexema != "}" {
				match(Token.token_type)
			} else {
				if Token.lexema == "}" {
					match(Token.token_type)
				}

				break
			}
		}
	}
	if t != nil {
		t.child[0] = b_expresion()
		if t.child[0] == nil {
			//fmt.Printf("Voy en el Token %s, en la linea %d \n", Token.lexema, Token.linea)
			t = nil
			for {
				if Token.lexema != "}" {
					match(Token.token_type)
				} else {
					/*if bandif == true {

					}*/
					match(Token.token_type)
					break
				}
			}
		}
		if t != nil {
			if Token.token_type == "TKN_RPAREN" {
				match("TKN_RPAREN")
			} else {
				TokenError = Token
				msg := fmt.Sprint("Error en la linea ", TokenError.linea, ", Se esperaba ')'")
				syntaxError(msg)
				t = nil
				for {
					if Token.lexema != "}" {
						match(Token.token_type)
					} else {
						match(Token.token_type)
						break
					}
				}
			}
			if t != nil {
				t.child[1] = bloque()
				if t.child[1] == nil {
					t = nil
					//fmt.Printf("Voy en el token %s, en la linea %d", Token.lexema, Token.linea)
					for {
						if Token.lexema != "}" && Token.lexema != "read" && Token.lexema != "write" {
							match(Token.token_type)
						} else {
							if Token.lexema == "}" {
								match(Token.token_type)
							}
							break
						}
					}
				}
			}

		}

	}
	bandwhile = false
	return t
}

func sent_If() *Node {
	bandsentenciaif = true
	var t *Node = newStmtNode(TipoSententencia[0])
	match("TKN_IF") //Despues de aqui podria ir un if
	if Token.token_type == "TKN_LPAREN" {
		match("TKN_LPAREN")
	} else {
		TokenError = Token
		msg := fmt.Sprint("Error en la linea ", TokenError.linea, ", Se esperaba '('")
		syntaxError(msg)
		t = nil
		for {
			if Token.token_type != "TKN_FI" {
				match(Token.token_type)
			} else {
				break
			}
		}
	}
	if t != nil {
		t.child[0] = b_expresion()
		//fmt.Printf("Voy en el token: %s, de la linea %d \n", Token.lexema, Token.linea)
		if t.child[0] == nil {
			t = nil
		}
	}
	if t != nil {
		if Token.token_type == "TKN_RPAREN" {
			match("TKN_RPAREN")
		} else {
			TokenError = Token
			msg := fmt.Sprint("Error en la linea ", TokenError.linea, ", Se esperaba ')'")
			syntaxError(msg)
			t = nil
			for {
				if Token.token_type != "TKN_FI" {
					match(Token.token_type)
				} else {
					break
				}
			}
		}
		if t != nil {
			if Token.token_type == "TKN_THEN" {
				match("TKN_THEN")
				bandsentenciaif = false
			} else {
				TokenError = Token
				msg := fmt.Sprint("Error en la linea ", TokenError.linea, ", Se esperaba 'then'")
				syntaxError(msg)
				t = nil
				for {
					if Token.token_type != "TKN_FI" {
						match(Token.token_type)
					} else {
						break
					}
				}

			}
		}

		if t != nil {
			t.child[1] = bloque()
			if t.child[1] == nil {
				t = nil
			}
		}
		if Token.token_type == "TKN_ELSE" {
			match("TKN_ELSE")
			if t != nil {
				t.child[2] = bloque()
			}
		}
		if Token.token_type == "TKN_FI" {
			match("TKN_FI")
		} else {
			TokenError = Token
			msg := fmt.Sprint("Error en la linea ", TokenError.linea, ", Se esperaba 'fi'")
			syntaxError(msg)
			t = nil
		}
	} else {
		if Token.token_type == "TKN_FI" {
			match("TKN_FI")
		} else {
			TokenError = Token
			msg := fmt.Sprint("Error en la linea ", TokenError.linea, ", Se esperaba 'fi'")
			syntaxError(msg)
		}
	}
	bandif = false
	return t

}

func bloque() *Node {
	var t *Node = nil
	var band int = 0
	if Token.token_type == "TKN_LLLAVE" {
		match("TKN_LLLAVE")
	} else {
		//match(Token.token_type)
		TokenError = Token
		msg := fmt.Sprint("Error en la linea ", TokenError.linea-1, " Se esperaba '{'")
		syntaxError(msg)
		t = nil
		band = 1
		if bandif == true && bandwhile == false {
			for {
				if Token.token_type != "TKN_FI" {
					match(Token.token_type)
				} else {
					break
				}
			}
		} else if bandwhile == true {
			for {
				if Token.lexema != "}" {
					match(Token.token_type)
				} else {
					break
				}
			}
		}

	}
	if band == 0 {
		t = lista_sentencias()
		//fmt.Printf("estoy en el Token %s, en la linea %d", Token.lexema, Token.linea)
	}

	if t != nil {
		if Token.token_type == "TKN_RLLAVE" {
			match("TKN_RLLAVE")
		} else {
			TokenError = Token
			msg := fmt.Sprint("Error en la linea ", TokenError.linea, " Se esperaba '}'")
			syntaxError(msg)
			t = nil
			if bandif == true && bandwhile == false {
				//fmt.Println("Aqui mero entre")
				for {
					if Token.token_type != "TKN_FI" {
						match(Token.token_type)
					} else {
						break
					}
				}
			} else if bandwhile == true {
				for {
					if Token.lexema != "}" && Token.lexema != "read" && Token.lexema != "write" && Token.lexema != "if" && Token.lexema != "while" {
						match(Token.token_type)
					} else {
						break
					}
				}
			}
		}
	}

	return t
}

func sent_read() *Node {
	var t *Node = newStmtNode(TipoSententencia[3])
	match("TKN_READ")
	if Token.token_type == "TKN_IDEN" {
		t.name = Token.lexema
		match("TKN_IDEN")
		if Token.token_type != "TKN_SEMICOLOM" {
			TokenError = Token
			msg := fmt.Sprint("Error en la linea ", TokenError.linea-1, ", Se esperaba ';'")
			syntaxError(msg)
			t = nil
			for {
				if Token.lexema != "}" && Token.lexema != ";" && Token.lexema != "write" && Token.lexema != "read" && Token.lexema != "while" && Token.lexema != "if" && Token.token_type != "TKN_IDEN" {
					//fmt.Println(Token.token_type)
					match(Token.token_type)
				} else {
					//fmt.Printf("voy en el token %s, en la linea %d \n", Token.lexema, Token.linea)
					break
				}
			}
		}
	} else {
		TokenError = Token
		msg := fmt.Sprint("Error en la linea ", TokenError.linea, ", Se esperaba un identificador")
		syntaxError(msg)
		for {
			if Token.lexema != "}" && Token.lexema != ";" && Token.lexema != "write" && Token.lexema != "read" && Token.lexema != "while" && Token.lexema != "if" {
				//fmt.Println(Token.token_type)
				match(Token.token_type)
			} else {
				break
			}
		}
		t = nil
	}

	return t
}

func sent_write() *Node {
	var t *Node = newStmtNode(TipoSententencia[4])
	match("TKN_WRITE")
	if t != nil {
		t.child[0] = b_expresion()
		if t.child[0] == nil {
			t = nil
		}
	}
	if Token.lexema != ";" {
		t = nil
		TokenError = Token
		msg := fmt.Sprint("Error en la linea ", TokenError.linea-1, ", Se esperaba ';'")
		syntaxError(msg)
		for {
			if Token.lexema != "}" && Token.lexema != ";" && Token.lexema != "read" && Token.lexema != "write" {
				match(Token.token_type)
			} else {
				break
			}
		}
	}
	return t
}

func asignacion() *Node {
	bandasign = true
	var t *Node = newStmtNode(TipoSententencia[2])
	t.name = Token.lexema
	match("TKN_IDEN")
	if Token.token_type == "TKN_ASIGN" {
		match("TKN_ASIGN")
		t.child[0] = b_expresion()
		if t.child[0] == nil {
			t = nil
		} else {
			if Token.lexema != ";" {
				TokenError = Token
				msg := fmt.Sprint("Error en la linea ", TokenError.linea-1, ", Se esperaba ';'")
				syntaxError(msg)
				t = nil
				for {
					if Token.lexema != "read" && Token.lexema != "}" && Token.lexema != "write" && Token.lexema != "if" && Token.token_type != "TKN_IDEN" && Token.lexema != "while" {
						match(Token.token_type)
					} else {
						break
					}
				}
			}
		}

	} else {
		TokenError = Token
		msg := fmt.Sprint("Error en la linea ", TokenError.linea, ", Se esperaba '='")
		syntaxError(msg)
		if Token.lexema != ";" {
			TokenError = Token
			msg := fmt.Sprint("Error en la linea ", TokenError.linea, ", Se esperaba ';'")
			syntaxError(msg)
			t = nil
		}
		for {
			if Token.token_type != "TKN_SEMICOLOM" {
				//fmt.Println(Token.token_type)
				match(Token.token_type)
			} else {
				break
			}
		}
		t = nil
	}
	bandasign = false

	return t

}

func b_expresion() *Node {
	var t *Node = b_term()
	if t != nil {
		for {
			if Token.token_type == "TKN_OR" {
				var p *Node = newExpNode(TipoExpresion[3])
				p.child[0] = t
				p.op = Token.token_type
				t = p
				match(Token.token_type)
				t.child[1] = b_term()
				if t.child[1] == nil {
					t = nil
					p = nil
				}
			} else if Token.lexema != "AND" && Token.token_type != "TKN_SEMICOLOM" && Token.lexema != ")" && Token.lexema != "}" && Token.lexema != "then" && Token.lexema != "{" && Token.lexema != "read" && Token.lexema != "write" && Token.lexema != "if" && Token.lexema != "while" && Token.token_type != "TKN_IDEN" {
				TokenError = Token
				msg := fmt.Sprint("Error en la linea ", TokenError.linea, " Token ", TokenError.lexema, " no valido, Se esperaba una expresion valida ")
				syntaxError(msg)
				t = nil
				if bandif == true && bandwhile == false {
					for {
						if Token.token_type != "TKN_FI" {
							match(Token.token_type)
						} else {
							break
						}
					}
				} else if bandwhile == true {
					for {
						if Token.lexema != "}" {
							match(Token.token_type)
						} else {
							//fmt.Printf("voy en el Token %s, en la linea % \n", Token.lexema, Token.linea)
							break
						}
					}
				}
				break
			} else {
				break
			}
		}
	}
	return t
}
func b_term() *Node {
	var t *Node = not_factor()
	if t != nil {
		for {
			if Token.token_type == "TKN_AND" {
				var p *Node = newExpNode(TipoExpresion[3])
				p.child[0] = t
				p.op = Token.token_type
				t = p
				match(Token.token_type)
				p.child[1] = not_factor()
				if p.child[1] == nil {
					t = nil
				}
			} else {
				break
			}
		}
	}

	return t
}

func not_factor() *Node {
	var t *Node = nil
	if Token.token_type == "TKN_NOT" {
		t = newExpNode(TipoExpresion[3])
		t.op = Token.token_type
		match(Token.token_type)
		t.child[0] = b_factor()
		if t.child[0] == nil {
			t = nil
		}
	} else {
		t = b_factor()
	}
	return t
}

func b_factor() *Node {
	var t *Node = nil
	if Token.token_type == "TKN_TRUE" || Token.token_type == "TKN_FALSE" {
		t = newExpNode(TipoExpresion[4])
		t.name = Token.lexema
		match(Token.token_type)
	} else if Token.token_type == "TKN_LPAREN" || Token.token_type == "TKN_NUM" || Token.token_type == "TKN_IDEN" {
		t = relacion()

	} else {
		TokenError = Token
		msg := fmt.Sprint("Error en la linea ", TokenError.linea, ", Se esperaba un expresion valida")
		syntaxError(msg)
		if bandif == true && bandwhile == false {
			for {
				if Token.token_type != "TKN_FI" && Token.lexema != "read" && Token.lexema != " write" && Token.lexema != "while" && Token.lexema != "if" && Token.token_type != "TKN_IDEN" && Token.lexema != ";" {
					match(Token.token_type)
				} else {
					break
				}
			}
		} else if bandwhile == true {
			for {
				if Token.lexema != "}" && Token.lexema != "read" && Token.lexema != " write" && Token.lexema != "while" && Token.lexema != "if" && Token.token_type != "TKN_IDEN" && Token.lexema != ";" {
					match(Token.token_type)
				} else {
					break
				}
			}
		} else {
			for {
				if Token.lexema != "read" && Token.lexema != " write" && Token.lexema != "while" && Token.lexema != "if" && Token.token_type != "TKN_IDEN" && Token.lexema != ";" {
					//fmt.Println(Token.token_type)
					match(Token.token_type)
				} else {
					break
				}
			}
		}

		t = nil
	}
	return t
}

func relacion() *Node {
	var t *Node = exp()
	if t != nil {
		for {
			if Token.lexema == "<=" || Token.lexema == "<" || Token.lexema == ">" || Token.lexema == ">=" || Token.lexema == "==" || Token.lexema == "!=" {
				var p *Node = newExpNode(TipoExpresion[0])
				p.child[0] = t
				p.op = Token.token_type
				t = p
				match(Token.token_type)
				t.child[1] = exp()
				if t.child[1] == nil {
					t = nil
				}
			} else {
				break
			}
		}
	}
	return t
}

func exp() *Node {
	var t *Node = term()
	if t != nil {
		for {
			if Token.token_type == "TKN_MINUS" || Token.token_type == "TKN_ADD" {
				var p *Node = newExpNode(TipoExpresion[0])
				p.child[0] = t
				p.op = Token.token_type
				t = p
				match(Token.token_type)
				t.child[1] = term()
				if t.child[1] == nil {
					t = nil
					p = nil
				}

			} else if Token.token_type != "TKN_MUL" && Token.token_type != "TKN_DIV" && Token.token_type != "TKN_SEMICOLOM" && Token.lexema != "<=" && Token.lexema != "<" && Token.lexema != ">" && Token.lexema != ">=" && Token.lexema != "==" && Token.lexema != "!=" && Token.lexema != ")" && Token.lexema != "{" && Token.lexema != "then" && Token.lexema != "AND" && Token.lexema != "OR" && Token.lexema != "}" && Token.lexema != "write" && Token.lexema != "read" && Token.lexema != "while" && Token.token_type != "TKN_IDEN" {
				TokenError = Token
				msg := fmt.Sprint("Error en la linea ", TokenError.linea, " Token ", TokenError.lexema, " no valido, Se esperaba una expresion valida")
				syntaxError(msg)
				if bandif == true && bandwhile == false && bandasign == false {
					for {
						if Token.token_type != "TKN_FI" && Token.lexema != ";" && Token.lexema != "read" && Token.lexema != "write" && Token.lexema != "if" && Token.lexema != "while" && Token.lexema != "TKN_IDEN" && Token.lexema != ")" {
							match(Token.token_type)
						} else {
							if Token.lexema == ")" || Token.lexema == "then" {
								for {
									if Token.lexema != "fi" {
										match(Token.token_type)
									} else {
										break
									}
								}
							}
							break
						}
					}
				} else if bandwhile == true {
					for {
						if Token.lexema != "}" && Token.lexema != ";" && Token.lexema != "read" && Token.lexema != "write" && Token.lexema != "if" && Token.lexema != "while" && Token.lexema != "TKN_IDEN" && Token.lexema != ")" {
							match(Token.token_type)
						} else {
							if Token.lexema == ")" || Token.lexema == "{" {
								for {
									if Token.token_type != "}" {
										match(Token.token_type)
									} else {
										break
									}
								}
							}
							break
						}
					}
				} else if bandasign == true {
					for {
						if Token.lexema != ";" && Token.lexema != ")" && Token.lexema != "read" && Token.lexema != "write" && Token.lexema != "if" && Token.lexema != "while" {
							match(Token.token_type)
						} else {
							break
						}
					}
				} else {
					for {
						if Token.lexema != ";" && Token.lexema != "read" && Token.lexema != "write" && Token.lexema != "if" && Token.lexema != "while" && Token.lexema != "TKN_IDEN" && Token.lexema != ")" {
							//fmt.Println(Token.token_type)
							match(Token.token_type)
						} else {
							break
						}
					}
				}

				t = nil
				break
			} else {
				break
			}
		}
	}
	return t
}

func term() *Node {
	var t *Node = Factor()
	if t != nil {
		for {
			if Token.token_type == "TKN_MUL" || Token.token_type == "TKN_DIV" {

				var p *Node = newExpNode(TipoExpresion[0])
				p.child[0] = t
				p.op = Token.token_type
				t = p
				match(Token.token_type)
				p.child[1] = Factor()
				if p.child[1] == nil {
					t = nil
					p = nil
				}

			} else {

				//
				break
			}
		}
	}
	return t
}

func Factor() *Node {
	var t *Node = nil
	if Token.token_type == "TKN_LPAREN" {
		match("TKN_LPAREN")
		t = b_expresion()
		if Token.token_type != "TKN_RPAREN" {
			TokenError = Token
			msg := fmt.Sprint("Error en la linea ", TokenError.linea, " Se esperaba ')'")
			syntaxError(msg)
			for {
				if Token.token_type != "TKN_SEMICOLOM" {
					//fmt.Println(Token.token_type)
					match(Token.token_type)
				} else {
					break
				}
			}
			t = nil
		} else {
			match("TKN_RPAREN")
		}

	} else if Token.token_type == "TKN_NUM" {
		t = newExpNode(TipoExpresion[1])
		num, err := strconv.Atoi(Token.lexema)
		//t.val = num
		if err == nil {
			t.val = num
		}
		match("TKN_NUM")

	} else if Token.token_type == "TKN_IDEN" {
		t = newExpNode(TipoExpresion[2])
		t.name = Token.lexema
		match("TKN_IDEN")
	} else {
		TokenError = Token
		msg := fmt.Sprint("Error en la linea ", TokenError.linea, " Token ", TokenError.lexema, " no valido, Se esperaba una expresion valida")
		syntaxError(msg)
		if bandif == true {
			for {
				if Token.token_type != "TKN_FI" && Token.lexema != "read" && Token.lexema != ";" && Token.lexema != "write" && Token.lexema != "if" && Token.lexema != "while" {
					match(Token.token_type)
				} else {
					break
				}
			}
		} else {
			for {
				if Token.token_type != "TKN_SEMICOLOM" {
					//fmt.Println(Token.token_type)
					match(Token.token_type)
				} else {
					break
				}
			}
		}

		t = nil
	}

	return t
}
