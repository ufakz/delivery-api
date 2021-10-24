package docs

import "github.com/farouqu/delivery-api/delivery"

//swagger:route /delivery_order delivery_order idOrder
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
