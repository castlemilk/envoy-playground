init:
	virtualenv venv

start:
	( \
		source venv/bin/activate; \
		FLASK_APP=flaskr FLASK_ENV=development flask run --port=8080; \
	)

build:
	docker-compose build
run: build
	docker-compose up