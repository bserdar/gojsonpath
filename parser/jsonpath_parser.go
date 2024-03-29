// Code generated from Jsonpath.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Jsonpath

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 49, 436,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 3, 2, 5, 2, 53, 10, 2, 3, 2, 7, 2, 56,
	10, 2, 12, 2, 14, 2, 59, 11, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 5, 7, 83, 10, 7, 3, 8, 3, 8, 5, 8, 87, 10, 8, 3, 8,
	3, 8, 5, 8, 91, 10, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 7, 9, 98, 10, 9, 12,
	9, 14, 9, 101, 11, 9, 3, 10, 3, 10, 5, 10, 105, 10, 10, 3, 11, 5, 11, 108,
	10, 11, 3, 11, 3, 11, 5, 11, 112, 10, 11, 3, 11, 5, 11, 115, 10, 11, 3,
	11, 5, 11, 118, 10, 11, 3, 12, 3, 12, 5, 12, 122, 10, 12, 3, 13, 3, 13,
	3, 13, 3, 13, 7, 13, 128, 10, 13, 12, 13, 14, 13, 131, 11, 13, 3, 13, 5,
	13, 134, 10, 13, 5, 13, 136, 10, 13, 3, 13, 3, 13, 3, 14, 5, 14, 141, 10,
	14, 3, 14, 3, 14, 5, 14, 145, 10, 14, 3, 15, 3, 15, 3, 15, 7, 15, 150,
	10, 15, 12, 15, 14, 15, 153, 11, 15, 3, 16, 3, 16, 3, 16, 5, 16, 158, 10,
	16, 3, 16, 3, 16, 3, 16, 5, 16, 163, 10, 16, 3, 16, 3, 16, 3, 16, 5, 16,
	168, 10, 16, 3, 16, 3, 16, 3, 16, 5, 16, 173, 10, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 5, 16, 189, 10, 16, 3, 16, 3, 16, 5, 16, 193, 10, 16, 3, 16, 3,
	16, 5, 16, 197, 10, 16, 3, 16, 3, 16, 3, 16, 5, 16, 202, 10, 16, 3, 16,
	3, 16, 5, 16, 206, 10, 16, 3, 16, 3, 16, 3, 16, 5, 16, 211, 10, 16, 3,
	16, 3, 16, 5, 16, 215, 10, 16, 3, 16, 3, 16, 3, 16, 5, 16, 220, 10, 16,
	3, 16, 3, 16, 5, 16, 224, 10, 16, 3, 16, 3, 16, 3, 16, 5, 16, 229, 10,
	16, 3, 16, 3, 16, 5, 16, 233, 10, 16, 3, 16, 3, 16, 3, 16, 5, 16, 238,
	10, 16, 3, 16, 3, 16, 5, 16, 242, 10, 16, 3, 16, 3, 16, 3, 16, 5, 16, 247,
	10, 16, 3, 16, 3, 16, 5, 16, 251, 10, 16, 3, 16, 3, 16, 3, 16, 5, 16, 256,
	10, 16, 3, 16, 3, 16, 5, 16, 260, 10, 16, 3, 16, 3, 16, 3, 16, 5, 16, 265,
	10, 16, 3, 16, 3, 16, 5, 16, 269, 10, 16, 3, 16, 3, 16, 3, 16, 5, 16, 274,
	10, 16, 3, 16, 3, 16, 5, 16, 278, 10, 16, 3, 16, 3, 16, 3, 16, 5, 16, 283,
	10, 16, 3, 16, 3, 16, 5, 16, 287, 10, 16, 3, 16, 3, 16, 3, 16, 5, 16, 292,
	10, 16, 3, 16, 3, 16, 5, 16, 296, 10, 16, 3, 16, 3, 16, 3, 16, 5, 16, 301,
	10, 16, 3, 16, 3, 16, 5, 16, 305, 10, 16, 3, 16, 3, 16, 3, 16, 5, 16, 310,
	10, 16, 3, 16, 3, 16, 3, 16, 5, 16, 315, 10, 16, 3, 16, 3, 16, 5, 16, 319,
	10, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	7, 16, 330, 10, 16, 12, 16, 14, 16, 333, 11, 16, 3, 17, 3, 17, 5, 17, 337,
	10, 17, 3, 17, 3, 17, 5, 17, 341, 10, 17, 3, 17, 3, 17, 3, 18, 7, 18, 346,
	10, 18, 12, 18, 14, 18, 349, 11, 18, 3, 18, 5, 18, 352, 10, 18, 3, 18,
	5, 18, 355, 10, 18, 3, 18, 5, 18, 358, 10, 18, 3, 18, 6, 18, 361, 10, 18,
	13, 18, 14, 18, 362, 3, 18, 5, 18, 366, 10, 18, 3, 18, 7, 18, 369, 10,
	18, 12, 18, 14, 18, 372, 11, 18, 3, 18, 7, 18, 375, 10, 18, 12, 18, 14,
	18, 378, 11, 18, 3, 19, 5, 19, 381, 10, 19, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 20, 3, 20, 7, 20, 389, 10, 20, 12, 20, 14, 20, 392, 11, 20, 3, 20, 5,
	20, 395, 10, 20, 5, 20, 397, 10, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 412, 10,
	21, 3, 21, 5, 21, 415, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 5, 22, 424, 10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 430, 10,
	23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 2, 3, 30, 26, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
	2, 8, 4, 2, 11, 11, 19, 20, 3, 2, 14, 15, 3, 2, 22, 25, 3, 2, 26, 29, 4,
	2, 4, 4, 39, 39, 4, 2, 43, 43, 45, 47, 2, 515, 2, 52, 3, 2, 2, 2, 4, 60,
	3, 2, 2, 2, 6, 62, 3, 2, 2, 2, 8, 64, 3, 2, 2, 2, 10, 66, 3, 2, 2, 2, 12,
	82, 3, 2, 2, 2, 14, 84, 3, 2, 2, 2, 16, 94, 3, 2, 2, 2, 18, 104, 3, 2,
	2, 2, 20, 107, 3, 2, 2, 2, 22, 121, 3, 2, 2, 2, 24, 123, 3, 2, 2, 2, 26,
	140, 3, 2, 2, 2, 28, 146, 3, 2, 2, 2, 30, 188, 3, 2, 2, 2, 32, 334, 3,
	2, 2, 2, 34, 347, 3, 2, 2, 2, 36, 380, 3, 2, 2, 2, 38, 384, 3, 2, 2, 2,
	40, 414, 3, 2, 2, 2, 42, 423, 3, 2, 2, 2, 44, 429, 3, 2, 2, 2, 46, 431,
	3, 2, 2, 2, 48, 433, 3, 2, 2, 2, 50, 53, 5, 4, 3, 2, 51, 53, 5, 6, 4, 2,
	52, 50, 3, 2, 2, 2, 52, 51, 3, 2, 2, 2, 53, 57, 3, 2, 2, 2, 54, 56, 5,
	12, 7, 2, 55, 54, 3, 2, 2, 2, 56, 59, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57,
	58, 3, 2, 2, 2, 58, 3, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 60, 61, 7, 3, 2,
	2, 61, 5, 3, 2, 2, 2, 62, 63, 7, 4, 2, 2, 63, 7, 3, 2, 2, 2, 64, 65, 7,
	5, 2, 2, 65, 9, 3, 2, 2, 2, 66, 67, 7, 6, 2, 2, 67, 11, 3, 2, 2, 2, 68,
	83, 5, 14, 8, 2, 69, 70, 5, 8, 5, 2, 70, 71, 5, 14, 8, 2, 71, 83, 3, 2,
	2, 2, 72, 73, 5, 8, 5, 2, 73, 74, 5, 22, 12, 2, 74, 83, 3, 2, 2, 2, 75,
	76, 5, 10, 6, 2, 76, 77, 5, 14, 8, 2, 77, 83, 3, 2, 2, 2, 78, 79, 5, 10,
	6, 2, 79, 80, 5, 22, 12, 2, 80, 83, 3, 2, 2, 2, 81, 83, 5, 10, 6, 2, 82,
	68, 3, 2, 2, 2, 82, 69, 3, 2, 2, 2, 82, 72, 3, 2, 2, 2, 82, 75, 3, 2, 2,
	2, 82, 78, 3, 2, 2, 2, 82, 81, 3, 2, 2, 2, 83, 13, 3, 2, 2, 2, 84, 86,
	7, 7, 2, 2, 85, 87, 7, 49, 2, 2, 86, 85, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2,
	87, 88, 3, 2, 2, 2, 88, 90, 5, 16, 9, 2, 89, 91, 7, 49, 2, 2, 90, 89, 3,
	2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 93, 7, 8, 2, 2, 93,
	15, 3, 2, 2, 2, 94, 99, 5, 18, 10, 2, 95, 96, 7, 9, 2, 2, 96, 98, 5, 18,
	10, 2, 97, 95, 3, 2, 2, 2, 98, 101, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 99,
	100, 3, 2, 2, 2, 100, 17, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 102, 105, 5,
	22, 12, 2, 103, 105, 5, 20, 11, 2, 104, 102, 3, 2, 2, 2, 104, 103, 3, 2,
	2, 2, 105, 19, 3, 2, 2, 2, 106, 108, 5, 30, 16, 2, 107, 106, 3, 2, 2, 2,
	107, 108, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 111, 7, 10, 2, 2, 110,
	112, 5, 30, 16, 2, 111, 110, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 114,
	3, 2, 2, 2, 113, 115, 7, 10, 2, 2, 114, 113, 3, 2, 2, 2, 114, 115, 3, 2,
	2, 2, 115, 117, 3, 2, 2, 2, 116, 118, 5, 30, 16, 2, 117, 116, 3, 2, 2,
	2, 117, 118, 3, 2, 2, 2, 118, 21, 3, 2, 2, 2, 119, 122, 5, 30, 16, 2, 120,
	122, 7, 11, 2, 2, 121, 119, 3, 2, 2, 2, 121, 120, 3, 2, 2, 2, 122, 23,
	3, 2, 2, 2, 123, 135, 7, 12, 2, 2, 124, 129, 5, 26, 14, 2, 125, 126, 7,
	9, 2, 2, 126, 128, 5, 26, 14, 2, 127, 125, 3, 2, 2, 2, 128, 131, 3, 2,
	2, 2, 129, 127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 133, 3, 2, 2, 2,
	131, 129, 3, 2, 2, 2, 132, 134, 7, 9, 2, 2, 133, 132, 3, 2, 2, 2, 133,
	134, 3, 2, 2, 2, 134, 136, 3, 2, 2, 2, 135, 124, 3, 2, 2, 2, 135, 136,
	3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 138, 7, 13, 2, 2, 138, 25, 3, 2,
	2, 2, 139, 141, 7, 40, 2, 2, 140, 139, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2,
	141, 144, 3, 2, 2, 2, 142, 145, 5, 30, 16, 2, 143, 145, 5, 46, 24, 2, 144,
	142, 3, 2, 2, 2, 144, 143, 3, 2, 2, 2, 145, 27, 3, 2, 2, 2, 146, 151, 5,
	30, 16, 2, 147, 148, 7, 9, 2, 2, 148, 150, 5, 30, 16, 2, 149, 147, 3, 2,
	2, 2, 150, 153, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2,
	152, 29, 3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 154, 155, 8, 16, 1, 2, 155,
	157, 7, 14, 2, 2, 156, 158, 7, 49, 2, 2, 157, 156, 3, 2, 2, 2, 157, 158,
	3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 189, 5, 30, 16, 26, 160, 162, 7,
	15, 2, 2, 161, 163, 7, 49, 2, 2, 162, 161, 3, 2, 2, 2, 162, 163, 3, 2,
	2, 2, 163, 164, 3, 2, 2, 2, 164, 189, 5, 30, 16, 25, 165, 167, 7, 16, 2,
	2, 166, 168, 7, 49, 2, 2, 167, 166, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168,
	169, 3, 2, 2, 2, 169, 189, 5, 30, 16, 24, 170, 172, 7, 17, 2, 2, 171, 173,
	7, 49, 2, 2, 172, 171, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 174, 3, 2,
	2, 2, 174, 189, 5, 30, 16, 23, 175, 189, 5, 2, 2, 2, 176, 189, 5, 46, 24,
	2, 177, 189, 5, 44, 23, 2, 178, 189, 5, 32, 17, 2, 179, 189, 5, 38, 20,
	2, 180, 181, 7, 12, 2, 2, 181, 182, 5, 30, 16, 2, 182, 183, 7, 13, 2, 2,
	183, 189, 3, 2, 2, 2, 184, 185, 7, 36, 2, 2, 185, 186, 5, 30, 16, 2, 186,
	187, 7, 13, 2, 2, 187, 189, 3, 2, 2, 2, 188, 154, 3, 2, 2, 2, 188, 160,
	3, 2, 2, 2, 188, 165, 3, 2, 2, 2, 188, 170, 3, 2, 2, 2, 188, 175, 3, 2,
	2, 2, 188, 176, 3, 2, 2, 2, 188, 177, 3, 2, 2, 2, 188, 178, 3, 2, 2, 2,
	188, 179, 3, 2, 2, 2, 188, 180, 3, 2, 2, 2, 188, 184, 3, 2, 2, 2, 189,
	331, 3, 2, 2, 2, 190, 192, 12, 29, 2, 2, 191, 193, 7, 49, 2, 2, 192, 191,
	3, 2, 2, 2, 192, 193, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 196, 7, 5,
	2, 2, 195, 197, 7, 49, 2, 2, 196, 195, 3, 2, 2, 2, 196, 197, 3, 2, 2, 2,
	197, 198, 3, 2, 2, 2, 198, 330, 5, 30, 16, 30, 199, 201, 12, 22, 2, 2,
	200, 202, 7, 49, 2, 2, 201, 200, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202,
	203, 3, 2, 2, 2, 203, 205, 7, 18, 2, 2, 204, 206, 7, 49, 2, 2, 205, 204,
	3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 330, 5, 30,
	16, 22, 208, 210, 12, 21, 2, 2, 209, 211, 7, 49, 2, 2, 210, 209, 3, 2,
	2, 2, 210, 211, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 214, 9, 2, 2, 2,
	213, 215, 7, 49, 2, 2, 214, 213, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215,
	216, 3, 2, 2, 2, 216, 330, 5, 30, 16, 22, 217, 219, 12, 20, 2, 2, 218,
	220, 7, 49, 2, 2, 219, 218, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 221,
	3, 2, 2, 2, 221, 223, 9, 3, 2, 2, 222, 224, 7, 49, 2, 2, 223, 222, 3, 2,
	2, 2, 223, 224, 3, 2, 2, 2, 224, 225, 3, 2, 2, 2, 225, 330, 5, 30, 16,
	21, 226, 228, 12, 19, 2, 2, 227, 229, 7, 49, 2, 2, 228, 227, 3, 2, 2, 2,
	228, 229, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 232, 7, 21, 2, 2, 231,
	233, 7, 49, 2, 2, 232, 231, 3, 2, 2, 2, 232, 233, 3, 2, 2, 2, 233, 234,
	3, 2, 2, 2, 234, 330, 5, 30, 16, 20, 235, 237, 12, 18, 2, 2, 236, 238,
	7, 49, 2, 2, 237, 236, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238, 239, 3, 2,
	2, 2, 239, 241, 9, 4, 2, 2, 240, 242, 7, 49, 2, 2, 241, 240, 3, 2, 2, 2,
	241, 242, 3, 2, 2, 2, 242, 243, 3, 2, 2, 2, 243, 330, 5, 30, 16, 19, 244,
	246, 12, 17, 2, 2, 245, 247, 7, 49, 2, 2, 246, 245, 3, 2, 2, 2, 246, 247,
	3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 250, 7, 48, 2, 2, 249, 251, 7, 49,
	2, 2, 250, 249, 3, 2, 2, 2, 250, 251, 3, 2, 2, 2, 251, 252, 3, 2, 2, 2,
	252, 330, 5, 30, 16, 18, 253, 255, 12, 16, 2, 2, 254, 256, 7, 49, 2, 2,
	255, 254, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257,
	259, 9, 5, 2, 2, 258, 260, 7, 49, 2, 2, 259, 258, 3, 2, 2, 2, 259, 260,
	3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261, 330, 5, 30, 16, 17, 262, 264, 12,
	15, 2, 2, 263, 265, 7, 49, 2, 2, 264, 263, 3, 2, 2, 2, 264, 265, 3, 2,
	2, 2, 265, 266, 3, 2, 2, 2, 266, 268, 7, 30, 2, 2, 267, 269, 7, 49, 2,
	2, 268, 267, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 270, 3, 2, 2, 2, 270,
	330, 5, 30, 16, 16, 271, 273, 12, 14, 2, 2, 272, 274, 7, 49, 2, 2, 273,
	272, 3, 2, 2, 2, 273, 274, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 277,
	7, 31, 2, 2, 276, 278, 7, 49, 2, 2, 277, 276, 3, 2, 2, 2, 277, 278, 3,
	2, 2, 2, 278, 279, 3, 2, 2, 2, 279, 330, 5, 30, 16, 15, 280, 282, 12, 13,
	2, 2, 281, 283, 7, 49, 2, 2, 282, 281, 3, 2, 2, 2, 282, 283, 3, 2, 2, 2,
	283, 284, 3, 2, 2, 2, 284, 286, 7, 32, 2, 2, 285, 287, 7, 49, 2, 2, 286,
	285, 3, 2, 2, 2, 286, 287, 3, 2, 2, 2, 287, 288, 3, 2, 2, 2, 288, 330,
	5, 30, 16, 14, 289, 291, 12, 12, 2, 2, 290, 292, 7, 49, 2, 2, 291, 290,
	3, 2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 293, 3, 2, 2, 2, 293, 295, 7, 33,
	2, 2, 294, 296, 7, 49, 2, 2, 295, 294, 3, 2, 2, 2, 295, 296, 3, 2, 2, 2,
	296, 297, 3, 2, 2, 2, 297, 330, 5, 30, 16, 13, 298, 300, 12, 11, 2, 2,
	299, 301, 7, 49, 2, 2, 300, 299, 3, 2, 2, 2, 300, 301, 3, 2, 2, 2, 301,
	302, 3, 2, 2, 2, 302, 304, 7, 34, 2, 2, 303, 305, 7, 49, 2, 2, 304, 303,
	3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 305, 306, 3, 2, 2, 2, 306, 330, 5, 30,
	16, 12, 307, 309, 12, 10, 2, 2, 308, 310, 7, 49, 2, 2, 309, 308, 3, 2,
	2, 2, 309, 310, 3, 2, 2, 2, 310, 311, 3, 2, 2, 2, 311, 312, 7, 35, 2, 2,
	312, 314, 5, 30, 16, 2, 313, 315, 7, 49, 2, 2, 314, 313, 3, 2, 2, 2, 314,
	315, 3, 2, 2, 2, 315, 316, 3, 2, 2, 2, 316, 318, 7, 10, 2, 2, 317, 319,
	7, 49, 2, 2, 318, 317, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319, 320, 3, 2,
	2, 2, 320, 321, 5, 30, 16, 11, 321, 330, 3, 2, 2, 2, 322, 323, 12, 28,
	2, 2, 323, 324, 7, 7, 2, 2, 324, 325, 5, 28, 15, 2, 325, 326, 7, 8, 2,
	2, 326, 330, 3, 2, 2, 2, 327, 328, 12, 27, 2, 2, 328, 330, 5, 24, 13, 2,
	329, 190, 3, 2, 2, 2, 329, 199, 3, 2, 2, 2, 329, 208, 3, 2, 2, 2, 329,
	217, 3, 2, 2, 2, 329, 226, 3, 2, 2, 2, 329, 235, 3, 2, 2, 2, 329, 244,
	3, 2, 2, 2, 329, 253, 3, 2, 2, 2, 329, 262, 3, 2, 2, 2, 329, 271, 3, 2,
	2, 2, 329, 280, 3, 2, 2, 2, 329, 289, 3, 2, 2, 2, 329, 298, 3, 2, 2, 2,
	329, 307, 3, 2, 2, 2, 329, 322, 3, 2, 2, 2, 329, 327, 3, 2, 2, 2, 330,
	333, 3, 2, 2, 2, 331, 329, 3, 2, 2, 2, 331, 332, 3, 2, 2, 2, 332, 31, 3,
	2, 2, 2, 333, 331, 3, 2, 2, 2, 334, 336, 7, 7, 2, 2, 335, 337, 7, 49, 2,
	2, 336, 335, 3, 2, 2, 2, 336, 337, 3, 2, 2, 2, 337, 338, 3, 2, 2, 2, 338,
	340, 5, 34, 18, 2, 339, 341, 7, 49, 2, 2, 340, 339, 3, 2, 2, 2, 340, 341,
	3, 2, 2, 2, 341, 342, 3, 2, 2, 2, 342, 343, 7, 8, 2, 2, 343, 33, 3, 2,
	2, 2, 344, 346, 7, 9, 2, 2, 345, 344, 3, 2, 2, 2, 346, 349, 3, 2, 2, 2,
	347, 345, 3, 2, 2, 2, 347, 348, 3, 2, 2, 2, 348, 351, 3, 2, 2, 2, 349,
	347, 3, 2, 2, 2, 350, 352, 7, 49, 2, 2, 351, 350, 3, 2, 2, 2, 351, 352,
	3, 2, 2, 2, 352, 354, 3, 2, 2, 2, 353, 355, 5, 36, 19, 2, 354, 353, 3,
	2, 2, 2, 354, 355, 3, 2, 2, 2, 355, 370, 3, 2, 2, 2, 356, 358, 7, 49, 2,
	2, 357, 356, 3, 2, 2, 2, 357, 358, 3, 2, 2, 2, 358, 360, 3, 2, 2, 2, 359,
	361, 7, 9, 2, 2, 360, 359, 3, 2, 2, 2, 361, 362, 3, 2, 2, 2, 362, 360,
	3, 2, 2, 2, 362, 363, 3, 2, 2, 2, 363, 365, 3, 2, 2, 2, 364, 366, 7, 49,
	2, 2, 365, 364, 3, 2, 2, 2, 365, 366, 3, 2, 2, 2, 366, 367, 3, 2, 2, 2,
	367, 369, 5, 36, 19, 2, 368, 357, 3, 2, 2, 2, 369, 372, 3, 2, 2, 2, 370,
	368, 3, 2, 2, 2, 370, 371, 3, 2, 2, 2, 371, 376, 3, 2, 2, 2, 372, 370,
	3, 2, 2, 2, 373, 375, 7, 9, 2, 2, 374, 373, 3, 2, 2, 2, 375, 378, 3, 2,
	2, 2, 376, 374, 3, 2, 2, 2, 376, 377, 3, 2, 2, 2, 377, 35, 3, 2, 2, 2,
	378, 376, 3, 2, 2, 2, 379, 381, 7, 40, 2, 2, 380, 379, 3, 2, 2, 2, 380,
	381, 3, 2, 2, 2, 381, 382, 3, 2, 2, 2, 382, 383, 5, 30, 16, 2, 383, 37,
	3, 2, 2, 2, 384, 396, 7, 37, 2, 2, 385, 390, 5, 40, 21, 2, 386, 387, 7,
	9, 2, 2, 387, 389, 5, 40, 21, 2, 388, 386, 3, 2, 2, 2, 389, 392, 3, 2,
	2, 2, 390, 388, 3, 2, 2, 2, 390, 391, 3, 2, 2, 2, 391, 394, 3, 2, 2, 2,
	392, 390, 3, 2, 2, 2, 393, 395, 7, 9, 2, 2, 394, 393, 3, 2, 2, 2, 394,
	395, 3, 2, 2, 2, 395, 397, 3, 2, 2, 2, 396, 385, 3, 2, 2, 2, 396, 397,
	3, 2, 2, 2, 397, 398, 3, 2, 2, 2, 398, 399, 7, 38, 2, 2, 399, 39, 3, 2,
	2, 2, 400, 401, 5, 42, 22, 2, 401, 402, 7, 10, 2, 2, 402, 403, 5, 30, 16,
	2, 403, 415, 3, 2, 2, 2, 404, 405, 7, 7, 2, 2, 405, 406, 5, 30, 16, 2,
	406, 407, 7, 8, 2, 2, 407, 408, 7, 10, 2, 2, 408, 409, 5, 30, 16, 2, 409,
	415, 3, 2, 2, 2, 410, 412, 7, 40, 2, 2, 411, 410, 3, 2, 2, 2, 411, 412,
	3, 2, 2, 2, 412, 413, 3, 2, 2, 2, 413, 415, 5, 30, 16, 2, 414, 400, 3,
	2, 2, 2, 414, 404, 3, 2, 2, 2, 414, 411, 3, 2, 2, 2, 415, 41, 3, 2, 2,
	2, 416, 424, 5, 46, 24, 2, 417, 424, 7, 44, 2, 2, 418, 424, 5, 48, 25,
	2, 419, 420, 7, 7, 2, 2, 420, 421, 5, 30, 16, 2, 421, 422, 7, 8, 2, 2,
	422, 424, 3, 2, 2, 2, 423, 416, 3, 2, 2, 2, 423, 417, 3, 2, 2, 2, 423,
	418, 3, 2, 2, 2, 423, 419, 3, 2, 2, 2, 424, 43, 3, 2, 2, 2, 425, 430, 7,
	41, 2, 2, 426, 430, 7, 42, 2, 2, 427, 430, 7, 44, 2, 2, 428, 430, 5, 48,
	25, 2, 429, 425, 3, 2, 2, 2, 429, 426, 3, 2, 2, 2, 429, 427, 3, 2, 2, 2,
	429, 428, 3, 2, 2, 2, 430, 45, 3, 2, 2, 2, 431, 432, 9, 6, 2, 2, 432, 47,
	3, 2, 2, 2, 433, 434, 9, 7, 2, 2, 434, 49, 3, 2, 2, 2, 74, 52, 57, 82,
	86, 90, 99, 104, 107, 111, 114, 117, 121, 129, 133, 135, 140, 144, 151,
	157, 162, 167, 172, 188, 192, 196, 201, 205, 210, 214, 219, 223, 228, 232,
	237, 241, 246, 250, 255, 259, 264, 268, 273, 277, 282, 286, 291, 295, 300,
	304, 309, 314, 318, 329, 331, 336, 340, 347, 351, 354, 357, 362, 365, 370,
	376, 380, 390, 394, 396, 411, 414, 423, 429,
}
var literalNames = []string{
	"", "'$'", "'@'", "'.'", "'..'", "'['", "']'", "','", "':'", "'*'", "'('",
	"')'", "'+'", "'-'", "'~'", "'!'", "'**'", "'/'", "'%'", "'??'", "'<'",
	"'>'", "'<='", "'>='", "'=='", "'!='", "'==='", "'!=='", "'&'", "'^'",
	"'|'", "'&&'", "'||'", "'?'", "'?('", "'{'", "'}'", "", "'...'", "'null'",
	"", "", "", "", "", "", "'in'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "Identifier", "Ellipsis", "NullLiteral", "BooleanLiteral", "DecimalLiteral",
	"StringLiteral", "HexIntegerLiteral", "OctalIntegerLiteral2", "BinaryIntegerLiteral",
	"In", "SP",
}

