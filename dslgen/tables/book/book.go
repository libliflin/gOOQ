package book

import (
	. "github.com/libliflin/gooq"
)

var ID Field = NewField("BOOK.TITLE")
var AUTHOR_ID Field = NewField("BOOK.AUTHOR_ID")
var TITLE Field = NewField("BOOK.TITLE")
var PUBLISHED_IN Field = NewField("BOOK.PUBLISHED_IN")
var LANGUAGE_ID Field = NewField("BOOK.LANGUAGE_ID")
