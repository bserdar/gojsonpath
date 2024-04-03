// Code generated from Jsonpath.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Jsonpath

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type JsonpathParser struct {
	*antlr.BaseParser
}

var JsonpathParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jsonpathParserInit() {
	staticData := &JsonpathParserStaticData
	staticData.LiteralNames = []string{
		"", "'/'", "'*'", "'$'", "'@'", "'('", "','", "')'", "':'", "'.'", "'..'",
		"'['", "']'", "'+'", "'-'", "'~'", "'!'", "'**'", "'%'", "'<'", "'>'",
		"'<='", "'>='", "'in'", "'=='", "'!='", "'==='", "'!=='", "'&'", "'^'",
		"'|'", "'&&'", "'||'", "'?('", "'...'", "'null'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"Ellipsis", "NullLiteral", "BooleanLiteral", "DecimalLiteral", "StringLiteral",
		"HexIntegerLiteral", "OctalIntegerLiteral2", "BinaryIntegerLiteral",
		"SP", "Identifier",
	}
	staticData.RuleNames = []string{
		"path", "simplePath", "simplePathExpr", "selector", "arguments", "slice",
		"indexExpression", "expression", "arrayLiteral", "elementList", "literal",
		"identifier", "numericLiteral",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 43, 372, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 3, 0, 29, 8, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 3, 1, 36, 8, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 3,
		2, 44, 8, 2, 1, 2, 1, 2, 3, 2, 48, 8, 2, 1, 2, 1, 2, 3, 2, 52, 8, 2, 1,
		2, 5, 2, 55, 8, 2, 10, 2, 12, 2, 58, 9, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 5, 4, 66, 8, 4, 10, 4, 12, 4, 69, 9, 4, 1, 4, 3, 4, 72, 8, 4, 3,
		4, 74, 8, 4, 1, 4, 1, 4, 1, 5, 3, 5, 79, 8, 5, 1, 5, 1, 5, 3, 5, 83, 8,
		5, 1, 5, 3, 5, 86, 8, 5, 1, 5, 3, 5, 89, 8, 5, 1, 6, 1, 6, 3, 6, 93, 8,
		6, 1, 7, 1, 7, 1, 7, 3, 7, 98, 8, 7, 1, 7, 1, 7, 1, 7, 3, 7, 103, 8, 7,
		1, 7, 1, 7, 1, 7, 3, 7, 108, 8, 7, 1, 7, 1, 7, 1, 7, 3, 7, 113, 8, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 3, 7, 128, 8, 7, 1, 7, 1, 7, 3, 7, 132, 8, 7, 1, 7, 1, 7, 3, 7, 136,
		8, 7, 1, 7, 1, 7, 1, 7, 3, 7, 141, 8, 7, 1, 7, 1, 7, 3, 7, 145, 8, 7, 1,
		7, 1, 7, 1, 7, 3, 7, 150, 8, 7, 1, 7, 1, 7, 3, 7, 154, 8, 7, 1, 7, 1, 7,
		1, 7, 3, 7, 159, 8, 7, 1, 7, 1, 7, 3, 7, 163, 8, 7, 1, 7, 1, 7, 1, 7, 3,
		7, 168, 8, 7, 1, 7, 1, 7, 3, 7, 172, 8, 7, 1, 7, 1, 7, 1, 7, 3, 7, 177,
		8, 7, 1, 7, 1, 7, 3, 7, 181, 8, 7, 1, 7, 1, 7, 1, 7, 3, 7, 186, 8, 7, 1,
		7, 1, 7, 3, 7, 190, 8, 7, 1, 7, 1, 7, 1, 7, 3, 7, 195, 8, 7, 1, 7, 1, 7,
		3, 7, 199, 8, 7, 1, 7, 1, 7, 1, 7, 3, 7, 204, 8, 7, 1, 7, 1, 7, 3, 7, 208,
		8, 7, 1, 7, 1, 7, 1, 7, 3, 7, 213, 8, 7, 1, 7, 1, 7, 3, 7, 217, 8, 7, 1,
		7, 1, 7, 1, 7, 3, 7, 222, 8, 7, 1, 7, 1, 7, 3, 7, 226, 8, 7, 1, 7, 1, 7,
		1, 7, 3, 7, 231, 8, 7, 1, 7, 1, 7, 3, 7, 235, 8, 7, 1, 7, 1, 7, 1, 7, 3,
		7, 240, 8, 7, 1, 7, 1, 7, 3, 7, 244, 8, 7, 1, 7, 1, 7, 1, 7, 3, 7, 249,
		8, 7, 1, 7, 1, 7, 1, 7, 3, 7, 254, 8, 7, 1, 7, 1, 7, 3, 7, 258, 8, 7, 1,
		7, 1, 7, 3, 7, 262, 8, 7, 1, 7, 1, 7, 3, 7, 266, 8, 7, 1, 7, 5, 7, 269,
		8, 7, 10, 7, 12, 7, 272, 9, 7, 1, 7, 3, 7, 275, 8, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 3, 7, 281, 8, 7, 1, 7, 1, 7, 3, 7, 285, 8, 7, 1, 7, 1, 7, 3, 7, 289,
		8, 7, 1, 7, 1, 7, 3, 7, 293, 8, 7, 1, 7, 1, 7, 3, 7, 297, 8, 7, 1, 7, 5,
		7, 300, 8, 7, 10, 7, 12, 7, 303, 9, 7, 1, 7, 3, 7, 306, 8, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 5, 7, 312, 8, 7, 10, 7, 12, 7, 315, 9, 7, 1, 8, 1, 8, 3,
		8, 319, 8, 8, 1, 8, 1, 8, 3, 8, 323, 8, 8, 1, 8, 1, 8, 1, 9, 5, 9, 328,
		8, 9, 10, 9, 12, 9, 331, 9, 9, 1, 9, 3, 9, 334, 8, 9, 1, 9, 3, 9, 337,
		8, 9, 1, 9, 3, 9, 340, 8, 9, 1, 9, 4, 9, 343, 8, 9, 11, 9, 12, 9, 344,
		1, 9, 3, 9, 348, 8, 9, 1, 9, 5, 9, 351, 8, 9, 10, 9, 12, 9, 354, 9, 9,
		1, 9, 5, 9, 357, 8, 9, 10, 9, 12, 9, 360, 9, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 3, 10, 366, 8, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 0, 2, 4, 14,
		13, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 0, 6, 1, 0, 2, 4, 2,
		0, 1, 2, 18, 18, 1, 0, 13, 14, 1, 0, 19, 22, 1, 0, 24, 27, 2, 0, 37, 37,
		39, 41, 456, 0, 26, 1, 0, 0, 0, 2, 32, 1, 0, 0, 0, 4, 43, 1, 0, 0, 0, 6,
		59, 1, 0, 0, 0, 8, 61, 1, 0, 0, 0, 10, 78, 1, 0, 0, 0, 12, 92, 1, 0, 0,
		0, 14, 127, 1, 0, 0, 0, 16, 316, 1, 0, 0, 0, 18, 329, 1, 0, 0, 0, 20, 365,
		1, 0, 0, 0, 22, 367, 1, 0, 0, 0, 24, 369, 1, 0, 0, 0, 26, 28, 3, 14, 7,
		0, 27, 29, 5, 42, 0, 0, 28, 27, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 30,
		1, 0, 0, 0, 30, 31, 5, 0, 0, 1, 31, 1, 1, 0, 0, 0, 32, 33, 5, 1, 0, 0,
		33, 35, 3, 4, 2, 0, 34, 36, 5, 42, 0, 0, 35, 34, 1, 0, 0, 0, 35, 36, 1,
		0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 38, 5, 0, 0, 1, 38, 3, 1, 0, 0, 0, 39,
		40, 6, 2, -1, 0, 40, 44, 3, 20, 10, 0, 41, 44, 3, 22, 11, 0, 42, 44, 5,
		2, 0, 0, 43, 39, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 43, 42, 1, 0, 0, 0, 44,
		56, 1, 0, 0, 0, 45, 47, 10, 4, 0, 0, 46, 48, 5, 42, 0, 0, 47, 46, 1, 0,
		0, 0, 47, 48, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 51, 5, 1, 0, 0, 50, 52,
		5, 42, 0, 0, 51, 50, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0,
		53, 55, 3, 4, 2, 5, 54, 45, 1, 0, 0, 0, 55, 58, 1, 0, 0, 0, 56, 54, 1,
		0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 5, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 59,
		60, 7, 0, 0, 0, 60, 7, 1, 0, 0, 0, 61, 73, 5, 5, 0, 0, 62, 67, 3, 14, 7,
		0, 63, 64, 5, 6, 0, 0, 64, 66, 3, 14, 7, 0, 65, 63, 1, 0, 0, 0, 66, 69,
		1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 71, 1, 0, 0, 0,
		69, 67, 1, 0, 0, 0, 70, 72, 5, 6, 0, 0, 71, 70, 1, 0, 0, 0, 71, 72, 1,
		0, 0, 0, 72, 74, 1, 0, 0, 0, 73, 62, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74,
		75, 1, 0, 0, 0, 75, 76, 5, 7, 0, 0, 76, 9, 1, 0, 0, 0, 77, 79, 3, 14, 7,
		0, 78, 77, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 82,
		5, 8, 0, 0, 81, 83, 3, 14, 7, 0, 82, 81, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0,
		83, 85, 1, 0, 0, 0, 84, 86, 5, 8, 0, 0, 85, 84, 1, 0, 0, 0, 85, 86, 1,
		0, 0, 0, 86, 88, 1, 0, 0, 0, 87, 89, 3, 14, 7, 0, 88, 87, 1, 0, 0, 0, 88,
		89, 1, 0, 0, 0, 89, 11, 1, 0, 0, 0, 90, 93, 3, 10, 5, 0, 91, 93, 3, 14,
		7, 0, 92, 90, 1, 0, 0, 0, 92, 91, 1, 0, 0, 0, 93, 13, 1, 0, 0, 0, 94, 95,
		6, 7, -1, 0, 95, 97, 5, 13, 0, 0, 96, 98, 5, 42, 0, 0, 97, 96, 1, 0, 0,
		0, 97, 98, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 128, 3, 14, 7, 21, 100,
		102, 5, 14, 0, 0, 101, 103, 5, 42, 0, 0, 102, 101, 1, 0, 0, 0, 102, 103,
		1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 128, 3, 14, 7, 20, 105, 107, 5,
		15, 0, 0, 106, 108, 5, 42, 0, 0, 107, 106, 1, 0, 0, 0, 107, 108, 1, 0,
		0, 0, 108, 109, 1, 0, 0, 0, 109, 128, 3, 14, 7, 19, 110, 112, 5, 16, 0,
		0, 111, 113, 5, 42, 0, 0, 112, 111, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113,
		114, 1, 0, 0, 0, 114, 128, 3, 14, 7, 18, 115, 128, 3, 20, 10, 0, 116, 128,
		3, 16, 8, 0, 117, 128, 3, 22, 11, 0, 118, 128, 3, 6, 3, 0, 119, 120, 5,
		5, 0, 0, 120, 121, 3, 14, 7, 0, 121, 122, 5, 7, 0, 0, 122, 128, 1, 0, 0,
		0, 123, 124, 5, 33, 0, 0, 124, 125, 3, 14, 7, 0, 125, 126, 5, 7, 0, 0,
		126, 128, 1, 0, 0, 0, 127, 94, 1, 0, 0, 0, 127, 100, 1, 0, 0, 0, 127, 105,
		1, 0, 0, 0, 127, 110, 1, 0, 0, 0, 127, 115, 1, 0, 0, 0, 127, 116, 1, 0,
		0, 0, 127, 117, 1, 0, 0, 0, 127, 118, 1, 0, 0, 0, 127, 119, 1, 0, 0, 0,
		127, 123, 1, 0, 0, 0, 128, 313, 1, 0, 0, 0, 129, 131, 10, 27, 0, 0, 130,
		132, 5, 42, 0, 0, 131, 130, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 133,
		1, 0, 0, 0, 133, 135, 5, 9, 0, 0, 134, 136, 5, 42, 0, 0, 135, 134, 1, 0,
		0, 0, 135, 136, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 312, 3, 14, 7, 28,
		138, 140, 10, 26, 0, 0, 139, 141, 5, 42, 0, 0, 140, 139, 1, 0, 0, 0, 140,
		141, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 144, 5, 10, 0, 0, 143, 145,
		5, 42, 0, 0, 144, 143, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 146, 1, 0,
		0, 0, 146, 312, 3, 14, 7, 27, 147, 149, 10, 17, 0, 0, 148, 150, 5, 42,
		0, 0, 149, 148, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0,
		151, 153, 5, 17, 0, 0, 152, 154, 5, 42, 0, 0, 153, 152, 1, 0, 0, 0, 153,
		154, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 312, 3, 14, 7, 17, 156, 158,
		10, 16, 0, 0, 157, 159, 5, 42, 0, 0, 158, 157, 1, 0, 0, 0, 158, 159, 1,
		0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 162, 7, 1, 0, 0, 161, 163, 5, 42, 0,
		0, 162, 161, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164,
		312, 3, 14, 7, 17, 165, 167, 10, 15, 0, 0, 166, 168, 5, 42, 0, 0, 167,
		166, 1, 0, 0, 0, 167, 168, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 171,
		7, 2, 0, 0, 170, 172, 5, 42, 0, 0, 171, 170, 1, 0, 0, 0, 171, 172, 1, 0,
		0, 0, 172, 173, 1, 0, 0, 0, 173, 312, 3, 14, 7, 16, 174, 176, 10, 14, 0,
		0, 175, 177, 5, 42, 0, 0, 176, 175, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177,
		178, 1, 0, 0, 0, 178, 180, 7, 3, 0, 0, 179, 181, 5, 42, 0, 0, 180, 179,
		1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 312, 3, 14,
		7, 15, 183, 185, 10, 13, 0, 0, 184, 186, 5, 42, 0, 0, 185, 184, 1, 0, 0,
		0, 185, 186, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 189, 5, 23, 0, 0, 188,
		190, 5, 42, 0, 0, 189, 188, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 191,
		1, 0, 0, 0, 191, 312, 3, 14, 7, 14, 192, 194, 10, 12, 0, 0, 193, 195, 5,
		42, 0, 0, 194, 193, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 196, 1, 0, 0,
		0, 196, 198, 7, 4, 0, 0, 197, 199, 5, 42, 0, 0, 198, 197, 1, 0, 0, 0, 198,
		199, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 312, 3, 14, 7, 13, 201, 203,
		10, 11, 0, 0, 202, 204, 5, 42, 0, 0, 203, 202, 1, 0, 0, 0, 203, 204, 1,
		0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 207, 5, 28, 0, 0, 206, 208, 5, 42,
		0, 0, 207, 206, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0,
		209, 312, 3, 14, 7, 12, 210, 212, 10, 10, 0, 0, 211, 213, 5, 42, 0, 0,
		212, 211, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214,
		216, 5, 29, 0, 0, 215, 217, 5, 42, 0, 0, 216, 215, 1, 0, 0, 0, 216, 217,
		1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218, 312, 3, 14, 7, 11, 219, 221, 10,
		9, 0, 0, 220, 222, 5, 42, 0, 0, 221, 220, 1, 0, 0, 0, 221, 222, 1, 0, 0,
		0, 222, 223, 1, 0, 0, 0, 223, 225, 5, 30, 0, 0, 224, 226, 5, 42, 0, 0,
		225, 224, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227,
		312, 3, 14, 7, 10, 228, 230, 10, 8, 0, 0, 229, 231, 5, 42, 0, 0, 230, 229,
		1, 0, 0, 0, 230, 231, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 234, 5, 31,
		0, 0, 233, 235, 5, 42, 0, 0, 234, 233, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0,
		235, 236, 1, 0, 0, 0, 236, 312, 3, 14, 7, 9, 237, 239, 10, 7, 0, 0, 238,
		240, 5, 42, 0, 0, 239, 238, 1, 0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 241,
		1, 0, 0, 0, 241, 243, 5, 32, 0, 0, 242, 244, 5, 42, 0, 0, 243, 242, 1,
		0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 312, 3, 14, 7,
		8, 246, 248, 10, 25, 0, 0, 247, 249, 5, 42, 0, 0, 248, 247, 1, 0, 0, 0,
		248, 249, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 312, 5, 10, 0, 0, 251,
		253, 10, 24, 0, 0, 252, 254, 5, 42, 0, 0, 253, 252, 1, 0, 0, 0, 253, 254,
		1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 257, 5, 11, 0, 0, 256, 258, 5, 42,
		0, 0, 257, 256, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0,
		259, 270, 3, 12, 6, 0, 260, 262, 5, 42, 0, 0, 261, 260, 1, 0, 0, 0, 261,
		262, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 265, 5, 6, 0, 0, 264, 266,
		5, 42, 0, 0, 265, 264, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 267, 1, 0,
		0, 0, 267, 269, 3, 12, 6, 0, 268, 261, 1, 0, 0, 0, 269, 272, 1, 0, 0, 0,
		270, 268, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 274, 1, 0, 0, 0, 272,
		270, 1, 0, 0, 0, 273, 275, 5, 42, 0, 0, 274, 273, 1, 0, 0, 0, 274, 275,
		1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276, 277, 5, 12, 0, 0, 277, 312, 1, 0,
		0, 0, 278, 280, 10, 23, 0, 0, 279, 281, 5, 42, 0, 0, 280, 279, 1, 0, 0,
		0, 280, 281, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 284, 5, 10, 0, 0, 283,
		285, 5, 42, 0, 0, 284, 283, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 286,
		1, 0, 0, 0, 286, 288, 5, 11, 0, 0, 287, 289, 5, 42, 0, 0, 288, 287, 1,
		0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 301, 3, 12, 6,
		0, 291, 293, 5, 42, 0, 0, 292, 291, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293,
		294, 1, 0, 0, 0, 294, 296, 5, 6, 0, 0, 295, 297, 5, 42, 0, 0, 296, 295,
		1, 0, 0, 0, 296, 297, 1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 300, 3, 12,
		6, 0, 299, 292, 1, 0, 0, 0, 300, 303, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0,
		301, 302, 1, 0, 0, 0, 302, 305, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 304,
		306, 5, 42, 0, 0, 305, 304, 1, 0, 0, 0, 305, 306, 1, 0, 0, 0, 306, 307,
		1, 0, 0, 0, 307, 308, 5, 12, 0, 0, 308, 312, 1, 0, 0, 0, 309, 310, 10,
		22, 0, 0, 310, 312, 3, 8, 4, 0, 311, 129, 1, 0, 0, 0, 311, 138, 1, 0, 0,
		0, 311, 147, 1, 0, 0, 0, 311, 156, 1, 0, 0, 0, 311, 165, 1, 0, 0, 0, 311,
		174, 1, 0, 0, 0, 311, 183, 1, 0, 0, 0, 311, 192, 1, 0, 0, 0, 311, 201,
		1, 0, 0, 0, 311, 210, 1, 0, 0, 0, 311, 219, 1, 0, 0, 0, 311, 228, 1, 0,
		0, 0, 311, 237, 1, 0, 0, 0, 311, 246, 1, 0, 0, 0, 311, 251, 1, 0, 0, 0,
		311, 278, 1, 0, 0, 0, 311, 309, 1, 0, 0, 0, 312, 315, 1, 0, 0, 0, 313,
		311, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 15, 1, 0, 0, 0, 315, 313, 1,
		0, 0, 0, 316, 318, 5, 11, 0, 0, 317, 319, 5, 42, 0, 0, 318, 317, 1, 0,
		0, 0, 318, 319, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 322, 3, 18, 9, 0,
		321, 323, 5, 42, 0, 0, 322, 321, 1, 0, 0, 0, 322, 323, 1, 0, 0, 0, 323,
		324, 1, 0, 0, 0, 324, 325, 5, 12, 0, 0, 325, 17, 1, 0, 0, 0, 326, 328,
		5, 6, 0, 0, 327, 326, 1, 0, 0, 0, 328, 331, 1, 0, 0, 0, 329, 327, 1, 0,
		0, 0, 329, 330, 1, 0, 0, 0, 330, 333, 1, 0, 0, 0, 331, 329, 1, 0, 0, 0,
		332, 334, 5, 42, 0, 0, 333, 332, 1, 0, 0, 0, 333, 334, 1, 0, 0, 0, 334,
		336, 1, 0, 0, 0, 335, 337, 3, 14, 7, 0, 336, 335, 1, 0, 0, 0, 336, 337,
		1, 0, 0, 0, 337, 352, 1, 0, 0, 0, 338, 340, 5, 42, 0, 0, 339, 338, 1, 0,
		0, 0, 339, 340, 1, 0, 0, 0, 340, 342, 1, 0, 0, 0, 341, 343, 5, 6, 0, 0,
		342, 341, 1, 0, 0, 0, 343, 344, 1, 0, 0, 0, 344, 342, 1, 0, 0, 0, 344,
		345, 1, 0, 0, 0, 345, 347, 1, 0, 0, 0, 346, 348, 5, 42, 0, 0, 347, 346,
		1, 0, 0, 0, 347, 348, 1, 0, 0, 0, 348, 349, 1, 0, 0, 0, 349, 351, 3, 14,
		7, 0, 350, 339, 1, 0, 0, 0, 351, 354, 1, 0, 0, 0, 352, 350, 1, 0, 0, 0,
		352, 353, 1, 0, 0, 0, 353, 358, 1, 0, 0, 0, 354, 352, 1, 0, 0, 0, 355,
		357, 5, 6, 0, 0, 356, 355, 1, 0, 0, 0, 357, 360, 1, 0, 0, 0, 358, 356,
		1, 0, 0, 0, 358, 359, 1, 0, 0, 0, 359, 19, 1, 0, 0, 0, 360, 358, 1, 0,
		0, 0, 361, 366, 5, 35, 0, 0, 362, 366, 5, 36, 0, 0, 363, 366, 3, 24, 12,
		0, 364, 366, 5, 38, 0, 0, 365, 361, 1, 0, 0, 0, 365, 362, 1, 0, 0, 0, 365,
		363, 1, 0, 0, 0, 365, 364, 1, 0, 0, 0, 366, 21, 1, 0, 0, 0, 367, 368, 5,
		43, 0, 0, 368, 23, 1, 0, 0, 0, 369, 370, 7, 5, 0, 0, 370, 25, 1, 0, 0,
		0, 72, 28, 35, 43, 47, 51, 56, 67, 71, 73, 78, 82, 85, 88, 92, 97, 102,
		107, 112, 127, 131, 135, 140, 144, 149, 153, 158, 162, 167, 171, 176, 180,
		185, 189, 194, 198, 203, 207, 212, 216, 221, 225, 230, 234, 239, 243, 248,
		253, 257, 261, 265, 270, 274, 280, 284, 288, 292, 296, 301, 305, 311, 313,
		318, 322, 329, 333, 336, 339, 344, 347, 352, 358, 365,
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

// JsonpathParserInit initializes any static state used to implement JsonpathParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJsonpathParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JsonpathParserInit() {
	staticData := &JsonpathParserStaticData
	staticData.once.Do(jsonpathParserInit)
}

// NewJsonpathParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJsonpathParser(input antlr.TokenStream) *JsonpathParser {
	JsonpathParserInit()
	this := new(JsonpathParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &JsonpathParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Jsonpath.g4"

	return this
}

// JsonpathParser tokens.
const (
	JsonpathParserEOF                  = antlr.TokenEOF
	JsonpathParserT__0                 = 1
	JsonpathParserT__1                 = 2
	JsonpathParserT__2                 = 3
	JsonpathParserT__3                 = 4
	JsonpathParserT__4                 = 5
	JsonpathParserT__5                 = 6
	JsonpathParserT__6                 = 7
	JsonpathParserT__7                 = 8
	JsonpathParserT__8                 = 9
	JsonpathParserT__9                 = 10
	JsonpathParserT__10                = 11
	JsonpathParserT__11                = 12
	JsonpathParserT__12                = 13
	JsonpathParserT__13                = 14
	JsonpathParserT__14                = 15
	JsonpathParserT__15                = 16
	JsonpathParserT__16                = 17
	JsonpathParserT__17                = 18
	JsonpathParserT__18                = 19
	JsonpathParserT__19                = 20
	JsonpathParserT__20                = 21
	JsonpathParserT__21                = 22
	JsonpathParserT__22                = 23
	JsonpathParserT__23                = 24
	JsonpathParserT__24                = 25
	JsonpathParserT__25                = 26
	JsonpathParserT__26                = 27
	JsonpathParserT__27                = 28
	JsonpathParserT__28                = 29
	JsonpathParserT__29                = 30
	JsonpathParserT__30                = 31
	JsonpathParserT__31                = 32
	JsonpathParserT__32                = 33
	JsonpathParserEllipsis             = 34
	JsonpathParserNullLiteral          = 35
	JsonpathParserBooleanLiteral       = 36
	JsonpathParserDecimalLiteral       = 37
	JsonpathParserStringLiteral        = 38
	JsonpathParserHexIntegerLiteral    = 39
	JsonpathParserOctalIntegerLiteral2 = 40
	JsonpathParserBinaryIntegerLiteral = 41
	JsonpathParserSP                   = 42
	JsonpathParserIdentifier           = 43
)

// JsonpathParser rules.
const (
	JsonpathParserRULE_path            = 0
	JsonpathParserRULE_simplePath      = 1
	JsonpathParserRULE_simplePathExpr  = 2
	JsonpathParserRULE_selector        = 3
	JsonpathParserRULE_arguments       = 4
	JsonpathParserRULE_slice           = 5
	JsonpathParserRULE_indexExpression = 6
	JsonpathParserRULE_expression      = 7
	JsonpathParserRULE_arrayLiteral    = 8
	JsonpathParserRULE_elementList     = 9
	JsonpathParserRULE_literal         = 10
	JsonpathParserRULE_identifier      = 11
	JsonpathParserRULE_numericLiteral  = 12
)

// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	EOF() antlr.TerminalNode
	SP() antlr.TerminalNode

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_path
	return p
}

