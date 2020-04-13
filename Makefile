SHELL := /bin/bash

##@ Generator
gen-api:  ## Generate the API package from the JSON specification
	$(eval input  ?= /)
	$(eval output ?= api)
	@printf "\033[2m→ Generating API package from specification ($(version):$(build_hash))...\033[0m\n"
	@{ \
		set -e; \
		trap "test -d .git && git checkout --quiet $(PWD)/internal/cmd/generate/go.mod" INT TERM EXIT; \
		export ELASTICSEARCH_BUILD_VERSION=$(version) && \
		export ELASTICSEARCH_BUILD_HASH=$(build_hash) && \
		cd internal/cmd/generate && \
		go run main.go apisource --input '$(PWD)/$(input)/rest-api-spec/api/*.json' --output '$(PWD)/$(output)' && \
		go run main.go apistruct --output '$(PWD)/$(output)'; \
	}