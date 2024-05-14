package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// リクエストボディの構造体を定義
type RequestBody struct {
	Hello string `json:"hello"`
}

// レスポンスボディの構造体を定義
type ResponseBody struct {
	Message string `json:"message"`
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody RequestBody

	// リクエストボディをパース
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid request body",
		}, nil
	}

	// データを加工（例: "hello" の値を大文字にする）
	processedData := fmt.Sprintf("Hello, %s!", requestBody.Hello)

	// レスポンスボディを作成
	responseBody := ResponseBody{
		Message: processedData,
	}

	// レスポンスボディをJSONに変換
	responseBodyJSON, err := json.Marshal(responseBody)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(responseBodyJSON),
	}, nil
}

func main() {
	lambda.Start(handler)
}
