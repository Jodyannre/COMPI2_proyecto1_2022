// Code generated from c:\Users\Joddie\Documents\GitHub\COMPI2_proyecto1_2022\Back\parser\Nlexer.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 91, 599,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4,
	81, 9, 81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85, 4, 86,
	9, 86, 4, 87, 9, 87, 4, 88, 9, 88, 4, 89, 9, 89, 4, 90, 9, 90, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34,
	3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37,
	3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3,
	39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42,
	3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3,
	44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45,
	3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 6, 47, 437, 10, 47, 13,
	47, 14, 47, 438, 3, 48, 6, 48, 442, 10, 48, 13, 48, 14, 48, 443, 3, 48,
	3, 48, 6, 48, 448, 10, 48, 13, 48, 14, 48, 449, 3, 49, 3, 49, 7, 49, 454,
	10, 49, 12, 49, 14, 49, 457, 11, 49, 3, 50, 5, 50, 460, 10, 50, 3, 50,
	7, 50, 463, 10, 50, 12, 50, 14, 50, 466, 11, 50, 3, 51, 3, 51, 3, 52, 3,
	52, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 56, 3, 56,
	3, 56, 3, 57, 3, 57, 3, 57, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 60, 3,
	60, 3, 61, 3, 61, 3, 62, 3, 62, 3, 63, 3, 63, 3, 63, 3, 64, 3, 64, 3, 65,
	3, 65, 3, 65, 3, 66, 3, 66, 3, 67, 3, 67, 3, 67, 3, 68, 3, 68, 3, 68, 3,
	69, 3, 69, 3, 69, 3, 70, 3, 70, 3, 71, 3, 71, 3, 72, 3, 72, 3, 73, 3, 73,
	3, 74, 3, 74, 3, 74, 3, 75, 3, 75, 3, 76, 3, 76, 3, 77, 3, 77, 3, 78, 3,
	78, 3, 79, 3, 79, 3, 80, 3, 80, 3, 81, 3, 81, 3, 82, 3, 82, 3, 83, 3, 83,
	3, 84, 3, 84, 3, 85, 3, 85, 7, 85, 549, 10, 85, 12, 85, 14, 85, 552, 11,
	85, 3, 85, 3, 85, 3, 86, 3, 86, 3, 86, 5, 86, 559, 10, 86, 3, 87, 3, 87,
	3, 87, 5, 87, 564, 10, 87, 3, 87, 3, 87, 3, 88, 6, 88, 569, 10, 88, 13,
	88, 14, 88, 570, 3, 88, 3, 88, 3, 89, 3, 89, 3, 89, 3, 89, 7, 89, 579,
	10, 89, 12, 89, 14, 89, 582, 11, 89, 3, 89, 3, 89, 3, 89, 3, 89, 3, 89,
	3, 90, 3, 90, 3, 90, 3, 90, 7, 90, 593, 10, 90, 12, 90, 14, 90, 596, 11,
	90, 3, 90, 3, 90, 4, 550, 580, 2, 91, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26,
	51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35,
	69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44,
	87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97, 50, 99, 51, 101, 52, 103, 53,
	105, 54, 107, 55, 109, 56, 111, 57, 113, 58, 115, 59, 117, 60, 119, 61,
	121, 62, 123, 63, 125, 64, 127, 65, 129, 66, 131, 67, 133, 68, 135, 69,
	137, 70, 139, 71, 141, 72, 143, 73, 145, 74, 147, 75, 149, 76, 151, 77,
	153, 78, 155, 79, 157, 80, 159, 81, 161, 82, 163, 83, 165, 84, 167, 85,
	169, 86, 171, 87, 173, 88, 175, 89, 177, 90, 179, 91, 3, 2, 9, 3, 2, 50,
	59, 3, 2, 67, 92, 5, 2, 50, 59, 67, 92, 99, 124, 6, 2, 50, 59, 67, 92,
	97, 97, 99, 124, 10, 2, 35, 35, 37, 40, 42, 47, 49, 49, 60, 66, 96, 97,
	125, 125, 127, 127, 6, 2, 11, 12, 15, 15, 34, 34, 94, 94, 4, 2, 12, 12,
	15, 15, 2, 610, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2,
	2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2,
	2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2,
	2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3,
	2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47,
	3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2,
	55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2,
	2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2,
	2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2,
	2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3,
	2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93,
	3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2,
	101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2,
	2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 3, 2, 2, 2, 2, 115,
	3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2, 2, 2, 2, 121, 3, 2, 2, 2,
	2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 2, 127, 3, 2, 2, 2, 2, 129, 3,
	2, 2, 2, 2, 131, 3, 2, 2, 2, 2, 133, 3, 2, 2, 2, 2, 135, 3, 2, 2, 2, 2,
	137, 3, 2, 2, 2, 2, 139, 3, 2, 2, 2, 2, 141, 3, 2, 2, 2, 2, 143, 3, 2,
	2, 2, 2, 145, 3, 2, 2, 2, 2, 147, 3, 2, 2, 2, 2, 149, 3, 2, 2, 2, 2, 151,
	3, 2, 2, 2, 2, 153, 3, 2, 2, 2, 2, 155, 3, 2, 2, 2, 2, 157, 3, 2, 2, 2,
	2, 159, 3, 2, 2, 2, 2, 161, 3, 2, 2, 2, 2, 163, 3, 2, 2, 2, 2, 165, 3,
	2, 2, 2, 2, 167, 3, 2, 2, 2, 2, 169, 3, 2, 2, 2, 2, 171, 3, 2, 2, 2, 2,
	173, 3, 2, 2, 2, 2, 175, 3, 2, 2, 2, 2, 177, 3, 2, 2, 2, 2, 179, 3, 2,
	2, 2, 3, 181, 3, 2, 2, 2, 5, 186, 3, 2, 2, 2, 7, 191, 3, 2, 2, 2, 9, 195,
	3, 2, 2, 2, 11, 199, 3, 2, 2, 2, 13, 204, 3, 2, 2, 2, 15, 211, 3, 2, 2,
	2, 17, 217, 3, 2, 2, 2, 19, 222, 3, 2, 2, 2, 21, 227, 3, 2, 2, 2, 23, 231,
	3, 2, 2, 2, 25, 234, 3, 2, 2, 2, 27, 238, 3, 2, 2, 2, 29, 242, 3, 2, 2,
	2, 31, 246, 3, 2, 2, 2, 33, 250, 3, 2, 2, 2, 35, 257, 3, 2, 2, 2, 37, 267,
	3, 2, 2, 2, 39, 276, 3, 2, 2, 2, 41, 281, 3, 2, 2, 2, 43, 287, 3, 2, 2,
	2, 45, 296, 3, 2, 2, 2, 47, 299, 3, 2, 2, 2, 49, 303, 3, 2, 2, 2, 51, 308,
	3, 2, 2, 2, 53, 314, 3, 2, 2, 2, 55, 318, 3, 2, 2, 2, 57, 322, 3, 2, 2,
	2, 59, 327, 3, 2, 2, 2, 61, 334, 3, 2, 2, 2, 63, 343, 3, 2, 2, 2, 65, 350,
	3, 2, 2, 2, 67, 359, 3, 2, 2, 2, 69, 373, 3, 2, 2, 2, 71, 376, 3, 2, 2,
	2, 73, 381, 3, 2, 2, 2, 75, 387, 3, 2, 2, 2, 77, 392, 3, 2, 2, 2, 79, 398,
	3, 2, 2, 2, 81, 402, 3, 2, 2, 2, 83, 405, 3, 2, 2, 2, 85, 412, 3, 2, 2,
	2, 87, 418, 3, 2, 2, 2, 89, 427, 3, 2, 2, 2, 91, 431, 3, 2, 2, 2, 93, 436,
	3, 2, 2, 2, 95, 441, 3, 2, 2, 2, 97, 451, 3, 2, 2, 2, 99, 459, 3, 2, 2,
	2, 101, 467, 3, 2, 2, 2, 103, 469, 3, 2, 2, 2, 105, 471, 3, 2, 2, 2, 107,
	474, 3, 2, 2, 2, 109, 476, 3, 2, 2, 2, 111, 479, 3, 2, 2, 2, 113, 482,
	3, 2, 2, 2, 115, 485, 3, 2, 2, 2, 117, 487, 3, 2, 2, 2, 119, 490, 3, 2,
	2, 2, 121, 492, 3, 2, 2, 2, 123, 494, 3, 2, 2, 2, 125, 496, 3, 2, 2, 2,
	127, 499, 3, 2, 2, 2, 129, 501, 3, 2, 2, 2, 131, 504, 3, 2, 2, 2, 133,
	506, 3, 2, 2, 2, 135, 509, 3, 2, 2, 2, 137, 512, 3, 2, 2, 2, 139, 515,
	3, 2, 2, 2, 141, 517, 3, 2, 2, 2, 143, 519, 3, 2, 2, 2, 145, 521, 3, 2,
	2, 2, 147, 523, 3, 2, 2, 2, 149, 526, 3, 2, 2, 2, 151, 528, 3, 2, 2, 2,
	153, 530, 3, 2, 2, 2, 155, 532, 3, 2, 2, 2, 157, 534, 3, 2, 2, 2, 159,
	536, 3, 2, 2, 2, 161, 538, 3, 2, 2, 2, 163, 540, 3, 2, 2, 2, 165, 542,
	3, 2, 2, 2, 167, 544, 3, 2, 2, 2, 169, 546, 3, 2, 2, 2, 171, 558, 3, 2,
	2, 2, 173, 560, 3, 2, 2, 2, 175, 568, 3, 2, 2, 2, 177, 574, 3, 2, 2, 2,
	179, 588, 3, 2, 2, 2, 181, 182, 7, 100, 2, 2, 182, 183, 7, 113, 2, 2, 183,
	184, 7, 113, 2, 2, 184, 185, 7, 110, 2, 2, 185, 4, 3, 2, 2, 2, 186, 187,
	7, 101, 2, 2, 187, 188, 7, 106, 2, 2, 188, 189, 7, 99, 2, 2, 189, 190,
	7, 116, 2, 2, 190, 6, 3, 2, 2, 2, 191, 192, 7, 104, 2, 2, 192, 193, 7,
	56, 2, 2, 193, 194, 7, 54, 2, 2, 194, 8, 3, 2, 2, 2, 195, 196, 7, 107,
	2, 2, 196, 197, 7, 56, 2, 2, 197, 198, 7, 54, 2, 2, 198, 10, 3, 2, 2, 2,
	199, 200, 7, 40, 2, 2, 200, 201, 7, 117, 2, 2, 201, 202, 7, 118, 2, 2,
	202, 203, 7, 116, 2, 2, 203, 12, 3, 2, 2, 2, 204, 205, 7, 85, 2, 2, 205,
	206, 7, 118, 2, 2, 206, 207, 7, 116, 2, 2, 207, 208, 7, 107, 2, 2, 208,
	209, 7, 112, 2, 2, 209, 210, 7, 105, 2, 2, 210, 14, 3, 2, 2, 2, 211, 212,
	7, 119, 2, 2, 212, 213, 7, 117, 2, 2, 213, 214, 7, 107, 2, 2, 214, 215,
	7, 124, 2, 2, 215, 216, 7, 103, 2, 2, 216, 16, 3, 2, 2, 2, 217, 218, 7,
	111, 2, 2, 218, 219, 7, 99, 2, 2, 219, 220, 7, 107, 2, 2, 220, 221, 7,
	112, 2, 2, 221, 18, 3, 2, 2, 2, 222, 223, 7, 114, 2, 2, 223, 224, 7, 113,
	2, 2, 224, 225, 7, 121, 2, 2, 225, 226, 7, 104, 2, 2, 226, 20, 3, 2, 2,
	2, 227, 228, 7, 114, 2, 2, 228, 229, 7, 113, 2, 2, 229, 230, 7, 121, 2,
	2, 230, 22, 3, 2, 2, 2, 231, 232, 7, 99, 2, 2, 232, 233, 7, 117, 2, 2,
	233, 24, 3, 2, 2, 2, 234, 235, 7, 88, 2, 2, 235, 236, 7, 103, 2, 2, 236,
	237, 7, 101, 2, 2, 237, 26, 3, 2, 2, 2, 238, 239, 7, 120, 2, 2, 239, 240,
	7, 103, 2, 2, 240, 241, 7, 101, 2, 2, 241, 28, 3, 2, 2, 2, 242, 243, 7,
	111, 2, 2, 243, 244, 7, 119, 2, 2, 244, 245, 7, 118, 2, 2, 245, 30, 3,
	2, 2, 2, 246, 247, 7, 110, 2, 2, 247, 248, 7, 103, 2, 2, 248, 249, 7, 118,
	2, 2, 249, 32, 3, 2, 2, 2, 250, 251, 7, 117, 2, 2, 251, 252, 7, 118, 2,
	2, 252, 253, 7, 116, 2, 2, 253, 254, 7, 119, 2, 2, 254, 255, 7, 101, 2,
	2, 255, 256, 7, 118, 2, 2, 256, 34, 3, 2, 2, 2, 257, 258, 7, 118, 2, 2,
	258, 259, 7, 113, 2, 2, 259, 260, 7, 97, 2, 2, 260, 261, 7, 117, 2, 2,
	261, 262, 7, 118, 2, 2, 262, 263, 7, 116, 2, 2, 263, 264, 7, 107, 2, 2,
	264, 265, 7, 112, 2, 2, 265, 266, 7, 105, 2, 2, 266, 36, 3, 2, 2, 2, 267,
	268, 7, 118, 2, 2, 268, 269, 7, 113, 2, 2, 269, 270, 7, 97, 2, 2, 270,
	271, 7, 113, 2, 2, 271, 272, 7, 121, 2, 2, 272, 273, 7, 112, 2, 2, 273,
	274, 7, 103, 2, 2, 274, 275, 7, 102, 2, 2, 275, 38, 3, 2, 2, 2, 276, 277,
	7, 118, 2, 2, 277, 278, 7, 116, 2, 2, 278, 279, 7, 119, 2, 2, 279, 280,
	7, 103, 2, 2, 280, 40, 3, 2, 2, 2, 281, 282, 7, 104, 2, 2, 282, 283, 7,
	99, 2, 2, 283, 284, 7, 110, 2, 2, 284, 285, 7, 117, 2, 2, 285, 286, 7,
	103, 2, 2, 286, 42, 3, 2, 2, 2, 287, 288, 7, 114, 2, 2, 288, 289, 7, 116,
	2, 2, 289, 290, 7, 107, 2, 2, 290, 291, 7, 112, 2, 2, 291, 292, 7, 118,
	2, 2, 292, 293, 7, 110, 2, 2, 293, 294, 7, 112, 2, 2, 294, 295, 7, 35,
	2, 2, 295, 44, 3, 2, 2, 2, 296, 297, 7, 104, 2, 2, 297, 298, 7, 112, 2,
	2, 298, 46, 3, 2, 2, 2, 299, 300, 7, 99, 2, 2, 300, 301, 7, 100, 2, 2,
	301, 302, 7, 117, 2, 2, 302, 48, 3, 2, 2, 2, 303, 304, 7, 117, 2, 2, 304,
	305, 7, 115, 2, 2, 305, 306, 7, 116, 2, 2, 306, 307, 7, 118, 2, 2, 307,
	50, 3, 2, 2, 2, 308, 309, 7, 101, 2, 2, 309, 310, 7, 110, 2, 2, 310, 311,
	7, 113, 2, 2, 311, 312, 7, 112, 2, 2, 312, 313, 7, 103, 2, 2, 313, 52,
	3, 2, 2, 2, 314, 315, 7, 112, 2, 2, 315, 316, 7, 103, 2, 2, 316, 317, 7,
	121, 2, 2, 317, 54, 3, 2, 2, 2, 318, 319, 7, 110, 2, 2, 319, 320, 7, 103,
	2, 2, 320, 321, 7, 112, 2, 2, 321, 56, 3, 2, 2, 2, 322, 323, 7, 114, 2,
	2, 323, 324, 7, 119, 2, 2, 324, 325, 7, 117, 2, 2, 325, 326, 7, 106, 2,
	2, 326, 58, 3, 2, 2, 2, 327, 328, 7, 116, 2, 2, 328, 329, 7, 103, 2, 2,
	329, 330, 7, 111, 2, 2, 330, 331, 7, 113, 2, 2, 331, 332, 7, 120, 2, 2,
	332, 333, 7, 103, 2, 2, 333, 60, 3, 2, 2, 2, 334, 335, 7, 101, 2, 2, 335,
	336, 7, 113, 2, 2, 336, 337, 7, 112, 2, 2, 337, 338, 7, 118, 2, 2, 338,
	339, 7, 99, 2, 2, 339, 340, 7, 107, 2, 2, 340, 341, 7, 112, 2, 2, 341,
	342, 7, 117, 2, 2, 342, 62, 3, 2, 2, 2, 343, 344, 7, 107, 2, 2, 344, 345,
	7, 112, 2, 2, 345, 346, 7, 117, 2, 2, 346, 347, 7, 103, 2, 2, 347, 348,
	7, 116, 2, 2, 348, 349, 7, 118, 2, 2, 349, 64, 3, 2, 2, 2, 350, 351, 7,
	101, 2, 2, 351, 352, 7, 99, 2, 2, 352, 353, 7, 114, 2, 2, 353, 354, 7,
	99, 2, 2, 354, 355, 7, 101, 2, 2, 355, 356, 7, 107, 2, 2, 356, 357, 7,
	118, 2, 2, 357, 358, 7, 123, 2, 2, 358, 66, 3, 2, 2, 2, 359, 360, 7, 121,
	2, 2, 360, 361, 7, 107, 2, 2, 361, 362, 7, 118, 2, 2, 362, 363, 7, 106,
	2, 2, 363, 364, 7, 97, 2, 2, 364, 365, 7, 101, 2, 2, 365, 366, 7, 99, 2,
	2, 366, 367, 7, 114, 2, 2, 367, 368, 7, 99, 2, 2, 368, 369, 7, 101, 2,
	2, 369, 370, 7, 107, 2, 2, 370, 371, 7, 118, 2, 2, 371, 372, 7, 123, 2,
	2, 372, 68, 3, 2, 2, 2, 373, 374, 7, 107, 2, 2, 374, 375, 7, 104, 2, 2,
	375, 70, 3, 2, 2, 2, 376, 377, 7, 103, 2, 2, 377, 378, 7, 110, 2, 2, 378,
	379, 7, 117, 2, 2, 379, 380, 7, 103, 2, 2, 380, 72, 3, 2, 2, 2, 381, 382,
	7, 111, 2, 2, 382, 383, 7, 99, 2, 2, 383, 384, 7, 118, 2, 2, 384, 385,
	7, 101, 2, 2, 385, 386, 7, 106, 2, 2, 386, 74, 3, 2, 2, 2, 387, 388, 7,
	110, 2, 2, 388, 389, 7, 113, 2, 2, 389, 390, 7, 113, 2, 2, 390, 391, 7,
	114, 2, 2, 391, 76, 3, 2, 2, 2, 392, 393, 7, 121, 2, 2, 393, 394, 7, 106,
	2, 2, 394, 395, 7, 107, 2, 2, 395, 396, 7, 110, 2, 2, 396, 397, 7, 103,
	2, 2, 397, 78, 3, 2, 2, 2, 398, 399, 7, 104, 2, 2, 399, 400, 7, 113, 2,
	2, 400, 401, 7, 116, 2, 2, 401, 80, 3, 2, 2, 2, 402, 403, 7, 107, 2, 2,
	403, 404, 7, 112, 2, 2, 404, 82, 3, 2, 2, 2, 405, 406, 7, 116, 2, 2, 406,
	407, 7, 103, 2, 2, 407, 408, 7, 118, 2, 2, 408, 409, 7, 119, 2, 2, 409,
	410, 7, 116, 2, 2, 410, 411, 7, 112, 2, 2, 411, 84, 3, 2, 2, 2, 412, 413,
	7, 100, 2, 2, 413, 414, 7, 116, 2, 2, 414, 415, 7, 103, 2, 2, 415, 416,
	7, 99, 2, 2, 416, 417, 7, 109, 2, 2, 417, 86, 3, 2, 2, 2, 418, 419, 7,
	101, 2, 2, 419, 420, 7, 113, 2, 2, 420, 421, 7, 112, 2, 2, 421, 422, 7,
	118, 2, 2, 422, 423, 7, 107, 2, 2, 423, 424, 7, 112, 2, 2, 424, 425, 7,
	119, 2, 2, 425, 426, 7, 103, 2, 2, 426, 88, 3, 2, 2, 2, 427, 428, 7, 111,
	2, 2, 428, 429, 7, 113, 2, 2, 429, 430, 7, 102, 2, 2, 430, 90, 3, 2, 2,
	2, 431, 432, 7, 114, 2, 2, 432, 433, 7, 119, 2, 2, 433, 434, 7, 100, 2,
	2, 434, 92, 3, 2, 2, 2, 435, 437, 9, 2, 2, 2, 436, 435, 3, 2, 2, 2, 437,
	438, 3, 2, 2, 2, 438, 436, 3, 2, 2, 2, 438, 439, 3, 2, 2, 2, 439, 94, 3,
	2, 2, 2, 440, 442, 9, 2, 2, 2, 441, 440, 3, 2, 2, 2, 442, 443, 3, 2, 2,
	2, 443, 441, 3, 2, 2, 2, 443, 444, 3, 2, 2, 2, 444, 445, 3, 2, 2, 2, 445,
	447, 7, 48, 2, 2, 446, 448, 9, 2, 2, 2, 447, 446, 3, 2, 2, 2, 448, 449,
	3, 2, 2, 2, 449, 447, 3, 2, 2, 2, 449, 450, 3, 2, 2, 2, 450, 96, 3, 2,
	2, 2, 451, 455, 9, 3, 2, 2, 452, 454, 9, 4, 2, 2, 453, 452, 3, 2, 2, 2,
	454, 457, 3, 2, 2, 2, 455, 453, 3, 2, 2, 2, 455, 456, 3, 2, 2, 2, 456,
	98, 3, 2, 2, 2, 457, 455, 3, 2, 2, 2, 458, 460, 9, 4, 2, 2, 459, 458, 3,
	2, 2, 2, 460, 464, 3, 2, 2, 2, 461, 463, 9, 5, 2, 2, 462, 461, 3, 2, 2,
	2, 463, 466, 3, 2, 2, 2, 464, 462, 3, 2, 2, 2, 464, 465, 3, 2, 2, 2, 465,
	100, 3, 2, 2, 2, 466, 464, 3, 2, 2, 2, 467, 468, 7, 97, 2, 2, 468, 102,
	3, 2, 2, 2, 469, 470, 7, 126, 2, 2, 470, 104, 3, 2, 2, 2, 471, 472, 7,
	126, 2, 2, 472, 473, 7, 126, 2, 2, 473, 106, 3, 2, 2, 2, 474, 475, 7, 40,
	2, 2, 475, 108, 3, 2, 2, 2, 476, 477, 7, 40, 2, 2, 477, 478, 7, 40, 2,
	2, 478, 110, 3, 2, 2, 2, 479, 480, 7, 60, 2, 2, 480, 481, 7, 65, 2, 2,
	481, 112, 3, 2, 2, 2, 482, 483, 7, 60, 2, 2, 483, 484, 7, 60, 2, 2, 484,
	114, 3, 2, 2, 2, 485, 486, 7, 60, 2, 2, 486, 116, 3, 2, 2, 2, 487, 488,
	7, 48, 2, 2, 488, 489, 7, 48, 2, 2, 489, 118, 3, 2, 2, 2, 490, 491, 7,
	48, 2, 2, 491, 120, 3, 2, 2, 2, 492, 493, 7, 46, 2, 2, 493, 122, 3, 2,
	2, 2, 494, 495, 7, 61, 2, 2, 495, 124, 3, 2, 2, 2, 496, 497, 7, 64, 2,
	2, 497, 498, 7, 63, 2, 2, 498, 126, 3, 2, 2, 2, 499, 500, 7, 64, 2, 2,
	500, 128, 3, 2, 2, 2, 501, 502, 7, 62, 2, 2, 502, 503, 7, 63, 2, 2, 503,
	130, 3, 2, 2, 2, 504, 505, 7, 62, 2, 2, 505, 132, 3, 2, 2, 2, 506, 507,
	7, 63, 2, 2, 507, 508, 7, 63, 2, 2, 508, 134, 3, 2, 2, 2, 509, 510, 7,
	63, 2, 2, 510, 511, 7, 64, 2, 2, 511, 136, 3, 2, 2, 2, 512, 513, 7, 35,
	2, 2, 513, 514, 7, 63, 2, 2, 514, 138, 3, 2, 2, 2, 515, 516, 7, 63, 2,
	2, 516, 140, 3, 2, 2, 2, 517, 518, 7, 39, 2, 2, 518, 142, 3, 2, 2, 2, 519,
	520, 7, 44, 2, 2, 520, 144, 3, 2, 2, 2, 521, 522, 7, 49, 2, 2, 522, 146,
	3, 2, 2, 2, 523, 524, 7, 47, 2, 2, 524, 525, 7, 64, 2, 2, 525, 148, 3,
	2, 2, 2, 526, 527, 7, 47, 2, 2, 527, 150, 3, 2, 2, 2, 528, 529, 7, 45,
	2, 2, 529, 152, 3, 2, 2, 2, 530, 531, 7, 35, 2, 2, 531, 154, 3, 2, 2, 2,
	532, 533, 7, 65, 2, 2, 533, 156, 3, 2, 2, 2, 534, 535, 7, 42, 2, 2, 535,
	158, 3, 2, 2, 2, 536, 537, 7, 43, 2, 2, 537, 160, 3, 2, 2, 2, 538, 539,
	7, 125, 2, 2, 539, 162, 3, 2, 2, 2, 540, 541, 7, 127, 2, 2, 541, 164, 3,
	2, 2, 2, 542, 543, 7, 93, 2, 2, 543, 166, 3, 2, 2, 2, 544, 545, 7, 95,
	2, 2, 545, 168, 3, 2, 2, 2, 546, 550, 7, 36, 2, 2, 547, 549, 11, 2, 2,
	2, 548, 547, 3, 2, 2, 2, 549, 552, 3, 2, 2, 2, 550, 551, 3, 2, 2, 2, 550,
	548, 3, 2, 2, 2, 551, 553, 3, 2, 2, 2, 552, 550, 3, 2, 2, 2, 553, 554,
	7, 36, 2, 2, 554, 170, 3, 2, 2, 2, 555, 559, 9, 6, 2, 2, 556, 559, 5, 167,
	84, 2, 557, 559, 5, 165, 83, 2, 558, 555, 3, 2, 2, 2, 558, 556, 3, 2, 2,
	2, 558, 557, 3, 2, 2, 2, 559, 172, 3, 2, 2, 2, 560, 563, 7, 41, 2, 2, 561,
	564, 5, 171, 86, 2, 562, 564, 9, 4, 2, 2, 563, 561, 3, 2, 2, 2, 563, 562,
	3, 2, 2, 2, 564, 565, 3, 2, 2, 2, 565, 566, 7, 41, 2, 2, 566, 174, 3, 2,
	2, 2, 567, 569, 9, 7, 2, 2, 568, 567, 3, 2, 2, 2, 569, 570, 3, 2, 2, 2,
	570, 568, 3, 2, 2, 2, 570, 571, 3, 2, 2, 2, 571, 572, 3, 2, 2, 2, 572,
	573, 8, 88, 2, 2, 573, 176, 3, 2, 2, 2, 574, 575, 7, 49, 2, 2, 575, 576,
	7, 44, 2, 2, 576, 580, 3, 2, 2, 2, 577, 579, 11, 2, 2, 2, 578, 577, 3,
	2, 2, 2, 579, 582, 3, 2, 2, 2, 580, 581, 3, 2, 2, 2, 580, 578, 3, 2, 2,
	2, 581, 583, 3, 2, 2, 2, 582, 580, 3, 2, 2, 2, 583, 584, 7, 44, 2, 2, 584,
	585, 7, 49, 2, 2, 585, 586, 3, 2, 2, 2, 586, 587, 8, 89, 2, 2, 587, 178,
	3, 2, 2, 2, 588, 589, 7, 49, 2, 2, 589, 590, 7, 49, 2, 2, 590, 594, 3,
	2, 2, 2, 591, 593, 10, 8, 2, 2, 592, 591, 3, 2, 2, 2, 593, 596, 3, 2, 2,
	2, 594, 592, 3, 2, 2, 2, 594, 595, 3, 2, 2, 2, 595, 597, 3, 2, 2, 2, 596,
	594, 3, 2, 2, 2, 597, 598, 8, 90, 2, 2, 598, 180, 3, 2, 2, 2, 17, 2, 438,
	443, 449, 453, 455, 459, 462, 464, 550, 558, 563, 570, 580, 594, 3, 8,
	2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
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

