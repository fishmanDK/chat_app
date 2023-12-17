package grpc

import (
	"context"
	"fmt"
	fishman_auth_v1 "github.com/fishmanDK/proto-auth/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	api fishman_auth_v1.AuthClient
}

func NewClient(ctx context.Context, addr string) (*Client, error) {
	const op = "clients.auth.grpc.NewClient"
	cc, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{
		api: fishman_auth_v1.NewAuthClient(cc),
	}, nil
}

func (c *Client) Auth(ctx context.Context, user *fishman_auth_v1.User) {
	const op = "clients.auth.grpc.Auth"

	resp, err := c.api.Authentication(ctx, user)
	if err != nil {
		return
	}

	fmt.Println(resp.Tokens)
}
