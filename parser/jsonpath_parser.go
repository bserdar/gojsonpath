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
		"", "'$'", "'@'", "'*'", "'('", "','", "')'", "':'", "'.'", "'..'",
		"'['", "']'", "'+'", "'-'", "'~'", "'!'", "'**'", "'/'", "'%'", "'<'",
		"'>'", "'<='", "'>='", "'in'", "'=='", "'!='", "'==='", "'!=='", "'&'",
		"'^'", "'|'", "'&&'", "'||'", "'?'", "'?('", "'{'", "'}'", "'...'",
		"'null'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "Ellipsis", "NullLiteral", "BooleanLiteral", "DecimalLiteral",
		"StringLiteral", "HexIntegerLiteral", "OctalIntegerLiteral2", "BinaryIntegerLiteral",
		"SP", "Identifier",
	}
	staticData.RuleNames = []string{
		"path", "selector", "arguments", "slice", "indexExpression", "expression",
		"arrayLiteral", "elementList", "objectLiteral", "propertyAssignment",
		"propertyName", "literal", "identifier", "numericLiteral",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 46, 396, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 3, 0, 31, 8,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 41, 8, 2, 10,
		2, 12, 2, 44, 9, 2, 1, 2, 3, 2, 47, 8, 2, 3, 2, 49, 8, 2, 1, 2, 1, 2, 1,
		3, 3, 3, 54, 8, 3, 1, 3, 1, 3, 3, 3, 58, 8, 3, 1, 3, 3, 3, 61, 8, 3, 1,
		3, 3, 3, 64, 8, 3, 1, 4, 1, 4, 3, 4, 68, 8, 4, 1, 5, 1, 5, 1, 5, 3, 5,
		73, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 78, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 83,
		8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 88, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 104, 8, 5, 1,
		5, 1, 5, 3, 5, 108, 8, 5, 1, 5, 1, 5, 3, 5, 112, 8, 5, 1, 5, 1, 5, 1, 5,
		3, 5, 117, 8, 5, 1, 5, 1, 5, 3, 5, 121, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 126,
		8, 5, 1, 5, 1, 5, 3, 5, 130, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 135, 8, 5, 1,
		5, 1, 5, 3, 5, 139, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 144, 8, 5, 1, 5, 1, 5,
		3, 5, 148, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 153, 8, 5, 1, 5, 1, 5, 3, 5, 157,
		8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 162, 8, 5, 1, 5, 1, 5, 3, 5, 166, 8, 5, 1,
		5, 1, 5, 1, 5, 3, 5, 171, 8, 5, 1, 5, 1, 5, 3, 5, 175, 8, 5, 1, 5, 1, 5,
		1, 5, 3, 5, 180, 8, 5, 1, 5, 1, 5, 3, 5, 184, 8, 5, 1, 5, 1, 5, 1, 5, 3,
		5, 189, 8, 5, 1, 5, 1, 5, 3, 5, 193, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 198,
		8, 5, 1, 5, 1, 5, 3, 5, 202, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 207, 8, 5, 1,
		5, 1, 5, 3, 5, 211, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 216, 8, 5, 1, 5, 1, 5,
		3, 5, 220, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 225, 8, 5, 1, 5, 1, 5, 1, 5, 3,
		5, 230, 8, 5, 1, 5, 1, 5, 3, 5, 234, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5,
		240, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5, 245, 8, 5, 1, 5, 1, 5, 3, 5, 249, 8,
		5, 1, 5, 1, 5, 3, 5, 253, 8, 5, 1, 5, 1, 5, 3, 5, 257, 8, 5, 1, 5, 5, 5,
		260, 8, 5, 10, 5, 12, 5, 263, 9, 5, 1, 5, 3, 5, 266, 8, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 3, 5, 272, 8, 5, 1, 5, 1, 5, 3, 5, 276, 8, 5, 1, 5, 1, 5, 3,
		5, 280, 8, 5, 1, 5, 1, 5, 3, 5, 284, 8, 5, 1, 5, 1, 5, 3, 5, 288, 8, 5,
		1, 5, 5, 5, 291, 8, 5, 10, 5, 12, 5, 294, 9, 5, 1, 5, 3, 5, 297, 8, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 303, 8, 5, 10, 5, 12, 5, 306, 9, 5, 1, 6,
		1, 6, 3, 6, 310, 8, 6, 1, 6, 1, 6, 3, 6, 314, 8, 6, 1, 6, 1, 6, 1, 7, 5,
		7, 319, 8, 7, 10, 7, 12, 7, 322, 9, 7, 1, 7, 3, 7, 325, 8, 7, 1, 7, 3,
		7, 328, 8, 7, 1, 7, 3, 7, 331, 8, 7, 1, 7, 4, 7, 334, 8, 7, 11, 7, 12,
		7, 335, 1, 7, 3, 7, 339, 8, 7, 1, 7, 5, 7, 342, 8, 7, 10, 7, 12, 7, 345,
		9, 7, 1, 7, 5, 7, 348, 8, 7, 10, 7, 12, 7, 351, 9, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 5, 8, 357, 8, 8, 10, 8, 12, 8, 360, 9, 8, 1, 8, 3, 8, 363, 8, 8,
		3, 8, 365, 8, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 3, 9, 379, 8, 9, 1, 10, 1, 10, 1, 10, 3, 10, 384, 8,
		10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 390, 8, 11, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 13, 0, 1, 10, 14, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 0, 6, 1, 0, 1, 3, 2, 0, 3, 3, 17, 18, 1, 0, 12, 13, 1, 0, 19, 22,
		1, 0, 24, 27, 2, 0, 40, 40, 42, 44, 484, 0, 28, 1, 0, 0, 0, 2, 34, 1, 0,
		0, 0, 4, 36, 1, 0, 0, 0, 6, 53, 1, 0, 0, 0, 8, 67, 1, 0, 0, 0, 10, 103,
		1, 0, 0, 0, 12, 307, 1, 0, 0, 0, 14, 320, 1, 0, 0, 0, 16, 352, 1, 0, 0,
		0, 18, 378, 1, 0, 0, 0, 20, 383, 1, 0, 0, 0, 22, 389, 1, 0, 0, 0, 24, 391,
		1, 0, 0, 0, 26, 393, 1, 0, 0, 0, 28, 30, 3, 10, 5, 0, 29, 31, 5, 45, 0,
		0, 30, 29, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 33,
		5, 0, 0, 1, 33, 1, 1, 0, 0, 0, 34, 35, 7, 0, 0, 0, 35, 3, 1, 0, 0, 0, 36,
		48, 5, 4, 0, 0, 37, 42, 3, 10, 5, 0, 38, 39, 5, 5, 0, 0, 39, 41, 3, 10,
		5, 0, 40, 38, 1, 0, 0, 0, 41, 44, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 42, 43,
		1, 0, 0, 0, 43, 46, 1, 0, 0, 0, 44, 42, 1, 0, 0, 0, 45, 47, 5, 5, 0, 0,
		46, 45, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 49, 1, 0, 0, 0, 48, 37, 1,
		0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51, 5, 6, 0, 0, 51,
		5, 1, 0, 0, 0, 52, 54, 3, 10, 5, 0, 53, 52, 1, 0, 0, 0, 53, 54, 1, 0, 0,
		0, 54, 55, 1, 0, 0, 0, 55, 57, 5, 7, 0, 0, 56, 58, 3, 10, 5, 0, 57, 56,
		1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 60, 1, 0, 0, 0, 59, 61, 5, 7, 0, 0,
		60, 59, 1, 0, 0, 0, 60, 61, 1, 0, 0, 0, 61, 63, 1, 0, 0, 0, 62, 64, 3,
		10, 5, 0, 63, 62, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 7, 1, 0, 0, 0, 65,
		68, 3, 6, 3, 0, 66, 68, 3, 10, 5, 0, 67, 65, 1, 0, 0, 0, 67, 66, 1, 0,
		0, 0, 68, 9, 1, 0, 0, 0, 69, 70, 6, 5, -1, 0, 70, 72, 5, 12, 0, 0, 71,
		73, 5, 45, 0, 0, 72, 71, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 74, 1, 0,
		0, 0, 74, 104, 3, 10, 5, 23, 75, 77, 5, 13, 0, 0, 76, 78, 5, 45, 0, 0,
		77, 76, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 104, 3,
		10, 5, 22, 80, 82, 5, 14, 0, 0, 81, 83, 5, 45, 0, 0, 82, 81, 1, 0, 0, 0,
		82, 83, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 104, 3, 10, 5, 21, 85, 87,
		5, 15, 0, 0, 86, 88, 5, 45, 0, 0, 87, 86, 1, 0, 0, 0, 87, 88, 1, 0, 0,
		0, 88, 89, 1, 0, 0, 0, 89, 104, 3, 10, 5, 20, 90, 104, 3, 22, 11, 0, 91,
		104, 3, 12, 6, 0, 92, 104, 3, 16, 8, 0, 93, 104, 3, 24, 12, 0, 94, 104,
		3, 2, 1, 0, 95, 96, 5, 4, 0, 0, 96, 97, 3, 10, 5, 0, 97, 98, 5, 6, 0, 0,
		98, 104, 1, 0, 0, 0, 99, 100, 5, 34, 0, 0, 100, 101, 3, 10, 5, 0, 101,
		102, 5, 6, 0, 0, 102, 104, 1, 0, 0, 0, 103, 69, 1, 0, 0, 0, 103, 75, 1,
		0, 0, 0, 103, 80, 1, 0, 0, 0, 103, 85, 1, 0, 0, 0, 103, 90, 1, 0, 0, 0,
		103, 91, 1, 0, 0, 0, 103, 92, 1, 0, 0, 0, 103, 93, 1, 0, 0, 0, 103, 94,
		1, 0, 0, 0, 103, 95, 1, 0, 0, 0, 103, 99, 1, 0, 0, 0, 104, 304, 1, 0, 0,
		0, 105, 107, 10, 29, 0, 0, 106, 108, 5, 45, 0, 0, 107, 106, 1, 0, 0, 0,
		107, 108, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 111, 5, 8, 0, 0, 110,
		112, 5, 45, 0, 0, 111, 110, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 113,
		1, 0, 0, 0, 113, 303, 3, 10, 5, 30, 114, 116, 10, 28, 0, 0, 115, 117, 5,
		45, 0, 0, 116, 115, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 118, 1, 0, 0,
		0, 118, 120, 5, 9, 0, 0, 119, 121, 5, 45, 0, 0, 120, 119, 1, 0, 0, 0, 120,
		121, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 303, 3, 10, 5, 29, 123, 125,
		10, 19, 0, 0, 124, 126, 5, 45, 0, 0, 125, 124, 1, 0, 0, 0, 125, 126, 1,
		0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 129, 5, 16, 0, 0, 128, 130, 5, 45,
		0, 0, 129, 128, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0,
		131, 303, 3, 10, 5, 19, 132, 134, 10, 18, 0, 0, 133, 135, 5, 45, 0, 0,
		134, 133, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 136, 1, 0, 0, 0, 136,
		138, 7, 1, 0, 0, 137, 139, 5, 45, 0, 0, 138, 137, 1, 0, 0, 0, 138, 139,
		1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 303, 3, 10, 5, 19, 141, 143, 10,
		17, 0, 0, 142, 144, 5, 45, 0, 0, 143, 142, 1, 0, 0, 0, 143, 144, 1, 0,
		0, 0, 144, 145, 1, 0, 0, 0, 145, 147, 7, 2, 0, 0, 146, 148, 5, 45, 0, 0,
		147, 146, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149,
		303, 3, 10, 5, 18, 150, 152, 10, 16, 0, 0, 151, 153, 5, 45, 0, 0, 152,
		151, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 156,
		7, 3, 0, 0, 155, 157, 5, 45, 0, 0, 156, 155, 1, 0, 0, 0, 156, 157, 1, 0,
		0, 0, 157, 158, 1, 0, 0, 0, 158, 303, 3, 10, 5, 17, 159, 161, 10, 15, 0,
		0, 160, 162, 5, 45, 0, 0, 161, 160, 1, 0, 0, 0, 161, 162, 1, 0, 0, 0, 162,
		163, 1, 0, 0, 0, 163, 165, 5, 23, 0, 0, 164, 166, 5, 45, 0, 0, 165, 164,
		1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 303, 3, 10,
		5, 16, 168, 170, 10, 14, 0, 0, 169, 171, 5, 45, 0, 0, 170, 169, 1, 0, 0,
		0, 170, 171, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 174, 7, 4, 0, 0, 173,
		175, 5, 45, 0, 0, 174, 173, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 176,
		1, 0, 0, 0, 176, 303, 3, 10, 5, 15, 177, 179, 10, 13, 0, 0, 178, 180, 5,
		45, 0, 0, 179, 178, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 181, 1, 0, 0,
		0, 181, 183, 5, 28, 0, 0, 182, 184, 5, 45, 0, 0, 183, 182, 1, 0, 0, 0,
		183, 184, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 303, 3, 10, 5, 14, 186,
		188, 10, 12, 0, 0, 187, 189, 5, 45, 0, 0, 188, 187, 1, 0, 0, 0, 188, 189,
		1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 192, 5, 29, 0, 0, 191, 193, 5, 45,
		0, 0, 192, 191, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0,
		194, 303, 3, 10, 5, 13, 195, 197, 10, 11, 0, 0, 196, 198, 5, 45, 0, 0,
		197, 196, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199,
		201, 5, 30, 0, 0, 200, 202, 5, 45, 0, 0, 201, 200, 1, 0, 0, 0, 201, 202,
		1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 303, 3, 10, 5, 12, 204, 206, 10,
		10, 0, 0, 205, 207, 5, 45, 0, 0, 206, 205, 1, 0, 0, 0, 206, 207, 1, 0,
		0, 0, 207, 208, 1, 0, 0, 0, 208, 210, 5, 31, 0, 0, 209, 211, 5, 45, 0,
		0, 210, 209, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212,
		303, 3, 10, 5, 11, 213, 215, 10, 9, 0, 0, 214, 216, 5, 45, 0, 0, 215, 214,
		1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 219, 5, 32,
		0, 0, 218, 220, 5, 45, 0, 0, 219, 218, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0,
		220, 221, 1, 0, 0, 0, 221, 303, 3, 10, 5, 10, 222, 224, 10, 8, 0, 0, 223,
		225, 5, 45, 0, 0, 224, 223, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 226,
		1, 0, 0, 0, 226, 227, 5, 33, 0, 0, 227, 229, 3, 10, 5, 0, 228, 230, 5,
		45, 0, 0, 229, 228, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 231, 1, 0, 0,
		0, 231, 233, 5, 7, 0, 0, 232, 234, 5, 45, 0, 0, 233, 232, 1, 0, 0, 0, 233,
		234, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 236, 3, 10, 5, 9, 236, 303,
		1, 0, 0, 0, 237, 239, 10, 27, 0, 0, 238, 240, 5, 45, 0, 0, 239, 238, 1,
		0, 0, 0, 239, 240, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 303, 5, 9, 0,
		0, 242, 244, 10, 26, 0, 0, 243, 245, 5, 45, 0, 0, 244, 243, 1, 0, 0, 0,
		244, 245, 1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 248, 5, 10, 0, 0, 247,
		249, 5, 45, 0, 0, 248, 247, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 250,
		1, 0, 0, 0, 250, 261, 3, 8, 4, 0, 251, 253, 5, 45, 0, 0, 252, 251, 1, 0,
		0, 0, 252, 253, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 256, 5, 5, 0, 0,
		255, 257, 5, 45, 0, 0, 256, 255, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257,
		258, 1, 0, 0, 0, 258, 260, 3, 8, 4, 0, 259, 252, 1, 0, 0, 0, 260, 263,
		1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 265, 1, 0,
		0, 0, 263, 261, 1, 0, 0, 0, 264, 266, 5, 45, 0, 0, 265, 264, 1, 0, 0, 0,
		265, 266, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 268, 5, 11, 0, 0, 268,
		303, 1, 0, 0, 0, 269, 271, 10, 25, 0, 0, 270, 272, 5, 45, 0, 0, 271, 270,
		1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 273, 1, 0, 0, 0, 273, 275, 5, 9,
		0, 0, 274, 276, 5, 45, 0, 0, 275, 274, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0,
		276, 277, 1, 0, 0, 0, 277, 279, 5, 10, 0, 0, 278, 280, 5, 45, 0, 0, 279,
		278, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 292,
		3, 8, 4, 0, 282, 284, 5, 45, 0, 0, 283, 282, 1, 0, 0, 0, 283, 284, 1, 0,
		0, 0, 284, 285, 1, 0, 0, 0, 285, 287, 5, 5, 0, 0, 286, 288, 5, 45, 0, 0,
		287, 286, 1, 0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289,
		291, 3, 8, 4, 0, 290, 283, 1, 0, 0, 0, 291, 294, 1, 0, 0, 0, 292, 290,
		1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293, 296, 1, 0, 0, 0, 294, 292, 1, 0,
		0, 0, 295, 297, 5, 45, 0, 0, 296, 295, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0,
		297, 298, 1, 0, 0, 0, 298, 299, 5, 11, 0, 0, 299, 303, 1, 0, 0, 0, 300,
		301, 10, 24, 0, 0, 301, 303, 3, 4, 2, 0, 302, 105, 1, 0, 0, 0, 302, 114,
		1, 0, 0, 0, 302, 123, 1, 0, 0, 0, 302, 132, 1, 0, 0, 0, 302, 141, 1, 0,
		0, 0, 302, 150, 1, 0, 0, 0, 302, 159, 1, 0, 0, 0, 302, 168, 1, 0, 0, 0,
		302, 177, 1, 0, 0, 0, 302, 186, 1, 0, 0, 0, 302, 195, 1, 0, 0, 0, 302,
		204, 1, 0, 0, 0, 302, 213, 1, 0, 0, 0, 302, 222, 1, 0, 0, 0, 302, 237,
		1, 0, 0, 0, 302, 242, 1, 0, 0, 0, 302, 269, 1, 0, 0, 0, 302, 300, 1, 0,
		0, 0, 303, 306, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0,
		305, 11, 1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 307, 309, 5, 10, 0, 0, 308,
		310, 5, 45, 0, 0, 309, 308, 1, 0, 0, 0, 309, 310, 1, 0, 0, 0, 310, 311,
		1, 0, 0, 0, 311, 313, 3, 14, 7, 0, 312, 314, 5, 45, 0, 0, 313, 312, 1,
		0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 316, 5, 11, 0,
		0, 316, 13, 1, 0, 0, 0, 317, 319, 5, 5, 0, 0, 318, 317, 1, 0, 0, 0, 319,
		322, 1, 0, 0, 0, 320, 318, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321, 324,
		1, 0, 0, 0, 322, 320, 1, 0, 0, 0, 323, 325, 5, 45, 0, 0, 324, 323, 1, 0,
		0, 0, 324, 325, 1, 0, 0, 0, 325, 327, 1, 0, 0, 0, 326, 328, 3, 10, 5, 0,
		327, 326, 1, 0, 0, 0, 327, 328, 1, 0, 0, 0, 328, 343, 1, 0, 0, 0, 329,
		331, 5, 45, 0, 0, 330, 329, 1, 0, 0, 0, 330, 331, 1, 0, 0, 0, 331, 333,
		1, 0, 0, 0, 332, 334, 5, 5, 0, 0, 333, 332, 1, 0, 0, 0, 334, 335, 1, 0,
		0, 0, 335, 333, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336, 338, 1, 0, 0, 0,
		337, 339, 5, 45, 0, 0, 338, 337, 1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339,
		340, 1, 0, 0, 0, 340, 342, 3, 10, 5, 0, 341, 330, 1, 0, 0, 0, 342, 345,
		1, 0, 0, 0, 343, 341, 1, 0, 0, 0, 343, 344, 1, 0, 0, 0, 344, 349, 1, 0,
		0, 0, 345, 343, 1, 0, 0, 0, 346, 348, 5, 5, 0, 0, 347, 346, 1, 0, 0, 0,
		348, 351, 1, 0, 0, 0, 349, 347, 1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350,
		15, 1, 0, 0, 0, 351, 349, 1, 0, 0, 0, 352, 364, 5, 35, 0, 0, 353, 358,
		3, 18, 9, 0, 354, 355, 5, 5, 0, 0, 355, 357, 3, 18, 9, 0, 356, 354, 1,
		0, 0, 0, 357, 360, 1, 0, 0, 0, 358, 356, 1, 0, 0, 0, 358, 359, 1, 0, 0,
		0, 359, 362, 1, 0, 0, 0, 360, 358, 1, 0, 0, 0, 361, 363, 5, 5, 0, 0, 362,
		361, 1, 0, 0, 0, 362, 363, 1, 0, 0, 0, 363, 365, 1, 0, 0, 0, 364, 353,
		1, 0, 0, 0, 364, 365, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 367, 5, 36,
		0, 0, 367, 17, 1, 0, 0, 0, 368, 369, 3, 20, 10, 0, 369, 370, 5, 7, 0, 0,
		370, 371, 3, 10, 5, 0, 371, 379, 1, 0, 0, 0, 372, 373, 5, 10, 0, 0, 373,
		374, 3, 10, 5, 0, 374, 375, 5, 11, 0, 0, 375, 376, 5, 7, 0, 0, 376, 377,
		3, 10, 5, 0, 377, 379, 1, 0, 0, 0, 378, 368, 1, 0, 0, 0, 378, 372, 1, 0,
		0, 0, 379, 19, 1, 0, 0, 0, 380, 384, 3, 24, 12, 0, 381, 384, 5, 41, 0,
		0, 382, 384, 3, 26, 13, 0, 383, 380, 1, 0, 0, 0, 383, 381, 1, 0, 0, 0,
		383, 382, 1, 0, 0, 0, 384, 21, 1, 0, 0, 0, 385, 390, 5, 38, 0, 0, 386,
		390, 5, 39, 0, 0, 387, 390, 3, 26, 13, 0, 388, 390, 5, 41, 0, 0, 389, 385,
		1, 0, 0, 0, 389, 386, 1, 0, 0, 0, 389, 387, 1, 0, 0, 0, 389, 388, 1, 0,
		0, 0, 390, 23, 1, 0, 0, 0, 391, 392, 5, 46, 0, 0, 392, 25, 1, 0, 0, 0,
		393, 394, 7, 5, 0, 0, 394, 27, 1, 0, 0, 0, 75, 30, 42, 46, 48, 53, 57,
		60, 63, 67, 72, 77, 82, 87, 103, 107, 111, 116, 120, 125, 129, 134, 138,
		143, 147, 152, 156, 161, 165, 170, 174, 179, 183, 188, 192, 197, 201, 206,
		210, 215, 219, 224, 229, 233, 239, 244, 248, 252, 256, 261, 265, 271, 275,
		279, 283, 287, 292, 296, 302, 304, 309, 313, 320, 324, 327, 330, 335, 338,
		343, 349, 358, 362, 364, 378, 383, 389,
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
	JsonpathParserT__33                = 34
	JsonpathParserT__34                = 35
	JsonpathParserT__35                = 36
	JsonpathParserEllipsis             = 37
	JsonpathParserNullLiteral          = 38
	JsonpathParserBooleanLiteral       = 39
	JsonpathParserDecimalLiteral       = 40
	JsonpathParserStringLiteral        = 41
	JsonpathParserHexIntegerLiteral    = 42
	JsonpathParserOctalIntegerLiteral2 = 43
	JsonpathParserBinaryIntegerLiteral = 44
	JsonpathParserSP                   = 45
	JsonpathParserIdentifier           = 46
)

