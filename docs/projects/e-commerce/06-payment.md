---
title: 支付集成
difficulty: advanced
duration: "4-5小时"
prerequisites: ["订单系统"]
tags: ["支付", "支付宝", "微信支付", "回调"]
---

# 支付集成

本章节将实现支付功能，包括支付宝和微信支付的集成、支付回调处理和支付状态管理。

## 📋 学习目标

完成本章节后，你将能够：

- [ ] 集成支付宝支付接口
- [ ] 集成微信支付接口
- [ ] 处理支付回调
- [ ] 实现支付状态管理
- [ ] 处理支付异常情况
- [ ] 实现退款功能

## 💳 支付服务接口

创建 `pkg/payment/payment.go`:

```go
package payment

import (
	"context"
)

type PaymentService interface {
	CreatePayment(orderID uint, amount float64, method string) (*PaymentRequest, error)
	VerifyCallback(data map[string]interface{}, method string) (*CallbackResult, error)
	QueryPayment(paymentNo string, method string) (*PaymentStatus, error)
	Refund(paymentNo string, amount float64, method string) error
}

type PaymentRequest struct {
	PaymentNo string
	PayURL    string
	QRCode    string
	Params    map[string]interface{}
}

type CallbackResult struct {
	PaymentNo     string
	TransactionID string
	Amount        float64
	Status        string
	Success       bool
}

type PaymentStatus struct {
	PaymentNo     string
	TransactionID string
	Status        string
	Amount        float64
}
```

## 🔵 支付宝集成

创建 `pkg/payment/alipay.go`:

```go
package payment

import (
	"encoding/json"
	"fmt"
	"github.com/smartwalle/alipay/v3"
)

type AlipayService struct {
	client *alipay.Client
	appID  string
}

func NewAlipayService(appID, privateKey, publicKey string) (PaymentService, error) {
	client, err := alipay.New(appID, privateKey, false)
	if err != nil {
		return nil, err
	}

	// 加载公钥
	if err = client.LoadAliPayPublicKey(publicKey); err != nil {
		return nil, err
	}

	return &AlipayService{
		client: client,
		appID:  appID,
	}, nil
}

func (s *AlipayService) CreatePayment(orderID uint, amount float64, method string) (*PaymentRequest, error) {
	var p alipay.TradePagePay
	p.NotifyURL = "https://your-domain.com/api/payment/alipay/notify"
	p.ReturnURL = "https://your-domain.com/api/payment/alipay/return"
	p.Subject = fmt.Sprintf("订单支付-%d", orderID)
	p.OutTradeNo = generatePaymentNo(orderID)
	p.TotalAmount = fmt.Sprintf("%.2f", amount)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := s.client.TradePagePay(p)
	if err != nil {
		return nil, err
	}

	return &PaymentRequest{
		PaymentNo: p.OutTradeNo,
		PayURL:    url.String(),
		Params:    make(map[string]interface{}),
	}, nil
}

func (s *AlipayService) VerifyCallback(data map[string]interface{}, method string) (*CallbackResult, error) {
	var notification alipay.Notification
	if err := json.Unmarshal([]byte(data["body"].(string)), &notification); err != nil {
		return nil, err
	}

	// 验证签名
	if err := s.client.VerifySign(notification); err != nil {
		return nil, err
	}

	result := &CallbackResult{
		PaymentNo:     notification.OutTradeNo,
		TransactionID: notification.TradeNo,
		Amount:        parseFloat(notification.TotalAmount),
		Status:        notification.TradeStatus,
		Success:       notification.TradeStatus == "TRADE_SUCCESS",
	}

	return result, nil
}
```

## 🟢 微信支付集成

创建 `pkg/payment/wechat.go`:

