package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "token", "value1")
	bookHotel(ctx, "hotel")
}

func bookHotel(ctx context.Context, name string) {
	token := ctx.Value("token")
	fmt.Println(token)
}