func InitEmptyPathContext(p *PathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_path
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PathContext) EOF() antlr.TerminalNode {
	return s.GetToken(JsonpathParserEOF, 0)
}

func (s *PathContext) SP() antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, 0)
}

func (s *PathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterPath(s)
	}
}

func (s *PathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitPath(s)
	}
}

func (p *JsonpathParser) Path() (localctx IPathContext) {
	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JsonpathParserRULE_path)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.expression(0)
	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JsonpathParserSP {
		{
			p.SetState(27)
			p.Match(JsonpathParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(30)
		p.Match(JsonpathParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimplePathContext is an interface to support dynamic dispatch.
type ISimplePathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SimplePathExpr() ISimplePathExprContext
	EOF() antlr.TerminalNode
	SP() antlr.TerminalNode

	// IsSimplePathContext differentiates from other interfaces.
	IsSimplePathContext()
}

type SimplePathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimplePathContext() *SimplePathContext {
	var p = new(SimplePathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_simplePath
	return p
}

func InitEmptySimplePathContext(p *SimplePathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_simplePath
}

func (*SimplePathContext) IsSimplePathContext() {}

func NewSimplePathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimplePathContext {
	var p = new(SimplePathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_simplePath

	return p
}

func (s *SimplePathContext) GetParser() antlr.Parser { return s.parser }

func (s *SimplePathContext) SimplePathExpr() ISimplePathExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimplePathExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimplePathExprContext)
}

func (s *SimplePathContext) EOF() antlr.TerminalNode {
	return s.GetToken(JsonpathParserEOF, 0)
}

func (s *SimplePathContext) SP() antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, 0)
}

func (s *SimplePathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimplePathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimplePathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterSimplePath(s)
	}
}

