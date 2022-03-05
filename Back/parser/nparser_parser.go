// Code generated from Nparser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Nparser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "github.com/colegno/arraylist"
import "Back/analizador/Ast"
import "Back/analizador/expresiones"
import "Back/analizador/instrucciones"
import "Back/analizador/exp_ins"
import "Back/analizador/transferencia"
import "Back/analizador/bucles"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 91, 463,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 7, 4,
	68, 10, 4, 12, 4, 14, 4, 71, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 107, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 5, 6, 176, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 5, 7, 188, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 214, 10, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 7, 8, 246, 10, 8, 12, 8, 14, 8, 249, 11, 8, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 5, 9, 265, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 293, 10,
	10, 3, 11, 6, 11, 296, 10, 11, 13, 11, 14, 11, 297, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5,
	13, 334, 10, 13, 3, 14, 6, 14, 337, 10, 14, 13, 14, 14, 14, 338, 3, 14,
	3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 358, 10, 16, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 6, 18, 368, 10, 18, 13, 18, 14,
	18, 369, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 386, 10, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 7, 20, 393, 10, 20, 12, 20, 14, 20, 396, 11, 20, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 6, 22, 406, 10, 22, 13,
	22, 14, 22, 407, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 424, 10, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 7, 24, 431, 10, 24, 12, 24, 14, 24, 434, 11, 24,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 442, 10, 25, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 450, 10, 26, 3, 27, 3, 27, 3, 27,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 2, 5, 14,
	38, 46, 30, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
	34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 2, 7, 4, 2, 76, 76, 78,
	78, 3, 2, 72, 74, 3, 2, 76, 77, 4, 2, 64, 68, 70, 70, 4, 2, 68, 68, 70,
	70, 2, 490, 2, 58, 3, 2, 2, 2, 4, 61, 3, 2, 2, 2, 6, 69, 3, 2, 2, 2, 8,
	106, 3, 2, 2, 2, 10, 175, 3, 2, 2, 2, 12, 187, 3, 2, 2, 2, 14, 213, 3,
	2, 2, 2, 16, 264, 3, 2, 2, 2, 18, 292, 3, 2, 2, 2, 20, 295, 3, 2, 2, 2,
	22, 301, 3, 2, 2, 2, 24, 333, 3, 2, 2, 2, 26, 336, 3, 2, 2, 2, 28, 342,
	3, 2, 2, 2, 30, 357, 3, 2, 2, 2, 32, 359, 3, 2, 2, 2, 34, 367, 3, 2, 2,
	2, 36, 373, 3, 2, 2, 2, 38, 385, 3, 2, 2, 2, 40, 397, 3, 2, 2, 2, 42, 405,
	3, 2, 2, 2, 44, 411, 3, 2, 2, 2, 46, 423, 3, 2, 2, 2, 48, 441, 3, 2, 2,
	2, 50, 449, 3, 2, 2, 2, 52, 451, 3, 2, 2, 2, 54, 454, 3, 2, 2, 2, 56, 458,
	3, 2, 2, 2, 58, 59, 5, 6, 4, 2, 59, 60, 8, 2, 1, 2, 60, 3, 3, 2, 2, 2,
	61, 62, 7, 82, 2, 2, 62, 63, 5, 6, 4, 2, 63, 64, 7, 83, 2, 2, 64, 65, 8,
	3, 1, 2, 65, 5, 3, 2, 2, 2, 66, 68, 5, 8, 5, 2, 67, 66, 3, 2, 2, 2, 68,
	71, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 72, 3, 2, 2,
	2, 71, 69, 3, 2, 2, 2, 72, 73, 8, 4, 1, 2, 73, 7, 3, 2, 2, 2, 74, 75, 5,
	14, 8, 2, 75, 76, 8, 5, 1, 2, 76, 107, 3, 2, 2, 2, 77, 78, 5, 10, 6, 2,
	78, 79, 7, 63, 2, 2, 79, 80, 8, 5, 1, 2, 80, 107, 3, 2, 2, 2, 81, 82, 5,
	12, 7, 2, 82, 83, 7, 63, 2, 2, 83, 84, 8, 5, 1, 2, 84, 107, 3, 2, 2, 2,
	85, 86, 5, 18, 10, 2, 86, 87, 8, 5, 1, 2, 87, 107, 3, 2, 2, 2, 88, 89,
	5, 32, 17, 2, 89, 90, 8, 5, 1, 2, 90, 107, 3, 2, 2, 2, 91, 92, 5, 54, 28,
	2, 92, 93, 8, 5, 1, 2, 93, 107, 3, 2, 2, 2, 94, 95, 5, 50, 26, 2, 95, 96,
	7, 63, 2, 2, 96, 97, 8, 5, 1, 2, 97, 107, 3, 2, 2, 2, 98, 99, 5, 52, 27,
	2, 99, 100, 7, 63, 2, 2, 100, 101, 8, 5, 1, 2, 101, 107, 3, 2, 2, 2, 102,
	103, 5, 48, 25, 2, 103, 104, 7, 63, 2, 2, 104, 105, 8, 5, 1, 2, 105, 107,
	3, 2, 2, 2, 106, 74, 3, 2, 2, 2, 106, 77, 3, 2, 2, 2, 106, 81, 3, 2, 2,
	2, 106, 85, 3, 2, 2, 2, 106, 88, 3, 2, 2, 2, 106, 91, 3, 2, 2, 2, 106,
	94, 3, 2, 2, 2, 106, 98, 3, 2, 2, 2, 106, 102, 3, 2, 2, 2, 107, 9, 3, 2,
	2, 2, 108, 109, 7, 17, 2, 2, 109, 110, 7, 51, 2, 2, 110, 111, 7, 71, 2,
	2, 111, 112, 5, 14, 8, 2, 112, 113, 8, 6, 1, 2, 113, 176, 3, 2, 2, 2, 114,
	115, 7, 17, 2, 2, 115, 116, 7, 51, 2, 2, 116, 117, 7, 71, 2, 2, 117, 118,
	5, 30, 16, 2, 118, 119, 8, 6, 1, 2, 119, 176, 3, 2, 2, 2, 120, 121, 7,
	17, 2, 2, 121, 122, 7, 51, 2, 2, 122, 123, 7, 59, 2, 2, 123, 124, 5, 16,
	9, 2, 124, 125, 7, 71, 2, 2, 125, 126, 5, 14, 8, 2, 126, 127, 8, 6, 1,
	2, 127, 176, 3, 2, 2, 2, 128, 129, 7, 17, 2, 2, 129, 130, 7, 51, 2, 2,
	130, 131, 7, 59, 2, 2, 131, 132, 5, 16, 9, 2, 132, 133, 7, 71, 2, 2, 133,
	134, 5, 30, 16, 2, 134, 135, 8, 6, 1, 2, 135, 176, 3, 2, 2, 2, 136, 137,
	7, 17, 2, 2, 137, 138, 7, 16, 2, 2, 138, 139, 7, 51, 2, 2, 139, 140, 7,
	71, 2, 2, 140, 141, 5, 14, 8, 2, 141, 142, 8, 6, 1, 2, 142, 176, 3, 2,
	2, 2, 143, 144, 7, 17, 2, 2, 144, 145, 7, 16, 2, 2, 145, 146, 7, 51, 2,
	2, 146, 147, 7, 71, 2, 2, 147, 148, 5, 30, 16, 2, 148, 149, 8, 6, 1, 2,
	149, 176, 3, 2, 2, 2, 150, 151, 7, 17, 2, 2, 151, 152, 7, 16, 2, 2, 152,
	153, 7, 51, 2, 2, 153, 154, 7, 59, 2, 2, 154, 155, 5, 16, 9, 2, 155, 156,
	8, 6, 1, 2, 156, 176, 3, 2, 2, 2, 157, 158, 7, 17, 2, 2, 158, 159, 7, 16,
	2, 2, 159, 160, 7, 51, 2, 2, 160, 161, 7, 59, 2, 2, 161, 162, 5, 16, 9,
	2, 162, 163, 7, 71, 2, 2, 163, 164, 5, 14, 8, 2, 164, 165, 8, 6, 1, 2,
	165, 176, 3, 2, 2, 2, 166, 167, 7, 17, 2, 2, 167, 168, 7, 16, 2, 2, 168,
	169, 7, 51, 2, 2, 169, 170, 7, 59, 2, 2, 170, 171, 5, 16, 9, 2, 171, 172,
	7, 71, 2, 2, 172, 173, 5, 30, 16, 2, 173, 174, 8, 6, 1, 2, 174, 176, 3,
	2, 2, 2, 175, 108, 3, 2, 2, 2, 175, 114, 3, 2, 2, 2, 175, 120, 3, 2, 2,
	2, 175, 128, 3, 2, 2, 2, 175, 136, 3, 2, 2, 2, 175, 143, 3, 2, 2, 2, 175,
	150, 3, 2, 2, 2, 175, 157, 3, 2, 2, 2, 175, 166, 3, 2, 2, 2, 176, 11, 3,
	2, 2, 2, 177, 178, 7, 51, 2, 2, 178, 179, 7, 71, 2, 2, 179, 180, 5, 14,
	8, 2, 180, 181, 8, 7, 1, 2, 181, 188, 3, 2, 2, 2, 182, 183, 7, 51, 2, 2,
	183, 184, 7, 71, 2, 2, 184, 185, 5, 24, 13, 2, 185, 186, 8, 7, 1, 2, 186,
	188, 3, 2, 2, 2, 187, 177, 3, 2, 2, 2, 187, 182, 3, 2, 2, 2, 188, 13, 3,
	2, 2, 2, 189, 190, 8, 8, 1, 2, 190, 191, 9, 2, 2, 2, 191, 192, 5, 14, 8,
	17, 192, 193, 8, 8, 1, 2, 193, 214, 3, 2, 2, 2, 194, 195, 7, 80, 2, 2,
	195, 196, 5, 14, 8, 2, 196, 197, 7, 81, 2, 2, 197, 198, 8, 8, 1, 2, 198,
	214, 3, 2, 2, 2, 199, 200, 7, 51, 2, 2, 200, 214, 8, 8, 1, 2, 201, 202,
	7, 21, 2, 2, 202, 214, 8, 8, 1, 2, 203, 204, 7, 22, 2, 2, 204, 214, 8,
	8, 1, 2, 205, 206, 7, 88, 2, 2, 206, 214, 8, 8, 1, 2, 207, 208, 7, 49,
	2, 2, 208, 214, 8, 8, 1, 2, 209, 210, 7, 48, 2, 2, 210, 214, 8, 8, 1, 2,
	211, 212, 7, 86, 2, 2, 212, 214, 8, 8, 1, 2, 213, 189, 3, 2, 2, 2, 213,
	194, 3, 2, 2, 2, 213, 199, 3, 2, 2, 2, 213, 201, 3, 2, 2, 2, 213, 203,
	3, 2, 2, 2, 213, 205, 3, 2, 2, 2, 213, 207, 3, 2, 2, 2, 213, 209, 3, 2,
	2, 2, 213, 211, 3, 2, 2, 2, 214, 247, 3, 2, 2, 2, 215, 216, 12, 16, 2,
	2, 216, 217, 9, 3, 2, 2, 217, 218, 5, 14, 8, 17, 218, 219, 8, 8, 1, 2,
	219, 246, 3, 2, 2, 2, 220, 221, 12, 15, 2, 2, 221, 222, 9, 4, 2, 2, 222,
	223, 5, 14, 8, 16, 223, 224, 8, 8, 1, 2, 224, 246, 3, 2, 2, 2, 225, 226,
	12, 14, 2, 2, 226, 227, 9, 5, 2, 2, 227, 228, 5, 14, 8, 15, 228, 229, 8,
	8, 1, 2, 229, 246, 3, 2, 2, 2, 230, 231, 12, 13, 2, 2, 231, 232, 9, 6,
	2, 2, 232, 233, 5, 14, 8, 14, 233, 234, 8, 8, 1, 2, 234, 246, 3, 2, 2,
	2, 235, 236, 12, 12, 2, 2, 236, 237, 7, 56, 2, 2, 237, 238, 5, 14, 8, 13,
	238, 239, 8, 8, 1, 2, 239, 246, 3, 2, 2, 2, 240, 241, 12, 11, 2, 2, 241,
	242, 7, 54, 2, 2, 242, 243, 5, 14, 8, 12, 243, 244, 8, 8, 1, 2, 244, 246,
	3, 2, 2, 2, 245, 215, 3, 2, 2, 2, 245, 220, 3, 2, 2, 2, 245, 225, 3, 2,
	2, 2, 245, 230, 3, 2, 2, 2, 245, 235, 3, 2, 2, 2, 245, 240, 3, 2, 2, 2,
	246, 249, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248,
	15, 3, 2, 2, 2, 249, 247, 3, 2, 2, 2, 250, 251, 7, 3, 2, 2, 251, 265, 8,
	9, 1, 2, 252, 253, 7, 4, 2, 2, 253, 265, 8, 9, 1, 2, 254, 255, 7, 6, 2,
	2, 255, 265, 8, 9, 1, 2, 256, 257, 7, 5, 2, 2, 257, 265, 8, 9, 1, 2, 258,
	259, 7, 7, 2, 2, 259, 265, 8, 9, 1, 2, 260, 261, 7, 8, 2, 2, 261, 265,
	8, 9, 1, 2, 262, 263, 7, 9, 2, 2, 263, 265, 8, 9, 1, 2, 264, 250, 3, 2,
	2, 2, 264, 252, 3, 2, 2, 2, 264, 254, 3, 2, 2, 2, 264, 256, 3, 2, 2, 2,
	264, 258, 3, 2, 2, 2, 264, 260, 3, 2, 2, 2, 264, 262, 3, 2, 2, 2, 265,
	17, 3, 2, 2, 2, 266, 267, 7, 36, 2, 2, 267, 268, 5, 14, 8, 2, 268, 269,
	5, 4, 3, 2, 269, 270, 8, 10, 1, 2, 270, 293, 3, 2, 2, 2, 271, 272, 7, 36,
	2, 2, 272, 273, 5, 14, 8, 2, 273, 274, 5, 4, 3, 2, 274, 275, 7, 37, 2,
	2, 275, 276, 5, 4, 3, 2, 276, 277, 8, 10, 1, 2, 277, 293, 3, 2, 2, 2, 278,
	279, 7, 36, 2, 2, 279, 280, 5, 14, 8, 2, 280, 281, 5, 4, 3, 2, 281, 282,
	5, 20, 11, 2, 282, 283, 8, 10, 1, 2, 283, 293, 3, 2, 2, 2, 284, 285, 7,
	36, 2, 2, 285, 286, 5, 14, 8, 2, 286, 287, 5, 4, 3, 2, 287, 288, 5, 20,
	11, 2, 288, 289, 7, 37, 2, 2, 289, 290, 5, 4, 3, 2, 290, 291, 8, 10, 1,
	2, 291, 293, 3, 2, 2, 2, 292, 266, 3, 2, 2, 2, 292, 271, 3, 2, 2, 2, 292,
	278, 3, 2, 2, 2, 292, 284, 3, 2, 2, 2, 293, 19, 3, 2, 2, 2, 294, 296, 5,
	22, 12, 2, 295, 294, 3, 2, 2, 2, 296, 297, 3, 2, 2, 2, 297, 295, 3, 2,
	2, 2, 297, 298, 3, 2, 2, 2, 298, 299, 3, 2, 2, 2, 299, 300, 8, 11, 1, 2,
	300, 21, 3, 2, 2, 2, 301, 302, 7, 37, 2, 2, 302, 303, 7, 36, 2, 2, 303,
	304, 5, 14, 8, 2, 304, 305, 5, 4, 3, 2, 305, 306, 8, 12, 1, 2, 306, 23,
	3, 2, 2, 2, 307, 308, 7, 36, 2, 2, 308, 309, 5, 14, 8, 2, 309, 310, 5,
	4, 3, 2, 310, 311, 8, 13, 1, 2, 311, 334, 3, 2, 2, 2, 312, 313, 7, 36,
	2, 2, 313, 314, 5, 14, 8, 2, 314, 315, 5, 4, 3, 2, 315, 316, 7, 37, 2,
	2, 316, 317, 5, 4, 3, 2, 317, 318, 8, 13, 1, 2, 318, 334, 3, 2, 2, 2, 319,
	320, 7, 36, 2, 2, 320, 321, 5, 14, 8, 2, 321, 322, 5, 4, 3, 2, 322, 323,
	5, 26, 14, 2, 323, 324, 8, 13, 1, 2, 324, 334, 3, 2, 2, 2, 325, 326, 7,
	36, 2, 2, 326, 327, 5, 14, 8, 2, 327, 328, 5, 4, 3, 2, 328, 329, 5, 26,
	14, 2, 329, 330, 7, 37, 2, 2, 330, 331, 5, 4, 3, 2, 331, 332, 8, 13, 1,
	2, 332, 334, 3, 2, 2, 2, 333, 307, 3, 2, 2, 2, 333, 312, 3, 2, 2, 2, 333,
	319, 3, 2, 2, 2, 333, 325, 3, 2, 2, 2, 334, 25, 3, 2, 2, 2, 335, 337, 5,
	28, 15, 2, 336, 335, 3, 2, 2, 2, 337, 338, 3, 2, 2, 2, 338, 336, 3, 2,
	2, 2, 338, 339, 3, 2, 2, 2, 339, 340, 3, 2, 2, 2, 340, 341, 8, 14, 1, 2,
	341, 27, 3, 2, 2, 2, 342, 343, 7, 37, 2, 2, 343, 344, 7, 36, 2, 2, 344,
	345, 5, 14, 8, 2, 345, 346, 5, 4, 3, 2, 346, 347, 8, 15, 1, 2, 347, 29,
	3, 2, 2, 2, 348, 349, 5, 24, 13, 2, 349, 350, 8, 16, 1, 2, 350, 358, 3,
	2, 2, 2, 351, 352, 5, 40, 21, 2, 352, 353, 8, 16, 1, 2, 353, 358, 3, 2,
	2, 2, 354, 355, 5, 56, 29, 2, 355, 356, 8, 16, 1, 2, 356, 358, 3, 2, 2,
	2, 357, 348, 3, 2, 2, 2, 357, 351, 3, 2, 2, 2, 357, 354, 3, 2, 2, 2, 358,
	31, 3, 2, 2, 2, 359, 360, 7, 38, 2, 2, 360, 361, 5, 14, 8, 2, 361, 362,
	7, 82, 2, 2, 362, 363, 5, 34, 18, 2, 363, 364, 7, 83, 2, 2, 364, 365, 8,
	17, 1, 2, 365, 33, 3, 2, 2, 2, 366, 368, 5, 36, 19, 2, 367, 366, 3, 2,
	2, 2, 368, 369, 3, 2, 2, 2, 369, 367, 3, 2, 2, 2, 369, 370, 3, 2, 2, 2,
	370, 371, 3, 2, 2, 2, 371, 372, 8, 18, 1, 2, 372, 35, 3, 2, 2, 2, 373,
	374, 5, 38, 20, 2, 374, 375, 7, 69, 2, 2, 375, 376, 5, 4, 3, 2, 376, 377,
	7, 62, 2, 2, 377, 378, 8, 19, 1, 2, 378, 37, 3, 2, 2, 2, 379, 380, 8, 20,
	1, 2, 380, 381, 5, 14, 8, 2, 381, 382, 8, 20, 1, 2, 382, 386, 3, 2, 2,
	2, 383, 384, 7, 52, 2, 2, 384, 386, 8, 20, 1, 2, 385, 379, 3, 2, 2, 2,
	385, 383, 3, 2, 2, 2, 386, 394, 3, 2, 2, 2, 387, 388, 12, 5, 2, 2, 388,
	389, 7, 53, 2, 2, 389, 390, 5, 14, 8, 2, 390, 391, 8, 20, 1, 2, 391, 393,
	3, 2, 2, 2, 392, 387, 3, 2, 2, 2, 393, 396, 3, 2, 2, 2, 394, 392, 3, 2,
	2, 2, 394, 395, 3, 2, 2, 2, 395, 39, 3, 2, 2, 2, 396, 394, 3, 2, 2, 2,
	397, 398, 7, 38, 2, 2, 398, 399, 5, 14, 8, 2, 399, 400, 7, 82, 2, 2, 400,
	401, 5, 42, 22, 2, 401, 402, 7, 83, 2, 2, 402, 403, 8, 21, 1, 2, 403, 41,
	3, 2, 2, 2, 404, 406, 5, 44, 23, 2, 405, 404, 3, 2, 2, 2, 406, 407, 3,
	2, 2, 2, 407, 405, 3, 2, 2, 2, 407, 408, 3, 2, 2, 2, 408, 409, 3, 2, 2,
	2, 409, 410, 8, 22, 1, 2, 410, 43, 3, 2, 2, 2, 411, 412, 5, 46, 24, 2,
	412, 413, 7, 69, 2, 2, 413, 414, 5, 4, 3, 2, 414, 415, 7, 62, 2, 2, 415,
	416, 8, 23, 1, 2, 416, 45, 3, 2, 2, 2, 417, 418, 8, 24, 1, 2, 418, 419,
	5, 14, 8, 2, 419, 420, 8, 24, 1, 2, 420, 424, 3, 2, 2, 2, 421, 422, 7,
	52, 2, 2, 422, 424, 8, 24, 1, 2, 423, 417, 3, 2, 2, 2, 423, 421, 3, 2,
	2, 2, 424, 432, 3, 2, 2, 2, 425, 426, 12, 5, 2, 2, 426, 427, 7, 53, 2,
	2, 427, 428, 5, 14, 8, 2, 428, 429, 8, 24, 1, 2, 429, 431, 3, 2, 2, 2,
	430, 425, 3, 2, 2, 2, 431, 434, 3, 2, 2, 2, 432, 430, 3, 2, 2, 2, 432,
	433, 3, 2, 2, 2, 433, 47, 3, 2, 2, 2, 434, 432, 3, 2, 2, 2, 435, 436, 7,
	43, 2, 2, 436, 442, 8, 25, 1, 2, 437, 438, 7, 43, 2, 2, 438, 439, 5, 14,
	8, 2, 439, 440, 8, 25, 1, 2, 440, 442, 3, 2, 2, 2, 441, 435, 3, 2, 2, 2,
	441, 437, 3, 2, 2, 2, 442, 49, 3, 2, 2, 2, 443, 444, 7, 44, 2, 2, 444,
	450, 8, 26, 1, 2, 445, 446, 7, 44, 2, 2, 446, 447, 5, 14, 8, 2, 447, 448,
	8, 26, 1, 2, 448, 450, 3, 2, 2, 2, 449, 443, 3, 2, 2, 2, 449, 445, 3, 2,
	2, 2, 450, 51, 3, 2, 2, 2, 451, 452, 7, 45, 2, 2, 452, 453, 8, 27, 1, 2,
	453, 53, 3, 2, 2, 2, 454, 455, 7, 39, 2, 2, 455, 456, 5, 4, 3, 2, 456,
	457, 8, 28, 1, 2, 457, 55, 3, 2, 2, 2, 458, 459, 7, 39, 2, 2, 459, 460,
	5, 4, 3, 2, 460, 461, 8, 29, 1, 2, 461, 57, 3, 2, 2, 2, 23, 69, 106, 175,
	187, 213, 245, 247, 264, 292, 297, 333, 338, 357, 369, 385, 394, 407, 423,
	432, 441, 449,
}
var literalNames = []string{
	"", "'bool'", "'char'", "'f64'", "'i64'", "'&str'", "'String'", "'usize'",
	"'main'", "'powf'", "'pow'", "'as'", "'Vec'", "'vec'", "'mut'", "'let'",
	"'struct'", "'to_string'", "'to_owned'", "'true'", "'false'", "'println'",
	"'fn'", "'abs'", "'sqrt'", "'clone'", "'new'", "'len'", "'push'", "'remove'",
	"'contains'", "'insert'", "'capacity'", "'with_capacity'", "'if'", "'else'",
	"'match'", "'loop'", "'while'", "'for'", "'in'", "'return'", "'break'",
	"'continue'", "'mod'", "'pub'", "", "", "", "", "'_'", "'|'", "'||'", "'&'",
	"'&&'", "':?'", "'::'", "':'", "'..'", "'.'", "','", "';'", "'>='", "'>'",
	"'<='", "'<'", "'=='", "'=>'", "'!='", "'='", "'%'", "'*'", "'/'", "'->'",
	"'-'", "'+'", "'!'", "'?'", "'('", "')'", "'{'", "'}'", "'['", "']'",
}
var symbolicNames = []string{
	"", "BOOL", "CHAR", "F64", "I64", "STR", "STRING", "USIZE", "MAIN", "POWF",
	"POW", "AS", "VEC", "VEC_M", "MUT", "LET", "STRUCT", "TO_STRING", "TO_OWNED",
	"TRUE", "FALSE", "PRINT", "FN", "ABS", "SQRT", "CLONE", "NEW", "LEN", "PUSH",
	"REMOVE", "CONTAINS", "INSERT", "CAPACITY", "WITH_CAPACITY", "IF", "ELSE",
	"MATCH", "LOOP", "WHILE", "FOR", "IN", "RETURN", "BREAK", "CONTINUE", "MOD",
	"PUB", "NUMERO", "DECIMAL", "ID_CAMEL", "ID", "DEFAULT", "O", "OR", "AMPERSAND",
	"AND", "PRINT_OP_DEBUG", "DOBLE_DOSPUNTOS", "DOSPUNTOS", "RANGO", "PUNTO",
	"COMA", "PUNTOCOMA", "MAYOR_I", "MAYOR", "MENOR_I", "MENOR", "IGUALDAD",
	"CASE", "DISTINTO", "IGUAL", "MODULO", "MULTIPLICACION", "DIVISION", "FN_TIPO_RETORNO",
	"RESTA", "SUMA", "NOT", "PREGUNTA", "PAR_IZQ", "PAR_DER", "LLAVE_IZQ",
	"LLAVE_DER", "CORCHETE_IZQ", "CORCHETE_DER", "CADENA", "ASCII", "CARACTER",
	"WHITESPACE", "COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"inicio", "bloque", "instrucciones", "instruccion", "declaracion", "asignacion",
	"expresion", "tipo_dato", "control_if", "bloque_else_if", "else_if", "control_if_exp",
	"bloque_else_if_exp", "else_if_exp", "control_expresion", "control_match",
	"control_case", "cases", "case_match", "control_match_exp", "control_case_exp",
	"cases_exp", "case_match_exp", "ireturn", "ibreak", "icontinue", "control_loop",
	"control_loop_exp",
}

