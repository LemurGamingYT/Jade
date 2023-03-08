// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Jade

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type JadeParser struct {
	*antlr.BaseParser
}

var jadeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jadeParserInit() {
	staticData := &jadeParserStaticData
	staticData.literalNames = []string{
		"", "'('", "')'", "'{'", "'}'", "'['", "']'", "'='", "'.'", "','", "':'",
		"';'", "'->'", "'<-'", "'=>'", "'?'", "", "", "", "", "'func'", "'if'",
		"'else'", "'while'", "'import'", "'from'", "'override'", "'new'", "'class'",
		"'public'", "'private'", "'return'", "'break'", "'continue'", "'undefine'",
		"'const'", "", "", "", "", "", "'null'", "'''",
	}
	staticData.symbolicNames = []string{
		"", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LLIST", "RLIST", "ASSIGN",
		"DOT", "COMMA", "COLON", "SEMI", "LARROW", "RARROW", "ARROWASSIGN",
		"QUESTION", "NOT", "PREDTWO", "PREDONE", "COMPARATIVE", "FUNC", "IF",
		"ELSE", "WHILE", "IMPORT", "FROM", "OVERRIDE", "NEW", "CLASS", "PUBLIC",
		"PRIVATE", "RETURN", "BREAK", "CONTINUE", "UNDEFINE", "CONST", "ALL_KWS",
		"WS", "COMMENT", "MULTILINECOMMENT", "BOOL", "NUL", "APOSTROPHE", "ID",
		"INT", "FLOAT", "STRING",
	}
	staticData.ruleNames = []string{
		"parse", "block", "stmt", "iterationStmt", "functionStmt", "allStmts",
		"importStmt", "ifStmt", "whileStmt", "condition", "undefineStmt", "inheritList",
		"classdef", "args", "params", "call", "varAssign", "funcAssign", "getAttributes",
		"funcExpr", "expr", "array", "atom",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 46, 287, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 1, 0, 5, 0, 48, 8, 0, 10, 0, 12, 0, 51, 9, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 5, 1, 57, 8, 1, 10, 1, 12, 1, 60, 9, 1, 1, 1, 1,
		1, 3, 1, 64, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2,
		74, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 85,
		8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 91, 8, 6, 10, 6, 12, 6, 94, 9, 6, 1,
		6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 107,
		8, 7, 10, 7, 12, 7, 110, 9, 7, 1, 7, 1, 7, 3, 7, 114, 8, 7, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 125, 8, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 5, 11, 135, 8, 11, 10, 11, 12,
		11, 138, 9, 11, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 144, 8, 12, 1, 13, 1,
		13, 1, 13, 5, 13, 149, 8, 13, 10, 13, 12, 13, 152, 9, 13, 1, 14, 3, 14,
		155, 8, 14, 1, 14, 1, 14, 1, 14, 3, 14, 160, 8, 14, 1, 14, 5, 14, 163,
		8, 14, 10, 14, 12, 14, 166, 9, 14, 1, 15, 1, 15, 1, 15, 3, 15, 171, 8,
		15, 1, 15, 1, 15, 1, 16, 3, 16, 176, 8, 16, 1, 16, 3, 16, 179, 8, 16, 1,
		16, 3, 16, 182, 8, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 3, 17, 189, 8,
		17, 1, 17, 3, 17, 192, 8, 17, 1, 17, 3, 17, 195, 8, 17, 1, 17, 3, 17, 198,
		8, 17, 1, 17, 1, 17, 1, 17, 3, 17, 203, 8, 17, 1, 17, 1, 17, 1, 17, 3,
		17, 208, 8, 17, 1, 17, 3, 17, 211, 8, 17, 1, 17, 1, 17, 1, 17, 3, 17, 216,
		8, 17, 1, 17, 1, 17, 1, 17, 3, 17, 221, 8, 17, 1, 18, 1, 18, 1, 18, 1,
		18, 1, 18, 3, 18, 228, 8, 18, 1, 18, 5, 18, 231, 8, 18, 10, 18, 12, 18,
		234, 9, 18, 1, 19, 1, 19, 1, 19, 3, 19, 239, 8, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 3, 20, 256, 8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 5, 20, 267, 8, 20, 10, 20, 12, 20, 270, 9, 20, 1,
		21, 1, 21, 3, 21, 274, 8, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 3, 22, 285, 8, 22, 1, 22, 0, 1, 40, 23, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 0, 2, 1, 0, 32, 33, 1, 0, 29, 30, 317, 0, 49, 1, 0, 0, 0, 2, 63, 1,
		0, 0, 0, 4, 73, 1, 0, 0, 0, 6, 75, 1, 0, 0, 0, 8, 77, 1, 0, 0, 0, 10, 84,
		1, 0, 0, 0, 12, 86, 1, 0, 0, 0, 14, 98, 1, 0, 0, 0, 16, 115, 1, 0, 0, 0,
		18, 124, 1, 0, 0, 0, 20, 126, 1, 0, 0, 0, 22, 131, 1, 0, 0, 0, 24, 139,
		1, 0, 0, 0, 26, 145, 1, 0, 0, 0, 28, 154, 1, 0, 0, 0, 30, 167, 1, 0, 0,
		0, 32, 175, 1, 0, 0, 0, 34, 220, 1, 0, 0, 0, 36, 222, 1, 0, 0, 0, 38, 235,
		1, 0, 0, 0, 40, 255, 1, 0, 0, 0, 42, 271, 1, 0, 0, 0, 44, 284, 1, 0, 0,
		0, 46, 48, 3, 4, 2, 0, 47, 46, 1, 0, 0, 0, 48, 51, 1, 0, 0, 0, 49, 47,
		1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 52, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0,
		52, 53, 5, 0, 0, 1, 53, 1, 1, 0, 0, 0, 54, 58, 5, 3, 0, 0, 55, 57, 3, 4,
		2, 0, 56, 55, 1, 0, 0, 0, 57, 60, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 59,
		1, 0, 0, 0, 59, 61, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 61, 64, 5, 4, 0, 0,
		62, 64, 3, 4, 2, 0, 63, 54, 1, 0, 0, 0, 63, 62, 1, 0, 0, 0, 64, 3, 1, 0,
		0, 0, 65, 74, 3, 30, 15, 0, 66, 74, 3, 32, 16, 0, 67, 74, 3, 34, 17, 0,
		68, 74, 3, 40, 20, 0, 69, 74, 3, 24, 12, 0, 70, 74, 3, 10, 5, 0, 71, 74,
		3, 6, 3, 0, 72, 74, 3, 8, 4, 0, 73, 65, 1, 0, 0, 0, 73, 66, 1, 0, 0, 0,
		73, 67, 1, 0, 0, 0, 73, 68, 1, 0, 0, 0, 73, 69, 1, 0, 0, 0, 73, 70, 1,
		0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 72, 1, 0, 0, 0, 74, 5, 1, 0, 0, 0, 75,
		76, 7, 0, 0, 0, 76, 7, 1, 0, 0, 0, 77, 78, 5, 31, 0, 0, 78, 79, 3, 40,
		20, 0, 79, 9, 1, 0, 0, 0, 80, 85, 3, 14, 7, 0, 81, 85, 3, 16, 8, 0, 82,
		85, 3, 20, 10, 0, 83, 85, 3, 12, 6, 0, 84, 80, 1, 0, 0, 0, 84, 81, 1, 0,
		0, 0, 84, 82, 1, 0, 0, 0, 84, 83, 1, 0, 0, 0, 85, 11, 1, 0, 0, 0, 86, 87,
		5, 24, 0, 0, 87, 92, 5, 46, 0, 0, 88, 89, 5, 9, 0, 0, 89, 91, 5, 46, 0,
		0, 90, 88, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93,
		1, 0, 0, 0, 93, 95, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 95, 96, 5, 25, 0, 0,
		96, 97, 5, 46, 0, 0, 97, 13, 1, 0, 0, 0, 98, 99, 5, 21, 0, 0, 99, 100,
		3, 18, 9, 0, 100, 108, 3, 2, 1, 0, 101, 102, 5, 22, 0, 0, 102, 103, 5,
		21, 0, 0, 103, 104, 3, 18, 9, 0, 104, 105, 3, 2, 1, 0, 105, 107, 1, 0,
		0, 0, 106, 101, 1, 0, 0, 0, 107, 110, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0,
		108, 109, 1, 0, 0, 0, 109, 113, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 111,
		112, 5, 22, 0, 0, 112, 114, 3, 2, 1, 0, 113, 111, 1, 0, 0, 0, 113, 114,
		1, 0, 0, 0, 114, 15, 1, 0, 0, 0, 115, 116, 5, 23, 0, 0, 116, 117, 3, 18,
		9, 0, 117, 118, 3, 2, 1, 0, 118, 17, 1, 0, 0, 0, 119, 120, 5, 1, 0, 0,
		120, 121, 3, 40, 20, 0, 121, 122, 5, 2, 0, 0, 122, 125, 1, 0, 0, 0, 123,
		125, 3, 40, 20, 0, 124, 119, 1, 0, 0, 0, 124, 123, 1, 0, 0, 0, 125, 19,
		1, 0, 0, 0, 126, 127, 5, 34, 0, 0, 127, 128, 5, 1, 0, 0, 128, 129, 5, 43,
		0, 0, 129, 130, 5, 2, 0, 0, 130, 21, 1, 0, 0, 0, 131, 136, 5, 43, 0, 0,
		132, 133, 5, 9, 0, 0, 133, 135, 5, 43, 0, 0, 134, 132, 1, 0, 0, 0, 135,
		138, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 23, 1,
		0, 0, 0, 138, 136, 1, 0, 0, 0, 139, 140, 5, 28, 0, 0, 140, 141, 5, 43,
		0, 0, 141, 143, 5, 13, 0, 0, 142, 144, 3, 22, 11, 0, 143, 142, 1, 0, 0,
		0, 143, 144, 1, 0, 0, 0, 144, 25, 1, 0, 0, 0, 145, 150, 3, 40, 20, 0, 146,
		147, 5, 9, 0, 0, 147, 149, 3, 40, 20, 0, 148, 146, 1, 0, 0, 0, 149, 152,
		1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 27, 1, 0,
		0, 0, 152, 150, 1, 0, 0, 0, 153, 155, 5, 36, 0, 0, 154, 153, 1, 0, 0, 0,
		154, 155, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 164, 5, 43, 0, 0, 157,
		159, 5, 9, 0, 0, 158, 160, 5, 36, 0, 0, 159, 158, 1, 0, 0, 0, 159, 160,
		1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 163, 5, 43, 0, 0, 162, 157, 1, 0,
		0, 0, 163, 166, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0,
		165, 29, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 167, 168, 5, 43, 0, 0, 168,
		170, 5, 1, 0, 0, 169, 171, 3, 26, 13, 0, 170, 169, 1, 0, 0, 0, 170, 171,
		1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 173, 5, 2, 0, 0, 173, 31, 1, 0,
		0, 0, 174, 176, 7, 1, 0, 0, 175, 174, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0,
		176, 178, 1, 0, 0, 0, 177, 179, 5, 35, 0, 0, 178, 177, 1, 0, 0, 0, 178,
		179, 1, 0, 0, 0, 179, 181, 1, 0, 0, 0, 180, 182, 5, 36, 0, 0, 181, 180,
		1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 184, 5, 43,
		0, 0, 184, 185, 5, 7, 0, 0, 185, 186, 3, 40, 20, 0, 186, 33, 1, 0, 0, 0,
		187, 189, 7, 1, 0, 0, 188, 187, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189,
		194, 1, 0, 0, 0, 190, 192, 5, 36, 0, 0, 191, 190, 1, 0, 0, 0, 191, 192,
		1, 0, 0, 0, 192, 195, 1, 0, 0, 0, 193, 195, 5, 20, 0, 0, 194, 191, 1, 0,
		0, 0, 194, 193, 1, 0, 0, 0, 195, 197, 1, 0, 0, 0, 196, 198, 5, 26, 0, 0,
		197, 196, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199,
		200, 5, 43, 0, 0, 200, 202, 5, 1, 0, 0, 201, 203, 3, 28, 14, 0, 202, 201,
		1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 204, 1, 0, 0, 0, 204, 205, 5, 2,
		0, 0, 205, 221, 3, 2, 1, 0, 206, 208, 7, 1, 0, 0, 207, 206, 1, 0, 0, 0,
		207, 208, 1, 0, 0, 0, 208, 210, 1, 0, 0, 0, 209, 211, 5, 36, 0, 0, 210,
		209, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 213,
		5, 43, 0, 0, 213, 215, 5, 1, 0, 0, 214, 216, 3, 28, 14, 0, 215, 214, 1,
		0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 218, 5, 2, 0,
		0, 218, 219, 5, 14, 0, 0, 219, 221, 3, 2, 1, 0, 220, 188, 1, 0, 0, 0, 220,
		207, 1, 0, 0, 0, 221, 35, 1, 0, 0, 0, 222, 232, 3, 44, 22, 0, 223, 224,
		5, 8, 0, 0, 224, 225, 5, 43, 0, 0, 225, 227, 5, 1, 0, 0, 226, 228, 3, 26,
		13, 0, 227, 226, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0,
		229, 231, 5, 2, 0, 0, 230, 223, 1, 0, 0, 0, 231, 234, 1, 0, 0, 0, 232,
		230, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 37, 1, 0, 0, 0, 234, 232, 1,
		0, 0, 0, 235, 236, 5, 43, 0, 0, 236, 238, 5, 1, 0, 0, 237, 239, 3, 28,
		14, 0, 238, 237, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0,
		240, 241, 5, 2, 0, 0, 241, 242, 5, 14, 0, 0, 242, 243, 3, 2, 1, 0, 243,
		39, 1, 0, 0, 0, 244, 245, 6, 20, -1, 0, 245, 256, 3, 44, 22, 0, 246, 256,
		3, 30, 15, 0, 247, 248, 5, 16, 0, 0, 248, 256, 3, 40, 20, 7, 249, 256,
		3, 36, 18, 0, 250, 256, 3, 38, 19, 0, 251, 252, 5, 1, 0, 0, 252, 253, 3,
		40, 20, 0, 253, 254, 5, 2, 0, 0, 254, 256, 1, 0, 0, 0, 255, 244, 1, 0,
		0, 0, 255, 246, 1, 0, 0, 0, 255, 247, 1, 0, 0, 0, 255, 249, 1, 0, 0, 0,
		255, 250, 1, 0, 0, 0, 255, 251, 1, 0, 0, 0, 256, 268, 1, 0, 0, 0, 257,
		258, 10, 6, 0, 0, 258, 259, 5, 18, 0, 0, 259, 267, 3, 40, 20, 7, 260, 261,
		10, 5, 0, 0, 261, 262, 5, 17, 0, 0, 262, 267, 3, 40, 20, 6, 263, 264, 10,
		4, 0, 0, 264, 265, 5, 19, 0, 0, 265, 267, 3, 40, 20, 5, 266, 257, 1, 0,
		0, 0, 266, 260, 1, 0, 0, 0, 266, 263, 1, 0, 0, 0, 267, 270, 1, 0, 0, 0,
		268, 266, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269, 41, 1, 0, 0, 0, 270, 268,
		1, 0, 0, 0, 271, 273, 5, 3, 0, 0, 272, 274, 3, 26, 13, 0, 273, 272, 1,
		0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 275, 1, 0, 0, 0, 275, 276, 5, 4, 0,
		0, 276, 43, 1, 0, 0, 0, 277, 285, 3, 42, 21, 0, 278, 285, 5, 43, 0, 0,
		279, 285, 5, 44, 0, 0, 280, 285, 5, 45, 0, 0, 281, 285, 5, 46, 0, 0, 282,
		285, 5, 41, 0, 0, 283, 285, 5, 40, 0, 0, 284, 277, 1, 0, 0, 0, 284, 278,
		1, 0, 0, 0, 284, 279, 1, 0, 0, 0, 284, 280, 1, 0, 0, 0, 284, 281, 1, 0,
		0, 0, 284, 282, 1, 0, 0, 0, 284, 283, 1, 0, 0, 0, 285, 45, 1, 0, 0, 0,
		36, 49, 58, 63, 73, 84, 92, 108, 113, 124, 136, 143, 150, 154, 159, 164,
		170, 175, 178, 181, 188, 191, 194, 197, 202, 207, 210, 215, 220, 227, 232,
		238, 255, 266, 268, 273, 284,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// JadeParserInit initializes any static state used to implement JadeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJadeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JadeParserInit() {
	staticData := &jadeParserStaticData
	staticData.once.Do(jadeParserInit)
}

// NewJadeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJadeParser(input antlr.TokenStream) *JadeParser {
	JadeParserInit()
	this := new(JadeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &jadeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// JadeParser tokens.
const (
	JadeParserEOF              = antlr.TokenEOF
	JadeParserLPAREN           = 1
	JadeParserRPAREN           = 2
	JadeParserLBRACE           = 3
	JadeParserRBRACE           = 4
	JadeParserLLIST            = 5
	JadeParserRLIST            = 6
	JadeParserASSIGN           = 7
	JadeParserDOT              = 8
	JadeParserCOMMA            = 9
	JadeParserCOLON            = 10
	JadeParserSEMI             = 11
	JadeParserLARROW           = 12
	JadeParserRARROW           = 13
	JadeParserARROWASSIGN      = 14
	JadeParserQUESTION         = 15
	JadeParserNOT              = 16
	JadeParserPREDTWO          = 17
	JadeParserPREDONE          = 18
	JadeParserCOMPARATIVE      = 19
	JadeParserFUNC             = 20
	JadeParserIF               = 21
	JadeParserELSE             = 22
	JadeParserWHILE            = 23
	JadeParserIMPORT           = 24
	JadeParserFROM             = 25
	JadeParserOVERRIDE         = 26
	JadeParserNEW              = 27
	JadeParserCLASS            = 28
	JadeParserPUBLIC           = 29
	JadeParserPRIVATE          = 30
	JadeParserRETURN           = 31
	JadeParserBREAK            = 32
	JadeParserCONTINUE         = 33
	JadeParserUNDEFINE         = 34
	JadeParserCONST            = 35
	JadeParserALL_KWS          = 36
	JadeParserWS               = 37
	JadeParserCOMMENT          = 38
	JadeParserMULTILINECOMMENT = 39
	JadeParserBOOL             = 40
	JadeParserNUL              = 41
	JadeParserAPOSTROPHE       = 42
	JadeParserID               = 43
	JadeParserINT              = 44
	JadeParserFLOAT            = 45
	JadeParserSTRING           = 46
)

// JadeParser rules.
const (
	JadeParserRULE_parse         = 0
	JadeParserRULE_block         = 1
	JadeParserRULE_stmt          = 2
	JadeParserRULE_iterationStmt = 3
	JadeParserRULE_functionStmt  = 4
	JadeParserRULE_allStmts      = 5
	JadeParserRULE_importStmt    = 6
	JadeParserRULE_ifStmt        = 7
	JadeParserRULE_whileStmt     = 8
	JadeParserRULE_condition     = 9
	JadeParserRULE_undefineStmt  = 10
	JadeParserRULE_inheritList   = 11
	JadeParserRULE_classdef      = 12
	JadeParserRULE_args          = 13
	JadeParserRULE_params        = 14
	JadeParserRULE_call          = 15
	JadeParserRULE_varAssign     = 16
	JadeParserRULE_funcAssign    = 17
	JadeParserRULE_getAttributes = 18
	JadeParserRULE_funcExpr      = 19
	JadeParserRULE_expr          = 20
	JadeParserRULE_array         = 21
	JadeParserRULE_atom          = 22
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(JadeParserEOF, 0)
}

func (s *ParseContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ParseContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitParse(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) Parse() (localctx IParseContext) {
	this := p
	_ = this

	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JadeParserRULE_parse)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135377196220426) != 0 {
		{
			p.SetState(46)
			p.Stmt()
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(52)
		p.Match(JadeParserEOF)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(JadeParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(JadeParserRBRACE, 0)
}

func (s *BlockContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JadeParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(54)
			p.Match(JadeParserLBRACE)
		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135377196220426) != 0 {
			{
				p.SetState(55)
				p.Stmt()
			}

			p.SetState(60)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(61)
			p.Match(JadeParserRBRACE)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)
			p.Stmt()
		}

	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Call() ICallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}

func (s *StmtContext) VarAssign() IVarAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarAssignContext)
}

