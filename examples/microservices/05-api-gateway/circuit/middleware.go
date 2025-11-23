package circuit

import (
	"errors"
	"net/http"
	"sync"
	"time"
)

// CircuitBreakerMiddleware 熔断器中间件
func CircuitBreakerMiddleware(maxFailures int, resetTimeout time.Duration) func(http.Handler) http.Handler {
	breakers := make(map[string]*CircuitBreaker)
	mu := sync.RWMutex{}

	getBreaker := func(key string) *CircuitBreaker {
		mu.RLock()
		breaker, exists := breakers[key]
		mu.RUnlock()

		if !exists {
			mu.Lock()
			breaker = NewCircuitBreaker(maxFailures, resetTimeout)
			breakers[key] = breaker
			mu.Unlock()
		}

		return breaker
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 根据路径获取熔断器
			breaker := getBreaker(r.URL.Path)

			// 检查熔断器状态
			if breaker.GetState() == StateOpen {
				http.Error(w, "服务暂时不可用", http.StatusServiceUnavailable)
				return
			}

			// 执行请求
			err := breaker.Call(func() error {
				// 创建响应写入器以捕获错误
				wrapped := &responseWriter{
					ResponseWriter: w,
					statusCode:     http.StatusOK,
				}

				next.ServeHTTP(wrapped, r)

				// 如果状态码表示错误，返回错误
				if wrapped.statusCode >= http.StatusInternalServerError {
					return errors.New("服务错误")
				}

				return nil
			})

			if err == ErrCircuitOpen {
				http.Error(w, "服务暂时不可用", http.StatusServiceUnavailable)
			}
		})
	}
}

// responseWriter 包装 ResponseWriter
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