type Nparser struct {
	*antlr.BaseParser
}

// NewNparser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *Nparser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewNparser(input antlr.TokenStream) *Nparser {
	this := new(Nparser)
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
	this.GrammarFileName = "Nparser.g4"

	return this
}

// Nparser tokens.
const (
	NparserEOF             = antlr.TokenEOF
	NparserBOOL            = 1
	NparserCHAR            = 2
	NparserF64             = 3
	NparserI64             = 4
	NparserSTR             = 5
	NparserSTRING          = 6
	NparserUSIZE           = 7
	NparserMAIN            = 8
	NparserPOWF            = 9
	NparserPOW             = 10
	NparserAS              = 11
	NparserVEC             = 12
	NparserVEC_M           = 13
	NparserMUT             = 14
	NparserLET             = 15
	NparserSTRUCT          = 16
	NparserTO_STRING       = 17
	NparserTO_OWNED        = 18
	NparserTRUE            = 19
	NparserFALSE           = 20
	NparserPRINT           = 21
	NparserFN              = 22
	NparserABS             = 23
	NparserSQRT            = 24
	NparserCLONE           = 25
	NparserNEW             = 26
	NparserLEN             = 27
	NparserPUSH            = 28
	NparserREMOVE          = 29
	NparserCONTAINS        = 30
	NparserINSERT          = 31
	NparserCAPACITY        = 32
	NparserWITH_CAPACITY   = 33
	NparserIF              = 34
	NparserELSE            = 35
	NparserMATCH           = 36
	NparserLOOP            = 37
	NparserWHILE           = 38
	NparserFOR             = 39
	NparserIN              = 40
	NparserRETURN          = 41
	NparserBREAK           = 42
	NparserCONTINUE        = 43
	NparserMOD             = 44
	NparserPUB             = 45
	NparserNUMERO          = 46
	NparserDECIMAL         = 47
	NparserID_CAMEL        = 48
	NparserID              = 49
	NparserDEFAULT         = 50
	NparserO               = 51
	NparserOR              = 52
	NparserAMPERSAND       = 53
	NparserAND             = 54
	NparserPRINT_OP_DEBUG  = 55
	NparserDOBLE_DOSPUNTOS = 56
	NparserDOSPUNTOS       = 57
	NparserRANGO           = 58
	NparserPUNTO           = 59
	NparserCOMA            = 60
	NparserPUNTOCOMA       = 61
	NparserMAYOR_I         = 62
	NparserMAYOR           = 63
	NparserMENOR_I         = 64
	NparserMENOR           = 65
	NparserIGUALDAD        = 66
	NparserCASE            = 67
	NparserDISTINTO        = 68
	NparserIGUAL           = 69
	NparserMODULO          = 70
	NparserMULTIPLICACION  = 71
	NparserDIVISION        = 72
	NparserFN_TIPO_RETORNO = 73
	NparserRESTA           = 74
	NparserSUMA            = 75
	NparserNOT             = 76
	NparserPREGUNTA        = 77
	NparserPAR_IZQ         = 78
	NparserPAR_DER         = 79
	NparserLLAVE_IZQ       = 80
	NparserLLAVE_DER       = 81
	NparserCORCHETE_IZQ    = 82
	NparserCORCHETE_DER    = 83
	NparserCADENA          = 84
	NparserASCII           = 85
	NparserCARACTER        = 86
	NparserWHITESPACE      = 87
	NparserCOMMENT         = 88
	NparserLINE_COMMENT    = 89
)

// Nparser rules.
const (
	NparserRULE_inicio             = 0
	NparserRULE_bloque             = 1
	NparserRULE_instrucciones      = 2
	NparserRULE_instruccion        = 3
	NparserRULE_declaracion        = 4
	NparserRULE_asignacion         = 5
	NparserRULE_expresion          = 6
	NparserRULE_tipo_dato          = 7
	NparserRULE_control_if         = 8
	NparserRULE_bloque_else_if     = 9
	NparserRULE_else_if            = 10
	NparserRULE_control_if_exp     = 11
	NparserRULE_bloque_else_if_exp = 12
	NparserRULE_else_if_exp        = 13
	NparserRULE_control_expresion  = 14
	NparserRULE_control_match      = 15
	NparserRULE_control_case       = 16
	NparserRULE_cases              = 17
	NparserRULE_case_match         = 18
	NparserRULE_control_match_exp  = 19
	NparserRULE_control_case_exp   = 20
	NparserRULE_cases_exp          = 21
	NparserRULE_case_match_exp     = 22
	NparserRULE_ireturn            = 23
	NparserRULE_ibreak             = 24
	NparserRULE_icontinue          = 25
	NparserRULE_control_loop       = 26
	NparserRULE_control_loop_exp   = 27
)

// IInicioContext is an interface to support dynamic dispatch.
type IInicioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetLista returns the lista attribute.
	GetLista() *arraylist.List

	// SetLista sets the lista attribute.
	SetLista(*arraylist.List)

	// IsInicioContext differentiates from other interfaces.
	IsInicioContext()
}

type InicioContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	lista          *arraylist.List
	_instrucciones IInstruccionesContext
}

func NewEmptyInicioContext() *InicioContext {
	var p = new(InicioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_inicio
	return p
}

func (*InicioContext) IsInicioContext() {}

func NewInicioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InicioContext {
	var p = new(InicioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_inicio

	return p
}

func (s *InicioContext) GetParser() antlr.Parser { return s.parser }

func (s *InicioContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *InicioContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *InicioContext) GetLista() *arraylist.List { return s.lista }

func (s *InicioContext) SetLista(v *arraylist.List) { s.lista = v }

func (s *InicioContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *InicioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InicioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InicioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterInicio(s)
	}
}

func (s *InicioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitInicio(s)
	}
}

func (p *Nparser) Inicio() (localctx IInicioContext) {
	this := p
	_ = this

	localctx = NewInicioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NparserRULE_inicio)

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
		p.SetState(56)

		var _x = p.Instrucciones()

		localctx.(*InicioContext)._instrucciones = _x
	}

	localctx.(*InicioContext).lista = localctx.(*InicioContext).Get_instrucciones().GetList()

	return localctx
}

// IBloqueContext is an interface to support dynamic dispatch.
type IBloqueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetList returns the list attribute.
	GetList() *arraylist.List

	// SetList sets the list attribute.
	SetList(*arraylist.List)

	// IsBloqueContext differentiates from other interfaces.
	IsBloqueContext()
}

type BloqueContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	list           *arraylist.List
	_instrucciones IInstruccionesContext
}

func NewEmptyBloqueContext() *BloqueContext {
	var p = new(BloqueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_bloque
	return p
}

func (*BloqueContext) IsBloqueContext() {}

func NewBloqueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueContext {
	var p = new(BloqueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_bloque

	return p
}

func (s *BloqueContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *BloqueContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *BloqueContext) GetList() *arraylist.List { return s.list }

func (s *BloqueContext) SetList(v *arraylist.List) { s.list = v }

func (s *BloqueContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(NparserLLAVE_IZQ, 0)
}

func (s *BloqueContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *BloqueContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(NparserLLAVE_DER, 0)
}

func (s *BloqueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloqueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterBloque(s)
	}
}

func (s *BloqueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitBloque(s)
	}
}

func (p *Nparser) Bloque() (localctx IBloqueContext) {
	this := p
	_ = this

	localctx = NewBloqueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NparserRULE_bloque)

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
		p.SetState(59)
		p.Match(NparserLLAVE_IZQ)
	}
	{
		p.SetState(60)

		var _x = p.Instrucciones()

		localctx.(*BloqueContext)._instrucciones = _x
	}
	{
		p.SetState(61)
		p.Match(NparserLLAVE_DER)
	}

	localctx.(*BloqueContext).list = localctx.(*BloqueContext).Get_instrucciones().GetList()

	return localctx
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetE returns the e rule context list.
	GetE() []IInstruccionContext

	// SetE sets the e rule context list.
	SetE([]IInstruccionContext)

	// GetList returns the list attribute.
	GetList() *arraylist.List

	// SetList sets the list attribute.
	SetList(*arraylist.List)

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	list         *arraylist.List
	_instruccion IInstruccionContext
	e            []IInstruccionContext
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_instrucciones
	return p
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *InstruccionesContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *InstruccionesContext) GetE() []IInstruccionContext { return s.e }

func (s *InstruccionesContext) SetE(v []IInstruccionContext) { s.e = v }

func (s *InstruccionesContext) GetList() *arraylist.List { return s.list }

func (s *InstruccionesContext) SetList(v *arraylist.List) { s.list = v }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionContext)(nil)).Elem())
	var tst = make([]IInstruccionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionContext)
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterInstrucciones(s)
	}
}

func (s *InstruccionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitInstrucciones(s)
	}
}

func (p *Nparser) Instrucciones() (localctx IInstruccionesContext) {
	this := p
	_ = this

	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NparserRULE_instrucciones)

	localctx.(*InstruccionesContext).list = arraylist.New()

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
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NparserLET)|(1<<NparserTRUE)|(1<<NparserFALSE))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(NparserIF-34))|(1<<(NparserMATCH-34))|(1<<(NparserLOOP-34))|(1<<(NparserRETURN-34))|(1<<(NparserBREAK-34))|(1<<(NparserCONTINUE-34))|(1<<(NparserNUMERO-34))|(1<<(NparserDECIMAL-34))|(1<<(NparserID-34)))) != 0) || (((_la-74)&-(0x1f+1)) == 0 && ((1<<uint((_la-74)))&((1<<(NparserRESTA-74))|(1<<(NparserNOT-74))|(1<<(NparserPAR_IZQ-74))|(1<<(NparserCADENA-74))|(1<<(NparserCARACTER-74)))) != 0) {
		{
			p.SetState(64)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*InstruccionesContext).GetE()
	for _, e := range listInt {
		localctx.(*InstruccionesContext).list.Add(e.GetEx())
	}

	return localctx
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Get_declaracion returns the _declaracion rule contexts.
	Get_declaracion() IDeclaracionContext

	// Get_asignacion returns the _asignacion rule contexts.
	Get_asignacion() IAsignacionContext

	// Get_control_if returns the _control_if rule contexts.
	Get_control_if() IControl_ifContext

	// Get_control_match returns the _control_match rule contexts.
	Get_control_match() IControl_matchContext

	// Get_control_loop returns the _control_loop rule contexts.
	Get_control_loop() IControl_loopContext

	// Get_ibreak returns the _ibreak rule contexts.
	Get_ibreak() IIbreakContext

	// Get_icontinue returns the _icontinue rule contexts.
	Get_icontinue() IIcontinueContext

	// Get_ireturn returns the _ireturn rule contexts.
	Get_ireturn() IIreturnContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// Set_declaracion sets the _declaracion rule contexts.
	Set_declaracion(IDeclaracionContext)

	// Set_asignacion sets the _asignacion rule contexts.
	Set_asignacion(IAsignacionContext)

	// Set_control_if sets the _control_if rule contexts.
	Set_control_if(IControl_ifContext)

	// Set_control_match sets the _control_match rule contexts.
	Set_control_match(IControl_matchContext)

	// Set_control_loop sets the _control_loop rule contexts.
	Set_control_loop(IControl_loopContext)

	// Set_ibreak sets the _ibreak rule contexts.
	Set_ibreak(IIbreakContext)

	// Set_icontinue sets the _icontinue rule contexts.
	Set_icontinue(IIcontinueContext)

	// Set_ireturn sets the _ireturn rule contexts.
	Set_ireturn(IIreturnContext)

	// GetEx returns the ex attribute.
	GetEx() interface{}

	// SetEx sets the ex attribute.
	SetEx(interface{})

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	ex             interface{}
	_expresion     IExpresionContext
	_declaracion   IDeclaracionContext
	_asignacion    IAsignacionContext
	_control_if    IControl_ifContext
	_control_match IControl_matchContext
	_control_loop  IControl_loopContext
	_ibreak        IIbreakContext
	_icontinue     IIcontinueContext
	_ireturn       IIreturnContext
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_instruccion
	return p
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *InstruccionContext) Get_declaracion() IDeclaracionContext { return s._declaracion }

