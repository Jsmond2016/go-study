# Go 学习项目构建脚本

.PHONY: help fmt lint test build clean docs-serve docs-build install-tools check-modules

# 默认目标
help:
	@echo "Go 学习项目构建脚本"
	@echo ""
	@echo "可用命令:"
	@echo "  fmt         - 格式化所有代码"
	@echo "  lint        - 代码质量检查"
	@echo "  test        - 运行所有测试"
	@echo "  build       - 构建所有示例"
	@echo "  clean       - 清理构建文件"
	@echo "  docs-serve  - 启动文档开发服务器"
	@echo "  docs-build  - 构建文档"
	@echo "  install-tools - 安装开发工具"
	@echo "  check-modules - 检查模块依赖"

# 格式化代码
fmt:
	@echo "格式化代码..."
	go work sync
	go fmt ./...
	@echo "代码格式化完成"

# 代码质量检查
lint:
	@echo "执行代码质量检查..."
	@if ! command -v golangci-lint &> /dev/null; then \
		echo "golangci-lint 未安装，请运行 'make install-tools'"; \
		exit 1; \
	fi
	golangci-lint run

# 运行测试
test:
	@echo "运行所有测试..."
	go work sync
	go test -v -race ./...

# 构建所有示例
build:
	@echo "构建所有示例..."
	go work sync
	@for dir in examples/*/; do \
		if [ -f "$$dir/go.mod" ]; then \
			echo "构建 $$dir"; \
			cd "$$dir" && go build -o ../../bin/$$(basename $$dir) . && cd ../..; \
		fi \
	done
	@echo "构建完成"

# 清理构建文件
clean:
	@echo "清理构建文件..."
	rm -rf bin/
	@echo "清理完成"

# 启动文档开发服务器
docs-serve:
	@echo "启动文档开发服务器..."
	@if ! command -v pnpm &> /dev/null; then \
		echo "pnpm 未安装，请先安装 pnpm"; \
		exit 1; \
	fi
	pnpm dev

# 构建文档
docs-build:
	@echo "构建文档..."
	@if ! command -v pnpm &> /dev/null; then \
		echo "pnpm 未安装，请先安装 pnpm"; \
		exit 1; \
	fi
	pnpm build

# 安装开发工具
install-tools:
	@echo "安装开发工具..."
	@echo "安装 golangci-lint..."
	@if ! command -v golangci-lint &> /dev/null; then \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	else \
		echo "golangci-lint 已安装"; \
	fi
	@echo "安装 goimports..."
	@if ! command -v goimports &> /dev/null; then \
		go install golang.org/x/tools/cmd/goimports@latest; \
	else \
		echo "goimports 已安装"; \
	fi
	@echo "开发工具安装完成"

# 检查模块依赖
check-modules:
	@echo "检查模块依赖..."
	go work sync
	@for dir in examples/*/; do \
		if [ -f "$$dir/go.mod" ]; then \
			echo "检查 $$dir 依赖..."; \
			cd "$$dir" && go mod tidy && go mod verify && cd ../..; \
		fi \
	done
	@echo "模块依赖检查完成"

# 初始化项目
init: install-tools check-modules fmt
	@echo "项目初始化完成"

# 快速检查（格式化 + 代码检查）
check: fmt lint
	@echo "快速检查完成"

# 完整检查（格式化 + 代码检查 + 测试）
full-check: check test
	@echo "完整检查完成"