var ruleNames = []string{
	"path", "root_selector", "current_element_selector", "child_selector",
	"recursive_descent", "path_element", "bracketed_selector", "union", "unionPart",
	"slice", "selector", "arguments", "argument", "expressionSequence", "singleExpression",
	"arrayLiteral", "elementList", "arrayElement", "objectLiteral", "propertyAssignment",
	"propertyName", "literal", "identifier", "numericLiteral",
}

type JsonpathParser struct {
	*antlr.BaseParser
}

// NewJsonpathParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *JsonpathParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewJsonpathParser(input antlr.TokenStream) *JsonpathParser {
	this := new(JsonpathParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
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
	JsonpathParserIdentifier           = 37
	JsonpathParserEllipsis             = 38
	JsonpathParserNullLiteral          = 39
	JsonpathParserBooleanLiteral       = 40
	JsonpathParserDecimalLiteral       = 41
	JsonpathParserStringLiteral        = 42
	JsonpathParserHexIntegerLiteral    = 43
	JsonpathParserOctalIntegerLiteral2 = 44
	JsonpathParserBinaryIntegerLiteral = 45
	JsonpathParserIn                   = 46
	JsonpathParserSP                   = 47
)

// JsonpathParser rules.
const (
	JsonpathParserRULE_path                     = 0
	JsonpathParserRULE_root_selector            = 1
	JsonpathParserRULE_current_element_selector = 2
	JsonpathParserRULE_child_selector           = 3
	JsonpathParserRULE_recursive_descent        = 4
	JsonpathParserRULE_path_element             = 5
	JsonpathParserRULE_bracketed_selector       = 6
	JsonpathParserRULE_union                    = 7
	JsonpathParserRULE_unionPart                = 8
	JsonpathParserRULE_slice                    = 9
	JsonpathParserRULE_selector                 = 10
	JsonpathParserRULE_arguments                = 11
	JsonpathParserRULE_argument                 = 12
	JsonpathParserRULE_expressionSequence       = 13
	JsonpathParserRULE_singleExpression         = 14
	JsonpathParserRULE_arrayLiteral             = 15
	JsonpathParserRULE_elementList              = 16
	JsonpathParserRULE_arrayElement             = 17
	JsonpathParserRULE_objectLiteral            = 18
	JsonpathParserRULE_propertyAssignment       = 19
	JsonpathParserRULE_propertyName             = 20
	JsonpathParserRULE_literal                  = 21
	JsonpathParserRULE_identifier               = 22
	JsonpathParserRULE_numericLiteral           = 23
)

// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_path
	return p
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) Root_selector() IRoot_selectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRoot_selectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRoot_selectorContext)
}

func (s *PathContext) Current_element_selector() ICurrent_element_selectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICurrent_element_selectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICurrent_element_selectorContext)
}

func (s *PathContext) AllPath_element() []IPath_elementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPath_elementContext)(nil)).Elem())
	var tst = make([]IPath_elementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPath_elementContext)
		}
	}

	return tst
}

func (s *PathContext) Path_element(i int) IPath_elementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPath_elementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPath_elementContext)
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
	p.SetState(50)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonpathParserT__0:
		{
			p.SetState(48)
			p.Root_selector()
		}

	case JsonpathParserT__1:
		{
			p.SetState(49)
			p.Current_element_selector()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(52)
				p.Path_element()
			}

		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IRoot_selectorContext is an interface to support dynamic dispatch.
type IRoot_selectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRoot_selectorContext differentiates from other interfaces.
	IsRoot_selectorContext()
}

type Root_selectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRoot_selectorContext() *Root_selectorContext {
	var p = new(Root_selectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_root_selector
	return p
}

func (*Root_selectorContext) IsRoot_selectorContext() {}

func NewRoot_selectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Root_selectorContext {
	var p = new(Root_selectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_root_selector

	return p
}

func (s *Root_selectorContext) GetParser() antlr.Parser { return s.parser }
func (s *Root_selectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Root_selectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Root_selectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterRoot_selector(s)
	}
}

func (s *Root_selectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitRoot_selector(s)
	}
}