func (s *InstruccionContext) Get_asignacion() IAsignacionContext { return s._asignacion }

func (s *InstruccionContext) Get_control_if() IControl_ifContext { return s._control_if }

func (s *InstruccionContext) Get_control_match() IControl_matchContext { return s._control_match }

func (s *InstruccionContext) Get_control_loop() IControl_loopContext { return s._control_loop }

func (s *InstruccionContext) Get_ibreak() IIbreakContext { return s._ibreak }

func (s *InstruccionContext) Get_icontinue() IIcontinueContext { return s._icontinue }

func (s *InstruccionContext) Get_ireturn() IIreturnContext { return s._ireturn }

func (s *InstruccionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *InstruccionContext) Set_declaracion(v IDeclaracionContext) { s._declaracion = v }

func (s *InstruccionContext) Set_asignacion(v IAsignacionContext) { s._asignacion = v }

func (s *InstruccionContext) Set_control_if(v IControl_ifContext) { s._control_if = v }

func (s *InstruccionContext) Set_control_match(v IControl_matchContext) { s._control_match = v }

func (s *InstruccionContext) Set_control_loop(v IControl_loopContext) { s._control_loop = v }

func (s *InstruccionContext) Set_ibreak(v IIbreakContext) { s._ibreak = v }

func (s *InstruccionContext) Set_icontinue(v IIcontinueContext) { s._icontinue = v }

func (s *InstruccionContext) Set_ireturn(v IIreturnContext) { s._ireturn = v }

func (s *InstruccionContext) GetEx() interface{} { return s.ex }

func (s *InstruccionContext) SetEx(v interface{}) { s.ex = v }

func (s *InstruccionContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *InstruccionContext) Declaracion() IDeclaracionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaracionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaracionContext)
}

func (s *InstruccionContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(NparserPUNTOCOMA, 0)
}

func (s *InstruccionContext) Asignacion() IAsignacionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsignacionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *InstruccionContext) Control_if() IControl_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControl_ifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControl_ifContext)
}

func (s *InstruccionContext) Control_match() IControl_matchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControl_matchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControl_matchContext)
}

func (s *InstruccionContext) Control_loop() IControl_loopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControl_loopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControl_loopContext)
}

func (s *InstruccionContext) Ibreak() IIbreakContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIbreakContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIbreakContext)
}

func (s *InstruccionContext) Icontinue() IIcontinueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIcontinueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIcontinueContext)
}

func (s *InstruccionContext) Ireturn() IIreturnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIreturnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIreturnContext)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterInstruccion(s)
	}
}

func (s *InstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitInstruccion(s)
	}
}

func (p *Nparser) Instruccion() (localctx IInstruccionContext) {
	this := p
	_ = this

	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NparserRULE_instruccion)

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

	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)

			var _x = p.expresion(0)

			localctx.(*InstruccionContext)._expresion = _x
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_expresion().GetEx()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)

			var _x = p.Declaracion()

			localctx.(*InstruccionContext)._declaracion = _x
		}
		{
			p.SetState(76)
			p.Match(NparserPUNTOCOMA)
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_declaracion().GetEx()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(79)

			var _x = p.Asignacion()

			localctx.(*InstruccionContext)._asignacion = _x
		}
		{
			p.SetState(80)
			p.Match(NparserPUNTOCOMA)
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_asignacion().GetEx()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(83)

			var _x = p.Control_if()

			localctx.(*InstruccionContext)._control_if = _x
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_control_if().GetEx()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(86)

			var _x = p.Control_match()

			localctx.(*InstruccionContext)._control_match = _x
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_control_match().GetEx()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(89)

			var _x = p.Control_loop()

			localctx.(*InstruccionContext)._control_loop = _x
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_control_loop().GetEx()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(92)

			var _x = p.Ibreak()

			localctx.(*InstruccionContext)._ibreak = _x
		}
		{
			p.SetState(93)
			p.Match(NparserPUNTOCOMA)
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_ibreak().GetEx()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(96)

			var _x = p.Icontinue()

			localctx.(*InstruccionContext)._icontinue = _x
		}
		{
			p.SetState(97)
			p.Match(NparserPUNTOCOMA)
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_icontinue().GetEx()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(100)

			var _x = p.Ireturn()

			localctx.(*InstruccionContext)._ireturn = _x
		}
		{
			p.SetState(101)
			p.Match(NparserPUNTOCOMA)
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_ireturn().GetEx()

	}

	return localctx
}

// IDeclaracionContext is an interface to support dynamic dispatch.
type IDeclaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_LET returns the _LET token.
	Get_LET() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_LET sets the _LET token.
	Set_LET(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Get_control_expresion returns the _control_expresion rule contexts.
	Get_control_expresion() IControl_expresionContext

	// Get_tipo_dato returns the _tipo_dato rule contexts.
	Get_tipo_dato() ITipo_datoContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// Set_control_expresion sets the _control_expresion rule contexts.
	Set_control_expresion(IControl_expresionContext)

	// Set_tipo_dato sets the _tipo_dato rule contexts.
	Set_tipo_dato(ITipo_datoContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsDeclaracionContext differentiates from other interfaces.
	IsDeclaracionContext()
}

type DeclaracionContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	ex                 Ast.Instruccion
	_LET               antlr.Token
	_ID                antlr.Token
	_expresion         IExpresionContext
	_control_expresion IControl_expresionContext
	_tipo_dato         ITipo_datoContext
}

func NewEmptyDeclaracionContext() *DeclaracionContext {
	var p = new(DeclaracionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_declaracion
	return p
}

func (*DeclaracionContext) IsDeclaracionContext() {}

func NewDeclaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionContext {
	var p = new(DeclaracionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_declaracion

	return p
}

func (s *DeclaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionContext) Get_LET() antlr.Token { return s._LET }

func (s *DeclaracionContext) Get_ID() antlr.Token { return s._ID }

func (s *DeclaracionContext) Set_LET(v antlr.Token) { s._LET = v }

func (s *DeclaracionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *DeclaracionContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *DeclaracionContext) Get_control_expresion() IControl_expresionContext {
	return s._control_expresion
}

func (s *DeclaracionContext) Get_tipo_dato() ITipo_datoContext { return s._tipo_dato }

func (s *DeclaracionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *DeclaracionContext) Set_control_expresion(v IControl_expresionContext) {
	s._control_expresion = v
}

func (s *DeclaracionContext) Set_tipo_dato(v ITipo_datoContext) { s._tipo_dato = v }

func (s *DeclaracionContext) GetEx() Ast.Instruccion { return s.ex }

func (s *DeclaracionContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *DeclaracionContext) LET() antlr.TerminalNode {
	return s.GetToken(NparserLET, 0)
}

func (s *DeclaracionContext) ID() antlr.TerminalNode {
	return s.GetToken(NparserID, 0)
}

func (s *DeclaracionContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(NparserIGUAL, 0)
}

func (s *DeclaracionContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *DeclaracionContext) Control_expresion() IControl_expresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControl_expresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControl_expresionContext)
}

func (s *DeclaracionContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(NparserDOSPUNTOS, 0)
}

func (s *DeclaracionContext) Tipo_dato() ITipo_datoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITipo_datoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITipo_datoContext)
}

func (s *DeclaracionContext) MUT() antlr.TerminalNode {
	return s.GetToken(NparserMUT, 0)
}

func (s *DeclaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterDeclaracion(s)
	}
}

func (s *DeclaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitDeclaracion(s)
	}
}

func (p *Nparser) Declaracion() (localctx IDeclaracionContext) {
	this := p
	_ = this

	localctx = NewDeclaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NparserRULE_declaracion)

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

	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(106)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(107)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(108)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(109)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}

		fila := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetColumn()
			}
		}())
		localctx.(*DeclaracionContext).ex = instrucciones.NewDeclaracion((func() string {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetText()
			}
		}()), Ast.INDEFINIDO,
			false, false, Ast.VOID, localctx.(*DeclaracionContext).Get_expresion().GetEx(), fila, columna)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(113)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(114)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(115)

			var _x = p.Control_expresion()

			localctx.(*DeclaracionContext)._control_expresion = _x
		}

		fila := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetColumn()
			}
		}())
		localctx.(*DeclaracionContext).ex = instrucciones.NewDeclaracion((func() string {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetText()
			}
		}()), Ast.INDEFINIDO,
			false, false, Ast.VOID, localctx.(*DeclaracionContext).Get_control_expresion().GetEx(), fila, columna)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(119)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(120)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(121)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(122)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(123)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}

		fila := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetColumn()
			}
		}())
		localctx.(*DeclaracionContext).ex = instrucciones.NewDeclaracion((func() string {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_tipo_dato().GetEx(),
			false, false, Ast.VOID, localctx.(*DeclaracionContext).Get_expresion().GetEx(), fila, columna)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(126)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(127)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(128)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(129)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(130)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(131)

			var _x = p.Control_expresion()

			localctx.(*DeclaracionContext)._control_expresion = _x
		}

		fila := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetColumn()
			}
		}())
		localctx.(*DeclaracionContext).ex = instrucciones.NewDeclaracion((func() string {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_tipo_dato().GetEx(),
			false, false, Ast.VOID, localctx.(*DeclaracionContext).Get_control_expresion().GetEx(), fila, columna)

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(134)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(135)
			p.Match(NparserMUT)
		}
		{
			p.SetState(136)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(137)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(138)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}

		fila := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetColumn()
			}
		}())
		localctx.(*DeclaracionContext).ex = instrucciones.NewDeclaracion((func() string {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetText()
			}
		}()), Ast.INDEFINIDO,
			true, false, Ast.VOID, localctx.(*DeclaracionContext).Get_expresion().GetEx(), fila, columna)

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(141)
			p.Match(NparserLET)
		}
		{
			p.SetState(142)
			p.Match(NparserMUT)
		}
		{
			p.SetState(143)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(144)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(145)

			var _x = p.Control_expresion()

			localctx.(*DeclaracionContext)._control_expresion = _x
		}

		fila := (func() int {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetColumn()
			}
		}())
		localctx.(*DeclaracionContext).ex = instrucciones.NewDeclaracion((func() string {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetText()
			}
		}()), Ast.INDEFINIDO,
			true, false, Ast.VOID, localctx.(*DeclaracionContext).Get_control_expresion().GetEx(), fila, columna)

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(148)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(149)
			p.Match(NparserMUT)
		}
		{
			p.SetState(150)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(151)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(152)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}

		fila := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetColumn()
			}
		}())
		valor := expresiones.NewPrimitivo(nil, Ast.NULL, fila, columna)
		localctx.(*DeclaracionContext).ex = instrucciones.NewDeclaracion((func() string {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_tipo_dato().GetEx(),
			true, false, Ast.VOID, valor, fila, columna)

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(155)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(156)
			p.Match(NparserMUT)
		}
		{
			p.SetState(157)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(158)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(159)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(160)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(161)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}

		fila := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetColumn()
			}
		}())
		localctx.(*DeclaracionContext).ex = instrucciones.NewDeclaracion((func() string {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_tipo_dato().GetEx(),
			true, false, Ast.VOID, localctx.(*DeclaracionContext).Get_expresion().GetEx(), fila, columna)

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(164)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(165)
			p.Match(NparserMUT)
		}
		{
			p.SetState(166)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(167)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(168)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(169)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(170)

			var _x = p.Control_expresion()

			localctx.(*DeclaracionContext)._control_expresion = _x
		}

		fila := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*DeclaracionContext).Get_LET() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_LET().GetColumn()
			}
		}())
		localctx.(*DeclaracionContext).ex = instrucciones.NewDeclaracion((func() string {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_tipo_dato().GetEx(),
			true, false, Ast.VOID, localctx.(*DeclaracionContext).Get_control_expresion().GetEx(), fila, columna)

	}

	return localctx
}

// IAsignacionContext is an interface to support dynamic dispatch.
type IAsignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Get_control_if_exp returns the _control_if_exp rule contexts.
	Get_control_if_exp() IControl_if_expContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// Set_control_if_exp sets the _control_if_exp rule contexts.
	Set_control_if_exp(IControl_if_expContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	ex              Ast.Instruccion
	_ID             antlr.Token
	_expresion      IExpresionContext
	_control_if_exp IControl_if_expContext
}

func NewEmptyAsignacionContext() *AsignacionContext {
	var p = new(AsignacionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_asignacion
	return p
}

func (*AsignacionContext) IsAsignacionContext() {}

func NewAsignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionContext {
	var p = new(AsignacionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_asignacion

	return p
}

func (s *AsignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionContext) Get_ID() antlr.Token { return s._ID }

func (s *AsignacionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *AsignacionContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *AsignacionContext) Get_control_if_exp() IControl_if_expContext { return s._control_if_exp }

func (s *AsignacionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *AsignacionContext) Set_control_if_exp(v IControl_if_expContext) { s._control_if_exp = v }

func (s *AsignacionContext) GetEx() Ast.Instruccion { return s.ex }

func (s *AsignacionContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *AsignacionContext) ID() antlr.TerminalNode {
	return s.GetToken(NparserID, 0)
}

func (s *AsignacionContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(NparserIGUAL, 0)
}

func (s *AsignacionContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsignacionContext) Control_if_exp() IControl_if_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControl_if_expContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControl_if_expContext)
}

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterAsignacion(s)
	}
}

func (s *AsignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitAsignacion(s)
	}
}

func (p *Nparser) Asignacion() (localctx IAsignacionContext) {
	this := p
	_ = this

	localctx = NewAsignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NparserRULE_asignacion)

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

	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(175)

			var _m = p.Match(NparserID)

			localctx.(*AsignacionContext)._ID = _m
		}
		{
			p.SetState(176)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(177)

			var _x = p.expresion(0)

			localctx.(*AsignacionContext)._expresion = _x
		}

		fila := (func() int {
			if localctx.(*AsignacionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*AsignacionContext).Get_ID().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*AsignacionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*AsignacionContext).Get_ID().GetColumn()
			}
		}())
		localctx.(*AsignacionContext).ex = instrucciones.NewAsignacion((func() string {
			if localctx.(*AsignacionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*AsignacionContext).Get_ID().GetText()
			}
		}()), localctx.(*AsignacionContext).Get_expresion().GetEx(), fila, columna)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(180)

			var _m = p.Match(NparserID)

			localctx.(*AsignacionContext)._ID = _m
		}
		{
			p.SetState(181)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(182)

			var _x = p.Control_if_exp()

			localctx.(*AsignacionContext)._control_if_exp = _x
		}

		fila := (func() int {
			if localctx.(*AsignacionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*AsignacionContext).Get_ID().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*AsignacionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*AsignacionContext).Get_ID().GetColumn()
			}
		}())
		localctx.(*AsignacionContext).ex = instrucciones.NewAsignacion((func() string {
			if localctx.(*AsignacionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*AsignacionContext).Get_ID().GetText()
			}
		}()), localctx.(*AsignacionContext).Get_control_if_exp().GetEx(), fila, columna)

	}

	return localctx
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_CARACTER returns the _CARACTER token.
	Get_CARACTER() antlr.Token

	// Get_DECIMAL returns the _DECIMAL token.
	Get_DECIMAL() antlr.Token

	// Get_NUMERO returns the _NUMERO token.
	Get_NUMERO() antlr.Token

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Get_AND returns the _AND token.
	Get_AND() antlr.Token

	// Get_OR returns the _OR token.
	Get_OR() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_CARACTER sets the _CARACTER token.
	Set_CARACTER(antlr.Token)

	// Set_DECIMAL sets the _DECIMAL token.
	Set_DECIMAL(antlr.Token)

	// Set_NUMERO sets the _NUMERO token.
	Set_NUMERO(antlr.Token)

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// Set_AND sets the _AND token.
	Set_AND(antlr.Token)

	// Set_OR sets the _OR token.
	Set_OR(antlr.Token)

	// GetOp_izq returns the op_izq rule contexts.
	GetOp_izq() IExpresionContext

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// GetOp_der returns the op_der rule contexts.
	GetOp_der() IExpresionContext

	// SetOp_izq sets the op_izq rule contexts.
	SetOp_izq(IExpresionContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// SetOp_der sets the op_der rule contexts.
	SetOp_der(IExpresionContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Expresion

	// SetEx sets the ex attribute.
	SetEx(Ast.Expresion)

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	ex         Ast.Expresion
	op_izq     IExpresionContext
	op         antlr.Token
	_expresion IExpresionContext
	_ID        antlr.Token
	_CARACTER  antlr.Token
	_DECIMAL   antlr.Token
	_NUMERO    antlr.Token
	_CADENA    antlr.Token
	op_der     IExpresionContext
	_AND       antlr.Token
	_OR        antlr.Token
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_expresion
	return p
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) GetOp() antlr.Token { return s.op }

func (s *ExpresionContext) Get_ID() antlr.Token { return s._ID }

func (s *ExpresionContext) Get_CARACTER() antlr.Token { return s._CARACTER }

func (s *ExpresionContext) Get_DECIMAL() antlr.Token { return s._DECIMAL }

func (s *ExpresionContext) Get_NUMERO() antlr.Token { return s._NUMERO }

func (s *ExpresionContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *ExpresionContext) Get_AND() antlr.Token { return s._AND }

func (s *ExpresionContext) Get_OR() antlr.Token { return s._OR }

func (s *ExpresionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpresionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ExpresionContext) Set_CARACTER(v antlr.Token) { s._CARACTER = v }

func (s *ExpresionContext) Set_DECIMAL(v antlr.Token) { s._DECIMAL = v }

func (s *ExpresionContext) Set_NUMERO(v antlr.Token) { s._NUMERO = v }

func (s *ExpresionContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *ExpresionContext) Set_AND(v antlr.Token) { s._AND = v }

func (s *ExpresionContext) Set_OR(v antlr.Token) { s._OR = v }

func (s *ExpresionContext) GetOp_izq() IExpresionContext { return s.op_izq }

func (s *ExpresionContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *ExpresionContext) GetOp_der() IExpresionContext { return s.op_der }

func (s *ExpresionContext) SetOp_izq(v IExpresionContext) { s.op_izq = v }

func (s *ExpresionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *ExpresionContext) SetOp_der(v IExpresionContext) { s.op_der = v }

func (s *ExpresionContext) GetEx() Ast.Expresion { return s.ex }

func (s *ExpresionContext) SetEx(v Ast.Expresion) { s.ex = v }

func (s *ExpresionContext) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *ExpresionContext) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpresionContext) RESTA() antlr.TerminalNode {
	return s.GetToken(NparserRESTA, 0)
}

func (s *ExpresionContext) NOT() antlr.TerminalNode {
	return s.GetToken(NparserNOT, 0)
}

func (s *ExpresionContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(NparserPAR_IZQ, 0)
}

func (s *ExpresionContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(NparserPAR_DER, 0)
}

func (s *ExpresionContext) ID() antlr.TerminalNode {
	return s.GetToken(NparserID, 0)
}

