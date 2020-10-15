// Code generated from JSPNL.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // JSPNL

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 43, 262,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 5, 2, 48, 10, 2, 3, 2, 7, 2, 51, 10, 2, 12, 2, 14, 2, 54, 11, 2,
	3, 3, 3, 3, 5, 3, 58, 10, 3, 3, 4, 3, 4, 3, 4, 5, 4, 63, 10, 4, 3, 5, 3,
	5, 3, 5, 5, 5, 68, 10, 5, 3, 5, 3, 5, 3, 5, 5, 5, 73, 10, 5, 3, 5, 5, 5,
	76, 10, 5, 3, 5, 3, 5, 3, 5, 5, 5, 81, 10, 5, 3, 5, 5, 5, 84, 10, 5, 5,
	5, 86, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 93, 10, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 7, 9, 106, 10, 9,
	12, 9, 14, 9, 109, 11, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	12, 3, 12, 3, 13, 7, 13, 120, 10, 13, 12, 13, 14, 13, 123, 11, 13, 3, 13,
	3, 13, 5, 13, 127, 10, 13, 3, 13, 3, 13, 5, 13, 131, 10, 13, 7, 13, 133,
	10, 13, 12, 13, 14, 13, 136, 11, 13, 3, 13, 3, 13, 3, 14, 3, 14, 5, 14,
	142, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 5, 15, 154, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 190, 10, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 222,
	10, 16, 12, 16, 14, 16, 225, 11, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 18, 3, 18, 5, 18, 234, 10, 18, 3, 19, 3, 19, 3, 19, 7, 19, 239, 10,
	19, 12, 19, 14, 19, 242, 11, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3,
	22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	5, 23, 260, 10, 23, 3, 23, 2, 3, 30, 24, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 2, 8, 3, 2, 5, 8, 3,
	2, 17, 19, 3, 2, 20, 23, 3, 2, 18, 19, 3, 2, 24, 27, 3, 2, 28, 29, 2, 277,
	2, 52, 3, 2, 2, 2, 4, 57, 3, 2, 2, 2, 6, 62, 3, 2, 2, 2, 8, 85, 3, 2, 2,
	2, 10, 87, 3, 2, 2, 2, 12, 94, 3, 2, 2, 2, 14, 98, 3, 2, 2, 2, 16, 102,
	3, 2, 2, 2, 18, 110, 3, 2, 2, 2, 20, 112, 3, 2, 2, 2, 22, 116, 3, 2, 2,
	2, 24, 121, 3, 2, 2, 2, 26, 141, 3, 2, 2, 2, 28, 153, 3, 2, 2, 2, 30, 189,
	3, 2, 2, 2, 32, 226, 3, 2, 2, 2, 34, 233, 3, 2, 2, 2, 36, 235, 3, 2, 2,
	2, 38, 243, 3, 2, 2, 2, 40, 245, 3, 2, 2, 2, 42, 248, 3, 2, 2, 2, 44, 259,
	3, 2, 2, 2, 46, 48, 5, 4, 3, 2, 47, 46, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2,
	48, 49, 3, 2, 2, 2, 49, 51, 7, 40, 2, 2, 50, 47, 3, 2, 2, 2, 51, 54, 3,
	2, 2, 2, 52, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 3, 3, 2, 2, 2, 54,
	52, 3, 2, 2, 2, 55, 58, 5, 6, 4, 2, 56, 58, 5, 26, 14, 2, 57, 55, 3, 2,
	2, 2, 57, 56, 3, 2, 2, 2, 58, 5, 3, 2, 2, 2, 59, 63, 5, 8, 5, 2, 60, 63,
	5, 10, 6, 2, 61, 63, 5, 12, 7, 2, 62, 59, 3, 2, 2, 2, 62, 60, 3, 2, 2,
	2, 62, 61, 3, 2, 2, 2, 63, 7, 3, 2, 2, 2, 64, 65, 7, 3, 2, 2, 65, 67, 7,
	36, 2, 2, 66, 68, 5, 14, 8, 2, 67, 66, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2,
	68, 86, 3, 2, 2, 2, 69, 70, 7, 4, 2, 2, 70, 72, 7, 36, 2, 2, 71, 73, 5,
	14, 8, 2, 72, 71, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 75, 3, 2, 2, 2, 74,
	76, 5, 22, 12, 2, 75, 74, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 86, 3, 2,
	2, 2, 77, 78, 7, 36, 2, 2, 78, 80, 7, 36, 2, 2, 79, 81, 5, 14, 8, 2, 80,
	79, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 83, 3, 2, 2, 2, 82, 84, 5, 22,
	12, 2, 83, 82, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 86, 3, 2, 2, 2, 85,
	64, 3, 2, 2, 2, 85, 69, 3, 2, 2, 2, 85, 77, 3, 2, 2, 2, 86, 9, 3, 2, 2,
	2, 87, 88, 9, 2, 2, 2, 88, 89, 7, 36, 2, 2, 89, 90, 7, 9, 2, 2, 90, 92,
	7, 36, 2, 2, 91, 93, 5, 14, 8, 2, 92, 91, 3, 2, 2, 2, 92, 93, 3, 2, 2,
	2, 93, 11, 3, 2, 2, 2, 94, 95, 7, 10, 2, 2, 95, 96, 7, 36, 2, 2, 96, 97,
	5, 30, 16, 2, 97, 13, 3, 2, 2, 2, 98, 99, 7, 11, 2, 2, 99, 100, 5, 16,
	9, 2, 100, 101, 7, 12, 2, 2, 101, 15, 3, 2, 2, 2, 102, 107, 5, 18, 10,
	2, 103, 104, 7, 13, 2, 2, 104, 106, 5, 16, 9, 2, 105, 103, 3, 2, 2, 2,
	106, 109, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108,
	17, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 110, 111, 5, 20, 11, 2, 111, 19,
	3, 2, 2, 2, 112, 113, 7, 36, 2, 2, 113, 114, 7, 14, 2, 2, 114, 115, 5,
	30, 16, 2, 115, 21, 3, 2, 2, 2, 116, 117, 5, 24, 13, 2, 117, 23, 3, 2,
	2, 2, 118, 120, 7, 40, 2, 2, 119, 118, 3, 2, 2, 2, 120, 123, 3, 2, 2, 2,
	121, 119, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 124, 3, 2, 2, 2, 123,
	121, 3, 2, 2, 2, 124, 126, 7, 15, 2, 2, 125, 127, 5, 26, 14, 2, 126, 125,
	3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 134, 3, 2, 2, 2, 128, 130, 7, 40,
	2, 2, 129, 131, 5, 26, 14, 2, 130, 129, 3, 2, 2, 2, 130, 131, 3, 2, 2,
	2, 131, 133, 3, 2, 2, 2, 132, 128, 3, 2, 2, 2, 133, 136, 3, 2, 2, 2, 134,
	132, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 137, 3, 2, 2, 2, 136, 134,
	3, 2, 2, 2, 137, 138, 7, 16, 2, 2, 138, 25, 3, 2, 2, 2, 139, 142, 5, 28,
	15, 2, 140, 142, 5, 30, 16, 2, 141, 139, 3, 2, 2, 2, 141, 140, 3, 2, 2,
	2, 142, 27, 3, 2, 2, 2, 143, 144, 7, 36, 2, 2, 144, 145, 7, 14, 2, 2, 145,
	146, 5, 30, 16, 2, 146, 147, 8, 15, 1, 2, 147, 154, 3, 2, 2, 2, 148, 149,
	5, 40, 21, 2, 149, 150, 7, 14, 2, 2, 150, 151, 5, 30, 16, 2, 151, 152,
	8, 15, 1, 2, 152, 154, 3, 2, 2, 2, 153, 143, 3, 2, 2, 2, 153, 148, 3, 2,
	2, 2, 154, 29, 3, 2, 2, 2, 155, 156, 8, 16, 1, 2, 156, 157, 9, 3, 2, 2,
	157, 158, 5, 30, 16, 16, 158, 159, 8, 16, 1, 2, 159, 190, 3, 2, 2, 2, 160,
	161, 7, 32, 2, 2, 161, 162, 7, 11, 2, 2, 162, 163, 5, 30, 16, 2, 163, 164,
	7, 13, 2, 2, 164, 165, 5, 30, 16, 2, 165, 166, 7, 13, 2, 2, 166, 167, 5,
	30, 16, 2, 167, 168, 7, 12, 2, 2, 168, 169, 8, 16, 1, 2, 169, 190, 3, 2,
	2, 2, 170, 171, 5, 32, 17, 2, 171, 172, 8, 16, 1, 2, 172, 190, 3, 2, 2,
	2, 173, 174, 5, 40, 21, 2, 174, 175, 8, 16, 1, 2, 175, 190, 3, 2, 2, 2,
	176, 177, 5, 42, 22, 2, 177, 178, 8, 16, 1, 2, 178, 190, 3, 2, 2, 2, 179,
	180, 5, 44, 23, 2, 180, 181, 8, 16, 1, 2, 181, 190, 3, 2, 2, 2, 182, 183,
	7, 36, 2, 2, 183, 190, 8, 16, 1, 2, 184, 185, 7, 11, 2, 2, 185, 186, 5,
	30, 16, 2, 186, 187, 7, 12, 2, 2, 187, 188, 8, 16, 1, 2, 188, 190, 3, 2,
	2, 2, 189, 155, 3, 2, 2, 2, 189, 160, 3, 2, 2, 2, 189, 170, 3, 2, 2, 2,
	189, 173, 3, 2, 2, 2, 189, 176, 3, 2, 2, 2, 189, 179, 3, 2, 2, 2, 189,
	182, 3, 2, 2, 2, 189, 184, 3, 2, 2, 2, 190, 223, 3, 2, 2, 2, 191, 192,
	12, 15, 2, 2, 192, 193, 9, 4, 2, 2, 193, 194, 5, 30, 16, 16, 194, 195,
	8, 16, 1, 2, 195, 222, 3, 2, 2, 2, 196, 197, 12, 14, 2, 2, 197, 198, 9,
	5, 2, 2, 198, 199, 5, 30, 16, 15, 199, 200, 8, 16, 1, 2, 200, 222, 3, 2,
	2, 2, 201, 202, 12, 13, 2, 2, 202, 203, 9, 6, 2, 2, 203, 204, 5, 30, 16,
	14, 204, 205, 8, 16, 1, 2, 205, 222, 3, 2, 2, 2, 206, 207, 12, 12, 2, 2,
	207, 208, 9, 7, 2, 2, 208, 209, 5, 30, 16, 13, 209, 210, 8, 16, 1, 2, 210,
	222, 3, 2, 2, 2, 211, 212, 12, 11, 2, 2, 212, 213, 7, 30, 2, 2, 213, 214,
	5, 30, 16, 12, 214, 215, 8, 16, 1, 2, 215, 222, 3, 2, 2, 2, 216, 217, 12,
	10, 2, 2, 217, 218, 7, 31, 2, 2, 218, 219, 5, 30, 16, 11, 219, 220, 8,
	16, 1, 2, 220, 222, 3, 2, 2, 2, 221, 191, 3, 2, 2, 2, 221, 196, 3, 2, 2,
	2, 221, 201, 3, 2, 2, 2, 221, 206, 3, 2, 2, 2, 221, 211, 3, 2, 2, 2, 221,
	216, 3, 2, 2, 2, 222, 225, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 223, 224,
	3, 2, 2, 2, 224, 31, 3, 2, 2, 2, 225, 223, 3, 2, 2, 2, 226, 227, 7, 36,
	2, 2, 227, 228, 7, 11, 2, 2, 228, 229, 5, 34, 18, 2, 229, 230, 7, 12, 2,
	2, 230, 33, 3, 2, 2, 2, 231, 234, 5, 36, 19, 2, 232, 234, 5, 16, 9, 2,
	233, 231, 3, 2, 2, 2, 233, 232, 3, 2, 2, 2, 234, 35, 3, 2, 2, 2, 235, 240,
	5, 38, 20, 2, 236, 237, 7, 13, 2, 2, 237, 239, 5, 36, 19, 2, 238, 236,
	3, 2, 2, 2, 239, 242, 3, 2, 2, 2, 240, 238, 3, 2, 2, 2, 240, 241, 3, 2,
	2, 2, 241, 37, 3, 2, 2, 2, 242, 240, 3, 2, 2, 2, 243, 244, 5, 30, 16, 2,
	244, 39, 3, 2, 2, 2, 245, 246, 7, 33, 2, 2, 246, 247, 7, 36, 2, 2, 247,
	41, 3, 2, 2, 2, 248, 249, 7, 34, 2, 2, 249, 250, 7, 36, 2, 2, 250, 43,
	3, 2, 2, 2, 251, 252, 7, 37, 2, 2, 252, 260, 8, 23, 1, 2, 253, 254, 7,
	38, 2, 2, 254, 260, 8, 23, 1, 2, 255, 256, 7, 35, 2, 2, 256, 260, 8, 23,
	1, 2, 257, 258, 7, 39, 2, 2, 258, 260, 8, 23, 1, 2, 259, 251, 3, 2, 2,
	2, 259, 253, 3, 2, 2, 2, 259, 255, 3, 2, 2, 2, 259, 257, 3, 2, 2, 2, 260,
	45, 3, 2, 2, 2, 26, 47, 52, 57, 62, 67, 72, 75, 80, 83, 85, 92, 107, 121,
	126, 130, 134, 141, 153, 189, 221, 223, 233, 240, 259,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'place'", "'trans'", "'arc'", "'iarc'", "'oarc'", "'harc'", "'to'",
	"'reward'", "'('", "')'", "','", "'='", "'{'", "'}'", "'!'", "'+'", "'-'",
	"'*'", "'/'", "'div'", "'mod'", "'<'", "'<='", "'>'", "'>='", "'=='", "'!='",
	"'&&'", "'||'", "'ifelse'", "'#'", "'?'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "LOGICAL",
	"ID", "INT", "FLOAT", "STRING", "NEWLINE", "WS", "LINE_COMMENT", "BLOCK_COMMENT",
}