```go
package payment

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
)

type WechatPayService struct {
	client *core.Client
	appID  string
	mchID  string
}

func NewWechatPayService(appID, mchID, apiKey string) (PaymentService, error) {
	// 初始化微信支付客户端
	// 这里需要配置证书等
	client := &core.Client{} // 实际需要配置

	return &WechatPayService{
		client: client,
		appID:  appID,
		mchID:  mchID,
	}, nil
}

func (s *WechatPayService) CreatePayment(orderID uint, amount float64, method string) (*PaymentRequest, error) {
	svc := jsapi.JsapiApiService{Client: s.client}

	request := jsapi.PrepayRequest{
		Appid:       core.String(s.appID),
		Mchid:       core.String(s.mchID),
		Description: core.String(fmt.Sprintf("订单支付-%d", orderID)),
		OutTradeNo:  core.String(generatePaymentNo(orderID)),
		NotifyUrl:   core.String("https://your-domain.com/api/payment/wechat/notify"),
		Amount: &jsapi.Amount{
			Total: core.Int64(int64(amount * 100)), // 转换为分
		},
	}

	resp, err := svc.Prepay(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return &PaymentRequest{
		PaymentNo: *request.OutTradeNo,
		Params: map[string]interface{}{
			"prepay_id": resp.PrepayId,
		},
	}, nil
}
```

## 📝 支付处理器

创建 `internal/handler/payment.go`:

```go
package handler

import (
	"net/http"
	"strconv"
	"blog-system/internal/service"
	"blog-system/pkg/payment"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentService payment.PaymentService
	orderService   service.OrderService
}

func NewPaymentHandler(paymentService payment.PaymentService, orderService service.OrderService) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
		orderService:   orderService,
	}
}

func (h *PaymentHandler) Create(c *gin.Context) {
	orderID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	method := c.Query("method") // alipay or wechat

	order, err := h.orderService.GetByID(uint(orderID), c.GetUint("user_id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "订单不存在",
		})
		return
	}

	paymentReq, err := h.paymentService.CreatePayment(order.ID, order.TotalAmount, method)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "创建支付失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    paymentReq,
	})
}

func (h *PaymentHandler) AlipayCallback(c *gin.Context) {
	// 获取回调数据
	data := make(map[string]interface{})
	c.ShouldBind(&data)

	result, err := h.paymentService.VerifyCallback(data, "alipay")
	if err != nil {
		c.String(http.StatusBadRequest, "FAIL")
		return
	}

	// 更新订单状态
	if result.Success {
		h.orderService.UpdatePaymentStatus(result.PaymentNo, "paid")
	}

	c.String(http.StatusOK, "SUCCESS")
}

func (h *PaymentHandler) WechatCallback(c *gin.Context) {
	// 处理微信支付回调
	// 类似支付宝回调处理
}
```

## 🔧 路由设置

```go
func setupPaymentRoutes(r *gin.RouterGroup, paymentHandler *handler.PaymentHandler) {
	payment := r.Group("/payment")
	payment.Use(auth.AuthMiddleware())
	{
		payment.POST("/orders/:id", paymentHandler.Create)
		payment.GET("/orders/:id/status", paymentHandler.QueryStatus)
	}

	// 支付回调（不需要认证）
	callback := r.Group("/payment")
	{
		callback.POST("/alipay/notify", paymentHandler.AlipayCallback)
		callback.POST("/wechat/notify", paymentHandler.WechatCallback)
	}
}
```

## 📝 API 使用示例

### 创建支付

```bash
curl -X POST http://localhost:8080/api/payment/orders/1?method=alipay \
  -H "Authorization: Bearer <token>"
```

### 查询支付状态

```bash
curl http://localhost:8080/api/payment/orders/1/status \
  -H "Authorization: Bearer <token>"
```

## 💡 最佳实践

### 1. 支付安全

- 验证支付金额
- 验证订单状态
- 使用HTTPS
- 验证签名

### 2. 支付回调

- 幂等性处理
- 异步处理
- 重试机制
- 日志记录

### 3. 异常处理

- 支付超时
- 支付失败
- 重复支付
- 退款处理

## ⏭️ 下一步

支付集成完成后，下一步是：

- [库存管理](./07-inventory.md) - 实现库存管理功能

---

**🎉 支付集成完成！** 现在你可以开始实现库存管理了。