func (s *StmtContext) FuncAssign() IFuncAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncAssignContext)
}

func (s *StmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StmtContext) Classdef() IClassdefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassdefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassdefContext)
}

func (s *StmtContext) AllStmts() IAllStmtsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAllStmtsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAllStmtsContext)
}

func (s *StmtContext) IterationStmt() IIterationStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIterationStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIterationStmtContext)
}

func (s *StmtContext) FunctionStmt() IFunctionStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionStmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) Stmt() (localctx IStmtContext) {
	this := p
	_ = this

	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JadeParserRULE_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Call()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.VarAssign()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.FuncAssign()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(68)
			p.expr(0)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(69)
			p.Classdef()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(70)
			p.AllStmts()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(71)
			p.IterationStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(72)
			p.FunctionStmt()
		}

	}

	return localctx
}

// IIterationStmtContext is an interface to support dynamic dispatch.
type IIterationStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIterationStmtContext differentiates from other interfaces.
	IsIterationStmtContext()
}

type IterationStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIterationStmtContext() *IterationStmtContext {
	var p = new(IterationStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_iterationStmt
	return p
}

func (*IterationStmtContext) IsIterationStmtContext() {}

func NewIterationStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IterationStmtContext {
	var p = new(IterationStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_iterationStmt

	return p
}

func (s *IterationStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IterationStmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(JadeParserBREAK, 0)
}

func (s *IterationStmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(JadeParserCONTINUE, 0)
}

func (s *IterationStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IterationStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IterationStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitIterationStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) IterationStmt() (localctx IIterationStmtContext) {
	this := p
	_ = this

	localctx = NewIterationStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JadeParserRULE_iterationStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		_la = p.GetTokenStream().LA(1)

		if !(_la == JadeParserBREAK || _la == JadeParserCONTINUE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFunctionStmtContext is an interface to support dynamic dispatch.
type IFunctionStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionStmtContext differentiates from other interfaces.
	IsFunctionStmtContext()
}

type FunctionStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionStmtContext() *FunctionStmtContext {
	var p = new(FunctionStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_functionStmt
	return p
}

func (*FunctionStmtContext) IsFunctionStmtContext() {}

func NewFunctionStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionStmtContext {
	var p = new(FunctionStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_functionStmt

	return p
}

func (s *FunctionStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionStmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(JadeParserRETURN, 0)
}

func (s *FunctionStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunctionStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitFunctionStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) FunctionStmt() (localctx IFunctionStmtContext) {
	this := p
	_ = this

	localctx = NewFunctionStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JadeParserRULE_functionStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(JadeParserRETURN)
	}
	{
		p.SetState(78)
		p.expr(0)
	}

	return localctx
}

// IAllStmtsContext is an interface to support dynamic dispatch.
type IAllStmtsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAllStmtsContext differentiates from other interfaces.
	IsAllStmtsContext()
}

type AllStmtsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllStmtsContext() *AllStmtsContext {
	var p = new(AllStmtsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_allStmts
	return p
}

func (*AllStmtsContext) IsAllStmtsContext() {}

func NewAllStmtsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllStmtsContext {
	var p = new(AllStmtsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_allStmts

	return p
}

func (s *AllStmtsContext) GetParser() antlr.Parser { return s.parser }

func (s *AllStmtsContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *AllStmtsContext) WhileStmt() IWhileStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStmtContext)
}

func (s *AllStmtsContext) UndefineStmt() IUndefineStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUndefineStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUndefineStmtContext)
}

func (s *AllStmtsContext) ImportStmt() IImportStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportStmtContext)
}

