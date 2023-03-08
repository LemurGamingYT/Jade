// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type JadeLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var jadelexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jadelexerLexerInit() {
	staticData := &jadelexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"LPAREN", "RPAREN", "LBRACE", "RBRACE", "LLIST", "RLIST", "ASSIGN",
		"DOT", "COMMA", "COLON", "SEMI", "LARROW", "RARROW", "ARROWASSIGN",
		"QUESTION", "NOT", "PREDTWO", "PREDONE", "COMPARATIVE", "FUNC", "IF",
		"ELSE", "WHILE", "IMPORT", "FROM", "OVERRIDE", "NEW", "CLASS", "PUBLIC",
		"PRIVATE", "RETURN", "BREAK", "CONTINUE", "UNDEFINE", "CONST", "ALL_KWS",
		"WS", "COMMENT", "MULTILINECOMMENT", "BOOL", "NUL", "APOSTROPHE", "ID",
		"INT", "FLOAT", "STRING",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 46, 382, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 1, 15, 3, 15, 131, 8, 15, 1, 16, 1, 16, 1, 16, 3, 16, 136, 8, 16, 1,
		17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 3, 18, 160, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 282, 8,
		35, 1, 36, 4, 36, 285, 8, 36, 11, 36, 12, 36, 286, 1, 36, 1, 36, 1, 37,
		1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 297, 8, 37, 1, 37, 5, 37, 300,
		8, 37, 10, 37, 12, 37, 303, 9, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 1,
		38, 5, 38, 311, 8, 38, 10, 38, 12, 38, 314, 9, 38, 1, 38, 1, 38, 1, 38,
		1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 3, 39, 330, 8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41,
		1, 42, 1, 42, 5, 42, 341, 8, 42, 10, 42, 12, 42, 344, 9, 42, 1, 43, 3,
		43, 347, 8, 43, 1, 43, 4, 43, 350, 8, 43, 11, 43, 12, 43, 351, 1, 44, 3,
		44, 355, 8, 44, 1, 44, 5, 44, 358, 8, 44, 10, 44, 12, 44, 361, 9, 44, 1,
		44, 1, 44, 4, 44, 365, 8, 44, 11, 44, 12, 44, 366, 1, 45, 1, 45, 3, 45,
		371, 8, 45, 1, 45, 5, 45, 374, 8, 45, 10, 45, 12, 45, 377, 9, 45, 1, 45,
		1, 45, 3, 45, 381, 8, 45, 2, 312, 375, 0, 46, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23,
		47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32,
		65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41,
		83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 1, 0, 8, 2, 0, 43, 43, 45, 45,
		3, 0, 37, 37, 42, 42, 47, 47, 2, 0, 60, 60, 62, 62, 3, 0, 9, 10, 13, 13,
		32, 32, 2, 0, 10, 10, 13, 13, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48,
		57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57, 410, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0,
		0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0,
		0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1,
		0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57,
		1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0,
		65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0,
		0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0,
		0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0,
		0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 1, 93, 1, 0, 0, 0, 3, 95, 1,
		0, 0, 0, 5, 97, 1, 0, 0, 0, 7, 99, 1, 0, 0, 0, 9, 101, 1, 0, 0, 0, 11,
		103, 1, 0, 0, 0, 13, 105, 1, 0, 0, 0, 15, 107, 1, 0, 0, 0, 17, 109, 1,
		0, 0, 0, 19, 111, 1, 0, 0, 0, 21, 113, 1, 0, 0, 0, 23, 115, 1, 0, 0, 0,
		25, 118, 1, 0, 0, 0, 27, 121, 1, 0, 0, 0, 29, 124, 1, 0, 0, 0, 31, 130,
		1, 0, 0, 0, 33, 135, 1, 0, 0, 0, 35, 137, 1, 0, 0, 0, 37, 159, 1, 0, 0,
		0, 39, 161, 1, 0, 0, 0, 41, 166, 1, 0, 0, 0, 43, 169, 1, 0, 0, 0, 45, 174,
		1, 0, 0, 0, 47, 180, 1, 0, 0, 0, 49, 187, 1, 0, 0, 0, 51, 192, 1, 0, 0,
		0, 53, 201, 1, 0, 0, 0, 55, 205, 1, 0, 0, 0, 57, 211, 1, 0, 0, 0, 59, 218,
		1, 0, 0, 0, 61, 226, 1, 0, 0, 0, 63, 233, 1, 0, 0, 0, 65, 239, 1, 0, 0,
		0, 67, 248, 1, 0, 0, 0, 69, 257, 1, 0, 0, 0, 71, 281, 1, 0, 0, 0, 73, 284,
		1, 0, 0, 0, 75, 296, 1, 0, 0, 0, 77, 306, 1, 0, 0, 0, 79, 329, 1, 0, 0,
		0, 81, 331, 1, 0, 0, 0, 83, 336, 1, 0, 0, 0, 85, 338, 1, 0, 0, 0, 87, 346,
		1, 0, 0, 0, 89, 354, 1, 0, 0, 0, 91, 370, 1, 0, 0, 0, 93, 94, 5, 40, 0,
		0, 94, 2, 1, 0, 0, 0, 95, 96, 5, 41, 0, 0, 96, 4, 1, 0, 0, 0, 97, 98, 5,
		123, 0, 0, 98, 6, 1, 0, 0, 0, 99, 100, 5, 125, 0, 0, 100, 8, 1, 0, 0, 0,
		101, 102, 5, 91, 0, 0, 102, 10, 1, 0, 0, 0, 103, 104, 5, 93, 0, 0, 104,
		12, 1, 0, 0, 0, 105, 106, 5, 61, 0, 0, 106, 14, 1, 0, 0, 0, 107, 108, 5,
		46, 0, 0, 108, 16, 1, 0, 0, 0, 109, 110, 5, 44, 0, 0, 110, 18, 1, 0, 0,
		0, 111, 112, 5, 58, 0, 0, 112, 20, 1, 0, 0, 0, 113, 114, 5, 59, 0, 0, 114,
		22, 1, 0, 0, 0, 115, 116, 5, 45, 0, 0, 116, 117, 5, 62, 0, 0, 117, 24,
		1, 0, 0, 0, 118, 119, 5, 60, 0, 0, 119, 120, 5, 45, 0, 0, 120, 26, 1, 0,
		0, 0, 121, 122, 5, 61, 0, 0, 122, 123, 5, 62, 0, 0, 123, 28, 1, 0, 0, 0,
		124, 125, 5, 63, 0, 0, 125, 30, 1, 0, 0, 0, 126, 131, 5, 33, 0, 0, 127,
		128, 5, 110, 0, 0, 128, 129, 5, 111, 0, 0, 129, 131, 5, 116, 0, 0, 130,
		126, 1, 0, 0, 0, 130, 127, 1, 0, 0, 0, 131, 32, 1, 0, 0, 0, 132, 136, 7,
		0, 0, 0, 133, 134, 5, 42, 0, 0, 134, 136, 5, 42, 0, 0, 135, 132, 1, 0,
		0, 0, 135, 133, 1, 0, 0, 0, 136, 34, 1, 0, 0, 0, 137, 138, 7, 1, 0, 0,
		138, 36, 1, 0, 0, 0, 139, 140, 5, 61, 0, 0, 140, 160, 5, 61, 0, 0, 141,
		142, 5, 33, 0, 0, 142, 160, 5, 61, 0, 0, 143, 144, 5, 126, 0, 0, 144, 160,
		5, 61, 0, 0, 145, 160, 7, 2, 0, 0, 146, 147, 5, 60, 0, 0, 147, 160, 5,
		61, 0, 0, 148, 149, 5, 62, 0, 0, 149, 160, 5, 61, 0, 0, 150, 151, 5, 97,
		0, 0, 151, 152, 5, 110, 0, 0, 152, 160, 5, 100, 0, 0, 153, 154, 5, 38,
		0, 0, 154, 160, 5, 38, 0, 0, 155, 156, 5, 111, 0, 0, 156, 160, 5, 114,
		0, 0, 157, 158, 5, 124, 0, 0, 158, 160, 5, 124, 0, 0, 159, 139, 1, 0, 0,
		0, 159, 141, 1, 0, 0, 0, 159, 143, 1, 0, 0, 0, 159, 145, 1, 0, 0, 0, 159,
		146, 1, 0, 0, 0, 159, 148, 1, 0, 0, 0, 159, 150, 1, 0, 0, 0, 159, 153,
		1, 0, 0, 0, 159, 155, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160, 38, 1, 0,
		0, 0, 161, 162, 5, 102, 0, 0, 162, 163, 5, 117, 0, 0, 163, 164, 5, 110,
		0, 0, 164, 165, 5, 99, 0, 0, 165, 40, 1, 0, 0, 0, 166, 167, 5, 105, 0,
		0, 167, 168, 5, 102, 0, 0, 168, 42, 1, 0, 0, 0, 169, 170, 5, 101, 0, 0,
		170, 171, 5, 108, 0, 0, 171, 172, 5, 115, 0, 0, 172, 173, 5, 101, 0, 0,
		173, 44, 1, 0, 0, 0, 174, 175, 5, 119, 0, 0, 175, 176, 5, 104, 0, 0, 176,
		177, 5, 105, 0, 0, 177, 178, 5, 108, 0, 0, 178, 179, 5, 101, 0, 0, 179,
		46, 1, 0, 0, 0, 180, 181, 5, 105, 0, 0, 181, 182, 5, 109, 0, 0, 182, 183,
		5, 112, 0, 0, 183, 184, 5, 111, 0, 0, 184, 185, 5, 114, 0, 0, 185, 186,
		5, 116, 0, 0, 186, 48, 1, 0, 0, 0, 187, 188, 5, 102, 0, 0, 188, 189, 5,
		114, 0, 0, 189, 190, 5, 111, 0, 0, 190, 191, 5, 109, 0, 0, 191, 50, 1,
		0, 0, 0, 192, 193, 5, 111, 0, 0, 193, 194, 5, 118, 0, 0, 194, 195, 5, 101,
		0, 0, 195, 196, 5, 114, 0, 0, 196, 197, 5, 114, 0, 0, 197, 198, 5, 105,
		0, 0, 198, 199, 5, 100, 0, 0, 199, 200, 5, 101, 0, 0, 200, 52, 1, 0, 0,
		0, 201, 202, 5, 110, 0, 0, 202, 203, 5, 101, 0, 0, 203, 204, 5, 119, 0,
		0, 204, 54, 1, 0, 0, 0, 205, 206, 5, 99, 0, 0, 206, 207, 5, 108, 0, 0,
		207, 208, 5, 97, 0, 0, 208, 209, 5, 115, 0, 0, 209, 210, 5, 115, 0, 0,
		210, 56, 1, 0, 0, 0, 211, 212, 5, 112, 0, 0, 212, 213, 5, 117, 0, 0, 213,
		214, 5, 98, 0, 0, 214, 215, 5, 108, 0, 0, 215, 216, 5, 105, 0, 0, 216,
		217, 5, 99, 0, 0, 217, 58, 1, 0, 0, 0, 218, 219, 5, 112, 0, 0, 219, 220,
		5, 114, 0, 0, 220, 221, 5, 105, 0, 0, 221, 222, 5, 118, 0, 0, 222, 223,
		5, 97, 0, 0, 223, 224, 5, 116, 0, 0, 224, 225, 5, 101, 0, 0, 225, 60, 1,
		0, 0, 0, 226, 227, 5, 114, 0, 0, 227, 228, 5, 101, 0, 0, 228, 229, 5, 116,
		0, 0, 229, 230, 5, 117, 0, 0, 230, 231, 5, 114, 0, 0, 231, 232, 5, 110,
		0, 0, 232, 62, 1, 0, 0, 0, 233, 234, 5, 98, 0, 0, 234, 235, 5, 114, 0,
		0, 235, 236, 5, 101, 0, 0, 236, 237, 5, 97, 0, 0, 237, 238, 5, 107, 0,
		0, 238, 64, 1, 0, 0, 0, 239, 240, 5, 99, 0, 0, 240, 241, 5, 111, 0, 0,
		241, 242, 5, 110, 0, 0, 242, 243, 5, 116, 0, 0, 243, 244, 5, 105, 0, 0,
		244, 245, 5, 110, 0, 0, 245, 246, 5, 117, 0, 0, 246, 247, 5, 101, 0, 0,
		247, 66, 1, 0, 0, 0, 248, 249, 5, 117, 0, 0, 249, 250, 5, 110, 0, 0, 250,
		251, 5, 100, 0, 0, 251, 252, 5, 101, 0, 0, 252, 253, 5, 102, 0, 0, 253,
		254, 5, 105, 0, 0, 254, 255, 5, 110, 0, 0, 255, 256, 5, 101, 0, 0, 256,
		68, 1, 0, 0, 0, 257, 258, 5, 99, 0, 0, 258, 259, 5, 111, 0, 0, 259, 260,
		5, 110, 0, 0, 260, 261, 5, 115, 0, 0, 261, 262, 5, 116, 0, 0, 262, 70,
		1, 0, 0, 0, 263, 264, 5, 115, 0, 0, 264, 265, 5, 116, 0, 0, 265, 266, 5,
		114, 0, 0, 266, 267, 5, 105, 0, 0, 267, 268, 5, 110, 0, 0, 268, 282, 5,
		103, 0, 0, 269, 270, 5, 105, 0, 0, 270, 271, 5, 110, 0, 0, 271, 282, 5,
		116, 0, 0, 272, 273, 5, 102, 0, 0, 273, 274, 5, 108, 0, 0, 274, 275, 5,
		111, 0, 0, 275, 276, 5, 97, 0, 0, 276, 282, 5, 116, 0, 0, 277, 278, 5,
		98, 0, 0, 278, 279, 5, 111, 0, 0, 279, 280, 5, 111, 0, 0, 280, 282, 5,
		108, 0, 0, 281, 263, 1, 0, 0, 0, 281, 269, 1, 0, 0, 0, 281, 272, 1, 0,
		0, 0, 281, 277, 1, 0, 0, 0, 282, 72, 1, 0, 0, 0, 283, 285, 7, 3, 0, 0,
		284, 283, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 284, 1, 0, 0, 0, 286,
		287, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 289, 6, 36, 0, 0, 289, 74,
		1, 0, 0, 0, 290, 291, 5, 47, 0, 0, 291, 297, 5, 47, 0, 0, 292, 293, 5,
		60, 0, 0, 293, 294, 5, 45, 0, 0, 294, 297, 5, 45, 0, 0, 295, 297, 5, 35,
		0, 0, 296, 290, 1, 0, 0, 0, 296, 292, 1, 0, 0, 0, 296, 295, 1, 0, 0, 0,
		297, 301, 1, 0, 0, 0, 298, 300, 8, 4, 0, 0, 299, 298, 1, 0, 0, 0, 300,
		303, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 304,
		1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 304, 305, 6, 37, 0, 0, 305, 76, 1, 0,
		0, 0, 306, 307, 5, 47, 0, 0, 307, 308, 5, 42, 0, 0, 308, 312, 1, 0, 0,
		0, 309, 311, 9, 0, 0, 0, 310, 309, 1, 0, 0, 0, 311, 314, 1, 0, 0, 0, 312,
		313, 1, 0, 0, 0, 312, 310, 1, 0, 0, 0, 313, 315, 1, 0, 0, 0, 314, 312,
		1, 0, 0, 0, 315, 316, 5, 42, 0, 0, 316, 317, 5, 47, 0, 0, 317, 318, 1,
		0, 0, 0, 318, 319, 6, 38, 0, 0, 319, 78, 1, 0, 0, 0, 320, 321, 5, 116,
		0, 0, 321, 322, 5, 114, 0, 0, 322, 323, 5, 117, 0, 0, 323, 330, 5, 101,
		0, 0, 324, 325, 5, 102, 0, 0, 325, 326, 5, 97, 0, 0, 326, 327, 5, 108,
		0, 0, 327, 328, 5, 115, 0, 0, 328, 330, 5, 101, 0, 0, 329, 320, 1, 0, 0,
		0, 329, 324, 1, 0, 0, 0, 330, 80, 1, 0, 0, 0, 331, 332, 5, 110, 0, 0, 332,
		333, 5, 117, 0, 0, 333, 334, 5, 108, 0, 0, 334, 335, 5, 108, 0, 0, 335,
		82, 1, 0, 0, 0, 336, 337, 5, 39, 0, 0, 337, 84, 1, 0, 0, 0, 338, 342, 7,
		5, 0, 0, 339, 341, 7, 6, 0, 0, 340, 339, 1, 0, 0, 0, 341, 344, 1, 0, 0,
		0, 342, 340, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 86, 1, 0, 0, 0, 344,
		342, 1, 0, 0, 0, 345, 347, 5, 45, 0, 0, 346, 345, 1, 0, 0, 0, 346, 347,
		1, 0, 0, 0, 347, 349, 1, 0, 0, 0, 348, 350, 7, 7, 0, 0, 349, 348, 1, 0,
		0, 0, 350, 351, 1, 0, 0, 0, 351, 349, 1, 0, 0, 0, 351, 352, 1, 0, 0, 0,
		352, 88, 1, 0, 0, 0, 353, 355, 5, 45, 0, 0, 354, 353, 1, 0, 0, 0, 354,
		355, 1, 0, 0, 0, 355, 359, 1, 0, 0, 0, 356, 358, 7, 7, 0, 0, 357, 356,
		1, 0, 0, 0, 358, 361, 1, 0, 0, 0, 359, 357, 1, 0, 0, 0, 359, 360, 1, 0,
		0, 0, 360, 362, 1, 0, 0, 0, 361, 359, 1, 0, 0, 0, 362, 364, 5, 46, 0, 0,
		363, 365, 7, 7, 0, 0, 364, 363, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366,
		364, 1, 0, 0, 0, 366, 367, 1, 0, 0, 0, 367, 90, 1, 0, 0, 0, 368, 371, 5,
		34, 0, 0, 369, 371, 3, 83, 41, 0, 370, 368, 1, 0, 0, 0, 370, 369, 1, 0,
		0, 0, 371, 375, 1, 0, 0, 0, 372, 374, 9, 0, 0, 0, 373, 372, 1, 0, 0, 0,
		374, 377, 1, 0, 0, 0, 375, 376, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 376,
		380, 1, 0, 0, 0, 377, 375, 1, 0, 0, 0, 378, 381, 5, 34, 0, 0, 379, 381,
		3, 83, 41, 0, 380, 378, 1, 0, 0, 0, 380, 379, 1, 0, 0, 0, 381, 92, 1, 0,
		0, 0, 19, 0, 130, 135, 159, 281, 286, 296, 301, 312, 329, 342, 346, 351,
		354, 359, 366, 370, 375, 380, 1, 6, 0, 0,
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

