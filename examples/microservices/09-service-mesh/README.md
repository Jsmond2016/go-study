# Service Mesh with Istio - Go 应用示例

本示例展示如何将 Go 应用集成到 Istio 服务网格中，并利用 Istio 的流量管理、安全和可观测性功能。

## 目录结构

```
09-service-mesh/
├── main.go                  # Go 应用程序（支持 Istio）
├── Dockerfile               # Docker 镜像构建文件
├── k8s-deployment.yaml      # Kubernetes 部署配置
├── istio-traffic.yaml       # Istio 流量管理配置
├── istio-security.yaml      # Istio 安全策略配置
├── Makefile                 # 构建和部署脚本
└── README.md                # 本文件
```

## 功能特性

### 1. Go 应用集成 Istio

应用实现了以下 Istio 集成特性：

- ✅ **追踪 Header 传播**: 自动传播 Istio 追踪 header（x-request-id, x-b3-* 等）
- ✅ **Prometheus 指标**: 暴露 `/metrics` 端点供 Istio 采集
- ✅ **健康检查**: `/health` 和 `/ready` 端点供 Kubernetes 探针使用
- ✅ **版本标识**: 支持多版本部署（v1, v2, v3）
- ✅ **环境变量配置**: 从环境变量读取配置，适配容器化部署

### 2. Istio 流量管理

#### A/B 测试

通过权重分配实现流量分流：
- 80% 流量 → v1 版本
- 20% 流量 → v2 版本

#### 金丝雀发布

基于 HTTP Header 的路由：
- 带 `x-user-type: premium` header → v2 版本
- 其他用户 → 按权重分配

#### 熔断机制

配置了异常检测（Circuit Breaking）：
- 连续错误 5 次触发熔断
- 熔断时间 30 秒
- 最大隔离比例 50%

### 3. Istio 安全策略

- **mTLS**: 强制服务间通信加密
- **授权策略**: 基于服务身份的访问控制
- **JWT 认证**: 支持 JWT token 验证

## 快速开始

### 前置条件

1. Kubernetes 集群（Minikube、Kind 或云服务）
2. 安装 kubectl
3. 安装 Istio

#### 安装 Istio

```bash
# 下载 Istio
curl -L https://istio.io/downloadIstio | sh -
cd istio-*

# 安装 Istio（demo 配置包含所有组件）
istioctl install --set profile=demo -y

# 启用 default 命名空间的自动 sidecar 注入
kubectl label namespace default istio-injection=enabled
```

### 构建和部署

#### 1. 构建 Docker 镜像

```bash
# 构建 v1 版本
docker build -t demo-service:v1 --build-arg SERVICE_VERSION=v1 .

# 构建 v2 版本
docker build -t demo-service:v2 --build-arg SERVICE_VERSION=v2 .
```

#### 2. 部署到 Kubernetes

```bash
# 部署应用（v1 和 v2 版本）
kubectl apply -f k8s-deployment.yaml

# 查看 Pod 状态（应该看到 2/2 READY，说明 Istio sidecar 已注入）
kubectl get pods

# 示例输出：
# NAME                               READY   STATUS    RESTARTS   AGE
# demo-service-v1-xxxxxxxxxx-xxxxx   2/2     Running   0          1m
# demo-service-v2-xxxxxxxxxx-xxxxx   2/2     Running   0          1m
```

#### 3. 应用 Istio 配置

```bash
# 应用流量管理规则
kubectl apply -f istio-traffic.yaml

# 应用安全策略
kubectl apply -f istio-security.yaml
```

#### 4. 访问应用

```bash
# 获取 Ingress Gateway 地址
export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].port}')
export GATEWAY_URL=$INGRESS_HOST:$INGRESS_PORT

# 访问应用
curl http://$GATEWAY_URL/api/hello

# 查看版本信息
curl http://$GATEWAY_URL/api/version

# 测试特性端点（观察不同版本的响应）
curl http://$GATEWAY_URL/api/feature
```

## 测试场景

### 场景 1: A/B 测试

多次调用同一接口，观察不同版本的响应（80% v1, 20% v2）：

```bash
for i in {1..10}; do
  curl -s http://$GATEWAY_URL/api/feature | jq .version
done
```

### 场景 2: 金丝雀发布