func (s *AllStmtsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllStmtsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AllStmtsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitAllStmts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) AllStmts() (localctx IAllStmtsContext) {
	this := p
	_ = this

	localctx = NewAllStmtsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JadeParserRULE_allStmts)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(84)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JadeParserIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.IfStmt()
		}

	case JadeParserWHILE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.WhileStmt()
		}

	case JadeParserUNDEFINE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(82)
			p.UndefineStmt()
		}

	case JadeParserIMPORT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(83)
			p.ImportStmt()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IImportStmtContext is an interface to support dynamic dispatch.
type IImportStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportStmtContext differentiates from other interfaces.
	IsImportStmtContext()
}

type ImportStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStmtContext() *ImportStmtContext {
	var p = new(ImportStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_importStmt
	return p
}

func (*ImportStmtContext) IsImportStmtContext() {}

func NewImportStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStmtContext {
	var p = new(ImportStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_importStmt

	return p
}

func (s *ImportStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStmtContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(JadeParserIMPORT, 0)
}

func (s *ImportStmtContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(JadeParserSTRING)
}

func (s *ImportStmtContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(JadeParserSTRING, i)
}

func (s *ImportStmtContext) FROM() antlr.TerminalNode {
	return s.GetToken(JadeParserFROM, 0)
}

func (s *ImportStmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JadeParserCOMMA)
}

func (s *ImportStmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JadeParserCOMMA, i)
}

