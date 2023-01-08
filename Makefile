build:
	@go build
	@sudo mv 403_bypass /usr/bin/
	@echo -e "\e[0;32m[+]\e[0m 403_bypass has been added to /usr/bin."

uninstall:
	@sudo rm /usr/bin/403_bypass
	@echo -e "\e[0;31m[-]\e[0m Removed /usr/bin/403_bypass."