func (p *JsonpathParser) Root_selector() (localctx IRoot_selectorContext) {
	localctx = NewRoot_selectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JsonpathParserRULE_root_selector)

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
		p.SetState(58)
		p.Match(JsonpathParserT__0)
	}

	return localctx
}

// ICurrent_element_selectorContext is an interface to support dynamic dispatch.
type ICurrent_element_selectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCurrent_element_selectorContext differentiates from other interfaces.
	IsCurrent_element_selectorContext()
}

type Current_element_selectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCurrent_element_selectorContext() *Current_element_selectorContext {
	var p = new(Current_element_selectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_current_element_selector
	return p
}

func (*Current_element_selectorContext) IsCurrent_element_selectorContext() {}

func NewCurrent_element_selectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Current_element_selectorContext {
	var p = new(Current_element_selectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_current_element_selector

	return p
}

func (s *Current_element_selectorContext) GetParser() antlr.Parser { return s.parser }
func (s *Current_element_selectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Current_element_selectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Current_element_selectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterCurrent_element_selector(s)
	}
}

func (s *Current_element_selectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitCurrent_element_selector(s)
	}
}

func (p *JsonpathParser) Current_element_selector() (localctx ICurrent_element_selectorContext) {
	localctx = NewCurrent_element_selectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JsonpathParserRULE_current_element_selector)

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
		p.SetState(60)
		p.Match(JsonpathParserT__1)
	}

	return localctx
}

// IChild_selectorContext is an interface to support dynamic dispatch.
type IChild_selectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChild_selectorContext differentiates from other interfaces.
	IsChild_selectorContext()
}

type Child_selectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChild_selectorContext() *Child_selectorContext {
	var p = new(Child_selectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_child_selector
	return p
}

func (*Child_selectorContext) IsChild_selectorContext() {}

func NewChild_selectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Child_selectorContext {
	var p = new(Child_selectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_child_selector

	return p
}

func (s *Child_selectorContext) GetParser() antlr.Parser { return s.parser }
func (s *Child_selectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Child_selectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Child_selectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterChild_selector(s)
	}
}

func (s *Child_selectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitChild_selector(s)
	}
}

func (p *JsonpathParser) Child_selector() (localctx IChild_selectorContext) {
	localctx = NewChild_selectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JsonpathParserRULE_child_selector)

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
		p.SetState(62)
		p.Match(JsonpathParserT__2)
	}

	return localctx
}

// IRecursive_descentContext is an interface to support dynamic dispatch.
type IRecursive_descentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRecursive_descentContext differentiates from other interfaces.
	IsRecursive_descentContext()
}

type Recursive_descentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecursive_descentContext() *Recursive_descentContext {
	var p = new(Recursive_descentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_recursive_descent
	return p
}

func (*Recursive_descentContext) IsRecursive_descentContext() {}

func NewRecursive_descentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Recursive_descentContext {
	var p = new(Recursive_descentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_recursive_descent

	return p
}

func (s *Recursive_descentContext) GetParser() antlr.Parser { return s.parser }
func (s *Recursive_descentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Recursive_descentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Recursive_descentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterRecursive_descent(s)
	}
}

func (s *Recursive_descentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitRecursive_descent(s)
	}
}

func (p *JsonpathParser) Recursive_descent() (localctx IRecursive_descentContext) {
	localctx = NewRecursive_descentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JsonpathParserRULE_recursive_descent)

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
		p.SetState(64)
		p.Match(JsonpathParserT__3)
	}

	return localctx
}

// IPath_elementContext is an interface to support dynamic dispatch.
type IPath_elementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPath_elementContext differentiates from other interfaces.
	IsPath_elementContext()
}

type Path_elementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPath_elementContext() *Path_elementContext {
	var p = new(Path_elementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_path_element
	return p
}

func (*Path_elementContext) IsPath_elementContext() {}

func NewPath_elementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Path_elementContext {
	var p = new(Path_elementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_path_element

	return p
}

func (s *Path_elementContext) GetParser() antlr.Parser { return s.parser }

func (s *Path_elementContext) Bracketed_selector() IBracketed_selectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBracketed_selectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBracketed_selectorContext)
}

func (s *Path_elementContext) Child_selector() IChild_selectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChild_selectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChild_selectorContext)
}

func (s *Path_elementContext) Selector() ISelectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *Path_elementContext) Recursive_descent() IRecursive_descentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecursive_descentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRecursive_descentContext)
}

func (s *Path_elementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Path_elementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Path_elementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterPath_element(s)
	}
}

func (s *Path_elementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitPath_element(s)
	}
}

func (p *JsonpathParser) Path_element() (localctx IPath_elementContext) {
	localctx = NewPath_elementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JsonpathParserRULE_path_element)

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

	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(66)
			p.Bracketed_selector()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)
			p.Child_selector()
		}
		{
			p.SetState(68)
			p.Bracketed_selector()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(70)
			p.Child_selector()
		}
		{
			p.SetState(71)
			p.Selector()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(73)
			p.Recursive_descent()
		}
		{
			p.SetState(74)
			p.Bracketed_selector()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(76)
			p.Recursive_descent()
		}
		{
			p.SetState(77)
			p.Selector()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(79)
			p.Recursive_descent()
		}

	}

	return localctx
}

// IBracketed_selectorContext is an interface to support dynamic dispatch.
type IBracketed_selectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBracketed_selectorContext differentiates from other interfaces.
	IsBracketed_selectorContext()
}

type Bracketed_selectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBracketed_selectorContext() *Bracketed_selectorContext {
	var p = new(Bracketed_selectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_bracketed_selector
	return p
}

func (*Bracketed_selectorContext) IsBracketed_selectorContext() {}

func NewBracketed_selectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bracketed_selectorContext {
	var p = new(Bracketed_selectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_bracketed_selector

	return p
}

func (s *Bracketed_selectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Bracketed_selectorContext) Union() IUnionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionContext)
}

