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
import "Back/analizador/Ast/simbolos"
import "Back/analizador/fn_primitivas"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 91, 803,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 7, 4, 92, 10, 4, 12, 4, 14, 4, 95, 11, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 5, 5, 149, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 276, 10,
	6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 5, 7, 350, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 5, 8, 362, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 401, 10, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 451, 10, 9, 12,
	9, 14, 9, 454, 11, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 470, 10, 10, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 498, 10, 11, 3, 12, 6, 12, 501,
	10, 12, 13, 12, 14, 12, 502, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 539, 10, 14, 3,
	15, 6, 15, 542, 10, 15, 13, 15, 14, 15, 543, 3, 15, 3, 15, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 5, 17, 563, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 19, 6, 19, 573, 10, 19, 13, 19, 14, 19, 574, 3, 19, 3,
	19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 5, 21, 591, 10, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 7,
	21, 598, 10, 21, 12, 21, 14, 21, 601, 11, 21, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 23, 6, 23, 611, 10, 23, 13, 23, 14, 23, 612, 3,
	23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 25, 5, 25, 629, 10, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	25, 7, 25, 636, 10, 25, 12, 25, 14, 25, 639, 11, 25, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 5, 26, 647, 10, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 27, 5, 27, 655, 10, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 7, 33, 691, 10, 33, 12,
	33, 14, 33, 694, 11, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 7, 35, 710, 10, 35, 12,
	35, 14, 35, 713, 11, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36,
	3, 36, 3, 36, 3, 36, 3, 36, 5, 36, 726, 10, 36, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 5, 37, 738, 10, 37, 3, 38,
	3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 7, 38, 749, 10,
	38, 12, 38, 14, 38, 752, 11, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3,
	39, 3, 39, 3, 39, 3, 39, 3, 39, 5, 39, 764, 10, 39, 3, 40, 3, 40, 3, 40,
	3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 7, 40, 775, 10, 40, 12, 40, 14,
	40, 778, 11, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41,
	3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3,
	41, 3, 41, 3, 41, 5, 41, 801, 10, 41, 3, 41, 2, 9, 16, 40, 48, 64, 68,
	74, 78, 42, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
	34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68,
	70, 72, 74, 76, 78, 80, 2, 7, 4, 2, 76, 76, 78, 78, 3, 2, 72, 74, 3, 2,
	76, 77, 4, 2, 64, 68, 70, 70, 4, 2, 68, 68, 70, 70, 2, 852, 2, 82, 3, 2,
	2, 2, 4, 85, 3, 2, 2, 2, 6, 93, 3, 2, 2, 2, 8, 148, 3, 2, 2, 2, 10, 275,
	3, 2, 2, 2, 12, 349, 3, 2, 2, 2, 14, 361, 3, 2, 2, 2, 16, 400, 3, 2, 2,
	2, 18, 469, 3, 2, 2, 2, 20, 497, 3, 2, 2, 2, 22, 500, 3, 2, 2, 2, 24, 506,
	3, 2, 2, 2, 26, 538, 3, 2, 2, 2, 28, 541, 3, 2, 2, 2, 30, 547, 3, 2, 2,
	2, 32, 562, 3, 2, 2, 2, 34, 564, 3, 2, 2, 2, 36, 572, 3, 2, 2, 2, 38, 578,
	3, 2, 2, 2, 40, 590, 3, 2, 2, 2, 42, 602, 3, 2, 2, 2, 44, 610, 3, 2, 2,
	2, 46, 616, 3, 2, 2, 2, 48, 628, 3, 2, 2, 2, 50, 646, 3, 2, 2, 2, 52, 654,
	3, 2, 2, 2, 54, 656, 3, 2, 2, 2, 56, 659, 3, 2, 2, 2, 58, 663, 3, 2, 2,
	2, 60, 667, 3, 2, 2, 2, 62, 673, 3, 2, 2, 2, 64, 681, 3, 2, 2, 2, 66, 695,
	3, 2, 2, 2, 68, 700, 3, 2, 2, 2, 70, 725, 3, 2, 2, 2, 72, 737, 3, 2, 2,
	2, 74, 739, 3, 2, 2, 2, 76, 763, 3, 2, 2, 2, 78, 765, 3, 2, 2, 2, 80, 800,
	3, 2, 2, 2, 82, 83, 5, 6, 4, 2, 83, 84, 8, 2, 1, 2, 84, 3, 3, 2, 2, 2,
	85, 86, 7, 82, 2, 2, 86, 87, 5, 6, 4, 2, 87, 88, 7, 83, 2, 2, 88, 89, 8,
	3, 1, 2, 89, 5, 3, 2, 2, 2, 90, 92, 5, 8, 5, 2, 91, 90, 3, 2, 2, 2, 92,
	95, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 96, 3, 2, 2,
	2, 95, 93, 3, 2, 2, 2, 96, 97, 8, 4, 1, 2, 97, 7, 3, 2, 2, 2, 98, 99, 5,
	72, 37, 2, 99, 100, 7, 63, 2, 2, 100, 101, 8, 5, 1, 2, 101, 149, 3, 2,
	2, 2, 102, 103, 5, 16, 9, 2, 103, 104, 8, 5, 1, 2, 104, 149, 3, 2, 2, 2,
	105, 106, 5, 10, 6, 2, 106, 107, 7, 63, 2, 2, 107, 108, 8, 5, 1, 2, 108,
	149, 3, 2, 2, 2, 109, 110, 5, 12, 7, 2, 110, 111, 8, 5, 1, 2, 111, 149,
	3, 2, 2, 2, 112, 113, 5, 14, 8, 2, 113, 114, 7, 63, 2, 2, 114, 115, 8,
	5, 1, 2, 115, 149, 3, 2, 2, 2, 116, 117, 5, 20, 11, 2, 117, 118, 8, 5,
	1, 2, 118, 149, 3, 2, 2, 2, 119, 120, 5, 34, 18, 2, 120, 121, 8, 5, 1,
	2, 121, 149, 3, 2, 2, 2, 122, 123, 5, 56, 29, 2, 123, 124, 8, 5, 1, 2,
	124, 149, 3, 2, 2, 2, 125, 126, 5, 66, 34, 2, 126, 127, 8, 5, 1, 2, 127,
	149, 3, 2, 2, 2, 128, 129, 5, 52, 27, 2, 129, 130, 7, 63, 2, 2, 130, 131,
	8, 5, 1, 2, 131, 149, 3, 2, 2, 2, 132, 133, 5, 54, 28, 2, 133, 134, 7,
	63, 2, 2, 134, 135, 8, 5, 1, 2, 135, 149, 3, 2, 2, 2, 136, 137, 5, 50,
	26, 2, 137, 138, 7, 63, 2, 2, 138, 139, 8, 5, 1, 2, 139, 149, 3, 2, 2,
	2, 140, 141, 5, 60, 31, 2, 141, 142, 7, 63, 2, 2, 142, 143, 8, 5, 1, 2,
	143, 149, 3, 2, 2, 2, 144, 145, 5, 62, 32, 2, 145, 146, 7, 63, 2, 2, 146,
	147, 8, 5, 1, 2, 147, 149, 3, 2, 2, 2, 148, 98, 3, 2, 2, 2, 148, 102, 3,
	2, 2, 2, 148, 105, 3, 2, 2, 2, 148, 109, 3, 2, 2, 2, 148, 112, 3, 2, 2,
	2, 148, 116, 3, 2, 2, 2, 148, 119, 3, 2, 2, 2, 148, 122, 3, 2, 2, 2, 148,
	125, 3, 2, 2, 2, 148, 128, 3, 2, 2, 2, 148, 132, 3, 2, 2, 2, 148, 136,
	3, 2, 2, 2, 148, 140, 3, 2, 2, 2, 148, 144, 3, 2, 2, 2, 149, 9, 3, 2, 2,
	2, 150, 151, 7, 17, 2, 2, 151, 152, 7, 51, 2, 2, 152, 153, 7, 71, 2, 2,
	153, 154, 5, 16, 9, 2, 154, 155, 8, 6, 1, 2, 155, 276, 3, 2, 2, 2, 156,
	157, 7, 17, 2, 2, 157, 158, 7, 51, 2, 2, 158, 159, 7, 71, 2, 2, 159, 160,
	5, 32, 17, 2, 160, 161, 8, 6, 1, 2, 161, 276, 3, 2, 2, 2, 162, 163, 7,
	17, 2, 2, 163, 164, 7, 51, 2, 2, 164, 165, 7, 59, 2, 2, 165, 166, 5, 18,
	10, 2, 166, 167, 7, 71, 2, 2, 167, 168, 5, 16, 9, 2, 168, 169, 8, 6, 1,
	2, 169, 276, 3, 2, 2, 2, 170, 171, 7, 17, 2, 2, 171, 172, 7, 51, 2, 2,
	172, 173, 7, 59, 2, 2, 173, 174, 5, 18, 10, 2, 174, 175, 7, 71, 2, 2, 175,
	176, 5, 32, 17, 2, 176, 177, 8, 6, 1, 2, 177, 276, 3, 2, 2, 2, 178, 179,
	7, 17, 2, 2, 179, 180, 7, 16, 2, 2, 180, 181, 7, 51, 2, 2, 181, 182, 7,
	71, 2, 2, 182, 183, 5, 16, 9, 2, 183, 184, 8, 6, 1, 2, 184, 276, 3, 2,
	2, 2, 185, 186, 7, 17, 2, 2, 186, 187, 7, 16, 2, 2, 187, 188, 7, 51, 2,
	2, 188, 189, 7, 71, 2, 2, 189, 190, 5, 32, 17, 2, 190, 191, 8, 6, 1, 2,
	191, 276, 3, 2, 2, 2, 192, 193, 7, 17, 2, 2, 193, 194, 7, 16, 2, 2, 194,
	195, 7, 51, 2, 2, 195, 196, 7, 59, 2, 2, 196, 197, 5, 18, 10, 2, 197, 198,
	8, 6, 1, 2, 198, 276, 3, 2, 2, 2, 199, 200, 7, 17, 2, 2, 200, 201, 7, 16,
	2, 2, 201, 202, 7, 51, 2, 2, 202, 203, 7, 59, 2, 2, 203, 204, 5, 18, 10,
	2, 204, 205, 7, 71, 2, 2, 205, 206, 5, 16, 9, 2, 206, 207, 8, 6, 1, 2,
	207, 276, 3, 2, 2, 2, 208, 209, 7, 17, 2, 2, 209, 210, 7, 16, 2, 2, 210,
	211, 7, 51, 2, 2, 211, 212, 7, 59, 2, 2, 212, 213, 5, 18, 10, 2, 213, 214,
	7, 71, 2, 2, 214, 215, 5, 32, 17, 2, 215, 216, 8, 6, 1, 2, 216, 276, 3,
	2, 2, 2, 217, 218, 7, 17, 2, 2, 218, 219, 7, 51, 2, 2, 219, 220, 7, 71,
	2, 2, 220, 221, 5, 16, 9, 2, 221, 222, 8, 6, 1, 2, 222, 276, 3, 2, 2, 2,
	223, 224, 7, 17, 2, 2, 224, 225, 7, 16, 2, 2, 225, 226, 7, 51, 2, 2, 226,
	227, 7, 71, 2, 2, 227, 228, 5, 16, 9, 2, 228, 229, 8, 6, 1, 2, 229, 276,
	3, 2, 2, 2, 230, 231, 7, 17, 2, 2, 231, 232, 7, 51, 2, 2, 232, 233, 7,
	59, 2, 2, 233, 234, 7, 14, 2, 2, 234, 235, 7, 67, 2, 2, 235, 236, 5, 18,
	10, 2, 236, 237, 7, 65, 2, 2, 237, 238, 7, 71, 2, 2, 238, 239, 5, 16, 9,
	2, 239, 240, 8, 6, 1, 2, 240, 276, 3, 2, 2, 2, 241, 242, 7, 17, 2, 2, 242,
	243, 7, 16, 2, 2, 243, 244, 7, 51, 2, 2, 244, 245, 7, 59, 2, 2, 245, 246,
	7, 14, 2, 2, 246, 247, 7, 67, 2, 2, 247, 248, 5, 18, 10, 2, 248, 249, 7,
	65, 2, 2, 249, 250, 7, 71, 2, 2, 250, 251, 5, 16, 9, 2, 251, 252, 8, 6,
	1, 2, 252, 276, 3, 2, 2, 2, 253, 254, 7, 17, 2, 2, 254, 255, 7, 16, 2,
	2, 255, 256, 7, 51, 2, 2, 256, 257, 7, 59, 2, 2, 257, 258, 7, 14, 2, 2,
	258, 259, 7, 67, 2, 2, 259, 260, 5, 18, 10, 2, 260, 261, 7, 65, 2, 2, 261,
	262, 7, 71, 2, 2, 262, 263, 5, 16, 9, 2, 263, 264, 8, 6, 1, 2, 264, 276,
	3, 2, 2, 2, 265, 266, 7, 17, 2, 2, 266, 267, 7, 51, 2, 2, 267, 268, 7,
	59, 2, 2, 268, 269, 7, 14, 2, 2, 269, 270, 7, 67, 2, 2, 270, 271, 5, 18,
	10, 2, 271, 272, 7, 65, 2, 2, 272, 273, 7, 71, 2, 2, 273, 274, 8, 6, 1,
	2, 274, 276, 3, 2, 2, 2, 275, 150, 3, 2, 2, 2, 275, 156, 3, 2, 2, 2, 275,
	162, 3, 2, 2, 2, 275, 170, 3, 2, 2, 2, 275, 178, 3, 2, 2, 2, 275, 185,
	3, 2, 2, 2, 275, 192, 3, 2, 2, 2, 275, 199, 3, 2, 2, 2, 275, 208, 3, 2,
	2, 2, 275, 217, 3, 2, 2, 2, 275, 223, 3, 2, 2, 2, 275, 230, 3, 2, 2, 2,
	275, 241, 3, 2, 2, 2, 275, 253, 3, 2, 2, 2, 275, 265, 3, 2, 2, 2, 276,
	11, 3, 2, 2, 2, 277, 278, 7, 47, 2, 2, 278, 279, 7, 24, 2, 2, 279, 280,
	7, 51, 2, 2, 280, 281, 7, 80, 2, 2, 281, 282, 7, 81, 2, 2, 282, 283, 7,
	75, 2, 2, 283, 284, 5, 18, 10, 2, 284, 285, 5, 4, 3, 2, 285, 286, 8, 7,
	1, 2, 286, 350, 3, 2, 2, 2, 287, 288, 7, 47, 2, 2, 288, 289, 7, 24, 2,
	2, 289, 290, 7, 51, 2, 2, 290, 291, 7, 80, 2, 2, 291, 292, 7, 81, 2, 2,
	292, 293, 5, 4, 3, 2, 293, 294, 8, 7, 1, 2, 294, 350, 3, 2, 2, 2, 295,
	296, 7, 24, 2, 2, 296, 297, 7, 51, 2, 2, 297, 298, 7, 80, 2, 2, 298, 299,
	7, 81, 2, 2, 299, 300, 7, 75, 2, 2, 300, 301, 5, 18, 10, 2, 301, 302, 5,
	4, 3, 2, 302, 303, 8, 7, 1, 2, 303, 350, 3, 2, 2, 2, 304, 305, 7, 24, 2,
	2, 305, 306, 7, 51, 2, 2, 306, 307, 7, 80, 2, 2, 307, 308, 7, 81, 2, 2,
	308, 309, 5, 4, 3, 2, 309, 310, 8, 7, 1, 2, 310, 350, 3, 2, 2, 2, 311,
	312, 7, 47, 2, 2, 312, 313, 7, 24, 2, 2, 313, 314, 7, 51, 2, 2, 314, 315,
	7, 80, 2, 2, 315, 316, 5, 68, 35, 2, 316, 317, 7, 81, 2, 2, 317, 318, 7,
	75, 2, 2, 318, 319, 5, 18, 10, 2, 319, 320, 5, 4, 3, 2, 320, 321, 8, 7,
	1, 2, 321, 350, 3, 2, 2, 2, 322, 323, 7, 47, 2, 2, 323, 324, 7, 24, 2,
	2, 324, 325, 7, 51, 2, 2, 325, 326, 7, 80, 2, 2, 326, 327, 5, 68, 35, 2,
	327, 328, 7, 81, 2, 2, 328, 329, 5, 4, 3, 2, 329, 330, 8, 7, 1, 2, 330,
	350, 3, 2, 2, 2, 331, 332, 7, 24, 2, 2, 332, 333, 7, 51, 2, 2, 333, 334,
	7, 80, 2, 2, 334, 335, 5, 68, 35, 2, 335, 336, 7, 81, 2, 2, 336, 337, 7,
	75, 2, 2, 337, 338, 5, 18, 10, 2, 338, 339, 5, 4, 3, 2, 339, 340, 8, 7,
	1, 2, 340, 350, 3, 2, 2, 2, 341, 342, 7, 24, 2, 2, 342, 343, 7, 51, 2,
	2, 343, 344, 7, 80, 2, 2, 344, 345, 5, 68, 35, 2, 345, 346, 7, 81, 2, 2,
	346, 347, 5, 4, 3, 2, 347, 348, 8, 7, 1, 2, 348, 350, 3, 2, 2, 2, 349,
	277, 3, 2, 2, 2, 349, 287, 3, 2, 2, 2, 349, 295, 3, 2, 2, 2, 349, 304,
	3, 2, 2, 2, 349, 311, 3, 2, 2, 2, 349, 322, 3, 2, 2, 2, 349, 331, 3, 2,
	2, 2, 349, 341, 3, 2, 2, 2, 350, 13, 3, 2, 2, 2, 351, 352, 7, 51, 2, 2,
	352, 353, 7, 71, 2, 2, 353, 354, 5, 16, 9, 2, 354, 355, 8, 8, 1, 2, 355,
	362, 3, 2, 2, 2, 356, 357, 7, 51, 2, 2, 357, 358, 7, 71, 2, 2, 358, 359,
	5, 32, 17, 2, 359, 360, 8, 8, 1, 2, 360, 362, 3, 2, 2, 2, 361, 351, 3,
	2, 2, 2, 361, 356, 3, 2, 2, 2, 362, 15, 3, 2, 2, 2, 363, 364, 8, 9, 1,
	2, 364, 365, 9, 2, 2, 2, 365, 366, 5, 16, 9, 23, 366, 367, 8, 9, 1, 2,
	367, 401, 3, 2, 2, 2, 368, 369, 7, 80, 2, 2, 369, 370, 5, 16, 9, 2, 370,
	371, 7, 81, 2, 2, 371, 372, 8, 9, 1, 2, 372, 401, 3, 2, 2, 2, 373, 374,
	7, 80, 2, 2, 374, 375, 5, 16, 9, 2, 375, 376, 7, 13, 2, 2, 376, 377, 5,
	18, 10, 2, 377, 378, 7, 81, 2, 2, 378, 379, 8, 9, 1, 2, 379, 401, 3, 2,
	2, 2, 380, 381, 5, 72, 37, 2, 381, 382, 8, 9, 1, 2, 382, 401, 3, 2, 2,
	2, 383, 384, 5, 80, 41, 2, 384, 385, 8, 9, 1, 2, 385, 401, 3, 2, 2, 2,
	386, 387, 7, 51, 2, 2, 387, 401, 8, 9, 1, 2, 388, 389, 7, 21, 2, 2, 389,
	401, 8, 9, 1, 2, 390, 391, 7, 22, 2, 2, 391, 401, 8, 9, 1, 2, 392, 393,
	7, 88, 2, 2, 393, 401, 8, 9, 1, 2, 394, 395, 7, 49, 2, 2, 395, 401, 8,
	9, 1, 2, 396, 397, 7, 48, 2, 2, 397, 401, 8, 9, 1, 2, 398, 399, 7, 86,
	2, 2, 399, 401, 8, 9, 1, 2, 400, 363, 3, 2, 2, 2, 400, 368, 3, 2, 2, 2,
	400, 373, 3, 2, 2, 2, 400, 380, 3, 2, 2, 2, 400, 383, 3, 2, 2, 2, 400,
	386, 3, 2, 2, 2, 400, 388, 3, 2, 2, 2, 400, 390, 3, 2, 2, 2, 400, 392,
	3, 2, 2, 2, 400, 394, 3, 2, 2, 2, 400, 396, 3, 2, 2, 2, 400, 398, 3, 2,
	2, 2, 401, 452, 3, 2, 2, 2, 402, 403, 12, 19, 2, 2, 403, 404, 9, 3, 2,
	2, 404, 405, 5, 16, 9, 20, 405, 406, 8, 9, 1, 2, 406, 451, 3, 2, 2, 2,
	407, 408, 12, 18, 2, 2, 408, 409, 9, 4, 2, 2, 409, 410, 5, 16, 9, 19, 410,
	411, 8, 9, 1, 2, 411, 451, 3, 2, 2, 2, 412, 413, 12, 17, 2, 2, 413, 414,
	9, 5, 2, 2, 414, 415, 5, 16, 9, 18, 415, 416, 8, 9, 1, 2, 416, 451, 3,
	2, 2, 2, 417, 418, 12, 16, 2, 2, 418, 419, 9, 6, 2, 2, 419, 420, 5, 16,
	9, 17, 420, 421, 8, 9, 1, 2, 421, 451, 3, 2, 2, 2, 422, 423, 12, 15, 2,
	2, 423, 424, 7, 56, 2, 2, 424, 425, 5, 16, 9, 16, 425, 426, 8, 9, 1, 2,
	426, 451, 3, 2, 2, 2, 427, 428, 12, 14, 2, 2, 428, 429, 7, 54, 2, 2, 429,
	430, 5, 16, 9, 15, 430, 431, 8, 9, 1, 2, 431, 451, 3, 2, 2, 2, 432, 433,
	12, 22, 2, 2, 433, 434, 7, 61, 2, 2, 434, 435, 7, 19, 2, 2, 435, 436, 7,
	80, 2, 2, 436, 437, 7, 81, 2, 2, 437, 451, 8, 9, 1, 2, 438, 439, 12, 21,
	2, 2, 439, 440, 7, 61, 2, 2, 440, 441, 7, 26, 2, 2, 441, 442, 7, 80, 2,
	2, 442, 443, 7, 81, 2, 2, 443, 451, 8, 9, 1, 2, 444, 445, 12, 20, 2, 2,
	445, 446, 7, 61, 2, 2, 446, 447, 7, 25, 2, 2, 447, 448, 7, 80, 2, 2, 448,
	449, 7, 81, 2, 2, 449, 451, 8, 9, 1, 2, 450, 402, 3, 2, 2, 2, 450, 407,
	3, 2, 2, 2, 450, 412, 3, 2, 2, 2, 450, 417, 3, 2, 2, 2, 450, 422, 3, 2,
	2, 2, 450, 427, 3, 2, 2, 2, 450, 432, 3, 2, 2, 2, 450, 438, 3, 2, 2, 2,
	450, 444, 3, 2, 2, 2, 451, 454, 3, 2, 2, 2, 452, 450, 3, 2, 2, 2, 452,
	453, 3, 2, 2, 2, 453, 17, 3, 2, 2, 2, 454, 452, 3, 2, 2, 2, 455, 456, 7,
	3, 2, 2, 456, 470, 8, 10, 1, 2, 457, 458, 7, 4, 2, 2, 458, 470, 8, 10,
	1, 2, 459, 460, 7, 6, 2, 2, 460, 470, 8, 10, 1, 2, 461, 462, 7, 5, 2, 2,
	462, 470, 8, 10, 1, 2, 463, 464, 7, 7, 2, 2, 464, 470, 8, 10, 1, 2, 465,
	466, 7, 8, 2, 2, 466, 470, 8, 10, 1, 2, 467, 468, 7, 9, 2, 2, 468, 470,
	8, 10, 1, 2, 469, 455, 3, 2, 2, 2, 469, 457, 3, 2, 2, 2, 469, 459, 3, 2,
	2, 2, 469, 461, 3, 2, 2, 2, 469, 463, 3, 2, 2, 2, 469, 465, 3, 2, 2, 2,
	469, 467, 3, 2, 2, 2, 470, 19, 3, 2, 2, 2, 471, 472, 7, 36, 2, 2, 472,
	473, 5, 16, 9, 2, 473, 474, 5, 4, 3, 2, 474, 475, 8, 11, 1, 2, 475, 498,
	3, 2, 2, 2, 476, 477, 7, 36, 2, 2, 477, 478, 5, 16, 9, 2, 478, 479, 5,
	4, 3, 2, 479, 480, 7, 37, 2, 2, 480, 481, 5, 4, 3, 2, 481, 482, 8, 11,
	1, 2, 482, 498, 3, 2, 2, 2, 483, 484, 7, 36, 2, 2, 484, 485, 5, 16, 9,
	2, 485, 486, 5, 4, 3, 2, 486, 487, 5, 22, 12, 2, 487, 488, 8, 11, 1, 2,
	488, 498, 3, 2, 2, 2, 489, 490, 7, 36, 2, 2, 490, 491, 5, 16, 9, 2, 491,
	492, 5, 4, 3, 2, 492, 493, 5, 22, 12, 2, 493, 494, 7, 37, 2, 2, 494, 495,
	5, 4, 3, 2, 495, 496, 8, 11, 1, 2, 496, 498, 3, 2, 2, 2, 497, 471, 3, 2,
	2, 2, 497, 476, 3, 2, 2, 2, 497, 483, 3, 2, 2, 2, 497, 489, 3, 2, 2, 2,
	498, 21, 3, 2, 2, 2, 499, 501, 5, 24, 13, 2, 500, 499, 3, 2, 2, 2, 501,
	502, 3, 2, 2, 2, 502, 500, 3, 2, 2, 2, 502, 503, 3, 2, 2, 2, 503, 504,
	3, 2, 2, 2, 504, 505, 8, 12, 1, 2, 505, 23, 3, 2, 2, 2, 506, 507, 7, 37,
	2, 2, 507, 508, 7, 36, 2, 2, 508, 509, 5, 16, 9, 2, 509, 510, 5, 4, 3,
	2, 510, 511, 8, 13, 1, 2, 511, 25, 3, 2, 2, 2, 512, 513, 7, 36, 2, 2, 513,
	514, 5, 16, 9, 2, 514, 515, 5, 4, 3, 2, 515, 516, 8, 14, 1, 2, 516, 539,
	3, 2, 2, 2, 517, 518, 7, 36, 2, 2, 518, 519, 5, 16, 9, 2, 519, 520, 5,
	4, 3, 2, 520, 521, 7, 37, 2, 2, 521, 522, 5, 4, 3, 2, 522, 523, 8, 14,
	1, 2, 523, 539, 3, 2, 2, 2, 524, 525, 7, 36, 2, 2, 525, 526, 5, 16, 9,
	2, 526, 527, 5, 4, 3, 2, 527, 528, 5, 28, 15, 2, 528, 529, 8, 14, 1, 2,
	529, 539, 3, 2, 2, 2, 530, 531, 7, 36, 2, 2, 531, 532, 5, 16, 9, 2, 532,
	533, 5, 4, 3, 2, 533, 534, 5, 28, 15, 2, 534, 535, 7, 37, 2, 2, 535, 536,
	5, 4, 3, 2, 536, 537, 8, 14, 1, 2, 537, 539, 3, 2, 2, 2, 538, 512, 3, 2,
	2, 2, 538, 517, 3, 2, 2, 2, 538, 524, 3, 2, 2, 2, 538, 530, 3, 2, 2, 2,
	539, 27, 3, 2, 2, 2, 540, 542, 5, 30, 16, 2, 541, 540, 3, 2, 2, 2, 542,
	543, 3, 2, 2, 2, 543, 541, 3, 2, 2, 2, 543, 544, 3, 2, 2, 2, 544, 545,
	3, 2, 2, 2, 545, 546, 8, 15, 1, 2, 546, 29, 3, 2, 2, 2, 547, 548, 7, 37,
	2, 2, 548, 549, 7, 36, 2, 2, 549, 550, 5, 16, 9, 2, 550, 551, 5, 4, 3,
	2, 551, 552, 8, 16, 1, 2, 552, 31, 3, 2, 2, 2, 553, 554, 5, 26, 14, 2,
	554, 555, 8, 17, 1, 2, 555, 563, 3, 2, 2, 2, 556, 557, 5, 42, 22, 2, 557,
	558, 8, 17, 1, 2, 558, 563, 3, 2, 2, 2, 559, 560, 5, 58, 30, 2, 560, 561,
	8, 17, 1, 2, 561, 563, 3, 2, 2, 2, 562, 553, 3, 2, 2, 2, 562, 556, 3, 2,
	2, 2, 562, 559, 3, 2, 2, 2, 563, 33, 3, 2, 2, 2, 564, 565, 7, 38, 2, 2,
	565, 566, 5, 16, 9, 2, 566, 567, 7, 82, 2, 2, 567, 568, 5, 36, 19, 2, 568,
	569, 7, 83, 2, 2, 569, 570, 8, 18, 1, 2, 570, 35, 3, 2, 2, 2, 571, 573,
	5, 38, 20, 2, 572, 571, 3, 2, 2, 2, 573, 574, 3, 2, 2, 2, 574, 572, 3,
	2, 2, 2, 574, 575, 3, 2, 2, 2, 575, 576, 3, 2, 2, 2, 576, 577, 8, 19, 1,
	2, 577, 37, 3, 2, 2, 2, 578, 579, 5, 40, 21, 2, 579, 580, 7, 69, 2, 2,
	580, 581, 5, 4, 3, 2, 581, 582, 7, 62, 2, 2, 582, 583, 8, 20, 1, 2, 583,
	39, 3, 2, 2, 2, 584, 585, 8, 21, 1, 2, 585, 586, 5, 16, 9, 2, 586, 587,
	8, 21, 1, 2, 587, 591, 3, 2, 2, 2, 588, 589, 7, 52, 2, 2, 589, 591, 8,
	21, 1, 2, 590, 584, 3, 2, 2, 2, 590, 588, 3, 2, 2, 2, 591, 599, 3, 2, 2,
	2, 592, 593, 12, 5, 2, 2, 593, 594, 7, 53, 2, 2, 594, 595, 5, 16, 9, 2,
	595, 596, 8, 21, 1, 2, 596, 598, 3, 2, 2, 2, 597, 592, 3, 2, 2, 2, 598,
	601, 3, 2, 2, 2, 599, 597, 3, 2, 2, 2, 599, 600, 3, 2, 2, 2, 600, 41, 3,
	2, 2, 2, 601, 599, 3, 2, 2, 2, 602, 603, 7, 38, 2, 2, 603, 604, 5, 16,
	9, 2, 604, 605, 7, 82, 2, 2, 605, 606, 5, 44, 23, 2, 606, 607, 7, 83, 2,
	2, 607, 608, 8, 22, 1, 2, 608, 43, 3, 2, 2, 2, 609, 611, 5, 46, 24, 2,
	610, 609, 3, 2, 2, 2, 611, 612, 3, 2, 2, 2, 612, 610, 3, 2, 2, 2, 612,
	613, 3, 2, 2, 2, 613, 614, 3, 2, 2, 2, 614, 615, 8, 23, 1, 2, 615, 45,
	3, 2, 2, 2, 616, 617, 5, 48, 25, 2, 617, 618, 7, 69, 2, 2, 618, 619, 5,
	4, 3, 2, 619, 620, 7, 62, 2, 2, 620, 621, 8, 24, 1, 2, 621, 47, 3, 2, 2,
	2, 622, 623, 8, 25, 1, 2, 623, 624, 5, 16, 9, 2, 624, 625, 8, 25, 1, 2,
	625, 629, 3, 2, 2, 2, 626, 627, 7, 52, 2, 2, 627, 629, 8, 25, 1, 2, 628,
	622, 3, 2, 2, 2, 628, 626, 3, 2, 2, 2, 629, 637, 3, 2, 2, 2, 630, 631,
	12, 5, 2, 2, 631, 632, 7, 53, 2, 2, 632, 633, 5, 16, 9, 2, 633, 634, 8,
	25, 1, 2, 634, 636, 3, 2, 2, 2, 635, 630, 3, 2, 2, 2, 636, 639, 3, 2, 2,
	2, 637, 635, 3, 2, 2, 2, 637, 638, 3, 2, 2, 2, 638, 49, 3, 2, 2, 2, 639,
	637, 3, 2, 2, 2, 640, 641, 7, 43, 2, 2, 641, 647, 8, 26, 1, 2, 642, 643,
	7, 43, 2, 2, 643, 644, 5, 16, 9, 2, 644, 645, 8, 26, 1, 2, 645, 647, 3,
	2, 2, 2, 646, 640, 3, 2, 2, 2, 646, 642, 3, 2, 2, 2, 647, 51, 3, 2, 2,
	2, 648, 649, 7, 44, 2, 2, 649, 655, 8, 27, 1, 2, 650, 651, 7, 44, 2, 2,
	651, 652, 5, 16, 9, 2, 652, 653, 8, 27, 1, 2, 653, 655, 3, 2, 2, 2, 654,
	648, 3, 2, 2, 2, 654, 650, 3, 2, 2, 2, 655, 53, 3, 2, 2, 2, 656, 657, 7,
	45, 2, 2, 657, 658, 8, 28, 1, 2, 658, 55, 3, 2, 2, 2, 659, 660, 7, 39,
	2, 2, 660, 661, 5, 4, 3, 2, 661, 662, 8, 29, 1, 2, 662, 57, 3, 2, 2, 2,
	663, 664, 7, 39, 2, 2, 664, 665, 5, 4, 3, 2, 665, 666, 8, 30, 1, 2, 666,
	59, 3, 2, 2, 2, 667, 668, 7, 23, 2, 2, 668, 669, 7, 80, 2, 2, 669, 670,
	5, 16, 9, 2, 670, 671, 7, 81, 2, 2, 671, 672, 8, 31, 1, 2, 672, 61, 3,
	2, 2, 2, 673, 674, 7, 23, 2, 2, 674, 675, 7, 80, 2, 2, 675, 676, 7, 86,
	2, 2, 676, 677, 7, 62, 2, 2, 677, 678, 5, 64, 33, 2, 678, 679, 7, 81, 2,
	2, 679, 680, 8, 32, 1, 2, 680, 63, 3, 2, 2, 2, 681, 682, 8, 33, 1, 2, 682,
	683, 5, 16, 9, 2, 683, 684, 8, 33, 1, 2, 684, 692, 3, 2, 2, 2, 685, 686,
	12, 4, 2, 2, 686, 687, 7, 62, 2, 2, 687, 688, 5, 16, 9, 2, 688, 689, 8,
	33, 1, 2, 689, 691, 3, 2, 2, 2, 690, 685, 3, 2, 2, 2, 691, 694, 3, 2, 2,
	2, 692, 690, 3, 2, 2, 2, 692, 693, 3, 2, 2, 2, 693, 65, 3, 2, 2, 2, 694,
	692, 3, 2, 2, 2, 695, 696, 7, 40, 2, 2, 696, 697, 5, 16, 9, 2, 697, 698,
	5, 4, 3, 2, 698, 699, 8, 34, 1, 2, 699, 67, 3, 2, 2, 2, 700, 701, 8, 35,
	1, 2, 701, 702, 5, 70, 36, 2, 702, 703, 8, 35, 1, 2, 703, 711, 3, 2, 2,
	2, 704, 705, 12, 4, 2, 2, 705, 706, 7, 62, 2, 2, 706, 707, 5, 70, 36, 2,
	707, 708, 8, 35, 1, 2, 708, 710, 3, 2, 2, 2, 709, 704, 3, 2, 2, 2, 710,
	713, 3, 2, 2, 2, 711, 709, 3, 2, 2, 2, 711, 712, 3, 2, 2, 2, 712, 69, 3,
	2, 2, 2, 713, 711, 3, 2, 2, 2, 714, 715, 7, 16, 2, 2, 715, 716, 7, 51,
	2, 2, 716, 717, 7, 59, 2, 2, 717, 718, 5, 18, 10, 2, 718, 719, 8, 36, 1,
	2, 719, 726, 3, 2, 2, 2, 720, 721, 7, 51, 2, 2, 721, 722, 7, 59, 2, 2,
	722, 723, 5, 18, 10, 2, 723, 724, 8, 36, 1, 2, 724, 726, 3, 2, 2, 2, 725,
	714, 3, 2, 2, 2, 725, 720, 3, 2, 2, 2, 726, 71, 3, 2, 2, 2, 727, 728, 7,
	51, 2, 2, 728, 729, 7, 80, 2, 2, 729, 730, 5, 74, 38, 2, 730, 731, 7, 81,
	2, 2, 731, 732, 8, 37, 1, 2, 732, 738, 3, 2, 2, 2, 733, 734, 7, 51, 2,
	2, 734, 735, 7, 80, 2, 2, 735, 736, 7, 81, 2, 2, 736, 738, 8, 37, 1, 2,
	737, 727, 3, 2, 2, 2, 737, 733, 3, 2, 2, 2, 738, 73, 3, 2, 2, 2, 739, 740,
	8, 38, 1, 2, 740, 741, 5, 76, 39, 2, 741, 742, 8, 38, 1, 2, 742, 750, 3,
	2, 2, 2, 743, 744, 12, 4, 2, 2, 744, 745, 7, 62, 2, 2, 745, 746, 5, 76,
	39, 2, 746, 747, 8, 38, 1, 2, 747, 749, 3, 2, 2, 2, 748, 743, 3, 2, 2,
	2, 749, 752, 3, 2, 2, 2, 750, 748, 3, 2, 2, 2, 750, 751, 3, 2, 2, 2, 751,
	75, 3, 2, 2, 2, 752, 750, 3, 2, 2, 2, 753, 754, 5, 16, 9, 2, 754, 755,
	8, 39, 1, 2, 755, 764, 3, 2, 2, 2, 756, 757, 7, 55, 2, 2, 757, 758, 7,
	16, 2, 2, 758, 759, 7, 51, 2, 2, 759, 764, 8, 39, 1, 2, 760, 761, 7, 55,
	2, 2, 761, 762, 7, 51, 2, 2, 762, 764, 8, 39, 1, 2, 763, 753, 3, 2, 2,
	2, 763, 756, 3, 2, 2, 2, 763, 760, 3, 2, 2, 2, 764, 77, 3, 2, 2, 2, 765,
	766, 8, 40, 1, 2, 766, 767, 5, 16, 9, 2, 767, 768, 8, 40, 1, 2, 768, 776,
	3, 2, 2, 2, 769, 770, 12, 4, 2, 2, 770, 771, 7, 62, 2, 2, 771, 772, 5,
	16, 9, 2, 772, 773, 8, 40, 1, 2, 773, 775, 3, 2, 2, 2, 774, 769, 3, 2,
	2, 2, 775, 778, 3, 2, 2, 2, 776, 774, 3, 2, 2, 2, 776, 777, 3, 2, 2, 2,
	777, 79, 3, 2, 2, 2, 778, 776, 3, 2, 2, 2, 779, 780, 7, 14, 2, 2, 780,
	781, 7, 58, 2, 2, 781, 782, 7, 28, 2, 2, 782, 783, 7, 80, 2, 2, 783, 784,
	7, 81, 2, 2, 784, 801, 8, 41, 1, 2, 785, 786, 7, 15, 2, 2, 786, 787, 7,
	78, 2, 2, 787, 788, 7, 84, 2, 2, 788, 789, 5, 78, 40, 2, 789, 790, 7, 85,
	2, 2, 790, 791, 8, 41, 1, 2, 791, 801, 3, 2, 2, 2, 792, 793, 7, 14, 2,
	2, 793, 794, 7, 58, 2, 2, 794, 795, 7, 35, 2, 2, 795, 796, 7, 80, 2, 2,
	796, 797, 5, 16, 9, 2, 797, 798, 7, 81, 2, 2, 798, 799, 8, 41, 1, 2, 799,
	801, 3, 2, 2, 2, 800, 779, 3, 2, 2, 2, 800, 785, 3, 2, 2, 2, 800, 792,
	3, 2, 2, 2, 801, 81, 3, 2, 2, 2, 32, 93, 148, 275, 349, 361, 400, 450,
	452, 469, 497, 502, 538, 543, 562, 574, 590, 599, 612, 628, 637, 646, 654,
	692, 711, 725, 737, 750, 763, 776, 800,
}
var literalNames = []string{
	"", "'bool'", "'char'", "'f64'", "'i64'", "'&str'", "'String'", "'usize'",
	"'main'", "'powf'", "'pow'", "'as'", "'Vec'", "'vec'", "'mut'", "'let'",
	"'struct'", "'to_string'", "'to_owned'", "'true'", "'false'", "'println!'",
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
	"inicio", "bloque", "instrucciones", "instruccion", "declaracion", "declaracion_funcion",
	"asignacion", "expresion", "tipo_dato", "control_if", "bloque_else_if",
	"else_if", "control_if_exp", "bloque_else_if_exp", "else_if_exp", "control_expresion",
	"control_match", "control_case", "cases", "case_match", "control_match_exp",
	"control_case_exp", "cases_exp", "case_match_exp", "ireturn", "ibreak",
	"icontinue", "control_loop", "control_loop_exp", "printNormal", "printFormato",
	"elementosPrint", "control_while", "parametros_funcion", "parametro", "llamada_funcion",
	"parametros_llamada", "parametro_llamada_referencia", "elementos_vector",
	"metodos_iniciar_vector",
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
	NparserRULE_inicio                       = 0
	NparserRULE_bloque                       = 1
	NparserRULE_instrucciones                = 2
	NparserRULE_instruccion                  = 3
	NparserRULE_declaracion                  = 4
	NparserRULE_declaracion_funcion          = 5
	NparserRULE_asignacion                   = 6
	NparserRULE_expresion                    = 7
	NparserRULE_tipo_dato                    = 8
	NparserRULE_control_if                   = 9
	NparserRULE_bloque_else_if               = 10
	NparserRULE_else_if                      = 11
	NparserRULE_control_if_exp               = 12
	NparserRULE_bloque_else_if_exp           = 13
	NparserRULE_else_if_exp                  = 14
	NparserRULE_control_expresion            = 15
	NparserRULE_control_match                = 16
	NparserRULE_control_case                 = 17
	NparserRULE_cases                        = 18
	NparserRULE_case_match                   = 19
	NparserRULE_control_match_exp            = 20
	NparserRULE_control_case_exp             = 21
	NparserRULE_cases_exp                    = 22
	NparserRULE_case_match_exp               = 23
	NparserRULE_ireturn                      = 24
	NparserRULE_ibreak                       = 25
	NparserRULE_icontinue                    = 26
	NparserRULE_control_loop                 = 27
	NparserRULE_control_loop_exp             = 28
	NparserRULE_printNormal                  = 29
	NparserRULE_printFormato                 = 30
	NparserRULE_elementosPrint               = 31
	NparserRULE_control_while                = 32
	NparserRULE_parametros_funcion           = 33
	NparserRULE_parametro                    = 34
	NparserRULE_llamada_funcion              = 35
	NparserRULE_parametros_llamada           = 36
	NparserRULE_parametro_llamada_referencia = 37
	NparserRULE_elementos_vector             = 38
	NparserRULE_metodos_iniciar_vector       = 39
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
		p.SetState(80)

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
		p.SetState(83)
		p.Match(NparserLLAVE_IZQ)
	}
	{
		p.SetState(84)

		var _x = p.Instrucciones()

		localctx.(*BloqueContext)._instrucciones = _x
	}
	{
		p.SetState(85)
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
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NparserVEC)|(1<<NparserVEC_M)|(1<<NparserLET)|(1<<NparserTRUE)|(1<<NparserFALSE)|(1<<NparserPRINT)|(1<<NparserFN))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(NparserIF-34))|(1<<(NparserMATCH-34))|(1<<(NparserLOOP-34))|(1<<(NparserWHILE-34))|(1<<(NparserRETURN-34))|(1<<(NparserBREAK-34))|(1<<(NparserCONTINUE-34))|(1<<(NparserPUB-34))|(1<<(NparserNUMERO-34))|(1<<(NparserDECIMAL-34))|(1<<(NparserID-34)))) != 0) || (((_la-74)&-(0x1f+1)) == 0 && ((1<<uint((_la-74)))&((1<<(NparserRESTA-74))|(1<<(NparserNOT-74))|(1<<(NparserPAR_IZQ-74))|(1<<(NparserCADENA-74))|(1<<(NparserCARACTER-74)))) != 0) {
		{
			p.SetState(88)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(93)
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

	// Get_llamada_funcion returns the _llamada_funcion rule contexts.
	Get_llamada_funcion() ILlamada_funcionContext

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Get_declaracion returns the _declaracion rule contexts.
	Get_declaracion() IDeclaracionContext

	// Get_declaracion_funcion returns the _declaracion_funcion rule contexts.
	Get_declaracion_funcion() IDeclaracion_funcionContext

	// Get_asignacion returns the _asignacion rule contexts.
	Get_asignacion() IAsignacionContext

	// Get_control_if returns the _control_if rule contexts.
	Get_control_if() IControl_ifContext

	// Get_control_match returns the _control_match rule contexts.
	Get_control_match() IControl_matchContext

	// Get_control_loop returns the _control_loop rule contexts.
	Get_control_loop() IControl_loopContext

	// Get_control_while returns the _control_while rule contexts.
	Get_control_while() IControl_whileContext

	// Get_ibreak returns the _ibreak rule contexts.
	Get_ibreak() IIbreakContext

	// Get_icontinue returns the _icontinue rule contexts.
	Get_icontinue() IIcontinueContext

	// Get_ireturn returns the _ireturn rule contexts.
	Get_ireturn() IIreturnContext

	// Get_printNormal returns the _printNormal rule contexts.
	Get_printNormal() IPrintNormalContext

	// Get_printFormato returns the _printFormato rule contexts.
	Get_printFormato() IPrintFormatoContext

	// Set_llamada_funcion sets the _llamada_funcion rule contexts.
	Set_llamada_funcion(ILlamada_funcionContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// Set_declaracion sets the _declaracion rule contexts.
	Set_declaracion(IDeclaracionContext)

	// Set_declaracion_funcion sets the _declaracion_funcion rule contexts.
	Set_declaracion_funcion(IDeclaracion_funcionContext)

	// Set_asignacion sets the _asignacion rule contexts.
	Set_asignacion(IAsignacionContext)

	// Set_control_if sets the _control_if rule contexts.
	Set_control_if(IControl_ifContext)

	// Set_control_match sets the _control_match rule contexts.
	Set_control_match(IControl_matchContext)

	// Set_control_loop sets the _control_loop rule contexts.
	Set_control_loop(IControl_loopContext)

	// Set_control_while sets the _control_while rule contexts.
	Set_control_while(IControl_whileContext)

	// Set_ibreak sets the _ibreak rule contexts.
	Set_ibreak(IIbreakContext)

	// Set_icontinue sets the _icontinue rule contexts.
	Set_icontinue(IIcontinueContext)

	// Set_ireturn sets the _ireturn rule contexts.
	Set_ireturn(IIreturnContext)

	// Set_printNormal sets the _printNormal rule contexts.
	Set_printNormal(IPrintNormalContext)

	// Set_printFormato sets the _printFormato rule contexts.
	Set_printFormato(IPrintFormatoContext)

	// GetEx returns the ex attribute.
	GetEx() interface{}

	// SetEx sets the ex attribute.
	SetEx(interface{})

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser               antlr.Parser
	ex                   interface{}
	_llamada_funcion     ILlamada_funcionContext
	_expresion           IExpresionContext
	_declaracion         IDeclaracionContext
	_declaracion_funcion IDeclaracion_funcionContext
	_asignacion          IAsignacionContext
	_control_if          IControl_ifContext
	_control_match       IControl_matchContext
	_control_loop        IControl_loopContext
	_control_while       IControl_whileContext
	_ibreak              IIbreakContext
	_icontinue           IIcontinueContext
	_ireturn             IIreturnContext
	_printNormal         IPrintNormalContext
	_printFormato        IPrintFormatoContext
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

func (s *InstruccionContext) Get_llamada_funcion() ILlamada_funcionContext { return s._llamada_funcion }

func (s *InstruccionContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *InstruccionContext) Get_declaracion() IDeclaracionContext { return s._declaracion }

func (s *InstruccionContext) Get_declaracion_funcion() IDeclaracion_funcionContext {
	return s._declaracion_funcion
}

func (s *InstruccionContext) Get_asignacion() IAsignacionContext { return s._asignacion }

func (s *InstruccionContext) Get_control_if() IControl_ifContext { return s._control_if }

func (s *InstruccionContext) Get_control_match() IControl_matchContext { return s._control_match }

func (s *InstruccionContext) Get_control_loop() IControl_loopContext { return s._control_loop }

func (s *InstruccionContext) Get_control_while() IControl_whileContext { return s._control_while }

func (s *InstruccionContext) Get_ibreak() IIbreakContext { return s._ibreak }

func (s *InstruccionContext) Get_icontinue() IIcontinueContext { return s._icontinue }

func (s *InstruccionContext) Get_ireturn() IIreturnContext { return s._ireturn }

func (s *InstruccionContext) Get_printNormal() IPrintNormalContext { return s._printNormal }

func (s *InstruccionContext) Get_printFormato() IPrintFormatoContext { return s._printFormato }

func (s *InstruccionContext) Set_llamada_funcion(v ILlamada_funcionContext) { s._llamada_funcion = v }

func (s *InstruccionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *InstruccionContext) Set_declaracion(v IDeclaracionContext) { s._declaracion = v }

func (s *InstruccionContext) Set_declaracion_funcion(v IDeclaracion_funcionContext) {
	s._declaracion_funcion = v
}

func (s *InstruccionContext) Set_asignacion(v IAsignacionContext) { s._asignacion = v }

func (s *InstruccionContext) Set_control_if(v IControl_ifContext) { s._control_if = v }

func (s *InstruccionContext) Set_control_match(v IControl_matchContext) { s._control_match = v }

func (s *InstruccionContext) Set_control_loop(v IControl_loopContext) { s._control_loop = v }

func (s *InstruccionContext) Set_control_while(v IControl_whileContext) { s._control_while = v }

func (s *InstruccionContext) Set_ibreak(v IIbreakContext) { s._ibreak = v }

func (s *InstruccionContext) Set_icontinue(v IIcontinueContext) { s._icontinue = v }

func (s *InstruccionContext) Set_ireturn(v IIreturnContext) { s._ireturn = v }

func (s *InstruccionContext) Set_printNormal(v IPrintNormalContext) { s._printNormal = v }

func (s *InstruccionContext) Set_printFormato(v IPrintFormatoContext) { s._printFormato = v }

func (s *InstruccionContext) GetEx() interface{} { return s.ex }

func (s *InstruccionContext) SetEx(v interface{}) { s.ex = v }

func (s *InstruccionContext) Llamada_funcion() ILlamada_funcionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILlamada_funcionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILlamada_funcionContext)
}

func (s *InstruccionContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(NparserPUNTOCOMA, 0)
}

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

func (s *InstruccionContext) Declaracion_funcion() IDeclaracion_funcionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaracion_funcionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaracion_funcionContext)
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

