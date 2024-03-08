

test-verbose: create-temp ## Run verbose all tests
	go test -v -count=1 -race -coverprofile=tmp/coverage-report-unit.cov -covermode=atomic ./... 

test-unit: create-temp ## run tests	
	go test -count=1 -race -coverprofile=tmp/coverage-report-unit.cov -covermode=atomic --tags=unit ./...

test-unit-verbose: create-temp ## Run unit tests
	go test -v -count=1 -race -coverprofile=tmp/coverage-report-unit.cov -covermode=atomic --tags=unit ./...
	
test: create-temp ## create temp folder 
	go test -count=1 -race -coverprofile=tmp/coverage-report.cov -covermode=atomic  ./...

test-clear-cache: ## clear test cache 
	go clean -testcache

test-e2e: ## run e2e tests
	@echo APPLICATION_URL=${APPLICATION_URL}
	@echo NEW_VERSION=${NEW_VERSION}
	@echo TEST_TIMEOUT=${TEST_TIMEOUT}
	go test --tags=e2e -v ./...

test-e2e-clear-cache: ## clear test cache 
	go clean -testcache

test-e2e-local: test-e2e-clear-cache ## run e2e tests local
	export
	APPLICATION_URL=http://localhost:8080 \
	APPLICATION_VERSION= \
	TEST_TIMEOUT=0 \
	go test --tags=e2e -v ./...


test-race:
	echo START > tmp/test-race.txt
	@for i in `seq 1 100`; \
		do echo [ $$i ]============================================ >> tmp/test-race.txt && \
			make test --silent >> tmp/test-race.txt; \
	done;
	echo "------------------------------" >> tmp/test-race.txt
	