func (s *Bracketed_selectorContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserSP)
}

func (s *Bracketed_selectorContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, i)
}

func (s *Bracketed_selectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bracketed_selectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bracketed_selectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterBracketed_selector(s)
	}
}

func (s *Bracketed_selectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitBracketed_selector(s)
	}
}

func (p *JsonpathParser) Bracketed_selector() (localctx IBracketed_selectorContext) {
	localctx = NewBracketed_selectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JsonpathParserRULE_bracketed_selector)
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
		p.SetState(82)
		p.Match(JsonpathParserT__4)
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JsonpathParserSP {
		{
			p.SetState(83)
			p.Match(JsonpathParserSP)
		}

	}
	{
		p.SetState(86)
		p.Union()
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JsonpathParserSP {
		{
			p.SetState(87)
			p.Match(JsonpathParserSP)
		}

	}
	{
		p.SetState(90)
		p.Match(JsonpathParserT__5)
	}

	return localctx
}

// IUnionContext is an interface to support dynamic dispatch.
type IUnionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionContext differentiates from other interfaces.
	IsUnionContext()
}

type UnionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionContext() *UnionContext {
	var p = new(UnionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_union
	return p
}

func (*UnionContext) IsUnionContext() {}

func NewUnionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionContext {
	var p = new(UnionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_union

	return p
}

func (s *UnionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionContext) AllUnionPart() []IUnionPartContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnionPartContext)(nil)).Elem())
	var tst = make([]IUnionPartContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnionPartContext)
		}
	}

	return tst
}

func (s *UnionContext) UnionPart(i int) IUnionPartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionPartContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnionPartContext)
}

func (s *UnionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterUnion(s)
	}
}

func (s *UnionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitUnion(s)
	}
}

func (p *JsonpathParser) Union() (localctx IUnionContext) {
	localctx = NewUnionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JsonpathParserRULE_union)
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
		p.SetState(92)
		p.UnionPart()
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JsonpathParserT__6 {
		{
			p.SetState(93)
			p.Match(JsonpathParserT__6)
		}
		{
			p.SetState(94)
			p.UnionPart()
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUnionPartContext is an interface to support dynamic dispatch.
type IUnionPartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionPartContext differentiates from other interfaces.
	IsUnionPartContext()
}

type UnionPartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionPartContext() *UnionPartContext {
	var p = new(UnionPartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_unionPart
	return p
}

func (*UnionPartContext) IsUnionPartContext() {}

func NewUnionPartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionPartContext {
	var p = new(UnionPartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_unionPart

	return p
}

func (s *UnionPartContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionPartContext) Selector() ISelectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *UnionPartContext) Slice() ISliceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISliceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISliceContext)
}

func (s *UnionPartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionPartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionPartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterUnionPart(s)
	}
}

func (s *UnionPartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitUnionPart(s)
	}
}

func (p *JsonpathParser) UnionPart() (localctx IUnionPartContext) {
	localctx = NewUnionPartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JsonpathParserRULE_unionPart)

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

	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.Selector()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)
			p.Slice()
		}

	}

	return localctx
}

// ISliceContext is an interface to support dynamic dispatch.
type ISliceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSliceContext differentiates from other interfaces.
	IsSliceContext()
}

type SliceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceContext() *SliceContext {
	var p = new(SliceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_slice
	return p
}

func (*SliceContext) IsSliceContext() {}

func NewSliceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceContext {
	var p = new(SliceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_slice

	return p
}

func (s *SliceContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *SliceContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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
	p.EnterRule(localctx, 18, JsonpathParserRULE_slice)
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
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonpathParserT__0)|(1<<JsonpathParserT__1)|(1<<JsonpathParserT__4)|(1<<JsonpathParserT__9)|(1<<JsonpathParserT__11)|(1<<JsonpathParserT__12)|(1<<JsonpathParserT__13)|(1<<JsonpathParserT__14))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(JsonpathParserT__33-34))|(1<<(JsonpathParserT__34-34))|(1<<(JsonpathParserIdentifier-34))|(1<<(JsonpathParserNullLiteral-34))|(1<<(JsonpathParserBooleanLiteral-34))|(1<<(JsonpathParserDecimalLiteral-34))|(1<<(JsonpathParserStringLiteral-34))|(1<<(JsonpathParserHexIntegerLiteral-34))|(1<<(JsonpathParserOctalIntegerLiteral2-34))|(1<<(JsonpathParserBinaryIntegerLiteral-34)))) != 0) {
		{
			p.SetState(104)
			p.singleExpression(0)
		}

	}
	{
		p.SetState(107)
		p.Match(JsonpathParserT__7)
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(108)
			p.singleExpression(0)
		}

	}

	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JsonpathParserT__7 {
		{
			p.SetState(111)
			p.Match(JsonpathParserT__7)
		}

	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonpathParserT__0)|(1<<JsonpathParserT__1)|(1<<JsonpathParserT__4)|(1<<JsonpathParserT__9)|(1<<JsonpathParserT__11)|(1<<JsonpathParserT__12)|(1<<JsonpathParserT__13)|(1<<JsonpathParserT__14))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(JsonpathParserT__33-34))|(1<<(JsonpathParserT__34-34))|(1<<(JsonpathParserIdentifier-34))|(1<<(JsonpathParserNullLiteral-34))|(1<<(JsonpathParserBooleanLiteral-34))|(1<<(JsonpathParserDecimalLiteral-34))|(1<<(JsonpathParserStringLiteral-34))|(1<<(JsonpathParserHexIntegerLiteral-34))|(1<<(JsonpathParserOctalIntegerLiteral2-34))|(1<<(JsonpathParserBinaryIntegerLiteral-34)))) != 0) {
		{
			p.SetState(114)
			p.singleExpression(0)
		}

	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorContext() *SelectorContext {
	var p = new(SelectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_selector
	return p
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_selector

	return p
}

func (s *SelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

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
	p.EnterRule(localctx, 20, JsonpathParserRULE_selector)

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

	p.SetState(119)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonpathParserT__0, JsonpathParserT__1, JsonpathParserT__4, JsonpathParserT__9, JsonpathParserT__11, JsonpathParserT__12, JsonpathParserT__13, JsonpathParserT__14, JsonpathParserT__33, JsonpathParserT__34, JsonpathParserIdentifier, JsonpathParserNullLiteral, JsonpathParserBooleanLiteral, JsonpathParserDecimalLiteral, JsonpathParserStringLiteral, JsonpathParserHexIntegerLiteral, JsonpathParserOctalIntegerLiteral2, JsonpathParserBinaryIntegerLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(117)
			p.singleExpression(0)
		}

	case JsonpathParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.Match(JsonpathParserT__8)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *ArgumentsContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
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
	p.EnterRule(localctx, 22, JsonpathParserRULE_arguments)
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
		p.SetState(121)
		p.Match(JsonpathParserT__9)
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonpathParserT__0)|(1<<JsonpathParserT__1)|(1<<JsonpathParserT__4)|(1<<JsonpathParserT__9)|(1<<JsonpathParserT__11)|(1<<JsonpathParserT__12)|(1<<JsonpathParserT__13)|(1<<JsonpathParserT__14))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(JsonpathParserT__33-34))|(1<<(JsonpathParserT__34-34))|(1<<(JsonpathParserIdentifier-34))|(1<<(JsonpathParserEllipsis-34))|(1<<(JsonpathParserNullLiteral-34))|(1<<(JsonpathParserBooleanLiteral-34))|(1<<(JsonpathParserDecimalLiteral-34))|(1<<(JsonpathParserStringLiteral-34))|(1<<(JsonpathParserHexIntegerLiteral-34))|(1<<(JsonpathParserOctalIntegerLiteral2-34))|(1<<(JsonpathParserBinaryIntegerLiteral-34)))) != 0) {
		{
			p.SetState(122)
			p.Argument()
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(123)
					p.Match(JsonpathParserT__6)
				}
				{
					p.SetState(124)
					p.Argument()
				}

			}
			p.SetState(129)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonpathParserT__6 {
			{
				p.SetState(130)
				p.Match(JsonpathParserT__6)
			}

		}

	}
	{
		p.SetState(135)
		p.Match(JsonpathParserT__10)
	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ArgumentContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ArgumentContext) Ellipsis() antlr.TerminalNode {
	return s.GetToken(JsonpathParserEllipsis, 0)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *JsonpathParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, JsonpathParserRULE_argument)
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
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JsonpathParserEllipsis {
		{
			p.SetState(137)
			p.Match(JsonpathParserEllipsis)
		}

	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(140)
			p.singleExpression(0)
		}

	case 2:
		{
			p.SetState(141)
			p.Identifier()
		}

	}

	return localctx
}

// IExpressionSequenceContext is an interface to support dynamic dispatch.
type IExpressionSequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionSequenceContext differentiates from other interfaces.
	IsExpressionSequenceContext()
}

type ExpressionSequenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionSequenceContext() *ExpressionSequenceContext {
	var p = new(ExpressionSequenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_expressionSequence
	return p
}

func (*ExpressionSequenceContext) IsExpressionSequenceContext() {}

func NewExpressionSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionSequenceContext {
	var p = new(ExpressionSequenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_expressionSequence

	return p
}

func (s *ExpressionSequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionSequenceContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionSequenceContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ExpressionSequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionSequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterExpressionSequence(s)
	}
}

func (s *ExpressionSequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitExpressionSequence(s)
	}
}

func (p *JsonpathParser) ExpressionSequence() (localctx IExpressionSequenceContext) {
	localctx = NewExpressionSequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, JsonpathParserRULE_expressionSequence)
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
		p.SetState(144)
		p.singleExpression(0)
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JsonpathParserT__6 {
		{
			p.SetState(145)
			p.Match(JsonpathParserT__6)
		}
		{
			p.SetState(146)
			p.singleExpression(0)
		}

		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISingleExpressionContext is an interface to support dynamic dispatch.
type ISingleExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleExpressionContext differentiates from other interfaces.
	IsSingleExpressionContext()
}

type SingleExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleExpressionContext() *SingleExpressionContext {
	var p = new(SingleExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_singleExpression
	return p
}

func (*SingleExpressionContext) IsSingleExpressionContext() {}

func NewSingleExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleExpressionContext {
	var p = new(SingleExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_singleExpression

	return p
}

func (s *SingleExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleExpressionContext) CopyFrom(ctx *SingleExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SingleExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TernaryExpressionContext struct {
	*SingleExpressionContext
}

func NewTernaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TernaryExpressionContext {
	var p = new(TernaryExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *TernaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *TernaryExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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
	*SingleExpressionContext
}

func NewChainExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ChainExpressionContext {
	var p = new(ChainExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ChainExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChainExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *ChainExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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
	*SingleExpressionContext
}

func NewLogicalAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalAndExpressionContext {
	var p = new(LogicalAndExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *LogicalAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *LogicalAndExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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
	*SingleExpressionContext
}

func NewPowerExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowerExpressionContext {
	var p = new(PowerExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *PowerExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *PowerExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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
	*SingleExpressionContext
}

func NewObjectLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectLiteralExpressionContext {
	var p = new(ObjectLiteralExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ObjectLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectLiteralExpressionContext) ObjectLiteral() IObjectLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectLiteralContext)(nil)).Elem(), 0)

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
	*SingleExpressionContext
}

func NewInExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InExpressionContext {
	var p = new(InExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *InExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *InExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *InExpressionContext) In() antlr.TerminalNode {
	return s.GetToken(JsonpathParserIn, 0)
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
	*SingleExpressionContext
}

func NewLogicalOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalOrExpressionContext {
	var p = new(LogicalOrExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *LogicalOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *LogicalOrExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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
	*SingleExpressionContext
}

func NewNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExpressionContext {
	var p = new(NotExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *NotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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

type ArgumentsExpressionContext struct {
	*SingleExpressionContext
}

func NewArgumentsExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArgumentsExpressionContext {
	var p = new(ArgumentsExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ArgumentsExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ArgumentsExpressionContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

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
	*SingleExpressionContext
}

func NewUnaryMinusExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryMinusExpressionContext {
	var p = new(UnaryMinusExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *UnaryMinusExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMinusExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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

type PathExpressionContext struct {
	*SingleExpressionContext
}

func NewPathExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PathExpressionContext {
	var p = new(PathExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *PathExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathExpressionContext) Path() IPathContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *PathExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterPathExpression(s)
	}
}

func (s *PathExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitPathExpression(s)
	}
}

type UnaryPlusExpressionContext struct {
	*SingleExpressionContext
}

func NewUnaryPlusExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryPlusExpressionContext {
	var p = new(UnaryPlusExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *UnaryPlusExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryPlusExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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
	*SingleExpressionContext
}

func NewFilterExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FilterExpressionContext {
	var p = new(FilterExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *FilterExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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
	*SingleExpressionContext
}

func NewEqualityExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *EqualityExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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
	*SingleExpressionContext
}

func NewBitXOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitXOrExpressionContext {
	var p = new(BitXOrExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *BitXOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitXOrExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *BitXOrExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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
	*SingleExpressionContext
}

func NewMultiplicativeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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
	*SingleExpressionContext
}

func NewParenthesizedExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesizedExpressionContext {
	var p = new(ParenthesizedExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ParenthesizedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesizedExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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
	*SingleExpressionContext
}

func NewAdditiveExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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
	*SingleExpressionContext
}

func NewRelationalExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *RelationalExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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

type BitNotExpressionContext struct {
	*SingleExpressionContext
}

func NewBitNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitNotExpressionContext {
	var p = new(BitNotExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *BitNotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitNotExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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
	*SingleExpressionContext
}

func NewLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *LiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpressionContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

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
	*SingleExpressionContext
}

func NewArrayLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayLiteralExpressionContext {
	var p = new(ArrayLiteralExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ArrayLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralExpressionContext) ArrayLiteral() IArrayLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayLiteralContext)(nil)).Elem(), 0)

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
	*SingleExpressionContext
}

func NewMemberIndexExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberIndexExpressionContext {
	var p = new(MemberIndexExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *MemberIndexExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberIndexExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *MemberIndexExpressionContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
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
	*SingleExpressionContext
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

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
	*SingleExpressionContext
}

func NewBitAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitAndExpressionContext {
	var p = new(BitAndExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *BitAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitAndExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *BitAndExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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
	*SingleExpressionContext
}

func NewBitOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitOrExpressionContext {
	var p = new(BitOrExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *BitOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitOrExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *BitOrExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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

type CoalesceExpressionContext struct {
	*SingleExpressionContext
}

func NewCoalesceExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CoalesceExpressionContext {
	var p = new(CoalesceExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *CoalesceExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoalesceExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *CoalesceExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *CoalesceExpressionContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonpathParserSP)
}

func (s *CoalesceExpressionContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonpathParserSP, i)
}

func (s *CoalesceExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterCoalesceExpression(s)
	}
}

func (s *CoalesceExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitCoalesceExpression(s)
	}
}

func (p *JsonpathParser) SingleExpression() (localctx ISingleExpressionContext) {
	return p.singleExpression(0)
}

func (p *JsonpathParser) singleExpression(_p int) (localctx ISingleExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewSingleExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISingleExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 28
	p.EnterRecursionRule(localctx, 28, JsonpathParserRULE_singleExpression, _p)
	var _la int

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
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		localctx = NewUnaryPlusExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(153)
			p.Match(JsonpathParserT__11)
		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonpathParserSP {
			{
				p.SetState(154)
				p.Match(JsonpathParserSP)
			}

		}
		{
			p.SetState(157)
			p.singleExpression(24)
		}

	case 2:
		localctx = NewUnaryMinusExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(158)
			p.Match(JsonpathParserT__12)
		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonpathParserSP {
			{
				p.SetState(159)
				p.Match(JsonpathParserSP)
			}

		}
		{
			p.SetState(162)
			p.singleExpression(23)
		}

	case 3:
		localctx = NewBitNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(163)
			p.Match(JsonpathParserT__13)
		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonpathParserSP {
			{
				p.SetState(164)
				p.Match(JsonpathParserSP)
			}

		}
		{
			p.SetState(167)
			p.singleExpression(22)
		}

	case 4:
		localctx = NewNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(168)
			p.Match(JsonpathParserT__14)
		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonpathParserSP {
			{
				p.SetState(169)
				p.Match(JsonpathParserSP)
			}

		}
		{
			p.SetState(172)
			p.singleExpression(21)
		}

	case 5:
		localctx = NewPathExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(173)
			p.Path()
		}

	case 6:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(174)
			p.Identifier()
		}

	case 7:
		localctx = NewLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(175)
			p.Literal()
		}

	case 8:
		localctx = NewArrayLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(176)
			p.ArrayLiteral()
		}

	case 9:
		localctx = NewObjectLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(177)
			p.ObjectLiteral()
		}

	case 10:
		localctx = NewParenthesizedExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(178)
			p.Match(JsonpathParserT__9)
		}
		{
			p.SetState(179)
			p.singleExpression(0)
		}
		{
			p.SetState(180)
			p.Match(JsonpathParserT__10)
		}

	case 11:
		localctx = NewFilterExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(182)
			p.Match(JsonpathParserT__33)
		}
		{
			p.SetState(183)
			p.singleExpression(0)
		}
		{
			p.SetState(184)
			p.Match(JsonpathParserT__10)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(327)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext()) {
			case 1:
				localctx = NewChainExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_singleExpression)
				p.SetState(188)

				if !(p.Precpred(p.GetParserRuleContext(), 27)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 27)", ""))
				}
				p.SetState(190)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(189)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(192)
					p.Match(JsonpathParserT__2)
				}
				p.SetState(194)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(193)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(196)
					p.singleExpression(28)
				}

			case 2:
				localctx = NewPowerExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_singleExpression)
				p.SetState(197)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
				}
				p.SetState(199)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(198)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(201)
					p.Match(JsonpathParserT__15)
				}
				p.SetState(203)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(202)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(205)
					p.singleExpression(20)
				}

			case 3:
				localctx = NewMultiplicativeExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_singleExpression)
				p.SetState(206)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				p.SetState(208)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(207)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(210)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonpathParserT__8)|(1<<JsonpathParserT__16)|(1<<JsonpathParserT__17))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(212)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(211)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(214)
					p.singleExpression(20)
				}

			case 4:
				localctx = NewAdditiveExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_singleExpression)
				p.SetState(215)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				p.SetState(217)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(216)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(219)
					_la = p.GetTokenStream().LA(1)

					if !(_la == JsonpathParserT__11 || _la == JsonpathParserT__12) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(221)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(220)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(223)
					p.singleExpression(19)
				}

			case 5:
				localctx = NewCoalesceExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_singleExpression)
				p.SetState(224)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				p.SetState(226)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(225)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(228)
					p.Match(JsonpathParserT__18)
				}
				p.SetState(230)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(229)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(232)
					p.singleExpression(18)
				}

			case 6:
				localctx = NewRelationalExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_singleExpression)
				p.SetState(233)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				p.SetState(235)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(234)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(237)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonpathParserT__19)|(1<<JsonpathParserT__20)|(1<<JsonpathParserT__21)|(1<<JsonpathParserT__22))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(239)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(238)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(241)
					p.singleExpression(17)
				}

			case 7:
				localctx = NewInExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_singleExpression)
				p.SetState(242)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				p.SetState(244)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(243)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(246)
					p.Match(JsonpathParserIn)
				}
				p.SetState(248)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(247)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(250)
					p.singleExpression(16)
				}

			case 8:
				localctx = NewEqualityExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_singleExpression)
				p.SetState(251)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				p.SetState(253)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(252)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(255)
					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonpathParserT__23)|(1<<JsonpathParserT__24)|(1<<JsonpathParserT__25)|(1<<JsonpathParserT__26))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(257)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(256)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(259)
					p.singleExpression(15)
				}

			case 9:
				localctx = NewBitAndExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_singleExpression)
				p.SetState(260)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				p.SetState(262)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(261)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(264)
					p.Match(JsonpathParserT__27)
				}
				p.SetState(266)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(265)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(268)
					p.singleExpression(14)
				}

			case 10:
				localctx = NewBitXOrExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_singleExpression)
				p.SetState(269)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				p.SetState(271)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(270)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(273)
					p.Match(JsonpathParserT__28)
				}
				p.SetState(275)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(274)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(277)
					p.singleExpression(13)
				}

			case 11:
				localctx = NewBitOrExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_singleExpression)
				p.SetState(278)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				p.SetState(280)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(279)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(282)
					p.Match(JsonpathParserT__29)
				}
				p.SetState(284)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(283)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(286)
					p.singleExpression(12)
				}

			case 12:
				localctx = NewLogicalAndExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_singleExpression)
				p.SetState(287)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				p.SetState(289)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(288)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(291)
					p.Match(JsonpathParserT__30)
				}
				p.SetState(293)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(292)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(295)
					p.singleExpression(11)
				}

			case 13:
				localctx = NewLogicalOrExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_singleExpression)
				p.SetState(296)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				p.SetState(298)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(297)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(300)
					p.Match(JsonpathParserT__31)
				}
				p.SetState(302)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(301)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(304)
					p.singleExpression(10)
				}

			case 14:
				localctx = NewTernaryExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_singleExpression)
				p.SetState(305)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				p.SetState(307)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(306)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(309)
					p.Match(JsonpathParserT__32)
				}
				{
					p.SetState(310)
					p.singleExpression(0)
				}
				p.SetState(312)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(311)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(314)
					p.Match(JsonpathParserT__7)
				}
				p.SetState(316)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == JsonpathParserSP {
					{
						p.SetState(315)
						p.Match(JsonpathParserSP)
					}

				}
				{
					p.SetState(318)
					p.singleExpression(9)
				}

			case 15:
				localctx = NewMemberIndexExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_singleExpression)
				p.SetState(320)

				if !(p.Precpred(p.GetParserRuleContext(), 26)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 26)", ""))
				}
				{
					p.SetState(321)
					p.Match(JsonpathParserT__4)
				}
				{
					p.SetState(322)
					p.ExpressionSequence()
				}
				{
					p.SetState(323)
					p.Match(JsonpathParserT__5)
				}

			case 16:
				localctx = NewArgumentsExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonpathParserRULE_singleExpression)
				p.SetState(325)

				if !(p.Precpred(p.GetParserRuleContext(), 25)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 25)", ""))
				}
				{
					p.SetState(326)
					p.Arguments()
				}

			}

		}
		p.SetState(331)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext())
	}

	return localctx
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_arrayLiteral
	return p
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) ElementList() IElementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementListContext)(nil)).Elem(), 0)

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
	p.EnterRule(localctx, 30, JsonpathParserRULE_arrayLiteral)
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
		p.SetState(332)
		p.Match(JsonpathParserT__4)
	}
	p.SetState(334)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 54, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(333)
			p.Match(JsonpathParserSP)
		}

	}
	{
		p.SetState(336)
		p.ElementList()
	}
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JsonpathParserSP {
		{
			p.SetState(337)
			p.Match(JsonpathParserSP)
		}

	}
	{
		p.SetState(340)
		p.Match(JsonpathParserT__5)
	}

	return localctx
}

