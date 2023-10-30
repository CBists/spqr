// Code generated by goyacc -o gram.go -p yy gram.y. DO NOT EDIT.

//line gram.y:3
package spqrparser

import __yyfmt__ "fmt"

//line gram.y:3

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

//line gram.y:23
type yySymType struct {
	yys   int
	str   string
	byte  byte
	bytes []byte
	int   int
	bool  bool
	empty struct{}

	statement Statement
	show      *Show

	drop   *Drop
	create *Create

	kill   *Kill
	lock   *Lock
	unlock *Unlock

	ds            *DataspaceDefinition
	kr            *KeyRangeDefinition
	shard         *ShardDefinition
	sharding_rule *ShardingRuleDefinition

	register_router   *RegisterRouter
	unregister_router *UnregisterRouter

	split *SplitKeyRange
	move  *MoveKeyRange
	unite *UniteKeyRange

	shutdown *Shutdown
	listen   *Listen

	trace     *TraceStmt
	stoptrace *StopTraceStmt

	entrieslist []ShardingRuleEntry
	shruleEntry ShardingRuleEntry

	sharding_rule_selector *ShardingRuleSelector
	key_range_selector     *KeyRangeSelector

	colref ColumnRef
	where  WhereClauseNode
}

const IDENT = 57346
const COMMAND = 57347
const SHOW = 57348
const KILL = 57349
const WHERE = 57350
const OR = 57351
const AND = 57352
const TEQ = 57353
const SCONST = 57354
const TSEMICOLON = 57355
const TOPENBR = 57356
const TCLOSEBR = 57357
const SHUTDOWN = 57358
const LISTEN = 57359
const REGISTER = 57360
const UNREGISTER = 57361
const ROUTER = 57362
const ROUTE = 57363
const CREATE = 57364
const ADD = 57365
const DROP = 57366
const LOCK = 57367
const UNLOCK = 57368
const SPLIT = 57369
const MOVE = 57370
const COMPOSE = 57371
const SHARDING = 57372
const COLUMN = 57373
const TABLE = 57374
const HASH = 57375
const FUNCTION = 57376
const KEY = 57377
const RANGE = 57378
const DATASPACE = 57379
const SHARDS = 57380
const KEY_RANGES = 57381
const ROUTERS = 57382
const SHARD = 57383
const HOST = 57384
const SHARDING_RULES = 57385
const RULE = 57386
const COLUMNS = 57387
const VERSION = 57388
const BY = 57389
const FROM = 57390
const TO = 57391
const WITH = 57392
const UNITE = 57393
const ALL = 57394
const ADDRESS = 57395
const CLIENT = 57396
const START = 57397
const STOP = 57398
const TRACE = 57399
const MESSAGES = 57400
const OP = 57401

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"COMMAND",
	"SHOW",
	"KILL",
	"WHERE",
	"OR",
	"AND",
	"TEQ",
	"SCONST",
	"TSEMICOLON",
	"TOPENBR",
	"TCLOSEBR",
	"SHUTDOWN",
	"LISTEN",
	"REGISTER",
	"UNREGISTER",
	"ROUTER",
	"ROUTE",
	"CREATE",
	"ADD",
	"DROP",
	"LOCK",
	"UNLOCK",
	"SPLIT",
	"MOVE",
	"COMPOSE",
	"SHARDING",
	"COLUMN",
	"TABLE",
	"HASH",
	"FUNCTION",
	"KEY",
	"RANGE",
	"DATASPACE",
	"SHARDS",
	"KEY_RANGES",
	"ROUTERS",
	"SHARD",
	"HOST",
	"SHARDING_RULES",
	"RULE",
	"COLUMNS",
	"VERSION",
	"BY",
	"FROM",
	"TO",
	"WITH",
	"UNITE",
	"ALL",
	"ADDRESS",
	"CLIENT",
	"START",
	"STOP",
	"TRACE",
	"MESSAGES",
	"OP",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line gram.y:624

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 183

