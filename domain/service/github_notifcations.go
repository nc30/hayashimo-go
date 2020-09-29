package service

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func GetNotifcationLength(ctx context.Context, access_token string) (int, error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: access_token},
	)

	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	n, _, e := client.Activity.ListNotifications(ctx, nil)

	return len(n), e
}
