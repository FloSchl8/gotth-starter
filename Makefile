run: build
	@./bin/app

build:
	@go build -tags dev -o bin/app .

templ-watch:
	@templ generate --watch --proxy=http://localhost:8080
	
css:
	@tailwindcss -i package/views/css/app.css -o public/styles.css --watch
