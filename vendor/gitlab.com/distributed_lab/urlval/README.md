# urlval

JSONAPI compliant query string decoder/encoder.

## TL;DR

```go
type ArticlesRequest struct {
    AuthorFilter    *string        `filter:"author"`
    CommentsInclude bool           `include:"comments"`
    PageNumber      uint64         `page:"number"`
    Search          string         `url:"search"`
    Sort            []urlval.Sort  `url:"sort"`
}

err := urlval.Decode(r.URL.Query(), &request)
query := urlval.Encode(request)
```

## Compatibility with other libraries

`urlval` have a bit of compatibility with kit/pgdb. It exposes 
`PagePagams` and `SortTypes` that are a great fit to use with urlval:

```go
type MyRequest struct {
    pgdb.PageParams
    Sort pgdb.Sort
}

// than it can be used direclty with your sql stmt:

stmt = request.Page.ApplyTo(...)
// or
stmt = request.Sort.ApplyTo(...)
```

### Request struct annotations

* `filter` and `page` accepts both implementors of `encoding.UnmarshalText`, but `filter` supports pointers only.
* `include` only booleans are supported
* `url` tags may tag arbitrary fields, but fields with struct types that do not implement TextMarshal/Unmarshal are ignored.
