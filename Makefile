#! /usr/bin/make
#
# Makefile for goa cellar example pimped out govendor
#
# Targets:
# - clean     delete all generated files
# - generate  (re)generate all goagen-generated files.
# - build     compile executable
#
# Meta targets:
# - all is the default target, it runs all the targets in the order above.
#
CURRENT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
design_path := "$(CURRENT_DIR)/design"



.PHONY: swagger
swagger:
	@echo "generating $(design_path)..."
	@go run gen/main.go swagger

