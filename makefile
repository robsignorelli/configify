PACKAGE=github.com/robsignorelli/configify

TESTING_FLAGS=
ifeq ($(VERBOSE),true)
	TESTING_FLAGS=-v
endif

#
# Runs through our suite of all unit tests
#
test:
	go test $(TESTING_FLAGS) -timeout 5s $(PACKAGE)/...
#
# Runs through our suite of all unit tests
#
coverage:
	go test $(TESTING_FLAGS) -cover -timeout 5s $(PACKAGE)/...

#
# Runs our benchmarks
#
benchmark:
	go test -bench=.