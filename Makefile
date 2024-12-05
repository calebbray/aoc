build:
	go build -o bin/aoc

run: build
	@./bin/aoc $(-day)

test:
	go test -v ./...