func (s *SimplePathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitSimplePath(s)
	}
}

func (p *JsonpathParser) SimplePath() (localctx ISimplePathContext) {
	localctx = NewSimplePathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JsonpathParserRULE_simplePath)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Match(JsonpathParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(33)
		p.simplePathExpr(0)
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JsonpathParserSP {
		{
			p.SetState(34)
			p.Match(JsonpathParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(37)
		p.Match(JsonpathParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimplePathExprContext is an interface to support dynamic dispatch.
type ISimplePathExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal() ILiteralContext
	Identifier() IIdentifierContext
	AllSimplePathExpr() []ISimplePathExprContext
	SimplePathExpr(i int) ISimplePathExprContext
	AllSP() []antlr.TerminalNode
	SP(i int) antlr.TerminalNode

	// IsSimplePathExprContext differentiates from other interfaces.
	IsSimplePathExprContext()
}

type SimplePathExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimplePathExprContext() *SimplePathExprContext {
	var p = new(SimplePathExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_simplePathExpr
	return p
}

func InitEmptySimplePathExprContext(p *SimplePathExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_simplePathExpr
}

func (*SimplePathExprContext) IsSimplePathExprContext() {}

func NewSimplePathExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimplePathExprContext {
	var p = new(SimplePathExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_simplePathExpr

	return p
}

func (s *SimplePathExprContext) GetParser() antlr.Parser { return s.parser }

func (s *SimplePathExprContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *SimplePathExprContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *SimplePathExprContext) AllSimplePathExpr() []ISimplePathExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimplePathExprContext); ok {
			len++
		}
	}

	tst := make([]ISimplePathExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimplePathExprContext); ok {
			tst[i] = t.(ISimplePathExprContext)
			i++
		}
	}

	return tst
}

func (s *SimplePathExprContext) SimplePathExpr(i int) ISimplePathExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimplePathExprContext); ok {
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

	return t.(ISimplePathExprContext)
}

func (s *SimplePathExprContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserSP)
}