var lexerSymbolicNames = []string{
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

var lexerRuleNames = []string{
	"BOOL", "CHAR", "F64", "I64", "STR", "STRING", "USIZE", "MAIN", "POWF",
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

type Nlexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewNlexer(input antlr.CharStream) *Nlexer {

	l := new(Nlexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Nlexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Nlexer tokens.
const (
	NlexerBOOL            = 1
	NlexerCHAR            = 2
	NlexerF64             = 3
	NlexerI64             = 4
	NlexerSTR             = 5
	NlexerSTRING          = 6
	NlexerUSIZE           = 7
	NlexerMAIN            = 8
	NlexerPOWF            = 9
	NlexerPOW             = 10
	NlexerAS              = 11
	NlexerVEC             = 12
	NlexerVEC_M           = 13
	NlexerMUT             = 14
	NlexerLET             = 15
	NlexerSTRUCT          = 16
	NlexerTO_STRING       = 17
	NlexerTO_OWNED        = 18
	NlexerTRUE            = 19
	NlexerFALSE           = 20
	NlexerPRINT           = 21
	NlexerFN              = 22
	NlexerABS             = 23
	NlexerSQRT            = 24
	NlexerCLONE           = 25
	NlexerNEW             = 26
	NlexerLEN             = 27
	NlexerPUSH            = 28
	NlexerREMOVE          = 29
	NlexerCONTAINS        = 30
	NlexerINSERT          = 31
	NlexerCAPACITY        = 32
	NlexerWITH_CAPACITY   = 33
	NlexerIF              = 34
	NlexerELSE            = 35
	NlexerMATCH           = 36
	NlexerLOOP            = 37
	NlexerWHILE           = 38
	NlexerFOR             = 39
	NlexerIN              = 40
	NlexerRETURN          = 41
	NlexerBREAK           = 42
	NlexerCONTINUE        = 43
	NlexerMOD             = 44
	NlexerPUB             = 45
	NlexerNUMERO          = 46
	NlexerDECIMAL         = 47
	NlexerID_CAMEL        = 48
	NlexerID              = 49
	NlexerDEFAULT         = 50
	NlexerO               = 51
	NlexerOR              = 52
	NlexerAMPERSAND       = 53
	NlexerAND             = 54
	NlexerPRINT_OP_DEBUG  = 55
	NlexerDOBLE_DOSPUNTOS = 56
	NlexerDOSPUNTOS       = 57
	NlexerRANGO           = 58
	NlexerPUNTO           = 59
	NlexerCOMA            = 60
	NlexerPUNTOCOMA       = 61
	NlexerMAYOR_I         = 62
	NlexerMAYOR           = 63
	NlexerMENOR_I         = 64
	NlexerMENOR           = 65
	NlexerIGUALDAD        = 66
	NlexerCASE            = 67
	NlexerDISTINTO        = 68
	NlexerIGUAL           = 69
	NlexerMODULO          = 70
	NlexerMULTIPLICACION  = 71
	NlexerDIVISION        = 72
	NlexerFN_TIPO_RETORNO = 73
	NlexerRESTA           = 74
	NlexerSUMA            = 75
	NlexerNOT             = 76
	NlexerPREGUNTA        = 77
	NlexerPAR_IZQ         = 78
	NlexerPAR_DER         = 79
	NlexerLLAVE_IZQ       = 80
	NlexerLLAVE_DER       = 81
	NlexerCORCHETE_IZQ    = 82
	NlexerCORCHETE_DER    = 83
	NlexerCADENA          = 84
	NlexerASCII           = 85
	NlexerCARACTER        = 86
	NlexerWHITESPACE      = 87
	NlexerCOMMENT         = 88
	NlexerLINE_COMMENT    = 89
)
