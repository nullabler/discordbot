SHELL := /bin/bash
MAKEFLAGS += --silent
ARGS = $(filter-out $@,$(MAKECMDGOALS))

.default: help

include scripts/*/*.mk
include scripts/*.mk