func (s *InstruccionContext) Control_while() IControl_whileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControl_whileContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControl_whileContext)
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

func (s *InstruccionContext) PrintNormal() IPrintNormalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrintNormalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrintNormalContext)
}

func (s *InstruccionContext) PrintFormato() IPrintFormatoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrintFormatoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrintFormatoContext)
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

	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(96)

			var _x = p.Llamada_funcion()

			localctx.(*InstruccionContext)._llamada_funcion = _x
		}
		{
			p.SetState(97)
			p.Match(NparserPUNTOCOMA)
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_llamada_funcion().GetEx()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(100)

			var _x = p.expresion(0)

			localctx.(*InstruccionContext)._expresion = _x
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_expresion().GetEx()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(103)

			var _x = p.Declaracion()

			localctx.(*InstruccionContext)._declaracion = _x
		}
		{
			p.SetState(104)
			p.Match(NparserPUNTOCOMA)
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_declaracion().GetEx()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(107)

			var _x = p.Declaracion_funcion()

			localctx.(*InstruccionContext)._declaracion_funcion = _x
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_declaracion_funcion().GetEx()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(110)

			var _x = p.Asignacion()

			localctx.(*InstruccionContext)._asignacion = _x
		}
		{
			p.SetState(111)
			p.Match(NparserPUNTOCOMA)
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_asignacion().GetEx()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(114)

			var _x = p.Control_if()

			localctx.(*InstruccionContext)._control_if = _x
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_control_if().GetEx()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(117)

			var _x = p.Control_match()

			localctx.(*InstruccionContext)._control_match = _x
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_control_match().GetEx()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(120)

			var _x = p.Control_loop()

			localctx.(*InstruccionContext)._control_loop = _x
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_control_loop().GetEx()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(123)

			var _x = p.Control_while()

			localctx.(*InstruccionContext)._control_while = _x
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_control_while().GetEx()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(126)

			var _x = p.Ibreak()

			localctx.(*InstruccionContext)._ibreak = _x
		}
		{
			p.SetState(127)
			p.Match(NparserPUNTOCOMA)
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_ibreak().GetEx()

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(130)

			var _x = p.Icontinue()

			localctx.(*InstruccionContext)._icontinue = _x
		}
		{
			p.SetState(131)
			p.Match(NparserPUNTOCOMA)
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_icontinue().GetEx()

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(134)

			var _x = p.Ireturn()

			localctx.(*InstruccionContext)._ireturn = _x
		}
		{
			p.SetState(135)
			p.Match(NparserPUNTOCOMA)
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_ireturn().GetEx()

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(138)

			var _x = p.PrintNormal()

			localctx.(*InstruccionContext)._printNormal = _x
		}
		{
			p.SetState(139)
			p.Match(NparserPUNTOCOMA)
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_printNormal().GetEx()

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(142)

			var _x = p.PrintFormato()

			localctx.(*InstruccionContext)._printFormato = _x
		}
		{
			p.SetState(143)
			p.Match(NparserPUNTOCOMA)
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_printFormato().GetEx()

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

func (s *DeclaracionContext) VEC() antlr.TerminalNode {
	return s.GetToken(NparserVEC, 0)
}

func (s *DeclaracionContext) MENOR() antlr.TerminalNode {
	return s.GetToken(NparserMENOR, 0)
}

func (s *DeclaracionContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(NparserMAYOR, 0)
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

	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(149)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(150)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(151)

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
			p.SetState(154)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(155)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(156)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(157)

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
			p.SetState(160)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(161)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(162)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(163)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(164)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(165)

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
			p.SetState(168)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(169)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(170)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(171)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(172)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(173)

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
			p.SetState(176)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(177)
			p.Match(NparserMUT)
		}
		{
			p.SetState(178)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(179)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(180)

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
			p.SetState(183)
			p.Match(NparserLET)
		}
		{
			p.SetState(184)
			p.Match(NparserMUT)
		}
		{
			p.SetState(185)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(186)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(187)

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
			p.SetState(190)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(191)
			p.Match(NparserMUT)
		}
		{
			p.SetState(192)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(193)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(194)

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
			p.SetState(197)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(198)
			p.Match(NparserMUT)
		}
		{
			p.SetState(199)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(200)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(201)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(202)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(203)

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
			p.SetState(206)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(207)
			p.Match(NparserMUT)
		}
		{
			p.SetState(208)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(209)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(210)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(211)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(212)

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

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(215)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(216)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(217)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(218)

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
		}()), Ast.VECTOR, false, false, Ast.VOID, localctx.(*DeclaracionContext).Get_expresion().GetEx(), fila, columna)

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(221)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(222)
			p.Match(NparserMUT)
		}
		{
			p.SetState(223)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(224)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(225)

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
		}()), Ast.VECTOR, true, false, Ast.VOID, localctx.(*DeclaracionContext).Get_expresion().GetEx(), fila, columna)

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(228)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(229)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(230)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(231)
			p.Match(NparserVEC)
		}
		{
			p.SetState(232)
			p.Match(NparserMENOR)
		}
		{
			p.SetState(233)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(234)
			p.Match(NparserMAYOR)
		}
		{
			p.SetState(235)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(236)

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
		}()), Ast.VECTOR, false, false, localctx.(*DeclaracionContext).Get_tipo_dato().GetEx(), localctx.(*DeclaracionContext).Get_expresion().GetEx(), fila, columna)

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(239)

			var _m = p.Match(NparserLET)

			localctx.(*DeclaracionContext)._LET = _m
		}
		{
			p.SetState(240)
			p.Match(NparserMUT)
		}
		{
			p.SetState(241)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(242)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(243)
			p.Match(NparserVEC)
		}
		{
			p.SetState(244)
			p.Match(NparserMENOR)
		}
		{
			p.SetState(245)

			var _x = p.Tipo_dato()

			localctx.(*DeclaracionContext)._tipo_dato = _x
		}
		{
			p.SetState(246)
			p.Match(NparserMAYOR)
		}
		{
			p.SetState(247)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(248)

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
		}()), Ast.VECTOR, true, false, localctx.(*DeclaracionContext).Get_tipo_dato().GetEx(), localctx.(*DeclaracionContext).Get_expresion().GetEx(), fila, columna)

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(251)
			p.Match(NparserLET)
		}
		{
			p.SetState(252)
			p.Match(NparserMUT)
		}
		{
			p.SetState(253)
			p.Match(NparserID)
		}
		{
			p.SetState(254)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(255)
			p.Match(NparserVEC)
		}
		{
			p.SetState(256)
			p.Match(NparserMENOR)
		}
		{
			p.SetState(257)
			p.Tipo_dato()
		}
		{
			p.SetState(258)
			p.Match(NparserMAYOR)
		}
		{
			p.SetState(259)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(260)
			p.expresion(0)
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(263)
			p.Match(NparserLET)
		}
		{
			p.SetState(264)
			p.Match(NparserID)
		}
		{
			p.SetState(265)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(266)
			p.Match(NparserVEC)
		}
		{
			p.SetState(267)
			p.Match(NparserMENOR)
		}
		{
			p.SetState(268)
			p.Tipo_dato()
		}
		{
			p.SetState(269)
			p.Match(NparserMAYOR)
		}
		{
			p.SetState(270)
			p.Match(NparserIGUAL)
		}

	}

	return localctx
}

