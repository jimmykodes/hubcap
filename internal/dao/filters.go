package dao

import (
	"fmt"
	"strings"
)

const (
	like = "like"
	eq   = "="
)

type fields map[string]bool
type SearchFilters map[string]interface{}

func (sf SearchFilters) whereClause(searchFields, filterFields fields, startingI int) *whereClause {
	wc := &whereClause{
		i: startingI,
	}
	for key, val := range sf {
		if _, ok := searchFields[key]; ok {
			wc.or(key, like)
			wc.addArg(fmt.Sprintf("%%%v%%", val))
		} else if _, ok := filterFields[key]; ok {
			wc.and(key, eq)
			wc.addArg(val)
		}
	}
	return wc
}

type whereClause struct {
	firstOpWritten bool
	i              int
	qb             strings.Builder
	args           []interface{}
}

func (wc *whereClause) query() string {
	return wc.qb.String()
}

func (wc *whereClause) writeClause(clause, operand string) {
	if !wc.firstOpWritten {
		operand = ""
		wc.firstOpWritten = true
	}
	wc.i++
	wc.qb.WriteString(fmt.Sprintf("%s %s\n", operand, clause))
}

func (wc *whereClause) and(f, operand string) {
	wc.writeClause(format(f, operand, wc.i), "and")
}
func (wc *whereClause) or(f, operand string) {
	wc.writeClause(format(f, operand, wc.i), "or")
}

func (wc *whereClause) addArg(arg interface{}) {
	wc.args = append(wc.args, arg)
}
func format(f, operand string, i int) string {
	return fmt.Sprintf("`%s` %s $%d", f, operand, i)
}