带特定 header 的请求会路由到 v2：

```bash
# Premium 用户 → v2
curl -H "x-user-type: premium" http://$GATEWAY_URL/api/feature | jq .

# 普通用户 → 按权重分配
curl http://$GATEWAY_URL/api/feature | jq .
```

### 场景 3: 熔断测试

```bash
# 使用压测工具触发熔断
kubectl run fortio --image=fortio/fortio --image-pull-policy=IfNotPresent -i -t -- \
  load -c 3 -qps 0 -n 30 -loglevel Warning http://demo-service:8080/api/hello
```

### 场景 4: 查看追踪

访问 Jaeger UI 查看分布式追踪：

```bash
# 端口转发 Jaeger
istioctl dashboard jaeger
```

### 场景 5: 查看指标

访问 Grafana 查看服务指标：

```bash
# 端口转发 Grafana
istioctl dashboard grafana
```

### 场景 6: 查看服务网格拓扑

访问 Kiali 查看服务拓扑和流量：

```bash
# 端口转发 Kiali
istioctl dashboard kiali
```

## API 端点说明

| 端点 | 方法 | 说明 |
|------|------|------|
| `/health` | GET | 健康检查（Liveness Probe） |
| `/ready` | GET | 就绪检查（Readiness Probe） |
| `/metrics` | GET | Prometheus 指标 |
| `/api/hello` | GET | 简单问候接口 |
| `/api/version` | GET | 返回服务版本和 Pod 信息 |
| `/api/feature` | GET | 不同版本有不同特性（演示版本差异） |
| `/api/call-service` | GET | 调用下游服务（演示服务间调用） |

## Istio 配置说明

### VirtualService（虚拟服务）

定义流量路由规则：

```yaml
apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: demo-service
spec:
  hosts:
  - demo-service
  http:
  # 基于 header 的路由
  - match:
    - headers:
        x-user-type:
          exact: premium
    route:
    - destination:
        host: demo-service
        subset: v2
  # 基于权重的路由
  - route:
    - destination:
        host: demo-service
        subset: v1
      weight: 80
    - destination:
        host: demo-service
        subset: v2
      weight: 20
```

### DestinationRule（目标规则）

定义服务子集和流量策略：

```yaml
apiVersion: networking.istio.io/v1beta1
kind: DestinationRule
metadata:
  name: demo-service
spec:
  host: demo-service
  trafficPolicy:
    loadBalancer:
      simple: LEAST_REQUEST
    outlierDetection:
      consecutiveErrors: 5
      interval: 30s
  subsets:
  - name: v1
    labels:
      version: v1
  - name: v2
    labels:
      version: v2
```

### PeerAuthentication（对等认证）

启用 mTLS：

```yaml
apiVersion: security.istio.io/v1beta1
kind: PeerAuthentication
metadata:
  name: default
spec:
  mtls:
    mode: STRICT
```

## Go 应用关键代码

### 1. 追踪 Header 传播

Istio 需要应用程序手动传播追踪 header：

```go
func istioHeadersMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tracingHeaders := []string{
            "x-request-id",
            "x-b3-traceid",
            "x-b3-spanid",
            "x-b3-parentspanid",
            "x-b3-sampled",
            "x-b3-flags",
        }

        headers := make(map[string]string)
        for _, h := range tracingHeaders {
            if value := c.GetHeader(h); value != "" {
                headers[h] = value
            }
        }

        c.Set("tracing-headers", headers)
        c.Next()
    }
}
```

### 2. 服务间调用

调用下游服务时传播 header：

```go
func callServiceHandler(c *gin.Context) {
    client := &http.Client{Timeout: 5 * time.Second}
    req, _ := http.NewRequest("GET", targetService+"/api/hello", nil)

    // 传播追踪 header
    if headers, exists := c.Get("tracing-headers"); exists {
        for k, v := range headers.(map[string]string) {
            req.Header.Set(k, v)
        }
    }

    resp, err := client.Do(req)
    // ...
}
```

### 3. Prometheus 指标

暴露自定义指标：

