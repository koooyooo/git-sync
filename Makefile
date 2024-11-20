.PHONY: install

install:
	@ go build -o git-sync main.go && sudo mv git-sync /usr/local/bin/git-sync