func (s *ExpresionContext) TRUE() antlr.TerminalNode {
	return s.GetToken(NparserTRUE, 0)
}

func (s *ExpresionContext) FALSE() antlr.TerminalNode {
	return s.GetToken(NparserFALSE, 0)
}

func (s *ExpresionContext) CARACTER() antlr.TerminalNode {
	return s.GetToken(NparserCARACTER, 0)
}

func (s *ExpresionContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(NparserDECIMAL, 0)
}

func (s *ExpresionContext) NUMERO() antlr.TerminalNode {
	return s.GetToken(NparserNUMERO, 0)
}

func (s *ExpresionContext) CADENA() antlr.TerminalNode {
	return s.GetToken(NparserCADENA, 0)
}

func (s *ExpresionContext) MULTIPLICACION() antlr.TerminalNode {
	return s.GetToken(NparserMULTIPLICACION, 0)
}

func (s *ExpresionContext) DIVISION() antlr.TerminalNode {
	return s.GetToken(NparserDIVISION, 0)
}

func (s *ExpresionContext) MODULO() antlr.TerminalNode {
	return s.GetToken(NparserMODULO, 0)
}

func (s *ExpresionContext) SUMA() antlr.TerminalNode {
	return s.GetToken(NparserSUMA, 0)
}

func (s *ExpresionContext) MAYOR_I() antlr.TerminalNode {
	return s.GetToken(NparserMAYOR_I, 0)
}

func (s *ExpresionContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(NparserMAYOR, 0)
}

func (s *ExpresionContext) MENOR_I() antlr.TerminalNode {
	return s.GetToken(NparserMENOR_I, 0)
}

func (s *ExpresionContext) MENOR() antlr.TerminalNode {
	return s.GetToken(NparserMENOR, 0)
}

func (s *ExpresionContext) IGUALDAD() antlr.TerminalNode {
	return s.GetToken(NparserIGUALDAD, 0)
}

func (s *ExpresionContext) DISTINTO() antlr.TerminalNode {
	return s.GetToken(NparserDISTINTO, 0)
}

func (s *ExpresionContext) AND() antlr.TerminalNode {
	return s.GetToken(NparserAND, 0)
}

func (s *ExpresionContext) OR() antlr.TerminalNode {
	return s.GetToken(NparserOR, 0)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterExpresion(s)
	}
}

func (s *ExpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitExpresion(s)
	}
}

func (p *Nparser) Expresion() (localctx IExpresionContext) {
	return p.expresion(0)
}

func (p *Nparser) expresion(_p int) (localctx IExpresionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, NparserRULE_expresion, _p)
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
	p.SetState(211)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserRESTA, NparserNOT:
		{
			p.SetState(188)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpresionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == NparserRESTA || _la == NparserNOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpresionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(189)

			var _x = p.expresion(15)

			localctx.(*ExpresionContext).op_izq = _x
			localctx.(*ExpresionContext)._expresion = _x
		}

		fila := (func() int {
			if localctx.(*ExpresionContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).GetOp().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*ExpresionContext).GetOp() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).GetOp().GetColumn()
			}
		}())
		localctx.(*ExpresionContext).ex = expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_izq().GetEx(), (func() string {
			if localctx.(*ExpresionContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).GetOp().GetText()
			}
		}()), nil, true, fila, columna)

	case NparserPAR_IZQ:
		{
			p.SetState(192)
			p.Match(NparserPAR_IZQ)
		}
		{
			p.SetState(193)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext)._expresion = _x
		}
		{
			p.SetState(194)
			p.Match(NparserPAR_DER)
		}

		localctx.(*ExpresionContext).ex = localctx.(*ExpresionContext).Get_expresion().GetEx()

	case NparserID:
		{
			p.SetState(197)

			var _m = p.Match(NparserID)

			localctx.(*ExpresionContext)._ID = _m
		}

		id := (func() string {
			if localctx.(*ExpresionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_ID().GetText()
			}
		}())
		fila := (func() int {
			if localctx.(*ExpresionContext).Get_DECIMAL() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_DECIMAL().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*ExpresionContext).Get_DECIMAL() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_DECIMAL().GetColumn()
			}
		}())
		localctx.(*ExpresionContext).ex = expresiones.NewIdentificador(id, Ast.IDENTIFICADOR, fila, columna)

	case NparserTRUE:
		{
			p.SetState(199)
			p.Match(NparserTRUE)
		}

		valor := true
		fila := (func() int {
			if localctx.(*ExpresionContext).Get_DECIMAL() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_DECIMAL().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*ExpresionContext).Get_DECIMAL() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_DECIMAL().GetColumn()
			}
		}())
		localctx.(*ExpresionContext).ex = expresiones.NewPrimitivo(valor, Ast.BOOLEAN, fila, columna)

	case NparserFALSE:
		{
			p.SetState(201)
			p.Match(NparserFALSE)
		}

		valor := false
		fila := (func() int {
			if localctx.(*ExpresionContext).Get_DECIMAL() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_DECIMAL().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*ExpresionContext).Get_DECIMAL() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_DECIMAL().GetColumn()
			}
		}())
		localctx.(*ExpresionContext).ex = expresiones.NewPrimitivo(valor, Ast.BOOLEAN, fila, columna)

	case NparserCARACTER:
		{
			p.SetState(203)

			var _m = p.Match(NparserCARACTER)

			localctx.(*ExpresionContext)._CARACTER = _m
		}

		valor := (func() string {
			if localctx.(*ExpresionContext).Get_CARACTER() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_CARACTER().GetText()
			}
		}())
		valor = valor[1 : len(valor)-1]
		fila := (func() int {
			if localctx.(*ExpresionContext).Get_DECIMAL() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_DECIMAL().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*ExpresionContext).Get_DECIMAL() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_DECIMAL().GetColumn()
			}
		}())
		localctx.(*ExpresionContext).ex = expresiones.NewPrimitivo(valor, Ast.CHAR, fila, columna)

	case NparserDECIMAL:
		{
			p.SetState(205)

			var _m = p.Match(NparserDECIMAL)

			localctx.(*ExpresionContext)._DECIMAL = _m
		}

		valor, err := strconv.ParseFloat((func() string {
			if localctx.(*ExpresionContext).Get_DECIMAL() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_DECIMAL().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println("Hay un error en el get número")
			fmt.Println(err)
		}
		fila := (func() int {
			if localctx.(*ExpresionContext).Get_DECIMAL() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_DECIMAL().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*ExpresionContext).Get_DECIMAL() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_DECIMAL().GetColumn()
			}
		}())
		localctx.(*ExpresionContext).ex = expresiones.NewPrimitivo(valor, Ast.F64, fila, columna)

	case NparserNUMERO:
		{
			p.SetState(207)

			var _m = p.Match(NparserNUMERO)

			localctx.(*ExpresionContext)._NUMERO = _m
		}

		valor, err := strconv.Atoi((func() string {
			if localctx.(*ExpresionContext).Get_NUMERO() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_NUMERO().GetText()
			}
		}()))
		if err != nil {
			fmt.Println("Hay un error en el get número")
			fmt.Println(err)
		}
		fila := (func() int {
			if localctx.(*ExpresionContext).Get_NUMERO() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_NUMERO().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*ExpresionContext).Get_NUMERO() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_NUMERO().GetColumn()
			}
		}())
		localctx.(*ExpresionContext).ex = expresiones.NewPrimitivo(valor, Ast.I64, fila, columna)

	case NparserCADENA:
		{
			p.SetState(209)

			var _m = p.Match(NparserCADENA)

			localctx.(*ExpresionContext)._CADENA = _m
		}

		fila := (func() int {
			if localctx.(*ExpresionContext).Get_CADENA() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_CADENA().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*ExpresionContext).Get_CADENA() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_CADENA().GetColumn()
			}
		}())
		valor := (func() string {
			if localctx.(*ExpresionContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_CADENA().GetText()
			}
		}())
		valor = valor[1 : len(valor)-1]
		localctx.(*ExpresionContext).ex = expresiones.NewPrimitivo(valor, Ast.STR, fila, columna)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(243)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).op_izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(213)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(214)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-70)&-(0x1f+1)) == 0 && ((1<<uint((_la-70)))&((1<<(NparserMODULO-70))|(1<<(NparserMULTIPLICACION-70))|(1<<(NparserDIVISION-70)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(215)

					var _x = p.expresion(15)

					localctx.(*ExpresionContext).op_der = _x
					localctx.(*ExpresionContext)._expresion = _x
				}

				fila := (func() int {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOp().GetLine()
					}
				}())
				columna := (func() int {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOp().GetColumn()
					}
				}())
				localctx.(*ExpresionContext).ex = expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_izq().GetEx(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetOp_der().GetEx(), false, fila, columna)

			case 2:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).op_izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(218)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(219)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == NparserRESTA || _la == NparserSUMA) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(220)

					var _x = p.expresion(14)

					localctx.(*ExpresionContext).op_der = _x
					localctx.(*ExpresionContext)._expresion = _x
				}

				fila := (func() int {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOp().GetLine()
					}
				}())
				columna := (func() int {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOp().GetColumn()
					}
				}())
				localctx.(*ExpresionContext).ex = expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_izq().GetEx(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetOp_der().GetEx(), false, fila, columna)

			case 3:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).op_izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(223)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(224)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-62)&-(0x1f+1)) == 0 && ((1<<uint((_la-62)))&((1<<(NparserMAYOR_I-62))|(1<<(NparserMAYOR-62))|(1<<(NparserMENOR_I-62))|(1<<(NparserMENOR-62))|(1<<(NparserIGUALDAD-62))|(1<<(NparserDISTINTO-62)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(225)

					var _x = p.expresion(13)

					localctx.(*ExpresionContext).op_der = _x
					localctx.(*ExpresionContext)._expresion = _x
				}

				fila := (func() int {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOp().GetLine()
					}
				}())
				columna := (func() int {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOp().GetColumn()
					}
				}())
				localctx.(*ExpresionContext).ex = expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_izq().GetEx(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetOp_der().GetEx(), false, fila, columna)

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).op_izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(228)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(229)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == NparserIGUALDAD || _la == NparserDISTINTO) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(230)

					var _x = p.expresion(12)

					localctx.(*ExpresionContext).op_der = _x
					localctx.(*ExpresionContext)._expresion = _x
				}

				fila := (func() int {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOp().GetLine()
					}
				}())
				columna := (func() int {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOp().GetColumn()
					}
				}())
				localctx.(*ExpresionContext).ex = expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_izq().GetEx(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetOp_der().GetEx(), false, fila, columna)

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).op_izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(233)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(234)

					var _m = p.Match(NparserAND)

					localctx.(*ExpresionContext)._AND = _m
				}
				{
					p.SetState(235)

					var _x = p.expresion(11)

					localctx.(*ExpresionContext).op_der = _x
					localctx.(*ExpresionContext)._expresion = _x
				}

				fila := (func() int {
					if localctx.(*ExpresionContext).Get_AND() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_AND().GetLine()
					}
				}())
				columna := (func() int {
					if localctx.(*ExpresionContext).Get_AND() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_AND().GetColumn()
					}
				}())
				localctx.(*ExpresionContext).ex = expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_izq().GetEx(), (func() string {
					if localctx.(*ExpresionContext).Get_AND() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).Get_AND().GetText()
					}
				}()), localctx.(*ExpresionContext).GetOp_der().GetEx(), false, fila, columna)

			case 6:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).op_izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(238)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(239)

					var _m = p.Match(NparserOR)

					localctx.(*ExpresionContext)._OR = _m
				}
				{
					p.SetState(240)

					var _x = p.expresion(10)

					localctx.(*ExpresionContext).op_der = _x
					localctx.(*ExpresionContext)._expresion = _x
				}

				fila := (func() int {
					if localctx.(*ExpresionContext).Get_OR() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_OR().GetLine()
					}
				}())
				columna := (func() int {
					if localctx.(*ExpresionContext).Get_OR() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_OR().GetColumn()
					}
				}())
				localctx.(*ExpresionContext).ex = expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_izq().GetEx(), (func() string {
					if localctx.(*ExpresionContext).Get_OR() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).Get_OR().GetText()
					}
				}()), localctx.(*ExpresionContext).GetOp_der().GetEx(), false, fila, columna)

			}

		}
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// ITipo_datoContext is an interface to support dynamic dispatch.
type ITipo_datoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEx returns the ex attribute.
	GetEx() Ast.TipoDato

	// SetEx sets the ex attribute.
	SetEx(Ast.TipoDato)

	// IsTipo_datoContext differentiates from other interfaces.
	IsTipo_datoContext()
}

type Tipo_datoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ex     Ast.TipoDato
}

func NewEmptyTipo_datoContext() *Tipo_datoContext {
	var p = new(Tipo_datoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_tipo_dato
	return p
}

func (*Tipo_datoContext) IsTipo_datoContext() {}

func NewTipo_datoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tipo_datoContext {
	var p = new(Tipo_datoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_tipo_dato

	return p
}

func (s *Tipo_datoContext) GetParser() antlr.Parser { return s.parser }

func (s *Tipo_datoContext) GetEx() Ast.TipoDato { return s.ex }

func (s *Tipo_datoContext) SetEx(v Ast.TipoDato) { s.ex = v }

func (s *Tipo_datoContext) BOOL() antlr.TerminalNode {
	return s.GetToken(NparserBOOL, 0)
}

func (s *Tipo_datoContext) CHAR() antlr.TerminalNode {
	return s.GetToken(NparserCHAR, 0)
}

func (s *Tipo_datoContext) I64() antlr.TerminalNode {
	return s.GetToken(NparserI64, 0)
}

func (s *Tipo_datoContext) F64() antlr.TerminalNode {
	return s.GetToken(NparserF64, 0)
}

func (s *Tipo_datoContext) STR() antlr.TerminalNode {
	return s.GetToken(NparserSTR, 0)
}

func (s *Tipo_datoContext) STRING() antlr.TerminalNode {
	return s.GetToken(NparserSTRING, 0)
}

func (s *Tipo_datoContext) USIZE() antlr.TerminalNode {
	return s.GetToken(NparserUSIZE, 0)
}

func (s *Tipo_datoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_datoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tipo_datoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterTipo_dato(s)
	}
}

func (s *Tipo_datoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitTipo_dato(s)
	}
}

func (p *Nparser) Tipo_dato() (localctx ITipo_datoContext) {
	this := p
	_ = this

	localctx = NewTipo_datoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NparserRULE_tipo_dato)

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

	p.SetState(262)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserBOOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(248)
			p.Match(NparserBOOL)
		}
		localctx.(*Tipo_datoContext).ex = Ast.BOOLEAN

	case NparserCHAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(250)
			p.Match(NparserCHAR)
		}
		localctx.(*Tipo_datoContext).ex = Ast.CHAR

	case NparserI64:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(252)
			p.Match(NparserI64)
		}
		localctx.(*Tipo_datoContext).ex = Ast.I64

	case NparserF64:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(254)
			p.Match(NparserF64)
		}
		localctx.(*Tipo_datoContext).ex = Ast.F64

	case NparserSTR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(256)
			p.Match(NparserSTR)
		}
		localctx.(*Tipo_datoContext).ex = Ast.STR

	case NparserSTRING:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(258)
			p.Match(NparserSTRING)
		}
		localctx.(*Tipo_datoContext).ex = Ast.STRING

	case NparserUSIZE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(260)
			p.Match(NparserUSIZE)
		}
		localctx.(*Tipo_datoContext).ex = Ast.USIZE

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IControl_ifContext is an interface to support dynamic dispatch.
type IControl_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IF returns the _IF token.
	Get_IF() antlr.Token

	// Set_IF sets the _IF token.
	Set_IF(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// GetBloqueIf returns the bloqueIf rule contexts.
	GetBloqueIf() IBloqueContext

	// GetBloqueElse returns the bloqueElse rule contexts.
	GetBloqueElse() IBloqueContext

	// Get_bloque_else_if returns the _bloque_else_if rule contexts.
	Get_bloque_else_if() IBloque_else_ifContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// SetBloqueIf sets the bloqueIf rule contexts.
	SetBloqueIf(IBloqueContext)

	// SetBloqueElse sets the bloqueElse rule contexts.
	SetBloqueElse(IBloqueContext)

	// Set_bloque_else_if sets the _bloque_else_if rule contexts.
	Set_bloque_else_if(IBloque_else_ifContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsControl_ifContext differentiates from other interfaces.
	IsControl_ifContext()
}

type Control_ifContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	ex              Ast.Instruccion
	_IF             antlr.Token
	_expresion      IExpresionContext
	bloqueIf        IBloqueContext
	bloqueElse      IBloqueContext
	_bloque_else_if IBloque_else_ifContext
}

func NewEmptyControl_ifContext() *Control_ifContext {
	var p = new(Control_ifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_control_if
	return p
}

func (*Control_ifContext) IsControl_ifContext() {}

func NewControl_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Control_ifContext {
	var p = new(Control_ifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_control_if

	return p
}

func (s *Control_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *Control_ifContext) Get_IF() antlr.Token { return s._IF }

func (s *Control_ifContext) Set_IF(v antlr.Token) { s._IF = v }

func (s *Control_ifContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *Control_ifContext) GetBloqueIf() IBloqueContext { return s.bloqueIf }

func (s *Control_ifContext) GetBloqueElse() IBloqueContext { return s.bloqueElse }

func (s *Control_ifContext) Get_bloque_else_if() IBloque_else_ifContext { return s._bloque_else_if }

func (s *Control_ifContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *Control_ifContext) SetBloqueIf(v IBloqueContext) { s.bloqueIf = v }

func (s *Control_ifContext) SetBloqueElse(v IBloqueContext) { s.bloqueElse = v }

func (s *Control_ifContext) Set_bloque_else_if(v IBloque_else_ifContext) { s._bloque_else_if = v }

func (s *Control_ifContext) GetEx() Ast.Instruccion { return s.ex }

func (s *Control_ifContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *Control_ifContext) IF() antlr.TerminalNode {
	return s.GetToken(NparserIF, 0)
}

func (s *Control_ifContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Control_ifContext) AllBloque() []IBloqueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBloqueContext)(nil)).Elem())
	var tst = make([]IBloqueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBloqueContext)
		}
	}

	return tst
}

func (s *Control_ifContext) Bloque(i int) IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Control_ifContext) ELSE() antlr.TerminalNode {
	return s.GetToken(NparserELSE, 0)
}

func (s *Control_ifContext) Bloque_else_if() IBloque_else_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloque_else_ifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloque_else_ifContext)
}

func (s *Control_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Control_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Control_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterControl_if(s)
	}
}

func (s *Control_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitControl_if(s)
	}
}

