package yotpo

import (
	"fmt"
	"testing"
)

func TestYotpoCreateReview(t *testing.T) {
}

func TestRetreiveAllReviews(t *testing.T) {

	yClient := NewYotpoClient("AppIdString", "ApiSecretKey")
	response, _ := yClient.RetrieveAllReviews(nil)

	for _, r := range response.GetReviews {
		fmt.Println(r.Name)
	}
}
