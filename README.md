#Slice vs Map Testing

This code is to test making a map vs slice and performing a key lookup to
determine the the trade off points in speed.

## Running

This benchmark can be run by:

```
go test -bench=.
```

It is based on original work by Graham King

https://gist.github.com/grahamking/dbc90cb3f45c8fea2ba6

http://www.darkcoding.net/software/go-slice-search-vs-map-lookup/
