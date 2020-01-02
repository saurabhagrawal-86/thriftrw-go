// Code generated by goyacc -l thrift.y. DO NOT EDIT.

// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
package internal

import (
	__yyfmt__ "fmt"

	"go.uber.org/thriftrw/ast"
)

type yySymType struct {
	yys int
	// Used to record line numbers when the line number at the start point is
	// required.
	line int

	docstring string

	// Holds the final AST for the file.
	prog *ast.Program

	// Other intermediate variables:

	bul bool
	str string
	i64 int64
	dub float64

	fieldType     ast.Type
	structType    ast.StructureType
	baseTypeID    ast.BaseTypeID
	fieldRequired ast.Requiredness

	field  *ast.Field
	fields []*ast.Field

	header  ast.Header
	headers []ast.Header

	function  *ast.Function
	functions []*ast.Function

	enumItem  *ast.EnumItem
	enumItems []*ast.EnumItem

	definition  ast.Definition
	definitions []ast.Definition

	typeAnnotations []*ast.Annotation

	constantValue    ast.ConstantValue
	constantValues   []ast.ConstantValue
	constantMapItems []ast.ConstantMapItem
}

const IDENTIFIER = 57346
const LITERAL = 57347
const INTCONSTANT = 57348
const DUBCONSTANT = 57349
const NAMESPACE = 57350
const INCLUDE = 57351
const VOID = 57352
const BOOL = 57353
const BYTE = 57354
const I8 = 57355
const I16 = 57356
const I32 = 57357
const I64 = 57358
const DOUBLE = 57359
const STRING = 57360
const BINARY = 57361
const MAP = 57362
const LIST = 57363
const SET = 57364
const ONEWAY = 57365
const TYPEDEF = 57366
const STRUCT = 57367
const UNION = 57368
const EXCEPTION = 57369
const EXTENDS = 57370
const THROWS = 57371
const SERVICE = 57372
const ENUM = 57373
const CONST = 57374
const REQUIRED = 57375
const OPTIONAL = 57376
const TRUE = 57377
const FALSE = 57378

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENTIFIER",
	"LITERAL",
	"INTCONSTANT",
	"DUBCONSTANT",
	"NAMESPACE",
	"INCLUDE",
	"VOID",
	"BOOL",
	"BYTE",
	"I8",
	"I16",
	"I32",
	"I64",
	"DOUBLE",
	"STRING",
	"BINARY",
	"MAP",
	"LIST",
	"SET",
	"ONEWAY",
	"TYPEDEF",
	"STRUCT",
	"UNION",
	"EXCEPTION",
	"EXTENDS",
	"THROWS",
	"SERVICE",
	"ENUM",
	"CONST",
	"REQUIRED",
	"OPTIONAL",
	"TRUE",
	"FALSE",
	"'*'",
	"'='",
	"'{'",
	"'}'",
	"':'",
	"'('",
	"')'",
	"'<'",
	"','",
	"'>'",
	"'['",
	"']'",
	"';'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	8, 70,
	9, 70,
	-2, 8,
	-1, 3,
	1, 1,
	-2, 70,
}

const yyPrivate = 57344

const yyLast = 187