func (s *ImportStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitImportStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) ImportStmt() (localctx IImportStmtContext) {
	this := p
	_ = this

	localctx = NewImportStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JadeParserRULE_importStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(JadeParserIMPORT)
	}
	{
		p.SetState(87)
		p.Match(JadeParserSTRING)
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JadeParserCOMMA {
		{
			p.SetState(88)
			p.Match(JadeParserCOMMA)
		}
		{
			p.SetState(89)
			p.Match(JadeParserSTRING)
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(95)
		p.Match(JadeParserFROM)
	}
	{
		p.SetState(96)
		p.Match(JadeParserSTRING)
	}

	return localctx
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_ifStmt
	return p
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) AllIF() []antlr.TerminalNode {
	return s.GetTokens(JadeParserIF)
}

func (s *IfStmtContext) IF(i int) antlr.TerminalNode {
	return s.GetToken(JadeParserIF, i)
}

func (s *IfStmtContext) AllCondition() []IConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionContext); ok {
			len++
		}
	}

	tst := make([]IConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionContext); ok {
			tst[i] = t.(IConditionContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Condition(i int) IConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *IfStmtContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStmtContext) AllELSE() []antlr.TerminalNode {
	return s.GetTokens(JadeParserELSE)
}

func (s *IfStmtContext) ELSE(i int) antlr.TerminalNode {
	return s.GetToken(JadeParserELSE, i)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) IfStmt() (localctx IIfStmtContext) {
	this := p
	_ = this

	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JadeParserRULE_ifStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(JadeParserIF)
	}
	{
		p.SetState(99)
		p.Condition()
	}
	{
		p.SetState(100)
		p.Block()
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(101)
				p.Match(JadeParserELSE)
			}
			{
				p.SetState(102)
				p.Match(JadeParserIF)
			}
			{
				p.SetState(103)
				p.Condition()
			}
			{
				p.SetState(104)
				p.Block()
			}

		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(111)
			p.Match(JadeParserELSE)
		}
		{
			p.SetState(112)
			p.Block()
		}

	}

	return localctx
}

