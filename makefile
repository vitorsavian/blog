build-bahamut:
	@echo "Building Bahamut..."
	# using this as a workaround
	cd ./bahamut && go build -o ./dist/bahamut && cd ..