func (p *Nparser) Control_if() (localctx IControl_ifContext) {
	this := p
	_ = this

	localctx = NewControl_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NparserRULE_control_if)

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

	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(264)

			var _m = p.Match(NparserIF)

			localctx.(*Control_ifContext)._IF = _m
		}
		{
			p.SetState(265)

			var _x = p.expresion(0)

			localctx.(*Control_ifContext)._expresion = _x
		}
		{
			p.SetState(266)

			var _x = p.Bloque()

			localctx.(*Control_ifContext).bloqueIf = _x
		}

		fila := (func() int {
			if localctx.(*Control_ifContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Control_ifContext).Get_IF().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Control_ifContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Control_ifContext).Get_IF().GetColumn()
			}
		}())
		columna++
		lista_null := arraylist.New()
		localctx.(*Control_ifContext).ex = exp_ins.NewIF(localctx.(*Control_ifContext).Get_expresion().GetEx(), localctx.(*Control_ifContext).GetBloqueIf().GetList(), lista_null, Ast.IF, fila, columna, false)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(269)

			var _m = p.Match(NparserIF)

			localctx.(*Control_ifContext)._IF = _m
		}
		{
			p.SetState(270)

			var _x = p.expresion(0)

			localctx.(*Control_ifContext)._expresion = _x
		}
		{
			p.SetState(271)

			var _x = p.Bloque()

			localctx.(*Control_ifContext).bloqueIf = _x
		}
		{
			p.SetState(272)
			p.Match(NparserELSE)
		}
		{
			p.SetState(273)

			var _x = p.Bloque()

			localctx.(*Control_ifContext).bloqueElse = _x
		}

		fila := (func() int {
			if localctx.(*Control_ifContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Control_ifContext).Get_IF().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Control_ifContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Control_ifContext).Get_IF().GetColumn()
			}
		}())
		columna++
		lista_entonces := arraylist.New()
		lista_null := arraylist.New()
		Else := exp_ins.NewIF(localctx.(*Control_ifContext).Get_expresion().GetEx(), localctx.(*Control_ifContext).GetBloqueElse().GetList(), lista_null, Ast.ELSE, fila, columna, false)
		lista_entonces.Add(Else)
		localctx.(*Control_ifContext).ex = exp_ins.NewIF(localctx.(*Control_ifContext).Get_expresion().GetEx(), localctx.(*Control_ifContext).GetBloqueIf().GetList(), lista_entonces, Ast.IF, fila, columna, false)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(276)

			var _m = p.Match(NparserIF)

			localctx.(*Control_ifContext)._IF = _m
		}
		{
			p.SetState(277)

			var _x = p.expresion(0)

			localctx.(*Control_ifContext)._expresion = _x
		}
		{
			p.SetState(278)

			var _x = p.Bloque()

			localctx.(*Control_ifContext).bloqueIf = _x
		}
		{
			p.SetState(279)

			var _x = p.Bloque_else_if()

			localctx.(*Control_ifContext)._bloque_else_if = _x
		}

		fila := (func() int {
			if localctx.(*Control_ifContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Control_ifContext).Get_IF().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Control_ifContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Control_ifContext).Get_IF().GetColumn()
			}
		}())
		columna++
		lista_entonces := localctx.(*Control_ifContext).Get_bloque_else_if().GetList()
		localctx.(*Control_ifContext).ex = exp_ins.NewIF(localctx.(*Control_ifContext).Get_expresion().GetEx(), localctx.(*Control_ifContext).GetBloqueIf().GetList(), lista_entonces, Ast.IF, fila, columna, false)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(282)

			var _m = p.Match(NparserIF)

			localctx.(*Control_ifContext)._IF = _m
		}
		{
			p.SetState(283)

			var _x = p.expresion(0)

			localctx.(*Control_ifContext)._expresion = _x
		}
		{
			p.SetState(284)

			var _x = p.Bloque()

			localctx.(*Control_ifContext).bloqueIf = _x
		}
		{
			p.SetState(285)

			var _x = p.Bloque_else_if()

			localctx.(*Control_ifContext)._bloque_else_if = _x
		}
		{
			p.SetState(286)
			p.Match(NparserELSE)
		}
		{
			p.SetState(287)

			var _x = p.Bloque()

			localctx.(*Control_ifContext).bloqueElse = _x
		}

		fila := (func() int {
			if localctx.(*Control_ifContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Control_ifContext).Get_IF().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Control_ifContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Control_ifContext).Get_IF().GetColumn()
			}
		}())
		columna++
		lista_null := arraylist.New()
		Else := exp_ins.NewIF(localctx.(*Control_ifContext).Get_expresion().GetEx(), localctx.(*Control_ifContext).GetBloqueElse().GetList(), lista_null, Ast.ELSE, fila, columna, false)
		lista_entonces := localctx.(*Control_ifContext).Get_bloque_else_if().GetList()
		lista_entonces.Add(Else)
		localctx.(*Control_ifContext).ex = exp_ins.NewIF(localctx.(*Control_ifContext).Get_expresion().GetEx(), localctx.(*Control_ifContext).GetBloqueIf().GetList(), lista_entonces, Ast.IF, fila, columna, false)

	}

	return localctx
}

// IBloque_else_ifContext is an interface to support dynamic dispatch.
type IBloque_else_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_else_if returns the _else_if rule contexts.
	Get_else_if() IElse_ifContext

	// Set_else_if sets the _else_if rule contexts.
	Set_else_if(IElse_ifContext)

	// GetLista returns the lista rule context list.
	GetLista() []IElse_ifContext

	// SetLista sets the lista rule context list.
	SetLista([]IElse_ifContext)

	// GetList returns the list attribute.
	GetList() *arraylist.List

	// SetList sets the list attribute.
	SetList(*arraylist.List)

	// IsBloque_else_ifContext differentiates from other interfaces.
	IsBloque_else_ifContext()
}

type Bloque_else_ifContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	list     *arraylist.List
	_else_if IElse_ifContext
	lista    []IElse_ifContext
}

func NewEmptyBloque_else_ifContext() *Bloque_else_ifContext {
	var p = new(Bloque_else_ifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_bloque_else_if
	return p
}

func (*Bloque_else_ifContext) IsBloque_else_ifContext() {}

func NewBloque_else_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bloque_else_ifContext {
	var p = new(Bloque_else_ifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_bloque_else_if

	return p
}

func (s *Bloque_else_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *Bloque_else_ifContext) Get_else_if() IElse_ifContext { return s._else_if }

func (s *Bloque_else_ifContext) Set_else_if(v IElse_ifContext) { s._else_if = v }

func (s *Bloque_else_ifContext) GetLista() []IElse_ifContext { return s.lista }

func (s *Bloque_else_ifContext) SetLista(v []IElse_ifContext) { s.lista = v }

func (s *Bloque_else_ifContext) GetList() *arraylist.List { return s.list }

func (s *Bloque_else_ifContext) SetList(v *arraylist.List) { s.list = v }

func (s *Bloque_else_ifContext) AllElse_if() []IElse_ifContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElse_ifContext)(nil)).Elem())
	var tst = make([]IElse_ifContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElse_ifContext)
		}
	}

	return tst
}

func (s *Bloque_else_ifContext) Else_if(i int) IElse_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElse_ifContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElse_ifContext)
}

func (s *Bloque_else_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bloque_else_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bloque_else_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterBloque_else_if(s)
	}
}

func (s *Bloque_else_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitBloque_else_if(s)
	}
}

func (p *Nparser) Bloque_else_if() (localctx IBloque_else_ifContext) {
	this := p
	_ = this

	localctx = NewBloque_else_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, NparserRULE_bloque_else_if)
	localctx.(*Bloque_else_ifContext).list = arraylist.New()

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
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(292)

				var _x = p.Else_if()

				localctx.(*Bloque_else_ifContext)._else_if = _x
			}
			localctx.(*Bloque_else_ifContext).lista = append(localctx.(*Bloque_else_ifContext).lista, localctx.(*Bloque_else_ifContext)._else_if)

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	listas := localctx.(*Bloque_else_ifContext).GetLista()
	for _, e := range listas {
		localctx.(*Bloque_else_ifContext).list.Add(e.GetEx())
	}

	return localctx
}

// IElse_ifContext is an interface to support dynamic dispatch.
type IElse_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ELSE returns the _ELSE token.
	Get_ELSE() antlr.Token

	// Set_ELSE sets the _ELSE token.
	Set_ELSE(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Get_bloque returns the _bloque rule contexts.
	Get_bloque() IBloqueContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// Set_bloque sets the _bloque rule contexts.
	Set_bloque(IBloqueContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsElse_ifContext differentiates from other interfaces.
	IsElse_ifContext()
}

type Else_ifContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	ex         Ast.Instruccion
	_ELSE      antlr.Token
	_expresion IExpresionContext
	_bloque    IBloqueContext
}

func NewEmptyElse_ifContext() *Else_ifContext {
	var p = new(Else_ifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_else_if
	return p
}

func (*Else_ifContext) IsElse_ifContext() {}

func NewElse_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_ifContext {
	var p = new(Else_ifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_else_if

	return p
}

func (s *Else_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_ifContext) Get_ELSE() antlr.Token { return s._ELSE }

func (s *Else_ifContext) Set_ELSE(v antlr.Token) { s._ELSE = v }

func (s *Else_ifContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *Else_ifContext) Get_bloque() IBloqueContext { return s._bloque }

func (s *Else_ifContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *Else_ifContext) Set_bloque(v IBloqueContext) { s._bloque = v }

func (s *Else_ifContext) GetEx() Ast.Instruccion { return s.ex }

func (s *Else_ifContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *Else_ifContext) ELSE() antlr.TerminalNode {
	return s.GetToken(NparserELSE, 0)
}

func (s *Else_ifContext) IF() antlr.TerminalNode {
	return s.GetToken(NparserIF, 0)
}

func (s *Else_ifContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Else_ifContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Else_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Else_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterElse_if(s)
	}
}

func (s *Else_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitElse_if(s)
	}
}

func (p *Nparser) Else_if() (localctx IElse_ifContext) {
	this := p
	_ = this

	localctx = NewElse_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, NparserRULE_else_if)

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
		p.SetState(299)

		var _m = p.Match(NparserELSE)

		localctx.(*Else_ifContext)._ELSE = _m
	}
	{
		p.SetState(300)
		p.Match(NparserIF)
	}
	{
		p.SetState(301)

		var _x = p.expresion(0)

		localctx.(*Else_ifContext)._expresion = _x
	}
	{
		p.SetState(302)

		var _x = p.Bloque()

		localctx.(*Else_ifContext)._bloque = _x
	}

	fila := (func() int {
		if localctx.(*Else_ifContext).Get_ELSE() == nil {
			return 0
		} else {
			return localctx.(*Else_ifContext).Get_ELSE().GetLine()
		}
	}())
	columna := (func() int {
		if localctx.(*Else_ifContext).Get_ELSE() == nil {
			return 0
		} else {
			return localctx.(*Else_ifContext).Get_ELSE().GetColumn()
		}
	}())
	columna++
	lista_null := arraylist.New()
	localctx.(*Else_ifContext).ex = exp_ins.NewIF(localctx.(*Else_ifContext).Get_expresion().GetEx(), localctx.(*Else_ifContext).Get_bloque().GetList(), lista_null, Ast.ELSEIF, fila, columna, false)

	return localctx
}

// IControl_if_expContext is an interface to support dynamic dispatch.
type IControl_if_expContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IF returns the _IF token.
	Get_IF() antlr.Token

	// Set_IF sets the _IF token.
	Set_IF(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// GetBloqueIf returns the bloqueIf rule contexts.
	GetBloqueIf() IBloqueContext

	// GetBloqueElse returns the bloqueElse rule contexts.
	GetBloqueElse() IBloqueContext

	// Get_bloque_else_if_exp returns the _bloque_else_if_exp rule contexts.
	Get_bloque_else_if_exp() IBloque_else_if_expContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// SetBloqueIf sets the bloqueIf rule contexts.
	SetBloqueIf(IBloqueContext)

	// SetBloqueElse sets the bloqueElse rule contexts.
	SetBloqueElse(IBloqueContext)

	// Set_bloque_else_if_exp sets the _bloque_else_if_exp rule contexts.
	Set_bloque_else_if_exp(IBloque_else_if_expContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsControl_if_expContext differentiates from other interfaces.
	IsControl_if_expContext()
}

type Control_if_expContext struct {
	*antlr.BaseParserRuleContext
	parser              antlr.Parser
	ex                  Ast.Instruccion
	_IF                 antlr.Token
	_expresion          IExpresionContext
	bloqueIf            IBloqueContext
	bloqueElse          IBloqueContext
	_bloque_else_if_exp IBloque_else_if_expContext
}

func NewEmptyControl_if_expContext() *Control_if_expContext {
	var p = new(Control_if_expContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_control_if_exp
	return p
}

func (*Control_if_expContext) IsControl_if_expContext() {}

func NewControl_if_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Control_if_expContext {
	var p = new(Control_if_expContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_control_if_exp

	return p
}

func (s *Control_if_expContext) GetParser() antlr.Parser { return s.parser }

func (s *Control_if_expContext) Get_IF() antlr.Token { return s._IF }

func (s *Control_if_expContext) Set_IF(v antlr.Token) { s._IF = v }

func (s *Control_if_expContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *Control_if_expContext) GetBloqueIf() IBloqueContext { return s.bloqueIf }

func (s *Control_if_expContext) GetBloqueElse() IBloqueContext { return s.bloqueElse }

func (s *Control_if_expContext) Get_bloque_else_if_exp() IBloque_else_if_expContext {
	return s._bloque_else_if_exp
}

func (s *Control_if_expContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *Control_if_expContext) SetBloqueIf(v IBloqueContext) { s.bloqueIf = v }

func (s *Control_if_expContext) SetBloqueElse(v IBloqueContext) { s.bloqueElse = v }

func (s *Control_if_expContext) Set_bloque_else_if_exp(v IBloque_else_if_expContext) {
	s._bloque_else_if_exp = v
}

func (s *Control_if_expContext) GetEx() Ast.Instruccion { return s.ex }

func (s *Control_if_expContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *Control_if_expContext) IF() antlr.TerminalNode {
	return s.GetToken(NparserIF, 0)
}

func (s *Control_if_expContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Control_if_expContext) AllBloque() []IBloqueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBloqueContext)(nil)).Elem())
	var tst = make([]IBloqueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBloqueContext)
		}
	}

	return tst
}

func (s *Control_if_expContext) Bloque(i int) IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Control_if_expContext) ELSE() antlr.TerminalNode {
	return s.GetToken(NparserELSE, 0)
}

func (s *Control_if_expContext) Bloque_else_if_exp() IBloque_else_if_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloque_else_if_expContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloque_else_if_expContext)
}

func (s *Control_if_expContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Control_if_expContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Control_if_expContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterControl_if_exp(s)
	}
}

func (s *Control_if_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitControl_if_exp(s)
	}
}

func (p *Nparser) Control_if_exp() (localctx IControl_if_expContext) {
	this := p
	_ = this

	localctx = NewControl_if_expContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, NparserRULE_control_if_exp)

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

	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(305)

			var _m = p.Match(NparserIF)

			localctx.(*Control_if_expContext)._IF = _m
		}
		{
			p.SetState(306)

			var _x = p.expresion(0)

			localctx.(*Control_if_expContext)._expresion = _x
		}
		{
			p.SetState(307)

			var _x = p.Bloque()

			localctx.(*Control_if_expContext).bloqueIf = _x
		}

		fila := (func() int {
			if localctx.(*Control_if_expContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Control_if_expContext).Get_IF().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Control_if_expContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Control_if_expContext).Get_IF().GetColumn()
			}
		}())
		columna++
		lista_null := arraylist.New()
		localctx.(*Control_if_expContext).ex = exp_ins.NewIF(localctx.(*Control_if_expContext).Get_expresion().GetEx(), localctx.(*Control_if_expContext).GetBloqueIf().GetList(), lista_null, Ast.IF_EXPRESION, fila, columna, true)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(310)

			var _m = p.Match(NparserIF)

			localctx.(*Control_if_expContext)._IF = _m
		}
		{
			p.SetState(311)

			var _x = p.expresion(0)

			localctx.(*Control_if_expContext)._expresion = _x
		}
		{
			p.SetState(312)

			var _x = p.Bloque()

			localctx.(*Control_if_expContext).bloqueIf = _x
		}
		{
			p.SetState(313)
			p.Match(NparserELSE)
		}
		{
			p.SetState(314)

			var _x = p.Bloque()

			localctx.(*Control_if_expContext).bloqueElse = _x
		}

		fila := (func() int {
			if localctx.(*Control_if_expContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Control_if_expContext).Get_IF().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Control_if_expContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Control_if_expContext).Get_IF().GetColumn()
			}
		}())
		columna++
		lista_entonces := arraylist.New()
		lista_null := arraylist.New()
		Else := exp_ins.NewIF(localctx.(*Control_if_expContext).Get_expresion().GetEx(), localctx.(*Control_if_expContext).GetBloqueElse().GetList(), lista_null, Ast.ELSE_EXPRESION, fila, columna, true)
		lista_entonces.Add(Else)
		localctx.(*Control_if_expContext).ex = exp_ins.NewIF(localctx.(*Control_if_expContext).Get_expresion().GetEx(), localctx.(*Control_if_expContext).GetBloqueIf().GetList(), lista_entonces, Ast.IF_EXPRESION, fila, columna, true)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(317)

			var _m = p.Match(NparserIF)

			localctx.(*Control_if_expContext)._IF = _m
		}
		{
			p.SetState(318)

			var _x = p.expresion(0)

			localctx.(*Control_if_expContext)._expresion = _x
		}
		{
			p.SetState(319)

			var _x = p.Bloque()

			localctx.(*Control_if_expContext).bloqueIf = _x
		}
		{
			p.SetState(320)

			var _x = p.Bloque_else_if_exp()

			localctx.(*Control_if_expContext)._bloque_else_if_exp = _x
		}

		fila := (func() int {
			if localctx.(*Control_if_expContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Control_if_expContext).Get_IF().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Control_if_expContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Control_if_expContext).Get_IF().GetColumn()
			}
		}())
		columna++
		lista_entonces := localctx.(*Control_if_expContext).Get_bloque_else_if_exp().GetList()
		localctx.(*Control_if_expContext).ex = exp_ins.NewIF(localctx.(*Control_if_expContext).Get_expresion().GetEx(), localctx.(*Control_if_expContext).GetBloqueIf().GetList(), lista_entonces, Ast.IF_EXPRESION, fila, columna, true)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(323)

			var _m = p.Match(NparserIF)

			localctx.(*Control_if_expContext)._IF = _m
		}
		{
			p.SetState(324)

			var _x = p.expresion(0)

			localctx.(*Control_if_expContext)._expresion = _x
		}
		{
			p.SetState(325)

			var _x = p.Bloque()

			localctx.(*Control_if_expContext).bloqueIf = _x
		}
		{
			p.SetState(326)

			var _x = p.Bloque_else_if_exp()

			localctx.(*Control_if_expContext)._bloque_else_if_exp = _x
		}
		{
			p.SetState(327)
			p.Match(NparserELSE)
		}
		{
			p.SetState(328)

			var _x = p.Bloque()

			localctx.(*Control_if_expContext).bloqueElse = _x
		}

		fila := (func() int {
			if localctx.(*Control_if_expContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Control_if_expContext).Get_IF().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Control_if_expContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Control_if_expContext).Get_IF().GetColumn()
			}
		}())
		columna++
		lista_null := arraylist.New()
		Else := exp_ins.NewIF(localctx.(*Control_if_expContext).Get_expresion().GetEx(), localctx.(*Control_if_expContext).GetBloqueElse().GetList(), lista_null, Ast.ELSE_EXPRESION, fila, columna, true)
		lista_entonces := localctx.(*Control_if_expContext).Get_bloque_else_if_exp().GetList()
		lista_entonces.Add(Else)
		localctx.(*Control_if_expContext).ex = exp_ins.NewIF(localctx.(*Control_if_expContext).Get_expresion().GetEx(), localctx.(*Control_if_expContext).GetBloqueIf().GetList(), lista_entonces, Ast.IF_EXPRESION, fila, columna, true)

	}

	return localctx
}

