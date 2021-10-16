SHELL := /bin/bash
MAKEFLAGS += --silent
ARGS = $(filter-out $@,$(MAKECMDGOALS))

.default: help

include .dev/*/*.mk
include .dev/*.mk
