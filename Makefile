TEST_PKGS := 

install:
	go build farm.e-pedion.com/repo/cache/client

pkg_cache:
	@echo "Add cache pkg for tests"
	$(eval TEST_PKGS += "farm.e-pedion.com/repo/cache")

pkg_test: pkg_cache
	@echo "TEST_PKGS=$(TEST_PKGS)"

test:
	@if [ "$(TEST_PKGS)" == "" ]; then \
	    echo "Building without TEST_PKGS" ;\
	    go test farm.e-pedion.com/repo/cache ;\
	else \
	    echo "Building with TEST_PKGS=$(TEST_PKGS)" ;\
	    go test $(TEST_PKGS) ;\
	fi

build: install
	@echo "Built successfully"