var ruleNames = []string{
	"prog", "statement", "declaration", "node_declaration", "arc_declaration",
	"reward_declaration", "node_options", "option_list", "option_value", "label_expression",
	"update_block", "simple_block", "simple", "assign_expression", "expression",
	"function_expression", "function_args", "args_list", "args_value", "ntoken_expression",
	"enable_expression", "literal_expression",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type JSPNLParser struct {
	*antlr.BaseParser
}

func NewJSPNLParser(input antlr.TokenStream) *JSPNLParser {
	this := new(JSPNLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "JSPNL.g4"

	return this
}

// JSPNLParser tokens.
const (
	JSPNLParserEOF           = antlr.TokenEOF
	JSPNLParserT__0          = 1
	JSPNLParserT__1          = 2
	JSPNLParserT__2          = 3
	JSPNLParserT__3          = 4
	JSPNLParserT__4          = 5
	JSPNLParserT__5          = 6
	JSPNLParserT__6          = 7
	JSPNLParserT__7          = 8
	JSPNLParserT__8          = 9
	JSPNLParserT__9          = 10
	JSPNLParserT__10         = 11
	JSPNLParserT__11         = 12
	JSPNLParserT__12         = 13
	JSPNLParserT__13         = 14
	JSPNLParserT__14         = 15
	JSPNLParserT__15         = 16
	JSPNLParserT__16         = 17
	JSPNLParserT__17         = 18
	JSPNLParserT__18         = 19
	JSPNLParserT__19         = 20
	JSPNLParserT__20         = 21
	JSPNLParserT__21         = 22
	JSPNLParserT__22         = 23
	JSPNLParserT__23         = 24
	JSPNLParserT__24         = 25
	JSPNLParserT__25         = 26
	JSPNLParserT__26         = 27
	JSPNLParserT__27         = 28
	JSPNLParserT__28         = 29
	JSPNLParserT__29         = 30
	JSPNLParserT__30         = 31
	JSPNLParserT__31         = 32
	JSPNLParserLOGICAL       = 33
	JSPNLParserID            = 34
	JSPNLParserINT           = 35
	JSPNLParserFLOAT         = 36
	JSPNLParserSTRING        = 37
	JSPNLParserNEWLINE       = 38
	JSPNLParserWS            = 39
	JSPNLParserLINE_COMMENT  = 40
	JSPNLParserBLOCK_COMMENT = 41
)

// JSPNLParser rules.
const (
	JSPNLParserRULE_prog                = 0
	JSPNLParserRULE_statement           = 1
	JSPNLParserRULE_declaration         = 2
	JSPNLParserRULE_node_declaration    = 3
	JSPNLParserRULE_arc_declaration     = 4
	JSPNLParserRULE_reward_declaration  = 5
	JSPNLParserRULE_node_options        = 6
	JSPNLParserRULE_option_list         = 7
	JSPNLParserRULE_option_value        = 8
	JSPNLParserRULE_label_expression    = 9
	JSPNLParserRULE_update_block        = 10
	JSPNLParserRULE_simple_block        = 11
	JSPNLParserRULE_simple              = 12
	JSPNLParserRULE_assign_expression   = 13
	JSPNLParserRULE_expression          = 14
	JSPNLParserRULE_function_expression = 15
	JSPNLParserRULE_function_args       = 16
	JSPNLParserRULE_args_list           = 17
	JSPNLParserRULE_args_value          = 18
	JSPNLParserRULE_ntoken_expression   = 19
	JSPNLParserRULE_enable_expression   = 20
	JSPNLParserRULE_literal_expression  = 21
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSPNLParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSPNLParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(JSPNLParserNEWLINE)
}

func (s *ProgContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(JSPNLParserNEWLINE, i)
}

func (s *ProgContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ProgContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *JSPNLParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JSPNLParserRULE_prog)
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
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JSPNLParserT__0)|(1<<JSPNLParserT__1)|(1<<JSPNLParserT__2)|(1<<JSPNLParserT__3)|(1<<JSPNLParserT__4)|(1<<JSPNLParserT__5)|(1<<JSPNLParserT__7)|(1<<JSPNLParserT__8)|(1<<JSPNLParserT__14)|(1<<JSPNLParserT__15)|(1<<JSPNLParserT__16)|(1<<JSPNLParserT__29)|(1<<JSPNLParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(JSPNLParserT__31-32))|(1<<(JSPNLParserLOGICAL-32))|(1<<(JSPNLParserID-32))|(1<<(JSPNLParserINT-32))|(1<<(JSPNLParserFLOAT-32))|(1<<(JSPNLParserSTRING-32))|(1<<(JSPNLParserNEWLINE-32)))) != 0) {
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JSPNLParserT__0)|(1<<JSPNLParserT__1)|(1<<JSPNLParserT__2)|(1<<JSPNLParserT__3)|(1<<JSPNLParserT__4)|(1<<JSPNLParserT__5)|(1<<JSPNLParserT__7)|(1<<JSPNLParserT__8)|(1<<JSPNLParserT__14)|(1<<JSPNLParserT__15)|(1<<JSPNLParserT__16)|(1<<JSPNLParserT__29)|(1<<JSPNLParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(JSPNLParserT__31-32))|(1<<(JSPNLParserLOGICAL-32))|(1<<(JSPNLParserID-32))|(1<<(JSPNLParserINT-32))|(1<<(JSPNLParserFLOAT-32))|(1<<(JSPNLParserSTRING-32)))) != 0) {
			{
				p.SetState(44)
				p.Statement()
			}

		}
		{
			p.SetState(47)
			p.Match(JSPNLParserNEWLINE)
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSPNLParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSPNLParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StatementContext) Simple() ISimpleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *JSPNLParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JSPNLParserRULE_statement)

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

	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.Declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Simple()
		}

	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSPNLParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSPNLParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) Node_declaration() INode_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INode_declarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INode_declarationContext)
}

