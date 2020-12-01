.Phony: dependencies run build

dependencies:
	go get -u github.com/gin-gonic/gin
	go get k8s.io/client-go

build-darwin:
	rm -rf dist || true
	mkdir dist
	cd src; env GOOS=darwin GOARCH=amd64 go build -v -o wheelhouse
	mv src/wheelhouse dist/wheelhouse
	cp -r resources/templates dist/templates
	cp -r resources/data dist/data

build-linux:
	rm -rf dist || true
	mkdir dist
	cd src; env GOOS=linux GOARCH=amd64 go build -v -o wheelhouse
	mv src/wheelhouse dist/wheelhouse
	cp -r resources/templates dist/templates
	cp -r resources/data dist/data

build-windows:
	rm -rf dist || true
	mkdir dist
	cd src; env GOOS=windows GOARCH=amd64 go build -v -o wheelhouse
	mv src/wheelhouse dist/wheelhouse
	cp -r resources/templates dist/templates
	cp -r resources/data dist/data

run: 
	cd dist; ./wheelhouse