// IBloque_else_if_expContext is an interface to support dynamic dispatch.
type IBloque_else_if_expContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_else_if_exp returns the _else_if_exp rule contexts.
	Get_else_if_exp() IElse_if_expContext

	// Set_else_if_exp sets the _else_if_exp rule contexts.
	Set_else_if_exp(IElse_if_expContext)

	// GetLista returns the lista rule context list.
	GetLista() []IElse_if_expContext

	// SetLista sets the lista rule context list.
	SetLista([]IElse_if_expContext)

	// GetList returns the list attribute.
	GetList() *arraylist.List

	// SetList sets the list attribute.
	SetList(*arraylist.List)

	// IsBloque_else_if_expContext differentiates from other interfaces.
	IsBloque_else_if_expContext()
}

type Bloque_else_if_expContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	list         *arraylist.List
	_else_if_exp IElse_if_expContext
	lista        []IElse_if_expContext
}

func NewEmptyBloque_else_if_expContext() *Bloque_else_if_expContext {
	var p = new(Bloque_else_if_expContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_bloque_else_if_exp
	return p
}

func (*Bloque_else_if_expContext) IsBloque_else_if_expContext() {}

func NewBloque_else_if_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bloque_else_if_expContext {
	var p = new(Bloque_else_if_expContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_bloque_else_if_exp

	return p
}

func (s *Bloque_else_if_expContext) GetParser() antlr.Parser { return s.parser }

func (s *Bloque_else_if_expContext) Get_else_if_exp() IElse_if_expContext { return s._else_if_exp }

func (s *Bloque_else_if_expContext) Set_else_if_exp(v IElse_if_expContext) { s._else_if_exp = v }

func (s *Bloque_else_if_expContext) GetLista() []IElse_if_expContext { return s.lista }

func (s *Bloque_else_if_expContext) SetLista(v []IElse_if_expContext) { s.lista = v }

func (s *Bloque_else_if_expContext) GetList() *arraylist.List { return s.list }

func (s *Bloque_else_if_expContext) SetList(v *arraylist.List) { s.list = v }

func (s *Bloque_else_if_expContext) AllElse_if_exp() []IElse_if_expContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElse_if_expContext)(nil)).Elem())
	var tst = make([]IElse_if_expContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElse_if_expContext)
		}
	}

	return tst
}

func (s *Bloque_else_if_expContext) Else_if_exp(i int) IElse_if_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElse_if_expContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElse_if_expContext)
}

func (s *Bloque_else_if_expContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bloque_else_if_expContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bloque_else_if_expContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterBloque_else_if_exp(s)
	}
}

func (s *Bloque_else_if_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitBloque_else_if_exp(s)
	}
}

func (p *Nparser) Bloque_else_if_exp() (localctx IBloque_else_if_expContext) {
	this := p
	_ = this

	localctx = NewBloque_else_if_expContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, NparserRULE_bloque_else_if_exp)
	localctx.(*Bloque_else_if_expContext).list = arraylist.New()

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
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(333)

				var _x = p.Else_if_exp()

				localctx.(*Bloque_else_if_expContext)._else_if_exp = _x
			}
			localctx.(*Bloque_else_if_expContext).lista = append(localctx.(*Bloque_else_if_expContext).lista, localctx.(*Bloque_else_if_expContext)._else_if_exp)

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(336)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	listas := localctx.(*Bloque_else_if_expContext).GetLista()
	for _, e := range listas {
		localctx.(*Bloque_else_if_expContext).list.Add(e.GetEx())
	}

	return localctx
}

// IElse_if_expContext is an interface to support dynamic dispatch.
type IElse_if_expContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ELSE returns the _ELSE token.
	Get_ELSE() antlr.Token

	// Set_ELSE sets the _ELSE token.
	Set_ELSE(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Get_bloque returns the _bloque rule contexts.
	Get_bloque() IBloqueContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// Set_bloque sets the _bloque rule contexts.
	Set_bloque(IBloqueContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsElse_if_expContext differentiates from other interfaces.
	IsElse_if_expContext()
}

type Else_if_expContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	ex         Ast.Instruccion
	_ELSE      antlr.Token
	_expresion IExpresionContext
	_bloque    IBloqueContext
}

func NewEmptyElse_if_expContext() *Else_if_expContext {
	var p = new(Else_if_expContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_else_if_exp
	return p
}

func (*Else_if_expContext) IsElse_if_expContext() {}

func NewElse_if_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_if_expContext {
	var p = new(Else_if_expContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_else_if_exp

	return p
}

func (s *Else_if_expContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_if_expContext) Get_ELSE() antlr.Token { return s._ELSE }

func (s *Else_if_expContext) Set_ELSE(v antlr.Token) { s._ELSE = v }

func (s *Else_if_expContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *Else_if_expContext) Get_bloque() IBloqueContext { return s._bloque }

func (s *Else_if_expContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *Else_if_expContext) Set_bloque(v IBloqueContext) { s._bloque = v }

func (s *Else_if_expContext) GetEx() Ast.Instruccion { return s.ex }

func (s *Else_if_expContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *Else_if_expContext) ELSE() antlr.TerminalNode {
	return s.GetToken(NparserELSE, 0)
}

func (s *Else_if_expContext) IF() antlr.TerminalNode {
	return s.GetToken(NparserIF, 0)
}

func (s *Else_if_expContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Else_if_expContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Else_if_expContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_if_expContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Else_if_expContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterElse_if_exp(s)
	}
}

func (s *Else_if_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitElse_if_exp(s)
	}
}

func (p *Nparser) Else_if_exp() (localctx IElse_if_expContext) {
	this := p
	_ = this

	localctx = NewElse_if_expContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, NparserRULE_else_if_exp)

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
		p.SetState(340)

		var _m = p.Match(NparserELSE)

		localctx.(*Else_if_expContext)._ELSE = _m
	}
	{
		p.SetState(341)
		p.Match(NparserIF)
	}
	{
		p.SetState(342)

		var _x = p.expresion(0)

		localctx.(*Else_if_expContext)._expresion = _x
	}
	{
		p.SetState(343)

		var _x = p.Bloque()

		localctx.(*Else_if_expContext)._bloque = _x
	}

	fila := (func() int {
		if localctx.(*Else_if_expContext).Get_ELSE() == nil {
			return 0
		} else {
			return localctx.(*Else_if_expContext).Get_ELSE().GetLine()
		}
	}())
	columna := (func() int {
		if localctx.(*Else_if_expContext).Get_ELSE() == nil {
			return 0
		} else {
			return localctx.(*Else_if_expContext).Get_ELSE().GetColumn()
		}
	}())
	columna++
	lista_null := arraylist.New()
	localctx.(*Else_if_expContext).ex = exp_ins.NewIF(localctx.(*Else_if_expContext).Get_expresion().GetEx(), localctx.(*Else_if_expContext).Get_bloque().GetList(), lista_null, Ast.ELSEIF_EXPRESION, fila, columna, true)

	return localctx
}

// IControl_expresionContext is an interface to support dynamic dispatch.
type IControl_expresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_control_if_exp returns the _control_if_exp rule contexts.
	Get_control_if_exp() IControl_if_expContext

	// Get_control_match_exp returns the _control_match_exp rule contexts.
	Get_control_match_exp() IControl_match_expContext

	// Get_control_loop_exp returns the _control_loop_exp rule contexts.
	Get_control_loop_exp() IControl_loop_expContext

	// Set_control_if_exp sets the _control_if_exp rule contexts.
	Set_control_if_exp(IControl_if_expContext)

	// Set_control_match_exp sets the _control_match_exp rule contexts.
	Set_control_match_exp(IControl_match_expContext)

	// Set_control_loop_exp sets the _control_loop_exp rule contexts.
	Set_control_loop_exp(IControl_loop_expContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsControl_expresionContext differentiates from other interfaces.
	IsControl_expresionContext()
}

type Control_expresionContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	ex                 Ast.Instruccion
	_control_if_exp    IControl_if_expContext
	_control_match_exp IControl_match_expContext
	_control_loop_exp  IControl_loop_expContext
}

func NewEmptyControl_expresionContext() *Control_expresionContext {
	var p = new(Control_expresionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_control_expresion
	return p
}

func (*Control_expresionContext) IsControl_expresionContext() {}

func NewControl_expresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Control_expresionContext {
	var p = new(Control_expresionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_control_expresion

	return p
}

func (s *Control_expresionContext) GetParser() antlr.Parser { return s.parser }

func (s *Control_expresionContext) Get_control_if_exp() IControl_if_expContext {
	return s._control_if_exp
}

func (s *Control_expresionContext) Get_control_match_exp() IControl_match_expContext {
	return s._control_match_exp
}

func (s *Control_expresionContext) Get_control_loop_exp() IControl_loop_expContext {
	return s._control_loop_exp
}

func (s *Control_expresionContext) Set_control_if_exp(v IControl_if_expContext) {
	s._control_if_exp = v
}

func (s *Control_expresionContext) Set_control_match_exp(v IControl_match_expContext) {
	s._control_match_exp = v
}

func (s *Control_expresionContext) Set_control_loop_exp(v IControl_loop_expContext) {
	s._control_loop_exp = v
}

func (s *Control_expresionContext) GetEx() Ast.Instruccion { return s.ex }

func (s *Control_expresionContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *Control_expresionContext) Control_if_exp() IControl_if_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControl_if_expContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControl_if_expContext)
}

func (s *Control_expresionContext) Control_match_exp() IControl_match_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControl_match_expContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControl_match_expContext)
}

func (s *Control_expresionContext) Control_loop_exp() IControl_loop_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControl_loop_expContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControl_loop_expContext)
}

func (s *Control_expresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Control_expresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Control_expresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterControl_expresion(s)
	}
}

func (s *Control_expresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitControl_expresion(s)
	}
}

func (p *Nparser) Control_expresion() (localctx IControl_expresionContext) {
	this := p
	_ = this

	localctx = NewControl_expresionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, NparserRULE_control_expresion)

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

	p.SetState(355)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(346)

			var _x = p.Control_if_exp()

			localctx.(*Control_expresionContext)._control_if_exp = _x
		}
		localctx.(*Control_expresionContext).ex = localctx.(*Control_expresionContext).Get_control_if_exp().GetEx()

	case NparserMATCH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(349)

			var _x = p.Control_match_exp()

			localctx.(*Control_expresionContext)._control_match_exp = _x
		}
		localctx.(*Control_expresionContext).ex = localctx.(*Control_expresionContext).Get_control_match_exp().GetEx()

	case NparserLOOP:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(352)

			var _x = p.Control_loop_exp()

			localctx.(*Control_expresionContext)._control_loop_exp = _x
		}
		localctx.(*Control_expresionContext).ex = localctx.(*Control_expresionContext).Get_control_loop_exp().GetEx()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IControl_matchContext is an interface to support dynamic dispatch.
type IControl_matchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_MATCH returns the _MATCH token.
	Get_MATCH() antlr.Token

	// Set_MATCH sets the _MATCH token.
	Set_MATCH(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Get_control_case returns the _control_case rule contexts.
	Get_control_case() IControl_caseContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// Set_control_case sets the _control_case rule contexts.
	Set_control_case(IControl_caseContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsControl_matchContext differentiates from other interfaces.
	IsControl_matchContext()
}

type Control_matchContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	ex            Ast.Instruccion
	_MATCH        antlr.Token
	_expresion    IExpresionContext
	_control_case IControl_caseContext
}

func NewEmptyControl_matchContext() *Control_matchContext {
	var p = new(Control_matchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_control_match
	return p
}

func (*Control_matchContext) IsControl_matchContext() {}

func NewControl_matchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Control_matchContext {
	var p = new(Control_matchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_control_match

	return p
}

func (s *Control_matchContext) GetParser() antlr.Parser { return s.parser }

func (s *Control_matchContext) Get_MATCH() antlr.Token { return s._MATCH }

func (s *Control_matchContext) Set_MATCH(v antlr.Token) { s._MATCH = v }

func (s *Control_matchContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *Control_matchContext) Get_control_case() IControl_caseContext { return s._control_case }

func (s *Control_matchContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *Control_matchContext) Set_control_case(v IControl_caseContext) { s._control_case = v }

func (s *Control_matchContext) GetEx() Ast.Instruccion { return s.ex }

func (s *Control_matchContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *Control_matchContext) MATCH() antlr.TerminalNode {
	return s.GetToken(NparserMATCH, 0)
}

func (s *Control_matchContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Control_matchContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(NparserLLAVE_IZQ, 0)
}

func (s *Control_matchContext) Control_case() IControl_caseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControl_caseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControl_caseContext)
}

func (s *Control_matchContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(NparserLLAVE_DER, 0)
}

func (s *Control_matchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Control_matchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Control_matchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterControl_match(s)
	}
}

func (s *Control_matchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitControl_match(s)
	}
}

func (p *Nparser) Control_match() (localctx IControl_matchContext) {
	this := p
	_ = this

	localctx = NewControl_matchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, NparserRULE_control_match)

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
		p.SetState(357)

		var _m = p.Match(NparserMATCH)

		localctx.(*Control_matchContext)._MATCH = _m
	}
	{
		p.SetState(358)

		var _x = p.expresion(0)

		localctx.(*Control_matchContext)._expresion = _x
	}
	{
		p.SetState(359)
		p.Match(NparserLLAVE_IZQ)
	}
	{
		p.SetState(360)

		var _x = p.Control_case()

		localctx.(*Control_matchContext)._control_case = _x
	}
	{
		p.SetState(361)
		p.Match(NparserLLAVE_DER)
	}

	fila := (func() int {
		if localctx.(*Control_matchContext).Get_MATCH() == nil {
			return 0
		} else {
			return localctx.(*Control_matchContext).Get_MATCH().GetLine()
		}
	}())
	columna := (func() int {
		if localctx.(*Control_matchContext).Get_MATCH() == nil {
			return 0
		} else {
			return localctx.(*Control_matchContext).Get_MATCH().GetLine()
		}
	}()) - 1
	localctx.(*Control_matchContext).ex = exp_ins.NewMatch(localctx.(*Control_matchContext).Get_expresion().GetEx(), localctx.(*Control_matchContext).Get_control_case().GetList(), Ast.MATCH, fila, columna)

	return localctx
}

// IControl_caseContext is an interface to support dynamic dispatch.
type IControl_caseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_cases returns the _cases rule contexts.
	Get_cases() ICasesContext

	// Set_cases sets the _cases rule contexts.
	Set_cases(ICasesContext)

	// GetLista returns the lista rule context list.
	GetLista() []ICasesContext

	// SetLista sets the lista rule context list.
	SetLista([]ICasesContext)

	// GetList returns the list attribute.
	GetList() *arraylist.List

	// SetList sets the list attribute.
	SetList(*arraylist.List)

	// IsControl_caseContext differentiates from other interfaces.
	IsControl_caseContext()
}

type Control_caseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	list   *arraylist.List
	_cases ICasesContext
	lista  []ICasesContext
}

func NewEmptyControl_caseContext() *Control_caseContext {
	var p = new(Control_caseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_control_case
	return p
}

func (*Control_caseContext) IsControl_caseContext() {}

func NewControl_caseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Control_caseContext {
	var p = new(Control_caseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_control_case

	return p
}

func (s *Control_caseContext) GetParser() antlr.Parser { return s.parser }

func (s *Control_caseContext) Get_cases() ICasesContext { return s._cases }

func (s *Control_caseContext) Set_cases(v ICasesContext) { s._cases = v }

func (s *Control_caseContext) GetLista() []ICasesContext { return s.lista }

func (s *Control_caseContext) SetLista(v []ICasesContext) { s.lista = v }

func (s *Control_caseContext) GetList() *arraylist.List { return s.list }

func (s *Control_caseContext) SetList(v *arraylist.List) { s.list = v }

func (s *Control_caseContext) AllCases() []ICasesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICasesContext)(nil)).Elem())
	var tst = make([]ICasesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICasesContext)
		}
	}

	return tst
}

func (s *Control_caseContext) Cases(i int) ICasesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICasesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICasesContext)
}

func (s *Control_caseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Control_caseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Control_caseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterControl_case(s)
	}
}

func (s *Control_caseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitControl_case(s)
	}
}

func (p *Nparser) Control_case() (localctx IControl_caseContext) {
	this := p
	_ = this

	localctx = NewControl_caseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, NparserRULE_control_case)
	localctx.(*Control_caseContext).list = arraylist.New()
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
	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-19)&-(0x1f+1)) == 0 && ((1<<uint((_la-19)))&((1<<(NparserTRUE-19))|(1<<(NparserFALSE-19))|(1<<(NparserNUMERO-19))|(1<<(NparserDECIMAL-19))|(1<<(NparserID-19))|(1<<(NparserDEFAULT-19)))) != 0) || (((_la-74)&-(0x1f+1)) == 0 && ((1<<uint((_la-74)))&((1<<(NparserRESTA-74))|(1<<(NparserNOT-74))|(1<<(NparserPAR_IZQ-74))|(1<<(NparserCADENA-74))|(1<<(NparserCARACTER-74)))) != 0) {
		{
			p.SetState(364)

			var _x = p.Cases()

			localctx.(*Control_caseContext)._cases = _x
		}
		localctx.(*Control_caseContext).lista = append(localctx.(*Control_caseContext).lista, localctx.(*Control_caseContext)._cases)

		p.SetState(367)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	fmt.Println("Entro a control case")
	listas := localctx.(*Control_caseContext).GetLista()
	for _, e := range listas {
		localctx.(*Control_caseContext).list.Add(e.GetEx())
	}

	return localctx
}

// ICasesContext is an interface to support dynamic dispatch.
type ICasesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CASE returns the _CASE token.
	Get_CASE() antlr.Token

	// Set_CASE sets the _CASE token.
	Set_CASE(antlr.Token)

	// Get_case_match returns the _case_match rule contexts.
	Get_case_match() ICase_matchContext

	// Get_bloque returns the _bloque rule contexts.
	Get_bloque() IBloqueContext

	// Set_case_match sets the _case_match rule contexts.
	Set_case_match(ICase_matchContext)

	// Set_bloque sets the _bloque rule contexts.
	Set_bloque(IBloqueContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsCasesContext differentiates from other interfaces.
	IsCasesContext()
}

type CasesContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	ex          Ast.Instruccion
	_case_match ICase_matchContext
	_CASE       antlr.Token
	_bloque     IBloqueContext
}