// IDeclaracion_funcionContext is an interface to support dynamic dispatch.
type IDeclaracion_funcionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_FN returns the _FN token.
	Get_FN() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_FN sets the _FN token.
	Set_FN(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_tipo_dato returns the _tipo_dato rule contexts.
	Get_tipo_dato() ITipo_datoContext

	// Get_bloque returns the _bloque rule contexts.
	Get_bloque() IBloqueContext

	// Get_parametros_funcion returns the _parametros_funcion rule contexts.
	Get_parametros_funcion() IParametros_funcionContext

	// Set_tipo_dato sets the _tipo_dato rule contexts.
	Set_tipo_dato(ITipo_datoContext)

	// Set_bloque sets the _bloque rule contexts.
	Set_bloque(IBloqueContext)

	// Set_parametros_funcion sets the _parametros_funcion rule contexts.
	Set_parametros_funcion(IParametros_funcionContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsDeclaracion_funcionContext differentiates from other interfaces.
	IsDeclaracion_funcionContext()
}

type Declaracion_funcionContext struct {
	*antlr.BaseParserRuleContext
	parser              antlr.Parser
	ex                  Ast.Instruccion
	_FN                 antlr.Token
	_ID                 antlr.Token
	_tipo_dato          ITipo_datoContext
	_bloque             IBloqueContext
	_parametros_funcion IParametros_funcionContext
}

func NewEmptyDeclaracion_funcionContext() *Declaracion_funcionContext {
	var p = new(Declaracion_funcionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_declaracion_funcion
	return p
}

func (*Declaracion_funcionContext) IsDeclaracion_funcionContext() {}

func NewDeclaracion_funcionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaracion_funcionContext {
	var p = new(Declaracion_funcionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_declaracion_funcion

	return p
}

func (s *Declaracion_funcionContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaracion_funcionContext) Get_FN() antlr.Token { return s._FN }

func (s *Declaracion_funcionContext) Get_ID() antlr.Token { return s._ID }

func (s *Declaracion_funcionContext) Set_FN(v antlr.Token) { s._FN = v }

func (s *Declaracion_funcionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Declaracion_funcionContext) Get_tipo_dato() ITipo_datoContext { return s._tipo_dato }

func (s *Declaracion_funcionContext) Get_bloque() IBloqueContext { return s._bloque }

func (s *Declaracion_funcionContext) Get_parametros_funcion() IParametros_funcionContext {
	return s._parametros_funcion
}

func (s *Declaracion_funcionContext) Set_tipo_dato(v ITipo_datoContext) { s._tipo_dato = v }

func (s *Declaracion_funcionContext) Set_bloque(v IBloqueContext) { s._bloque = v }

func (s *Declaracion_funcionContext) Set_parametros_funcion(v IParametros_funcionContext) {
	s._parametros_funcion = v
}

func (s *Declaracion_funcionContext) GetEx() Ast.Instruccion { return s.ex }

func (s *Declaracion_funcionContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *Declaracion_funcionContext) PUB() antlr.TerminalNode {
	return s.GetToken(NparserPUB, 0)
}

func (s *Declaracion_funcionContext) FN() antlr.TerminalNode {
	return s.GetToken(NparserFN, 0)
}

func (s *Declaracion_funcionContext) ID() antlr.TerminalNode {
	return s.GetToken(NparserID, 0)
}

func (s *Declaracion_funcionContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(NparserPAR_IZQ, 0)
}

func (s *Declaracion_funcionContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(NparserPAR_DER, 0)
}

func (s *Declaracion_funcionContext) FN_TIPO_RETORNO() antlr.TerminalNode {
	return s.GetToken(NparserFN_TIPO_RETORNO, 0)
}

func (s *Declaracion_funcionContext) Tipo_dato() ITipo_datoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITipo_datoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITipo_datoContext)
}

func (s *Declaracion_funcionContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Declaracion_funcionContext) Parametros_funcion() IParametros_funcionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametros_funcionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametros_funcionContext)
}