func (s *DeclarationContext) Arc_declaration() IArc_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArc_declarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArc_declarationContext)
}

func (s *DeclarationContext) Reward_declaration() IReward_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReward_declarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReward_declarationContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *JSPNLParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JSPNLParserRULE_declaration)

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

	switch p.GetTokenStream().LA(1) {
	case JSPNLParserT__0, JSPNLParserT__1, JSPNLParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(57)
			p.Node_declaration()
		}

	case JSPNLParserT__2, JSPNLParserT__3, JSPNLParserT__4, JSPNLParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(58)
			p.Arc_declaration()
		}

	case JSPNLParserT__7:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(59)
			p.Reward_declaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INode_declarationContext is an interface to support dynamic dispatch.
type INode_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNode returns the node token.
	GetNode() antlr.Token

	// GetId returns the id token.
	GetId() antlr.Token

	// SetNode sets the node token.
	SetNode(antlr.Token)

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsNode_declarationContext differentiates from other interfaces.
	IsNode_declarationContext()
}

type Node_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	node   antlr.Token
	id     antlr.Token
}

func NewEmptyNode_declarationContext() *Node_declarationContext {
	var p = new(Node_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSPNLParserRULE_node_declaration
	return p
}

func (*Node_declarationContext) IsNode_declarationContext() {}

func NewNode_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Node_declarationContext {
	var p = new(Node_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSPNLParserRULE_node_declaration

	return p
}

func (s *Node_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Node_declarationContext) GetNode() antlr.Token { return s.node }

func (s *Node_declarationContext) GetId() antlr.Token { return s.id }

func (s *Node_declarationContext) SetNode(v antlr.Token) { s.node = v }

func (s *Node_declarationContext) SetId(v antlr.Token) { s.id = v }

func (s *Node_declarationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(JSPNLParserID)
}

func (s *Node_declarationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(JSPNLParserID, i)
}

func (s *Node_declarationContext) Node_options() INode_optionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INode_optionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INode_optionsContext)
}

