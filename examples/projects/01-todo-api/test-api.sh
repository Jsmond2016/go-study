#!/bin/bash

# TODO API 测试脚本
# 使用此脚本测试所有 API 端点

BASE_URL="http://localhost:8080"
FAILED_TESTS=0

echo "🚀 TODO API 测试开始..."
echo "========================================"

# 测试函数
test_api() {
    local method=$1
    local url=$2
    local data=$3
    local description=$4
    local expected_status=${5:-200}

    echo ""
    echo "📝 测试: $description"
    echo "请求: $method $url"

    if [ -n "$data" ]; then
        echo "数据: $data"
        response=$(curl -s -w "\n%{http_code}" -X $method \
                   -H "Content-Type: application/json" \
                   -d "$data" \
                   "$BASE_URL$url")
    else
        response=$(curl -s -w "\n%{http_code}" -X $method "$BASE_URL$url")
    fi

    # 获取状态码（最后一行）
    status_code=$(echo "$response" | tail -n1)
    # 获取响应体（除最后一行外的所有行）
    response_body=$(echo "$response" | sed '$d')

    echo "状态码: $status_code"
    echo "响应: $response_body"

    if [ "$status_code" = "$expected_status" ]; then
        echo "✅ 测试通过"
    else
        echo "❌ 测试失败 (期望状态码: $expected_status)"
        FAILED_TESTS=$((FAILED_TESTS + 1))
    fi

    echo "----------------------------------------"
}

# 等待服务器启动
echo "⏳ 等待服务器启动..."
sleep 2

# 1. 健康检查
test_api "GET" "/api/health" "" "健康检查"

# 2. 获取初始任务列表
test_api "GET" "/api/todos" "" "获取任务列表"

# 3. 创建任务
test_api "POST" "/api/todos" \
    '{"title":"学习Go语言","description":"完成Go语言基础教程学习","priority":"high","due_date":"2024-01-15"}' \
    "创建任务" "201"

# 4. 获取指定任务
test_api "GET" "/api/todos/1" "" "获取任务详情"

# 5. 更新任务
test_api "PUT" "/api/todos/1" \
    '{"title":"学习Go语言（更新）","description":"Go语言学习进度更新","status":"completed"}' \
    "更新任务"

# 6. 切换任务状态
test_api "PATCH" "/api/todos/1/toggle" "" "切换任务状态"

# 7. 创建第二个任务
test_api "POST" "/api/todos" \
    '{"title":"学习数据库","description":"学习SQL和NoSQL数据库","priority":"medium"}' \
    "创建第二个任务" "201"

# 8. 创建第三个任务
test_api "POST" "/api/todos" \
    '{"title":"学习Docker","description":"学习容器化部署","priority":"low","due_date":"2024-02-01"}' \
    "创建第三个任务" "201"

# 9. 搜索任务
test_api "GET" "/api/todos?search=Go" "" "搜索包含'Go'的任务"

# 10. 按状态过滤
test_api "GET" "/api/todos?status=completed" "" "过滤已完成任务"

# 11. 按优先级过滤
test_api "GET" "/api/todos?priority=high" "" "过滤高优先级任务"

# 12. 分页测试
test_api "GET" "/api/todos?page=1&page_size=2" "" "分页查询测试"

# 13. 获取统计信息
test_api "GET" "/api/todos/statistics" "" "获取统计信息"

# 14. 删除任务
test_api "DELETE" "/api/todos/3" "" "删除任务"

# 15. 获取API文档
test_api "GET" "/api/docs" "" "获取API文档"

# 16. 测试错误情况 - 无效的任务ID
test_api "GET" "/api/todos/999" "" "获取不存在的任务" "404"

# 17. 测试错误情况 - 无效的请求体
test_api "POST" "/api/todos" \
    '{"invalid":"data"}' \
    "无效的请求体" "400"

# 18. 测试错误情况 - 删除不存在的任务
test_api "DELETE" "/api/todos/999" "" "删除不存在的任务" "404"

# 测试结果汇总
echo ""
echo "========================================"
echo "🏁 测试完成"
echo "========================================"

if [ $FAILED_TESTS -eq 0 ]; then
    echo "🎉 所有测试通过！"
    exit 0
else
    echo "❌ 有 $FAILED_TESTS 个测试失败"
    exit 1
fi