func (s *Declaracion_funcionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_funcionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declaracion_funcionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterDeclaracion_funcion(s)
	}
}

func (s *Declaracion_funcionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitDeclaracion_funcion(s)
	}
}

func (p *Nparser) Declaracion_funcion() (localctx IDeclaracion_funcionContext) {
	this := p
	_ = this

	localctx = NewDeclaracion_funcionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NparserRULE_declaracion_funcion)

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

	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(275)
			p.Match(NparserPUB)
		}
		{
			p.SetState(276)

			var _m = p.Match(NparserFN)

			localctx.(*Declaracion_funcionContext)._FN = _m
		}
		{
			p.SetState(277)

			var _m = p.Match(NparserID)

			localctx.(*Declaracion_funcionContext)._ID = _m
		}
		{
			p.SetState(278)
			p.Match(NparserPAR_IZQ)
		}
		{
			p.SetState(279)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(280)
			p.Match(NparserFN_TIPO_RETORNO)
		}
		{
			p.SetState(281)

			var _x = p.Tipo_dato()

			localctx.(*Declaracion_funcionContext)._tipo_dato = _x
		}
		{
			p.SetState(282)

			var _x = p.Bloque()

			localctx.(*Declaracion_funcionContext)._bloque = _x
		}

		parametros := arraylist.New()
		fila := (func() int {
			if localctx.(*Declaracion_funcionContext).Get_FN() == nil {
				return 0
			} else {
				return localctx.(*Declaracion_funcionContext).Get_FN().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Declaracion_funcionContext).Get_FN() == nil {
				return 0
			} else {
				return localctx.(*Declaracion_funcionContext).Get_FN().GetColumn()
			}
		}())
		funcion := simbolos.NewFuncion((func() string {
			if localctx.(*Declaracion_funcionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declaracion_funcionContext).Get_ID().GetText()
			}
		}()), Ast.FUNCION, localctx.(*Declaracion_funcionContext).Get_bloque().GetList(),
			parametros, localctx.(*Declaracion_funcionContext).Get_tipo_dato().GetEx(), true, fila, columna)
		localctx.(*Declaracion_funcionContext).ex = instrucciones.NewDeclaracion((func() string {
			if localctx.(*Declaracion_funcionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declaracion_funcionContext).Get_ID().GetText()
			}
		}()), Ast.FUNCION, false, true, Ast.VOID,
			funcion, fila, columna)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(285)
			p.Match(NparserPUB)
		}
		{
			p.SetState(286)

			var _m = p.Match(NparserFN)

			localctx.(*Declaracion_funcionContext)._FN = _m
		}
		{
			p.SetState(287)

			var _m = p.Match(NparserID)

			localctx.(*Declaracion_funcionContext)._ID = _m
		}
		{
			p.SetState(288)
			p.Match(NparserPAR_IZQ)
		}
		{
			p.SetState(289)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(290)

			var _x = p.Bloque()

			localctx.(*Declaracion_funcionContext)._bloque = _x
		}

		parametros := arraylist.New()
		fila := (func() int {
			if localctx.(*Declaracion_funcionContext).Get_FN() == nil {
				return 0
			} else {
				return localctx.(*Declaracion_funcionContext).Get_FN().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Declaracion_funcionContext).Get_FN() == nil {
				return 0
			} else {
				return localctx.(*Declaracion_funcionContext).Get_FN().GetColumn()
			}
		}())
		funcion := simbolos.NewFuncion((func() string {
			if localctx.(*Declaracion_funcionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declaracion_funcionContext).Get_ID().GetText()
			}
		}()), Ast.FUNCION, localctx.(*Declaracion_funcionContext).Get_bloque().GetList(),
			parametros, Ast.VOID, true, fila, columna)
		localctx.(*Declaracion_funcionContext).ex = instrucciones.NewDeclaracion((func() string {
			if localctx.(*Declaracion_funcionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declaracion_funcionContext).Get_ID().GetText()
			}
		}()), Ast.FUNCION, false, true, Ast.VOID,
			funcion, fila, columna)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(293)

			var _m = p.Match(NparserFN)

			localctx.(*Declaracion_funcionContext)._FN = _m
		}
		{
			p.SetState(294)

			var _m = p.Match(NparserID)

			localctx.(*Declaracion_funcionContext)._ID = _m
		}
		{
			p.SetState(295)
			p.Match(NparserPAR_IZQ)
		}
		{
			p.SetState(296)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(297)
			p.Match(NparserFN_TIPO_RETORNO)
		}
		{
			p.SetState(298)

			var _x = p.Tipo_dato()

			localctx.(*Declaracion_funcionContext)._tipo_dato = _x
		}
		{
			p.SetState(299)

			var _x = p.Bloque()

			localctx.(*Declaracion_funcionContext)._bloque = _x
		}

		fila := (func() int {
			if localctx.(*Declaracion_funcionContext).Get_FN() == nil {
				return 0
			} else {
				return localctx.(*Declaracion_funcionContext).Get_FN().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Declaracion_funcionContext).Get_FN() == nil {
				return 0
			} else {
				return localctx.(*Declaracion_funcionContext).Get_FN().GetColumn()
			}
		}())
		parametros := arraylist.New()
		funcion := simbolos.NewFuncion((func() string {
			if localctx.(*Declaracion_funcionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declaracion_funcionContext).Get_ID().GetText()
			}
		}()), Ast.FUNCION, localctx.(*Declaracion_funcionContext).Get_bloque().GetList(),
			parametros, localctx.(*Declaracion_funcionContext).Get_tipo_dato().GetEx(), false, fila, columna)
		localctx.(*Declaracion_funcionContext).ex = instrucciones.NewDeclaracion((func() string {
			if localctx.(*Declaracion_funcionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declaracion_funcionContext).Get_ID().GetText()
			}
		}()), Ast.FUNCION, false, false, Ast.VOID,
			funcion, fila, columna)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(302)

			var _m = p.Match(NparserFN)

			localctx.(*Declaracion_funcionContext)._FN = _m
		}
		{
			p.SetState(303)

			var _m = p.Match(NparserID)

			localctx.(*Declaracion_funcionContext)._ID = _m
		}
		{
			p.SetState(304)
			p.Match(NparserPAR_IZQ)
		}
		{
			p.SetState(305)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(306)

			var _x = p.Bloque()

			localctx.(*Declaracion_funcionContext)._bloque = _x
		}

		fila := (func() int {
			if localctx.(*Declaracion_funcionContext).Get_FN() == nil {
				return 0
			} else {
				return localctx.(*Declaracion_funcionContext).Get_FN().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Declaracion_funcionContext).Get_FN() == nil {
				return 0
			} else {
				return localctx.(*Declaracion_funcionContext).Get_FN().GetColumn()
			}
		}())
		parametros := arraylist.New()
		funcion := simbolos.NewFuncion((func() string {
			if localctx.(*Declaracion_funcionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declaracion_funcionContext).Get_ID().GetText()
			}
		}()), Ast.FUNCION, localctx.(*Declaracion_funcionContext).Get_bloque().GetList(),
			parametros, Ast.VOID, false, fila, columna)
		localctx.(*Declaracion_funcionContext).ex = instrucciones.NewDeclaracion((func() string {
			if localctx.(*Declaracion_funcionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declaracion_funcionContext).Get_ID().GetText()
			}
		}()), Ast.FUNCION, false, false, Ast.VOID,
			funcion, fila, columna)

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(309)
			p.Match(NparserPUB)
		}
		{
			p.SetState(310)

			var _m = p.Match(NparserFN)

			localctx.(*Declaracion_funcionContext)._FN = _m
		}
		{
			p.SetState(311)

			var _m = p.Match(NparserID)

			localctx.(*Declaracion_funcionContext)._ID = _m
		}
		{
			p.SetState(312)
			p.Match(NparserPAR_IZQ)
		}
		{
			p.SetState(313)

			var _x = p.parametros_funcion(0)

			localctx.(*Declaracion_funcionContext)._parametros_funcion = _x
		}
		{
			p.SetState(314)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(315)
			p.Match(NparserFN_TIPO_RETORNO)
		}
		{
			p.SetState(316)

			var _x = p.Tipo_dato()

			localctx.(*Declaracion_funcionContext)._tipo_dato = _x
		}
		{
			p.SetState(317)

			var _x = p.Bloque()

			localctx.(*Declaracion_funcionContext)._bloque = _x
		}

		fila := (func() int {
			if localctx.(*Declaracion_funcionContext).Get_FN() == nil {
				return 0
			} else {
				return localctx.(*Declaracion_funcionContext).Get_FN().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Declaracion_funcionContext).Get_FN() == nil {
				return 0
			} else {
				return localctx.(*Declaracion_funcionContext).Get_FN().GetColumn()
			}
		}())
		funcion := simbolos.NewFuncion((func() string {
			if localctx.(*Declaracion_funcionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declaracion_funcionContext).Get_ID().GetText()
			}
		}()), Ast.FUNCION, localctx.(*Declaracion_funcionContext).Get_bloque().GetList(),
			localctx.(*Declaracion_funcionContext).Get_parametros_funcion().GetList(), localctx.(*Declaracion_funcionContext).Get_tipo_dato().GetEx(), true, fila, columna)
		localctx.(*Declaracion_funcionContext).ex = instrucciones.NewDeclaracion((func() string {
			if localctx.(*Declaracion_funcionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declaracion_funcionContext).Get_ID().GetText()
			}
		}()), Ast.FUNCION, false, true, Ast.VOID,
			funcion, fila, columna)

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(320)
			p.Match(NparserPUB)
		}
		{
			p.SetState(321)

			var _m = p.Match(NparserFN)

			localctx.(*Declaracion_funcionContext)._FN = _m
		}
		{
			p.SetState(322)

			var _m = p.Match(NparserID)

			localctx.(*Declaracion_funcionContext)._ID = _m
		}
		{
			p.SetState(323)
			p.Match(NparserPAR_IZQ)
		}
		{
			p.SetState(324)

			var _x = p.parametros_funcion(0)

			localctx.(*Declaracion_funcionContext)._parametros_funcion = _x
		}
		{
			p.SetState(325)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(326)

			var _x = p.Bloque()

			localctx.(*Declaracion_funcionContext)._bloque = _x
		}

		fila := (func() int {
			if localctx.(*Declaracion_funcionContext).Get_FN() == nil {
				return 0
			} else {
				return localctx.(*Declaracion_funcionContext).Get_FN().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Declaracion_funcionContext).Get_FN() == nil {
				return 0
			} else {
				return localctx.(*Declaracion_funcionContext).Get_FN().GetColumn()
			}
		}())
		funcion := simbolos.NewFuncion((func() string {
			if localctx.(*Declaracion_funcionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declaracion_funcionContext).Get_ID().GetText()
			}
		}()), Ast.FUNCION, localctx.(*Declaracion_funcionContext).Get_bloque().GetList(),
			localctx.(*Declaracion_funcionContext).Get_parametros_funcion().GetList(), Ast.VOID, true, fila, columna)
		localctx.(*Declaracion_funcionContext).ex = instrucciones.NewDeclaracion((func() string {
			if localctx.(*Declaracion_funcionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declaracion_funcionContext).Get_ID().GetText()
			}
		}()), Ast.FUNCION, false, true, Ast.VOID,
			funcion, fila, columna)

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(329)

			var _m = p.Match(NparserFN)

			localctx.(*Declaracion_funcionContext)._FN = _m
		}
		{
			p.SetState(330)

			var _m = p.Match(NparserID)

			localctx.(*Declaracion_funcionContext)._ID = _m
		}
		{
			p.SetState(331)
			p.Match(NparserPAR_IZQ)
		}
		{
			p.SetState(332)

			var _x = p.parametros_funcion(0)

			localctx.(*Declaracion_funcionContext)._parametros_funcion = _x
		}
		{
			p.SetState(333)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(334)
			p.Match(NparserFN_TIPO_RETORNO)
		}
		{
			p.SetState(335)

			var _x = p.Tipo_dato()

			localctx.(*Declaracion_funcionContext)._tipo_dato = _x
		}
		{
			p.SetState(336)

			var _x = p.Bloque()

			localctx.(*Declaracion_funcionContext)._bloque = _x
		}

		fila := (func() int {
			if localctx.(*Declaracion_funcionContext).Get_FN() == nil {
				return 0
			} else {
				return localctx.(*Declaracion_funcionContext).Get_FN().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Declaracion_funcionContext).Get_FN() == nil {
				return 0
			} else {
				return localctx.(*Declaracion_funcionContext).Get_FN().GetColumn()
			}
		}())
		funcion := simbolos.NewFuncion((func() string {
			if localctx.(*Declaracion_funcionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declaracion_funcionContext).Get_ID().GetText()
			}
		}()), Ast.FUNCION, localctx.(*Declaracion_funcionContext).Get_bloque().GetList(),
			localctx.(*Declaracion_funcionContext).Get_parametros_funcion().GetList(), localctx.(*Declaracion_funcionContext).Get_tipo_dato().GetEx(), false, fila, columna)
		localctx.(*Declaracion_funcionContext).ex = instrucciones.NewDeclaracion((func() string {
			if localctx.(*Declaracion_funcionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declaracion_funcionContext).Get_ID().GetText()
			}
		}()), Ast.FUNCION, false, false, Ast.VOID,
			funcion, fila, columna)

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(339)

			var _m = p.Match(NparserFN)

			localctx.(*Declaracion_funcionContext)._FN = _m
		}
		{
			p.SetState(340)

			var _m = p.Match(NparserID)

			localctx.(*Declaracion_funcionContext)._ID = _m
		}
		{
			p.SetState(341)
			p.Match(NparserPAR_IZQ)
		}
		{
			p.SetState(342)

			var _x = p.parametros_funcion(0)

			localctx.(*Declaracion_funcionContext)._parametros_funcion = _x
		}
		{
			p.SetState(343)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(344)

			var _x = p.Bloque()

			localctx.(*Declaracion_funcionContext)._bloque = _x
		}

		fila := (func() int {
			if localctx.(*Declaracion_funcionContext).Get_FN() == nil {
				return 0
			} else {
				return localctx.(*Declaracion_funcionContext).Get_FN().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Declaracion_funcionContext).Get_FN() == nil {
				return 0
			} else {
				return localctx.(*Declaracion_funcionContext).Get_FN().GetColumn()
			}
		}())
		funcion := simbolos.NewFuncion((func() string {
			if localctx.(*Declaracion_funcionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declaracion_funcionContext).Get_ID().GetText()
			}
		}()), Ast.FUNCION, localctx.(*Declaracion_funcionContext).Get_bloque().GetList(),
			localctx.(*Declaracion_funcionContext).Get_parametros_funcion().GetList(), Ast.VOID, false, fila, columna)
		localctx.(*Declaracion_funcionContext).ex = instrucciones.NewDeclaracion((func() string {
			if localctx.(*Declaracion_funcionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Declaracion_funcionContext).Get_ID().GetText()
			}
		}()), Ast.FUNCION, false, false, Ast.VOID,
			funcion, fila, columna)

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

	// Get_control_expresion returns the _control_expresion rule contexts.
	Get_control_expresion() IControl_expresionContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// Set_control_expresion sets the _control_expresion rule contexts.
	Set_control_expresion(IControl_expresionContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	ex                 Ast.Instruccion
	_ID                antlr.Token
	_expresion         IExpresionContext
	_control_expresion IControl_expresionContext
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

func (s *AsignacionContext) Get_control_expresion() IControl_expresionContext {
	return s._control_expresion
}

func (s *AsignacionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *AsignacionContext) Set_control_expresion(v IControl_expresionContext) {
	s._control_expresion = v
}

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

func (s *AsignacionContext) Control_expresion() IControl_expresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControl_expresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControl_expresionContext)
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
	p.EnterRule(localctx, 12, NparserRULE_asignacion)

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

	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(349)

			var _m = p.Match(NparserID)

			localctx.(*AsignacionContext)._ID = _m
		}
		{
			p.SetState(350)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(351)

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
			p.SetState(354)

			var _m = p.Match(NparserID)

			localctx.(*AsignacionContext)._ID = _m
		}
		{
			p.SetState(355)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(356)

			var _x = p.Control_expresion()

			localctx.(*AsignacionContext)._control_expresion = _x
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
		}()), localctx.(*AsignacionContext).Get_control_expresion().GetEx(), fila, columna)

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

	// Get_PAR_IZQ returns the _PAR_IZQ token.
	Get_PAR_IZQ() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_TRUE returns the _TRUE token.
	Get_TRUE() antlr.Token

	// Get_FALSE returns the _FALSE token.
	Get_FALSE() antlr.Token

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

	// Get_PUNTO returns the _PUNTO token.
	Get_PUNTO() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Set_PAR_IZQ sets the _PAR_IZQ token.
	Set_PAR_IZQ(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_TRUE sets the _TRUE token.
	Set_TRUE(antlr.Token)

	// Set_FALSE sets the _FALSE token.
	Set_FALSE(antlr.Token)

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

	// Set_PUNTO sets the _PUNTO token.
	Set_PUNTO(antlr.Token)

	// GetE returns the e rule contexts.
	GetE() IExpresionContext

	// GetOp_izq returns the op_izq rule contexts.
	GetOp_izq() IExpresionContext

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Get_tipo_dato returns the _tipo_dato rule contexts.
	Get_tipo_dato() ITipo_datoContext

	// Get_llamada_funcion returns the _llamada_funcion rule contexts.
	Get_llamada_funcion() ILlamada_funcionContext

	// Get_metodos_iniciar_vector returns the _metodos_iniciar_vector rule contexts.
	Get_metodos_iniciar_vector() IMetodos_iniciar_vectorContext

	// GetOp_der returns the op_der rule contexts.
	GetOp_der() IExpresionContext

	// SetE sets the e rule contexts.
	SetE(IExpresionContext)

	// SetOp_izq sets the op_izq rule contexts.
	SetOp_izq(IExpresionContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// Set_tipo_dato sets the _tipo_dato rule contexts.
	Set_tipo_dato(ITipo_datoContext)

	// Set_llamada_funcion sets the _llamada_funcion rule contexts.
	Set_llamada_funcion(ILlamada_funcionContext)

	// Set_metodos_iniciar_vector sets the _metodos_iniciar_vector rule contexts.
	Set_metodos_iniciar_vector(IMetodos_iniciar_vectorContext)

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
	parser                  antlr.Parser
	ex                      Ast.Expresion
	e                       IExpresionContext
	op_izq                  IExpresionContext
	op                      antlr.Token
	_expresion              IExpresionContext
	_PAR_IZQ                antlr.Token
	_tipo_dato              ITipo_datoContext
	_llamada_funcion        ILlamada_funcionContext
	_metodos_iniciar_vector IMetodos_iniciar_vectorContext
	_ID                     antlr.Token
	_TRUE                   antlr.Token
	_FALSE                  antlr.Token
	_CARACTER               antlr.Token
	_DECIMAL                antlr.Token
	_NUMERO                 antlr.Token
	_CADENA                 antlr.Token
	op_der                  IExpresionContext
	_AND                    antlr.Token
	_OR                     antlr.Token
	_PUNTO                  antlr.Token
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

func (s *ExpresionContext) Get_PAR_IZQ() antlr.Token { return s._PAR_IZQ }

func (s *ExpresionContext) Get_ID() antlr.Token { return s._ID }

func (s *ExpresionContext) Get_TRUE() antlr.Token { return s._TRUE }

func (s *ExpresionContext) Get_FALSE() antlr.Token { return s._FALSE }

func (s *ExpresionContext) Get_CARACTER() antlr.Token { return s._CARACTER }

func (s *ExpresionContext) Get_DECIMAL() antlr.Token { return s._DECIMAL }

func (s *ExpresionContext) Get_NUMERO() antlr.Token { return s._NUMERO }

func (s *ExpresionContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *ExpresionContext) Get_AND() antlr.Token { return s._AND }

func (s *ExpresionContext) Get_OR() antlr.Token { return s._OR }

func (s *ExpresionContext) Get_PUNTO() antlr.Token { return s._PUNTO }

func (s *ExpresionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpresionContext) Set_PAR_IZQ(v antlr.Token) { s._PAR_IZQ = v }

func (s *ExpresionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ExpresionContext) Set_TRUE(v antlr.Token) { s._TRUE = v }

func (s *ExpresionContext) Set_FALSE(v antlr.Token) { s._FALSE = v }

func (s *ExpresionContext) Set_CARACTER(v antlr.Token) { s._CARACTER = v }

func (s *ExpresionContext) Set_DECIMAL(v antlr.Token) { s._DECIMAL = v }

func (s *ExpresionContext) Set_NUMERO(v antlr.Token) { s._NUMERO = v }

func (s *ExpresionContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *ExpresionContext) Set_AND(v antlr.Token) { s._AND = v }

func (s *ExpresionContext) Set_OR(v antlr.Token) { s._OR = v }

func (s *ExpresionContext) Set_PUNTO(v antlr.Token) { s._PUNTO = v }

func (s *ExpresionContext) GetE() IExpresionContext { return s.e }

func (s *ExpresionContext) GetOp_izq() IExpresionContext { return s.op_izq }

func (s *ExpresionContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *ExpresionContext) Get_tipo_dato() ITipo_datoContext { return s._tipo_dato }

func (s *ExpresionContext) Get_llamada_funcion() ILlamada_funcionContext { return s._llamada_funcion }

func (s *ExpresionContext) Get_metodos_iniciar_vector() IMetodos_iniciar_vectorContext {
	return s._metodos_iniciar_vector
}

func (s *ExpresionContext) GetOp_der() IExpresionContext { return s.op_der }

func (s *ExpresionContext) SetE(v IExpresionContext) { s.e = v }

func (s *ExpresionContext) SetOp_izq(v IExpresionContext) { s.op_izq = v }

func (s *ExpresionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *ExpresionContext) Set_tipo_dato(v ITipo_datoContext) { s._tipo_dato = v }

func (s *ExpresionContext) Set_llamada_funcion(v ILlamada_funcionContext) { s._llamada_funcion = v }

func (s *ExpresionContext) Set_metodos_iniciar_vector(v IMetodos_iniciar_vectorContext) {
	s._metodos_iniciar_vector = v
}

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

func (s *ExpresionContext) AS() antlr.TerminalNode {
	return s.GetToken(NparserAS, 0)
}

func (s *ExpresionContext) Tipo_dato() ITipo_datoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITipo_datoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITipo_datoContext)
}