// IElementListContext is an interface to support dynamic dispatch.
type IElementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementListContext differentiates from other interfaces.
	IsElementListContext()
}

type ElementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementListContext() *ElementListContext {
	var p = new(ElementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_elementList
	return p
}

func (*ElementListContext) IsElementListContext() {}

func NewElementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementListContext {
	var p = new(ElementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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

func (s *ElementListContext) AllArrayElement() []IArrayElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArrayElementContext)(nil)).Elem())
	var tst = make([]IArrayElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArrayElementContext)
		}
	}

	return tst
}

func (s *ElementListContext) ArrayElement(i int) IArrayElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArrayElementContext)
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
	p.EnterRule(localctx, 32, JsonpathParserRULE_elementList)
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
	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(342)
				p.Match(JsonpathParserT__6)
			}

		}
		p.SetState(347)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext())
	}
	p.SetState(349)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(348)
			p.Match(JsonpathParserSP)
		}

	}
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonpathParserT__0)|(1<<JsonpathParserT__1)|(1<<JsonpathParserT__4)|(1<<JsonpathParserT__9)|(1<<JsonpathParserT__11)|(1<<JsonpathParserT__12)|(1<<JsonpathParserT__13)|(1<<JsonpathParserT__14))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(JsonpathParserT__33-34))|(1<<(JsonpathParserT__34-34))|(1<<(JsonpathParserIdentifier-34))|(1<<(JsonpathParserEllipsis-34))|(1<<(JsonpathParserNullLiteral-34))|(1<<(JsonpathParserBooleanLiteral-34))|(1<<(JsonpathParserDecimalLiteral-34))|(1<<(JsonpathParserStringLiteral-34))|(1<<(JsonpathParserHexIntegerLiteral-34))|(1<<(JsonpathParserOctalIntegerLiteral2-34))|(1<<(JsonpathParserBinaryIntegerLiteral-34)))) != 0) {
		{
			p.SetState(351)
			p.ArrayElement()
		}

	}
	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 62, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(355)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == JsonpathParserSP {
				{
					p.SetState(354)
					p.Match(JsonpathParserSP)
				}

			}
			p.SetState(358)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == JsonpathParserT__6 {
				{
					p.SetState(357)
					p.Match(JsonpathParserT__6)
				}

				p.SetState(360)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(363)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == JsonpathParserSP {
				{
					p.SetState(362)
					p.Match(JsonpathParserSP)
				}

			}
			{
				p.SetState(365)
				p.ArrayElement()
			}

		}
		p.SetState(370)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 62, p.GetParserRuleContext())
	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JsonpathParserT__6 {
		{
			p.SetState(371)
			p.Match(JsonpathParserT__6)
		}

		p.SetState(376)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IArrayElementContext is an interface to support dynamic dispatch.
type IArrayElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayElementContext differentiates from other interfaces.
	IsArrayElementContext()
}

type ArrayElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayElementContext() *ArrayElementContext {
	var p = new(ArrayElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_arrayElement
	return p
}

func (*ArrayElementContext) IsArrayElementContext() {}

func NewArrayElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayElementContext {
	var p = new(ArrayElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_arrayElement

	return p
}

func (s *ArrayElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayElementContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ArrayElementContext) Ellipsis() antlr.TerminalNode {
	return s.GetToken(JsonpathParserEllipsis, 0)
}

func (s *ArrayElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterArrayElement(s)
	}
}

func (s *ArrayElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitArrayElement(s)
	}
}

func (p *JsonpathParser) ArrayElement() (localctx IArrayElementContext) {
	localctx = NewArrayElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, JsonpathParserRULE_arrayElement)
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
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JsonpathParserEllipsis {
		{
			p.SetState(377)
			p.Match(JsonpathParserEllipsis)
		}

	}
	{
		p.SetState(380)
		p.singleExpression(0)
	}

	return localctx
}

