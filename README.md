go test .
go test -v .
go test -cover .

go test -coverprofile=coverage.out
go tool cover -html=coverage.out 

go test -coverprofile=coverage.out && go tool cover -html=coverage.out

go test -run <name_of_the_testing_func>
go test -run <Test_isPrime>
