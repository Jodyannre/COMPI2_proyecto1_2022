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

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 91, 328,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 7, 4, 42, 10, 4, 12, 4, 14, 4, 45, 11, 4,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 63, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 5, 6, 132, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 5, 7, 144, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 170, 10, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 7, 8, 202, 10, 8, 12, 8, 14, 8, 205, 11, 8, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 5, 9, 221, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 257, 10, 10, 3, 11, 6, 11,
	260, 10, 11, 13, 11, 14, 11, 261, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 308, 10, 13, 3,
	14, 6, 14, 311, 10, 14, 13, 14, 14, 14, 312, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 2,
	3, 14, 17, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 2, 7,
	4, 2, 76, 76, 78, 78, 3, 2, 72, 74, 3, 2, 76, 77, 4, 2, 64, 68, 70, 70,
	4, 2, 68, 68, 70, 70, 2, 353, 2, 32, 3, 2, 2, 2, 4, 35, 3, 2, 2, 2, 6,
	43, 3, 2, 2, 2, 8, 62, 3, 2, 2, 2, 10, 131, 3, 2, 2, 2, 12, 143, 3, 2,
	2, 2, 14, 169, 3, 2, 2, 2, 16, 220, 3, 2, 2, 2, 18, 256, 3, 2, 2, 2, 20,
	259, 3, 2, 2, 2, 22, 265, 3, 2, 2, 2, 24, 307, 3, 2, 2, 2, 26, 310, 3,
	2, 2, 2, 28, 316, 3, 2, 2, 2, 30, 324, 3, 2, 2, 2, 32, 33, 5, 6, 4, 2,
	33, 34, 8, 2, 1, 2, 34, 3, 3, 2, 2, 2, 35, 36, 7, 82, 2, 2, 36, 37, 5,
	6, 4, 2, 37, 38, 7, 83, 2, 2, 38, 39, 8, 3, 1, 2, 39, 5, 3, 2, 2, 2, 40,
	42, 5, 8, 5, 2, 41, 40, 3, 2, 2, 2, 42, 45, 3, 2, 2, 2, 43, 41, 3, 2, 2,
	2, 43, 44, 3, 2, 2, 2, 44, 46, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 46, 47,
	8, 4, 1, 2, 47, 7, 3, 2, 2, 2, 48, 49, 5, 14, 8, 2, 49, 50, 8, 5, 1, 2,
	50, 63, 3, 2, 2, 2, 51, 52, 5, 10, 6, 2, 52, 53, 7, 63, 2, 2, 53, 54, 8,
	5, 1, 2, 54, 63, 3, 2, 2, 2, 55, 56, 5, 12, 7, 2, 56, 57, 7, 63, 2, 2,
	57, 58, 8, 5, 1, 2, 58, 63, 3, 2, 2, 2, 59, 60, 5, 18, 10, 2, 60, 61, 8,
	5, 1, 2, 61, 63, 3, 2, 2, 2, 62, 48, 3, 2, 2, 2, 62, 51, 3, 2, 2, 2, 62,
	55, 3, 2, 2, 2, 62, 59, 3, 2, 2, 2, 63, 9, 3, 2, 2, 2, 64, 65, 7, 17, 2,
	2, 65, 66, 7, 51, 2, 2, 66, 67, 7, 71, 2, 2, 67, 68, 5, 14, 8, 2, 68, 69,
	8, 6, 1, 2, 69, 132, 3, 2, 2, 2, 70, 71, 7, 17, 2, 2, 71, 72, 7, 51, 2,
	2, 72, 73, 7, 71, 2, 2, 73, 74, 5, 30, 16, 2, 74, 75, 8, 6, 1, 2, 75, 132,
	3, 2, 2, 2, 76, 77, 7, 17, 2, 2, 77, 78, 7, 51, 2, 2, 78, 79, 7, 59, 2,
	2, 79, 80, 5, 16, 9, 2, 80, 81, 7, 71, 2, 2, 81, 82, 5, 14, 8, 2, 82, 83,
	8, 6, 1, 2, 83, 132, 3, 2, 2, 2, 84, 85, 7, 17, 2, 2, 85, 86, 7, 51, 2,
	2, 86, 87, 7, 59, 2, 2, 87, 88, 5, 16, 9, 2, 88, 89, 7, 71, 2, 2, 89, 90,
	5, 30, 16, 2, 90, 91, 8, 6, 1, 2, 91, 132, 3, 2, 2, 2, 92, 93, 7, 17, 2,
	2, 93, 94, 7, 16, 2, 2, 94, 95, 7, 51, 2, 2, 95, 96, 7, 71, 2, 2, 96, 97,
	5, 14, 8, 2, 97, 98, 8, 6, 1, 2, 98, 132, 3, 2, 2, 2, 99, 100, 7, 17, 2,
	2, 100, 101, 7, 16, 2, 2, 101, 102, 7, 51, 2, 2, 102, 103, 7, 71, 2, 2,
	103, 104, 5, 30, 16, 2, 104, 105, 8, 6, 1, 2, 105, 132, 3, 2, 2, 2, 106,
	107, 7, 17, 2, 2, 107, 108, 7, 16, 2, 2, 108, 109, 7, 51, 2, 2, 109, 110,
	7, 59, 2, 2, 110, 111, 5, 16, 9, 2, 111, 112, 8, 6, 1, 2, 112, 132, 3,
	2, 2, 2, 113, 114, 7, 17, 2, 2, 114, 115, 7, 16, 2, 2, 115, 116, 7, 51,
	2, 2, 116, 117, 7, 59, 2, 2, 117, 118, 5, 16, 9, 2, 118, 119, 7, 71, 2,
	2, 119, 120, 5, 14, 8, 2, 120, 121, 8, 6, 1, 2, 121, 132, 3, 2, 2, 2, 122,
	123, 7, 17, 2, 2, 123, 124, 7, 16, 2, 2, 124, 125, 7, 51, 2, 2, 125, 126,
	7, 59, 2, 2, 126, 127, 5, 16, 9, 2, 127, 128, 7, 71, 2, 2, 128, 129, 5,
	30, 16, 2, 129, 130, 8, 6, 1, 2, 130, 132, 3, 2, 2, 2, 131, 64, 3, 2, 2,
	2, 131, 70, 3, 2, 2, 2, 131, 76, 3, 2, 2, 2, 131, 84, 3, 2, 2, 2, 131,
	92, 3, 2, 2, 2, 131, 99, 3, 2, 2, 2, 131, 106, 3, 2, 2, 2, 131, 113, 3,
	2, 2, 2, 131, 122, 3, 2, 2, 2, 132, 11, 3, 2, 2, 2, 133, 134, 7, 51, 2,
	2, 134, 135, 7, 71, 2, 2, 135, 136, 5, 14, 8, 2, 136, 137, 8, 7, 1, 2,
	137, 144, 3, 2, 2, 2, 138, 139, 7, 51, 2, 2, 139, 140, 7, 71, 2, 2, 140,
	141, 5, 24, 13, 2, 141, 142, 8, 7, 1, 2, 142, 144, 3, 2, 2, 2, 143, 133,
	3, 2, 2, 2, 143, 138, 3, 2, 2, 2, 144, 13, 3, 2, 2, 2, 145, 146, 8, 8,
	1, 2, 146, 147, 9, 2, 2, 2, 147, 148, 5, 14, 8, 17, 148, 149, 8, 8, 1,
	2, 149, 170, 3, 2, 2, 2, 150, 151, 7, 80, 2, 2, 151, 152, 5, 14, 8, 2,
	152, 153, 7, 81, 2, 2, 153, 154, 8, 8, 1, 2, 154, 170, 3, 2, 2, 2, 155,
	156, 7, 51, 2, 2, 156, 170, 8, 8, 1, 2, 157, 158, 7, 21, 2, 2, 158, 170,
	8, 8, 1, 2, 159, 160, 7, 22, 2, 2, 160, 170, 8, 8, 1, 2, 161, 162, 7, 88,
	2, 2, 162, 170, 8, 8, 1, 2, 163, 164, 7, 49, 2, 2, 164, 170, 8, 8, 1, 2,
	165, 166, 7, 48, 2, 2, 166, 170, 8, 8, 1, 2, 167, 168, 7, 86, 2, 2, 168,
	170, 8, 8, 1, 2, 169, 145, 3, 2, 2, 2, 169, 150, 3, 2, 2, 2, 169, 155,
	3, 2, 2, 2, 169, 157, 3, 2, 2, 2, 169, 159, 3, 2, 2, 2, 169, 161, 3, 2,
	2, 2, 169, 163, 3, 2, 2, 2, 169, 165, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2,
	170, 203, 3, 2, 2, 2, 171, 172, 12, 16, 2, 2, 172, 173, 9, 3, 2, 2, 173,
	174, 5, 14, 8, 17, 174, 175, 8, 8, 1, 2, 175, 202, 3, 2, 2, 2, 176, 177,
	12, 15, 2, 2, 177, 178, 9, 4, 2, 2, 178, 179, 5, 14, 8, 16, 179, 180, 8,
	8, 1, 2, 180, 202, 3, 2, 2, 2, 181, 182, 12, 14, 2, 2, 182, 183, 9, 5,
	2, 2, 183, 184, 5, 14, 8, 15, 184, 185, 8, 8, 1, 2, 185, 202, 3, 2, 2,
	2, 186, 187, 12, 13, 2, 2, 187, 188, 9, 6, 2, 2, 188, 189, 5, 14, 8, 14,
	189, 190, 8, 8, 1, 2, 190, 202, 3, 2, 2, 2, 191, 192, 12, 12, 2, 2, 192,
	193, 7, 56, 2, 2, 193, 194, 5, 14, 8, 13, 194, 195, 8, 8, 1, 2, 195, 202,
	3, 2, 2, 2, 196, 197, 12, 11, 2, 2, 197, 198, 7, 54, 2, 2, 198, 199, 5,
	14, 8, 12, 199, 200, 8, 8, 1, 2, 200, 202, 3, 2, 2, 2, 201, 171, 3, 2,
	2, 2, 201, 176, 3, 2, 2, 2, 201, 181, 3, 2, 2, 2, 201, 186, 3, 2, 2, 2,
	201, 191, 3, 2, 2, 2, 201, 196, 3, 2, 2, 2, 202, 205, 3, 2, 2, 2, 203,
	201, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2, 204, 15, 3, 2, 2, 2, 205, 203, 3,
	2, 2, 2, 206, 207, 7, 3, 2, 2, 207, 221, 8, 9, 1, 2, 208, 209, 7, 4, 2,
	2, 209, 221, 8, 9, 1, 2, 210, 211, 7, 6, 2, 2, 211, 221, 8, 9, 1, 2, 212,
	213, 7, 5, 2, 2, 213, 221, 8, 9, 1, 2, 214, 215, 7, 7, 2, 2, 215, 221,
	8, 9, 1, 2, 216, 217, 7, 8, 2, 2, 217, 221, 8, 9, 1, 2, 218, 219, 7, 9,
	2, 2, 219, 221, 8, 9, 1, 2, 220, 206, 3, 2, 2, 2, 220, 208, 3, 2, 2, 2,
	220, 210, 3, 2, 2, 2, 220, 212, 3, 2, 2, 2, 220, 214, 3, 2, 2, 2, 220,
	216, 3, 2, 2, 2, 220, 218, 3, 2, 2, 2, 221, 17, 3, 2, 2, 2, 222, 223, 7,
	36, 2, 2, 223, 224, 7, 80, 2, 2, 224, 225, 5, 14, 8, 2, 225, 226, 7, 81,
	2, 2, 226, 227, 5, 4, 3, 2, 227, 228, 8, 10, 1, 2, 228, 257, 3, 2, 2, 2,
	229, 230, 7, 36, 2, 2, 230, 231, 7, 80, 2, 2, 231, 232, 5, 14, 8, 2, 232,
	233, 7, 81, 2, 2, 233, 234, 5, 4, 3, 2, 234, 235, 7, 37, 2, 2, 235, 236,
	5, 4, 3, 2, 236, 237, 8, 10, 1, 2, 237, 257, 3, 2, 2, 2, 238, 239, 7, 36,
	2, 2, 239, 240, 7, 80, 2, 2, 240, 241, 5, 14, 8, 2, 241, 242, 7, 81, 2,
	2, 242, 243, 5, 4, 3, 2, 243, 244, 5, 20, 11, 2, 244, 245, 8, 10, 1, 2,
	245, 257, 3, 2, 2, 2, 246, 247, 7, 36, 2, 2, 247, 248, 7, 80, 2, 2, 248,
	249, 5, 14, 8, 2, 249, 250, 7, 81, 2, 2, 250, 251, 5, 4, 3, 2, 251, 252,
	5, 20, 11, 2, 252, 253, 7, 37, 2, 2, 253, 254, 5, 4, 3, 2, 254, 255, 8,
	10, 1, 2, 255, 257, 3, 2, 2, 2, 256, 222, 3, 2, 2, 2, 256, 229, 3, 2, 2,
	2, 256, 238, 3, 2, 2, 2, 256, 246, 3, 2, 2, 2, 257, 19, 3, 2, 2, 2, 258,
	260, 5, 22, 12, 2, 259, 258, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261, 259,
	3, 2, 2, 2, 261, 262, 3, 2, 2, 2, 262, 263, 3, 2, 2, 2, 263, 264, 8, 11,
	1, 2, 264, 21, 3, 2, 2, 2, 265, 266, 7, 37, 2, 2, 266, 267, 7, 36, 2, 2,
	267, 268, 7, 80, 2, 2, 268, 269, 5, 14, 8, 2, 269, 270, 7, 81, 2, 2, 270,
	271, 5, 4, 3, 2, 271, 272, 8, 12, 1, 2, 272, 23, 3, 2, 2, 2, 273, 274,
	7, 36, 2, 2, 274, 275, 7, 80, 2, 2, 275, 276, 5, 14, 8, 2, 276, 277, 7,
	81, 2, 2, 277, 278, 5, 4, 3, 2, 278, 279, 8, 13, 1, 2, 279, 308, 3, 2,
	2, 2, 280, 281, 7, 36, 2, 2, 281, 282, 7, 80, 2, 2, 282, 283, 5, 14, 8,
	2, 283, 284, 7, 81, 2, 2, 284, 285, 5, 4, 3, 2, 285, 286, 7, 37, 2, 2,
	286, 287, 5, 4, 3, 2, 287, 288, 8, 13, 1, 2, 288, 308, 3, 2, 2, 2, 289,
	290, 7, 36, 2, 2, 290, 291, 7, 80, 2, 2, 291, 292, 5, 14, 8, 2, 292, 293,
	7, 81, 2, 2, 293, 294, 5, 4, 3, 2, 294, 295, 5, 26, 14, 2, 295, 296, 8,
	13, 1, 2, 296, 308, 3, 2, 2, 2, 297, 298, 7, 36, 2, 2, 298, 299, 7, 80,
	2, 2, 299, 300, 5, 14, 8, 2, 300, 301, 7, 81, 2, 2, 301, 302, 5, 4, 3,
	2, 302, 303, 5, 26, 14, 2, 303, 304, 7, 37, 2, 2, 304, 305, 5, 4, 3, 2,
	305, 306, 8, 13, 1, 2, 306, 308, 3, 2, 2, 2, 307, 273, 3, 2, 2, 2, 307,
	280, 3, 2, 2, 2, 307, 289, 3, 2, 2, 2, 307, 297, 3, 2, 2, 2, 308, 25, 3,
	2, 2, 2, 309, 311, 5, 28, 15, 2, 310, 309, 3, 2, 2, 2, 311, 312, 3, 2,
	2, 2, 312, 310, 3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313, 314, 3, 2, 2, 2,
	314, 315, 8, 14, 1, 2, 315, 27, 3, 2, 2, 2, 316, 317, 7, 37, 2, 2, 317,
	318, 7, 36, 2, 2, 318, 319, 7, 80, 2, 2, 319, 320, 5, 14, 8, 2, 320, 321,
	7, 81, 2, 2, 321, 322, 5, 4, 3, 2, 322, 323, 8, 15, 1, 2, 323, 29, 3, 2,
	2, 2, 324, 325, 5, 24, 13, 2, 325, 326, 8, 16, 1, 2, 326, 31, 3, 2, 2,
	2, 14, 43, 62, 131, 143, 169, 201, 203, 220, 256, 261, 307, 312,
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
	"bloque_else_if_exp", "else_if_exp", "control_expresion",
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
		p.SetState(30)

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
		p.SetState(33)
		p.Match(NparserLLAVE_IZQ)
	}
	{
		p.SetState(34)

		var _x = p.Instrucciones()

		localctx.(*BloqueContext)._instrucciones = _x
	}
	{
		p.SetState(35)
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
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NparserLET)|(1<<NparserTRUE)|(1<<NparserFALSE))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(NparserIF-34))|(1<<(NparserNUMERO-34))|(1<<(NparserDECIMAL-34))|(1<<(NparserID-34)))) != 0) || (((_la-74)&-(0x1f+1)) == 0 && ((1<<uint((_la-74)))&((1<<(NparserRESTA-74))|(1<<(NparserNOT-74))|(1<<(NparserPAR_IZQ-74))|(1<<(NparserCADENA-74))|(1<<(NparserCARACTER-74)))) != 0) {
		{
			p.SetState(38)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(43)
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

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// Set_declaracion sets the _declaracion rule contexts.
	Set_declaracion(IDeclaracionContext)

	// Set_asignacion sets the _asignacion rule contexts.
	Set_asignacion(IAsignacionContext)

	// Set_control_if sets the _control_if rule contexts.
	Set_control_if(IControl_ifContext)

	// GetEx returns the ex attribute.
	GetEx() interface{}

	// SetEx sets the ex attribute.
	SetEx(interface{})

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	ex           interface{}
	_expresion   IExpresionContext
	_declaracion IDeclaracionContext
	_asignacion  IAsignacionContext
	_control_if  IControl_ifContext
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

func (s *InstruccionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *InstruccionContext) Set_declaracion(v IDeclaracionContext) { s._declaracion = v }

func (s *InstruccionContext) Set_asignacion(v IAsignacionContext) { s._asignacion = v }

func (s *InstruccionContext) Set_control_if(v IControl_ifContext) { s._control_if = v }

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

	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(46)

			var _x = p.expresion(0)

			localctx.(*InstruccionContext)._expresion = _x
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_expresion().GetEx()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)

			var _x = p.Declaracion()

			localctx.(*InstruccionContext)._declaracion = _x
		}
		{
			p.SetState(50)
			p.Match(NparserPUNTOCOMA)
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_declaracion().GetEx()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(53)

			var _x = p.Asignacion()

			localctx.(*InstruccionContext)._asignacion = _x
		}
		{
			p.SetState(54)
			p.Match(NparserPUNTOCOMA)
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_asignacion().GetEx()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(57)

			var _x = p.Control_if()

			localctx.(*InstruccionContext)._control_if = _x
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_control_if().GetEx()

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

	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(63)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(64)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(65)

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
			p.SetState(68)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(69)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(70)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(71)

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
			p.SetState(74)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(75)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(76)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(77)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(78)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(79)

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
			p.SetState(82)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(83)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(84)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(85)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(86)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(87)

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
			p.SetState(90)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(91)
			p.Match(NparserMUT)
		}
		{
			p.SetState(92)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(93)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(94)

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
			p.SetState(97)
			p.Match(NparserLET)
		}
		{
			p.SetState(98)
			p.Match(NparserMUT)
		}
		{
			p.SetState(99)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(100)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(101)

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
			p.SetState(104)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(105)
			p.Match(NparserMUT)
		}
		{
			p.SetState(106)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(107)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(108)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}

		valor := expresiones.NewPrimitivo(nil, Ast.NULL)
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
			true, false, Ast.VOID, valor, fila, columna)

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(111)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(112)
			p.Match(NparserMUT)
		}
		{
			p.SetState(113)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(114)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(115)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(116)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(117)

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
			p.SetState(120)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(121)
			p.Match(NparserMUT)
		}
		{
			p.SetState(122)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(123)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(124)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(125)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(126)

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

	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)

			var _m = p.Match(NparserID)

			localctx.(*AsignacionContext)._ID = _m
		}
		{
			p.SetState(132)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(133)

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
			p.SetState(136)

			var _m = p.Match(NparserID)

			localctx.(*AsignacionContext)._ID = _m
		}
		{
			p.SetState(137)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(138)

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
	p.SetState(167)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserRESTA, NparserNOT:
		{
			p.SetState(144)

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
			p.SetState(145)

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
			p.SetState(148)
			p.Match(NparserPAR_IZQ)
		}
		{
			p.SetState(149)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext)._expresion = _x
		}
		{
			p.SetState(150)
			p.Match(NparserPAR_DER)
		}

		localctx.(*ExpresionContext).ex = localctx.(*ExpresionContext).Get_expresion().GetEx()

	case NparserID:
		{
			p.SetState(153)

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
		localctx.(*ExpresionContext).ex = expresiones.NewIdentificador(id, Ast.IDENTIFICADOR)

	case NparserTRUE:
		{
			p.SetState(155)
			p.Match(NparserTRUE)
		}

		valor := true
		localctx.(*ExpresionContext).ex = expresiones.NewPrimitivo(valor, Ast.BOOLEAN)

	case NparserFALSE:
		{
			p.SetState(157)
			p.Match(NparserFALSE)
		}

		valor := false
		localctx.(*ExpresionContext).ex = expresiones.NewPrimitivo(valor, Ast.BOOLEAN)

	case NparserCARACTER:
		{
			p.SetState(159)

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
		localctx.(*ExpresionContext).ex = expresiones.NewPrimitivo(valor, Ast.CHAR)

	case NparserDECIMAL:
		{
			p.SetState(161)

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
		localctx.(*ExpresionContext).ex = expresiones.NewPrimitivo(valor, Ast.F64)

	case NparserNUMERO:
		{
			p.SetState(163)

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
		localctx.(*ExpresionContext).ex = expresiones.NewPrimitivo(valor, Ast.I64)

	case NparserCADENA:
		{
			p.SetState(165)

			var _m = p.Match(NparserCADENA)

			localctx.(*ExpresionContext)._CADENA = _m
		}

		valor := (func() string {
			if localctx.(*ExpresionContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_CADENA().GetText()
			}
		}())
		valor = valor[1 : len(valor)-1]
		localctx.(*ExpresionContext).ex = expresiones.NewPrimitivo(valor, Ast.STR)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(199)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).op_izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(169)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(170)

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
					p.SetState(171)

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
				p.SetState(174)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(175)

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
					p.SetState(176)

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
				p.SetState(179)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(180)

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
					p.SetState(181)

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
				p.SetState(184)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(185)

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
					p.SetState(186)

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
				p.SetState(189)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(190)

					var _m = p.Match(NparserAND)

					localctx.(*ExpresionContext)._AND = _m
				}
				{
					p.SetState(191)

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
				p.SetState(194)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(195)

					var _m = p.Match(NparserOR)

					localctx.(*ExpresionContext)._OR = _m
				}
				{
					p.SetState(196)

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
		p.SetState(203)
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

	p.SetState(218)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserBOOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(204)
			p.Match(NparserBOOL)
		}
		localctx.(*Tipo_datoContext).ex = Ast.BOOLEAN

	case NparserCHAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(206)
			p.Match(NparserCHAR)
		}
		localctx.(*Tipo_datoContext).ex = Ast.CHAR

	case NparserI64:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(208)
			p.Match(NparserI64)
		}
		localctx.(*Tipo_datoContext).ex = Ast.I64

	case NparserF64:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(210)
			p.Match(NparserF64)
		}
		localctx.(*Tipo_datoContext).ex = Ast.F64

	case NparserSTR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(212)
			p.Match(NparserSTR)
		}
		localctx.(*Tipo_datoContext).ex = Ast.STR

	case NparserSTRING:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(214)
			p.Match(NparserSTRING)
		}
		localctx.(*Tipo_datoContext).ex = Ast.STRING

	case NparserUSIZE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(216)
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

	// Get_PAR_IZQ returns the _PAR_IZQ token.
	Get_PAR_IZQ() antlr.Token

	// Set_IF sets the _IF token.
	Set_IF(antlr.Token)

	// Set_PAR_IZQ sets the _PAR_IZQ token.
	Set_PAR_IZQ(antlr.Token)

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
	_PAR_IZQ        antlr.Token
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