```go
var httpRequestsTotal = prometheus.NewCounterVec(
    prometheus.CounterOpts{
        Name: "http_requests_total",
        Help: "Total number of HTTP requests",
    },
    []string{"method", "endpoint", "status", "version"},
)

// 在中间件中记录
func prometheusMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        c.Next()

        duration := time.Since(start).Seconds()
        httpRequestsTotal.WithLabelValues(
            c.Request.Method,
            c.FullPath(),
            fmt.Sprintf("%d", c.Writer.Status()),
            serviceVersion,
        ).Inc()
    }
}
```

## 监控和可观测性

### Kiali（服务网格可视化）

```bash
istioctl dashboard kiali
```

功能：
- 服务拓扑图
- 实时流量
- 配置验证
- 健康检查

### Jaeger（分布式追踪）

```bash
istioctl dashboard jaeger
```

功能：
- 请求追踪
- 延迟分析
- 依赖关系图
- 错误追踪

### Grafana（指标可视化）

```bash
istioctl dashboard grafana
```

预置仪表盘：
- Istio Service Dashboard
- Istio Workload Dashboard
- Istio Mesh Dashboard

### Prometheus（指标存储）

```bash
istioctl dashboard prometheus
```

## 最佳实践

### 1. 应用程序设计

- ✅ 传播追踪 header（必须）
- ✅ 实现健康检查端点
- ✅ 暴露 Prometheus 指标
- ✅ 使用环境变量配置
- ✅ 优雅关闭（处理 SIGTERM）

### 2. 流量管理

- ✅ 从小比例开始金丝雀发布（5% → 10% → 50% → 100%）
- ✅ 设置合理的超时时间
- ✅ 配置重试策略
- ✅ 启用熔断保护

### 3. 安全

- ✅ 启用 mTLS
- ✅ 最小权限原则
- ✅ 定期更新证书
- ✅ 审计日志

### 4. 性能优化

- ✅ 调整连接池大小
- ✅ 优化 sidecar 资源限制
- ✅ 使用 HTTP/2
- ✅ 启用压缩

## 故障排查

### Pod 无法启动

```bash
# 查看 Pod 详情
kubectl describe pod <pod-name>

# 查看 sidecar 日志
kubectl logs <pod-name> -c istio-proxy

# 查看应用日志
kubectl logs <pod-name> -c demo-service
```

### 流量路由不生效

```bash
# 验证 VirtualService 配置
istioctl analyze

# 查看 Envoy 配置
istioctl proxy-config routes <pod-name>

# 查看路由详情
istioctl proxy-config routes <pod-name> --name <route-name> -o json
```

### mTLS 连接失败

```bash
# 检查 mTLS 状态
istioctl authn tls-check <pod-name> demo-service.default.svc.cluster.local

# 查看证书
istioctl proxy-config secret <pod-name>
```

## 清理资源

```bash
# 删除 Istio 配置
kubectl delete -f istio-security.yaml
kubectl delete -f istio-traffic.yaml

# 删除应用
kubectl delete -f k8s-deployment.yaml

# 卸载 Istio
istioctl uninstall --purge -y

# 删除命名空间标签
kubectl label namespace default istio-injection-
```

## 学习资源

- [Istio 官方文档](https://istio.io/latest/docs/)
- [Istio 最佳实践](https://istio.io/latest/docs/ops/best-practices/)
- [Envoy 代理文档](https://www.envoyproxy.io/docs/envoy/latest/)
- [服务网格书籍](https://www.servicemeshbook.com/)

## 进阶主题

1. **多集群服务网格**: 跨多个 Kubernetes 集群的服务网格
2. **虚拟机集成**: 将虚拟机工作负载纳入服务网格
3. **外部服务**: 管理对集群外部服务的访问
4. **Wasm 插件**: 使用 WebAssembly 扩展 Envoy
5. **GitOps 集成**: 使用 ArgoCD 管理 Istio 配置

## 常见问题

**Q: Istio 会增加多少延迟？**
A: 通常增加 1-3ms，取决于配置和网络环境。

**Q: 是否所有应用都需要修改代码？**
A: 基本功能不需要，但建议传播追踪 header 以获得完整的分布式追踪。

**Q: 如何监控 Istio 性能？**
A: 使用 Grafana 仪表盘查看 sidecar CPU/内存使用率和请求延迟。

**Q: 可以选择性地为某些 Pod 注入 sidecar 吗？**
A: 可以，使用 Pod annotation `sidecar.istio.io/inject: "false"` 排除特定 Pod。
