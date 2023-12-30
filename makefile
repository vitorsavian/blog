build-bahamut:
	@echo "Building Bahamut..."
	cd ./bahamut && go build -o ./dist/bahamut && cd ..

bahamut: build-bahamut
	@echo "Running Bahamut..."
	./bahamut/dist/bahamut http