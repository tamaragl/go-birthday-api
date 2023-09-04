export PROJECT_NAME := go-birthday-api
export CURRENT_PATH := $(shell pwd)

define build_binaries
	@echo "- Building binaries..."
	@GOOS=linux GOARCH=amd64 go build -o bin/getHello cmd/lambda_get/main.go
	@GOOS=linux GOARCH=amd64 go build -o bin/addUser cmd/lambda_put/main.go
	@echo "Finished building binaries"
endef

define zip_files
	@echo "- Zipping files..."
	@for file in bin/*; do \
		zip -j $$file.zip $$file; \
		rm $$file; \
	done
	@echo "Finished zipping files"
endef

define clean_up
	@echo "- Cleaning up..."
	@rm -rf bin
endef

define deploy_to_aws
	@echo "- Deploying to AWS..."
	@serverless deploy --stage dev
	@echo "Finished deploying to AWS"
endef

deploy-aws:
	@rm -rf bin/
	${build_binaries}
	${zip_files}
	${deploy_to_aws}
	${clean_up}

dev:
	# TODO: move build to Dockerfile
	GOARCH=amd64 GOOS=linux go build -o myapp ./cmd/server_http
	docker-compose -p ${PROJECT_NAME} -f ${CURRENT_PATH}/ops/docker/docker-compose.yml up -d --build --remove-orphans
nodev:
	docker-compose -p ${PROJECT_NAME} -f ${CURRENT_PATH}/ops/docker/docker-compose.yml down --remove-orphans
	rm myapp
logs:
	docker-compose -p ${PROJECT_NAME} -f ${CURRENT_PATH}/ops/docker/docker-compose.yml logs -f
