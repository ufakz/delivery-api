package docs

import "github.com/farouqu/delivery-api/delivery"

// swagger:route POST /delivery_estimate delivery_estimate idDelivery
// Fetches the estimate for a delivery order.
// responses:
//   200: estimateSuccessResponse
// swagger:parameters idDelivery
type estimateParamsWrapper struct {
	// Fetches the estimate for a delivery order.
	// in:body
	Body delivery.DeliveryEstimateRequest
}

// Successful estimate created response.
// swagger:response estimateSuccessResponse
type estimateResponseWrapper struct {
	// in:body
	Body delivery.EstimateResponse
}

//swagger:route POST /delivery_order delivery_order idOrder
//Creates a delivery order
//responses:
// 200: deliverySuccessResponse

//swagger:parameters idOrder
type deliveryParamsWrapper struct {
	//Creates a delivery order and returns the information
	//in:body
	Body delivery.CreateDeliveryRequest
}

//Successful delivery order
//swagger:response deliverySuccessResponse
type deliveryResponseWrapper struct {
	//in:body
	Body delivery.DeliveryResponse
}