func (s *Node_declarationContext) Update_block() IUpdate_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdate_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdate_blockContext)
}

func (s *Node_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Node_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Node_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.EnterNode_declaration(s)
	}
}

func (s *Node_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.ExitNode_declaration(s)
	}
}

func (p *JSPNLParser) Node_declaration() (localctx INode_declarationContext) {
	localctx = NewNode_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JSPNLParserRULE_node_declaration)
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

	p.SetState(83)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JSPNLParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)

			var _m = p.Match(JSPNLParserT__0)

			localctx.(*Node_declarationContext).node = _m
		}
		{
			p.SetState(63)

			var _m = p.Match(JSPNLParserID)

			localctx.(*Node_declarationContext).id = _m
		}
		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JSPNLParserT__8 {
			{
				p.SetState(64)
				p.Node_options()
			}

		}

	case JSPNLParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(67)

			var _m = p.Match(JSPNLParserT__1)

			localctx.(*Node_declarationContext).node = _m
		}
		{
			p.SetState(68)

			var _m = p.Match(JSPNLParserID)

			localctx.(*Node_declarationContext).id = _m
		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JSPNLParserT__8 {
			{
				p.SetState(69)
				p.Node_options()
			}

		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(72)
				p.Update_block()
			}

		}

	case JSPNLParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(75)

			var _m = p.Match(JSPNLParserID)

			localctx.(*Node_declarationContext).node = _m
		}
		{
			p.SetState(76)

			var _m = p.Match(JSPNLParserID)

			localctx.(*Node_declarationContext).id = _m
		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == JSPNLParserT__8 {
			{
				p.SetState(77)
				p.Node_options()
			}

		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(80)
				p.Update_block()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArc_declarationContext is an interface to support dynamic dispatch.
type IArc_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetArctype returns the arctype token.
	GetArctype() antlr.Token

	// GetSrcName returns the srcName token.
	GetSrcName() antlr.Token

	// GetDestName returns the destName token.
	GetDestName() antlr.Token

	// SetArctype sets the arctype token.
	SetArctype(antlr.Token)

	// SetSrcName sets the srcName token.
	SetSrcName(antlr.Token)

	// SetDestName sets the destName token.
	SetDestName(antlr.Token)

	// IsArc_declarationContext differentiates from other interfaces.
	IsArc_declarationContext()
}

type Arc_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	arctype  antlr.Token
	srcName  antlr.Token
	destName antlr.Token
}

func NewEmptyArc_declarationContext() *Arc_declarationContext {
	var p = new(Arc_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSPNLParserRULE_arc_declaration
	return p
}

func (*Arc_declarationContext) IsArc_declarationContext() {}

func NewArc_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Arc_declarationContext {
	var p = new(Arc_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSPNLParserRULE_arc_declaration

	return p
}

func (s *Arc_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Arc_declarationContext) GetArctype() antlr.Token { return s.arctype }

func (s *Arc_declarationContext) GetSrcName() antlr.Token { return s.srcName }

func (s *Arc_declarationContext) GetDestName() antlr.Token { return s.destName }

func (s *Arc_declarationContext) SetArctype(v antlr.Token) { s.arctype = v }

func (s *Arc_declarationContext) SetSrcName(v antlr.Token) { s.srcName = v }

func (s *Arc_declarationContext) SetDestName(v antlr.Token) { s.destName = v }

func (s *Arc_declarationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(JSPNLParserID)
}

func (s *Arc_declarationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(JSPNLParserID, i)
}

func (s *Arc_declarationContext) Node_options() INode_optionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INode_optionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INode_optionsContext)
}

func (s *Arc_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Arc_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Arc_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.EnterArc_declaration(s)
	}
}

func (s *Arc_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.ExitArc_declaration(s)
	}
}

func (p *JSPNLParser) Arc_declaration() (localctx IArc_declarationContext) {
	localctx = NewArc_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JSPNLParserRULE_arc_declaration)
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
		p.SetState(85)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Arc_declarationContext).arctype = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JSPNLParserT__2)|(1<<JSPNLParserT__3)|(1<<JSPNLParserT__4)|(1<<JSPNLParserT__5))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Arc_declarationContext).arctype = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(86)

		var _m = p.Match(JSPNLParserID)

		localctx.(*Arc_declarationContext).srcName = _m
	}
	{
		p.SetState(87)
		p.Match(JSPNLParserT__6)
	}
	{
		p.SetState(88)

		var _m = p.Match(JSPNLParserID)

		localctx.(*Arc_declarationContext).destName = _m
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == JSPNLParserT__8 {
		{
			p.SetState(89)
			p.Node_options()
		}

	}

	return localctx
}

// IReward_declarationContext is an interface to support dynamic dispatch.
type IReward_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsReward_declarationContext differentiates from other interfaces.
	IsReward_declarationContext()
}

type Reward_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyReward_declarationContext() *Reward_declarationContext {
	var p = new(Reward_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSPNLParserRULE_reward_declaration
	return p
}

func (*Reward_declarationContext) IsReward_declarationContext() {}

func NewReward_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Reward_declarationContext {
	var p = new(Reward_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSPNLParserRULE_reward_declaration

	return p
}

func (s *Reward_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Reward_declarationContext) GetId() antlr.Token { return s.id }

func (s *Reward_declarationContext) SetId(v antlr.Token) { s.id = v }

func (s *Reward_declarationContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Reward_declarationContext) ID() antlr.TerminalNode {
	return s.GetToken(JSPNLParserID, 0)
}

func (s *Reward_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Reward_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Reward_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.EnterReward_declaration(s)
	}
}

func (s *Reward_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.ExitReward_declaration(s)
	}
}

func (p *JSPNLParser) Reward_declaration() (localctx IReward_declarationContext) {
	localctx = NewReward_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JSPNLParserRULE_reward_declaration)

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
		p.Match(JSPNLParserT__7)
	}
	{
		p.SetState(93)

		var _m = p.Match(JSPNLParserID)

		localctx.(*Reward_declarationContext).id = _m
	}
	{
		p.SetState(94)
		p.expression(0)
	}

	return localctx
}

// INode_optionsContext is an interface to support dynamic dispatch.
type INode_optionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNode_optionsContext differentiates from other interfaces.
	IsNode_optionsContext()
}