// IWhileStmtContext is an interface to support dynamic dispatch.
type IWhileStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileStmtContext differentiates from other interfaces.
	IsWhileStmtContext()
}

type WhileStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStmtContext() *WhileStmtContext {
	var p = new(WhileStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_whileStmt
	return p
}

func (*WhileStmtContext) IsWhileStmtContext() {}

func NewWhileStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStmtContext {
	var p = new(WhileStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_whileStmt

	return p
}

func (s *WhileStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStmtContext) WHILE() antlr.TerminalNode {
	return s.GetToken(JadeParserWHILE, 0)
}

func (s *WhileStmtContext) Condition() IConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *WhileStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitWhileStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) WhileStmt() (localctx IWhileStmtContext) {
	this := p
	_ = this

	localctx = NewWhileStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JadeParserRULE_whileStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(JadeParserWHILE)
	}
	{
		p.SetState(116)
		p.Condition()
	}
	{
		p.SetState(117)
		p.Block()
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JadeParserLPAREN, 0)
}

func (s *ConditionContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConditionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JadeParserRPAREN, 0)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) Condition() (localctx IConditionContext) {
	this := p
	_ = this

	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JadeParserRULE_condition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Match(JadeParserLPAREN)
		}
		{
			p.SetState(120)
			p.expr(0)
		}
		{
			p.SetState(121)
			p.Match(JadeParserRPAREN)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)
			p.expr(0)
		}

	}

	return localctx
}

// IUndefineStmtContext is an interface to support dynamic dispatch.
type IUndefineStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUndefineStmtContext differentiates from other interfaces.
	IsUndefineStmtContext()
}

type UndefineStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUndefineStmtContext() *UndefineStmtContext {
	var p = new(UndefineStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_undefineStmt
	return p
}

func (*UndefineStmtContext) IsUndefineStmtContext() {}

func NewUndefineStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UndefineStmtContext {
	var p = new(UndefineStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_undefineStmt

	return p
}

func (s *UndefineStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *UndefineStmtContext) UNDEFINE() antlr.TerminalNode {
	return s.GetToken(JadeParserUNDEFINE, 0)
}

func (s *UndefineStmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JadeParserLPAREN, 0)
}

func (s *UndefineStmtContext) ID() antlr.TerminalNode {
	return s.GetToken(JadeParserID, 0)
}

func (s *UndefineStmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JadeParserRPAREN, 0)
}

func (s *UndefineStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UndefineStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UndefineStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitUndefineStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) UndefineStmt() (localctx IUndefineStmtContext) {
	this := p
	_ = this

	localctx = NewUndefineStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, JadeParserRULE_undefineStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(JadeParserUNDEFINE)
	}
	{
		p.SetState(127)
		p.Match(JadeParserLPAREN)
	}
	{
		p.SetState(128)
		p.Match(JadeParserID)
	}
	{
		p.SetState(129)
		p.Match(JadeParserRPAREN)
	}

	return localctx
}

// IInheritListContext is an interface to support dynamic dispatch.
type IInheritListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInheritListContext differentiates from other interfaces.
	IsInheritListContext()
}

type InheritListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInheritListContext() *InheritListContext {
	var p = new(InheritListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_inheritList
	return p
}

func (*InheritListContext) IsInheritListContext() {}

func NewInheritListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InheritListContext {
	var p = new(InheritListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_inheritList

	return p
}

func (s *InheritListContext) GetParser() antlr.Parser { return s.parser }

func (s *InheritListContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(JadeParserID)
}

func (s *InheritListContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(JadeParserID, i)
}

func (s *InheritListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JadeParserCOMMA)
}

func (s *InheritListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JadeParserCOMMA, i)
}