func (s *ExpresionContext) Llamada_funcion() ILlamada_funcionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILlamada_funcionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILlamada_funcionContext)
}

func (s *ExpresionContext) Metodos_iniciar_vector() IMetodos_iniciar_vectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMetodos_iniciar_vectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMetodos_iniciar_vectorContext)
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

func (s *ExpresionContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(NparserPUNTO, 0)
}

func (s *ExpresionContext) TO_STRING() antlr.TerminalNode {
	return s.GetToken(NparserTO_STRING, 0)
}

func (s *ExpresionContext) SQRT() antlr.TerminalNode {
	return s.GetToken(NparserSQRT, 0)
}

func (s *ExpresionContext) ABS() antlr.TerminalNode {
	return s.GetToken(NparserABS, 0)
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
	_startState := 14
	p.EnterRecursionRule(localctx, 14, NparserRULE_expresion, _p)
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
	p.SetState(398)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(362)

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
			p.SetState(363)

			var _x = p.expresion(21)

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

	case 2:
		{
			p.SetState(366)

			var _m = p.Match(NparserPAR_IZQ)

			localctx.(*ExpresionContext)._PAR_IZQ = _m
		}
		{
			p.SetState(367)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext)._expresion = _x
		}
		{
			p.SetState(368)
			p.Match(NparserPAR_DER)
		}

		localctx.(*ExpresionContext).ex = localctx.(*ExpresionContext).Get_expresion().GetEx()

	case 3:
		{
			p.SetState(371)

			var _m = p.Match(NparserPAR_IZQ)

			localctx.(*ExpresionContext)._PAR_IZQ = _m
		}
		{
			p.SetState(372)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext)._expresion = _x
		}
		{
			p.SetState(373)
			p.Match(NparserAS)
		}
		{
			p.SetState(374)

			var _x = p.Tipo_dato()

			localctx.(*ExpresionContext)._tipo_dato = _x
		}
		{
			p.SetState(375)
			p.Match(NparserPAR_DER)
		}

		//Cast
		fila := (func() int {
			if localctx.(*ExpresionContext).Get_PAR_IZQ() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_PAR_IZQ().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*ExpresionContext).Get_PAR_IZQ() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_PAR_IZQ().GetColumn()
			}
		}())
		localctx.(*ExpresionContext).ex = expresiones.NewCast(localctx.(*ExpresionContext).Get_expresion().GetEx(), Ast.CAST, localctx.(*ExpresionContext).Get_tipo_dato().GetEx(), fila, columna)

	case 4:
		{
			p.SetState(378)

			var _x = p.Llamada_funcion()

			localctx.(*ExpresionContext)._llamada_funcion = _x
		}

		localctx.(*ExpresionContext).ex = localctx.(*ExpresionContext).Get_llamada_funcion().GetEx()

	case 5:
		{
			p.SetState(381)

			var _x = p.Metodos_iniciar_vector()

			localctx.(*ExpresionContext)._metodos_iniciar_vector = _x
		}

		localctx.(*ExpresionContext).ex = localctx.(*ExpresionContext).Get_metodos_iniciar_vector().GetEx()

	case 6:
		{
			p.SetState(384)

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
			if localctx.(*ExpresionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_ID().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*ExpresionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_ID().GetColumn()
			}
		}())
		localctx.(*ExpresionContext).ex = expresiones.NewIdentificador(id, Ast.IDENTIFICADOR, fila, columna)

	case 7:
		{
			p.SetState(386)

			var _m = p.Match(NparserTRUE)

			localctx.(*ExpresionContext)._TRUE = _m
		}

		valor := true
		fila := (func() int {
			if localctx.(*ExpresionContext).Get_TRUE() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_TRUE().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*ExpresionContext).Get_TRUE() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_TRUE().GetColumn()
			}
		}())
		localctx.(*ExpresionContext).ex = expresiones.NewPrimitivo(valor, Ast.BOOLEAN, fila, columna)

	case 8:
		{
			p.SetState(388)

			var _m = p.Match(NparserFALSE)

			localctx.(*ExpresionContext)._FALSE = _m
		}

		valor := false
		fila := (func() int {
			if localctx.(*ExpresionContext).Get_FALSE() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_FALSE().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*ExpresionContext).Get_FALSE() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_FALSE().GetColumn()
			}
		}())
		localctx.(*ExpresionContext).ex = expresiones.NewPrimitivo(valor, Ast.BOOLEAN, fila, columna)

	case 9:
		{
			p.SetState(390)

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
			if localctx.(*ExpresionContext).Get_CARACTER() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_CARACTER().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*ExpresionContext).Get_CARACTER() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_CARACTER().GetColumn()
			}
		}())
		localctx.(*ExpresionContext).ex = expresiones.NewPrimitivo(valor, Ast.CHAR, fila, columna)

	case 10:
		{
			p.SetState(392)

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

	case 11:
		{
			p.SetState(394)

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

	case 12:
		{
			p.SetState(396)

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

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(448)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).op_izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(400)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(401)

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
					p.SetState(402)

					var _x = p.expresion(18)

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
				p.SetState(405)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(406)

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
					p.SetState(407)

					var _x = p.expresion(17)

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
				p.SetState(410)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(411)

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
					p.SetState(412)

					var _x = p.expresion(16)

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
				p.SetState(415)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(416)

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
					p.SetState(417)

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

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).op_izq = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(420)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(421)

					var _m = p.Match(NparserAND)

					localctx.(*ExpresionContext)._AND = _m
				}
				{
					p.SetState(422)

					var _x = p.expresion(14)

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
				p.SetState(425)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(426)

					var _m = p.Match(NparserOR)

					localctx.(*ExpresionContext)._OR = _m
				}
				{
					p.SetState(427)

					var _x = p.expresion(13)

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

			case 7:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(430)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
				}
				{
					p.SetState(431)

					var _m = p.Match(NparserPUNTO)

					localctx.(*ExpresionContext)._PUNTO = _m
				}
				{
					p.SetState(432)
					p.Match(NparserTO_STRING)
				}
				{
					p.SetState(433)

					var _m = p.Match(NparserPAR_IZQ)

					localctx.(*ExpresionContext)._PAR_IZQ = _m
				}
				{
					p.SetState(434)
					p.Match(NparserPAR_DER)
				}

				fila := (func() int {
					if localctx.(*ExpresionContext).Get_PUNTO() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_PUNTO().GetLine()
					}
				}())
				columna := (func() int {
					if localctx.(*ExpresionContext).Get_PUNTO() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_PUNTO().GetColumn()
					}
				}()) - 1
				localctx.(*ExpresionContext).ex = fn_primitivas.NewToString(Ast.LLAMADA_FUNCION, localctx.(*ExpresionContext).GetE().GetEx(), fila, columna)

			case 8:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(436)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				{
					p.SetState(437)

					var _m = p.Match(NparserPUNTO)

					localctx.(*ExpresionContext)._PUNTO = _m
				}
				{
					p.SetState(438)
					p.Match(NparserSQRT)
				}
				{
					p.SetState(439)

					var _m = p.Match(NparserPAR_IZQ)

					localctx.(*ExpresionContext)._PAR_IZQ = _m
				}
				{
					p.SetState(440)
					p.Match(NparserPAR_DER)
				}

				fila := (func() int {
					if localctx.(*ExpresionContext).Get_PUNTO() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_PUNTO().GetLine()
					}
				}())
				columna := (func() int {
					if localctx.(*ExpresionContext).Get_PUNTO() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_PUNTO().GetColumn()
					}
				}()) - 1
				localctx.(*ExpresionContext).ex = fn_primitivas.NewSqrt(Ast.LLAMADA_FUNCION, localctx.(*ExpresionContext).GetE().GetEx(), fila, columna)

			case 9:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(442)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(443)

					var _m = p.Match(NparserPUNTO)

					localctx.(*ExpresionContext)._PUNTO = _m
				}
				{
					p.SetState(444)
					p.Match(NparserABS)
				}
				{
					p.SetState(445)

					var _m = p.Match(NparserPAR_IZQ)

					localctx.(*ExpresionContext)._PAR_IZQ = _m
				}
				{
					p.SetState(446)
					p.Match(NparserPAR_DER)
				}

				fila := (func() int {
					if localctx.(*ExpresionContext).Get_PUNTO() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_PUNTO().GetLine()
					}
				}())
				columna := (func() int {
					if localctx.(*ExpresionContext).Get_PUNTO() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).Get_PUNTO().GetColumn()
					}
				}()) - 1
				localctx.(*ExpresionContext).ex = fn_primitivas.NewAbs(Ast.LLAMADA_FUNCION, localctx.(*ExpresionContext).GetE().GetEx(), fila, columna)

			}

		}
		p.SetState(452)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 16, NparserRULE_tipo_dato)

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

	p.SetState(467)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserBOOL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(453)
			p.Match(NparserBOOL)
		}
		localctx.(*Tipo_datoContext).ex = Ast.BOOLEAN

	case NparserCHAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(455)
			p.Match(NparserCHAR)
		}
		localctx.(*Tipo_datoContext).ex = Ast.CHAR

	case NparserI64:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(457)
			p.Match(NparserI64)
		}
		localctx.(*Tipo_datoContext).ex = Ast.I64

	case NparserF64:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(459)
			p.Match(NparserF64)
		}
		localctx.(*Tipo_datoContext).ex = Ast.F64

	case NparserSTR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(461)
			p.Match(NparserSTR)
		}
		localctx.(*Tipo_datoContext).ex = Ast.STR

	case NparserSTRING:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(463)
			p.Match(NparserSTRING)
		}
		localctx.(*Tipo_datoContext).ex = Ast.STRING

	case NparserUSIZE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(465)
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
	p.EnterRule(localctx, 18, NparserRULE_control_if)

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

	p.SetState(495)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(469)

			var _m = p.Match(NparserIF)

			localctx.(*Control_ifContext)._IF = _m
		}
		{
			p.SetState(470)

			var _x = p.expresion(0)

			localctx.(*Control_ifContext)._expresion = _x
		}
		{
			p.SetState(471)

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
			p.SetState(474)

			var _m = p.Match(NparserIF)

			localctx.(*Control_ifContext)._IF = _m
		}
		{
			p.SetState(475)

			var _x = p.expresion(0)

			localctx.(*Control_ifContext)._expresion = _x
		}
		{
			p.SetState(476)

			var _x = p.Bloque()

			localctx.(*Control_ifContext).bloqueIf = _x
		}
		{
			p.SetState(477)
			p.Match(NparserELSE)
		}
		{
			p.SetState(478)

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
			p.SetState(481)

			var _m = p.Match(NparserIF)

			localctx.(*Control_ifContext)._IF = _m
		}
		{
			p.SetState(482)

			var _x = p.expresion(0)

			localctx.(*Control_ifContext)._expresion = _x
		}
		{
			p.SetState(483)

			var _x = p.Bloque()

			localctx.(*Control_ifContext).bloqueIf = _x
		}
		{
			p.SetState(484)

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
			p.SetState(487)

			var _m = p.Match(NparserIF)

			localctx.(*Control_ifContext)._IF = _m
		}
		{
			p.SetState(488)

			var _x = p.expresion(0)

			localctx.(*Control_ifContext)._expresion = _x
		}
		{
			p.SetState(489)

			var _x = p.Bloque()

			localctx.(*Control_ifContext).bloqueIf = _x
		}
		{
			p.SetState(490)

			var _x = p.Bloque_else_if()

			localctx.(*Control_ifContext)._bloque_else_if = _x
		}
		{
			p.SetState(491)
			p.Match(NparserELSE)
		}
		{
			p.SetState(492)

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
	p.EnterRule(localctx, 20, NparserRULE_bloque_else_if)
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
	p.SetState(498)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(497)

				var _x = p.Else_if()

				localctx.(*Bloque_else_ifContext)._else_if = _x
			}
			localctx.(*Bloque_else_ifContext).lista = append(localctx.(*Bloque_else_ifContext).lista, localctx.(*Bloque_else_ifContext)._else_if)

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(500)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 22, NparserRULE_else_if)

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
		p.SetState(504)

		var _m = p.Match(NparserELSE)

		localctx.(*Else_ifContext)._ELSE = _m
	}
	{
		p.SetState(505)
		p.Match(NparserIF)
	}
	{
		p.SetState(506)

		var _x = p.expresion(0)

		localctx.(*Else_ifContext)._expresion = _x
	}
	{
		p.SetState(507)

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
	p.EnterRule(localctx, 24, NparserRULE_control_if_exp)

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

	p.SetState(536)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(510)

			var _m = p.Match(NparserIF)

			localctx.(*Control_if_expContext)._IF = _m
		}
		{
			p.SetState(511)

			var _x = p.expresion(0)

			localctx.(*Control_if_expContext)._expresion = _x
		}
		{
			p.SetState(512)

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
			p.SetState(515)

			var _m = p.Match(NparserIF)

			localctx.(*Control_if_expContext)._IF = _m
		}
		{
			p.SetState(516)

			var _x = p.expresion(0)

			localctx.(*Control_if_expContext)._expresion = _x
		}
		{
			p.SetState(517)

			var _x = p.Bloque()

			localctx.(*Control_if_expContext).bloqueIf = _x
		}
		{
			p.SetState(518)
			p.Match(NparserELSE)
		}
		{
			p.SetState(519)

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
			p.SetState(522)

			var _m = p.Match(NparserIF)

			localctx.(*Control_if_expContext)._IF = _m
		}
		{
			p.SetState(523)

			var _x = p.expresion(0)

			localctx.(*Control_if_expContext)._expresion = _x
		}
		{
			p.SetState(524)

			var _x = p.Bloque()

			localctx.(*Control_if_expContext).bloqueIf = _x
		}
		{
			p.SetState(525)

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
			p.SetState(528)

			var _m = p.Match(NparserIF)

			localctx.(*Control_if_expContext)._IF = _m
		}
		{
			p.SetState(529)

			var _x = p.expresion(0)

			localctx.(*Control_if_expContext)._expresion = _x
		}
		{
			p.SetState(530)

			var _x = p.Bloque()

			localctx.(*Control_if_expContext).bloqueIf = _x
		}
		{
			p.SetState(531)

			var _x = p.Bloque_else_if_exp()

			localctx.(*Control_if_expContext)._bloque_else_if_exp = _x
		}
		{
			p.SetState(532)
			p.Match(NparserELSE)
		}
		{
			p.SetState(533)

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
	p.EnterRule(localctx, 26, NparserRULE_bloque_else_if_exp)
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
	p.SetState(539)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(538)

				var _x = p.Else_if_exp()

				localctx.(*Bloque_else_if_expContext)._else_if_exp = _x
			}
			localctx.(*Bloque_else_if_expContext).lista = append(localctx.(*Bloque_else_if_expContext).lista, localctx.(*Bloque_else_if_expContext)._else_if_exp)

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(541)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 28, NparserRULE_else_if_exp)

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
		p.SetState(545)

		var _m = p.Match(NparserELSE)

		localctx.(*Else_if_expContext)._ELSE = _m
	}
	{
		p.SetState(546)
		p.Match(NparserIF)
	}
	{
		p.SetState(547)

		var _x = p.expresion(0)

		localctx.(*Else_if_expContext)._expresion = _x
	}
	{
		p.SetState(548)

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
	p.EnterRule(localctx, 30, NparserRULE_control_expresion)

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

	p.SetState(560)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(551)

			var _x = p.Control_if_exp()

			localctx.(*Control_expresionContext)._control_if_exp = _x
		}
		localctx.(*Control_expresionContext).ex = localctx.(*Control_expresionContext).Get_control_if_exp().GetEx()

	case NparserMATCH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(554)

			var _x = p.Control_match_exp()

			localctx.(*Control_expresionContext)._control_match_exp = _x
		}
		localctx.(*Control_expresionContext).ex = localctx.(*Control_expresionContext).Get_control_match_exp().GetEx()

	case NparserLOOP:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(557)

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
	p.EnterRule(localctx, 32, NparserRULE_control_match)

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
		p.SetState(562)

		var _m = p.Match(NparserMATCH)

		localctx.(*Control_matchContext)._MATCH = _m
	}
	{
		p.SetState(563)

		var _x = p.expresion(0)

		localctx.(*Control_matchContext)._expresion = _x
	}
	{
		p.SetState(564)
		p.Match(NparserLLAVE_IZQ)
	}
	{
		p.SetState(565)

		var _x = p.Control_case()

		localctx.(*Control_matchContext)._control_case = _x
	}
	{
		p.SetState(566)
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
	p.EnterRule(localctx, 34, NparserRULE_control_case)
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
	p.SetState(570)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NparserVEC)|(1<<NparserVEC_M)|(1<<NparserTRUE)|(1<<NparserFALSE))) != 0) || (((_la-46)&-(0x1f+1)) == 0 && ((1<<uint((_la-46)))&((1<<(NparserNUMERO-46))|(1<<(NparserDECIMAL-46))|(1<<(NparserID-46))|(1<<(NparserDEFAULT-46))|(1<<(NparserRESTA-46))|(1<<(NparserNOT-46)))) != 0) || (((_la-78)&-(0x1f+1)) == 0 && ((1<<uint((_la-78)))&((1<<(NparserPAR_IZQ-78))|(1<<(NparserCADENA-78))|(1<<(NparserCARACTER-78)))) != 0) {
		{
			p.SetState(569)

			var _x = p.Cases()

			localctx.(*Control_caseContext)._cases = _x
		}
		localctx.(*Control_caseContext).lista = append(localctx.(*Control_caseContext).lista, localctx.(*Control_caseContext)._cases)

		p.SetState(572)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

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
	p.EnterRule(localctx, 36, NparserRULE_cases)

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
		p.SetState(576)

		var _x = p.case_match(0)

		localctx.(*CasesContext)._case_match = _x
	}
	{
		p.SetState(577)

		var _m = p.Match(NparserCASE)

		localctx.(*CasesContext)._CASE = _m
	}
	{
		p.SetState(578)

		var _x = p.Bloque()

		localctx.(*CasesContext)._bloque = _x
	}
	{
		p.SetState(579)
		p.Match(NparserCOMA)
	}

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
	_startState := 38
	p.EnterRecursionRule(localctx, 38, NparserRULE_case_match, _p)
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
	p.SetState(588)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserVEC, NparserVEC_M, NparserTRUE, NparserFALSE, NparserNUMERO, NparserDECIMAL, NparserID, NparserRESTA, NparserNOT, NparserPAR_IZQ, NparserCADENA, NparserCARACTER:
		{
			p.SetState(583)

			var _x = p.expresion(0)

			localctx.(*Case_matchContext)._expresion = _x
		}

		localctx.(*Case_matchContext).list.Add(localctx.(*Case_matchContext).Get_expresion().GetEx())

	case NparserDEFAULT:
		{
			p.SetState(586)

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
	p.SetState(597)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCase_matchContext(p, _parentctx, _parentState)
			localctx.(*Case_matchContext).lista_cases = _prevctx
			p.PushNewRecursionContext(localctx, _startState, NparserRULE_case_match)
			p.SetState(590)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(591)
				p.Match(NparserO)
			}
			{
				p.SetState(592)

				var _x = p.expresion(0)

				localctx.(*Case_matchContext)._expresion = _x
			}

			localctx.(*Case_matchContext).GetLista_cases().GetList().Add(localctx.(*Case_matchContext).Get_expresion().GetEx())
			localctx.(*Case_matchContext).list = localctx.(*Case_matchContext).GetLista_cases().GetList()

		}
		p.SetState(599)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 40, NparserRULE_control_match_exp)

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
		p.SetState(600)

		var _m = p.Match(NparserMATCH)

		localctx.(*Control_match_expContext)._MATCH = _m
	}
	{
		p.SetState(601)

		var _x = p.expresion(0)

		localctx.(*Control_match_expContext)._expresion = _x
	}
	{
		p.SetState(602)
		p.Match(NparserLLAVE_IZQ)
	}
	{
		p.SetState(603)

		var _x = p.Control_case_exp()

		localctx.(*Control_match_expContext)._control_case_exp = _x
	}
	{
		p.SetState(604)
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
	p.EnterRule(localctx, 42, NparserRULE_control_case_exp)
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
	p.SetState(608)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NparserVEC)|(1<<NparserVEC_M)|(1<<NparserTRUE)|(1<<NparserFALSE))) != 0) || (((_la-46)&-(0x1f+1)) == 0 && ((1<<uint((_la-46)))&((1<<(NparserNUMERO-46))|(1<<(NparserDECIMAL-46))|(1<<(NparserID-46))|(1<<(NparserDEFAULT-46))|(1<<(NparserRESTA-46))|(1<<(NparserNOT-46)))) != 0) || (((_la-78)&-(0x1f+1)) == 0 && ((1<<uint((_la-78)))&((1<<(NparserPAR_IZQ-78))|(1<<(NparserCADENA-78))|(1<<(NparserCARACTER-78)))) != 0) {
		{
			p.SetState(607)

			var _x = p.Cases_exp()

			localctx.(*Control_case_expContext)._cases_exp = _x
		}
		localctx.(*Control_case_expContext).lista = append(localctx.(*Control_case_expContext).lista, localctx.(*Control_case_expContext)._cases_exp)

		p.SetState(610)
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
	p.EnterRule(localctx, 44, NparserRULE_cases_exp)

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
		p.SetState(614)

		var _x = p.case_match_exp(0)

		localctx.(*Cases_expContext)._case_match_exp = _x
	}
	{
		p.SetState(615)

		var _m = p.Match(NparserCASE)

		localctx.(*Cases_expContext)._CASE = _m
	}
	{
		p.SetState(616)

		var _x = p.Bloque()

		localctx.(*Cases_expContext)._bloque = _x
	}
	{
		p.SetState(617)
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
	_startState := 46
	p.EnterRecursionRule(localctx, 46, NparserRULE_case_match_exp, _p)
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
	p.SetState(626)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserVEC, NparserVEC_M, NparserTRUE, NparserFALSE, NparserNUMERO, NparserDECIMAL, NparserID, NparserRESTA, NparserNOT, NparserPAR_IZQ, NparserCADENA, NparserCARACTER:
		{
			p.SetState(621)

			var _x = p.expresion(0)

			localctx.(*Case_match_expContext)._expresion = _x
		}

		localctx.(*Case_match_expContext).list.Add(localctx.(*Case_match_expContext).Get_expresion().GetEx())

	case NparserDEFAULT:
		{
			p.SetState(624)

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
	p.SetState(635)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCase_match_expContext(p, _parentctx, _parentState)
			localctx.(*Case_match_expContext).lista_cases = _prevctx
			p.PushNewRecursionContext(localctx, _startState, NparserRULE_case_match_exp)
			p.SetState(628)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(629)
				p.Match(NparserO)
			}
			{
				p.SetState(630)

				var _x = p.expresion(0)

				localctx.(*Case_match_expContext)._expresion = _x
			}

			localctx.(*Case_match_expContext).GetLista_cases().GetList().Add(localctx.(*Case_match_expContext).Get_expresion().GetEx())
			localctx.(*Case_match_expContext).list = localctx.(*Case_match_expContext).GetLista_cases().GetList()

		}
		p.SetState(637)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 48, NparserRULE_ireturn)

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

	p.SetState(644)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(638)

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
			p.SetState(640)

			var _m = p.Match(NparserRETURN)

			localctx.(*IreturnContext)._RETURN = _m
		}
		{
			p.SetState(641)

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
	p.EnterRule(localctx, 50, NparserRULE_ibreak)

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

	p.SetState(652)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(646)

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
			p.SetState(648)

			var _m = p.Match(NparserBREAK)

			localctx.(*IbreakContext)._BREAK = _m
		}
		{
			p.SetState(649)

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
	p.EnterRule(localctx, 52, NparserRULE_icontinue)

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
		p.SetState(654)

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
	p.EnterRule(localctx, 54, NparserRULE_control_loop)

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
		p.SetState(657)

		var _m = p.Match(NparserLOOP)

		localctx.(*Control_loopContext)._LOOP = _m
	}
	{
		p.SetState(658)

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
	p.EnterRule(localctx, 56, NparserRULE_control_loop_exp)

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
		p.SetState(661)

		var _m = p.Match(NparserLOOP)

		localctx.(*Control_loop_expContext)._LOOP = _m
	}
	{
		p.SetState(662)

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

// IPrintNormalContext is an interface to support dynamic dispatch.
type IPrintNormalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_PRINT returns the _PRINT token.
	Get_PRINT() antlr.Token

	// Set_PRINT sets the _PRINT token.
	Set_PRINT(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsPrintNormalContext differentiates from other interfaces.
	IsPrintNormalContext()
}

type PrintNormalContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	ex         Ast.Instruccion
	_PRINT     antlr.Token
	_expresion IExpresionContext
}

