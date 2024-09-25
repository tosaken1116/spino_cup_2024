# Makefile

# mockgen コマンドのパス
MOCKGEN := mockgen

# プロジェクトのルートパッケージ名
PACKAGE_NAME := your_project_package_name

# モック生成対象の定義
# フォーマット: source_file:InterfaceName:destination_file
MOCKS := \
	internal/usecase/room_usecase.go:RoomUsecase:internal/mock/usecase/mock_room_usecase.go \
	internal/usecase/in_room_usecase.go:InRoomUsecase:internal/mock/usecase/mock_in_room_usecase.go \
	internal/domain/repository/room_repo.go:RoomRepository:internal/mock/domain/repository/mock_room_repository.go \
	internal/domain/repository/active_room_repo.go:ActiveRoomRepository:internal/mock/domain/repository/mock_active_room_repository.go

# デフォルトターゲット
.PHONY: all
all: generate-mocks

# モックを一括生成するターゲット
.PHONY: generate-mocks
generate-mocks:
	@echo "モックファイルを生成しています..."
	@mkdir -p internal/mock/usecase internal/mock/domain/repository
	@for mock in $(MOCKS); do \
		src=$$(echo $$mock | cut -d':' -f1); \
		iface=$$(echo $$mock | cut -d':' -f2); \
		out=$$(echo $$mock | cut -d':' -f3); \
		package_name=$$(basename $$(dirname $$out)); \
		echo "インターフェース $$iface のモックを生成中: $$out"; \
		$(MOCKGEN) -source=$$src -destination=$$out -package=$$package_name $$PACKAGE_NAME/$$src $$iface; \
	done
	@echo "モックファイルの生成が完了しました。"

# テスト実行前にモックを生成するターゲット
.PHONY: test
test: generate-mocks
	go test ./...

# クリーンアップターゲット（生成されたモックファイルを削除）
.PHONY: clean-mocks
clean-mocks:
	@echo "生成されたモックファイルを削除しています..."
	@rm -f internal/mock/usecase/mock_*.go
	@rm -f internal/mock/domain/repository/mock_*.go
	@echo "モックファイルの削除が完了しました。"

# その他のターゲット例（必要に応じて追加可能）
.PHONY: build
build: generate-mocks
	go build -o bin/app ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	# ここに linter コマンドを追加（例: golangci-lint）
	golangci-lint run