func (s *SimplePathExprContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, i)
}

func (s *SimplePathExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimplePathExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimplePathExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterSimplePathExpr(s)
	}
}

func (s *SimplePathExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitSimplePathExpr(s)
	}
}

func (p *JsonpathParser) SimplePathExpr() (localctx ISimplePathExprContext) {
	return p.simplePathExpr(0)
}

func (p *JsonpathParser) simplePathExpr(_p int) (localctx ISimplePathExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewSimplePathExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISimplePathExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, JsonpathParserRULE_simplePathExpr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JsonpathParserNullLiteral, JsonpathParserBooleanLiteral, JsonpathParserDecimalLiteral, JsonpathParserStringLiteral, JsonpathParserHexIntegerLiteral, JsonpathParserOctalIntegerLiteral2, JsonpathParserBinaryIntegerLiteral:
		{
			p.SetState(40)
			p.Literal()
		}

	case JsonpathParserIdentifier:
		{
			p.SetState(41)
			p.Identifier()
		}

	case JsonpathParserT__1:
		{
			p.SetState(42)
			p.Match(JsonpathParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewSimplePathExprContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_simplePathExpr)
			p.SetState(45)

			if !(p.Precpred(p.GetParserRuleContext(), 4)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				goto errorExit
			}
			p.SetState(47)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == JsonpathParserSP {
				{
					p.SetState(46)
					p.Match(JsonpathParserSP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(49)
				p.Match(JsonpathParserT__0)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(51)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == JsonpathParserSP {
				{
					p.SetState(50)
					p.Match(JsonpathParserSP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(53)
				p.simplePathExpr(5)
			}

		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectorContext is an interface to support dynamic dispatch.
type ISelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSelectorContext differentiates from other interfaces.
	IsSelectorContext()
}

type SelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorContext() *SelectorContext {
	var p = new(SelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_selector
	return p
}

func InitEmptySelectorContext(p *SelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_selector
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_selector

	return p
}

func (s *SelectorContext) GetParser() antlr.Parser { return s.parser }
func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterSelector(s)
	}
}

func (s *SelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitSelector(s)
	}
}

func (p *JsonpathParser) Selector() (localctx ISelectorContext) {
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JsonpathParserRULE_selector)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&28) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *JsonpathParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JsonpathParserRULE_arguments)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Match(JsonpathParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13168369854524) != 0 {
		{
			p.SetState(62)
			p.expression(0)
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(63)
					p.Match(JsonpathParserT__5)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(64)
					p.expression(0)
				}

			}
			p.SetState(69)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonpathParserT__5 {
			{
				p.SetState(70)
				p.Match(JsonpathParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(75)
		p.Match(JsonpathParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISliceContext is an interface to support dynamic dispatch.
type ISliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsSliceContext differentiates from other interfaces.
	IsSliceContext()
}

type SliceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceContext() *SliceContext {
	var p = new(SliceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_slice
	return p
}

func InitEmptySliceContext(p *SliceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_slice
}

func (*SliceContext) IsSliceContext() {}

func NewSliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceContext {
	var p = new(SliceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_slice

	return p
}

func (s *SliceContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *SliceContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *SliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterSlice(s)
	}
}

func (s *SliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitSlice(s)
	}
}

func (p *JsonpathParser) Slice() (localctx ISliceContext) {
	localctx = NewSliceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JsonpathParserRULE_slice)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13168369854524) != 0 {
		{
			p.SetState(77)
			p.expression(0)
		}

	}
	{
		p.SetState(80)
		p.Match(JsonpathParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(81)
			p.expression(0)
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JsonpathParserT__7 {
		{
			p.SetState(84)
			p.Match(JsonpathParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13168369854524) != 0 {
		{
			p.SetState(87)
			p.expression(0)
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndexExpressionContext is an interface to support dynamic dispatch.
type IIndexExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Slice() ISliceContext
	Expression() IExpressionContext

	// IsIndexExpressionContext differentiates from other interfaces.
	IsIndexExpressionContext()
}

type IndexExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexExpressionContext() *IndexExpressionContext {
	var p = new(IndexExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_indexExpression
	return p
}

func InitEmptyIndexExpressionContext(p *IndexExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_indexExpression
}

func (*IndexExpressionContext) IsIndexExpressionContext() {}

func NewIndexExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexExpressionContext {
	var p = new(IndexExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_indexExpression

	return p
}

func (s *IndexExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexExpressionContext) Slice() ISliceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceContext)
}

func (s *IndexExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterIndexExpression(s)
	}
}

func (s *IndexExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitIndexExpression(s)
	}
}

func (p *JsonpathParser) IndexExpression() (localctx IIndexExpressionContext) {
	localctx = NewIndexExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JsonpathParserRULE_indexExpression)
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)
			p.Slice()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
			p.expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ChainExpressionContext struct {
	ExpressionContext
}

func NewChainExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ChainExpressionContext {
	var p = new(ChainExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ChainExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChainExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ChainExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ChainExpressionContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserSP)
}

func (s *ChainExpressionContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, i)
}

func (s *ChainExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterChainExpression(s)
	}
}

func (s *ChainExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitChainExpression(s)
	}
}

type LogicalAndExpressionContext struct {
	ExpressionContext
}

func NewLogicalAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalAndExpressionContext {
	var p = new(LogicalAndExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LogicalAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicalAndExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *LogicalAndExpressionContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserSP)
}

func (s *LogicalAndExpressionContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, i)
}

func (s *LogicalAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterLogicalAndExpression(s)
	}
}

func (s *LogicalAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitLogicalAndExpression(s)
	}
}

