# Makefile
generate-swagger:
		swagger generate spec -o ./swagger.yml --scan-models

build:
		docker-compose build

run:
		docker-compose up