// Code generated by AfterShip SDK Generator. DO NOT EDIT.

package model

// MarkTrackingCompletedBySlugTrackingNumberQuery 
type MarkTrackingCompletedBySlugTrackingNumberQuery struct {
    // TrackingAccountNumber Additional field required by some carriers to retrieve the tracking info. The shipper’s carrier account number. Refer to our article on  for more details.
    TrackingAccountNumber string `url:"tracking_account_number,omitempty"`
    // TrackingOriginCountry Additional field required by some carriers to retrieve the tracking info. The origin country/region of the shipment. Refer to our article on  for more details.
    TrackingOriginCountry string `url:"tracking_origin_country,omitempty"`
    // TrackingDestinationCountry Additional field required by some carriers to retrieve the tracking info. The destination country/region of the shipment. Refer to our article on  for more details.
    TrackingDestinationCountry string `url:"tracking_destination_country,omitempty"`
    // TrackingKey Additional field required by some carriers to retrieve the tracking info. A type of tracking credential required by some carriers. Refer to our article on  for more details.
    TrackingKey string `url:"tracking_key,omitempty"`
    // TrackingPostalCode Additional field required by some carriers to retrieve the tracking info. The postal code of the recipient’s address. Refer to our article on  for more details.
    TrackingPostalCode string `url:"tracking_postal_code,omitempty"`
    // TrackingShipDate Additional field required by some carriers to retrieve the tracking info. The date the shipment was sent, using the format YYYYMMDD. Refer to our article on  for more details.
    TrackingShipDate string `url:"tracking_ship_date,omitempty"`
    // TrackingState Additional field required by some carriers to retrieve the tracking info. The state/province of the recipient’s address. Refer to our article on  for more details.
    TrackingState string `url:"tracking_state,omitempty"`
}