var yyAct = [...]int{

	109, 63, 116, 106, 94, 100, 115, 26, 27, 79,
	50, 49, 77, 72, 78, 62, 113, 29, 28, 33,
	34, 72, 72, 20, 19, 23, 24, 25, 30, 31,
	72, 98, 89, 160, 159, 152, 143, 88, 121, 87,
	72, 118, 71, 133, 118, 75, 118, 149, 81, 73,
	136, 164, 32, 163, 123, 119, 21, 22, 119, 82,
	119, 104, 85, 86, 99, 61, 80, 74, 76, 102,
	90, 91, 56, 151, 93, 96, 139, 51, 92, 101,
	42, 103, 105, 103, 97, 43, 72, 41, 110, 111,
	112, 44, 54, 95, 158, 157, 120, 52, 114, 122,
	70, 124, 55, 57, 69, 40, 39, 38, 66, 67,
	68, 129, 37, 36, 95, 134, 84, 72, 137, 140,
	141, 135, 65, 142, 72, 144, 48, 47, 46, 145,
	64, 59, 147, 45, 107, 148, 126, 150, 137, 131,
	35, 128, 127, 126, 1, 153, 132, 146, 128, 127,
	154, 18, 155, 17, 156, 16, 15, 14, 12, 13,
	161, 162, 8, 9, 165, 166, 138, 117, 6, 5,
	4, 3, 7, 11, 10, 60, 58, 53, 2, 108,
	130, 125, 83,
}
var yyPact = [...]int{

	1, -1000, 100, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 50,
	50, -46, -47, 62, 37, 37, 127, 11, 118, -1000,
	37, 37, 37, 84, 80, -1000, -1000, -1000, -1000, -1000,
	-1000, 113, 5, 31, 18, -1000, -1000, -1000, -1000, -40,
	-49, -1000, 30, -1000, 4, -1000, 23, -1000, 108, -1000,
	118, 118, -1000, -1000, -1000, -1000, -9, -12, -18, 113,
	26, -1000, -1000, 82, 36, -19, 22, -53, 113, -1000,
	17, 9, 113, -1000, 120, -1000, -1000, 113, 113, 113,
	-37, -1000, -1000, 61, 15, 113, -10, 118, 12, 118,
	-1000, -1000, -1000, -1000, -1000, -1000, 139, 120, 135, -1000,
	-4, -1000, -1000, 118, 15, 13, -1000, 43, 113, 113,
	-1000, 118, -13, 118, -1000, 120, -1000, -1000, -1000, 132,
	118, -1000, -1000, 118, -1000, 10, 113, -1000, -1000, 39,
	-1000, -1000, -14, 118, -1000, 139, -1000, -1000, -1000, 113,
	-1000, 113, 118, 74, -1000, -1000, 73, -15, -16, 113,
	113, 16, 14, 113, 113, -1000, -1000,
}
var yyPgo = [...]int{

	0, 182, 3, 181, 180, 179, 1, 0, 178, 177,
	77, 176, 175, 174, 173, 172, 171, 170, 169, 168,
	112, 107, 106, 105, 6, 2, 4, 167, 166, 163,
	162, 159, 158, 157, 156, 155, 153, 151, 144, 140,
}
var yyR1 = [...]int{

	0, 38, 39, 39, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	6, 6, 7, 3, 3, 3, 4, 4, 5, 2,
	2, 2, 1, 1, 11, 12, 15, 15, 15, 15,
	16, 16, 16, 16, 18, 18, 19, 17, 17, 17,
	17, 13, 30, 20, 21, 21, 21, 21, 24, 24,
	25, 26, 26, 27, 27, 28, 28, 22, 22, 22,
	22, 23, 23, 29, 9, 10, 33, 14, 14, 34,
	35, 32, 31, 36, 37, 37,
}
var yyR2 = [...]int{

	0, 2, 0, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
	3, 3, 0, 2, 1, 1, 2, 4, 2, 4,
	2, 2, 2, 2, 4, 4, 3, 2, 2, 2,
	2, 3, 2, 2, 7, 6, 5, 4, 1, 2,
	2, 2, 0, 2, 2, 3, 0, 12, 11, 10,
	9, 5, 4, 2, 3, 3, 6, 3, 3, 4,
	4, 2, 1, 5, 3, 3,
}
var yyChk = [...]int{

	-1000, -38, -8, -16, -17, -18, -19, -15, -30, -29,
	-13, -14, -32, -31, -33, -34, -35, -36, -37, 23,
	22, 55, 56, 24, 25, 26, 6, 7, 17, 16,
	27, 28, 51, 18, 19, -39, 13, -20, -21, -22,
	-23, 37, 30, 35, 41, -20, -21, -22, -23, 57,
	57, -10, 35, -9, 30, -10, 35, -10, -11, 4,
	-12, 54, 4, -6, 12, 4, -10, -10, -10, 20,
	20, -7, 4, 44, 36, -7, 50, 52, 54, 58,
	36, 44, 36, -1, 8, -6, -6, 48, 49, 50,
	-7, -7, 52, -7, -26, 32, -7, 48, 50, 42,
	58, -7, 52, -7, 52, -7, -2, 14, -5, -7,
	-7, -7, -7, 53, -26, -24, -25, -27, 31, 45,
	-7, 48, -6, 42, -6, -3, 4, 10, 9, -2,
	-4, 4, 11, 47, -6, -24, 37, -25, -28, 33,
	-7, -7, -6, 49, -6, -2, 15, -6, -6, 37,
	-7, 34, 49, -6, -7, -7, -6, 21, 21, 49,
	49, -7, -7, 37, 37, -7, -7,
}
var yyDef = [...]int{

	0, -2, 2, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 18, 19, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 82,
	0, 0, 0, 0, 0, 1, 3, 40, 41, 42,
	43, 0, 0, 0, 0, 47, 48, 49, 50, 0,
	0, 36, 0, 38, 0, 52, 0, 73, 32, 34,
	0, 0, 35, 81, 20, 21, 0, 0, 0, 0,
	0, 53, 22, 62, 0, 0, 0, 0, 0, 46,
	0, 0, 0, 51, 0, 77, 78, 0, 0, 0,
	0, 84, 85, 62, 0, 0, 0, 0, 0, 0,
	44, 45, 37, 75, 39, 74, 33, 0, 0, 28,
	0, 79, 80, 0, 0, 57, 58, 66, 0, 0,
	61, 0, 0, 0, 72, 0, 23, 24, 25, 0,
	0, 26, 27, 0, 83, 56, 0, 59, 60, 0,
	63, 64, 0, 0, 71, 31, 29, 30, 76, 0,
	55, 0, 0, 0, 54, 65, 0, 0, 0, 0,
	0, 70, 69, 0, 0, 68, 67,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

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

	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:165
		{
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:166
		{
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:171
		{
			setParseTree(yylex, yyDollar[1].create)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:175
		{
			setParseTree(yylex, yyDollar[1].create)
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:179
		{
			setParseTree(yylex, yyDollar[1].trace)
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:183
		{
			setParseTree(yylex, yyDollar[1].stoptrace)
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:187
		{
			setParseTree(yylex, yyDollar[1].drop)
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:191
		{
			setParseTree(yylex, yyDollar[1].lock)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:195
		{
			setParseTree(yylex, yyDollar[1].unlock)
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:199
		{
			setParseTree(yylex, yyDollar[1].show)
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:203
		{
			setParseTree(yylex, yyDollar[1].kill)
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:207
		{
			setParseTree(yylex, yyDollar[1].listen)
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:211
		{
			setParseTree(yylex, yyDollar[1].shutdown)
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:215
		{
			setParseTree(yylex, yyDollar[1].split)
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:219
		{
			setParseTree(yylex, yyDollar[1].move)
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:223
		{
			setParseTree(yylex, yyDollar[1].unite)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:227
		{
			setParseTree(yylex, yyDollar[1].register_router)
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:231
		{
			setParseTree(yylex, yyDollar[1].unregister_router)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:236
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:240
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:245
		{
			yyVAL.str = string(yyDollar[1].str)
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:251
		{
			yyVAL.str = yyDollar[1].str
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:253
		{
			yyVAL.str = "AND"
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:255
		{
			yyVAL.str = "OR"
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:260
		{
			yyVAL.str = yyDollar[1].str
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:262
		{
			yyVAL.str = "="
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:268
		{
			yyVAL.colref = ColumnRef{
				ColName: yyDollar[1].str,
			}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:276
		{
			yyVAL.where = yyDollar[2].where
		}
	case 30:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:279
		{
			yyVAL.where = WhereClauseLeaf{
				ColRef: yyDollar[1].colref,
				Op:     yyDollar[2].str,
				Value:  yyDollar[3].str,
			}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:287
		{
			yyVAL.where = WhereClauseOp{
				Op:    yyDollar[2].str,
				Left:  yyDollar[1].where,
				Right: yyDollar[3].where,
			}
		}
	case 32:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:297
		{
			yyVAL.where = WhereClauseEmpty{}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:301
		{
			yyVAL.where = yyDollar[2].where
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:308
		{
			switch v := strings.ToLower(string(yyDollar[1].str)); v {
			case DatabasesStr, RoutersStr, PoolsStr, ShardsStr, BackendConnectionsStr, KeyRangesStr, ShardingRules, ClientsStr, StatusStr, VersionStr:
				yyVAL.str = v
			default:
				yyVAL.str = UnsupportedStr
			}
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:319
		{
			switch v := string(yyDollar[1].str); v {
			case ClientStr:
				yyVAL.str = v
			default:
				yyVAL.str = "unsupp"
			}
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:330
		{
			yyVAL.drop = &Drop{Element: yyDollar[2].key_range_selector}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:335
		{
			yyVAL.drop = &Drop{Element: &KeyRangeSelector{KeyRangeID: `*`}}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:339
		{
			yyVAL.drop = &Drop{Element: yyDollar[2].sharding_rule_selector}
		}
	case 39:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:344
		{
			yyVAL.drop = &Drop{Element: &ShardingRuleSelector{ID: `*`}}
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:350
		{
			yyVAL.create = &Create{Element: yyDollar[2].ds}
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:355
		{
			yyVAL.create = &Create{Element: yyDollar[2].sharding_rule}
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:360
		{
			yyVAL.create = &Create{Element: yyDollar[2].kr}
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:364
		{
			yyVAL.create = &Create{Element: yyDollar[2].shard}
		}
	case 44:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:371
		{
			yyVAL.trace = &TraceStmt{All: true}
		}
	case 45:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:374
		{
			yyVAL.trace = &TraceStmt{
				Client: yyDollar[4].str,
			}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:382
		{
			yyVAL.stoptrace = &StopTraceStmt{}
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:389
		{
			yyVAL.create = &Create{Element: yyDollar[2].ds}
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:394
		{
			yyVAL.create = &Create{Element: yyDollar[2].sharding_rule}
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:399
		{
			yyVAL.create = &Create{Element: yyDollar[2].kr}
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:403
		{
			yyVAL.create = &Create{Element: yyDollar[2].shard}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:410
		{
			yyVAL.show = &Show{Cmd: yyDollar[2].str, Where: yyDollar[3].where}
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:416
		{
			yyVAL.lock = &Lock{KeyRangeID: yyDollar[2].key_range_selector.KeyRangeID}
		}
	case 53:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:424
		{
			yyVAL.ds = &DataspaceDefinition{ID: yyDollar[2].str}
		}
	case 54:
		yyDollar = yyS[yypt-7 : yypt+1]
//line gram.y:430
		{
			yyVAL.sharding_rule = &ShardingRuleDefinition{ID: yyDollar[3].str, TableName: yyDollar[4].str, Entries: yyDollar[5].entrieslist, Dataspace: yyDollar[7].str}
		}
	case 55:
		yyDollar = yyS[yypt-6 : yypt+1]
//line gram.y:435
		{
			str, err := randomHex(6)
			if err != nil {
				panic(err)
			}
			yyVAL.sharding_rule = &ShardingRuleDefinition{ID: "shrule" + str, TableName: yyDollar[3].str, Entries: yyDollar[4].entrieslist, Dataspace: yyDollar[6].str}
		}
	case 56:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:444
		{
			yyVAL.sharding_rule = &ShardingRuleDefinition{ID: yyDollar[3].str, TableName: yyDollar[4].str, Entries: yyDollar[5].entrieslist, Dataspace: "default"}
		}
	case 57:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:449
		{
			str, err := randomHex(6)
			if err != nil {
				panic(err)
			}
			yyVAL.sharding_rule = &ShardingRuleDefinition{ID: "shrule" + str, TableName: yyDollar[3].str, Entries: yyDollar[4].entrieslist, Dataspace: "default"}
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:458
		{
			yyVAL.entrieslist = make([]ShardingRuleEntry, 0)
			yyVAL.entrieslist = append(yyVAL.entrieslist, yyDollar[1].shruleEntry)
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:464
		{
			yyVAL.entrieslist = append(yyDollar[1].entrieslist, yyDollar[2].shruleEntry)
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:470
		{
			yyVAL.shruleEntry = ShardingRuleEntry{
				Column:       yyDollar[1].str,
				HashFunction: yyDollar[2].str,
			}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:479
		{
			yyVAL.str = yyDollar[2].str
		}
	case 62:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:482
		{
			yyVAL.str = ""
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:486
		{
			yyVAL.str = yyDollar[2].str
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:491
		{
			yyVAL.str = yyDollar[2].str
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:497
		{
			yyVAL.str = yyDollar[3].str
		}
	case 66:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:500
		{
			yyVAL.str = ""
		}
	case 67:
		yyDollar = yyS[yypt-12 : yypt+1]
//line gram.y:505
		{
			yyVAL.kr = &KeyRangeDefinition{LowerBound: []byte(yyDollar[5].str), UpperBound: []byte(yyDollar[7].str), ShardID: yyDollar[10].str, KeyRangeID: yyDollar[3].str, Dataspace: yyDollar[12].str}
		}
	case 68:
		yyDollar = yyS[yypt-11 : yypt+1]
//line gram.y:510
		{
			str, err := randomHex(6)
			if err != nil {
				panic(err)
			}
			yyVAL.kr = &KeyRangeDefinition{LowerBound: []byte(yyDollar[4].str), UpperBound: []byte(yyDollar[6].str), ShardID: yyDollar[9].str, KeyRangeID: "kr" + str, Dataspace: yyDollar[11].str}
		}
	case 69:
		yyDollar = yyS[yypt-10 : yypt+1]
//line gram.y:519
		{
			yyVAL.kr = &KeyRangeDefinition{LowerBound: []byte(yyDollar[5].str), UpperBound: []byte(yyDollar[7].str), ShardID: yyDollar[10].str, KeyRangeID: yyDollar[3].str, Dataspace: "default"}
		}
	case 70:
		yyDollar = yyS[yypt-9 : yypt+1]
//line gram.y:524
		{
			str, err := randomHex(6)
			if err != nil {
				panic(err)
			}
			yyVAL.kr = &KeyRangeDefinition{LowerBound: []byte(yyDollar[4].str), UpperBound: []byte(yyDollar[6].str), ShardID: yyDollar[9].str, KeyRangeID: "kr" + str, Dataspace: "default"}
		}
	case 71:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:535
		{
			yyVAL.shard = &ShardDefinition{Id: yyDollar[2].str, Hosts: []string{yyDollar[5].str}}
		}
	case 72:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:540
		{
			str, err := randomHex(6)
			if err != nil {
				panic(err)
			}
			yyVAL.shard = &ShardDefinition{Id: "shard" + str, Hosts: []string{yyDollar[4].str}}
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:551
		{
			yyVAL.unlock = &Unlock{KeyRangeID: yyDollar[2].key_range_selector.KeyRangeID}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:557
		{
			yyVAL.sharding_rule_selector = &ShardingRuleSelector{ID: yyDollar[3].str}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:563
		{
			yyVAL.key_range_selector = &KeyRangeSelector{KeyRangeID: yyDollar[3].str}
		}
	case 76:
		yyDollar = yyS[yypt-6 : yypt+1]
//line gram.y:569
		{
			yyVAL.split = &SplitKeyRange{KeyRangeID: yyDollar[2].key_range_selector.KeyRangeID, KeyRangeFromID: yyDollar[4].str, Border: []byte(yyDollar[6].str)}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:575
		{
			yyVAL.kill = &Kill{Cmd: yyDollar[2].str, Target: yyDollar[3].str}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:578
		{
			yyVAL.kill = &Kill{Cmd: "client", Target: yyDollar[3].str}
		}
	case 79:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:584
		{
			yyVAL.move = &MoveKeyRange{KeyRangeID: yyDollar[2].key_range_selector.KeyRangeID, DestShardID: yyDollar[4].str}
		}
	case 80:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:590
		{
			yyVAL.unite = &UniteKeyRange{KeyRangeIDL: yyDollar[2].key_range_selector.KeyRangeID, KeyRangeIDR: yyDollar[4].str}
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:596
		{
			yyVAL.listen = &Listen{addr: yyDollar[2].str}
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:602
		{
			yyVAL.shutdown = &Shutdown{}
		}
	case 83:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:610
		{
			yyVAL.register_router = &RegisterRouter{ID: yyDollar[3].str, Addr: yyDollar[5].str}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:616
		{
			yyVAL.unregister_router = &UnregisterRouter{ID: yyDollar[3].str}
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:621
		{
			yyVAL.unregister_router = &UnregisterRouter{ID: `*`}
		}
	}
	goto yystack /* stack new state and value */
}
