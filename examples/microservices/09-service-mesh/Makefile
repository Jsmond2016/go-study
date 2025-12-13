.PHONY: help build run docker-build docker-run k8s-deploy k8s-delete istio-apply istio-delete clean

# 默认目标
help:
	@echo "Service Mesh with Istio - Makefile"
	@echo ""
	@echo "可用命令:"
	@echo "  build         - 构建 Go 应用"
	@echo "  run           - 本地运行应用"
	@echo "  docker-build  - 构建 Docker 镜像（v1 和 v2）"
	@echo "  docker-run    - 本地运行 Docker 容器"
	@echo "  k8s-deploy    - 部署到 Kubernetes"
	@echo "  k8s-delete    - 删除 Kubernetes 部署"
	@echo "  istio-apply   - 应用 Istio 配置"
	@echo "  istio-delete  - 删除 Istio 配置"
	@echo "  test-traffic  - 测试流量路由"
	@echo "  clean         - 清理构建文件"

# 构建 Go 应用
build:
	@echo "构建 Go 应用..."
	go build -o bin/service main.go
	@echo "构建完成: bin/service"

# 本地运行
run:
	@echo "启动服务（默认 v1）..."
	SERVICE_VERSION=v1 PORT=8080 go run main.go

# 构建 Docker 镜像
docker-build:
	@echo "构建 Docker 镜像 v1..."
	docker build -t demo-service:v1 --build-arg VERSION=v1 .
	@echo "构建 Docker 镜像 v2..."
	docker build -t demo-service:v2 --build-arg VERSION=v2 .
	@echo "Docker 镜像构建完成"

# 本地运行 Docker 容器
docker-run:
	@echo "运行 Docker 容器..."
	docker run -d --name demo-service-v1 -p 8080:8080 \
		-e SERVICE_VERSION=v1 \
		demo-service:v1

# 部署到 Kubernetes
k8s-deploy:
	@echo "部署到 Kubernetes..."
	kubectl apply -f k8s-deployment.yaml
	@echo "等待 Pod 就绪..."
	kubectl wait --for=condition=ready pod -l app=demo-service --timeout=60s
	kubectl get pods -l app=demo-service

# 删除 Kubernetes 部署
k8s-delete:
	@echo "删除 Kubernetes 部署..."
	kubectl delete -f k8s-deployment.yaml

# 应用 Istio 配置
istio-apply:
	@echo "应用 Istio 流量管理配置..."
	kubectl apply -f istio-traffic.yaml
	@echo "应用 Istio 安全策略..."
	kubectl apply -f istio-security.yaml
	@echo "Istio 配置应用完成"

# 删除 Istio 配置
istio-delete:
	@echo "删除 Istio 配置..."
	kubectl delete -f istio-traffic.yaml
	kubectl delete -f istio-security.yaml

# 测试流量路由
test-traffic:
	@echo "测试流量路由..."
	@echo "获取 Ingress Gateway 地址..."
	@export INGRESS_HOST=$$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}'); \
	export INGRESS_PORT=$$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].port}'); \
	export GATEWAY_URL=$$INGRESS_HOST:$$INGRESS_PORT; \
	echo "Gateway URL: http://$$GATEWAY_URL"; \
	echo ""; \
	echo "测试 10 次请求，观察版本分布:"; \
	for i in {1..10}; do \
		curl -s http://$$GATEWAY_URL/api/feature | jq -r '.version'; \
	done

# 查看服务状态
status:
	@echo "服务状态:"
	@kubectl get pods -l app=demo-service
	@echo ""
	@echo "Istio VirtualService:"
	@kubectl get virtualservice
	@echo ""
	@echo "Istio DestinationRule:"
	@kubectl get destinationrule

# 查看日志
logs:
	@echo "查看服务日志（最新的 Pod）..."
	@kubectl logs -l app=demo-service -c demo-service --tail=50 -f

# 打开 Kiali 仪表盘
kiali:
	@echo "打开 Kiali 仪表盘..."
	istioctl dashboard kiali

# 打开 Jaeger 仪表盘
jaeger:
	@echo "打开 Jaeger 仪表盘..."
	istioctl dashboard jaeger

# 打开 Grafana 仪表盘
grafana:
	@echo "打开 Grafana 仪表盘..."
	istioctl dashboard grafana

# 清理
clean:
	@echo "清理构建文件..."
	rm -rf bin/
	docker rm -f demo-service-v1 2>/dev/null || true
	@echo "清理完成"

# 完整部署（构建 + 部署 + Istio）
deploy-all: docker-build k8s-deploy istio-apply
	@echo "完整部署完成！"
	@echo "运行 'make test-traffic' 测试流量路由"
	@echo "运行 'make kiali' 查看服务网格可视化"

# 完整清理
clean-all: istio-delete k8s-delete clean
	@echo "完整清理完成！"