func (s *InheritListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InheritListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InheritListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitInheritList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) InheritList() (localctx IInheritListContext) {
	this := p
	_ = this

	localctx = NewInheritListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, JadeParserRULE_inheritList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(JadeParserID)
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(132)
				p.Match(JadeParserCOMMA)
			}
			{
				p.SetState(133)
				p.Match(JadeParserID)
			}

		}
		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IClassdefContext is an interface to support dynamic dispatch.
type IClassdefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassdefContext differentiates from other interfaces.
	IsClassdefContext()
}

type ClassdefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassdefContext() *ClassdefContext {
	var p = new(ClassdefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_classdef
	return p
}

func (*ClassdefContext) IsClassdefContext() {}

func NewClassdefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassdefContext {
	var p = new(ClassdefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_classdef

	return p
}

func (s *ClassdefContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassdefContext) CLASS() antlr.TerminalNode {
	return s.GetToken(JadeParserCLASS, 0)
}

func (s *ClassdefContext) ID() antlr.TerminalNode {
	return s.GetToken(JadeParserID, 0)
}

func (s *ClassdefContext) RARROW() antlr.TerminalNode {
	return s.GetToken(JadeParserRARROW, 0)
}

func (s *ClassdefContext) InheritList() IInheritListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInheritListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInheritListContext)
}

func (s *ClassdefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassdefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassdefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitClassdef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) Classdef() (localctx IClassdefContext) {
	this := p
	_ = this

	localctx = NewClassdefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, JadeParserRULE_classdef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(JadeParserCLASS)
	}
	{
		p.SetState(140)
		p.Match(JadeParserID)
	}
	{
		p.SetState(141)
		p.Match(JadeParserRARROW)
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(142)
			p.InheritList()
		}

	}

	return localctx
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_args
	return p
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArgsContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArgsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JadeParserCOMMA)
}