// IObjectLiteralContext is an interface to support dynamic dispatch.
type IObjectLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectLiteralContext differentiates from other interfaces.
	IsObjectLiteralContext()
}

type ObjectLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectLiteralContext() *ObjectLiteralContext {
	var p = new(ObjectLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_objectLiteral
	return p
}

func (*ObjectLiteralContext) IsObjectLiteralContext() {}

func NewObjectLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectLiteralContext {
	var p = new(ObjectLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_objectLiteral

	return p
}

func (s *ObjectLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectLiteralContext) AllPropertyAssignment() []IPropertyAssignmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPropertyAssignmentContext)(nil)).Elem())
	var tst = make([]IPropertyAssignmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPropertyAssignmentContext)
		}
	}

	return tst
}

func (s *ObjectLiteralContext) PropertyAssignment(i int) IPropertyAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyAssignmentContext)(nil)).Elem(), i)

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
	p.EnterRule(localctx, 36, JsonpathParserRULE_objectLiteral)
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
		p.SetState(382)
		p.Match(JsonpathParserT__34)
	}
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JsonpathParserT__0)|(1<<JsonpathParserT__1)|(1<<JsonpathParserT__4)|(1<<JsonpathParserT__9)|(1<<JsonpathParserT__11)|(1<<JsonpathParserT__12)|(1<<JsonpathParserT__13)|(1<<JsonpathParserT__14))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(JsonpathParserT__33-34))|(1<<(JsonpathParserT__34-34))|(1<<(JsonpathParserIdentifier-34))|(1<<(JsonpathParserEllipsis-34))|(1<<(JsonpathParserNullLiteral-34))|(1<<(JsonpathParserBooleanLiteral-34))|(1<<(JsonpathParserDecimalLiteral-34))|(1<<(JsonpathParserStringLiteral-34))|(1<<(JsonpathParserHexIntegerLiteral-34))|(1<<(JsonpathParserOctalIntegerLiteral2-34))|(1<<(JsonpathParserBinaryIntegerLiteral-34)))) != 0) {
		{
			p.SetState(383)
			p.PropertyAssignment()
		}
		p.SetState(388)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 65, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(384)
					p.Match(JsonpathParserT__6)
				}
				{
					p.SetState(385)
					p.PropertyAssignment()
				}

			}
			p.SetState(390)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 65, p.GetParserRuleContext())
		}
		p.SetState(392)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonpathParserT__6 {
			{
				p.SetState(391)
				p.Match(JsonpathParserT__6)
			}

		}

	}
	{
		p.SetState(396)
		p.Match(JsonpathParserT__35)
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyAssignmentContext() *PropertyAssignmentContext {
	var p = new(PropertyAssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_propertyAssignment
	return p
}

func (*PropertyAssignmentContext) IsPropertyAssignmentContext() {}

func NewPropertyAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyAssignmentContext {
	var p = new(PropertyAssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_propertyAssignment

	return p
}

func (s *PropertyAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyAssignmentContext) CopyFrom(ctx *PropertyAssignmentContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PropertyAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PropertyExpressionAssignmentContext struct {
	*PropertyAssignmentContext
}

func NewPropertyExpressionAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropertyExpressionAssignmentContext {
	var p = new(PropertyExpressionAssignmentContext)

	p.PropertyAssignmentContext = NewEmptyPropertyAssignmentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropertyAssignmentContext))

	return p
}

func (s *PropertyExpressionAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyExpressionAssignmentContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *PropertyExpressionAssignmentContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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
	*PropertyAssignmentContext
}

func NewComputedPropertyExpressionAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComputedPropertyExpressionAssignmentContext {
	var p = new(ComputedPropertyExpressionAssignmentContext)

	p.PropertyAssignmentContext = NewEmptyPropertyAssignmentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropertyAssignmentContext))

	return p
}

func (s *ComputedPropertyExpressionAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComputedPropertyExpressionAssignmentContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *ComputedPropertyExpressionAssignmentContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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

type PropertyShorthandContext struct {
	*PropertyAssignmentContext
}

func NewPropertyShorthandContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropertyShorthandContext {
	var p = new(PropertyShorthandContext)

	p.PropertyAssignmentContext = NewEmptyPropertyAssignmentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropertyAssignmentContext))

	return p
}

func (s *PropertyShorthandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyShorthandContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *PropertyShorthandContext) Ellipsis() antlr.TerminalNode {
	return s.GetToken(JsonpathParserEllipsis, 0)
}

func (s *PropertyShorthandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.EnterPropertyShorthand(s)
	}
}

func (s *PropertyShorthandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonpathListener); ok {
		listenerT.ExitPropertyShorthand(s)
	}
}

func (p *JsonpathParser) PropertyAssignment() (localctx IPropertyAssignmentContext) {
	localctx = NewPropertyAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, JsonpathParserRULE_propertyAssignment)
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

	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 69, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPropertyExpressionAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(398)
			p.PropertyName()
		}
		{
			p.SetState(399)
			p.Match(JsonpathParserT__7)
		}
		{
			p.SetState(400)
			p.singleExpression(0)
		}

	case 2:
		localctx = NewComputedPropertyExpressionAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(402)
			p.Match(JsonpathParserT__4)
		}
		{
			p.SetState(403)
			p.singleExpression(0)
		}
		{
			p.SetState(404)
			p.Match(JsonpathParserT__5)
		}
		{
			p.SetState(405)
			p.Match(JsonpathParserT__7)
		}
		{
			p.SetState(406)
			p.singleExpression(0)
		}

	case 3:
		localctx = NewPropertyShorthandContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(409)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JsonpathParserEllipsis {
			{
				p.SetState(408)
				p.Match(JsonpathParserEllipsis)
			}

		}
		{
			p.SetState(411)
			p.singleExpression(0)
		}

	}

	return localctx
}

// IPropertyNameContext is an interface to support dynamic dispatch.
type IPropertyNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyNameContext differentiates from other interfaces.
	IsPropertyNameContext()
}

type PropertyNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyNameContext() *PropertyNameContext {
	var p = new(PropertyNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_propertyName
	return p
}

func (*PropertyNameContext) IsPropertyNameContext() {}

func NewPropertyNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyNameContext {
	var p = new(PropertyNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonpathParserRULE_propertyName

	return p
}

func (s *PropertyNameContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyNameContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *PropertyNameContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(JsonpathParserStringLiteral, 0)
}

func (s *PropertyNameContext) NumericLiteral() INumericLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *PropertyNameContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
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
	p.EnterRule(localctx, 40, JsonpathParserRULE_propertyName)

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

	p.SetState(421)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonpathParserT__1, JsonpathParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(414)
			p.Identifier()
		}

	case JsonpathParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(415)
			p.Match(JsonpathParserStringLiteral)
		}

	case JsonpathParserDecimalLiteral, JsonpathParserHexIntegerLiteral, JsonpathParserOctalIntegerLiteral2, JsonpathParserBinaryIntegerLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(416)
			p.NumericLiteral()
		}

	case JsonpathParserT__4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(417)
			p.Match(JsonpathParserT__4)
		}
		{
			p.SetState(418)
			p.singleExpression(0)
		}
		{
			p.SetState(419)
			p.Match(JsonpathParserT__5)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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

func (s *LiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(JsonpathParserStringLiteral, 0)
}

func (s *LiteralContext) NumericLiteral() INumericLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
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
	p.EnterRule(localctx, 42, JsonpathParserRULE_literal)

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

	p.SetState(427)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JsonpathParserNullLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(423)
			p.Match(JsonpathParserNullLiteral)
		}

	case JsonpathParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(424)
			p.Match(JsonpathParserBooleanLiteral)
		}

	case JsonpathParserStringLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(425)
			p.Match(JsonpathParserStringLiteral)
		}

	case JsonpathParserDecimalLiteral, JsonpathParserHexIntegerLiteral, JsonpathParserOctalIntegerLiteral2, JsonpathParserBinaryIntegerLiteral:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(426)
			p.NumericLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	p.EnterRule(localctx, 44, JsonpathParserRULE_identifier)
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
		p.SetState(429)
		_la = p.GetTokenStream().LA(1)

		if !(_la == JsonpathParserT__1 || _la == JsonpathParserIdentifier) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INumericLiteralContext is an interface to support dynamic dispatch.
type INumericLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericLiteralContext differentiates from other interfaces.
	IsNumericLiteralContext()
}

type NumericLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralContext() *NumericLiteralContext {
	var p = new(NumericLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JsonpathParserRULE_numericLiteral
	return p
}

func (*NumericLiteralContext) IsNumericLiteralContext() {}

func NewNumericLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralContext {
	var p = new(NumericLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	p.EnterRule(localctx, 46, JsonpathParserRULE_numericLiteral)
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
		p.SetState(431)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(JsonpathParserDecimalLiteral-41))|(1<<(JsonpathParserHexIntegerLiteral-41))|(1<<(JsonpathParserOctalIntegerLiteral2-41))|(1<<(JsonpathParserBinaryIntegerLiteral-41)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *JsonpathParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 14:
		var t *SingleExpressionContext = nil
		if localctx != nil {
			t = localctx.(*SingleExpressionContext)
		}
		return p.SingleExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *JsonpathParser) SingleExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 27)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 20)

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
		return p.Precpred(p.GetParserRuleContext(), 26)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 25)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
