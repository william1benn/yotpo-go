package yotpo

import (
	"encoding/json"
	"fmt"
	"time"
)

func (y *YotpoClient) CreateReview(createReview *CreateReviewMerchant) (CreateReviewMerchantResponse, error) {
	createReviewMerchantResponse := CreateReviewMerchantResponse{}
	payloadMarshal, err := json.Marshal(createReview)
	if err != nil {
		return CreateReviewMerchantResponse{}, err
	}
	response, err := y.PostRequest("/v1/widget/reviews", payloadMarshal, nil, nil)
	if err != nil {
		return CreateReviewMerchantResponse{}, err
	}

	err = json.NewDecoder(response.Body).Decode(&createReviewMerchantResponse)
	if err != nil {
		return CreateReviewMerchantResponse{}, err
	}
	return createReviewMerchantResponse, nil
}

func (y *YotpoClient) RetrieveAllReviews(params *GetReviewsUrlParams) (GetReviewsResponse, error) {
	getReviewsResponse := GetReviewsResponse{}
	endpointUrl := fmt.Sprintf("/v1/apps/%s/reviews", y.AppKey)
	uToken, err := y.CreateUtoken()
	CheckErrorFatal(err)
	response, err := y.GetRequest(endpointUrl, params, &uToken.AccessToken)
	if err != nil {
		return GetReviewsResponse{}, err
	}
	err = json.NewDecoder(response.Body).Decode(&getReviewsResponse)
	if err != nil {
		return GetReviewsResponse{}, err
	}
	return getReviewsResponse, nil
}

type CreateReviewMerchant struct {
	OrderMetadata       OrderMetadata    `json:"order_metadata,omitempty"`
	ProductMetadata     ProductMetadata  `json:"product_metadata,omitempty"`
	CustomerMetadata    CustomerMetadata `json:"customer_metadata,omitempty"`
	IsIncentivized      bool             `json:"is_incentivized,omitempty"`
	Appkey              string           `json:"appkey"`
	Domain              string           `json:"domain,omitempty"`
	Sku                 string           `json:"sku"`
	ProductTitle        string           `json:"product_title"`
	ProductDescription  string           `json:"product_description,omitempty"`
	ProductURL          string           `json:"product_url"`
	ProductImageURL     string           `json:"product_image_url,omitempty"`
	DisplayName         string           `json:"display_name"`
	Email               string           `json:"email"`
	ReviewContent       string           `json:"review_content"`
	IncentiveType       string           `json:"incentive_type,omitempty"`
	ReviewTitle         string           `json:"review_title"`
	ReviewScore         int32            `json:"review_score,omitempty"`
	Signature           string           `json:"signature,omitempty"`
	TimeStamp           time.Time        `json:"time_stamp,omitempty"`
	SubmissionTimeStamp time.Time        `json:"submission_time_stamp,omitempty"`
	ReviewerType        string           `json:"reviewer_type,omitempty"`
	DeliveryType        string           `json:"delivery_type,omitempty"`
}
type OrderMetadata struct {
	CouponUsed bool   `json:"coupon_used,omitempty"`
	Value      string `json:"value,omitempty"`
	Name       string `json:"name,omitempty"`
}
type CustomProperties struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
type ProductMetadata struct {
	CustomProperties CustomProperties `json:"custom_properties,omitempty"`
	Color            string           `json:"color,omitempty"`
	Size             string           `json:"size,omitempty"`
	Material         string           `json:"material,omitempty"`
	Model            string           `json:"model,omitempty"`
	Vendor           string           `json:"vendor,omitempty"`
	CouponUsed       bool             `json:"coupon_used,omitempty"`
}
type CustomerMetadata struct {
	CustomProperties CustomProperties `json:"custom_properties,omitempty"`
	State            string           `json:"state,omitempty"`
	Country          string           `json:"country,omitempty"`
	Address          string           `json:"address,omitempty"`
	PhoneNumber      string           `json:"phone_number,omitempty"`
}

type CreateReviewMerchantResponse struct {
	Code             int    `json:"code"`
	Message          string `json:"message"`
	ImageUploadToken string `json:"image_upload_token"`
}

type GetReviewsResponse struct {
	GetReviews `json:"reviews"`
}

type GetReviews []struct {
	Id             int         `json:"id"`
	Title          string      `json:"title"`
	Content        string      `json:"content"`
	Score          int         `json:"score"`
	VotesUp        int         `json:"votes_up"`
	VotesDown      int         `json:"votes_down"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	Sentiment      interface{} `json:"sentiment"`
	Sku            string      `json:"sku"`
	Name           string      `json:"name"`
	Email          string      `json:"email"`
	ReviewerType   string      `json:"reviewer_type"`
	Deleted        bool        `json:"deleted"`
	Archived       bool        `json:"archived"`
	Escalated      bool        `json:"escalated"`
	IsIncentivized bool        `json:"is_incentivized"`
}

type GetReviewsUrlParams struct {
	SinceId        string    `url:"since_id,omitempty"`
	SinceDate      time.Time `url:"since_date,omitempty"`
	SinceUpdatedAt time.Time `url:"since_updated_at,omitempty"`
	Count          string    `url:"count,omitempty"`
	Page           string    `url:"page,omitempty"`
	Deleted        bool      `url:"deleted,omitempty"`
	UserReference  string    `url:"user_reference,omitempty"`
}
