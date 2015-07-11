# gOOQ
Schema generated go-lang interfaces coupled with type-safe query language. (aka, a handicapped jOOQ ripoff for go)

## this project is currently under design.

For nulls, go has a [comma ok](https://golang.org/doc/effective_go.html#maps) idiom, which is why each field access 
will need to be a method. 

Go does not have constructors, so need to use factory methods.

For schema APIs, methods will be collected into interfaces. 

Since go doesn't have generics, a schema specific version of the query dsl will be need to be created.

Current Plan:

1. Do a pass by hand.
  1. Use [jOOQ's manual's db](http://www.jooq.org/doc/3.6/manual-single-page/#sample-database) 
  2. Create go apis as close to their manual as go will allow. 
2. Do a SQL-gen prototype.
  1. A manual Java to Go like [rsc's c2go](https://github.com/rsc/c2go) of [jOOQ's apache licensed stuff](http://www.jooq.org/javadoc/latest/org/jooq/Select.html).
  2. minimial required to get api done previously executing. 
3. Do a code-gen prototype.
  1. Connect to MySQL setup of the db from 1 and get the api written generated.
4. Polish up code-gen prototype with [WebERP](http://www.weberp.org/)'s schema
5. Polish up the SQL-gen with queries translated from [WebERP](http://www.weberp.org/)'s PHP.
6. Implement preparedStatment caches, batching, connection pooling and other performance items.


Notes:
* go get github.com/libliflin/gooq
* go build ./...
* go test

URLs:
[Effective Go](https://golang.org/doc/effective_go.html)
[libliflin/gOOQ](https://github.com/libliflin/gOOQ)
[How to Write Go Code](https://golang.org/doc/code.html)
[The jOOQ User Manual. Single Page.](http://www.jooq.org/doc/3.6/manual-single-page)
[DSL (jOOQ 3.6.2 API)](http://www.jooq.org/javadoc/3.6.x/org/jooq/impl/DSL.html)
[Getting Started - Go](https://golang.org/doc/install)
[The Go Programming Language Specification](https://golang.org/ref/spec)