func NewEmptyCasesContext() *CasesContext {
	var p = new(CasesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_cases
	return p
}

func (*CasesContext) IsCasesContext() {}

func NewCasesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasesContext {
	var p = new(CasesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_cases

	return p
}

func (s *CasesContext) GetParser() antlr.Parser { return s.parser }

func (s *CasesContext) Get_CASE() antlr.Token { return s._CASE }

func (s *CasesContext) Set_CASE(v antlr.Token) { s._CASE = v }

func (s *CasesContext) Get_case_match() ICase_matchContext { return s._case_match }

func (s *CasesContext) Get_bloque() IBloqueContext { return s._bloque }

func (s *CasesContext) Set_case_match(v ICase_matchContext) { s._case_match = v }

func (s *CasesContext) Set_bloque(v IBloqueContext) { s._bloque = v }

func (s *CasesContext) GetEx() Ast.Instruccion { return s.ex }

func (s *CasesContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *CasesContext) Case_match() ICase_matchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICase_matchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICase_matchContext)
}

func (s *CasesContext) CASE() antlr.TerminalNode {
	return s.GetToken(NparserCASE, 0)
}

func (s *CasesContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *CasesContext) COMA() antlr.TerminalNode {
	return s.GetToken(NparserCOMA, 0)
}

func (s *CasesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CasesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterCases(s)
	}
}

func (s *CasesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitCases(s)
	}
}

func (p *Nparser) Cases() (localctx ICasesContext) {
	this := p
	_ = this

	localctx = NewCasesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, NparserRULE_cases)

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
		p.SetState(371)

		var _x = p.case_match(0)

		localctx.(*CasesContext)._case_match = _x
	}
	{
		p.SetState(372)

		var _m = p.Match(NparserCASE)

		localctx.(*CasesContext)._CASE = _m
	}
	{
		p.SetState(373)

		var _x = p.Bloque()

		localctx.(*CasesContext)._bloque = _x
	}
	{
		p.SetState(374)
		p.Match(NparserCOMA)
	}

	fmt.Println("Entro cases")
	fila := (func() int {
		if localctx.(*CasesContext).Get_CASE() == nil {
			return 0
		} else {
			return localctx.(*CasesContext).Get_CASE().GetLine()
		}
	}())
	columna := (func() int {
		if localctx.(*CasesContext).Get_CASE() == nil {
			return 0
		} else {
			return localctx.(*CasesContext).Get_CASE().GetLine()
		}
	}()) - 1
	//Verificar si lo que vienen es un default
	listaTemp := localctx.(*CasesContext).Get_case_match().GetList()
	_, tipo := listaTemp.GetValue(0).(Ast.Abstracto).GetTipo()
	if tipo == Ast.DEFAULT {
		localctx.(*CasesContext).ex = exp_ins.NewCase(localctx.(*CasesContext).Get_case_match().GetList(), localctx.(*CasesContext).Get_bloque().GetList(), Ast.CASE, fila, columna, true)
	} else {
		localctx.(*CasesContext).ex = exp_ins.NewCase(localctx.(*CasesContext).Get_case_match().GetList(), localctx.(*CasesContext).Get_bloque().GetList(), Ast.CASE, fila, columna, false)
	}

	return localctx
}

// ICase_matchContext is an interface to support dynamic dispatch.
type ICase_matchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_DEFAULT returns the _DEFAULT token.
	Get_DEFAULT() antlr.Token

	// Set_DEFAULT sets the _DEFAULT token.
	Set_DEFAULT(antlr.Token)

	// GetLista_cases returns the lista_cases rule contexts.
	GetLista_cases() ICase_matchContext

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// SetLista_cases sets the lista_cases rule contexts.
	SetLista_cases(ICase_matchContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetList returns the list attribute.
	GetList() *arraylist.List

	// SetList sets the list attribute.
	SetList(*arraylist.List)

	// IsCase_matchContext differentiates from other interfaces.
	IsCase_matchContext()
}

type Case_matchContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	list        *arraylist.List
	lista_cases ICase_matchContext
	_expresion  IExpresionContext
	_DEFAULT    antlr.Token
}

func NewEmptyCase_matchContext() *Case_matchContext {
	var p = new(Case_matchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_case_match
	return p
}

func (*Case_matchContext) IsCase_matchContext() {}

func NewCase_matchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Case_matchContext {
	var p = new(Case_matchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_case_match

	return p
}

func (s *Case_matchContext) GetParser() antlr.Parser { return s.parser }

func (s *Case_matchContext) Get_DEFAULT() antlr.Token { return s._DEFAULT }

func (s *Case_matchContext) Set_DEFAULT(v antlr.Token) { s._DEFAULT = v }

func (s *Case_matchContext) GetLista_cases() ICase_matchContext { return s.lista_cases }

func (s *Case_matchContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *Case_matchContext) SetLista_cases(v ICase_matchContext) { s.lista_cases = v }

func (s *Case_matchContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *Case_matchContext) GetList() *arraylist.List { return s.list }

func (s *Case_matchContext) SetList(v *arraylist.List) { s.list = v }

func (s *Case_matchContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Case_matchContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(NparserDEFAULT, 0)
}

func (s *Case_matchContext) O() antlr.TerminalNode {
	return s.GetToken(NparserO, 0)
}

func (s *Case_matchContext) Case_match() ICase_matchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICase_matchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICase_matchContext)
}

func (s *Case_matchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Case_matchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Case_matchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterCase_match(s)
	}
}

func (s *Case_matchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitCase_match(s)
	}
}

func (p *Nparser) Case_match() (localctx ICase_matchContext) {
	return p.case_match(0)
}

func (p *Nparser) case_match(_p int) (localctx ICase_matchContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewCase_matchContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICase_matchContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, NparserRULE_case_match, _p)
	localctx.(*Case_matchContext).list = arraylist.New()

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
	p.SetState(383)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserTRUE, NparserFALSE, NparserNUMERO, NparserDECIMAL, NparserID, NparserRESTA, NparserNOT, NparserPAR_IZQ, NparserCADENA, NparserCARACTER:
		{
			p.SetState(378)

			var _x = p.expresion(0)

			localctx.(*Case_matchContext)._expresion = _x
		}

		localctx.(*Case_matchContext).list.Add(localctx.(*Case_matchContext).Get_expresion().GetEx())

	case NparserDEFAULT:
		{
			p.SetState(381)

			var _m = p.Match(NparserDEFAULT)

			localctx.(*Case_matchContext)._DEFAULT = _m
		}

		fila := (func() int {
			if localctx.(*Case_matchContext).Get_DEFAULT() == nil {
				return 0
			} else {
				return localctx.(*Case_matchContext).Get_DEFAULT().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Case_matchContext).Get_DEFAULT() == nil {
				return 0
			} else {
				return localctx.(*Case_matchContext).Get_DEFAULT().GetColumn()
			}
		}())
		expresion := expresiones.NewPrimitivo(false, Ast.DEFAULT, fila, columna)
		localctx.(*Case_matchContext).list.Add(expresion)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCase_matchContext(p, _parentctx, _parentState)
			localctx.(*Case_matchContext).lista_cases = _prevctx
			p.PushNewRecursionContext(localctx, _startState, NparserRULE_case_match)
			p.SetState(385)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(386)
				p.Match(NparserO)
			}
			{
				p.SetState(387)

				var _x = p.expresion(0)

				localctx.(*Case_matchContext)._expresion = _x
			}

			localctx.(*Case_matchContext).GetLista_cases().GetList().Add(localctx.(*Case_matchContext).Get_expresion().GetEx())
			localctx.(*Case_matchContext).list = localctx.(*Case_matchContext).GetLista_cases().GetList()

		}
		p.SetState(394)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}

	return localctx
}

// IControl_match_expContext is an interface to support dynamic dispatch.
type IControl_match_expContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_MATCH returns the _MATCH token.
	Get_MATCH() antlr.Token

	// Set_MATCH sets the _MATCH token.
	Set_MATCH(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Get_control_case_exp returns the _control_case_exp rule contexts.
	Get_control_case_exp() IControl_case_expContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// Set_control_case_exp sets the _control_case_exp rule contexts.
	Set_control_case_exp(IControl_case_expContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsControl_match_expContext differentiates from other interfaces.
	IsControl_match_expContext()
}

type Control_match_expContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	ex                Ast.Instruccion
	_MATCH            antlr.Token
	_expresion        IExpresionContext
	_control_case_exp IControl_case_expContext
}

func NewEmptyControl_match_expContext() *Control_match_expContext {
	var p = new(Control_match_expContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_control_match_exp
	return p
}

func (*Control_match_expContext) IsControl_match_expContext() {}

func NewControl_match_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Control_match_expContext {
	var p = new(Control_match_expContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_control_match_exp

	return p
}

func (s *Control_match_expContext) GetParser() antlr.Parser { return s.parser }

func (s *Control_match_expContext) Get_MATCH() antlr.Token { return s._MATCH }

func (s *Control_match_expContext) Set_MATCH(v antlr.Token) { s._MATCH = v }

func (s *Control_match_expContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *Control_match_expContext) Get_control_case_exp() IControl_case_expContext {
	return s._control_case_exp
}

func (s *Control_match_expContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *Control_match_expContext) Set_control_case_exp(v IControl_case_expContext) {
	s._control_case_exp = v
}

func (s *Control_match_expContext) GetEx() Ast.Instruccion { return s.ex }

func (s *Control_match_expContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *Control_match_expContext) MATCH() antlr.TerminalNode {
	return s.GetToken(NparserMATCH, 0)
}

func (s *Control_match_expContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Control_match_expContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(NparserLLAVE_IZQ, 0)
}

func (s *Control_match_expContext) Control_case_exp() IControl_case_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControl_case_expContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControl_case_expContext)
}

func (s *Control_match_expContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(NparserLLAVE_DER, 0)
}

func (s *Control_match_expContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Control_match_expContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Control_match_expContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterControl_match_exp(s)
	}
}

func (s *Control_match_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitControl_match_exp(s)
	}
}

func (p *Nparser) Control_match_exp() (localctx IControl_match_expContext) {
	this := p
	_ = this

	localctx = NewControl_match_expContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, NparserRULE_control_match_exp)

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
		p.SetState(395)

		var _m = p.Match(NparserMATCH)

		localctx.(*Control_match_expContext)._MATCH = _m
	}
	{
		p.SetState(396)

		var _x = p.expresion(0)

		localctx.(*Control_match_expContext)._expresion = _x
	}
	{
		p.SetState(397)
		p.Match(NparserLLAVE_IZQ)
	}
	{
		p.SetState(398)

		var _x = p.Control_case_exp()

		localctx.(*Control_match_expContext)._control_case_exp = _x
	}
	{
		p.SetState(399)
		p.Match(NparserLLAVE_DER)
	}

	fila := (func() int {
		if localctx.(*Control_match_expContext).Get_MATCH() == nil {
			return 0
		} else {
			return localctx.(*Control_match_expContext).Get_MATCH().GetLine()
		}
	}())
	columna := (func() int {
		if localctx.(*Control_match_expContext).Get_MATCH() == nil {
			return 0
		} else {
			return localctx.(*Control_match_expContext).Get_MATCH().GetLine()
		}
	}()) - 1
	fmt.Println("Entro a control match")
	localctx.(*Control_match_expContext).ex = exp_ins.NewMatch(localctx.(*Control_match_expContext).Get_expresion().GetEx(), localctx.(*Control_match_expContext).Get_control_case_exp().GetList(), Ast.MATCH_EXPRESION, fila, columna)

	return localctx
}

// IControl_case_expContext is an interface to support dynamic dispatch.
type IControl_case_expContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_cases_exp returns the _cases_exp rule contexts.
	Get_cases_exp() ICases_expContext

	// Set_cases_exp sets the _cases_exp rule contexts.
	Set_cases_exp(ICases_expContext)

	// GetLista returns the lista rule context list.
	GetLista() []ICases_expContext

	// SetLista sets the lista rule context list.
	SetLista([]ICases_expContext)

	// GetList returns the list attribute.
	GetList() *arraylist.List

	// SetList sets the list attribute.
	SetList(*arraylist.List)

	// IsControl_case_expContext differentiates from other interfaces.
	IsControl_case_expContext()
}

type Control_case_expContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	list       *arraylist.List
	_cases_exp ICases_expContext
	lista      []ICases_expContext
}

func NewEmptyControl_case_expContext() *Control_case_expContext {
	var p = new(Control_case_expContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_control_case_exp
	return p
}

func (*Control_case_expContext) IsControl_case_expContext() {}

func NewControl_case_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Control_case_expContext {
	var p = new(Control_case_expContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_control_case_exp

	return p
}

func (s *Control_case_expContext) GetParser() antlr.Parser { return s.parser }

func (s *Control_case_expContext) Get_cases_exp() ICases_expContext { return s._cases_exp }

func (s *Control_case_expContext) Set_cases_exp(v ICases_expContext) { s._cases_exp = v }

func (s *Control_case_expContext) GetLista() []ICases_expContext { return s.lista }

func (s *Control_case_expContext) SetLista(v []ICases_expContext) { s.lista = v }

func (s *Control_case_expContext) GetList() *arraylist.List { return s.list }

func (s *Control_case_expContext) SetList(v *arraylist.List) { s.list = v }

func (s *Control_case_expContext) AllCases_exp() []ICases_expContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICases_expContext)(nil)).Elem())
	var tst = make([]ICases_expContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICases_expContext)
		}
	}

	return tst
}

func (s *Control_case_expContext) Cases_exp(i int) ICases_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICases_expContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICases_expContext)
}

func (s *Control_case_expContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Control_case_expContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Control_case_expContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterControl_case_exp(s)
	}
}

func (s *Control_case_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitControl_case_exp(s)
	}
}

func (p *Nparser) Control_case_exp() (localctx IControl_case_expContext) {
	this := p
	_ = this

	localctx = NewControl_case_expContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, NparserRULE_control_case_exp)
	localctx.(*Control_case_expContext).list = arraylist.New()
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
	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-19)&-(0x1f+1)) == 0 && ((1<<uint((_la-19)))&((1<<(NparserTRUE-19))|(1<<(NparserFALSE-19))|(1<<(NparserNUMERO-19))|(1<<(NparserDECIMAL-19))|(1<<(NparserID-19))|(1<<(NparserDEFAULT-19)))) != 0) || (((_la-74)&-(0x1f+1)) == 0 && ((1<<uint((_la-74)))&((1<<(NparserRESTA-74))|(1<<(NparserNOT-74))|(1<<(NparserPAR_IZQ-74))|(1<<(NparserCADENA-74))|(1<<(NparserCARACTER-74)))) != 0) {
		{
			p.SetState(402)

			var _x = p.Cases_exp()

			localctx.(*Control_case_expContext)._cases_exp = _x
		}
		localctx.(*Control_case_expContext).lista = append(localctx.(*Control_case_expContext).lista, localctx.(*Control_case_expContext)._cases_exp)

		p.SetState(405)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listas := localctx.(*Control_case_expContext).GetLista()
	for _, e := range listas {
		localctx.(*Control_case_expContext).list.Add(e.GetEx())
	}

	return localctx
}

// ICases_expContext is an interface to support dynamic dispatch.
type ICases_expContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CASE returns the _CASE token.
	Get_CASE() antlr.Token

	// Set_CASE sets the _CASE token.
	Set_CASE(antlr.Token)

	// Get_case_match_exp returns the _case_match_exp rule contexts.
	Get_case_match_exp() ICase_match_expContext

	// Get_bloque returns the _bloque rule contexts.
	Get_bloque() IBloqueContext

	// Set_case_match_exp sets the _case_match_exp rule contexts.
	Set_case_match_exp(ICase_match_expContext)

	// Set_bloque sets the _bloque rule contexts.
	Set_bloque(IBloqueContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsCases_expContext differentiates from other interfaces.
	IsCases_expContext()
}

type Cases_expContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	ex              Ast.Instruccion
	_case_match_exp ICase_match_expContext
	_CASE           antlr.Token
	_bloque         IBloqueContext
}

func NewEmptyCases_expContext() *Cases_expContext {
	var p = new(Cases_expContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_cases_exp
	return p
}

func (*Cases_expContext) IsCases_expContext() {}

func NewCases_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cases_expContext {
	var p = new(Cases_expContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_cases_exp

	return p
}

func (s *Cases_expContext) GetParser() antlr.Parser { return s.parser }

func (s *Cases_expContext) Get_CASE() antlr.Token { return s._CASE }

func (s *Cases_expContext) Set_CASE(v antlr.Token) { s._CASE = v }

func (s *Cases_expContext) Get_case_match_exp() ICase_match_expContext { return s._case_match_exp }

func (s *Cases_expContext) Get_bloque() IBloqueContext { return s._bloque }

func (s *Cases_expContext) Set_case_match_exp(v ICase_match_expContext) { s._case_match_exp = v }

func (s *Cases_expContext) Set_bloque(v IBloqueContext) { s._bloque = v }

func (s *Cases_expContext) GetEx() Ast.Instruccion { return s.ex }

func (s *Cases_expContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *Cases_expContext) Case_match_exp() ICase_match_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICase_match_expContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICase_match_expContext)
}

func (s *Cases_expContext) CASE() antlr.TerminalNode {
	return s.GetToken(NparserCASE, 0)
}

func (s *Cases_expContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Cases_expContext) COMA() antlr.TerminalNode {
	return s.GetToken(NparserCOMA, 0)
}

func (s *Cases_expContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cases_expContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cases_expContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterCases_exp(s)
	}
}

func (s *Cases_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitCases_exp(s)
	}
}

func (p *Nparser) Cases_exp() (localctx ICases_expContext) {
	this := p
	_ = this

	localctx = NewCases_expContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, NparserRULE_cases_exp)

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
		p.SetState(409)

		var _x = p.case_match_exp(0)

		localctx.(*Cases_expContext)._case_match_exp = _x
	}
	{
		p.SetState(410)

		var _m = p.Match(NparserCASE)

		localctx.(*Cases_expContext)._CASE = _m
	}
	{
		p.SetState(411)

		var _x = p.Bloque()

		localctx.(*Cases_expContext)._bloque = _x
	}
	{
		p.SetState(412)
		p.Match(NparserCOMA)
	}

	fila := (func() int {
		if localctx.(*Cases_expContext).Get_CASE() == nil {
			return 0
		} else {
			return localctx.(*Cases_expContext).Get_CASE().GetLine()
		}
	}())
	columna := (func() int {
		if localctx.(*Cases_expContext).Get_CASE() == nil {
			return 0
		} else {
			return localctx.(*Cases_expContext).Get_CASE().GetLine()
		}
	}()) - 1
	//Verificar si lo que vienen es un default
	listaTemp := localctx.(*Cases_expContext).Get_case_match_exp().GetList()
	_, tipo := listaTemp.GetValue(0).(Ast.Abstracto).GetTipo()
	if tipo == Ast.DEFAULT {
		localctx.(*Cases_expContext).ex = exp_ins.NewCase(localctx.(*Cases_expContext).Get_case_match_exp().GetList(), localctx.(*Cases_expContext).Get_bloque().GetList(), Ast.CASE_EXPRESION, fila, columna, true)
	} else {
		localctx.(*Cases_expContext).ex = exp_ins.NewCase(localctx.(*Cases_expContext).Get_case_match_exp().GetList(), localctx.(*Cases_expContext).Get_bloque().GetList(), Ast.CASE_EXPRESION, fila, columna, false)
	}

	return localctx
}