type PowerExpressionContext struct {
	ExpressionContext
}

func NewPowerExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowerExpressionContext {
	var p = new(PowerExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PowerExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PowerExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *PowerExpressionContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserSP)
}

func (s *PowerExpressionContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, i)
}

func (s *PowerExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterPowerExpression(s)
	}
}

func (s *PowerExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitPowerExpression(s)
	}
}

type InExpressionContext struct {
	ExpressionContext
}

func NewInExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InExpressionContext {
	var p = new(InExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *InExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *InExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *InExpressionContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserSP)
}

func (s *InExpressionContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, i)
}

func (s *InExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterInExpression(s)
	}
}

func (s *InExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitInExpression(s)
	}
}

type LogicalOrExpressionContext struct {
	ExpressionContext
}

func NewLogicalOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalOrExpressionContext {
	var p = new(LogicalOrExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LogicalOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogicalOrExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *LogicalOrExpressionContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserSP)
}

func (s *LogicalOrExpressionContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, i)
}

func (s *LogicalOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterLogicalOrExpression(s)
	}
}

func (s *LogicalOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitLogicalOrExpression(s)
	}
}

type NotExpressionContext struct {
	ExpressionContext
}

func NewNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExpressionContext {
	var p = new(NotExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NotExpressionContext) SP() antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, 0)
}

func (s *NotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterNotExpression(s)
	}
}

func (s *NotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitNotExpression(s)
	}
}

type SelectorExpressionContext struct {
	ExpressionContext
}

func NewSelectorExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectorExpressionContext {
	var p = new(SelectorExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *SelectorExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorExpressionContext) Selector() ISelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *SelectorExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterSelectorExpression(s)
	}
}

