include makefile.dc.mk
include makefile.build.mk
include makefile.lint.mk
include makefile.migration.mk

GO_COVER_EXCLUDE="config|service.go|mock.go|mocks|pkg/storage/gen"
PG_DSN="host=localhost port=6432 dbname=wallet user=postgres password=postgres sslmode=disable"
LOCAL_BIN:=$(CURDIR)/bin

.PHONY: run
run:
	PG_DSN=$(PG_DSN) \
	go run ./cmd/wallet/main.go -debug

.PHONY: test
test:
	PG_DSN=$(PG_DSN) \
	$(LOCAL_BIN)/gotestsum --junitfile unit-tests.xml \
		--jsonfile json-report.txt \
		-- -count 3 -p 3 -race -cover -coverpkg=./internal/... -coverprofile=cover.tmp.out -covermode atomic ./internal/...

.PHONY: test_ci
test_ci:
	PG_DSN=$(PG_DSN_CI) \
	$(LOCAL_BIN)/gotestsum --junitfile unit-tests.xml \
		--jsonfile json-report.txt \
		-- -count 3 -p 3 -race -cover -coverpkg=./internal/... -coverprofile=cover.tmp.out -covermode atomic ./internal/...
	grep -vE ${GO_COVER_EXCLUDE} cover.tmp.out > cover.out || cp cover.tmp.out cover.out
	go tool cover -func=cover.out

.PHONY: e2e
e2e:
	PG_DSN=$(PG_DSN) \
	$(LOCAL_BIN)/gotestsum -f testname -- -count 3 -p 3 ./test/...

.PHONY: e2e_ci
e2e_ci: tags=$tags
e2e_ci:
	PG_DSN=$(PG_DSN_CI) \
	$(LOCAL_BIN)/gotestsum -f testname -- -count 3 -p 3 -tags=${tags} ./test/...

.PHONY: test-cov
test-cov:
	make test
	grep -vE ${GO_COVER_EXCLUDE} cover.tmp.out > cover.out || cp cover.tmp.out cover.out
	go tool cover -func=cover.out
	go tool cover -html=cover.out
	go tool cover -html=cover.out -o cover.html

.PHONY: start-wiremock
start-wiremock:
	@echo "Starting Wiremock..."
	@docker run --rm -d --name wiremock -p 8070:8080 wiremock/wiremock

.PHONY: test-cbr
test-cbr: start-wiremock
	ENV=test make test
	@echo "Stopping Wiremock..."
	@docker stop wiremock

.PHONY: test-cov-cbr
test-cov-cbr: start-wiremock
	ENV=test make test-cov
	@echo "Stopping Wiremock..."
	@docker stop wiremock

.PHONY: prom-refresh
prom-refresh:
	curl 'http://localhost:8428/-/reload'

.PHONY: prom-config
prom-config:
	curl 'http://localhost:8428/config'

.PHONY: prom-status
prom-status:
	curl 'http://localhost:8428/api/v1/targets'|jq '.data.activeTargets| .[] | {pool:.scrapePool, status:.health}'

.PHONY: dc-app-logs-follow
dc-app-logs-follow:
	docker logs -f 673531d1b947

produce:
	curl -X 'POST' \
      'http://localhost:8001/api/v1/kafka/produce' \
      -H 'accept: application/json' \
      -H 'Content-Type: application/json' \
      -d '{"operation": {"operationType": "TYPE_DEBIT", "amount": 0, "accountId": "efc3f026-e496-47f4-9b03-bc937cc4ad73", "externalId": "efc3f026-e496-47f4-9b03-bc937cc4ad73"}}'

consume:
	curl -X 'POST' 'http://localhost:8001/api/v1/kafka/consume' -d '{"count": 1000}' | jq