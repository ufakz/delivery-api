package delivery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	r := Router()

	r.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestCreateDeliveryOrder(t *testing.T) {
	reqBody := []byte(`{
        "api_key": "KTLgep79E5EnUjOmoHOyxuHBV21bZuxoCg5iHw9CfMqdo79dTjLTYZCRVHAu",
        "pickup_address": "Vita Towers, Kofo Abayomi Street, Lagos, Nigeria",
        "pickup_latitude": "6.435275500000001",
        "pickup_longitude": "3.4147874",
        "delivery_address": "21 Lugard Avenue, Lagos, Nigeria",
        "delivery_latitude": "6.456150699999999",
        "delivery_longitude": "3.4298536",
        "pickup_name": "Foo",
        "pickup_phone": "+23412345678",
        "pickup_email": "example@about.com",
        "delivery_name": "Bar",
        "delivery_phone": "+2341234567891",
        "delivery_email": "example+1@about.com",
        "description": "Food delivery"
	}`)

	req, _ := http.NewRequest("POST", "/delivery_order", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	fmt.Println(m)

	if m["order_id"] != "FFFXY-426043" {
		t.Errorf("Expected order id to be 'FFFXY-426043'. Got '%v'", m["order_id"])
	}

	if m["distance"] != 5.0 {
		t.Errorf("Expected distance to be 5. Got %v", m["distance"])
	}

	if m["time"] != 12.0 {
		t.Errorf("Expected time to be 12. Got %v", m["time"])
	}

	if m["fare"] != 500.0 {
		t.Errorf("Expected fare to be 500. Got %v", m["fare"])
	}

}
