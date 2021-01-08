install-deps:
	@echo " > \033[32mInstalling dependencies... \033[0m "
	./scripts/install_deps.sh

compile:
	@echo " > \033[32mCompiling contracts... \033[0m "
	npx truffle compile

bindings: compile
	@echo " > \033[32mCreating go bindings for ethereum contracts... \033[0m "
	./scripts/create_bindings.sh
