// Package gooq_test contains functions for testing gooq
package gooq_test

import (
	. "github.com/libliflin/gooq"
	. "github.com/libliflin/gooq/dslgen/schemas/default_schema"
	"github.com/libliflin/gooq/dslgen/tables/author"
	"github.com/libliflin/gooq/dslgen/tables/book"
	"testing"
)

// Test331gooqAsSQLBuilder tests Section 3.3.1 jOOQ as a SQL Builder.
// Found in jOOQ 3.6 User Manual.
// http://www.jooq.org/doc/3.6/manual-single-page/#jooq-as-a-standalone-sql-builder
func Test331gooqAsSQLBuilder(t *testing.T) {
	// Fetch a SQL string from a jOOQ Query in order to manually execute it with another tool.
	//String sql = create.select(field("BOOK.TITLE"), field("AUTHOR.FIRST_NAME"), field("AUTHOR.LAST_NAME"))
	//                   .from(table("BOOK"))
	//                   .join(table("AUTHOR"))
	//                   .on(field("BOOK.AUTHOR_ID").equal(field("AUTHOR.ID")))
	//                   .where(field("BOOK.PUBLISHED_IN").equal(1948))
	//                   .getSQL();
	got := Select(NewField("BOOK.TITLE"), NewField("AUTHOR.FIRST_NAME"), NewField("AUTHOR.LAST_NAME")).
		From(NewTable("BOOK")).
		Join(NewTable("AUTHOR")).
		On(NewField("BOOK.AUTHOR_ID").Equal(NewField("AUTHOR_ID"))).
		Where(NewField("BOOK.PUBLISHED_IN").Equal(Val(1948))).
		GetSQL()
	want := "select BOOK.TITLE, AUTHOR.FIRST_NAME, AUTHOR.LAST_NAME from BOOK join AUTHOR on BOOK.AUTHOR_ID = AUTHOR.ID where BOOK.PUBLISHED_IN = 1948"
	if got != want {
		t.Errorf("Select == %q, want %q", got, want)
	}
}

// Test332gooqAsSQLBuilderWithCodeGeneration tests Section 3.3.2 jOOQ as a SQL builder with code generation.
// Found in jOOQ 3.6 User Manual.
// http://www.jooq.org/doc/3.6/manual-single-page/#jooq-as-a-sql-builder-with-code-generation
func Test332gooqAsSQLBuilderWithCodeGeneration(t *testing.T) {
	// Fetch a SQL string from a jOOQ Query in order to manually execute it with another tool.
	// String sql = create.select(BOOK.TITLE, AUTHOR.FIRST_NAME, AUTHOR.LAST_NAME)
	//                    .from(BOOK)
	//                    .join(AUTHOR)
	//                    .on(BOOK.AUTHOR_ID.equal(AUTHOR.ID))
	//                    .where(BOOK.PUBLISHED_IN.equal(1948))
	//                    .getSQL();
	got := Select(book.TITLE, author.FIRST_NAME, author.LAST_NAME).
		From(BOOK).
		Join(AUTHOR).
		On(book.AUTHOR_ID.Equal(author.AUTHOR_ID)).
		Where(book.PUBLISHED_IN.Equal(Val(1948))).
		GetSQL()
	want := "select BOOK.TITLE, AUTHOR.FIRST_NAME, AUTHOR.LAST_NAME from BOOK join AUTHOR on BOOK.AUTHOR_ID = AUTHOR.ID where BOOK.PUBLISHED_IN = 1948"
	if got != want {
		t.Errorf("Select == %q, want %q", got, want)
	}
}
