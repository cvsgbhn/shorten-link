
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

ARGS = $(filter-out $@,$(MAKECMDGOALS))

%:
	@:

include ./makefiles/*.mk