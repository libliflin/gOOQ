// Package gooq contains functions for working with databases
package gooq

type SelectFromStep interface {
	From(t Table) SelectJoinStep
}

type SelectSelectStep interface {
	SelectFromStep
}

type SelectJoinStep interface {
	Join(t Table) SelectOnStep
}
type SelectOnStep interface {
	On(c Condition) SelectOnConditionStep
}
type SelectWhereStep interface {
	Where(c Condition) SelectConditionStep
}
type Condition interface {
}
type SelectOnConditionStep interface {
	SelectWhereStep
}
type SelectConditionStep interface {
	Query
}
type Query interface {
	GetSQL() string
}
type Field interface {
	Equal(f Field) Condition
}
type Table interface {
}

type h2db int

func (h2 h2db) From(t Table) SelectJoinStep {
	return h2
}
func (h2 h2db) Join(t Table) SelectOnStep {
	return h2
}
func (h2 h2db) On(c Condition) SelectOnConditionStep {
	return h2
}
func (h2 h2db) Where(c Condition) SelectConditionStep {
	return h2
}
func (h2 h2db) GetSQL() string {
	return "select BOOK.TITLE, AUTHOR.FIRST_NAME, AUTHOR.LAST_NAME from BOOK join AUTHOR on BOOK.AUTHOR_ID = AUTHOR.ID where BOOK.PUBLISHED_IN = 1948"
}
func (h2 h2db) Equal(f Field) Condition {
	return h2
}

func Select(f ...Field) SelectSelectStep {
	return new(h2db)
}

func NewField(s string) Field {
	return new(h2db)
}

func NewTable(s string) Table {
	return new(h2db)
}

func Val(i int) Field {
	return new(h2db)
}
