// Package gooq contains functions for working with databases
package gOOQ

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
	got := gooq.Select(gooq.Field("BOOK.TITLE"), gooq.Field("AUTHOR.FIRST_NAME"), gooq.Field("AUTHOR.LAST_NAME"))
						.From(gooq.Table("BOOK"))
						.Join(gooq.Table("AUTHOR"))
						.On(gooq.Field("BOOK.AUTHOR_ID").Equal(gooq.Field("AUTHOR_ID")))
						.Where(gooq.Field("BOOK.PUBLISHED_IN").Equal(1948))
						.GetSQL()
	want := "" // having trouble getting this, probably because I'm trying to do it in scala. 
	if got != want {
		t.Errorf("gooq.Select == %q, want %q", got, want)
	}
}