func NewEmptyPrintNormalContext() *PrintNormalContext {
	var p = new(PrintNormalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_printNormal
	return p
}

func (*PrintNormalContext) IsPrintNormalContext() {}

func NewPrintNormalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintNormalContext {
	var p = new(PrintNormalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_printNormal

	return p
}

func (s *PrintNormalContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintNormalContext) Get_PRINT() antlr.Token { return s._PRINT }

func (s *PrintNormalContext) Set_PRINT(v antlr.Token) { s._PRINT = v }

func (s *PrintNormalContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *PrintNormalContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *PrintNormalContext) GetEx() Ast.Instruccion { return s.ex }

func (s *PrintNormalContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *PrintNormalContext) PRINT() antlr.TerminalNode {
	return s.GetToken(NparserPRINT, 0)
}

func (s *PrintNormalContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(NparserPAR_IZQ, 0)
}

func (s *PrintNormalContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *PrintNormalContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(NparserPAR_DER, 0)
}

func (s *PrintNormalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintNormalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintNormalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterPrintNormal(s)
	}
}

func (s *PrintNormalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitPrintNormal(s)
	}
}

func (p *Nparser) PrintNormal() (localctx IPrintNormalContext) {
	this := p
	_ = this

	localctx = NewPrintNormalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, NparserRULE_printNormal)

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
		p.SetState(665)

		var _m = p.Match(NparserPRINT)

		localctx.(*PrintNormalContext)._PRINT = _m
	}
	{
		p.SetState(666)
		p.Match(NparserPAR_IZQ)
	}
	{
		p.SetState(667)

		var _x = p.expresion(0)

		localctx.(*PrintNormalContext)._expresion = _x
	}
	{
		p.SetState(668)
		p.Match(NparserPAR_DER)
	}

	fila := (func() int {
		if localctx.(*PrintNormalContext).Get_PRINT() == nil {
			return 0
		} else {
			return localctx.(*PrintNormalContext).Get_PRINT().GetLine()
		}
	}())
	columna := (func() int {
		if localctx.(*PrintNormalContext).Get_PRINT() == nil {
			return 0
		} else {
			return localctx.(*PrintNormalContext).Get_PRINT().GetColumn()
		}
	}())
	localctx.(*PrintNormalContext).ex = instrucciones.NewPrint(localctx.(*PrintNormalContext).Get_expresion().GetEx(), Ast.PRINT, fila, columna)

	return localctx
}

// IPrintFormatoContext is an interface to support dynamic dispatch.
type IPrintFormatoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_PRINT returns the _PRINT token.
	Get_PRINT() antlr.Token

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Set_PRINT sets the _PRINT token.
	Set_PRINT(antlr.Token)

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// GetExpresiones returns the expresiones rule contexts.
	GetExpresiones() IElementosPrintContext

	// SetExpresiones sets the expresiones rule contexts.
	SetExpresiones(IElementosPrintContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsPrintFormatoContext differentiates from other interfaces.
	IsPrintFormatoContext()
}

type PrintFormatoContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	ex          Ast.Instruccion
	_PRINT      antlr.Token
	_CADENA     antlr.Token
	expresiones IElementosPrintContext
}

func NewEmptyPrintFormatoContext() *PrintFormatoContext {
	var p = new(PrintFormatoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_printFormato
	return p
}

func (*PrintFormatoContext) IsPrintFormatoContext() {}

func NewPrintFormatoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintFormatoContext {
	var p = new(PrintFormatoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_printFormato

	return p
}

func (s *PrintFormatoContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintFormatoContext) Get_PRINT() antlr.Token { return s._PRINT }

func (s *PrintFormatoContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *PrintFormatoContext) Set_PRINT(v antlr.Token) { s._PRINT = v }

func (s *PrintFormatoContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *PrintFormatoContext) GetExpresiones() IElementosPrintContext { return s.expresiones }

func (s *PrintFormatoContext) SetExpresiones(v IElementosPrintContext) { s.expresiones = v }

func (s *PrintFormatoContext) GetEx() Ast.Instruccion { return s.ex }

func (s *PrintFormatoContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *PrintFormatoContext) PRINT() antlr.TerminalNode {
	return s.GetToken(NparserPRINT, 0)
}

func (s *PrintFormatoContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(NparserPAR_IZQ, 0)
}

func (s *PrintFormatoContext) CADENA() antlr.TerminalNode {
	return s.GetToken(NparserCADENA, 0)
}

func (s *PrintFormatoContext) COMA() antlr.TerminalNode {
	return s.GetToken(NparserCOMA, 0)
}

func (s *PrintFormatoContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(NparserPAR_DER, 0)
}

func (s *PrintFormatoContext) ElementosPrint() IElementosPrintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementosPrintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementosPrintContext)
}

func (s *PrintFormatoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintFormatoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintFormatoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterPrintFormato(s)
	}
}

func (s *PrintFormatoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitPrintFormato(s)
	}
}

func (p *Nparser) PrintFormato() (localctx IPrintFormatoContext) {
	this := p
	_ = this

	localctx = NewPrintFormatoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, NparserRULE_printFormato)

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
		p.SetState(671)

		var _m = p.Match(NparserPRINT)

		localctx.(*PrintFormatoContext)._PRINT = _m
	}
	{
		p.SetState(672)
		p.Match(NparserPAR_IZQ)
	}
	{
		p.SetState(673)

		var _m = p.Match(NparserCADENA)

		localctx.(*PrintFormatoContext)._CADENA = _m
	}
	{
		p.SetState(674)
		p.Match(NparserCOMA)
	}
	{
		p.SetState(675)

		var _x = p.elementosPrint(0)

		localctx.(*PrintFormatoContext).expresiones = _x
	}
	{
		p.SetState(676)
		p.Match(NparserPAR_DER)
	}

	fila := (func() int {
		if localctx.(*PrintFormatoContext).Get_PRINT() == nil {
			return 0
		} else {
			return localctx.(*PrintFormatoContext).Get_PRINT().GetLine()
		}
	}())
	columna := (func() int {
		if localctx.(*PrintFormatoContext).Get_PRINT() == nil {
			return 0
		} else {
			return localctx.(*PrintFormatoContext).Get_PRINT().GetColumn()
		}
	}())
	valor := (func() string {
		if localctx.(*PrintFormatoContext).Get_CADENA() == nil {
			return ""
		} else {
			return localctx.(*PrintFormatoContext).Get_CADENA().GetText()
		}
	}())
	valor = valor[1 : len(valor)-1]
	localctx.(*PrintFormatoContext).ex = instrucciones.NewPrintF(localctx.(*PrintFormatoContext).GetExpresiones().GetList(), valor, Ast.PRINTF, fila, columna)

	return localctx
}

// IElementosPrintContext is an interface to support dynamic dispatch.
type IElementosPrintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLista_elementos returns the lista_elementos rule contexts.
	GetLista_elementos() IElementosPrintContext

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// SetLista_elementos sets the lista_elementos rule contexts.
	SetLista_elementos(IElementosPrintContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetList returns the list attribute.
	GetList() *arraylist.List

	// SetList sets the list attribute.
	SetList(*arraylist.List)

	// IsElementosPrintContext differentiates from other interfaces.
	IsElementosPrintContext()
}

type ElementosPrintContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	list            *arraylist.List
	lista_elementos IElementosPrintContext
	_expresion      IExpresionContext
}

func NewEmptyElementosPrintContext() *ElementosPrintContext {
	var p = new(ElementosPrintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_elementosPrint
	return p
}

func (*ElementosPrintContext) IsElementosPrintContext() {}

func NewElementosPrintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementosPrintContext {
	var p = new(ElementosPrintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_elementosPrint

	return p
}

func (s *ElementosPrintContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementosPrintContext) GetLista_elementos() IElementosPrintContext { return s.lista_elementos }

func (s *ElementosPrintContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *ElementosPrintContext) SetLista_elementos(v IElementosPrintContext) { s.lista_elementos = v }

func (s *ElementosPrintContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *ElementosPrintContext) GetList() *arraylist.List { return s.list }

func (s *ElementosPrintContext) SetList(v *arraylist.List) { s.list = v }

func (s *ElementosPrintContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ElementosPrintContext) COMA() antlr.TerminalNode {
	return s.GetToken(NparserCOMA, 0)
}

func (s *ElementosPrintContext) ElementosPrint() IElementosPrintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementosPrintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementosPrintContext)
}

func (s *ElementosPrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementosPrintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementosPrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterElementosPrint(s)
	}
}

func (s *ElementosPrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitElementosPrint(s)
	}
}

func (p *Nparser) ElementosPrint() (localctx IElementosPrintContext) {
	return p.elementosPrint(0)
}

func (p *Nparser) elementosPrint(_p int) (localctx IElementosPrintContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewElementosPrintContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IElementosPrintContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 62
	p.EnterRecursionRule(localctx, 62, NparserRULE_elementosPrint, _p)
	localctx.(*ElementosPrintContext).list = arraylist.New()

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
	{
		p.SetState(680)

		var _x = p.expresion(0)

		localctx.(*ElementosPrintContext)._expresion = _x
	}

	localctx.(*ElementosPrintContext).list.Add(localctx.(*ElementosPrintContext).Get_expresion().GetEx())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(690)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewElementosPrintContext(p, _parentctx, _parentState)
			localctx.(*ElementosPrintContext).lista_elementos = _prevctx
			p.PushNewRecursionContext(localctx, _startState, NparserRULE_elementosPrint)
			p.SetState(683)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(684)
				p.Match(NparserCOMA)
			}
			{
				p.SetState(685)

				var _x = p.expresion(0)

				localctx.(*ElementosPrintContext)._expresion = _x
			}

			localctx.(*ElementosPrintContext).GetLista_elementos().GetList().Add(localctx.(*ElementosPrintContext).Get_expresion().GetEx())
			localctx.(*ElementosPrintContext).list = localctx.(*ElementosPrintContext).GetLista_elementos().GetList()

		}
		p.SetState(692)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}

	return localctx
}

// IControl_whileContext is an interface to support dynamic dispatch.
type IControl_whileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_WHILE returns the _WHILE token.
	Get_WHILE() antlr.Token

	// Set_WHILE sets the _WHILE token.
	Set_WHILE(antlr.Token)

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

	// IsControl_whileContext differentiates from other interfaces.
	IsControl_whileContext()
}

type Control_whileContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	ex         Ast.Instruccion
	_WHILE     antlr.Token
	_expresion IExpresionContext
	_bloque    IBloqueContext
}

