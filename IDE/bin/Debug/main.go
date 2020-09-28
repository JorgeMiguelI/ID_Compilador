package main

import (
	"fmt"
	"os"
)

func main() {
	var SyntaxTree *Node
	Scanner()
	SyntaxTree = Analyze()
	ImprimeArbol(SyntaxTree)
}

var espacios int = 0
var band int = 0

func printSpaces() {
	f, err := os.OpenFile("salidaSintactico.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	for i := 0; i < espacios; i++ {
		if band > 0 {
			fmt.Fprintf(f, "|___")
			fmt.Printf("|___")
		}

	}
}

func ImprimeArbol(tree *Node) {
	f, err := os.OpenFile("salidaSintactico.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	//band := 0
	espacios += 1
	for {
		if tree != nil {
			printSpaces()
			band += 1
			if tree.TipoNodo == TipoNodo[0] {
				switch tree.TipoSentencia {
				case TipoSententencia[0]:
					fmt.Printf("If \n")
					fmt.Fprintf(f, "If \n")
					break
				case TipoSententencia[1]:
					fmt.Printf("While \n")
					fmt.Fprintf(f, "while \n")
					break
				case TipoSententencia[2]:
					fmt.Printf("Assign to: %s \n", tree.name)
					fmt.Fprintf(f, "Assign to: %s \n", tree.name)
					break
				case TipoSententencia[3]:
					fmt.Printf("Read: %s \n", tree.name)
					fmt.Fprintf(f, "Read: %s \n", tree.name)
					break
				case TipoSententencia[4]:
					fmt.Printf("Write \n")
					fmt.Fprintf(f, "Write \n")
					break
				case TipoSententencia[5]:
					fmt.Printf("doUntil \n")
					fmt.Fprintf(f, "doUntil \n")
					break
				default:
					fmt.Printf("No se reconoce el tipo de nodo \n")
					break
				}
			} else if tree.TipoNodo == TipoNodo[1] {
				switch tree.TipoExpresion {
				case TipoExpresion[0]:
					fmt.Printf("Op: ")
					fmt.Fprintf(f, "Op: ")
					printToken(tree.op)
					break
				case TipoExpresion[1]:
					fmt.Printf("Const: %d \n", tree.val)
					fmt.Fprintf(f, "Const: %d \n", tree.val)
					break
				case TipoExpresion[2]:
					fmt.Printf("Id: %s \n", tree.name)
					fmt.Fprintf(f, "Id: %s \n", tree.name)
					break
				case TipoExpresion[3]:
					fmt.Printf("OpLogico: ")
					fmt.Fprintf(f, "OpLogico: ")
					printToken(tree.op)
					break
				case TipoExpresion[4]:
					fmt.Printf("Boolean: %s \n", tree.name)
					fmt.Fprintf(f, "Boolean: %s \n", tree.name)
					break
				default:
					fmt.Printf("No se reconoce el tipo de expresion \n")
					break
				}
			} else if tree.TipoNodo == TipoNodo[2] {
				switch tree.TipoDato {
				case TipoDatoDec[0]:
					fmt.Printf("int \n")
					fmt.Fprintf(f, "int: \n")
					break
				case TipoDatoDec[1]:
					fmt.Printf("float \n")
					fmt.Fprintf(f, "float: \n")
					break
				case TipoDatoDec[2]:
					fmt.Printf("bool \n")
					fmt.Fprintf(f, "bool: \n")
					break
				default:
					fmt.Printf("Error en la linea %d, Tipo de dato no reconocido \n", Token.linea)
					break
				}
			} else if tree.TipoNodo == TipoNodo[3] {
				fmt.Printf("%s \n", tree.name)
				fmt.Fprint(f, "program \n")
			} else {
				fmt.Printf("Tipo de nodo Desconocido \n")
			}
			for i := 0; i < 3; i++ {
				ImprimeArbol(tree.child[i])
			}
			tree = tree.sibling
		} else {
			break
		}
	}
	//fmt.Println(espacios)
	espacios -= 1
}

func printToken(Token string) {
	f, err := os.OpenFile("salidaSintactico.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	switch Token {
	case "TKN_IF":
	case "TKN_THEN":
	case "TKN_ELSE":
	case "TKN_END":
	case "TKN_REPEAT":
	case "TKN_UNTIL":
	case "TKN_READ":
	case "TKN_WRITE":
		fmt.Printf("Palabra Reservada: %s \n", Token)
		break
	case "TKN_ASIGN":
		fmt.Printf("= \n")
		fmt.Fprintf(f, "= \n")
		break
	case "TKN_MENOR":
		fmt.Printf("< \n")
		fmt.Fprintf(f, "< \n")
		break
	case "TKN_LPAREN":
		fmt.Printf("( \n")
		fmt.Fprintf(f, "( \n")
		break
	case "TKN_RPAREN":
		fmt.Printf(") \n")
		break
	case "TKN_SEMICOLOM":
		fmt.Printf("; \n")
		break
	case "TKN_ADD":
		fmt.Printf("+ \n")
		fmt.Fprintf(f, "+ \n")
		break
	case "TKN_MINUS":
		fmt.Printf("- \n")
		fmt.Fprintf(f, "- \n")
		break
	case "TKN_MUL":
		fmt.Printf("* \n")
		fmt.Fprintf(f, "* \n")
		break
	case "TKN_DIV":
		fmt.Printf("/ \n")
		fmt.Fprintf(f, "/ \n")
		break
	case "TKN_MENOR_IGUAL":
		fmt.Printf("<= \n")
		fmt.Fprintf(f, "<= \n")
		break
	case "TKN_MAYOR":
		fmt.Printf("> \n")
		fmt.Fprintf(f, "> \n")
		break
	case "TKN_MAYOR_IGUAL":
		fmt.Printf(">= \n")
		fmt.Fprintf(f, ">= \n")
		break
	case "TKN_COMP_IGUALDAD":
		fmt.Printf("== \n")
		fmt.Fprintf(f, "== \n")
		break
	case "TKN_DIFERENTE":
		fmt.Printf("!= \n")
		fmt.Fprintf(f, "!= \n")
		break
	case "TKN_EOF":
		fmt.Printf("EOF \n")
		break
	case "TKN_NOT":
		fmt.Printf("NOT \n")
		fmt.Fprintf(f, "NOT \n")
		break
	case "TKN_OR":
		fmt.Printf("OR \n")
		fmt.Fprintf(f, "OR \n")
		break
	case "TKN_AND":
		fmt.Printf("AND \n")
		fmt.Fprintf(f, "AND \n")
		break
	default:
		fmt.Print("No se reconoce el Token")

	}
}
