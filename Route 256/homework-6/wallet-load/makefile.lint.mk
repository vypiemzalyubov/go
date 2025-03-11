LINTER_CONFIG:=.golangci.pipeline.yaml
GOLANGCI_BIN := $(LOCAL_BIN)/golangci-lint
GOLANGCI_TAG ?= 1.57.2

.PHONY: .install-lint
.install-lint: export GOBIN := $(LOCAL_BIN)
.install-lint: ## Установить golangci-lint в текущую директорию с исполняемыми файлами
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info Downloading golangci-lint)
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_TAG)
endif

.PHONY: lint
lint: .install-lint
	$(GOLANGCI_BIN) run \
		--new-from-rev=origin/master \
		--config=${LINTER_CONFIG} \
		--sort-results \
		--max-issues-per-linter=1000 \
		--max-same-issues=1000 \
		./...

.PHONY: lint-full
lint-full: .install-lint
	$(GOLANGCI_BIN) run \
		--config=${LINTER_CONFIG} \
		--sort-results \
		--max-issues-per-linter=1000 \
		--max-same-issues=1000 \
		./...