func (s *ArgsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JadeParserCOMMA, i)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) Args() (localctx IArgsContext) {
	this := p
	_ = this

	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, JadeParserRULE_args)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.expr(0)
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JadeParserCOMMA {
		{
			p.SetState(146)
			p.Match(JadeParserCOMMA)
		}
		{
			p.SetState(147)
			p.expr(0)
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IParamsContext is an interface to support dynamic dispatch.
type IParamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamsContext differentiates from other interfaces.
	IsParamsContext()
}

type ParamsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamsContext() *ParamsContext {
	var p = new(ParamsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_params
	return p
}

func (*ParamsContext) IsParamsContext() {}

func NewParamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamsContext {
	var p = new(ParamsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_params

	return p
}

func (s *ParamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamsContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(JadeParserID)
}

func (s *ParamsContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(JadeParserID, i)
}

func (s *ParamsContext) AllALL_KWS() []antlr.TerminalNode {
	return s.GetTokens(JadeParserALL_KWS)
}

func (s *ParamsContext) ALL_KWS(i int) antlr.TerminalNode {
	return s.GetToken(JadeParserALL_KWS, i)
}

func (s *ParamsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JadeParserCOMMA)
}

func (s *ParamsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JadeParserCOMMA, i)
}

func (s *ParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitParams(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) Params() (localctx IParamsContext) {
	this := p
	_ = this

	localctx = NewParamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, JadeParserRULE_params)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JadeParserALL_KWS {
		{
			p.SetState(153)
			p.Match(JadeParserALL_KWS)
		}

	}
	{
		p.SetState(156)
		p.Match(JadeParserID)
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JadeParserCOMMA {
		{
			p.SetState(157)
			p.Match(JadeParserCOMMA)
		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JadeParserALL_KWS {
			{
				p.SetState(158)
				p.Match(JadeParserALL_KWS)
			}

		}
		{
			p.SetState(161)
			p.Match(JadeParserID)
		}

		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICallContext is an interface to support dynamic dispatch.
type ICallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCallContext differentiates from other interfaces.
	IsCallContext()
}

type CallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallContext() *CallContext {
	var p = new(CallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_call
	return p
}

func (*CallContext) IsCallContext() {}

func NewCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallContext {
	var p = new(CallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_call

	return p
}

func (s *CallContext) GetParser() antlr.Parser { return s.parser }

func (s *CallContext) ID() antlr.TerminalNode {
	return s.GetToken(JadeParserID, 0)
}

func (s *CallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JadeParserLPAREN, 0)
}

func (s *CallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JadeParserRPAREN, 0)
}

func (s *CallContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) Call() (localctx ICallContext) {
	this := p
	_ = this

	localctx = NewCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, JadeParserRULE_call)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(JadeParserID)
	}
	{
		p.SetState(168)
		p.Match(JadeParserLPAREN)
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135239930281994) != 0 {
		{
			p.SetState(169)
			p.Args()
		}

	}
	{
		p.SetState(172)
		p.Match(JadeParserRPAREN)
	}

	return localctx
}

// IVarAssignContext is an interface to support dynamic dispatch.
type IVarAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarAssignContext differentiates from other interfaces.
	IsVarAssignContext()
}

type VarAssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarAssignContext() *VarAssignContext {
	var p = new(VarAssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_varAssign
	return p
}

func (*VarAssignContext) IsVarAssignContext() {}

func NewVarAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarAssignContext {
	var p = new(VarAssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_varAssign

	return p
}

func (s *VarAssignContext) GetParser() antlr.Parser { return s.parser }

func (s *VarAssignContext) ID() antlr.TerminalNode {
	return s.GetToken(JadeParserID, 0)
}

func (s *VarAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(JadeParserASSIGN, 0)
}

func (s *VarAssignContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VarAssignContext) CONST() antlr.TerminalNode {
	return s.GetToken(JadeParserCONST, 0)
}

func (s *VarAssignContext) ALL_KWS() antlr.TerminalNode {
	return s.GetToken(JadeParserALL_KWS, 0)
}

func (s *VarAssignContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(JadeParserPUBLIC, 0)
}

func (s *VarAssignContext) PRIVATE() antlr.TerminalNode {
	return s.GetToken(JadeParserPRIVATE, 0)
}

func (s *VarAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarAssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitVarAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) VarAssign() (localctx IVarAssignContext) {
	this := p
	_ = this

	localctx = NewVarAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, JadeParserRULE_varAssign)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JadeParserPUBLIC || _la == JadeParserPRIVATE {
		{
			p.SetState(174)
			_la = p.GetTokenStream().LA(1)

			if !(_la == JadeParserPUBLIC || _la == JadeParserPRIVATE) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JadeParserCONST {
		{
			p.SetState(177)
			p.Match(JadeParserCONST)
		}

	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JadeParserALL_KWS {
		{
			p.SetState(180)
			p.Match(JadeParserALL_KWS)
		}

	}
	{
		p.SetState(183)
		p.Match(JadeParserID)
	}
	{
		p.SetState(184)
		p.Match(JadeParserASSIGN)
	}
	{
		p.SetState(185)
		p.expr(0)
	}

	return localctx
}

// IFuncAssignContext is an interface to support dynamic dispatch.
type IFuncAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncAssignContext differentiates from other interfaces.
	IsFuncAssignContext()
}

type FuncAssignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncAssignContext() *FuncAssignContext {
	var p = new(FuncAssignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_funcAssign
	return p
}

func (*FuncAssignContext) IsFuncAssignContext() {}

func NewFuncAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncAssignContext {
	var p = new(FuncAssignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_funcAssign

	return p
}

func (s *FuncAssignContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncAssignContext) ID() antlr.TerminalNode {
	return s.GetToken(JadeParserID, 0)
}

func (s *FuncAssignContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JadeParserLPAREN, 0)
}

func (s *FuncAssignContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JadeParserRPAREN, 0)
}

func (s *FuncAssignContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncAssignContext) FUNC() antlr.TerminalNode {
	return s.GetToken(JadeParserFUNC, 0)
}

func (s *FuncAssignContext) OVERRIDE() antlr.TerminalNode {
	return s.GetToken(JadeParserOVERRIDE, 0)
}

func (s *FuncAssignContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *FuncAssignContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(JadeParserPUBLIC, 0)
}

func (s *FuncAssignContext) PRIVATE() antlr.TerminalNode {
	return s.GetToken(JadeParserPRIVATE, 0)
}

func (s *FuncAssignContext) ALL_KWS() antlr.TerminalNode {
	return s.GetToken(JadeParserALL_KWS, 0)
}

func (s *FuncAssignContext) ARROWASSIGN() antlr.TerminalNode {
	return s.GetToken(JadeParserARROWASSIGN, 0)
}

func (s *FuncAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncAssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitFuncAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) FuncAssign() (localctx IFuncAssignContext) {
	this := p
	_ = this

	localctx = NewFuncAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, JadeParserRULE_funcAssign)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JadeParserPUBLIC || _la == JadeParserPRIVATE {
			{
				p.SetState(187)
				_la = p.GetTokenStream().LA(1)

				if !(_la == JadeParserPUBLIC || _la == JadeParserPRIVATE) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(194)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case JadeParserOVERRIDE, JadeParserALL_KWS, JadeParserID:
			p.SetState(191)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == JadeParserALL_KWS {
				{
					p.SetState(190)
					p.Match(JadeParserALL_KWS)
				}

			}

		case JadeParserFUNC:
			{
				p.SetState(193)
				p.Match(JadeParserFUNC)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JadeParserOVERRIDE {
			{
				p.SetState(196)
				p.Match(JadeParserOVERRIDE)
			}

		}
		{
			p.SetState(199)
			p.Match(JadeParserID)
		}
		{
			p.SetState(200)
			p.Match(JadeParserLPAREN)
		}
		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JadeParserALL_KWS || _la == JadeParserID {
			{
				p.SetState(201)
				p.Params()
			}

		}
		{
			p.SetState(204)
			p.Match(JadeParserRPAREN)
		}
		{
			p.SetState(205)
			p.Block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JadeParserPUBLIC || _la == JadeParserPRIVATE {
			{
				p.SetState(206)
				_la = p.GetTokenStream().LA(1)

				if !(_la == JadeParserPUBLIC || _la == JadeParserPRIVATE) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JadeParserALL_KWS {
			{
				p.SetState(209)
				p.Match(JadeParserALL_KWS)
			}

		}
		{
			p.SetState(212)
			p.Match(JadeParserID)
		}
		{
			p.SetState(213)
			p.Match(JadeParserLPAREN)
		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JadeParserALL_KWS || _la == JadeParserID {
			{
				p.SetState(214)
				p.Params()
			}

		}
		{
			p.SetState(217)
			p.Match(JadeParserRPAREN)
		}
		{
			p.SetState(218)
			p.Match(JadeParserARROWASSIGN)
		}
		{
			p.SetState(219)
			p.Block()
		}

	}

	return localctx
}

// IGetAttributesContext is an interface to support dynamic dispatch.
type IGetAttributesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGetAttributesContext differentiates from other interfaces.
	IsGetAttributesContext()
}

type GetAttributesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGetAttributesContext() *GetAttributesContext {
	var p = new(GetAttributesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_getAttributes
	return p
}

func (*GetAttributesContext) IsGetAttributesContext() {}

func NewGetAttributesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GetAttributesContext {
	var p = new(GetAttributesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_getAttributes

	return p
}

func (s *GetAttributesContext) GetParser() antlr.Parser { return s.parser }

func (s *GetAttributesContext) Atom() IAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *GetAttributesContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(JadeParserDOT)
}

func (s *GetAttributesContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(JadeParserDOT, i)
}

func (s *GetAttributesContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(JadeParserID)
}

func (s *GetAttributesContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(JadeParserID, i)
}

func (s *GetAttributesContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(JadeParserLPAREN)
}

func (s *GetAttributesContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(JadeParserLPAREN, i)
}

func (s *GetAttributesContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(JadeParserRPAREN)
}

func (s *GetAttributesContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(JadeParserRPAREN, i)
}

func (s *GetAttributesContext) AllArgs() []IArgsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgsContext); ok {
			len++
		}
	}

	tst := make([]IArgsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgsContext); ok {
			tst[i] = t.(IArgsContext)
			i++
		}
	}

	return tst
}

func (s *GetAttributesContext) Args(i int) IArgsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *GetAttributesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetAttributesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GetAttributesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitGetAttributes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) GetAttributes() (localctx IGetAttributesContext) {
	this := p
	_ = this

	localctx = NewGetAttributesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, JadeParserRULE_getAttributes)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.Atom()
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(223)
				p.Match(JadeParserDOT)
			}
			{
				p.SetState(224)
				p.Match(JadeParserID)
			}
			{
				p.SetState(225)
				p.Match(JadeParserLPAREN)
			}
			p.SetState(227)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135239930281994) != 0 {
				{
					p.SetState(226)
					p.Args()
				}

			}
			{
				p.SetState(229)
				p.Match(JadeParserRPAREN)
			}

		}
		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
	}

	return localctx
}

