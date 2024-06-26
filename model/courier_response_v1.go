// Code generated by AfterShip SDK Generator. DO NOT EDIT.

package model

// CourierResponseV1 Model of all couriers endpoint response
type CourierResponseV1 struct {
    // Meta Meta data
    Meta MetaV1 `json:"meta"`
    // Data 
    Data *DataCourierResponseV1 `json:"data"`
}

// DataCourierResponseV1 
type DataCourierResponseV1 struct {
    // Total Total count of courier objects
    Total int `json:"total,omitempty"`
    // Couriers Array of  object.
    Couriers []Courier `json:"couriers,omitempty"`
}