func NewEmptyControl_whileContext() *Control_whileContext {
	var p = new(Control_whileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_control_while
	return p
}

func (*Control_whileContext) IsControl_whileContext() {}

func NewControl_whileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Control_whileContext {
	var p = new(Control_whileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_control_while

	return p
}

func (s *Control_whileContext) GetParser() antlr.Parser { return s.parser }

func (s *Control_whileContext) Get_WHILE() antlr.Token { return s._WHILE }

func (s *Control_whileContext) Set_WHILE(v antlr.Token) { s._WHILE = v }

func (s *Control_whileContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *Control_whileContext) Get_bloque() IBloqueContext { return s._bloque }

func (s *Control_whileContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *Control_whileContext) Set_bloque(v IBloqueContext) { s._bloque = v }

func (s *Control_whileContext) GetEx() Ast.Instruccion { return s.ex }

func (s *Control_whileContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *Control_whileContext) WHILE() antlr.TerminalNode {
	return s.GetToken(NparserWHILE, 0)
}

func (s *Control_whileContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Control_whileContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Control_whileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Control_whileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Control_whileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterControl_while(s)
	}
}

func (s *Control_whileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitControl_while(s)
	}
}

func (p *Nparser) Control_while() (localctx IControl_whileContext) {
	this := p
	_ = this

	localctx = NewControl_whileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, NparserRULE_control_while)

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
		p.SetState(693)

		var _m = p.Match(NparserWHILE)

		localctx.(*Control_whileContext)._WHILE = _m
	}
	{
		p.SetState(694)

		var _x = p.expresion(0)

		localctx.(*Control_whileContext)._expresion = _x
	}
	{
		p.SetState(695)

		var _x = p.Bloque()

		localctx.(*Control_whileContext)._bloque = _x
	}

	fila := (func() int {
		if localctx.(*Control_whileContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*Control_whileContext).Get_WHILE().GetLine()
		}
	}())
	columna := (func() int {
		if localctx.(*Control_whileContext).Get_WHILE() == nil {
			return 0
		} else {
			return localctx.(*Control_whileContext).Get_WHILE().GetColumn()
		}
	}())
	localctx.(*Control_whileContext).ex = bucles.NewWhile(Ast.WHILE, localctx.(*Control_whileContext).Get_expresion().GetEx(), localctx.(*Control_whileContext).Get_bloque().GetList(), fila, columna)

	return localctx
}

// IParametros_funcionContext is an interface to support dynamic dispatch.
type IParametros_funcionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLista_elementos returns the lista_elementos rule contexts.
	GetLista_elementos() IParametros_funcionContext

	// Get_parametro returns the _parametro rule contexts.
	Get_parametro() IParametroContext

	// SetLista_elementos sets the lista_elementos rule contexts.
	SetLista_elementos(IParametros_funcionContext)

	// Set_parametro sets the _parametro rule contexts.
	Set_parametro(IParametroContext)

	// GetList returns the list attribute.
	GetList() *arraylist.List

	// SetList sets the list attribute.
	SetList(*arraylist.List)

	// IsParametros_funcionContext differentiates from other interfaces.
	IsParametros_funcionContext()
}

type Parametros_funcionContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	list            *arraylist.List
	lista_elementos IParametros_funcionContext
	_parametro      IParametroContext
}

func NewEmptyParametros_funcionContext() *Parametros_funcionContext {
	var p = new(Parametros_funcionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_parametros_funcion
	return p
}

func (*Parametros_funcionContext) IsParametros_funcionContext() {}

func NewParametros_funcionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parametros_funcionContext {
	var p = new(Parametros_funcionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_parametros_funcion

	return p
}

func (s *Parametros_funcionContext) GetParser() antlr.Parser { return s.parser }

func (s *Parametros_funcionContext) GetLista_elementos() IParametros_funcionContext {
	return s.lista_elementos
}

func (s *Parametros_funcionContext) Get_parametro() IParametroContext { return s._parametro }

func (s *Parametros_funcionContext) SetLista_elementos(v IParametros_funcionContext) {
	s.lista_elementos = v
}

func (s *Parametros_funcionContext) Set_parametro(v IParametroContext) { s._parametro = v }

func (s *Parametros_funcionContext) GetList() *arraylist.List { return s.list }

func (s *Parametros_funcionContext) SetList(v *arraylist.List) { s.list = v }

func (s *Parametros_funcionContext) Parametro() IParametroContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametroContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Parametros_funcionContext) COMA() antlr.TerminalNode {
	return s.GetToken(NparserCOMA, 0)
}

func (s *Parametros_funcionContext) Parametros_funcion() IParametros_funcionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametros_funcionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametros_funcionContext)
}

func (s *Parametros_funcionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parametros_funcionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parametros_funcionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterParametros_funcion(s)
	}
}

func (s *Parametros_funcionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitParametros_funcion(s)
	}
}

func (p *Nparser) Parametros_funcion() (localctx IParametros_funcionContext) {
	return p.parametros_funcion(0)
}

func (p *Nparser) parametros_funcion(_p int) (localctx IParametros_funcionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewParametros_funcionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IParametros_funcionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 66
	p.EnterRecursionRule(localctx, 66, NparserRULE_parametros_funcion, _p)
	localctx.(*Parametros_funcionContext).list = arraylist.New()

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
	{
		p.SetState(699)

		var _x = p.Parametro()

		localctx.(*Parametros_funcionContext)._parametro = _x
	}

	localctx.(*Parametros_funcionContext).list.Add(localctx.(*Parametros_funcionContext).Get_parametro().GetEx())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(709)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewParametros_funcionContext(p, _parentctx, _parentState)
			localctx.(*Parametros_funcionContext).lista_elementos = _prevctx
			p.PushNewRecursionContext(localctx, _startState, NparserRULE_parametros_funcion)
			p.SetState(702)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(703)
				p.Match(NparserCOMA)
			}
			{
				p.SetState(704)

				var _x = p.Parametro()

				localctx.(*Parametros_funcionContext)._parametro = _x
			}

			localctx.(*Parametros_funcionContext).GetLista_elementos().GetList().Add(localctx.(*Parametros_funcionContext).Get_parametro().GetEx())
			localctx.(*Parametros_funcionContext).list = localctx.(*Parametros_funcionContext).GetLista_elementos().GetList()

		}
		p.SetState(711)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// IParametroContext is an interface to support dynamic dispatch.
type IParametroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_MUT returns the _MUT token.
	Get_MUT() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_MUT sets the _MUT token.
	Set_MUT(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_tipo_dato returns the _tipo_dato rule contexts.
	Get_tipo_dato() ITipo_datoContext

	// Set_tipo_dato sets the _tipo_dato rule contexts.
	Set_tipo_dato(ITipo_datoContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Expresion

	// SetEx sets the ex attribute.
	SetEx(Ast.Expresion)

	// IsParametroContext differentiates from other interfaces.
	IsParametroContext()
}

type ParametroContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	ex         Ast.Expresion
	_MUT       antlr.Token
	_ID        antlr.Token
	_tipo_dato ITipo_datoContext
}

func NewEmptyParametroContext() *ParametroContext {
	var p = new(ParametroContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_parametro
	return p
}

func (*ParametroContext) IsParametroContext() {}

func NewParametroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametroContext {
	var p = new(ParametroContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_parametro

	return p
}

func (s *ParametroContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametroContext) Get_MUT() antlr.Token { return s._MUT }

func (s *ParametroContext) Get_ID() antlr.Token { return s._ID }

func (s *ParametroContext) Set_MUT(v antlr.Token) { s._MUT = v }

func (s *ParametroContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ParametroContext) Get_tipo_dato() ITipo_datoContext { return s._tipo_dato }

func (s *ParametroContext) Set_tipo_dato(v ITipo_datoContext) { s._tipo_dato = v }

func (s *ParametroContext) GetEx() Ast.Expresion { return s.ex }

func (s *ParametroContext) SetEx(v Ast.Expresion) { s.ex = v }

func (s *ParametroContext) MUT() antlr.TerminalNode {
	return s.GetToken(NparserMUT, 0)
}

func (s *ParametroContext) ID() antlr.TerminalNode {
	return s.GetToken(NparserID, 0)
}

func (s *ParametroContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(NparserDOSPUNTOS, 0)
}

func (s *ParametroContext) Tipo_dato() ITipo_datoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITipo_datoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITipo_datoContext)
}

func (s *ParametroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterParametro(s)
	}
}

func (s *ParametroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitParametro(s)
	}
}

func (p *Nparser) Parametro() (localctx IParametroContext) {
	this := p
	_ = this

	localctx = NewParametroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, NparserRULE_parametro)

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

	p.SetState(723)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserMUT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(712)

			var _m = p.Match(NparserMUT)

			localctx.(*ParametroContext)._MUT = _m
		}
		{
			p.SetState(713)

			var _m = p.Match(NparserID)

			localctx.(*ParametroContext)._ID = _m
		}
		{
			p.SetState(714)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(715)

			var _x = p.Tipo_dato()

			localctx.(*ParametroContext)._tipo_dato = _x
		}

		fila := (func() int {
			if localctx.(*ParametroContext).Get_MUT() == nil {
				return 0
			} else {
				return localctx.(*ParametroContext).Get_MUT().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*ParametroContext).Get_MUT() == nil {
				return 0
			} else {
				return localctx.(*ParametroContext).Get_MUT().GetColumn()
			}
		}())
		localctx.(*ParametroContext).ex = simbolos.NewParametro((func() string {
			if localctx.(*ParametroContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ParametroContext).Get_ID().GetText()
			}
		}()), Ast.PARAMETRO, localctx.(*ParametroContext).Get_tipo_dato().GetEx(), true, fila, columna)

	case NparserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(718)

			var _m = p.Match(NparserID)

			localctx.(*ParametroContext)._ID = _m
		}
		{
			p.SetState(719)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(720)

			var _x = p.Tipo_dato()

			localctx.(*ParametroContext)._tipo_dato = _x
		}

		fila := (func() int {
			if localctx.(*ParametroContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParametroContext).Get_ID().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*ParametroContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*ParametroContext).Get_ID().GetColumn()
			}
		}())
		localctx.(*ParametroContext).ex = simbolos.NewParametro((func() string {
			if localctx.(*ParametroContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ParametroContext).Get_ID().GetText()
			}
		}()), Ast.PARAMETRO, localctx.(*ParametroContext).Get_tipo_dato().GetEx(), false, fila, columna)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILlamada_funcionContext is an interface to support dynamic dispatch.
type ILlamada_funcionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_parametros_llamada returns the _parametros_llamada rule contexts.
	Get_parametros_llamada() IParametros_llamadaContext

	// Set_parametros_llamada sets the _parametros_llamada rule contexts.
	Set_parametros_llamada(IParametros_llamadaContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Expresion

	// SetEx sets the ex attribute.
	SetEx(Ast.Expresion)

	// IsLlamada_funcionContext differentiates from other interfaces.
	IsLlamada_funcionContext()
}

type Llamada_funcionContext struct {
	*antlr.BaseParserRuleContext
	parser              antlr.Parser
	ex                  Ast.Expresion
	_ID                 antlr.Token
	_parametros_llamada IParametros_llamadaContext
}

func NewEmptyLlamada_funcionContext() *Llamada_funcionContext {
	var p = new(Llamada_funcionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_llamada_funcion
	return p
}

func (*Llamada_funcionContext) IsLlamada_funcionContext() {}

func NewLlamada_funcionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Llamada_funcionContext {
	var p = new(Llamada_funcionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_llamada_funcion

	return p
}

func (s *Llamada_funcionContext) GetParser() antlr.Parser { return s.parser }

func (s *Llamada_funcionContext) Get_ID() antlr.Token { return s._ID }

func (s *Llamada_funcionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Llamada_funcionContext) Get_parametros_llamada() IParametros_llamadaContext {
	return s._parametros_llamada
}

func (s *Llamada_funcionContext) Set_parametros_llamada(v IParametros_llamadaContext) {
	s._parametros_llamada = v
}

func (s *Llamada_funcionContext) GetEx() Ast.Expresion { return s.ex }

func (s *Llamada_funcionContext) SetEx(v Ast.Expresion) { s.ex = v }

func (s *Llamada_funcionContext) ID() antlr.TerminalNode {
	return s.GetToken(NparserID, 0)
}

func (s *Llamada_funcionContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(NparserPAR_IZQ, 0)
}

func (s *Llamada_funcionContext) Parametros_llamada() IParametros_llamadaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametros_llamadaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametros_llamadaContext)
}

func (s *Llamada_funcionContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(NparserPAR_DER, 0)
}

func (s *Llamada_funcionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Llamada_funcionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Llamada_funcionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterLlamada_funcion(s)
	}
}

func (s *Llamada_funcionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitLlamada_funcion(s)
	}
}

func (p *Nparser) Llamada_funcion() (localctx ILlamada_funcionContext) {
	this := p
	_ = this

	localctx = NewLlamada_funcionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, NparserRULE_llamada_funcion)

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

	p.SetState(735)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(725)

			var _m = p.Match(NparserID)

			localctx.(*Llamada_funcionContext)._ID = _m
		}
		{
			p.SetState(726)
			p.Match(NparserPAR_IZQ)
		}
		{
			p.SetState(727)

			var _x = p.parametros_llamada(0)

			localctx.(*Llamada_funcionContext)._parametros_llamada = _x
		}
		{
			p.SetState(728)
			p.Match(NparserPAR_DER)
		}

		fila := (func() int {
			if localctx.(*Llamada_funcionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Llamada_funcionContext).Get_ID().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Llamada_funcionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Llamada_funcionContext).Get_ID().GetColumn()
			}
		}())
		id := expresiones.NewIdentificador((func() string {
			if localctx.(*Llamada_funcionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Llamada_funcionContext).Get_ID().GetText()
			}
		}()), Ast.IDENTIFICADOR, fila, columna)
		localctx.(*Llamada_funcionContext).ex = simbolos.NewLlamadaFuncion(id, localctx.(*Llamada_funcionContext).Get_parametros_llamada().GetList(), Ast.LLAMADA_FUNCION, fila, columna)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(731)

			var _m = p.Match(NparserID)

			localctx.(*Llamada_funcionContext)._ID = _m
		}
		{
			p.SetState(732)
			p.Match(NparserPAR_IZQ)
		}
		{
			p.SetState(733)
			p.Match(NparserPAR_DER)
		}

		fila := (func() int {
			if localctx.(*Llamada_funcionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Llamada_funcionContext).Get_ID().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Llamada_funcionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*Llamada_funcionContext).Get_ID().GetColumn()
			}
		}())
		params := arraylist.New()
		id := expresiones.NewIdentificador((func() string {
			if localctx.(*Llamada_funcionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Llamada_funcionContext).Get_ID().GetText()
			}
		}()), Ast.IDENTIFICADOR, fila, columna)
		localctx.(*Llamada_funcionContext).ex = simbolos.NewLlamadaFuncion(id, params, Ast.LLAMADA_FUNCION, fila, columna)

	}

	return localctx
}

// IParametros_llamadaContext is an interface to support dynamic dispatch.
type IParametros_llamadaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLista_elementos returns the lista_elementos rule contexts.
	GetLista_elementos() IParametros_llamadaContext

	// Get_parametro_llamada_referencia returns the _parametro_llamada_referencia rule contexts.
	Get_parametro_llamada_referencia() IParametro_llamada_referenciaContext

	// SetLista_elementos sets the lista_elementos rule contexts.
	SetLista_elementos(IParametros_llamadaContext)

	// Set_parametro_llamada_referencia sets the _parametro_llamada_referencia rule contexts.
	Set_parametro_llamada_referencia(IParametro_llamada_referenciaContext)

	// GetList returns the list attribute.
	GetList() *arraylist.List

	// SetList sets the list attribute.
	SetList(*arraylist.List)

	// IsParametros_llamadaContext differentiates from other interfaces.
	IsParametros_llamadaContext()
}

type Parametros_llamadaContext struct {
	*antlr.BaseParserRuleContext
	parser                        antlr.Parser
	list                          *arraylist.List
	lista_elementos               IParametros_llamadaContext
	_parametro_llamada_referencia IParametro_llamada_referenciaContext
}

func NewEmptyParametros_llamadaContext() *Parametros_llamadaContext {
	var p = new(Parametros_llamadaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_parametros_llamada
	return p
}

func (*Parametros_llamadaContext) IsParametros_llamadaContext() {}

func NewParametros_llamadaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parametros_llamadaContext {
	var p = new(Parametros_llamadaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_parametros_llamada

	return p
}

func (s *Parametros_llamadaContext) GetParser() antlr.Parser { return s.parser }

func (s *Parametros_llamadaContext) GetLista_elementos() IParametros_llamadaContext {
	return s.lista_elementos
}

func (s *Parametros_llamadaContext) Get_parametro_llamada_referencia() IParametro_llamada_referenciaContext {
	return s._parametro_llamada_referencia
}

func (s *Parametros_llamadaContext) SetLista_elementos(v IParametros_llamadaContext) {
	s.lista_elementos = v
}

func (s *Parametros_llamadaContext) Set_parametro_llamada_referencia(v IParametro_llamada_referenciaContext) {
	s._parametro_llamada_referencia = v
}

func (s *Parametros_llamadaContext) GetList() *arraylist.List { return s.list }

func (s *Parametros_llamadaContext) SetList(v *arraylist.List) { s.list = v }

func (s *Parametros_llamadaContext) Parametro_llamada_referencia() IParametro_llamada_referenciaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametro_llamada_referenciaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametro_llamada_referenciaContext)
}

func (s *Parametros_llamadaContext) COMA() antlr.TerminalNode {
	return s.GetToken(NparserCOMA, 0)
}

func (s *Parametros_llamadaContext) Parametros_llamada() IParametros_llamadaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametros_llamadaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametros_llamadaContext)
}

func (s *Parametros_llamadaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parametros_llamadaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parametros_llamadaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterParametros_llamada(s)
	}
}

func (s *Parametros_llamadaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitParametros_llamada(s)
	}
}

func (p *Nparser) Parametros_llamada() (localctx IParametros_llamadaContext) {
	return p.parametros_llamada(0)
}

func (p *Nparser) parametros_llamada(_p int) (localctx IParametros_llamadaContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewParametros_llamadaContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IParametros_llamadaContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 72
	p.EnterRecursionRule(localctx, 72, NparserRULE_parametros_llamada, _p)
	localctx.(*Parametros_llamadaContext).list = arraylist.New()

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
	{
		p.SetState(738)

		var _x = p.Parametro_llamada_referencia()

		localctx.(*Parametros_llamadaContext)._parametro_llamada_referencia = _x
	}

	localctx.(*Parametros_llamadaContext).list.Add(localctx.(*Parametros_llamadaContext).Get_parametro_llamada_referencia().GetEx())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(748)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewParametros_llamadaContext(p, _parentctx, _parentState)
			localctx.(*Parametros_llamadaContext).lista_elementos = _prevctx
			p.PushNewRecursionContext(localctx, _startState, NparserRULE_parametros_llamada)
			p.SetState(741)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(742)
				p.Match(NparserCOMA)
			}
			{
				p.SetState(743)

				var _x = p.Parametro_llamada_referencia()

				localctx.(*Parametros_llamadaContext)._parametro_llamada_referencia = _x
			}

			localctx.(*Parametros_llamadaContext).GetLista_elementos().GetList().Add(localctx.(*Parametros_llamadaContext).Get_parametro_llamada_referencia().GetEx())
			localctx.(*Parametros_llamadaContext).list = localctx.(*Parametros_llamadaContext).GetLista_elementos().GetList()

		}
		p.SetState(750)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
	}

	return localctx
}

