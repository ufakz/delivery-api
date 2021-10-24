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