type Node_optionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNode_optionsContext() *Node_optionsContext {
	var p = new(Node_optionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSPNLParserRULE_node_options
	return p
}

func (*Node_optionsContext) IsNode_optionsContext() {}

func NewNode_optionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Node_optionsContext {
	var p = new(Node_optionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSPNLParserRULE_node_options

	return p
}

func (s *Node_optionsContext) GetParser() antlr.Parser { return s.parser }

func (s *Node_optionsContext) Option_list() IOption_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOption_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOption_listContext)
}

func (s *Node_optionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Node_optionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Node_optionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.EnterNode_options(s)
	}
}

func (s *Node_optionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.ExitNode_options(s)
	}
}

func (p *JSPNLParser) Node_options() (localctx INode_optionsContext) {
	localctx = NewNode_optionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JSPNLParserRULE_node_options)

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
		p.SetState(96)
		p.Match(JSPNLParserT__8)
	}
	{
		p.SetState(97)
		p.Option_list()
	}
	{
		p.SetState(98)
		p.Match(JSPNLParserT__9)
	}

	return localctx
}

// IOption_listContext is an interface to support dynamic dispatch.
type IOption_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOption_listContext differentiates from other interfaces.
	IsOption_listContext()
}

type Option_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOption_listContext() *Option_listContext {
	var p = new(Option_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSPNLParserRULE_option_list
	return p
}

func (*Option_listContext) IsOption_listContext() {}

func NewOption_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Option_listContext {
	var p = new(Option_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSPNLParserRULE_option_list

	return p
}

func (s *Option_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Option_listContext) Option_value() IOption_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOption_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOption_valueContext)
}

func (s *Option_listContext) AllOption_list() []IOption_listContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOption_listContext)(nil)).Elem())
	var tst = make([]IOption_listContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOption_listContext)
		}
	}

	return tst
}

func (s *Option_listContext) Option_list(i int) IOption_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOption_listContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOption_listContext)
}

func (s *Option_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Option_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Option_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.EnterOption_list(s)
	}
}

func (s *Option_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.ExitOption_list(s)
	}
}

func (p *JSPNLParser) Option_list() (localctx IOption_listContext) {
	localctx = NewOption_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JSPNLParserRULE_option_list)

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
		p.SetState(100)
		p.Option_value()
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(101)
				p.Match(JSPNLParserT__10)
			}
			{
				p.SetState(102)
				p.Option_list()
			}

		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IOption_valueContext is an interface to support dynamic dispatch.
type IOption_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOption_valueContext differentiates from other interfaces.
	IsOption_valueContext()
}

type Option_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOption_valueContext() *Option_valueContext {
	var p = new(Option_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSPNLParserRULE_option_value
	return p
}

func (*Option_valueContext) IsOption_valueContext() {}

func NewOption_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Option_valueContext {
	var p = new(Option_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSPNLParserRULE_option_value

	return p
}

func (s *Option_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Option_valueContext) Label_expression() ILabel_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabel_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabel_expressionContext)
}

func (s *Option_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Option_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Option_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.EnterOption_value(s)
	}
}

func (s *Option_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.ExitOption_value(s)
	}
}

func (p *JSPNLParser) Option_value() (localctx IOption_valueContext) {
	localctx = NewOption_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JSPNLParserRULE_option_value)

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
		p.SetState(108)
		p.Label_expression()
	}

	return localctx
}

// ILabel_expressionContext is an interface to support dynamic dispatch.
type ILabel_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsLabel_expressionContext differentiates from other interfaces.
	IsLabel_expressionContext()
}

type Label_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyLabel_expressionContext() *Label_expressionContext {
	var p = new(Label_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSPNLParserRULE_label_expression
	return p
}

func (*Label_expressionContext) IsLabel_expressionContext() {}

func NewLabel_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Label_expressionContext {
	var p = new(Label_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSPNLParserRULE_label_expression

	return p
}

func (s *Label_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Label_expressionContext) GetId() antlr.Token { return s.id }

func (s *Label_expressionContext) SetId(v antlr.Token) { s.id = v }

func (s *Label_expressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Label_expressionContext) ID() antlr.TerminalNode {
	return s.GetToken(JSPNLParserID, 0)
}

func (s *Label_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Label_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Label_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.EnterLabel_expression(s)
	}
}

func (s *Label_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.ExitLabel_expression(s)
	}
}

func (p *JSPNLParser) Label_expression() (localctx ILabel_expressionContext) {
	localctx = NewLabel_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JSPNLParserRULE_label_expression)

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
		p.SetState(110)

		var _m = p.Match(JSPNLParserID)

		localctx.(*Label_expressionContext).id = _m
	}
	{
		p.SetState(111)
		p.Match(JSPNLParserT__11)
	}
	{
		p.SetState(112)
		p.expression(0)
	}

	return localctx
}

// IUpdate_blockContext is an interface to support dynamic dispatch.
type IUpdate_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdate_blockContext differentiates from other interfaces.
	IsUpdate_blockContext()
}

type Update_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdate_blockContext() *Update_blockContext {
	var p = new(Update_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSPNLParserRULE_update_block
	return p
}

func (*Update_blockContext) IsUpdate_blockContext() {}

func NewUpdate_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Update_blockContext {
	var p = new(Update_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSPNLParserRULE_update_block

	return p
}

func (s *Update_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Update_blockContext) Simple_block() ISimple_blockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_blockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_blockContext)
}

func (s *Update_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Update_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.EnterUpdate_block(s)
	}
}

func (s *Update_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.ExitUpdate_block(s)
	}
}

func (p *JSPNLParser) Update_block() (localctx IUpdate_blockContext) {
	localctx = NewUpdate_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, JSPNLParserRULE_update_block)

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
		p.SetState(114)
		p.Simple_block()
	}

	return localctx
}

// ISimple_blockContext is an interface to support dynamic dispatch.
type ISimple_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimple_blockContext differentiates from other interfaces.
	IsSimple_blockContext()
}

type Simple_blockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_blockContext() *Simple_blockContext {
	var p = new(Simple_blockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSPNLParserRULE_simple_block
	return p
}

func (*Simple_blockContext) IsSimple_blockContext() {}

func NewSimple_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_blockContext {
	var p = new(Simple_blockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSPNLParserRULE_simple_block

	return p
}

func (s *Simple_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_blockContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(JSPNLParserNEWLINE)
}

func (s *Simple_blockContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(JSPNLParserNEWLINE, i)
}

func (s *Simple_blockContext) AllSimple() []ISimpleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISimpleContext)(nil)).Elem())
	var tst = make([]ISimpleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISimpleContext)
		}
	}

	return tst
}

func (s *Simple_blockContext) Simple(i int) ISimpleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISimpleContext)
}

func (s *Simple_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.EnterSimple_block(s)
	}
}

func (s *Simple_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.ExitSimple_block(s)
	}
}

