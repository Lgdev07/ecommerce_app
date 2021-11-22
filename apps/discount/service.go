package discount

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type DiscountService struct{}

type DiscountServiceInt interface {
	GetProductDiscount(productID int) (float64, error)
}

func connect() *grpc.ClientConn {
	host := os.Getenv("DISCOUNT_SERVICE_HOST")
	port := os.Getenv("DISCOUNT_SERVICE_PORT")
	timeoutEnv, _ := strconv.Atoi(os.Getenv("DISCOUNT_SERVICE_TIMEOUT"))

	address := fmt.Sprintf("%s:%s", host, port)

	timeout := time.Duration(timeoutEnv) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	connection, err := grpc.DialContext(ctx, address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}

	return connection
}

func (grpcClient *DiscountService) GetProductDiscount(productID int) (float64, error) {
	connection := connect()

	defer connection.Close()

	client := NewDiscountClient(connection)

	request := &GetDiscountRequest{
		ProductID: int32(productID),
	}

	response, err := client.GetDiscount(context.Background(), request)
	if err != nil {
		return 0, err
	}

	return float64(response.Percentage), nil
}