func (s *Control_ifContext) Get_PAR_IZQ() antlr.Token { return s._PAR_IZQ }

func (s *Control_ifContext) Set_IF(v antlr.Token) { s._IF = v }

func (s *Control_ifContext) Set_PAR_IZQ(v antlr.Token) { s._PAR_IZQ = v }

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

func (s *Control_ifContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(NparserPAR_IZQ, 0)
}

func (s *Control_ifContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Control_ifContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(NparserPAR_DER, 0)
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

	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(220)

			var _m = p.Match(NparserIF)

			localctx.(*Control_ifContext)._IF = _m
		}
		{
			p.SetState(221)

			var _m = p.Match(NparserPAR_IZQ)

			localctx.(*Control_ifContext)._PAR_IZQ = _m
		}
		{
			p.SetState(222)

			var _x = p.expresion(0)

			localctx.(*Control_ifContext)._expresion = _x
		}
		{
			p.SetState(223)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(224)

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
			if localctx.(*Control_ifContext).Get_PAR_IZQ() == nil {
				return 0
			} else {
				return localctx.(*Control_ifContext).Get_PAR_IZQ().GetColumn()
			}
		}())
		columna++
		lista_null := arraylist.New()
		localctx.(*Control_ifContext).ex = exp_ins.NewIF(localctx.(*Control_ifContext).Get_expresion().GetEx(), localctx.(*Control_ifContext).GetBloqueIf().GetList(), lista_null, Ast.IF, fila, columna, false)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(227)

			var _m = p.Match(NparserIF)

			localctx.(*Control_ifContext)._IF = _m
		}
		{
			p.SetState(228)

			var _m = p.Match(NparserPAR_IZQ)

			localctx.(*Control_ifContext)._PAR_IZQ = _m
		}
		{
			p.SetState(229)

			var _x = p.expresion(0)

			localctx.(*Control_ifContext)._expresion = _x
		}
		{
			p.SetState(230)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(231)

			var _x = p.Bloque()

			localctx.(*Control_ifContext).bloqueIf = _x
		}
		{
			p.SetState(232)
			p.Match(NparserELSE)
		}
		{
			p.SetState(233)

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
			if localctx.(*Control_ifContext).Get_PAR_IZQ() == nil {
				return 0
			} else {
				return localctx.(*Control_ifContext).Get_PAR_IZQ().GetColumn()
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
			p.SetState(236)

			var _m = p.Match(NparserIF)

			localctx.(*Control_ifContext)._IF = _m
		}
		{
			p.SetState(237)

			var _m = p.Match(NparserPAR_IZQ)

			localctx.(*Control_ifContext)._PAR_IZQ = _m
		}
		{
			p.SetState(238)

			var _x = p.expresion(0)

			localctx.(*Control_ifContext)._expresion = _x
		}
		{
			p.SetState(239)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(240)

			var _x = p.Bloque()

			localctx.(*Control_ifContext).bloqueIf = _x
		}
		{
			p.SetState(241)

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
			if localctx.(*Control_ifContext).Get_PAR_IZQ() == nil {
				return 0
			} else {
				return localctx.(*Control_ifContext).Get_PAR_IZQ().GetColumn()
			}
		}())
		columna++
		lista_entonces := localctx.(*Control_ifContext).Get_bloque_else_if().GetList()
		localctx.(*Control_ifContext).ex = exp_ins.NewIF(localctx.(*Control_ifContext).Get_expresion().GetEx(), localctx.(*Control_ifContext).GetBloqueIf().GetList(), lista_entonces, Ast.IF, fila, columna, false)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(244)

			var _m = p.Match(NparserIF)

			localctx.(*Control_ifContext)._IF = _m
		}
		{
			p.SetState(245)

			var _m = p.Match(NparserPAR_IZQ)

			localctx.(*Control_ifContext)._PAR_IZQ = _m
		}
		{
			p.SetState(246)

			var _x = p.expresion(0)

			localctx.(*Control_ifContext)._expresion = _x
		}
		{
			p.SetState(247)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(248)

			var _x = p.Bloque()

			localctx.(*Control_ifContext).bloqueIf = _x
		}
		{
			p.SetState(249)

			var _x = p.Bloque_else_if()

			localctx.(*Control_ifContext)._bloque_else_if = _x
		}
		{
			p.SetState(250)
			p.Match(NparserELSE)
		}
		{
			p.SetState(251)

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
			if localctx.(*Control_ifContext).Get_PAR_IZQ() == nil {
				return 0
			} else {
				return localctx.(*Control_ifContext).Get_PAR_IZQ().GetColumn()
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
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(256)

				var _x = p.Else_if()

				localctx.(*Bloque_else_ifContext)._else_if = _x
			}
			localctx.(*Bloque_else_ifContext).lista = append(localctx.(*Bloque_else_ifContext).lista, localctx.(*Bloque_else_ifContext)._else_if)

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(259)
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

	// Get_PAR_IZQ returns the _PAR_IZQ token.
	Get_PAR_IZQ() antlr.Token

	// Set_ELSE sets the _ELSE token.
	Set_ELSE(antlr.Token)

	// Set_PAR_IZQ sets the _PAR_IZQ token.
	Set_PAR_IZQ(antlr.Token)

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
	_PAR_IZQ   antlr.Token
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

func (s *Else_ifContext) Get_PAR_IZQ() antlr.Token { return s._PAR_IZQ }

func (s *Else_ifContext) Set_ELSE(v antlr.Token) { s._ELSE = v }

func (s *Else_ifContext) Set_PAR_IZQ(v antlr.Token) { s._PAR_IZQ = v }

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

func (s *Else_ifContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(NparserPAR_IZQ, 0)
}

func (s *Else_ifContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Else_ifContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(NparserPAR_DER, 0)
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
		p.SetState(263)

		var _m = p.Match(NparserELSE)

		localctx.(*Else_ifContext)._ELSE = _m
	}
	{
		p.SetState(264)
		p.Match(NparserIF)
	}
	{
		p.SetState(265)

		var _m = p.Match(NparserPAR_IZQ)

		localctx.(*Else_ifContext)._PAR_IZQ = _m
	}
	{
		p.SetState(266)

		var _x = p.expresion(0)

		localctx.(*Else_ifContext)._expresion = _x
	}
	{
		p.SetState(267)
		p.Match(NparserPAR_DER)
	}
	{
		p.SetState(268)

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
		if localctx.(*Else_ifContext).Get_PAR_IZQ() == nil {
			return 0
		} else {
			return localctx.(*Else_ifContext).Get_PAR_IZQ().GetColumn()
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

	// Get_PAR_IZQ returns the _PAR_IZQ token.
	Get_PAR_IZQ() antlr.Token

	// Set_IF sets the _IF token.
	Set_IF(antlr.Token)

	// Set_PAR_IZQ sets the _PAR_IZQ token.
	Set_PAR_IZQ(antlr.Token)

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
	_PAR_IZQ            antlr.Token
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

func (s *Control_if_expContext) Get_PAR_IZQ() antlr.Token { return s._PAR_IZQ }

func (s *Control_if_expContext) Set_IF(v antlr.Token) { s._IF = v }

func (s *Control_if_expContext) Set_PAR_IZQ(v antlr.Token) { s._PAR_IZQ = v }

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

func (s *Control_if_expContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(NparserPAR_IZQ, 0)
}

func (s *Control_if_expContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Control_if_expContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(NparserPAR_DER, 0)
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

	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(271)

			var _m = p.Match(NparserIF)

			localctx.(*Control_if_expContext)._IF = _m
		}
		{
			p.SetState(272)

			var _m = p.Match(NparserPAR_IZQ)

			localctx.(*Control_if_expContext)._PAR_IZQ = _m
		}
		{
			p.SetState(273)

			var _x = p.expresion(0)

			localctx.(*Control_if_expContext)._expresion = _x
		}
		{
			p.SetState(274)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(275)

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
			if localctx.(*Control_if_expContext).Get_PAR_IZQ() == nil {
				return 0
			} else {
				return localctx.(*Control_if_expContext).Get_PAR_IZQ().GetColumn()
			}
		}())
		columna++
		lista_null := arraylist.New()
		localctx.(*Control_if_expContext).ex = exp_ins.NewIF(localctx.(*Control_if_expContext).Get_expresion().GetEx(), localctx.(*Control_if_expContext).GetBloqueIf().GetList(), lista_null, Ast.IF_EXPRESION, fila, columna, true)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(278)

			var _m = p.Match(NparserIF)

			localctx.(*Control_if_expContext)._IF = _m
		}
		{
			p.SetState(279)

			var _m = p.Match(NparserPAR_IZQ)

			localctx.(*Control_if_expContext)._PAR_IZQ = _m
		}
		{
			p.SetState(280)

			var _x = p.expresion(0)

			localctx.(*Control_if_expContext)._expresion = _x
		}
		{
			p.SetState(281)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(282)

			var _x = p.Bloque()

			localctx.(*Control_if_expContext).bloqueIf = _x
		}
		{
			p.SetState(283)
			p.Match(NparserELSE)
		}
		{
			p.SetState(284)

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
			if localctx.(*Control_if_expContext).Get_PAR_IZQ() == nil {
				return 0
			} else {
				return localctx.(*Control_if_expContext).Get_PAR_IZQ().GetColumn()
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
			p.SetState(287)

			var _m = p.Match(NparserIF)

			localctx.(*Control_if_expContext)._IF = _m
		}
		{
			p.SetState(288)

			var _m = p.Match(NparserPAR_IZQ)

			localctx.(*Control_if_expContext)._PAR_IZQ = _m
		}
		{
			p.SetState(289)

			var _x = p.expresion(0)

			localctx.(*Control_if_expContext)._expresion = _x
		}
		{
			p.SetState(290)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(291)

			var _x = p.Bloque()

			localctx.(*Control_if_expContext).bloqueIf = _x
		}
		{
			p.SetState(292)

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
			if localctx.(*Control_if_expContext).Get_PAR_IZQ() == nil {
				return 0
			} else {
				return localctx.(*Control_if_expContext).Get_PAR_IZQ().GetColumn()
			}
		}())
		columna++
		lista_entonces := localctx.(*Control_if_expContext).Get_bloque_else_if_exp().GetList()
		localctx.(*Control_if_expContext).ex = exp_ins.NewIF(localctx.(*Control_if_expContext).Get_expresion().GetEx(), localctx.(*Control_if_expContext).GetBloqueIf().GetList(), lista_entonces, Ast.IF_EXPRESION, fila, columna, true)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(295)

			var _m = p.Match(NparserIF)

			localctx.(*Control_if_expContext)._IF = _m
		}
		{
			p.SetState(296)

			var _m = p.Match(NparserPAR_IZQ)

			localctx.(*Control_if_expContext)._PAR_IZQ = _m
		}
		{
			p.SetState(297)

			var _x = p.expresion(0)

			localctx.(*Control_if_expContext)._expresion = _x
		}
		{
			p.SetState(298)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(299)

			var _x = p.Bloque()

			localctx.(*Control_if_expContext).bloqueIf = _x
		}
		{
			p.SetState(300)

			var _x = p.Bloque_else_if_exp()

			localctx.(*Control_if_expContext)._bloque_else_if_exp = _x
		}
		{
			p.SetState(301)
			p.Match(NparserELSE)
		}
		{
			p.SetState(302)

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
			if localctx.(*Control_if_expContext).Get_PAR_IZQ() == nil {
				return 0
			} else {
				return localctx.(*Control_if_expContext).Get_PAR_IZQ().GetColumn()
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
	p.SetState(308)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(307)

				var _x = p.Else_if_exp()

				localctx.(*Bloque_else_if_expContext)._else_if_exp = _x
			}
			localctx.(*Bloque_else_if_expContext).lista = append(localctx.(*Bloque_else_if_expContext).lista, localctx.(*Bloque_else_if_expContext)._else_if_exp)

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(310)
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

	// Get_PAR_IZQ returns the _PAR_IZQ token.
	Get_PAR_IZQ() antlr.Token

	// Set_ELSE sets the _ELSE token.
	Set_ELSE(antlr.Token)

	// Set_PAR_IZQ sets the _PAR_IZQ token.
	Set_PAR_IZQ(antlr.Token)

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
	_PAR_IZQ   antlr.Token
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

func (s *Else_if_expContext) Get_PAR_IZQ() antlr.Token { return s._PAR_IZQ }

func (s *Else_if_expContext) Set_ELSE(v antlr.Token) { s._ELSE = v }

func (s *Else_if_expContext) Set_PAR_IZQ(v antlr.Token) { s._PAR_IZQ = v }

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

func (s *Else_if_expContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(NparserPAR_IZQ, 0)
}

func (s *Else_if_expContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Else_if_expContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(NparserPAR_DER, 0)
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
		p.SetState(314)

		var _m = p.Match(NparserELSE)

		localctx.(*Else_if_expContext)._ELSE = _m
	}
	{
		p.SetState(315)
		p.Match(NparserIF)
	}
	{
		p.SetState(316)

		var _m = p.Match(NparserPAR_IZQ)

		localctx.(*Else_if_expContext)._PAR_IZQ = _m
	}
	{
		p.SetState(317)

		var _x = p.expresion(0)

		localctx.(*Else_if_expContext)._expresion = _x
	}
	{
		p.SetState(318)
		p.Match(NparserPAR_DER)
	}
	{
		p.SetState(319)

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
		if localctx.(*Else_if_expContext).Get_PAR_IZQ() == nil {
			return 0
		} else {
			return localctx.(*Else_if_expContext).Get_PAR_IZQ().GetColumn()
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

	// Set_control_if_exp sets the _control_if_exp rule contexts.
	Set_control_if_exp(IControl_if_expContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsControl_expresionContext differentiates from other interfaces.
	IsControl_expresionContext()
}

type Control_expresionContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	ex              Ast.Instruccion
	_control_if_exp IControl_if_expContext
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

func (s *Control_expresionContext) Set_control_if_exp(v IControl_if_expContext) {
	s._control_if_exp = v
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(322)

		var _x = p.Control_if_exp()

		localctx.(*Control_expresionContext)._control_if_exp = _x
	}
	localctx.(*Control_expresionContext).ex = localctx.(*Control_expresionContext).Get_control_if_exp().GetEx()

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