// IParametro_llamada_referenciaContext is an interface to support dynamic dispatch.
type IParametro_llamada_referenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_AMPERSAND returns the _AMPERSAND token.
	Get_AMPERSAND() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_AMPERSAND sets the _AMPERSAND token.
	Set_AMPERSAND(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// GetE returns the e rule contexts.
	GetE() IExpresionContext

	// SetE sets the e rule contexts.
	SetE(IExpresionContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Expresion

	// SetEx sets the ex attribute.
	SetEx(Ast.Expresion)

	// IsParametro_llamada_referenciaContext differentiates from other interfaces.
	IsParametro_llamada_referenciaContext()
}

type Parametro_llamada_referenciaContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	ex         Ast.Expresion
	e          IExpresionContext
	_AMPERSAND antlr.Token
	_ID        antlr.Token
}

func NewEmptyParametro_llamada_referenciaContext() *Parametro_llamada_referenciaContext {
	var p = new(Parametro_llamada_referenciaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_parametro_llamada_referencia
	return p
}

func (*Parametro_llamada_referenciaContext) IsParametro_llamada_referenciaContext() {}

func NewParametro_llamada_referenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parametro_llamada_referenciaContext {
	var p = new(Parametro_llamada_referenciaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_parametro_llamada_referencia

	return p
}

func (s *Parametro_llamada_referenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *Parametro_llamada_referenciaContext) Get_AMPERSAND() antlr.Token { return s._AMPERSAND }

func (s *Parametro_llamada_referenciaContext) Get_ID() antlr.Token { return s._ID }

func (s *Parametro_llamada_referenciaContext) Set_AMPERSAND(v antlr.Token) { s._AMPERSAND = v }

func (s *Parametro_llamada_referenciaContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *Parametro_llamada_referenciaContext) GetE() IExpresionContext { return s.e }

func (s *Parametro_llamada_referenciaContext) SetE(v IExpresionContext) { s.e = v }

func (s *Parametro_llamada_referenciaContext) GetEx() Ast.Expresion { return s.ex }

func (s *Parametro_llamada_referenciaContext) SetEx(v Ast.Expresion) { s.ex = v }

func (s *Parametro_llamada_referenciaContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Parametro_llamada_referenciaContext) AMPERSAND() antlr.TerminalNode {
	return s.GetToken(NparserAMPERSAND, 0)
}

func (s *Parametro_llamada_referenciaContext) MUT() antlr.TerminalNode {
	return s.GetToken(NparserMUT, 0)
}

func (s *Parametro_llamada_referenciaContext) ID() antlr.TerminalNode {
	return s.GetToken(NparserID, 0)
}

func (s *Parametro_llamada_referenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parametro_llamada_referenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parametro_llamada_referenciaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterParametro_llamada_referencia(s)
	}
}

func (s *Parametro_llamada_referenciaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitParametro_llamada_referencia(s)
	}
}

func (p *Nparser) Parametro_llamada_referencia() (localctx IParametro_llamada_referenciaContext) {
	this := p
	_ = this

	localctx = NewParametro_llamada_referenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, NparserRULE_parametro_llamada_referencia)

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

	p.SetState(761)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(751)

			var _x = p.expresion(0)

			localctx.(*Parametro_llamada_referenciaContext).e = _x
		}

		temp := localctx.(*Parametro_llamada_referenciaContext).GetE().GetEx()
		fila := temp.(Ast.Abstracto).GetFila()
		columna := temp.(Ast.Abstracto).GetColumna()
		localctx.(*Parametro_llamada_referenciaContext).ex = simbolos.NewValor(localctx.(*Parametro_llamada_referenciaContext).GetE().GetEx(), Ast.VALOR, false, false, fila, columna)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(754)

			var _m = p.Match(NparserAMPERSAND)

			localctx.(*Parametro_llamada_referenciaContext)._AMPERSAND = _m
		}
		{
			p.SetState(755)
			p.Match(NparserMUT)
		}
		{
			p.SetState(756)

			var _m = p.Match(NparserID)

			localctx.(*Parametro_llamada_referenciaContext)._ID = _m
		}

		fila := (func() int {
			if localctx.(*Parametro_llamada_referenciaContext).Get_AMPERSAND() == nil {
				return 0
			} else {
				return localctx.(*Parametro_llamada_referenciaContext).Get_AMPERSAND().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Parametro_llamada_referenciaContext).Get_AMPERSAND() == nil {
				return 0
			} else {
				return localctx.(*Parametro_llamada_referenciaContext).Get_AMPERSAND().GetColumn()
			}
		}())
		id := expresiones.NewIdentificador((func() string {
			if localctx.(*Parametro_llamada_referenciaContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Parametro_llamada_referenciaContext).Get_ID().GetText()
			}
		}()), Ast.IDENTIFICADOR, fila, columna)
		localctx.(*Parametro_llamada_referenciaContext).ex = simbolos.NewValor(id, Ast.VALOR, true, true, fila, columna)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(758)

			var _m = p.Match(NparserAMPERSAND)

			localctx.(*Parametro_llamada_referenciaContext)._AMPERSAND = _m
		}
		{
			p.SetState(759)

			var _m = p.Match(NparserID)

			localctx.(*Parametro_llamada_referenciaContext)._ID = _m
		}

		fila := (func() int {
			if localctx.(*Parametro_llamada_referenciaContext).Get_AMPERSAND() == nil {
				return 0
			} else {
				return localctx.(*Parametro_llamada_referenciaContext).Get_AMPERSAND().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Parametro_llamada_referenciaContext).Get_AMPERSAND() == nil {
				return 0
			} else {
				return localctx.(*Parametro_llamada_referenciaContext).Get_AMPERSAND().GetColumn()
			}
		}())
		id := expresiones.NewIdentificador((func() string {
			if localctx.(*Parametro_llamada_referenciaContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*Parametro_llamada_referenciaContext).Get_ID().GetText()
			}
		}()), Ast.IDENTIFICADOR, fila, columna)
		localctx.(*Parametro_llamada_referenciaContext).ex = simbolos.NewValor(id, Ast.VALOR, true, false, fila, columna)

	}

	return localctx
}

// IElementos_vectorContext is an interface to support dynamic dispatch.
type IElementos_vectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLista_elementos returns the lista_elementos rule contexts.
	GetLista_elementos() IElementos_vectorContext

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// SetLista_elementos sets the lista_elementos rule contexts.
	SetLista_elementos(IElementos_vectorContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetList returns the list attribute.
	GetList() *arraylist.List

	// SetList sets the list attribute.
	SetList(*arraylist.List)

	// IsElementos_vectorContext differentiates from other interfaces.
	IsElementos_vectorContext()
}

type Elementos_vectorContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	list            *arraylist.List
	lista_elementos IElementos_vectorContext
	_expresion      IExpresionContext
}

func NewEmptyElementos_vectorContext() *Elementos_vectorContext {
	var p = new(Elementos_vectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_elementos_vector
	return p
}

func (*Elementos_vectorContext) IsElementos_vectorContext() {}

func NewElementos_vectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Elementos_vectorContext {
	var p = new(Elementos_vectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_elementos_vector

	return p
}

func (s *Elementos_vectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Elementos_vectorContext) GetLista_elementos() IElementos_vectorContext {
	return s.lista_elementos
}

func (s *Elementos_vectorContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *Elementos_vectorContext) SetLista_elementos(v IElementos_vectorContext) {
	s.lista_elementos = v
}

func (s *Elementos_vectorContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *Elementos_vectorContext) GetList() *arraylist.List { return s.list }

func (s *Elementos_vectorContext) SetList(v *arraylist.List) { s.list = v }

func (s *Elementos_vectorContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Elementos_vectorContext) COMA() antlr.TerminalNode {
	return s.GetToken(NparserCOMA, 0)
}

func (s *Elementos_vectorContext) Elementos_vector() IElementos_vectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementos_vectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementos_vectorContext)
}

func (s *Elementos_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Elementos_vectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Elementos_vectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterElementos_vector(s)
	}
}

func (s *Elementos_vectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitElementos_vector(s)
	}
}

func (p *Nparser) Elementos_vector() (localctx IElementos_vectorContext) {
	return p.elementos_vector(0)
}

func (p *Nparser) elementos_vector(_p int) (localctx IElementos_vectorContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewElementos_vectorContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IElementos_vectorContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 76
	p.EnterRecursionRule(localctx, 76, NparserRULE_elementos_vector, _p)
	localctx.(*Elementos_vectorContext).list = arraylist.New()

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
	{
		p.SetState(764)

		var _x = p.expresion(0)

		localctx.(*Elementos_vectorContext)._expresion = _x
	}

	localctx.(*Elementos_vectorContext).list.Add(localctx.(*Elementos_vectorContext).Get_expresion().GetEx())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(774)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewElementos_vectorContext(p, _parentctx, _parentState)
			localctx.(*Elementos_vectorContext).lista_elementos = _prevctx
			p.PushNewRecursionContext(localctx, _startState, NparserRULE_elementos_vector)
			p.SetState(767)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(768)
				p.Match(NparserCOMA)
			}
			{
				p.SetState(769)

				var _x = p.expresion(0)

				localctx.(*Elementos_vectorContext)._expresion = _x
			}

			localctx.(*Elementos_vectorContext).GetLista_elementos().GetList().Add(localctx.(*Elementos_vectorContext).Get_expresion().GetEx())
			localctx.(*Elementos_vectorContext).list = localctx.(*Elementos_vectorContext).GetLista_elementos().GetList()

		}
		p.SetState(776)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
	}

	return localctx
}

// IMetodos_iniciar_vectorContext is an interface to support dynamic dispatch.
type IMetodos_iniciar_vectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_VEC returns the _VEC token.
	Get_VEC() antlr.Token

	// Get_VEC_M returns the _VEC_M token.
	Get_VEC_M() antlr.Token

	// Set_VEC sets the _VEC token.
	Set_VEC(antlr.Token)

	// Set_VEC_M sets the _VEC_M token.
	Set_VEC_M(antlr.Token)

	// GetE returns the e rule contexts.
	GetE() IElementos_vectorContext

	// Get_elementos_vector returns the _elementos_vector rule contexts.
	Get_elementos_vector() IElementos_vectorContext

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// SetE sets the e rule contexts.
	SetE(IElementos_vectorContext)

	// Set_elementos_vector sets the _elementos_vector rule contexts.
	Set_elementos_vector(IElementos_vectorContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Expresion

	// SetEx sets the ex attribute.
	SetEx(Ast.Expresion)

	// IsMetodos_iniciar_vectorContext differentiates from other interfaces.
	IsMetodos_iniciar_vectorContext()
}

type Metodos_iniciar_vectorContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	ex                Ast.Expresion
	_VEC              antlr.Token
	_VEC_M            antlr.Token
	e                 IElementos_vectorContext
	_elementos_vector IElementos_vectorContext
	_expresion        IExpresionContext
}

func NewEmptyMetodos_iniciar_vectorContext() *Metodos_iniciar_vectorContext {
	var p = new(Metodos_iniciar_vectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_metodos_iniciar_vector
	return p
}

func (*Metodos_iniciar_vectorContext) IsMetodos_iniciar_vectorContext() {}

func NewMetodos_iniciar_vectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Metodos_iniciar_vectorContext {
	var p = new(Metodos_iniciar_vectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_metodos_iniciar_vector

	return p
}

func (s *Metodos_iniciar_vectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Metodos_iniciar_vectorContext) Get_VEC() antlr.Token { return s._VEC }

func (s *Metodos_iniciar_vectorContext) Get_VEC_M() antlr.Token { return s._VEC_M }

func (s *Metodos_iniciar_vectorContext) Set_VEC(v antlr.Token) { s._VEC = v }

func (s *Metodos_iniciar_vectorContext) Set_VEC_M(v antlr.Token) { s._VEC_M = v }

func (s *Metodos_iniciar_vectorContext) GetE() IElementos_vectorContext { return s.e }

func (s *Metodos_iniciar_vectorContext) Get_elementos_vector() IElementos_vectorContext {
	return s._elementos_vector
}

func (s *Metodos_iniciar_vectorContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *Metodos_iniciar_vectorContext) SetE(v IElementos_vectorContext) { s.e = v }

func (s *Metodos_iniciar_vectorContext) Set_elementos_vector(v IElementos_vectorContext) {
	s._elementos_vector = v
}

func (s *Metodos_iniciar_vectorContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *Metodos_iniciar_vectorContext) GetEx() Ast.Expresion { return s.ex }

func (s *Metodos_iniciar_vectorContext) SetEx(v Ast.Expresion) { s.ex = v }

func (s *Metodos_iniciar_vectorContext) VEC() antlr.TerminalNode {
	return s.GetToken(NparserVEC, 0)
}

func (s *Metodos_iniciar_vectorContext) DOBLE_DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(NparserDOBLE_DOSPUNTOS, 0)
}

func (s *Metodos_iniciar_vectorContext) NEW() antlr.TerminalNode {
	return s.GetToken(NparserNEW, 0)
}

func (s *Metodos_iniciar_vectorContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(NparserPAR_IZQ, 0)
}

func (s *Metodos_iniciar_vectorContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(NparserPAR_DER, 0)
}

func (s *Metodos_iniciar_vectorContext) VEC_M() antlr.TerminalNode {
	return s.GetToken(NparserVEC_M, 0)
}

func (s *Metodos_iniciar_vectorContext) NOT() antlr.TerminalNode {
	return s.GetToken(NparserNOT, 0)
}

func (s *Metodos_iniciar_vectorContext) CORCHETE_IZQ() antlr.TerminalNode {
	return s.GetToken(NparserCORCHETE_IZQ, 0)
}

func (s *Metodos_iniciar_vectorContext) CORCHETE_DER() antlr.TerminalNode {
	return s.GetToken(NparserCORCHETE_DER, 0)
}

func (s *Metodos_iniciar_vectorContext) Elementos_vector() IElementos_vectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementos_vectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementos_vectorContext)
}

func (s *Metodos_iniciar_vectorContext) WITH_CAPACITY() antlr.TerminalNode {
	return s.GetToken(NparserWITH_CAPACITY, 0)
}

func (s *Metodos_iniciar_vectorContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Metodos_iniciar_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Metodos_iniciar_vectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Metodos_iniciar_vectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterMetodos_iniciar_vector(s)
	}
}

func (s *Metodos_iniciar_vectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitMetodos_iniciar_vector(s)
	}
}

func (p *Nparser) Metodos_iniciar_vector() (localctx IMetodos_iniciar_vectorContext) {
	this := p
	_ = this

	localctx = NewMetodos_iniciar_vectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, NparserRULE_metodos_iniciar_vector)

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

	p.SetState(798)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(777)

			var _m = p.Match(NparserVEC)

			localctx.(*Metodos_iniciar_vectorContext)._VEC = _m
		}
		{
			p.SetState(778)
			p.Match(NparserDOBLE_DOSPUNTOS)
		}
		{
			p.SetState(779)
			p.Match(NparserNEW)
		}
		{
			p.SetState(780)
			p.Match(NparserPAR_IZQ)
		}
		{
			p.SetState(781)
			p.Match(NparserPAR_DER)
		}

		fila := (func() int {
			if localctx.(*Metodos_iniciar_vectorContext).Get_VEC() == nil {
				return 0
			} else {
				return localctx.(*Metodos_iniciar_vectorContext).Get_VEC().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Metodos_iniciar_vectorContext).Get_VEC() == nil {
				return 0
			} else {
				return localctx.(*Metodos_iniciar_vectorContext).Get_VEC().GetColumn()
			}
		}())
		vacio := true
		listaTemp := arraylist.New()
		usize := expresiones.NewPrimitivo(0, Ast.USIZE, fila, columna)
		localctx.(*Metodos_iniciar_vectorContext).ex = expresiones.NewVector(Ast.VECTOR, listaTemp, Ast.INDEFINIDO,
			usize, false, false, vacio, fila, columna)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(783)

			var _m = p.Match(NparserVEC_M)

			localctx.(*Metodos_iniciar_vectorContext)._VEC_M = _m
		}
		{
			p.SetState(784)
			p.Match(NparserNOT)
		}
		{
			p.SetState(785)
			p.Match(NparserCORCHETE_IZQ)
		}
		{
			p.SetState(786)

			var _x = p.elementos_vector(0)

			localctx.(*Metodos_iniciar_vectorContext).e = _x
			localctx.(*Metodos_iniciar_vectorContext)._elementos_vector = _x
		}
		{
			p.SetState(787)
			p.Match(NparserCORCHETE_DER)
		}

		fila := (func() int {
			if localctx.(*Metodos_iniciar_vectorContext).Get_VEC_M() == nil {
				return 0
			} else {
				return localctx.(*Metodos_iniciar_vectorContext).Get_VEC_M().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Metodos_iniciar_vectorContext).Get_VEC_M() == nil {
				return 0
			} else {
				return localctx.(*Metodos_iniciar_vectorContext).Get_VEC_M().GetColumn()
			}
		}())
		vacio := true
		listaTemp := localctx.(*Metodos_iniciar_vectorContext).GetE().GetList()
		if listaTemp.Len() > 0 {
			vacio = false
		}
		localctx.(*Metodos_iniciar_vectorContext).ex = expresiones.NewVector(Ast.VECTOR, localctx.(*Metodos_iniciar_vectorContext).Get_elementos_vector().GetList(), Ast.INDEFINIDO,
			Ast.TipoRetornado{Tipo: Ast.LIBRE, Valor: true}, false, false, vacio, fila, columna)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(790)

			var _m = p.Match(NparserVEC)

			localctx.(*Metodos_iniciar_vectorContext)._VEC = _m
		}
		{
			p.SetState(791)
			p.Match(NparserDOBLE_DOSPUNTOS)
		}
		{
			p.SetState(792)
			p.Match(NparserWITH_CAPACITY)
		}
		{
			p.SetState(793)
			p.Match(NparserPAR_IZQ)
		}
		{
			p.SetState(794)

			var _x = p.expresion(0)

			localctx.(*Metodos_iniciar_vectorContext)._expresion = _x
		}
		{
			p.SetState(795)
			p.Match(NparserPAR_DER)
		}

		fila := (func() int {
			if localctx.(*Metodos_iniciar_vectorContext).Get_VEC() == nil {
				return 0
			} else {
				return localctx.(*Metodos_iniciar_vectorContext).Get_VEC().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Metodos_iniciar_vectorContext).Get_VEC() == nil {
				return 0
			} else {
				return localctx.(*Metodos_iniciar_vectorContext).Get_VEC().GetColumn()
			}
		}())
		vacio := true
		listaTemp := arraylist.New()
		localctx.(*Metodos_iniciar_vectorContext).ex = expresiones.NewVector(Ast.VECTOR, listaTemp, Ast.INDEFINIDO,
			localctx.(*Metodos_iniciar_vectorContext).Get_expresion().GetEx(), false, false, vacio, fila, columna)

	}

	return localctx
}

func (p *Nparser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 7:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	case 19:
		var t *Case_matchContext = nil
		if localctx != nil {
			t = localctx.(*Case_matchContext)
		}
		return p.Case_match_Sempred(t, predIndex)

	case 23:
		var t *Case_match_expContext = nil
		if localctx != nil {
			t = localctx.(*Case_match_expContext)
		}
		return p.Case_match_exp_Sempred(t, predIndex)

	case 31:
		var t *ElementosPrintContext = nil
		if localctx != nil {
			t = localctx.(*ElementosPrintContext)
		}
		return p.ElementosPrint_Sempred(t, predIndex)

	case 33:
		var t *Parametros_funcionContext = nil
		if localctx != nil {
			t = localctx.(*Parametros_funcionContext)
		}
		return p.Parametros_funcion_Sempred(t, predIndex)

	case 36:
		var t *Parametros_llamadaContext = nil
		if localctx != nil {
			t = localctx.(*Parametros_llamadaContext)
		}
		return p.Parametros_llamada_Sempred(t, predIndex)

	case 38:
		var t *Elementos_vectorContext = nil
		if localctx != nil {
			t = localctx.(*Elementos_vectorContext)
		}
		return p.Elementos_vector_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Nparser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 18)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Nparser) Case_match_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 9:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Nparser) Case_match_exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 10:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Nparser) ElementosPrint_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 11:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Nparser) Parametros_funcion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 12:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Nparser) Parametros_llamada_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 13:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Nparser) Elementos_vector_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 14:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
