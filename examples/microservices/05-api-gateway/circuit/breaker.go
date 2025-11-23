package circuit

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrCircuitOpen = errors.New("熔断器已打开")
)

// State 熔断器状态
type State int

const (
	StateClosed State = iota // 关闭：正常状态
	StateOpen                 // 打开：拒绝请求
	StateHalfOpen            // 半开：尝试恢复
)

// CircuitBreaker 熔断器
type CircuitBreaker struct {
	maxFailures  int           // 最大失败次数
	resetTimeout time.Duration // 重置超时
	state        State         // 当前状态
	failures     int           // 当前失败次数
	lastFailure  time.Time     // 最后失败时间
	mu           sync.RWMutex
}

// NewCircuitBreaker 创建熔断器
func NewCircuitBreaker(maxFailures int, resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		maxFailures:  maxFailures,
		resetTimeout: resetTimeout,
		state:        StateClosed,
	}
}

// Call 执行函数调用
func (cb *CircuitBreaker) Call(fn func() error) error {
	cb.mu.Lock()

	// 检查状态
	if cb.state == StateOpen {
		// 检查是否可以进入半开状态
		if time.Since(cb.lastFailure) >= cb.resetTimeout {
			cb.state = StateHalfOpen
			cb.failures = 0
		} else {
			cb.mu.Unlock()
			return ErrCircuitOpen
		}
	}

	cb.mu.Unlock()

	// 执行函数
	err := fn()

	cb.mu.Lock()
	defer cb.mu.Unlock()

	if err != nil {
		// 调用失败
		cb.failures++
		cb.lastFailure = time.Now()

		if cb.failures >= cb.maxFailures {
			cb.state = StateOpen
		}
		return err
	}

	// 调用成功
	if cb.state == StateHalfOpen {
		// 从半开状态恢复到关闭状态
		cb.state = StateClosed
		cb.failures = 0
	} else {
		// 重置失败计数
		cb.failures = 0
	}

	return nil
}

// GetState 获取当前状态
func (cb *CircuitBreaker) GetState() State {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.state
}

// Reset 手动重置熔断器
func (cb *CircuitBreaker) Reset() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.state = StateClosed
	cb.failures = 0
}

