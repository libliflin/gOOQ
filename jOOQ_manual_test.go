// Package gooq contains functions for working with databases
package gooq

import "testing"

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