// JsonpathParser rules.
const (
	JsonpathParserRULE_path               = 0
	JsonpathParserRULE_selector           = 1
	JsonpathParserRULE_arguments          = 2
	JsonpathParserRULE_slice              = 3
	JsonpathParserRULE_indexExpression    = 4
	JsonpathParserRULE_expression         = 5
	JsonpathParserRULE_arrayLiteral       = 6
	JsonpathParserRULE_elementList        = 7
	JsonpathParserRULE_objectLiteral      = 8
	JsonpathParserRULE_propertyAssignment = 9
	JsonpathParserRULE_propertyName       = 10
	JsonpathParserRULE_literal            = 11
	JsonpathParserRULE_identifier         = 12
	JsonpathParserRULE_numericLiteral     = 13
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
		p.SetState(28)
		p.expression(0)
	}
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JsonpathParserSP {
		{
			p.SetState(29)
			p.Match(JsonpathParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(32)
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
	p.EnterRule(localctx, 2, JsonpathParserRULE_selector)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14) != 0) {
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
	p.EnterRule(localctx, 4, JsonpathParserRULE_arguments)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.Match(JsonpathParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&105329778029598) != 0 {
		{
			p.SetState(37)
			p.expression(0)
		}
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(38)
					p.Match(JsonpathParserT__4)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(39)
					p.expression(0)
				}

			}
			p.SetState(44)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonpathParserT__4 {
			{
				p.SetState(45)
				p.Match(JsonpathParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(50)
		p.Match(JsonpathParserT__5)
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
	p.EnterRule(localctx, 6, JsonpathParserRULE_slice)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&105329778029598) != 0 {
		{
			p.SetState(52)
			p.expression(0)
		}

	}
	{
		p.SetState(55)
		p.Match(JsonpathParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(56)
			p.expression(0)
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JsonpathParserT__6 {
		{
			p.SetState(59)
			p.Match(JsonpathParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&105329778029598) != 0 {
		{
			p.SetState(62)
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
	p.EnterRule(localctx, 8, JsonpathParserRULE_indexExpression)
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Slice()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
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

type TernaryExpressionContext struct {
	ExpressionContext
}

func NewTernaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TernaryExpressionContext {
	var p = new(TernaryExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *TernaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryExpressionContext) AllExpression() []IExpressionContext {
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

func (s *TernaryExpressionContext) Expression(i int) IExpressionContext {
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

func (s *TernaryExpressionContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserSP)
}

func (s *TernaryExpressionContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, i)
}

func (s *TernaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterTernaryExpression(s)
	}
}

func (s *TernaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitTernaryExpression(s)
	}
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

type ObjectLiteralExpressionContext struct {
	ExpressionContext
}

func NewObjectLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectLiteralExpressionContext {
	var p = new(ObjectLiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ObjectLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectLiteralExpressionContext) ObjectLiteral() IObjectLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectLiteralContext)
}

func (s *ObjectLiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterObjectLiteralExpression(s)
	}
}

func (s *ObjectLiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitObjectLiteralExpression(s)
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
	_startState := 10
	p.EnterRecursionRule(localctx, 10, JsonpathParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JsonpathParserT__11:
		localctx = NewUnaryPlusExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(70)
			p.Match(JsonpathParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonpathParserSP {
			{
				p.SetState(71)
				p.Match(JsonpathParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(74)
			p.expression(23)
		}

	case JsonpathParserT__12:
		localctx = NewUnaryMinusExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(75)
			p.Match(JsonpathParserT__12)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonpathParserSP {
			{
				p.SetState(76)
				p.Match(JsonpathParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(79)
			p.expression(22)
		}

	case JsonpathParserT__13:
		localctx = NewBitNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(80)
			p.Match(JsonpathParserT__13)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonpathParserSP {
			{
				p.SetState(81)
				p.Match(JsonpathParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(84)
			p.expression(21)
		}

	case JsonpathParserT__14:
		localctx = NewNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(85)
			p.Match(JsonpathParserT__14)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonpathParserSP {
			{
				p.SetState(86)
				p.Match(JsonpathParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(89)
			p.expression(20)
		}

	case JsonpathParserNullLiteral, JsonpathParserBooleanLiteral, JsonpathParserDecimalLiteral, JsonpathParserStringLiteral, JsonpathParserHexIntegerLiteral, JsonpathParserOctalIntegerLiteral2, JsonpathParserBinaryIntegerLiteral:
		localctx = NewLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(90)
			p.Literal()
		}

	case JsonpathParserT__9:
		localctx = NewArrayLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(91)
			p.ArrayLiteral()
		}

	case JsonpathParserT__34:
		localctx = NewObjectLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(92)
			p.ObjectLiteral()
		}

	case JsonpathParserIdentifier:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(93)
			p.Identifier()
		}

	case JsonpathParserT__0, JsonpathParserT__1, JsonpathParserT__2:
		localctx = NewSelectorExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(94)
			p.Selector()
		}

	case JsonpathParserT__3:
		localctx = NewParenthesizedExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(95)
			p.Match(JsonpathParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)
			p.expression(0)
		}
		{
			p.SetState(97)
			p.Match(JsonpathParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JsonpathParserT__33:
		localctx = NewFilterExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(99)
			p.Match(JsonpathParserT__33)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.expression(0)
		}
		{
			p.SetState(101)
			p.Match(JsonpathParserT__5)
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
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(302)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 57, p.GetParserRuleContext()) {
			case 1:
				localctx = NewChainExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(105)

				if !(p.Precpred(p.GetParserRuleContext(), 29)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 29)", ""))
					goto errorExit
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
					p.Match(JsonpathParserT__7)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(111)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(110)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(113)
					p.expression(30)
				}

			case 2:
				localctx = NewRecursiveDescentExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(114)

				if !(p.Precpred(p.GetParserRuleContext(), 28)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 28)", ""))
					goto errorExit
				}
				p.SetState(116)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(115)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(118)
					p.Match(JsonpathParserT__8)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(120)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(119)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(122)
					p.expression(29)
				}

			case 3:
				localctx = NewPowerExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(123)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				p.SetState(125)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(124)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(127)
					p.Match(JsonpathParserT__15)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(129)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(128)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(131)
					p.expression(19)
				}

			case 4:
				localctx = NewMultiplicativeExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(132)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				p.SetState(134)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(133)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(136)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&393224) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(138)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(137)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(140)
					p.expression(19)
				}

			case 5:
				localctx = NewAdditiveExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(141)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				p.SetState(143)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(142)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(145)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JsonpathParserT__11 || _la == JsonpathParserT__12) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(147)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(146)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(149)
					p.expression(18)
				}

			case 6:
				localctx = NewRelationalExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(150)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				p.SetState(152)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(151)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(154)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7864320) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(156)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(155)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(158)
					p.expression(17)
				}

			case 7:
				localctx = NewInExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(159)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				p.SetState(161)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(160)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(163)
					p.Match(JsonpathParserT__22)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(165)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(164)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(167)
					p.expression(16)
				}

			case 8:
				localctx = NewEqualityExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(168)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				p.SetState(170)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(169)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(172)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&251658240) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(174)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(173)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(176)
					p.expression(15)
				}

			case 9:
				localctx = NewBitAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(177)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				p.SetState(179)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(178)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(181)
					p.Match(JsonpathParserT__27)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(183)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(182)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(185)
					p.expression(14)
				}

			case 10:
				localctx = NewBitXOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(186)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				p.SetState(188)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(187)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(190)
					p.Match(JsonpathParserT__28)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(192)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(191)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(194)
					p.expression(13)
				}

			case 11:
				localctx = NewBitOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(195)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				p.SetState(197)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(196)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(199)
					p.Match(JsonpathParserT__29)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(201)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(200)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(203)
					p.expression(12)
				}

			case 12:
				localctx = NewLogicalAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(204)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				p.SetState(206)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(205)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(208)
					p.Match(JsonpathParserT__30)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(210)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(209)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(212)
					p.expression(11)
				}

			case 13:
				localctx = NewLogicalOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(213)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				p.SetState(215)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(214)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(217)
					p.Match(JsonpathParserT__31)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(219)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(218)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(221)
					p.expression(10)
				}

			case 14:
				localctx = NewTernaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(222)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				p.SetState(224)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(223)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(226)
					p.Match(JsonpathParserT__32)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(227)
					p.expression(0)
				}
				p.SetState(229)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(228)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(231)
					p.Match(JsonpathParserT__6)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(233)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(232)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(235)
					p.expression(9)
				}

			case 15:
				localctx = NewRecursiveDescentTermExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(237)

				if !(p.Precpred(p.GetParserRuleContext(), 27)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 27)", ""))
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
					p.Match(JsonpathParserT__8)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 16:
				localctx = NewMemberIndexExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(242)

				if !(p.Precpred(p.GetParserRuleContext(), 26)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 26)", ""))
					goto errorExit
				}
				p.SetState(244)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(243)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(246)
					p.Match(JsonpathParserT__9)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
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
					p.IndexExpression()
				}
				p.SetState(261)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
				for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					if _alt == 1 {
						p.SetState(252)
						p.GetErrorHandler().Sync(p)
						if p.HasError() {
							goto errorExit
						}
						_la = p.GetTokenStream().LA(1)

						if _la == JsonpathParserSP {
							{
								p.SetState(251)
								p.Match(JsonpathParserSP)
								if p.HasError() {
									// Recognition error - abort rule
									goto errorExit
								}
							}

						}
						{
							p.SetState(254)
							p.Match(JsonpathParserT__4)
							if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
							}
						}
						p.SetState(256)
						p.GetErrorHandler().Sync(p)
						if p.HasError() {
							goto errorExit
						}
						_la = p.GetTokenStream().LA(1)

						if _la == JsonpathParserSP {
							{
								p.SetState(255)
								p.Match(JsonpathParserSP)
								if p.HasError() {
									// Recognition error - abort rule
									goto errorExit
								}
							}

						}
						{
							p.SetState(258)
							p.IndexExpression()
						}

					}
					p.SetState(263)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 48, p.GetParserRuleContext())
					if p.HasError() {
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
					p.Match(JsonpathParserT__10)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 17:
				localctx = NewRecursiveDescentMemberIndexExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(269)

				if !(p.Precpred(p.GetParserRuleContext(), 25)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 25)", ""))
					goto errorExit
				}
				p.SetState(271)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(270)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(273)
					p.Match(JsonpathParserT__8)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(275)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(274)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(277)
					p.Match(JsonpathParserT__9)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(279)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(278)
						p.Match(JsonpathParserSP)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(281)
					p.IndexExpression()
				}
				p.SetState(292)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
				for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					if _alt == 1 {
						p.SetState(283)
						p.GetErrorHandler().Sync(p)
						if p.HasError() {
							goto errorExit
						}
						_la = p.GetTokenStream().LA(1)

						if _la == JsonpathParserSP {
							{
								p.SetState(282)
								p.Match(JsonpathParserSP)
								if p.HasError() {
									// Recognition error - abort rule
									goto errorExit
								}
							}

						}
						{
							p.SetState(285)
							p.Match(JsonpathParserT__4)
							if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
							}
						}
						p.SetState(287)
						p.GetErrorHandler().Sync(p)
						if p.HasError() {
							goto errorExit
						}
						_la = p.GetTokenStream().LA(1)

						if _la == JsonpathParserSP {
							{
								p.SetState(286)
								p.Match(JsonpathParserSP)
								if p.HasError() {
									// Recognition error - abort rule
									goto errorExit
								}
							}

						}
						{
							p.SetState(289)
							p.IndexExpression()
						}

					}
					p.SetState(294)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 55, p.GetParserRuleContext())
					if p.HasError() {
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
					p.Match(JsonpathParserT__10)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 18:
				localctx = NewArgumentsExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_expression)
				p.SetState(300)

				if !(p.Precpred(p.GetParserRuleContext(), 24)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 24)", ""))
					goto errorExit
				}
				{
					p.SetState(301)
					p.Arguments()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 58, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 12, JsonpathParserRULE_arrayLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.Match(JsonpathParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(309)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 59, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(308)
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
		p.SetState(311)
		p.ElementList()
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JsonpathParserSP {
		{
			p.SetState(312)
			p.Match(JsonpathParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(315)
		p.Match(JsonpathParserT__10)
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
	p.EnterRule(localctx, 14, JsonpathParserRULE_elementList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(317)
				p.Match(JsonpathParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 61, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 62, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(323)
			p.Match(JsonpathParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&105329778029598) != 0 {
		{
			p.SetState(326)
			p.expression(0)
		}

	}
	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 67, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(330)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == JsonpathParserSP {
				{
					p.SetState(329)
					p.Match(JsonpathParserSP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(333)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == JsonpathParserT__4 {
				{
					p.SetState(332)
					p.Match(JsonpathParserT__4)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				p.SetState(335)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(338)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == JsonpathParserSP {
				{
					p.SetState(337)
					p.Match(JsonpathParserSP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(340)
				p.expression(0)
			}

		}
		p.SetState(345)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 67, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == JsonpathParserT__4 {
		{
			p.SetState(346)
			p.Match(JsonpathParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(351)
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

// IObjectLiteralContext is an interface to support dynamic dispatch.
type IObjectLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllPropertyAssignment() []IPropertyAssignmentContext
	PropertyAssignment(i int) IPropertyAssignmentContext

	// IsObjectLiteralContext differentiates from other interfaces.
	IsObjectLiteralContext()
}

type ObjectLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectLiteralContext() *ObjectLiteralContext {
	var p = new(ObjectLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_objectLiteral
	return p
}

func InitEmptyObjectLiteralContext(p *ObjectLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_objectLiteral
}

func (*ObjectLiteralContext) IsObjectLiteralContext() {}

func NewObjectLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectLiteralContext {
	var p = new(ObjectLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_objectLiteral

	return p
}

func (s *ObjectLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectLiteralContext) AllPropertyAssignment() []IPropertyAssignmentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPropertyAssignmentContext); ok {
			len++
		}
	}

	tst := make([]IPropertyAssignmentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPropertyAssignmentContext); ok {
			tst[i] = t.(IPropertyAssignmentContext)
			i++
		}
	}

	return tst
}

func (s *ObjectLiteralContext) PropertyAssignment(i int) IPropertyAssignmentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyAssignmentContext); ok {
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

	return t.(IPropertyAssignmentContext)
}

func (s *ObjectLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterObjectLiteral(s)
	}
}

func (s *ObjectLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitObjectLiteral(s)
	}
}

func (p *JsonpathParser) ObjectLiteral() (localctx IObjectLiteralContext) {
	localctx = NewObjectLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JsonpathParserRULE_objectLiteral)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(352)
		p.Match(JsonpathParserT__34)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&104453604639744) != 0 {
		{
			p.SetState(353)
			p.PropertyAssignment()
		}
		p.SetState(358)
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
				{
					p.SetState(354)
					p.Match(JsonpathParserT__4)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(355)
					p.PropertyAssignment()
				}

			}
			p.SetState(360)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 69, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(362)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonpathParserT__4 {
			{
				p.SetState(361)
				p.Match(JsonpathParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(366)
		p.Match(JsonpathParserT__35)
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

// IPropertyAssignmentContext is an interface to support dynamic dispatch.
type IPropertyAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPropertyAssignmentContext differentiates from other interfaces.
	IsPropertyAssignmentContext()
}

type PropertyAssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyAssignmentContext() *PropertyAssignmentContext {
	var p = new(PropertyAssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_propertyAssignment
	return p
}

func InitEmptyPropertyAssignmentContext(p *PropertyAssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_propertyAssignment
}

func (*PropertyAssignmentContext) IsPropertyAssignmentContext() {}

func NewPropertyAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyAssignmentContext {
	var p = new(PropertyAssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_propertyAssignment

	return p
}

func (s *PropertyAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyAssignmentContext) CopyAll(ctx *PropertyAssignmentContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PropertyAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PropertyExpressionAssignmentContext struct {
	PropertyAssignmentContext
}

func NewPropertyExpressionAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropertyExpressionAssignmentContext {
	var p = new(PropertyExpressionAssignmentContext)

	InitEmptyPropertyAssignmentContext(&p.PropertyAssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*PropertyAssignmentContext))

	return p
}

func (s *PropertyExpressionAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyExpressionAssignmentContext) PropertyName() IPropertyNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *PropertyExpressionAssignmentContext) Expression() IExpressionContext {
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

func (s *PropertyExpressionAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterPropertyExpressionAssignment(s)
	}
}

func (s *PropertyExpressionAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitPropertyExpressionAssignment(s)
	}
}

type ComputedPropertyExpressionAssignmentContext struct {
	PropertyAssignmentContext
}

func NewComputedPropertyExpressionAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComputedPropertyExpressionAssignmentContext {
	var p = new(ComputedPropertyExpressionAssignmentContext)

	InitEmptyPropertyAssignmentContext(&p.PropertyAssignmentContext)
	p.parser = parser
	p.CopyAll(ctx.(*PropertyAssignmentContext))

	return p
}

func (s *ComputedPropertyExpressionAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComputedPropertyExpressionAssignmentContext) AllExpression() []IExpressionContext {
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

func (s *ComputedPropertyExpressionAssignmentContext) Expression(i int) IExpressionContext {
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

func (s *ComputedPropertyExpressionAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterComputedPropertyExpressionAssignment(s)
	}
}

func (s *ComputedPropertyExpressionAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitComputedPropertyExpressionAssignment(s)
	}
}

func (p *JsonpathParser) PropertyAssignment() (localctx IPropertyAssignmentContext) {
	localctx = NewPropertyAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JsonpathParserRULE_propertyAssignment)
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JsonpathParserDecimalLiteral, JsonpathParserStringLiteral, JsonpathParserHexIntegerLiteral, JsonpathParserOctalIntegerLiteral2, JsonpathParserBinaryIntegerLiteral, JsonpathParserIdentifier:
		localctx = NewPropertyExpressionAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(368)
			p.PropertyName()
		}
		{
			p.SetState(369)
			p.Match(JsonpathParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(370)
			p.expression(0)
		}

	case JsonpathParserT__9:
		localctx = NewComputedPropertyExpressionAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(372)
			p.Match(JsonpathParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)
			p.expression(0)
		}
		{
			p.SetState(374)
			p.Match(JsonpathParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(375)
			p.Match(JsonpathParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(376)
			p.expression(0)
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

// IPropertyNameContext is an interface to support dynamic dispatch.
type IPropertyNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	StringLiteral() antlr.TerminalNode
	NumericLiteral() INumericLiteralContext

	// IsPropertyNameContext differentiates from other interfaces.
	IsPropertyNameContext()
}

type PropertyNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyNameContext() *PropertyNameContext {
	var p = new(PropertyNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_propertyName
	return p
}

func InitEmptyPropertyNameContext(p *PropertyNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonpathParserRULE_propertyName
}

func (*PropertyNameContext) IsPropertyNameContext() {}

func NewPropertyNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyNameContext {
	var p = new(PropertyNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_propertyName

	return p
}

func (s *PropertyNameContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyNameContext) Identifier() IIdentifierContext {
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

func (s *PropertyNameContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(JsonpathParserStringLiteral, 0)
}

func (s *PropertyNameContext) NumericLiteral() INumericLiteralContext {
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

func (s *PropertyNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterPropertyName(s)
	}
}

func (s *PropertyNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitPropertyName(s)
	}
}

func (p *JsonpathParser) PropertyName() (localctx IPropertyNameContext) {
	localctx = NewPropertyNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, JsonpathParserRULE_propertyName)
	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JsonpathParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(380)
			p.Identifier()
		}

	case JsonpathParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(381)
			p.Match(JsonpathParserStringLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JsonpathParserDecimalLiteral, JsonpathParserHexIntegerLiteral, JsonpathParserOctalIntegerLiteral2, JsonpathParserBinaryIntegerLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(382)
			p.NumericLiteral()
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
	p.EnterRule(localctx, 22, JsonpathParserRULE_literal)
	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JsonpathParserNullLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(385)
			p.Match(JsonpathParserNullLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JsonpathParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(386)
			p.Match(JsonpathParserBooleanLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JsonpathParserDecimalLiteral, JsonpathParserHexIntegerLiteral, JsonpathParserOctalIntegerLiteral2, JsonpathParserBinaryIntegerLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(387)
			p.NumericLiteral()
		}

	case JsonpathParserStringLiteral:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(388)
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
	p.EnterRule(localctx, 24, JsonpathParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(391)
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
	p.EnterRule(localctx, 26, JsonpathParserRULE_numericLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(393)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&31885837205504) != 0) {
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
	case 5:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *JsonpathParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 29)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 28)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 27)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 26)

	case 16:
		return p.Precpred(p.GetParserRuleContext(), 25)

	case 17:
		return p.Precpred(p.GetParserRuleContext(), 24)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
