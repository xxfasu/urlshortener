.PHONY: wire
wire:
	wire ./cmd/main/wire

.PHONY: convey
convey:
	goconvey -port 5555

.PHONY: mock

mock:
		# 查找所有 interface.go 文件
		REPOSITORY_INTERFACE_FILES := $(shell find internal/repository -type f -name "interface.go")
		FRONT_SERVICE_INTERFACE_FILES := $(shell find internal/service -type f -name "interface.go")
	@echo "Start generating mock files..."

	@for repository in $(REPOSITORY_INTERFACE_FILES); do \
		SERVICE_NAME=$$(basename $$(dirname $$repository)); \
		DEST_DIR=test/mocks/repository/mocks_$$SERVICE_NAME; \
		MOCK_FILE=$$DEST_DIR/mocks_$$SERVICE_NAME.go; \
		PACKAGE_NAME=mocks_$$SERVICE_NAME; \
		echo "Generate mock files for $$repository ..."; \
		mkdir -p $$DEST_DIR; \
		mockgen -source=$$repository -destination=$$MOCK_FILE -package=$$PACKAGE_NAME -exclude_interfaces=Reader,Writer ; \
		echo "The resulting mock file is located in $$MOCK_FILE"; \
	done

	@for admin_service in $(ADMIN_SERVICE_INTERFACE_FILES); do \
		SERVICE_NAME=$$(basename $$(dirname $$admin_service)); \
		DEST_DIR=test/mocks/service/admin_service/mocks_$$SERVICE_NAME; \
		MOCK_FILE=$$DEST_DIR/mocks_$$SERVICE_NAME.go; \
		PACKAGE_NAME=mocks_$$SERVICE_NAME; \
		echo "Generate mock files for $$admin_service ..."; \
		mkdir -p $$DEST_DIR; \
		mockgen -source=$$admin_service -destination=$$MOCK_FILE -package=$$PACKAGE_NAME; \
		echo "The resulting mock file is located in $$MOCK_FILE"; \
	done

	@for common_service in $(COMMON_SERVICE_INTERFACE_FILES); do \
  		SERVICE_NAME=$$(basename $$(dirname $$common_service)); \
  		DEST_DIR=test/mocks/service/common_service/mocks_$$SERVICE_NAME; \
  		MOCK_FILE=$$DEST_DIR/mocks_$$SERVICE_NAME.go; \
  		PACKAGE_NAME=mocks_$$SERVICE_NAME; \
  		echo "Generate mock files for $$common_service ..."; \
  		mkdir -p $$DEST_DIR; \
  		mockgen -source=$$common_service -destination=$$MOCK_FILE -package=$$PACKAGE_NAME; \
  		echo "The resulting mock file is located in $$MOCK_FILE"; \
  	done

	mockgen -source=internal/repository/db.go -destination=test/mocks/repository/mocks_transaction/mocks_transaction.go -package=mocks_transaction;
	@echo "mock file generation is complete!"
