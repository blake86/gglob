/*
@Time : 2020-04-15 10:53
@Author : lihao
@File : Range
@Software: GoLand
*/
package gglob

import (
	"fmt"
	"github.com/blake86/gglob/lex"
	"sort"
	"strconv"
)

func Compress(names []string) []string {
	return nil

}

/*
Usage: prefix{\}
example:
pre{001-3}post == pre{001-003}post == pre001post,pre002post,pre003post
pre{1,3}post == pre1post, pre3post
*/
//func Expand(str string) []string{
//	hyphenRegex:=`{([\d,-]*)}`
//	re:=regexp.MustCompile(hyphenRegex)
//
//	pads:=re.Split(str,-1)
//	matchs:=re.FindAllStringSubmatch(str,-1)
//
//	var exps [][]interface{}
//
//	for _,match:=range matchs{
//		var exp []interface{}
//
//		parts:=strings.Split(match[1],",")
//		for _,r:=range parts{
//			hyps:=strings.Split(r,"-")
//			if len(hyps) ==2 {
//				head,err:=strconv.Atoi(hyps[0])
//				if err!=nil{
//					return nil
//				}
//				tail,err:=strconv.Atoi(hyps[1])
//				if err!=nil{
//					return nil
//				}
//				format:="%0"+strconv.Itoa(len(hyps[0]))+"d"
//				for i:=head;i<=tail;i++{
//					exp=append(exp,fmt.Sprintf(format,i))
//				}
//			}else if len(hyps) == 1{
//				exp = append(exp,r)
//			}
//
//		}
//		exps = append(exps,exp)
//
//	}
//
//	var result []string
//	c := cartesian.Iter(exps[:]...)
//	for product := range c {
//		str:=""
//		for i:=0;i<len(pads)-1;i++{
//			str=str+pads[i]+product[i].(string)
//		}
//		str=str+pads[len(pads)-1]
//		result = append(result,str)
//
//	}
//
//	sort.Strings(result)
//	fmt.Println(result)
//	return result
//}

/*
prefix[1,3-12,14]mid[01-123,3-5]suffix

001 -> 最少3位
0010 -> 最少4位

preifx -> {1, 3-12, 14} -> mid ...

*/

const (
	leftCurl    = '{'
	rightCurl   = '}'
	leftRound   = '('
	rightRound  = ')'
	leftSquare  = '['
	rightSquare = ']'
	period      = '.'
	comma       = ','
	slash       = '/'
	equal       = '='
	quote       = '"'
	at          = '@'
	colon       = ':'
	lsThan      = '<'
	gtThan      = '>'
	star        = '*'
	native      = '-'
)

// Constants representing type of different graphql lexed items.
const (
	itemText lex.ItemType = 5 + iota // plain text
	itemComma
	itemFuncStart
	itemFuncEnd
	itemArgStart
	itemArgEnd
	itemNative
)

func StartNum(str string) (n int, num int, err error) {
	num, err = strconv.Atoi(str)
	if err != nil {
		return 0, 0, err
	}
	n = len(str)
	return
}

func lexRule(l *lex.Lexer) lex.StateFn {
	l.Mode = lexRule
	//l.ModeStack = append(l.ModeStack, lexRule)

	for {
		switch r := l.Next(); {
		case lex.IsNumber(r):
			l.AcceptRun(lex.IsNumber)
			l.Emit(itemArgStart)
			l.AcceptRun(lex.IsSpace)
			l.Ignore()
			if l.Peek() == native {
				l.Next()
				l.Emit(itemNative)
				l.AcceptRun(lex.IsSpace)
				l.Ignore()
				if !lex.IsNumber(l.Peek()) {
					return l.Errorf("'-' need number")
				}
				l.AcceptRun(lex.IsNumber)
				l.Emit(itemArgEnd)
			}
		case lex.IsSpace(r): // lex.IsEndOfLine(r) 不可换行
			l.Ignore()
		case r == comma:
			l.Emit(itemComma)
			l.AcceptRun(lex.IsSpace)
			l.Ignore()

			if !lex.IsNumber(l.Peek()) {
				return l.Errorf("',' need number")
			}

			////if l.Peek() == at {
			////	l.Next()
			////	if l.Next() == leftRound {
			////		l.Emit(itemFuncStart)
			////		l.BlockDepth++
			////		l.ModeStack = append(l.ModeStack, lexRule)
			////		return lexRule
			////	}
			////} else
			//if l.Peek() == lsThan {
			//	l.Next()
			//	if l.Next() == leftCurl {
			//		l.Emit(itemArgStart)
			//		l.ArgDepth++
			//		l.ModeStack = append(l.ModeStack, lexArgs)
			//		return lexArgs
			//	}
			//}
			//return l.Errorf("expect function/arg after comma")
			////return lexArgs(l)
		case r == rightSquare:
			l.Emit(itemFuncEnd)
			l.BlockDepth--
			if l.BlockDepth < 0 {
				return l.Errorf("']' to many")
			}
			return lexTopLevel

			//if l.ArgDepth > 0 {
			//	return l.Errorf("'-' has no end")
			//} else if l.ArgDepth == 0 {
			//	return lexTopLevel
			//} else {
			//	return l.Errorf("--")
			//}

			//if len(l.ModeStack) == 0 {
			//	return lexTopLevel(l)
			//}
			//last := l.ModeStack[len(l.ModeStack)-1]
			//l.ModeStack = l.ModeStack[:len(l.ModeStack)-1]
			//return last

			//if l.BlockDepth == 0 {
			//	return lexTopLevel(l)
			//} else if l.BlockDepth > 0 {
			//	return lexRule(l)
			//} else {
			//	return l.Errorf("", )
			//}
		//case r == at:
		//	if l.Peek() == leftRound {
		//		l.Backup()
		//		l.Emit(itemText)
		//		l.Next()
		//		l.Emit(itemAt)
		//		l.Next()
		//		l.Emit(itemLeftRound)
		//		l.BlockDepth++
		//		l.ModeStack = append(l.ModeStack, lexRule)
		//		return lexRule
		//	}
		//case r == lsThan:
		//	if l.Peek() == leftCurl {
		//		l.Backup()
		//
		//	}
		default:
			return l.Errorf("in rule function, expect ")
		}

	}
}