func (s *SelectorExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitSelectorExpression(s)
	}
}

type RecursiveDescentTermExpressionContext struct {
	ExpressionContext
}

func NewRecursiveDescentTermExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RecursiveDescentTermExpressionContext {
	var p = new(RecursiveDescentTermExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *RecursiveDescentTermExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecursiveDescentTermExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RecursiveDescentTermExpressionContext) SP() antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, 0)
}

func (s *RecursiveDescentTermExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterRecursiveDescentTermExpression(s)
	}
}

func (s *RecursiveDescentTermExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitRecursiveDescentTermExpression(s)
	}
}

type ArgumentsExpressionContext struct {
	ExpressionContext
}

func NewArgumentsExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArgumentsExpressionContext {
	var p = new(ArgumentsExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ArgumentsExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentsExpressionContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *ArgumentsExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterArgumentsExpression(s)
	}
}

func (s *ArgumentsExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitArgumentsExpression(s)
	}
}

type UnaryMinusExpressionContext struct {
	ExpressionContext
}

func NewUnaryMinusExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryMinusExpressionContext {
	var p = new(UnaryMinusExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryMinusExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMinusExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryMinusExpressionContext) SP() antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, 0)
}

func (s *UnaryMinusExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterUnaryMinusExpression(s)
	}
}

func (s *UnaryMinusExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitUnaryMinusExpression(s)
	}
}

type UnaryPlusExpressionContext struct {
	ExpressionContext
}

func NewUnaryPlusExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryPlusExpressionContext {
	var p = new(UnaryPlusExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryPlusExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryPlusExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryPlusExpressionContext) SP() antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, 0)
}

func (s *UnaryPlusExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterUnaryPlusExpression(s)
	}
}

func (s *UnaryPlusExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitUnaryPlusExpression(s)
	}
}

type FilterExpressionContext struct {
	ExpressionContext
}

func NewFilterExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FilterExpressionContext {
	var p = new(FilterExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FilterExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FilterExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterFilterExpression(s)
	}
}

func (s *FilterExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitFilterExpression(s)
	}
}

type EqualityExpressionContext struct {
	ExpressionContext
}

func NewEqualityExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *EqualityExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *EqualityExpressionContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserSP)
}

func (s *EqualityExpressionContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, i)
}

func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

type BitXOrExpressionContext struct {
	ExpressionContext
}

func NewBitXOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitXOrExpressionContext {
	var p = new(BitXOrExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BitXOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitXOrExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BitXOrExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *BitXOrExpressionContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserSP)
}

func (s *BitXOrExpressionContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, i)
}

func (s *BitXOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterBitXOrExpression(s)
	}
}

func (s *BitXOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitBitXOrExpression(s)
	}
}

type MultiplicativeExpressionContext struct {
	ExpressionContext
}

func NewMultiplicativeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *MultiplicativeExpressionContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserSP)
}

func (s *MultiplicativeExpressionContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, i)
}

func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitMultiplicativeExpression(s)
	}
}

type ParenthesizedExpressionContext struct {
	ExpressionContext
}

func NewParenthesizedExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesizedExpressionContext {
	var p = new(ParenthesizedExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParenthesizedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesizedExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenthesizedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterParenthesizedExpression(s)
	}
}

func (s *ParenthesizedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitParenthesizedExpression(s)
	}
}

type AdditiveExpressionContext struct {
	ExpressionContext
}

func NewAdditiveExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *AdditiveExpressionContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserSP)
}

func (s *AdditiveExpressionContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, i)
}

func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitAdditiveExpression(s)
	}
}

type RelationalExpressionContext struct {
	ExpressionContext
}

func NewRelationalExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RelationalExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *RelationalExpressionContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserSP)
}

func (s *RelationalExpressionContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, i)
}

func (s *RelationalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitRelationalExpression(s)
	}
}

type RecursiveDescentExpressionContext struct {
	ExpressionContext
}

func NewRecursiveDescentExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RecursiveDescentExpressionContext {
	var p = new(RecursiveDescentExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *RecursiveDescentExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecursiveDescentExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RecursiveDescentExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *RecursiveDescentExpressionContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserSP)
}

func (s *RecursiveDescentExpressionContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, i)
}

func (s *RecursiveDescentExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterRecursiveDescentExpression(s)
	}
}

func (s *RecursiveDescentExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitRecursiveDescentExpression(s)
	}
}

type BitNotExpressionContext struct {
	ExpressionContext
}

func NewBitNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitNotExpressionContext {
	var p = new(BitNotExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BitNotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitNotExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BitNotExpressionContext) SP() antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, 0)
}

func (s *BitNotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterBitNotExpression(s)
	}
}

func (s *BitNotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitBitNotExpression(s)
	}
}

type LiteralExpressionContext struct {
	ExpressionContext
}

func NewLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpressionContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitLiteralExpression(s)
	}
}

type ArrayLiteralExpressionContext struct {
	ExpressionContext
}

func NewArrayLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayLiteralExpressionContext {
	var p = new(ArrayLiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ArrayLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralExpressionContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *ArrayLiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterArrayLiteralExpression(s)
	}
}

func (s *ArrayLiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitArrayLiteralExpression(s)
	}
}

type MemberIndexExpressionContext struct {
	ExpressionContext
}

func NewMemberIndexExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberIndexExpressionContext {
	var p = new(MemberIndexExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MemberIndexExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberIndexExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MemberIndexExpressionContext) AllIndexExpression() []IIndexExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIndexExpressionContext); ok {
			len++
		}
	}

	tst := make([]IIndexExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIndexExpressionContext); ok {
			tst[i] = t.(IIndexExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MemberIndexExpressionContext) IndexExpression(i int) IIndexExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexExpressionContext); ok {
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

	return t.(IIndexExpressionContext)
}

func (s *MemberIndexExpressionContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserSP)
}

func (s *MemberIndexExpressionContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, i)
}

func (s *MemberIndexExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterMemberIndexExpression(s)
	}
}

func (s *MemberIndexExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitMemberIndexExpression(s)
	}
}

type IdentifierExpressionContext struct {
	ExpressionContext
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *IdentifierExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitIdentifierExpression(s)
	}
}

type BitAndExpressionContext struct {
	ExpressionContext
}

func NewBitAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitAndExpressionContext {
	var p = new(BitAndExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BitAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitAndExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BitAndExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *BitAndExpressionContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserSP)
}

func (s *BitAndExpressionContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, i)
}

func (s *BitAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterBitAndExpression(s)
	}
}

func (s *BitAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitBitAndExpression(s)
	}
}

type BitOrExpressionContext struct {
	ExpressionContext
}

func NewBitOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitOrExpressionContext {
	var p = new(BitOrExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BitOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitOrExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BitOrExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *BitOrExpressionContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserSP)
}

func (s *BitOrExpressionContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, i)
}

func (s *BitOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterBitOrExpression(s)
	}
}

func (s *BitOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitBitOrExpression(s)
	}
}

type RecursiveDescentMemberIndexExpressionContext struct {
	ExpressionContext
}

func NewRecursiveDescentMemberIndexExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RecursiveDescentMemberIndexExpressionContext {
	var p = new(RecursiveDescentMemberIndexExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *RecursiveDescentMemberIndexExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecursiveDescentMemberIndexExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RecursiveDescentMemberIndexExpressionContext) AllIndexExpression() []IIndexExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIndexExpressionContext); ok {
			len++
		}
	}

	tst := make([]IIndexExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIndexExpressionContext); ok {
			tst[i] = t.(IIndexExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RecursiveDescentMemberIndexExpressionContext) IndexExpression(i int) IIndexExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexExpressionContext); ok {
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

	return t.(IIndexExpressionContext)
}

func (s *RecursiveDescentMemberIndexExpressionContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserSP)
}

func (s *RecursiveDescentMemberIndexExpressionContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, i)
}

func (s *RecursiveDescentMemberIndexExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterRecursiveDescentMemberIndexExpression(s)
	}
}

func (s *RecursiveDescentMemberIndexExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitRecursiveDescentMemberIndexExpression(s)
	}
}

func (p *JsonpathParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *JsonpathParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, JsonpathParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JsonpathParserT__12:
		localctx = NewUnaryPlusExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(95)
			p.Match(JsonpathParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonpathParserSP {
			{
				p.SetState(96)
				p.Match(JsonpathParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(99)
			p.expression(21)
		}

	case JsonpathParserT__13:
		localctx = NewUnaryMinusExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(100)
			p.Match(JsonpathParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonpathParserSP {
			{
				p.SetState(101)
				p.Match(JsonpathParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(104)
			p.expression(20)
		}

	case JsonpathParserT__14:
		localctx = NewBitNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(105)
			p.Match(JsonpathParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonpathParserSP {
			{
				p.SetState(106)
				p.Match(JsonpathParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(109)
			p.expression(19)
		}

	case JsonpathParserT__15:
		localctx = NewNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(110)
			p.Match(JsonpathParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonpathParserSP {
			{
				p.SetState(111)
				p.Match(JsonpathParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(114)
			p.expression(18)
		}

	case JsonpathParserNullLiteral, JsonpathParserBooleanLiteral, JsonpathParserDecimalLiteral, JsonpathParserStringLiteral, JsonpathParserHexIntegerLiteral, JsonpathParserOctalIntegerLiteral2, JsonpathParserBinaryIntegerLiteral:
		localctx = NewLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(115)
			p.Literal()
		}

	case JsonpathParserT__10:
		localctx = NewArrayLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(116)
			p.ArrayLiteral()
		}

	case JsonpathParserIdentifier:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(117)
			p.Identifier()
		}

	case JsonpathParserT__1, JsonpathParserT__2, JsonpathParserT__3:
		localctx = NewSelectorExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(118)
			p.Selector()
		}

	case JsonpathParserT__4:
		localctx = NewParenthesizedExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(119)
			p.Match(JsonpathParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.expression(0)
		}
		{
			p.SetState(121)
			p.Match(JsonpathParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JsonpathParserT__32:
		localctx = NewFilterExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(123)
			p.Match(JsonpathParserT__32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.expression(0)
		}
		{
			p.SetState(125)
			p.Match(JsonpathParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(311)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext()) {
			case 1:
				localctx = NewChainExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(129)

				if !(p.Precpred(p.GetParserRuleContext(), 27)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 27)", ""))
					goto errorExit
				}
				p.SetState(131)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(130)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(133)
					p.Match(JsonpathParserT__8)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(135)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(134)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(137)
					p.expression(28)
				}

			case 2:
				localctx = NewRecursiveDescentExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(138)

				if !(p.Precpred(p.GetParserRuleContext(), 26)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 26)", ""))
					goto errorExit
				}
				p.SetState(140)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(139)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(142)
					p.Match(JsonpathParserT__9)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(144)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(143)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(146)
					p.expression(27)
				}

			case 3:
				localctx = NewPowerExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(147)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				p.SetState(149)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(148)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(151)
					p.Match(JsonpathParserT__16)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(153)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(152)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(155)
					p.expression(17)
				}

			case 4:
				localctx = NewMultiplicativeExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(156)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				p.SetState(158)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(157)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(160)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&262150) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(162)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(161)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(164)
					p.expression(17)
				}

			case 5:
				localctx = NewAdditiveExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(165)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				p.SetState(167)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(166)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(169)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JsonpathParserT__12 || _la == JsonpathParserT__13) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(171)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(170)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(173)
					p.expression(16)
				}

			case 6:
				localctx = NewRelationalExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(174)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				p.SetState(176)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(175)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(178)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7864320) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(180)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(179)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(182)
					p.expression(15)
				}

			case 7:
				localctx = NewInExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(183)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				p.SetState(185)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(184)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(187)
					p.Match(JsonpathParserT__22)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(189)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(188)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(191)
					p.expression(14)
				}

			case 8:
				localctx = NewEqualityExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(192)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				p.SetState(194)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(193)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(196)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&251658240) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(198)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(197)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(200)
					p.expression(13)
				}

			case 9:
				localctx = NewBitAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(201)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				p.SetState(203)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(202)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(205)
					p.Match(JsonpathParserT__27)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(207)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(206)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(209)
					p.expression(12)
				}

			case 10:
				localctx = NewBitXOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(210)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				p.SetState(212)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(211)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(214)
					p.Match(JsonpathParserT__28)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(216)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(215)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(218)
					p.expression(11)
				}

			case 11:
				localctx = NewBitOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(219)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				p.SetState(221)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(220)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(223)
					p.Match(JsonpathParserT__29)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(225)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(224)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(227)
					p.expression(10)
				}

			case 12:
				localctx = NewLogicalAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(228)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				p.SetState(230)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(229)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(232)
					p.Match(JsonpathParserT__30)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(234)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(233)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(236)
					p.expression(9)
				}

			case 13:
				localctx = NewLogicalOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(237)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				p.SetState(239)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(238)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(241)
					p.Match(JsonpathParserT__31)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(243)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(242)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(245)
					p.expression(8)
				}

			case 14:
				localctx = NewRecursiveDescentTermExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(246)

				if !(p.Precpred(p.GetParserRuleContext(), 25)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 25)", ""))
					goto errorExit
				}
				p.SetState(248)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(247)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(250)
					p.Match(JsonpathParserT__9)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 15:
				localctx = NewMemberIndexExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(251)

				if !(p.Precpred(p.GetParserRuleContext(), 24)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 24)", ""))
					goto errorExit
				}
				p.SetState(253)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(252)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(255)
					p.Match(JsonpathParserT__10)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(257)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(256)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(259)
					p.IndexExpression()
				}
				p.SetState(270)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
				for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					if _alt == 1 {
						p.SetState(261)
						p.GetErrorHandler().Sync(p)
						if p.HasError() {
							goto errorExit
						}
						_la = p.GetTokenStream().LA(1)

						if _la == JsonpathParserSP {
							{
								p.SetState(260)
								p.Match(JsonpathParserSP)
								if p.HasError() {
									// Recognition error - abort rule
									goto errorExit
								}
							}

						}
						{
							p.SetState(263)
							p.Match(JsonpathParserT__5)
							if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
							}
						}
						p.SetState(265)
						p.GetErrorHandler().Sync(p)
						if p.HasError() {
							goto errorExit
						}
						_la = p.GetTokenStream().LA(1)

						if _la == JsonpathParserSP {
							{
								p.SetState(264)
								p.Match(JsonpathParserSP)
								if p.HasError() {
									// Recognition error - abort rule
									goto errorExit
								}
							}

						}
						{
							p.SetState(267)
							p.IndexExpression()
						}

					}
					p.SetState(272)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext())
					if p.HasError() {
						goto errorExit
					}
				}
				p.SetState(274)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(273)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(276)
					p.Match(JsonpathParserT__11)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 16:
				localctx = NewRecursiveDescentMemberIndexExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(278)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
					goto errorExit
				}
				p.SetState(280)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(279)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(282)
					p.Match(JsonpathParserT__9)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(284)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(283)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(286)
					p.Match(JsonpathParserT__10)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(288)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(287)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(290)
					p.IndexExpression()
				}
				p.SetState(301)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
				for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					if _alt == 1 {
						p.SetState(292)
						p.GetErrorHandler().Sync(p)
						if p.HasError() {
							goto errorExit
						}
						_la = p.GetTokenStream().LA(1)

						if _la == JsonpathParserSP {
							{
								p.SetState(291)
								p.Match(JsonpathParserSP)
								if p.HasError() {
									// Recognition error - abort rule
									goto errorExit
								}
							}

						}
						{
							p.SetState(294)
							p.Match(JsonpathParserT__5)
							if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
							}
						}
						p.SetState(296)
						p.GetErrorHandler().Sync(p)
						if p.HasError() {
							goto errorExit
						}
						_la = p.GetTokenStream().LA(1)

						if _la == JsonpathParserSP {
							{
								p.SetState(295)
								p.Match(JsonpathParserSP)
								if p.HasError() {
									// Recognition error - abort rule
									goto errorExit
								}
							}

						}
						{
							p.SetState(298)
							p.IndexExpression()
						}

					}
					p.SetState(303)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext())
					if p.HasError() {
						goto errorExit
					}
				}
				p.SetState(305)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(304)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(307)
					p.Match(JsonpathParserT__11)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 17:
				localctx = NewArgumentsExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(309)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
					goto errorExit
				}
				{
					p.SetState(310)
					p.Arguments()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 60, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ElementList() IElementListContext
	AllSP() []antlr.TerminalNode
	SP(i int) antlr.TerminalNode

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_arrayLiteral
	return p
}

