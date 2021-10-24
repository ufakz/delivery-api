package delivery

type DeliveryEstimateRequest struct {
	ApiKey       string `json:"api_key"`
	PickupLat    string `json:"pickup_latitude"`
	PickupLong   string `json:"pickup_longitude"`
	DeliveryLat  string `json:"delivery_latitude"`
	DeliveryLong string `json:"delivery_longitude"`
}

type EstimateResponse struct {
	Distance int `json:"distance"`
	Time     int `json:"time"`
	Fare     int `json:"fare"`
}