// JadeLexerInit initializes any static state used to implement JadeLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewJadeLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func JadeLexerInit() {
	staticData := &jadelexerLexerStaticData
	staticData.once.Do(jadelexerLexerInit)
}

// NewJadeLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewJadeLexer(input antlr.CharStream) *JadeLexer {
	JadeLexerInit()
	l := new(JadeLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &jadelexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Jade.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JadeLexer tokens.
const (
	JadeLexerLPAREN           = 1
	JadeLexerRPAREN           = 2
	JadeLexerLBRACE           = 3
	JadeLexerRBRACE           = 4
	JadeLexerLLIST            = 5
	JadeLexerRLIST            = 6
	JadeLexerASSIGN           = 7
	JadeLexerDOT              = 8
	JadeLexerCOMMA            = 9
	JadeLexerCOLON            = 10
	JadeLexerSEMI             = 11
	JadeLexerLARROW           = 12
	JadeLexerRARROW           = 13
	JadeLexerARROWASSIGN      = 14
	JadeLexerQUESTION         = 15
	JadeLexerNOT              = 16
	JadeLexerPREDTWO          = 17
	JadeLexerPREDONE          = 18
	JadeLexerCOMPARATIVE      = 19
	JadeLexerFUNC             = 20
	JadeLexerIF               = 21
	JadeLexerELSE             = 22
	JadeLexerWHILE            = 23
	JadeLexerIMPORT           = 24
	JadeLexerFROM             = 25
	JadeLexerOVERRIDE         = 26
	JadeLexerNEW              = 27
	JadeLexerCLASS            = 28
	JadeLexerPUBLIC           = 29
	JadeLexerPRIVATE          = 30
	JadeLexerRETURN           = 31
	JadeLexerBREAK            = 32
	JadeLexerCONTINUE         = 33
	JadeLexerUNDEFINE         = 34
	JadeLexerCONST            = 35
	JadeLexerALL_KWS          = 36
	JadeLexerWS               = 37
	JadeLexerCOMMENT          = 38
	JadeLexerMULTILINECOMMENT = 39
	JadeLexerBOOL             = 40
	JadeLexerNUL              = 41
	JadeLexerAPOSTROPHE       = 42
	JadeLexerID               = 43
	JadeLexerINT              = 44
	JadeLexerFLOAT            = 45
	JadeLexerSTRING           = 46
)