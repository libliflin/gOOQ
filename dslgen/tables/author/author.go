package author

import (
	. "github.com/libliflin/gooq"
)

var AUTHOR_ID Field = NewField("AUTHOR.TITLE")
var FIRST_NAME Field = NewField("AUTHOR.FIRST_NAME")
var LAST_NAME Field = NewField("AUTHOR.LAST_NAME")
var DATE_OF_BIRTH Field = NewField("AUTHOR.DATE_OF_BIRTH")
var YEAR_OF_BIRTH Field = NewField("AUTHOR.YEAR_OF_BIRTH")
var DISTINGUISHED Field = NewField("AUTHOR.DISTINGUISHED")
