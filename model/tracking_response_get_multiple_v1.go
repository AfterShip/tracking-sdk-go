// Code generated by AfterShip SDK Generator. DO NOT EDIT.

package model

// TrackingResponseGetMultipleV1 Tracking response for getting tracking
type TrackingResponseGetMultipleV1 struct {
    // Meta Meta data
    Meta MetaV1 `json:"meta"`
    // Data 
    Data *DataTrackingResponseGetMultipleV1 `json:"data,omitempty"`
}

// DataTrackingResponseGetMultipleV1 
type DataTrackingResponseGetMultipleV1 struct {
    // Page Page to show. (Default: 1)
    Page int `json:"page,omitempty"`
    // Limit Number of trackings each page contain. (Default: 100, Max: 200)
    Limit int `json:"limit,omitempty"`
    // Count Number of returned trackings
    Count int `json:"count,omitempty"`
    // Keyword Search the content of the tracking record fields: `tracking_number`, `title`, `order_id`, `customer_name`, `custom_fields`, `order_id`, `emails`, `smses`
    Keyword string `json:"keyword,omitempty"`
    // Slug Unique 
    Slug string `json:"slug,omitempty"`
    // Origin Origin country/region of trackings. Use 
    Origin []string `json:"origin,omitempty"`
    // Destination Destination country/region of trackings. Use 
    Destination []string `json:"destination,omitempty"`
    // Tag Current status of tracking. (
    Tag TagV1 `json:"tag,omitempty"`
    // CreatedAtMin Start date and time of trackings created. AfterShip only stores data of 120 days.
    CreatedAtMin string `json:"created_at_min,omitempty"`
    // CreatedAtMax End date and time of trackings created.
    CreatedAtMax string `json:"created_at_max,omitempty"`
    // LastUpdatedAt Date and time the tracking was last updated
    LastUpdatedAt string `json:"last_updated_at,omitempty"`
    // ReturnToSender Whether or not the shipment is returned to sender. Value is `true` when any of its checkpoints has subtag `Exception_010` (returning to sender) or `Exception_011` (returned to sender). Otherwise value is `false`
    ReturnToSender []bool `json:"return_to_sender,omitempty"`
    // CourierDestinationCountryIso3 Destination country/region of the tracking detected from the courier. ISO Alpha-3 (three letters). Value will be `null` if the courier doesn't provide the destination country.
    CourierDestinationCountryIso3 []string `json:"courier_destination_country_iso3,omitempty"`
    // Trackings Array of 
    Trackings []Tracking `json:"trackings,omitempty"`
}