// IFuncExprContext is an interface to support dynamic dispatch.
type IFuncExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncExprContext differentiates from other interfaces.
	IsFuncExprContext()
}

type FuncExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncExprContext() *FuncExprContext {
	var p = new(FuncExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_funcExpr
	return p
}

func (*FuncExprContext) IsFuncExprContext() {}

func NewFuncExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncExprContext {
	var p = new(FuncExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_funcExpr

	return p
}

func (s *FuncExprContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncExprContext) ID() antlr.TerminalNode {
	return s.GetToken(JadeParserID, 0)
}

func (s *FuncExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JadeParserLPAREN, 0)
}

func (s *FuncExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JadeParserRPAREN, 0)
}

func (s *FuncExprContext) ARROWASSIGN() antlr.TerminalNode {
	return s.GetToken(JadeParserARROWASSIGN, 0)
}

func (s *FuncExprContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncExprContext) Params() IParamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamsContext)
}

func (s *FuncExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitFuncExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) FuncExpr() (localctx IFuncExprContext) {
	this := p
	_ = this

	localctx = NewFuncExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, JadeParserRULE_funcExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Match(JadeParserID)
	}
	{
		p.SetState(236)
		p.Match(JadeParserLPAREN)
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JadeParserALL_KWS || _la == JadeParserID {
		{
			p.SetState(237)
			p.Params()
		}

	}
	{
		p.SetState(240)
		p.Match(JadeParserRPAREN)
	}
	{
		p.SetState(241)
		p.Match(JadeParserARROWASSIGN)
	}
	{
		p.SetState(242)
		p.Block()
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) Atom() IAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ExprContext) Call() ICallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(JadeParserNOT, 0)
}

func (s *ExprContext) GetAttributes() IGetAttributesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGetAttributesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGetAttributesContext)
}

func (s *ExprContext) FuncExpr() IFuncExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncExprContext)
}

func (s *ExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(JadeParserLPAREN, 0)
}

func (s *ExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(JadeParserRPAREN, 0)
}

func (s *ExprContext) PREDONE() antlr.TerminalNode {
	return s.GetToken(JadeParserPREDONE, 0)
}

func (s *ExprContext) PREDTWO() antlr.TerminalNode {
	return s.GetToken(JadeParserPREDTWO, 0)
}

func (s *ExprContext) COMPARATIVE() antlr.TerminalNode {
	return s.GetToken(JadeParserCOMPARATIVE, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *JadeParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 40
	p.EnterRecursionRule(localctx, 40, JadeParserRULE_expr, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(245)
			p.Atom()
		}

	case 2:
		{
			p.SetState(246)
			p.Call()
		}

	case 3:
		{
			p.SetState(247)

			var _m = p.Match(JadeParserNOT)

			localctx.(*ExprContext).op = _m
		}
		{
			p.SetState(248)
			p.expr(7)
		}

	case 4:
		{
			p.SetState(249)
			p.GetAttributes()
		}

	case 5:
		{
			p.SetState(250)
			p.FuncExpr()
		}

	case 6:
		{
			p.SetState(251)
			p.Match(JadeParserLPAREN)
		}
		{
			p.SetState(252)
			p.expr(0)
		}
		{
			p.SetState(253)
			p.Match(JadeParserRPAREN)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(266)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JadeParserRULE_expr)
				p.SetState(257)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(258)

					var _m = p.Match(JadeParserPREDONE)

					localctx.(*ExprContext).op = _m
				}
				{
					p.SetState(259)
					p.expr(7)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JadeParserRULE_expr)
				p.SetState(260)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(261)

					var _m = p.Match(JadeParserPREDTWO)

					localctx.(*ExprContext).op = _m
				}
				{
					p.SetState(262)
					p.expr(6)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JadeParserRULE_expr)
				p.SetState(263)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(264)

					var _m = p.Match(JadeParserCOMPARATIVE)

					localctx.(*ExprContext).op = _m
				}
				{
					p.SetState(265)
					p.expr(5)
				}

			}

		}
		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(JadeParserLBRACE, 0)
}

func (s *ArrayContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(JadeParserRBRACE, 0)
}

func (s *ArrayContext) Args() IArgsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) Array() (localctx IArrayContext) {
	this := p
	_ = this

	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, JadeParserRULE_array)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(271)
		p.Match(JadeParserLBRACE)
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&135239930281994) != 0 {
		{
			p.SetState(272)
			p.Args()
		}

	}
	{
		p.SetState(275)
		p.Match(JadeParserRBRACE)
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JadeParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JadeParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) Array() IArrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *AtomContext) ID() antlr.TerminalNode {
	return s.GetToken(JadeParserID, 0)
}

func (s *AtomContext) INT() antlr.TerminalNode {
	return s.GetToken(JadeParserINT, 0)
}

func (s *AtomContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(JadeParserFLOAT, 0)
}

func (s *AtomContext) STRING() antlr.TerminalNode {
	return s.GetToken(JadeParserSTRING, 0)
}

func (s *AtomContext) NUL() antlr.TerminalNode {
	return s.GetToken(JadeParserNUL, 0)
}

func (s *AtomContext) BOOL() antlr.TerminalNode {
	return s.GetToken(JadeParserBOOL, 0)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JadeVisitor:
		return t.VisitAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JadeParser) Atom() (localctx IAtomContext) {
	this := p
	_ = this

	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, JadeParserRULE_atom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(284)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JadeParserLBRACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(277)
			p.Array()
		}

	case JadeParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(278)
			p.Match(JadeParserID)
		}

	case JadeParserINT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(279)
			p.Match(JadeParserINT)
		}

	case JadeParserFLOAT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(280)
			p.Match(JadeParserFLOAT)
		}

	case JadeParserSTRING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(281)
			p.Match(JadeParserSTRING)
		}

	case JadeParserNUL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(282)
			p.Match(JadeParserNUL)
		}

	case JadeParserBOOL:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(283)
			p.Match(JadeParserBOOL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *JadeParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 20:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *JadeParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