func (p *JSPNLParser) Simple_block() (localctx ISimple_blockContext) {
	localctx = NewSimple_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, JSPNLParserRULE_simple_block)
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
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JSPNLParserNEWLINE {
		{
			p.SetState(116)
			p.Match(JSPNLParserNEWLINE)
		}

		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(122)
		p.Match(JSPNLParserT__12)
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-9)&-(0x1f+1)) == 0 && ((1<<uint((_la-9)))&((1<<(JSPNLParserT__8-9))|(1<<(JSPNLParserT__14-9))|(1<<(JSPNLParserT__15-9))|(1<<(JSPNLParserT__16-9))|(1<<(JSPNLParserT__29-9))|(1<<(JSPNLParserT__30-9))|(1<<(JSPNLParserT__31-9))|(1<<(JSPNLParserLOGICAL-9))|(1<<(JSPNLParserID-9))|(1<<(JSPNLParserINT-9))|(1<<(JSPNLParserFLOAT-9))|(1<<(JSPNLParserSTRING-9)))) != 0 {
		{
			p.SetState(123)
			p.Simple()
		}

	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == JSPNLParserNEWLINE {
		{
			p.SetState(126)
			p.Match(JSPNLParserNEWLINE)
		}
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-9)&-(0x1f+1)) == 0 && ((1<<uint((_la-9)))&((1<<(JSPNLParserT__8-9))|(1<<(JSPNLParserT__14-9))|(1<<(JSPNLParserT__15-9))|(1<<(JSPNLParserT__16-9))|(1<<(JSPNLParserT__29-9))|(1<<(JSPNLParserT__30-9))|(1<<(JSPNLParserT__31-9))|(1<<(JSPNLParserLOGICAL-9))|(1<<(JSPNLParserID-9))|(1<<(JSPNLParserINT-9))|(1<<(JSPNLParserFLOAT-9))|(1<<(JSPNLParserSTRING-9)))) != 0 {
			{
				p.SetState(127)
				p.Simple()
			}

		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(135)
		p.Match(JSPNLParserT__13)
	}

	return localctx
}

// ISimpleContext is an interface to support dynamic dispatch.
type ISimpleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimpleContext differentiates from other interfaces.
	IsSimpleContext()
}

type SimpleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleContext() *SimpleContext {
	var p = new(SimpleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSPNLParserRULE_simple
	return p
}

func (*SimpleContext) IsSimpleContext() {}

func NewSimpleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleContext {
	var p = new(SimpleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSPNLParserRULE_simple

	return p
}

func (s *SimpleContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleContext) Assign_expression() IAssign_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssign_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssign_expressionContext)
}

func (s *SimpleContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.EnterSimple(s)
	}
}

func (s *SimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.ExitSimple(s)
	}
}

func (p *JSPNLParser) Simple() (localctx ISimpleContext) {
	localctx = NewSimpleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, JSPNLParserRULE_simple)

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

	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(137)
			p.Assign_expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(138)
			p.expression(0)
		}

	}

	return localctx
}

// IAssign_expressionContext is an interface to support dynamic dispatch.
type IAssign_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetExprtype returns the exprtype attribute.
	GetExprtype() int

	// SetExprtype sets the exprtype attribute.
	SetExprtype(int)

	// IsAssign_expressionContext differentiates from other interfaces.
	IsAssign_expressionContext()
}

type Assign_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	exprtype int
	id       antlr.Token
}

func NewEmptyAssign_expressionContext() *Assign_expressionContext {
	var p = new(Assign_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSPNLParserRULE_assign_expression
	return p
}

func (*Assign_expressionContext) IsAssign_expressionContext() {}

func NewAssign_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_expressionContext {
	var p = new(Assign_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSPNLParserRULE_assign_expression

	return p
}

func (s *Assign_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_expressionContext) GetId() antlr.Token { return s.id }

func (s *Assign_expressionContext) SetId(v antlr.Token) { s.id = v }

func (s *Assign_expressionContext) GetExprtype() int { return s.exprtype }

func (s *Assign_expressionContext) SetExprtype(v int) { s.exprtype = v }

func (s *Assign_expressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Assign_expressionContext) ID() antlr.TerminalNode {
	return s.GetToken(JSPNLParserID, 0)
}

func (s *Assign_expressionContext) Ntoken_expression() INtoken_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INtoken_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INtoken_expressionContext)
}

func (s *Assign_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assign_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.EnterAssign_expression(s)
	}
}

func (s *Assign_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.ExitAssign_expression(s)
	}
}

func (p *JSPNLParser) Assign_expression() (localctx IAssign_expressionContext) {
	localctx = NewAssign_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, JSPNLParserRULE_assign_expression)

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

	p.SetState(151)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JSPNLParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)

			var _m = p.Match(JSPNLParserID)

			localctx.(*Assign_expressionContext).id = _m
		}
		{
			p.SetState(142)
			p.Match(JSPNLParserT__11)
		}
		{
			p.SetState(143)
			p.expression(0)
		}
		localctx.(*Assign_expressionContext).SetExprtype(1)

	case JSPNLParserT__30:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)
			p.Ntoken_expression()
		}
		{
			p.SetState(147)
			p.Match(JSPNLParserT__11)
		}
		{
			p.SetState(148)
			p.expression(0)
		}
		localctx.(*Assign_expressionContext).SetExprtype(2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// GetId returns the id token.
	GetId() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetNodetype returns the nodetype attribute.
	GetNodetype() int

	// SetNodetype sets the nodetype attribute.
	SetNodetype(int)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	nodetype int
	op       antlr.Token
	id       antlr.Token
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSPNLParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSPNLParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetOp() antlr.Token { return s.op }

func (s *ExpressionContext) GetId() antlr.Token { return s.id }

func (s *ExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpressionContext) SetId(v antlr.Token) { s.id = v }

func (s *ExpressionContext) GetNodetype() int { return s.nodetype }

func (s *ExpressionContext) SetNodetype(v int) { s.nodetype = v }

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) Function_expression() IFunction_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_expressionContext)
}

func (s *ExpressionContext) Ntoken_expression() INtoken_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INtoken_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INtoken_expressionContext)
}

func (s *ExpressionContext) Enable_expression() IEnable_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnable_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnable_expressionContext)
}

func (s *ExpressionContext) Literal_expression() ILiteral_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteral_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteral_expressionContext)
}

