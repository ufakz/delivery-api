package delivery

type DeliveryEstimateRequest struct {
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

type CreateDeliveryRequest struct {
	PickupAdd     string `json:"pickup_address"`
	PickupLat     string `json:"pickup_latitue"`
	PickupLong    string `json:"pickup_longitude"`
	DeliveryAdd   string `json:"delivery_address"`
	DeliveryLat   string `json:"delivery_latitude"`
	DeliveryLong  string `json:"delivery_longitude"`
	PickupName    string `json:"pickup_name"`
	PickupPhone   string `json:"pickup_phone"`
	PickupEmail   string `json:"pickup_email,omitempty"`
	DeliveryName  string `json:"delivery_name"`
	DeliveryPhone string `json:"delivery_phone"`
	DeliveryEmail string `json:"delivery_email,omitempty"`
	Description   string `json:"descriptioin,omitempty"`
}

type DeliveryResponse struct {
	OrderID  string `json:"order_id"`
	Distance string `json:"distance"`
	Time     string `json:"time"`
	Fare     string `json:"fare"`
}