// ICase_match_expContext is an interface to support dynamic dispatch.
type ICase_match_expContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_DEFAULT returns the _DEFAULT token.
	Get_DEFAULT() antlr.Token

	// Set_DEFAULT sets the _DEFAULT token.
	Set_DEFAULT(antlr.Token)

	// GetLista_cases returns the lista_cases rule contexts.
	GetLista_cases() ICase_match_expContext

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// SetLista_cases sets the lista_cases rule contexts.
	SetLista_cases(ICase_match_expContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetList returns the list attribute.
	GetList() *arraylist.List

	// SetList sets the list attribute.
	SetList(*arraylist.List)

	// IsCase_match_expContext differentiates from other interfaces.
	IsCase_match_expContext()
}

type Case_match_expContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	list        *arraylist.List
	lista_cases ICase_match_expContext
	_expresion  IExpresionContext
	_DEFAULT    antlr.Token
}

func NewEmptyCase_match_expContext() *Case_match_expContext {
	var p = new(Case_match_expContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_case_match_exp
	return p
}

func (*Case_match_expContext) IsCase_match_expContext() {}

func NewCase_match_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Case_match_expContext {
	var p = new(Case_match_expContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_case_match_exp

	return p
}

func (s *Case_match_expContext) GetParser() antlr.Parser { return s.parser }

func (s *Case_match_expContext) Get_DEFAULT() antlr.Token { return s._DEFAULT }

func (s *Case_match_expContext) Set_DEFAULT(v antlr.Token) { s._DEFAULT = v }

func (s *Case_match_expContext) GetLista_cases() ICase_match_expContext { return s.lista_cases }

func (s *Case_match_expContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *Case_match_expContext) SetLista_cases(v ICase_match_expContext) { s.lista_cases = v }

func (s *Case_match_expContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *Case_match_expContext) GetList() *arraylist.List { return s.list }

func (s *Case_match_expContext) SetList(v *arraylist.List) { s.list = v }

func (s *Case_match_expContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Case_match_expContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(NparserDEFAULT, 0)
}

func (s *Case_match_expContext) O() antlr.TerminalNode {
	return s.GetToken(NparserO, 0)
}

func (s *Case_match_expContext) Case_match_exp() ICase_match_expContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICase_match_expContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICase_match_expContext)
}

func (s *Case_match_expContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Case_match_expContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Case_match_expContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterCase_match_exp(s)
	}
}

func (s *Case_match_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitCase_match_exp(s)
	}
}

func (p *Nparser) Case_match_exp() (localctx ICase_match_expContext) {
	return p.case_match_exp(0)
}

func (p *Nparser) case_match_exp(_p int) (localctx ICase_match_expContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewCase_match_expContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICase_match_expContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, NparserRULE_case_match_exp, _p)
	localctx.(*Case_match_expContext).list = arraylist.New()

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
	p.SetState(421)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserTRUE, NparserFALSE, NparserNUMERO, NparserDECIMAL, NparserID, NparserRESTA, NparserNOT, NparserPAR_IZQ, NparserCADENA, NparserCARACTER:
		{
			p.SetState(416)

			var _x = p.expresion(0)

			localctx.(*Case_match_expContext)._expresion = _x
		}

		localctx.(*Case_match_expContext).list.Add(localctx.(*Case_match_expContext).Get_expresion().GetEx())

	case NparserDEFAULT:
		{
			p.SetState(419)

			var _m = p.Match(NparserDEFAULT)

			localctx.(*Case_match_expContext)._DEFAULT = _m
		}

		fila := (func() int {
			if localctx.(*Case_match_expContext).Get_DEFAULT() == nil {
				return 0
			} else {
				return localctx.(*Case_match_expContext).Get_DEFAULT().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Case_match_expContext).Get_DEFAULT() == nil {
				return 0
			} else {
				return localctx.(*Case_match_expContext).Get_DEFAULT().GetColumn()
			}
		}())
		expresion := expresiones.NewPrimitivo(false, Ast.DEFAULT, fila, columna)
		localctx.(*Case_match_expContext).list.Add(expresion)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCase_match_expContext(p, _parentctx, _parentState)
			localctx.(*Case_match_expContext).lista_cases = _prevctx
			p.PushNewRecursionContext(localctx, _startState, NparserRULE_case_match_exp)
			p.SetState(423)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(424)
				p.Match(NparserO)
			}
			{
				p.SetState(425)

				var _x = p.expresion(0)

				localctx.(*Case_match_expContext)._expresion = _x
			}

			localctx.(*Case_match_expContext).GetLista_cases().GetList().Add(localctx.(*Case_match_expContext).Get_expresion().GetEx())
			localctx.(*Case_match_expContext).list = localctx.(*Case_match_expContext).GetLista_cases().GetList()

		}
		p.SetState(432)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}

	return localctx
}

// IIreturnContext is an interface to support dynamic dispatch.
type IIreturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RETURN returns the _RETURN token.
	Get_RETURN() antlr.Token

	// Set_RETURN sets the _RETURN token.
	Set_RETURN(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsIreturnContext differentiates from other interfaces.
	IsIreturnContext()
}

type IreturnContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	ex         Ast.Instruccion
	_RETURN    antlr.Token
	_expresion IExpresionContext
}

func NewEmptyIreturnContext() *IreturnContext {
	var p = new(IreturnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_ireturn
	return p
}

func (*IreturnContext) IsIreturnContext() {}

func NewIreturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IreturnContext {
	var p = new(IreturnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_ireturn

	return p
}

func (s *IreturnContext) GetParser() antlr.Parser { return s.parser }

func (s *IreturnContext) Get_RETURN() antlr.Token { return s._RETURN }

func (s *IreturnContext) Set_RETURN(v antlr.Token) { s._RETURN = v }

func (s *IreturnContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *IreturnContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *IreturnContext) GetEx() Ast.Instruccion { return s.ex }

func (s *IreturnContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *IreturnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(NparserRETURN, 0)
}

func (s *IreturnContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *IreturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IreturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IreturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterIreturn(s)
	}
}

func (s *IreturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitIreturn(s)
	}
}

func (p *Nparser) Ireturn() (localctx IIreturnContext) {
	this := p
	_ = this

	localctx = NewIreturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, NparserRULE_ireturn)

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

	p.SetState(439)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(433)

			var _m = p.Match(NparserRETURN)

			localctx.(*IreturnContext)._RETURN = _m
		}

		fila := (func() int {
			if localctx.(*IreturnContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*IreturnContext).Get_RETURN().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*IreturnContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*IreturnContext).Get_RETURN().GetColumn()
			}
		}())
		localctx.(*IreturnContext).ex = transferencia.NewReturn(Ast.RETURN, nil, fila, columna)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(435)

			var _m = p.Match(NparserRETURN)

			localctx.(*IreturnContext)._RETURN = _m
		}
		{
			p.SetState(436)

			var _x = p.expresion(0)

			localctx.(*IreturnContext)._expresion = _x
		}

		fila := (func() int {
			if localctx.(*IreturnContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*IreturnContext).Get_RETURN().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*IreturnContext).Get_RETURN() == nil {
				return 0
			} else {
				return localctx.(*IreturnContext).Get_RETURN().GetColumn()
			}
		}())
		localctx.(*IreturnContext).ex = transferencia.NewReturn(Ast.RETURN_EXPRESION, localctx.(*IreturnContext).Get_expresion().GetEx(), fila, columna)

	}

	return localctx
}

// IIbreakContext is an interface to support dynamic dispatch.
type IIbreakContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_BREAK returns the _BREAK token.
	Get_BREAK() antlr.Token

	// Set_BREAK sets the _BREAK token.
	Set_BREAK(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsIbreakContext differentiates from other interfaces.
	IsIbreakContext()
}

type IbreakContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	ex         Ast.Instruccion
	_BREAK     antlr.Token
	_expresion IExpresionContext
}

func NewEmptyIbreakContext() *IbreakContext {
	var p = new(IbreakContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_ibreak
	return p
}

func (*IbreakContext) IsIbreakContext() {}

func NewIbreakContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IbreakContext {
	var p = new(IbreakContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_ibreak

	return p
}

func (s *IbreakContext) GetParser() antlr.Parser { return s.parser }

func (s *IbreakContext) Get_BREAK() antlr.Token { return s._BREAK }

func (s *IbreakContext) Set_BREAK(v antlr.Token) { s._BREAK = v }

func (s *IbreakContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *IbreakContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *IbreakContext) GetEx() Ast.Instruccion { return s.ex }

func (s *IbreakContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *IbreakContext) BREAK() antlr.TerminalNode {
	return s.GetToken(NparserBREAK, 0)
}

func (s *IbreakContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *IbreakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IbreakContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IbreakContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterIbreak(s)
	}
}

func (s *IbreakContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitIbreak(s)
	}
}

func (p *Nparser) Ibreak() (localctx IIbreakContext) {
	this := p
	_ = this

	localctx = NewIbreakContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, NparserRULE_ibreak)

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

	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(441)

			var _m = p.Match(NparserBREAK)

			localctx.(*IbreakContext)._BREAK = _m
		}

		fila := (func() int {
			if localctx.(*IbreakContext).Get_BREAK() == nil {
				return 0
			} else {
				return localctx.(*IbreakContext).Get_BREAK().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*IbreakContext).Get_BREAK() == nil {
				return 0
			} else {
				return localctx.(*IbreakContext).Get_BREAK().GetColumn()
			}
		}())
		localctx.(*IbreakContext).ex = transferencia.NewBreak(Ast.BREAK, nil, fila, columna)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(443)

			var _m = p.Match(NparserBREAK)

			localctx.(*IbreakContext)._BREAK = _m
		}
		{
			p.SetState(444)

			var _x = p.expresion(0)

			localctx.(*IbreakContext)._expresion = _x
		}

		fila := (func() int {
			if localctx.(*IbreakContext).Get_BREAK() == nil {
				return 0
			} else {
				return localctx.(*IbreakContext).Get_BREAK().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*IbreakContext).Get_BREAK() == nil {
				return 0
			} else {
				return localctx.(*IbreakContext).Get_BREAK().GetColumn()
			}
		}())
		localctx.(*IbreakContext).ex = transferencia.NewBreak(Ast.BREAK_EXPRESION, localctx.(*IbreakContext).Get_expresion().GetEx(), fila, columna)

	}

	return localctx
}

// IIcontinueContext is an interface to support dynamic dispatch.
type IIcontinueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CONTINUE returns the _CONTINUE token.
	Get_CONTINUE() antlr.Token

	// Set_CONTINUE sets the _CONTINUE token.
	Set_CONTINUE(antlr.Token)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsIcontinueContext differentiates from other interfaces.
	IsIcontinueContext()
}

type IcontinueContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	ex        Ast.Instruccion
	_CONTINUE antlr.Token
}

func NewEmptyIcontinueContext() *IcontinueContext {
	var p = new(IcontinueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_icontinue
	return p
}

func (*IcontinueContext) IsIcontinueContext() {}

func NewIcontinueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IcontinueContext {
	var p = new(IcontinueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_icontinue

	return p
}

func (s *IcontinueContext) GetParser() antlr.Parser { return s.parser }

func (s *IcontinueContext) Get_CONTINUE() antlr.Token { return s._CONTINUE }

func (s *IcontinueContext) Set_CONTINUE(v antlr.Token) { s._CONTINUE = v }

func (s *IcontinueContext) GetEx() Ast.Instruccion { return s.ex }

func (s *IcontinueContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *IcontinueContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(NparserCONTINUE, 0)
}

func (s *IcontinueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IcontinueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IcontinueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterIcontinue(s)
	}
}

func (s *IcontinueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitIcontinue(s)
	}
}

func (p *Nparser) Icontinue() (localctx IIcontinueContext) {
	this := p
	_ = this

	localctx = NewIcontinueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, NparserRULE_icontinue)

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
		p.SetState(449)

		var _m = p.Match(NparserCONTINUE)

		localctx.(*IcontinueContext)._CONTINUE = _m
	}

	fila := (func() int {
		if localctx.(*IcontinueContext).Get_CONTINUE() == nil {
			return 0
		} else {
			return localctx.(*IcontinueContext).Get_CONTINUE().GetLine()
		}
	}())
	columna := (func() int {
		if localctx.(*IcontinueContext).Get_CONTINUE() == nil {
			return 0
		} else {
			return localctx.(*IcontinueContext).Get_CONTINUE().GetColumn()
		}
	}())
	localctx.(*IcontinueContext).ex = transferencia.NewContinue(fila, columna)

	return localctx
}

// IControl_loopContext is an interface to support dynamic dispatch.
type IControl_loopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_LOOP returns the _LOOP token.
	Get_LOOP() antlr.Token

	// Set_LOOP sets the _LOOP token.
	Set_LOOP(antlr.Token)

	// Get_bloque returns the _bloque rule contexts.
	Get_bloque() IBloqueContext

	// Set_bloque sets the _bloque rule contexts.
	Set_bloque(IBloqueContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsControl_loopContext differentiates from other interfaces.
	IsControl_loopContext()
}

type Control_loopContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	ex      Ast.Instruccion
	_LOOP   antlr.Token
	_bloque IBloqueContext
}

func NewEmptyControl_loopContext() *Control_loopContext {
	var p = new(Control_loopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_control_loop
	return p
}

func (*Control_loopContext) IsControl_loopContext() {}

func NewControl_loopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Control_loopContext {
	var p = new(Control_loopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_control_loop

	return p
}

func (s *Control_loopContext) GetParser() antlr.Parser { return s.parser }

func (s *Control_loopContext) Get_LOOP() antlr.Token { return s._LOOP }

func (s *Control_loopContext) Set_LOOP(v antlr.Token) { s._LOOP = v }

func (s *Control_loopContext) Get_bloque() IBloqueContext { return s._bloque }

func (s *Control_loopContext) Set_bloque(v IBloqueContext) { s._bloque = v }

func (s *Control_loopContext) GetEx() Ast.Instruccion { return s.ex }

func (s *Control_loopContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *Control_loopContext) LOOP() antlr.TerminalNode {
	return s.GetToken(NparserLOOP, 0)
}

func (s *Control_loopContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Control_loopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Control_loopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Control_loopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterControl_loop(s)
	}
}

func (s *Control_loopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitControl_loop(s)
	}
}

func (p *Nparser) Control_loop() (localctx IControl_loopContext) {
	this := p
	_ = this

	localctx = NewControl_loopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, NparserRULE_control_loop)

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
		p.SetState(452)

		var _m = p.Match(NparserLOOP)

		localctx.(*Control_loopContext)._LOOP = _m
	}
	{
		p.SetState(453)

		var _x = p.Bloque()

		localctx.(*Control_loopContext)._bloque = _x
	}

	fila := (func() int {
		if localctx.(*Control_loopContext).Get_LOOP() == nil {
			return 0
		} else {
			return localctx.(*Control_loopContext).Get_LOOP().GetLine()
		}
	}())
	columna := (func() int {
		if localctx.(*Control_loopContext).Get_LOOP() == nil {
			return 0
		} else {
			return localctx.(*Control_loopContext).Get_LOOP().GetColumn()
		}
	}())
	localctx.(*Control_loopContext).ex = bucles.NewLoop(Ast.LOOP, localctx.(*Control_loopContext).Get_bloque().GetList(), fila, columna)

	return localctx
}

// IControl_loop_expContext is an interface to support dynamic dispatch.
type IControl_loop_expContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_LOOP returns the _LOOP token.
	Get_LOOP() antlr.Token

	// Set_LOOP sets the _LOOP token.
	Set_LOOP(antlr.Token)

	// Get_bloque returns the _bloque rule contexts.
	Get_bloque() IBloqueContext

	// Set_bloque sets the _bloque rule contexts.
	Set_bloque(IBloqueContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsControl_loop_expContext differentiates from other interfaces.
	IsControl_loop_expContext()
}

type Control_loop_expContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	ex      Ast.Instruccion
	_LOOP   antlr.Token
	_bloque IBloqueContext
}

func NewEmptyControl_loop_expContext() *Control_loop_expContext {
	var p = new(Control_loop_expContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_control_loop_exp
	return p
}

func (*Control_loop_expContext) IsControl_loop_expContext() {}

func NewControl_loop_expContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Control_loop_expContext {
	var p = new(Control_loop_expContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_control_loop_exp

	return p
}

func (s *Control_loop_expContext) GetParser() antlr.Parser { return s.parser }

func (s *Control_loop_expContext) Get_LOOP() antlr.Token { return s._LOOP }

func (s *Control_loop_expContext) Set_LOOP(v antlr.Token) { s._LOOP = v }

func (s *Control_loop_expContext) Get_bloque() IBloqueContext { return s._bloque }

func (s *Control_loop_expContext) Set_bloque(v IBloqueContext) { s._bloque = v }

func (s *Control_loop_expContext) GetEx() Ast.Instruccion { return s.ex }

func (s *Control_loop_expContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *Control_loop_expContext) LOOP() antlr.TerminalNode {
	return s.GetToken(NparserLOOP, 0)
}

func (s *Control_loop_expContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Control_loop_expContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Control_loop_expContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Control_loop_expContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterControl_loop_exp(s)
	}
}

func (s *Control_loop_expContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitControl_loop_exp(s)
	}
}

func (p *Nparser) Control_loop_exp() (localctx IControl_loop_expContext) {
	this := p
	_ = this

	localctx = NewControl_loop_expContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, NparserRULE_control_loop_exp)

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
		p.SetState(456)

		var _m = p.Match(NparserLOOP)

		localctx.(*Control_loop_expContext)._LOOP = _m
	}
	{
		p.SetState(457)

		var _x = p.Bloque()

		localctx.(*Control_loop_expContext)._bloque = _x
	}

	fila := (func() int {
		if localctx.(*Control_loop_expContext).Get_LOOP() == nil {
			return 0
		} else {
			return localctx.(*Control_loop_expContext).Get_LOOP().GetLine()
		}
	}())
	columna := (func() int {
		if localctx.(*Control_loop_expContext).Get_LOOP() == nil {
			return 0
		} else {
			return localctx.(*Control_loop_expContext).Get_LOOP().GetColumn()
		}
	}())
	localctx.(*Control_loop_expContext).ex = bucles.NewLoop(Ast.LOOP_EXPRESION, localctx.(*Control_loop_expContext).Get_bloque().GetList(), fila, columna)

	return localctx
}

func (p *Nparser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 6:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	case 18:
		var t *Case_matchContext = nil
		if localctx != nil {
			t = localctx.(*Case_matchContext)
		}
		return p.Case_match_Sempred(t, predIndex)

	case 22:
		var t *Case_match_expContext = nil
		if localctx != nil {
			t = localctx.(*Case_match_expContext)
		}
		return p.Case_match_exp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Nparser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 9)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Nparser) Case_match_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Nparser) Case_match_exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