func (s *ExpressionContext) ID() antlr.TerminalNode {
	return s.GetToken(JSPNLParserID, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *JSPNLParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *JSPNLParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 28
	p.EnterRecursionRule(localctx, 28, JSPNLParserRULE_expression, _p)
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
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(154)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ExpressionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JSPNLParserT__14)|(1<<JSPNLParserT__15)|(1<<JSPNLParserT__16))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ExpressionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(155)
			p.expression(14)
		}
		localctx.(*ExpressionContext).SetNodetype(1)

	case 2:
		{
			p.SetState(158)

			var _m = p.Match(JSPNLParserT__29)

			localctx.(*ExpressionContext).op = _m
		}
		{
			p.SetState(159)
			p.Match(JSPNLParserT__8)
		}
		{
			p.SetState(160)
			p.expression(0)
		}
		{
			p.SetState(161)
			p.Match(JSPNLParserT__10)
		}
		{
			p.SetState(162)
			p.expression(0)
		}
		{
			p.SetState(163)
			p.Match(JSPNLParserT__10)
		}
		{
			p.SetState(164)
			p.expression(0)
		}
		{
			p.SetState(165)
			p.Match(JSPNLParserT__9)
		}
		localctx.(*ExpressionContext).SetNodetype(8)

	case 3:
		{
			p.SetState(168)
			p.Function_expression()
		}
		localctx.(*ExpressionContext).SetNodetype(9)

	case 4:
		{
			p.SetState(171)
			p.Ntoken_expression()
		}
		localctx.(*ExpressionContext).SetNodetype(10)

	case 5:
		{
			p.SetState(174)
			p.Enable_expression()
		}
		localctx.(*ExpressionContext).SetNodetype(14)

	case 6:
		{
			p.SetState(177)
			p.Literal_expression()
		}
		localctx.(*ExpressionContext).SetNodetype(11)

	case 7:
		{
			p.SetState(180)

			var _m = p.Match(JSPNLParserID)

			localctx.(*ExpressionContext).id = _m
		}
		localctx.(*ExpressionContext).SetNodetype(12)

	case 8:
		{
			p.SetState(182)
			p.Match(JSPNLParserT__8)
		}
		{
			p.SetState(183)
			p.expression(0)
		}
		{
			p.SetState(184)
			p.Match(JSPNLParserT__9)
		}
		localctx.(*ExpressionContext).SetNodetype(13)

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(219)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JSPNLParserRULE_expression)
				p.SetState(189)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(190)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JSPNLParserT__17)|(1<<JSPNLParserT__18)|(1<<JSPNLParserT__19)|(1<<JSPNLParserT__20))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(191)
					p.expression(14)
				}
				localctx.(*ExpressionContext).SetNodetype(2)

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JSPNLParserRULE_expression)
				p.SetState(194)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(195)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == JSPNLParserT__15 || _la == JSPNLParserT__16) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(196)
					p.expression(13)
				}
				localctx.(*ExpressionContext).SetNodetype(3)

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JSPNLParserRULE_expression)
				p.SetState(199)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(200)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<JSPNLParserT__21)|(1<<JSPNLParserT__22)|(1<<JSPNLParserT__23)|(1<<JSPNLParserT__24))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(201)
					p.expression(12)
				}
				localctx.(*ExpressionContext).SetNodetype(4)

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JSPNLParserRULE_expression)
				p.SetState(204)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(205)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == JSPNLParserT__25 || _la == JSPNLParserT__26) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(206)
					p.expression(11)
				}
				localctx.(*ExpressionContext).SetNodetype(5)

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JSPNLParserRULE_expression)
				p.SetState(209)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(210)

					var _m = p.Match(JSPNLParserT__27)

					localctx.(*ExpressionContext).op = _m
				}
				{
					p.SetState(211)
					p.expression(10)
				}
				localctx.(*ExpressionContext).SetNodetype(6)

			case 6:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, JSPNLParserRULE_expression)
				p.SetState(214)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(215)

					var _m = p.Match(JSPNLParserT__28)

					localctx.(*ExpressionContext).op = _m
				}
				{
					p.SetState(216)
					p.expression(9)
				}
				localctx.(*ExpressionContext).SetNodetype(7)

			}

		}
		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

// IFunction_expressionContext is an interface to support dynamic dispatch.
type IFunction_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsFunction_expressionContext differentiates from other interfaces.
	IsFunction_expressionContext()
}

type Function_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyFunction_expressionContext() *Function_expressionContext {
	var p = new(Function_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSPNLParserRULE_function_expression
	return p
}

func (*Function_expressionContext) IsFunction_expressionContext() {}

func NewFunction_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_expressionContext {
	var p = new(Function_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSPNLParserRULE_function_expression

	return p
}

func (s *Function_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_expressionContext) GetId() antlr.Token { return s.id }

func (s *Function_expressionContext) SetId(v antlr.Token) { s.id = v }

func (s *Function_expressionContext) Function_args() IFunction_argsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_argsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_argsContext)
}

func (s *Function_expressionContext) ID() antlr.TerminalNode {
	return s.GetToken(JSPNLParserID, 0)
}

func (s *Function_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.EnterFunction_expression(s)
	}
}

func (s *Function_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.ExitFunction_expression(s)
	}
}

func (p *JSPNLParser) Function_expression() (localctx IFunction_expressionContext) {
	localctx = NewFunction_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, JSPNLParserRULE_function_expression)

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
		p.SetState(224)

		var _m = p.Match(JSPNLParserID)

		localctx.(*Function_expressionContext).id = _m
	}
	{
		p.SetState(225)
		p.Match(JSPNLParserT__8)
	}
	{
		p.SetState(226)
		p.Function_args()
	}
	{
		p.SetState(227)
		p.Match(JSPNLParserT__9)
	}

	return localctx
}

// IFunction_argsContext is an interface to support dynamic dispatch.
type IFunction_argsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_argsContext differentiates from other interfaces.
	IsFunction_argsContext()
}

type Function_argsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_argsContext() *Function_argsContext {
	var p = new(Function_argsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSPNLParserRULE_function_args
	return p
}

func (*Function_argsContext) IsFunction_argsContext() {}

func NewFunction_argsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_argsContext {
	var p = new(Function_argsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSPNLParserRULE_function_args

	return p
}

func (s *Function_argsContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_argsContext) Args_list() IArgs_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgs_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgs_listContext)
}

func (s *Function_argsContext) Option_list() IOption_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOption_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOption_listContext)
}

func (s *Function_argsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_argsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_argsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.EnterFunction_args(s)
	}
}

func (s *Function_argsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.ExitFunction_args(s)
	}
}

func (p *JSPNLParser) Function_args() (localctx IFunction_argsContext) {
	localctx = NewFunction_argsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, JSPNLParserRULE_function_args)

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

	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(229)
			p.Args_list()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(230)
			p.Option_list()
		}

	}

	return localctx
}

// IArgs_listContext is an interface to support dynamic dispatch.
type IArgs_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgs_listContext differentiates from other interfaces.
	IsArgs_listContext()
}

type Args_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgs_listContext() *Args_listContext {
	var p = new(Args_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSPNLParserRULE_args_list
	return p
}

func (*Args_listContext) IsArgs_listContext() {}

func NewArgs_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Args_listContext {
	var p = new(Args_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSPNLParserRULE_args_list

	return p
}

func (s *Args_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Args_listContext) Args_value() IArgs_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgs_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgs_valueContext)
}

func (s *Args_listContext) AllArgs_list() []IArgs_listContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgs_listContext)(nil)).Elem())
	var tst = make([]IArgs_listContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgs_listContext)
		}
	}

	return tst
}

func (s *Args_listContext) Args_list(i int) IArgs_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgs_listContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgs_listContext)
}

func (s *Args_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Args_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Args_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.EnterArgs_list(s)
	}
}

func (s *Args_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.ExitArgs_list(s)
	}
}

