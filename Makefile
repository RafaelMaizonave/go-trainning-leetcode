.PHONY: list start done reset stats readme review test

ID ?=

list:
	go run ./cmd/leetcode list

start:
	@test -n "$(ID)" || (echo "use: make start ID=1" && exit 1)
	go run ./cmd/leetcode start $(ID)

done:
	@test -n "$(ID)" || (echo "use: make done ID=1" && exit 1)
	go run ./cmd/leetcode done $(ID)

reset:
	@test -n "$(ID)" || (echo "use: make reset ID=1" && exit 1)
	go run ./cmd/leetcode reset $(ID)

stats:
	go run ./cmd/leetcode stats

review:
	go run ./cmd/leetcode review $(ID)

readme:
	go run ./cmd/leetcode readme

test:
	go test ./...
