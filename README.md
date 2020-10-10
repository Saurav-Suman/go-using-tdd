# go-using-tdd

* Explore the Go language by writing tests
* Get a grounding with TDD. Go is a good language for learning TDD because it is a simple language to learn and testing is built-in
* Be confident that you'll be able to start writing robust, well-tested systems in Go
* Testing ( go test -v) , BenchMarking (go test -bench=.) , Code Coverage (go test -cover )

# run tests in all packages of the local project
go test ./...
# run tests only in package `slice` and its subpackages
go test ./slice/...
# run only tests that match pattern
go test ./... -run=TestSince/now