var yyAct = [...]int{

	30, 10, 66, 5, 7, 63, 64, 56, 122, 85,
	29, 71, 67, 68, 11, 11, 87, 13, 12, 12,
	124, 94, 93, 92, 60, 59, 58, 159, 151, 150,
	31, 126, 90, 57, 157, 57, 57, 144, 140, 127,
	129, 69, 70, 120, 83, 71, 67, 68, 80, 77,
	54, 89, 105, 52, 118, 65, 72, 51, 61, 88,
	17, 53, 55, 79, 82, 136, 137, 154, 104, 74,
	75, 76, 115, 134, 91, 69, 70, 9, 8, 113,
	96, 15, 14, 99, 132, 95, 102, 26, 98, 97,
	146, 101, 100, 16, 138, 112, 108, 86, 50, 35,
	34, 110, 111, 109, 33, 32, 28, 72, 121, 119,
	123, 27, 117, 153, 116, 128, 114, 103, 73, 107,
	125, 130, 72, 106, 131, 3, 6, 62, 78, 84,
	2, 4, 133, 81, 141, 21, 135, 139, 36, 1,
	0, 72, 142, 145, 0, 0, 143, 148, 82, 0,
	147, 72, 0, 152, 149, 0, 0, 0, 0, 82,
	155, 156, 0, 158, 19, 23, 24, 25, 40, 0,
	22, 20, 18, 0, 0, 41, 42, 43, 44, 45,
	46, 47, 48, 49, 37, 38, 39,
}
var yyPact = [...]int{

	-1000, -1000, -1000, -1000, -1000, 69, -31, -1000, 77, 56,
	-1000, -1000, -1000, 140, -1000, 82, 107, 102, -1000, -1000,
	101, 100, 96, -1000, -1000, -1000, -1000, -1000, -1000, 95,
	164, 94, 18, 14, 22, 24, -6, -18, -19, -20,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-6, -1000, -1000, -1000, -1000, 40, -1000, -1000, -1000, -1000,
	-1000, -1000, 9, 8, 4, 93, -1000, -1000, -1000, -1000,
	-1000, -1000, 12, -11, -22, -24, -25, -6, -31, -1000,
	-6, -31, -1000, -6, -31, 45, 13, -1000, -1000, -1000,
	-1000, 92, -1000, -6, -6, -1000, -1000, 91, -1000, -1000,
	73, -1000, -1000, 62, -1000, -1000, 6, 3, -30, -26,
	-1000, -1000, -7, -2, -1000, -1000, -1000, 0, -1000, -31,
	-1000, 40, 79, -1000, -6, -1000, 67, 32, 90, -6,
	-1000, -3, -31, -1000, -6, -1000, -1000, -1000, -5, -1000,
	40, -1000, -1000, 86, -1000, -31, -9, -15, -1000, -1000,
	40, 38, -6, -6, -8, -1000, -1000, -1000, -16, -1000,
}
var yyPgo = [...]int{

	0, 0, 9, 139, 10, 138, 136, 135, 133, 5,
	131, 130, 129, 6, 128, 127, 126, 125, 2, 123,
	119, 118, 7, 1, 117, 116, 113,
}
var yyR1 = [...]int{

	0, 3, 11, 11, 10, 10, 10, 10, 17, 17,
	16, 16, 16, 16, 16, 16, 7, 7, 7, 15,
	15, 14, 14, 9, 9, 8, 8, 6, 6, 6,
	13, 13, 12, 24, 24, 25, 25, 26, 26, 4,
	4, 4, 4, 4, 5, 5, 5, 5, 5, 5,
	5, 5, 5, 18, 18, 18, 18, 18, 18, 18,
	18, 19, 19, 20, 20, 22, 22, 21, 21, 21,
	1, 2, 23, 23, 23,
}
var yyR2 = [...]int{

	0, 2, 0, 2, 3, 4, 4, 4, 0, 3,
	7, 6, 8, 8, 8, 11, 1, 1, 1, 0,
	3, 4, 6, 0, 3, 8, 10, 1, 1, 0,
	0, 3, 10, 1, 0, 1, 1, 0, 4, 3,
	8, 6, 6, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 4,
	4, 0, 3, 0, 6, 0, 3, 0, 6, 4,
	0, 0, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -3, -11, -17, -10, -1, -16, -1, 9, 8,
	-23, 45, 49, -2, 5, 4, 37, 4, 32, 24,
	31, -7, 30, 25, 26, 27, 5, 4, 4, -4,
	-1, -4, 4, 4, 4, 4, -5, 20, 21, 22,
	4, 11, 12, 13, 14, 15, 16, 17, 18, 19,
	4, 39, 39, 39, 28, 38, -22, 42, 44, 44,
	44, -22, -15, -9, -13, -1, -18, 6, 7, 35,
	36, 5, -1, -21, -4, -4, -4, 40, -14, -1,
	40, -8, -1, 40, -12, -2, 4, 4, 47, 39,
	43, -1, 45, 46, 46, -22, -23, -2, -22, -23,
	-2, -22, -23, -24, 23, 39, -19, -20, 4, -4,
	-22, -22, 4, 6, -25, 10, -4, -13, 48, -18,
	40, -1, 38, -23, 46, -22, 38, 41, -1, 40,
	-23, -18, 5, -22, 6, -6, 33, 34, 4, -22,
	41, -23, -22, -4, 42, -18, 4, -9, -23, -22,
	38, 43, -18, -26, 29, -22, -22, 42, -9, 43,
}
var yyDef = [...]int{

	2, -2, -2, -2, 3, 0, 74, 71, 0, 0,
	9, 72, 73, 0, 4, 0, 0, 0, 70, 70,
	0, 0, 0, 16, 17, 18, 5, 6, 7, 0,
	0, 0, 0, 0, 0, 0, 65, 0, 0, 0,
	43, 44, 45, 46, 47, 48, 49, 50, 51, 52,
	65, 19, 23, 30, 70, 70, 39, 67, 70, 70,
	70, 11, 70, 70, 71, 0, 10, 53, 54, 55,
	56, 57, 0, 70, 0, 0, 0, 65, 74, 71,
	65, 74, 71, 65, 74, 34, 0, 58, 61, 63,
	66, 0, 70, 65, 65, 12, 20, 0, 13, 24,
	0, 14, 31, 70, 33, 30, 70, 70, 74, 0,
	41, 42, 65, 0, 70, 35, 36, 71, 59, 74,
	60, 70, 0, 69, 65, 21, 0, 29, 0, 65,
	62, 0, 74, 40, 65, 70, 27, 28, 0, 15,
	70, 68, 22, 0, 23, 74, 65, 70, 64, 25,
	70, 37, 65, 65, 0, 26, 32, 23, 70, 38,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	42, 43, 37, 3, 45, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 41, 49,
	44, 38, 46, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 47, 3, 48, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 39, 3, 40,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.prog = &ast.Program{Headers: yyDollar[1].headers, Definitions: yyDollar[2].definitions}
			yylex.(*lexer).program = yyVAL.prog
			return 0
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.headers = nil
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.headers = append(yyDollar[1].headers, yyDollar[2].header)
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.header = &ast.Include{
				Path: yyDollar[3].str,
				Line: yyDollar[1].line,
			}
		}
	case 5:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.header = &ast.Include{
				Name: yyDollar[3].str,
				Path: yyDollar[4].str,
				Line: yyDollar[1].line,
			}
		}
	case 6:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.header = &ast.Namespace{
				Scope: "*",
				Name:  yyDollar[4].str,
				Line:  yyDollar[1].line,
			}
		}
	case 7:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.header = &ast.Namespace{
				Scope: yyDollar[3].str,
				Name:  yyDollar[4].str,
				Line:  yyDollar[1].line,
			}
		}
	case 8:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.definitions = nil
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.definitions = append(yyDollar[1].definitions, yyDollar[2].definition)
		}
	case 10:
		yyDollar = yyS[yypt-7 : yypt+1]
		{
			yyVAL.definition = &ast.Constant{
				Name:  yyDollar[5].str,
				Type:  yyDollar[4].fieldType,
				Value: yyDollar[7].constantValue,
				Line:  yyDollar[1].line,
				Doc:   ParseDocstring(yyDollar[2].docstring),
			}
		}
	case 11:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.definition = &ast.Typedef{
				Name:        yyDollar[5].str,
				Type:        yyDollar[4].fieldType,
				Annotations: yyDollar[6].typeAnnotations,
				Line:        yyDollar[1].line,
				Doc:         ParseDocstring(yyDollar[2].docstring),
			}
		}
	case 12:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.definition = &ast.Enum{
				Name:        yyDollar[4].str,
				Items:       yyDollar[6].enumItems,
				Annotations: yyDollar[8].typeAnnotations,
				Line:        yyDollar[1].line,
				Doc:         ParseDocstring(yyDollar[2].docstring),
			}
		}
	case 13:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.definition = &ast.Struct{
				Name:        yyDollar[4].str,
				Type:        yyDollar[3].structType,
				Fields:      yyDollar[6].fields,
				Annotations: yyDollar[8].typeAnnotations,
				Line:        yyDollar[1].line,
				Doc:         ParseDocstring(yyDollar[2].docstring),
			}
		}
	case 14:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.definition = &ast.Service{
				Name:        yyDollar[4].str,
				Functions:   yyDollar[6].functions,
				Annotations: yyDollar[8].typeAnnotations,
				Line:        yyDollar[1].line,
				Doc:         ParseDocstring(yyDollar[2].docstring),
			}
		}
	case 15:
		yyDollar = yyS[yypt-11 : yypt+1]
		{
			parent := &ast.ServiceReference{
				Name: yyDollar[7].str,
				Line: yyDollar[6].line,
			}

			yyVAL.definition = &ast.Service{
				Name:        yyDollar[4].str,
				Functions:   yyDollar[9].functions,
				Parent:      parent,
				Annotations: yyDollar[11].typeAnnotations,
				Line:        yyDollar[1].line,
				Doc:         ParseDocstring(yyDollar[2].docstring),
			}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.structType = ast.StructType
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.structType = ast.UnionType
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.structType = ast.ExceptionType
		}
	case 19:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.enumItems = nil
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.enumItems = append(yyDollar[1].enumItems, yyDollar[2].enumItem)
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.enumItem = &ast.EnumItem{
				Name:        yyDollar[3].str,
				Annotations: yyDollar[4].typeAnnotations,
				Line:        yyDollar[1].line,
				Doc:         ParseDocstring(yyDollar[2].docstring),
			}
		}
	case 22:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			value := int(yyDollar[5].i64)
			yyVAL.enumItem = &ast.EnumItem{
				Name:        yyDollar[3].str,
				Value:       &value,
				Annotations: yyDollar[6].typeAnnotations,
				Line:        yyDollar[1].line,
				Doc:         ParseDocstring(yyDollar[2].docstring),
			}
		}
	case 23:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.fields = nil
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.fields = append(yyDollar[1].fields, yyDollar[2].field)
		}
	case 25:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.field = &ast.Field{
				ID:           int(yyDollar[3].i64),
				Name:         yyDollar[7].str,
				Type:         yyDollar[6].fieldType,
				Requiredness: yyDollar[5].fieldRequired,
				Annotations:  yyDollar[8].typeAnnotations,
				Line:         yyDollar[1].line,
				Doc:          ParseDocstring(yyDollar[2].docstring),
			}
		}
	case 26:
		yyDollar = yyS[yypt-10 : yypt+1]
		{
			yyVAL.field = &ast.Field{
				ID:           int(yyDollar[3].i64),
				Name:         yyDollar[7].str,
				Type:         yyDollar[6].fieldType,
				Requiredness: yyDollar[5].fieldRequired,
				Default:      yyDollar[9].constantValue,
				Annotations:  yyDollar[10].typeAnnotations,
				Line:         yyDollar[1].line,
				Doc:          ParseDocstring(yyDollar[2].docstring),
			}
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.fieldRequired = ast.Required
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.fieldRequired = ast.Optional
		}
	case 29:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.fieldRequired = ast.Unspecified
		}
	case 30:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.functions = nil
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.functions = append(yyDollar[1].functions, yyDollar[2].function)
		}
	case 32:
		yyDollar = yyS[yypt-10 : yypt+1]
		{
			yyVAL.function = &ast.Function{
				Name:        yyDollar[5].str,
				Parameters:  yyDollar[7].fields,
				ReturnType:  yyDollar[3].fieldType,
				Exceptions:  yyDollar[9].fields,
				OneWay:      yyDollar[2].bul,
				Annotations: yyDollar[10].typeAnnotations,
				Line:        yyDollar[4].line,
				Doc:         ParseDocstring(yyDollar[1].docstring),
			}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.bul = true
		}
	case 34:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.bul = false
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.fieldType = nil
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.fieldType = yyDollar[1].fieldType
		}
	case 37:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.fields = nil
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.fields = yyDollar[3].fields
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.fieldType = ast.BaseType{ID: yyDollar[2].baseTypeID, Annotations: yyDollar[3].typeAnnotations, Line: yyDollar[1].line}
		}
	case 40:
		yyDollar = yyS[yypt-8 : yypt+1]
		{
			yyVAL.fieldType = ast.MapType{KeyType: yyDollar[4].fieldType, ValueType: yyDollar[6].fieldType, Annotations: yyDollar[8].typeAnnotations, Line: yyDollar[1].line}
		}
	case 41:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.fieldType = ast.ListType{ValueType: yyDollar[4].fieldType, Annotations: yyDollar[6].typeAnnotations, Line: yyDollar[1].line}
		}
	case 42:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.fieldType = ast.SetType{ValueType: yyDollar[4].fieldType, Annotations: yyDollar[6].typeAnnotations, Line: yyDollar[1].line}
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.fieldType = ast.TypeReference{Name: yyDollar[2].str, Line: yyDollar[1].line}
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.baseTypeID = ast.BoolTypeID
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.baseTypeID = ast.I8TypeID
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.baseTypeID = ast.I8TypeID
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.baseTypeID = ast.I16TypeID
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.baseTypeID = ast.I32TypeID
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.baseTypeID = ast.I64TypeID
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.baseTypeID = ast.DoubleTypeID
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.baseTypeID = ast.StringTypeID
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.baseTypeID = ast.BinaryTypeID
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.constantValue = ast.ConstantInteger(yyDollar[1].i64)
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.constantValue = ast.ConstantDouble(yyDollar[1].dub)
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.constantValue = ast.ConstantBoolean(true)
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.constantValue = ast.ConstantBoolean(false)
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.constantValue = ast.ConstantString(yyDollar[1].str)
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.constantValue = ast.ConstantReference{Name: yyDollar[2].str, Line: yyDollar[1].line}
		}
	case 59:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.constantValue = ast.ConstantList{Items: yyDollar[3].constantValues, Line: yyDollar[1].line}
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.constantValue = ast.ConstantMap{Items: yyDollar[3].constantMapItems, Line: yyDollar[1].line}
		}
	case 61:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.constantValues = nil
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.constantValues = append(yyDollar[1].constantValues, yyDollar[2].constantValue)
		}
	case 63:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.constantMapItems = nil
		}
	case 64:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.constantMapItems = append(yyDollar[1].constantMapItems, ast.ConstantMapItem{Key: yyDollar[3].constantValue, Value: yyDollar[5].constantValue, Line: yyDollar[2].line})
		}
	case 65:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.typeAnnotations = nil
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.typeAnnotations = yyDollar[2].typeAnnotations
		}
	case 67:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.typeAnnotations = nil
		}
	case 68:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			yyVAL.typeAnnotations = append(yyDollar[1].typeAnnotations, &ast.Annotation{Name: yyDollar[3].str, Value: yyDollar[5].str, Line: yyDollar[2].line})
		}
	case 69:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			yyVAL.typeAnnotations = append(yyDollar[1].typeAnnotations, &ast.Annotation{Name: yyDollar[3].str, Line: yyDollar[2].line})
		}
	case 70:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.line = yylex.(*lexer).line
		}
	case 71:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.docstring = yylex.(*lexer).LastDocstring()
		}
	}
	goto yystack /* stack new state and value */
}