func (p *JSPNLParser) Args_list() (localctx IArgs_listContext) {
	localctx = NewArgs_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, JSPNLParserRULE_args_list)

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
		p.SetState(233)
		p.Args_value()
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(234)
				p.Match(JSPNLParserT__10)
			}
			{
				p.SetState(235)
				p.Args_list()
			}

		}
		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}

	return localctx
}

// IArgs_valueContext is an interface to support dynamic dispatch.
type IArgs_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgs_valueContext differentiates from other interfaces.
	IsArgs_valueContext()
}

type Args_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgs_valueContext() *Args_valueContext {
	var p = new(Args_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSPNLParserRULE_args_value
	return p
}

func (*Args_valueContext) IsArgs_valueContext() {}

func NewArgs_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Args_valueContext {
	var p = new(Args_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSPNLParserRULE_args_value

	return p
}

func (s *Args_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Args_valueContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Args_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Args_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Args_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.EnterArgs_value(s)
	}
}

func (s *Args_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.ExitArgs_value(s)
	}
}

func (p *JSPNLParser) Args_value() (localctx IArgs_valueContext) {
	localctx = NewArgs_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, JSPNLParserRULE_args_value)

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
		p.SetState(241)
		p.expression(0)
	}

	return localctx
}

// INtoken_expressionContext is an interface to support dynamic dispatch.
type INtoken_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsNtoken_expressionContext differentiates from other interfaces.
	IsNtoken_expressionContext()
}

type Ntoken_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyNtoken_expressionContext() *Ntoken_expressionContext {
	var p = new(Ntoken_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSPNLParserRULE_ntoken_expression
	return p
}

func (*Ntoken_expressionContext) IsNtoken_expressionContext() {}

func NewNtoken_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ntoken_expressionContext {
	var p = new(Ntoken_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSPNLParserRULE_ntoken_expression

	return p
}

func (s *Ntoken_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Ntoken_expressionContext) GetId() antlr.Token { return s.id }

func (s *Ntoken_expressionContext) SetId(v antlr.Token) { s.id = v }

func (s *Ntoken_expressionContext) ID() antlr.TerminalNode {
	return s.GetToken(JSPNLParserID, 0)
}

func (s *Ntoken_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ntoken_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ntoken_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.EnterNtoken_expression(s)
	}
}

func (s *Ntoken_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.ExitNtoken_expression(s)
	}
}

func (p *JSPNLParser) Ntoken_expression() (localctx INtoken_expressionContext) {
	localctx = NewNtoken_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, JSPNLParserRULE_ntoken_expression)

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
		p.SetState(243)
		p.Match(JSPNLParserT__30)
	}
	{
		p.SetState(244)

		var _m = p.Match(JSPNLParserID)

		localctx.(*Ntoken_expressionContext).id = _m
	}

	return localctx
}

// IEnable_expressionContext is an interface to support dynamic dispatch.
type IEnable_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsEnable_expressionContext differentiates from other interfaces.
	IsEnable_expressionContext()
}

type Enable_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyEnable_expressionContext() *Enable_expressionContext {
	var p = new(Enable_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSPNLParserRULE_enable_expression
	return p
}

func (*Enable_expressionContext) IsEnable_expressionContext() {}

func NewEnable_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Enable_expressionContext {
	var p = new(Enable_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSPNLParserRULE_enable_expression

	return p
}

func (s *Enable_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Enable_expressionContext) GetId() antlr.Token { return s.id }

func (s *Enable_expressionContext) SetId(v antlr.Token) { s.id = v }

func (s *Enable_expressionContext) ID() antlr.TerminalNode {
	return s.GetToken(JSPNLParserID, 0)
}

func (s *Enable_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Enable_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Enable_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.EnterEnable_expression(s)
	}
}

func (s *Enable_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.ExitEnable_expression(s)
	}
}

func (p *JSPNLParser) Enable_expression() (localctx IEnable_expressionContext) {
	localctx = NewEnable_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, JSPNLParserRULE_enable_expression)

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
		p.SetState(246)
		p.Match(JSPNLParserT__31)
	}
	{
		p.SetState(247)

		var _m = p.Match(JSPNLParserID)

		localctx.(*Enable_expressionContext).id = _m
	}

	return localctx
}

// ILiteral_expressionContext is an interface to support dynamic dispatch.
type ILiteral_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVal returns the val token.
	GetVal() antlr.Token

	// SetVal sets the val token.
	SetVal(antlr.Token)

	// GetLittype returns the littype attribute.
	GetLittype() int

	// SetLittype sets the littype attribute.
	SetLittype(int)

	// IsLiteral_expressionContext differentiates from other interfaces.
	IsLiteral_expressionContext()
}

type Literal_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	littype int
	val     antlr.Token
}

func NewEmptyLiteral_expressionContext() *Literal_expressionContext {
	var p = new(Literal_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSPNLParserRULE_literal_expression
	return p
}

func (*Literal_expressionContext) IsLiteral_expressionContext() {}

func NewLiteral_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Literal_expressionContext {
	var p = new(Literal_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSPNLParserRULE_literal_expression

	return p
}

func (s *Literal_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Literal_expressionContext) GetVal() antlr.Token { return s.val }

func (s *Literal_expressionContext) SetVal(v antlr.Token) { s.val = v }

func (s *Literal_expressionContext) GetLittype() int { return s.littype }

func (s *Literal_expressionContext) SetLittype(v int) { s.littype = v }

func (s *Literal_expressionContext) INT() antlr.TerminalNode {
	return s.GetToken(JSPNLParserINT, 0)
}

func (s *Literal_expressionContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(JSPNLParserFLOAT, 0)
}

func (s *Literal_expressionContext) LOGICAL() antlr.TerminalNode {
	return s.GetToken(JSPNLParserLOGICAL, 0)
}

func (s *Literal_expressionContext) STRING() antlr.TerminalNode {
	return s.GetToken(JSPNLParserSTRING, 0)
}

func (s *Literal_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Literal_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.EnterLiteral_expression(s)
	}
}

func (s *Literal_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSPNLListener); ok {
		listenerT.ExitLiteral_expression(s)
	}
}

func (p *JSPNLParser) Literal_expression() (localctx ILiteral_expressionContext) {
	localctx = NewLiteral_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, JSPNLParserRULE_literal_expression)

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

	p.SetState(257)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JSPNLParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(249)

			var _m = p.Match(JSPNLParserINT)

			localctx.(*Literal_expressionContext).val = _m
		}
		localctx.(*Literal_expressionContext).SetLittype(1)

	case JSPNLParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(251)

			var _m = p.Match(JSPNLParserFLOAT)

			localctx.(*Literal_expressionContext).val = _m
		}
		localctx.(*Literal_expressionContext).SetLittype(2)

	case JSPNLParserLOGICAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(253)

			var _m = p.Match(JSPNLParserLOGICAL)

			localctx.(*Literal_expressionContext).val = _m
		}
		localctx.(*Literal_expressionContext).SetLittype(3)

	case JSPNLParserSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(255)

			var _m = p.Match(JSPNLParserSTRING)

			localctx.(*Literal_expressionContext).val = _m
		}
		localctx.(*Literal_expressionContext).SetLittype(4)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *JSPNLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 14:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *JSPNLParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