func InitEmptyArrayLiteralContext(p *ArrayLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_arrayLiteral
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) ElementList() IElementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElementListContext)
}

func (s *ArrayLiteralContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserSP)
}

func (s *ArrayLiteralContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, i)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (p *JsonpathParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JsonpathParserRULE_arrayLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(316)
		p.Match(JsonpathParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(318)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(317)
			p.Match(JsonpathParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(320)
		p.ElementList()
	}
	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JsonpathParserSP {
		{
			p.SetState(321)
			p.Match(JsonpathParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(324)
		p.Match(JsonpathParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElementListContext is an interface to support dynamic dispatch.
type IElementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSP() []antlr.TerminalNode
	SP(i int) antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsElementListContext differentiates from other interfaces.
	IsElementListContext()
}

type ElementListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementListContext() *ElementListContext {
	var p = new(ElementListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_elementList
	return p
}

func InitEmptyElementListContext(p *ElementListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_elementList
}

func (*ElementListContext) IsElementListContext() {}

func NewElementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementListContext {
	var p = new(ElementListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_elementList

	return p
}

func (s *ElementListContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementListContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserSP)
}

func (s *ElementListContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, i)
}

func (s *ElementListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ElementListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ElementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterElementList(s)
	}
}

func (s *ElementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitElementList(s)
	}
}

func (p *JsonpathParser) ElementList() (localctx IElementListContext) {
	localctx = NewElementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JsonpathParserRULE_elementList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 63, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(326)
				p.Match(JsonpathParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(331)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 63, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(333)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 64, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(332)
			p.Match(JsonpathParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13168369854524) != 0 {
		{
			p.SetState(335)
			p.expression(0)
		}

	}
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 69, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(339)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == JsonpathParserSP {
				{
					p.SetState(338)
					p.Match(JsonpathParserSP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(342)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == JsonpathParserT__5 {
				{
					p.SetState(341)
					p.Match(JsonpathParserT__5)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(344)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(347)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == JsonpathParserSP {
				{
					p.SetState(346)
					p.Match(JsonpathParserSP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(349)
				p.expression(0)
			}

		}
		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 69, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JsonpathParserT__5 {
		{
			p.SetState(355)
			p.Match(JsonpathParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NullLiteral() antlr.TerminalNode
	BooleanLiteral() antlr.TerminalNode
	NumericLiteral() INumericLiteralContext
	StringLiteral() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) NullLiteral() antlr.TerminalNode {
	return s.GetToken(JsonpathParserNullLiteral, 0)
}

func (s *LiteralContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(JsonpathParserBooleanLiteral, 0)
}

func (s *LiteralContext) NumericLiteral() INumericLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *LiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(JsonpathParserStringLiteral, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *JsonpathParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, JsonpathParserRULE_literal)
	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JsonpathParserNullLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(361)
			p.Match(JsonpathParserNullLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JsonpathParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(362)
			p.Match(JsonpathParserBooleanLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JsonpathParserDecimalLiteral, JsonpathParserHexIntegerLiteral, JsonpathParserOctalIntegerLiteral2, JsonpathParserBinaryIntegerLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(363)
			p.NumericLiteral()
		}

	case JsonpathParserStringLiteral:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(364)
			p.Match(JsonpathParserStringLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(JsonpathParserIdentifier, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *JsonpathParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, JsonpathParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(367)
		p.Match(JsonpathParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumericLiteralContext is an interface to support dynamic dispatch.
type INumericLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DecimalLiteral() antlr.TerminalNode
	HexIntegerLiteral() antlr.TerminalNode
	OctalIntegerLiteral2() antlr.TerminalNode
	BinaryIntegerLiteral() antlr.TerminalNode

	// IsNumericLiteralContext differentiates from other interfaces.
	IsNumericLiteralContext()
}

type NumericLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralContext() *NumericLiteralContext {
	var p = new(NumericLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_numericLiteral
	return p
}

func InitEmptyNumericLiteralContext(p *NumericLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_numericLiteral
}

func (*NumericLiteralContext) IsNumericLiteralContext() {}

func NewNumericLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralContext {
	var p = new(NumericLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_numericLiteral

	return p
}

func (s *NumericLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericLiteralContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(JsonpathParserDecimalLiteral, 0)
}

func (s *NumericLiteralContext) HexIntegerLiteral() antlr.TerminalNode {
	return s.GetToken(JsonpathParserHexIntegerLiteral, 0)
}

func (s *NumericLiteralContext) OctalIntegerLiteral2() antlr.TerminalNode {
	return s.GetToken(JsonpathParserOctalIntegerLiteral2, 0)
}

func (s *NumericLiteralContext) BinaryIntegerLiteral() antlr.TerminalNode {
	return s.GetToken(JsonpathParserBinaryIntegerLiteral, 0)
}

func (s *NumericLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterNumericLiteral(s)
	}
}

func (s *NumericLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitNumericLiteral(s)
	}
}

func (p *JsonpathParser) NumericLiteral() (localctx INumericLiteralContext) {
	localctx = NewNumericLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, JsonpathParserRULE_numericLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(369)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3985729650688) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *JsonpathParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *SimplePathExprContext = nil
		if localctx != nil {
			t = localctx.(*SimplePathExprContext)
		}
		return p.SimplePathExpr_Sempred(t, predIndex)

	case 7:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *JsonpathParser) SimplePathExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *JsonpathParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 27)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 26)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 25)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 24)

	case 16:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 17:
		return p.Precpred(p.GetParserRuleContext(), 22)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