func lexTopLevel(l *lex.Lexer) lex.StateFn {
	// TODO(Aman): Find a way to identify different blocks in future. We only have
	// Upsert block right now. BlockDepth tells us nesting of blocks. Currently, only
	// the Upsert block has nested mutation/query/fragment blocks.
	//if l.BlockDepth != 0 {
	//	return lexUpsertBlock
	//}

	l.Mode = lexTopLevel
Loop:
	for {
		switch r := l.Next(); {
		//case r == leftCurl:
		//	l.Depth++ // one level down.
		//	l.Emit(itemLeftCurl)
		//	return lexQuery
		//case r == rightCurl:
		//	return l.Errorf("Too many right curl")

		case r == lex.EOF:
			break Loop
		//case r == '#':
		//	return lexComment
		case r == leftSquare:
			l.Backup()
			l.Emit(itemText)
			l.Next()
			l.Emit(itemFuncStart)
			l.BlockDepth++
			l.ModeStack = append(l.ModeStack, lexRule)
			return lexRule
		}
	}
	if l.Pos > l.Start {
		l.Emit(itemText)
	}
	l.Emit(lex.ItemEOF)
	return nil
}

type SegImpl interface {
	Segs() []string
}
type StrSeg string

func (ss StrSeg) Segs() []string {
	return []string{string(ss)}
}

type numSegSet map[int]int // num -> n
func (ns numSegSet) Add(num, n int) {
	if nv, ok := ns[num]; !ok || nv < n {
		ns[num] = n
	}
}
func (ns numSegSet) Segs() (l []string) {
	for num, n := range ns {
		l = append(l, fmt.Sprintf("%0*d", n, num))
	}
	return
}

type ListSeg []SegImpl

func (l ListSeg) Segs() (ret []string) {
	if len(l) == 0 {
		return nil
	}
	ret = l[0].Segs()
	if len(l) == 1 {
		return ret
	}

	for _, seg := range l[1:] {
		ss := seg.Segs()
		mid := make([]string, 0, len(ret)*len(ss))
		for _, ls := range ret {
			for _, s := range ss {
				mid = append(mid, ls+s)
			}
		}
		ret = mid
	}
	return
}

func parseFunc(it *lex.ItemIterator) (numSegSet, error) {
	var m = numSegSet{}
	var n, start, end int
	var err error
	for it.Next() {
		item := it.Item()
		switch item.Typ {
		case itemArgStart:
			n, start, err = StartNum(item.Val)
			if err != nil {
				return nil, item.Errorf("argsStart=%s fail:%v", item.Val, err)
			}
			end = start
			item2, ok := it.PeekOne()
			if ok && item2.Typ == itemNative {
				it.Next()
				it.Next()
				item2 = it.Item()
				end, err = strconv.Atoi(item2.Val)
				if err != nil {
					return nil, item2.Errorf("'-' need number, but got %v", item2.Val)
				}
			}
			it.Next()
			item2 = it.Item()
			var isReturn bool
			switch item2.Typ {
			case itemComma:
			case itemFuncEnd:
				isReturn = true
			default:
				return nil, item2.Errorf("startNumber got unexpect type:%v(%v)", item2.Typ, item2.Val)
			}
			if end < start {
				return nil, item2.Errorf("start=%d end=%d, need start <= end", start, end)
			}
			for i := start; i <= end; i++ {
				m.Add(i, n)
			}
			if isReturn {
				return m, nil
			}
		default:
			return nil, it.Errorf("unexpect type:%v(%V)", item.Typ, item.Val)
		}
	}
	return nil, it.Errorf("unexpect end")
}

func parseRoot(it *lex.ItemIterator) ([]string, error) {
	var list ListSeg
Loop:
	for it.Next() {
		item := it.Item()
		switch item.Typ {
		case itemText:
			list = append(list, StrSeg(item.Val))
		case itemFuncStart:
			m, err := parseFunc(it)
			if err != nil {
				return nil, err
			}
			list = append(list, m)
		case lex.ItemEOF:
			break Loop
		default:
			return nil, item.Errorf("unexpect item:%+v", item)
		}
	}
	ls := list.Segs()
	sort.Slice(ls, func(i, j int) bool {
		return ls[i] < ls[j]
	})
	return ls, nil
}

func Expand(str string) ([]string, error) {
	var lexer lex.Lexer
	lexer.Reset(str)
	lexer.Run(lexTopLevel)
	if err := lexer.ValidateResult(); err != nil {
		return nil, err
	}
	it := lexer.NewIterator()
	return parseRoot(it)
}
