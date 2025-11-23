package main

import (
	"fmt"
	"log"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "go-study/examples/microservices/02-protobuf/complex/proto"
)

func main() {
	fmt.Println("=== Protocol Buffers 复杂消息示例 ===")

	// 创建复杂消息
	product := &pb.Product{
		Id:          1,
		Name:        "Go Programming Book",
		Description: "A comprehensive guide to Go programming",
		Price:       49.99,
		Stock:       100,
		Status:      pb.ProductStatus_PRODUCT_STATUS_ACTIVE,
		Supplier: &pb.Supplier{
			Id:   1,
			Name: "Tech Books Inc.",
			Address: &pb.Address{
				Street:  "123 Tech Street",
				City:    "San Francisco",
				Country: "USA",
				ZipCode: 94105,
			},
			Metadata: map[string]string{
				"contact": "support@techbooks.com",
				"rating":  "4.5",
			},
		},
		Tags: []string{"programming", "go", "book"},
		Warehouses: []*pb.Address{
			{
				Street:  "456 Warehouse Ave",
				City:    "New York",
				Country: "USA",
				ZipCode: 10001,
			},
		},
		Attributes: map[string]string{
			"isbn":  "978-0123456789",
			"pages": "500",
		},
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.Now(),
		Discount: &pb.Product_Discount{
			Percentage: 10.0, // 10% 折扣
		},
	}

	fmt.Printf("原始产品: %+v\n", product)

	// 序列化
	data, err := proto.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("序列化大小: %d bytes\n", len(data))

	// 反序列化
	var newProduct pb.Product
	if err := proto.Unmarshal(data, &newProduct); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n反序列化产品:\n")
	fmt.Printf("  ID: %d\n", newProduct.GetId())
	fmt.Printf("  名称: %s\n", newProduct.GetName())
	fmt.Printf("  价格: $%.2f\n", newProduct.GetPrice())
	fmt.Printf("  状态: %v\n", newProduct.GetStatus())
	fmt.Printf("  供应商: %s\n", newProduct.GetSupplier().GetName())
	fmt.Printf("  标签: %v\n", newProduct.GetTags())
	fmt.Printf("  折扣: %.1f%%\n", newProduct.GetDiscount().(type) == &pb.Product_Percentage{})

	// 演示 oneof
	if percentage, ok := newProduct.GetDiscount().(*pb.Product_Percentage); ok {
		fmt.Printf("  折扣类型: 百分比, 值: %.1f%%\n", percentage.Percentage)
	} else if fixed, ok := newProduct.GetDiscount().(*pb.Product_FixedAmount); ok {
		fmt.Printf("  折扣类型: 固定金额, 值: $%.2f\n", fixed.FixedAmount)
	}
}

