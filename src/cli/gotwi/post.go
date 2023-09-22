package gotwi

import (
	"github.com/michimani/gotwi"
	"github.com/michimani/gotwi/tweet/managetweet"
	"github.com/michimani/gotwi/tweet/managetweet/types"
	"fmt"
	"context"
)

func PostTweet(t string) (string, error) {
	
	in := &gotwi.NewClientWithAccessTokenInput{
		AccessToken: "#",
	}

	c, err := gotwi.NewClientWithAccessToken(in)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c, "done")
	}

	tweet := &types.CreateInput {
		Text: gotwi.String(t),
	}

	res, err := managetweet.Create(context.Background(), c, tweet)

	if err != nil {
		return "", err
	}

	return gotwi.StringValue(res.Data.ID), nil
}