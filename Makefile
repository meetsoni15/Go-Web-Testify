cover:
	go test -cover ./...

coverprofile:
	go test -coverprofile=cover.out ./...

coverhtml:
	go tool cover -html=cover.out 

test:
	go test -v ./...

all: coverprofile coverhtml	 