GOSRC:=$(shell find . -name '*.go')

.PHONY: depgraph.svg

README.md: README.md~ depgraph.svg
	cp $< $@

README.html: README.md
	pandoc -f gfm -t html -s -M title="qbsllm – README" README.md > README.html

deps: depgraph.svg

depgraph.svg:
	go mod graph | gomodot -d v | dot -Tsvg -o $@

# → https://blog.golang.org/cover
cover: coverage.html

benchmark:
	go test -bench=.

cpuprof:
	go test -cpuprofile cpu.prof -bench BenchmarkExpandArgs
	go tool pprof -http :6060 qbsllm.test cpu.prof
# Read with '$ go tool pprof cpu.prof' >>> e.g. '(pprof) web'

coverage.html: coverage.out
	go tool cover -html=$< -o $@

coverage.out: $(GOSRC)
	go test -coverprofile=$@ ./... || true
#	go test -covermode=count -coverprofile=$@ || true

cov=$(shell go tool cover -func=coverage.out \
            | egrep '^total:' \
            | awk '{print $$3}' \
            | tr "%" " ")

README.md~: coverage.html
	awk -v cov=$(cov) '/^\[!\[Test Coverage]/{ \
        cov=sprintf("%.0f", cov); \
		printf "[![Test Coverage](https://img.shields.io/badge/coverage-"; \
		printf "%s%%25-", cov ;\
		if (cov < 50) { \
			printf "red" \
	    } else if (cov < 75) { \
			printf "orange" \
	    } else if (cov < 90) { \
			printf "yellow" \
	    } else { \
			printf "green" \
		}; \
        print ".svg)](file:coverage.html)" ;\
     	next \
	} \
	{ print }' README.md > $@
