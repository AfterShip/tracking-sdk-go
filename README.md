# AfterShip Tracking API library for Go

This library allows you to quickly and easily use the AfterShip Tracking API via Go.

For updates to this library, see our [GitHub release page](https://github.com/AfterShip/tracking-sdk-go/releases).

If you need support using AfterShip products, please contact support@aftership.com.

## Table of Contents

- [AfterShip Tracking API library for Go](#aftership-tracking-api-library-for-go)
  - [Table of Contents](#table-of-contents)
  - [Before you begin](#before-you-begin)
    - [API and SDK Version](#api-and-sdk-version)
  - [Quick Start](#quick-start)
    - [Installation](#installation)
  - [Constructor](#constructor)
    - [Example](#example)
  - [Rate Limiter](#rate-limiter)
  - [Error Handling](#error-handling)
    - [Error List](#error-list)
  - [Endpoints](#endpoints)
    - [/estimated-delivery-date](#estimated-delivery-date)
    - [/trackings](#trackings)
    - [/couriers](#couriers)
    - [/courier-connections](#courier-connections)
  - [Help](#help)
  - [License](#license)


## Before you begin

Before you begin to integrate:

- [Create an AfterShip account](https://admin.aftership.com/).
- [Create an API key](https://organization.automizely.com/api-keys).
- [Install Go](https://go.dev/dl/) version Go 1.16 or later.

### API and SDK Version

- SDK Version: 
- API Version: 2025-07
## Quick Start

### Installation
```bash
go get -u github.com/aftership/tracking-sdk-go/v10
```

## Constructor

Create AfterShip instance with options

| Name       | Type   | Required | Description                                                                                                                       |
| ---------- | ------ | -------- | --------------------------------------------------------------------------------------------------------------------------------- |
| api_key    | string | ✔        | Your AfterShip API key                                                                                                            |
| auth_type  | enum   |          | Default value: `AuthType.API_KEY` <br > AES authentication: `AuthType.AES` <br > RSA authentication: `AuthType.RSA`               |
| api_secret | string |          | Required if the authentication type is `AuthType.AES` or `AuthType.RSA`                                                           |
| domain     | string |          | AfterShip API domain. Default value: https://api.aftership.com                                                                    |
| user_agent | string |          | User-defined user-agent string, please follow [RFC9110](https://www.rfc-editor.org/rfc/rfc9110#field.user-agent) format standard. |
| proxy      | string |          | HTTP proxy URL to use for requests. <br > Default value: `null` <br > Example: `http://192.168.0.100:8888`                        |
| max_retry  | number |          | Number of retries for each request. Default value: 2. Min is 0, Max is 10.                                                        |
| timeout    | number |          | Timeout for each request in milliseconds.                                                                                         |

### Example

```go
package main

import (
    "fmt"
    "github.com/aftership/tracking-sdk-go/v10"
    
    "github.com/aftership/tracking-sdk-go/v10/operation"
)

func main() {
    sdk, err := tracking.New(tracking.WithApiKey("YOUR_API_KEY"))
    if err != nil {
        fmt.Println(err)
        return
    }

    query := operation.GetTrackingByIdQuery{}
    result, err := sdk.Tracking.GetTrackingById().
        BuildPath("valid_value",).
        BuildQuery(query).
        Execute()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result)
}
```

## Rate Limiter

See the [Rate Limit](https://www.aftership.com/docs/tracking/quickstart/rate-limit) to understand the AfterShip rate limit policy.

## Error Handling

The SDK will return an error object when there is any error during the request, with the following specification:

| Name            | Type   | Description                    |
| --------------- | ------ | ------------------------------ |
| message         | string | Detail message of the error    |
| code            | enum   | Error code enum for API Error. |
| meta_code       | number | API response meta code.        |
| status_code     | number | HTTP status code.              |
| response_body   | string | API response body.             |
| response_header | object | API response header.           |


### Error List

| code                              | meta_code       | status_code     | message                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| --------------------------------- | --------------- | --------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| INVALID_REQUEST | 400 | 400 | The request was invalid or cannot be otherwise served. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| INVALID_JSON | 4001 | 400 | Invalid JSON data. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| TRACKING_ALREADY_EXIST | 4003 | 400 | Tracking already exists. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| TRACKING_DOES_NOT_EXIST | 4004 | 404 | Tracking does not exist. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| TRACKING_NUMBER_INVALID | 4005 | 400 | The value of tracking_number is invalid. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| TRACKING_REQUIRED | 4006 | 400 | tracking object is required. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| TRACKING_NUMBER_REQUIRED | 4007 | 400 | tracking_number is required. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| VALUE_INVALID | 4008 | 400 | The value of [field_name] is invalid. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| VALUE_REQUIRED | 4009 | 400 | [field_name] is required. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| SLUG_INVALID | 4010 | 400 | The value of slug is invalid. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| MISSING_OR_INVALID_REQUIRED_FIELD | 4011 | 400 | Missing or invalid value of the required fields for this courier. Besides tracking_number, also required: [field_name] |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| BAD_COURIER | 4012 | 400 | The error message will be one of the following:1. Unable to import shipment as the carrier is not on your approved list for carrier auto-detection. Add the carrier here: https://admin.aftership.com/settings/couriers2. Unable to import shipment as we don&#39;t recognize the carrier from this tracking number.3. Unable to import shipment as the tracking number has an invalid format.4. Unable to import shipment as this carrier is no longer supported.5. Unable to import shipment as the tracking number does not belong to a carrier in that group. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| INACTIVE_RETRACK_NOT_ALLOWED | 4013 | 400 | Retrack is not allowed. You can only retrack an inactive tracking. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| NOTIFICATION_REQUIRED | 4014 | 400 | notification object is required. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| ID_INVALID | 4015 | 400 | The value of id is invalid. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| RETRACK_ONCE_ALLOWED | 4016 | 400 | Retrack is not allowed. You can only retrack each shipment once. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| TRACKING_NUMBER_FORMAT_INVALID | 4017 | 400 | The format of tracking_number is invalid. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| API_KEY_INVALID | 401 | 401 | The API Key is invalid. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| REQUEST_NOT_ALLOWED | 403 | 403 | The request is understood, but it has been refused or access is not allowed. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| NOT_FOUND | 404 | 404 | The URI requested is invalid or the resource requested does not exist. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| TOO_MANY_REQUEST | 429 | 429 | You have exceeded the API call rate limit. The default limit is 10 requests per second. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| INTERNAL_ERROR | 500 | 500 | Something went wrong on AfterShip&#39;s end. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| INTERNAL_ERROR | 502 | 502 | Something went wrong on AfterShip&#39;s end. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| INTERNAL_ERROR | 503 | 503 | Something went wrong on AfterShip&#39;s end. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| INTERNAL_ERROR | 504 | 504 | Something went wrong on AfterShip&#39;s end. |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |

## Endpoints

The AfterShip SDK has the following resource which are exactly the same as the API endpoints:

- EstimatedDeliveryDateResource
  - Prediction for the Estimated Delivery Date
  - Batch prediction for the Estimated Delivery Date
- TrackingResource
  - Get trackings
  - Create a tracking
  - Get a tracking by ID
  - Update a tracking by ID
  - Delete a tracking by ID
  - Retrack an expired tracking by ID
  - Mark tracking as completed by ID
- CourierResource
  - Get couriers
  - Detect courier
- CourierConnectionResource
  - Get courier connections
  - Create courier connections
  - Get courier connection by id
  - Update courier connection by id
  - Delete courier connection by id

### /estimated-delivery-date
**POST** /estimated-delivery-date/predict

```go
    body := model.EstimatedDeliveryDateRequest{}
    body.SetSlug("valid_value")
    originAddressVal := model.EstimatedDeliveryDateRequestOriginAddress{}
    body.SetOriginAddress(originAddressVal)
    destinationAddressVal := model.EstimatedDeliveryDateRequestDestinationAddress{}
    body.SetDestinationAddress(destinationAddressVal)
    result, err := sdk.EstimatedDeliveryDate.Predict().
        BuildBody(body).
        Execute()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result)
```

**POST** /estimated-delivery-date/predict-batch

```go
    body := model.PredictBatchRequest{}
    result, err := sdk.EstimatedDeliveryDate.PredictBatch().
        BuildBody(body).
        Execute()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result)
```

### /trackings
**GET** /trackings

```go
    query := operation.GetTrackingsQuery{}
    result, err := sdk.Tracking.GetTrackings().
        BuildQuery(query).
        Execute()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result)
```

**POST** /trackings

```go
    body := model.CreateTrackingRequest{}
    body.SetTrackingNumber("valid_value")
    result, err := sdk.Tracking.CreateTracking().
        BuildBody(body).
        Execute()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result)
```

**GET** /trackings/{id}

```go
    query := operation.GetTrackingByIdQuery{}
    result, err := sdk.Tracking.GetTrackingById().
        BuildPath("valid_value").
        BuildQuery(query).
        Execute()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result)
```

**PUT** /trackings/{id}

```go
    body := model.UpdateTrackingByIdRequest{}
    result, err := sdk.Tracking.UpdateTrackingById().
        BuildPath("valid_value").
        BuildBody(body).
        Execute()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result)
```

**DELETE** /trackings/{id}

```go
    result, err := sdk.Tracking.DeleteTrackingById().
        BuildPath("valid_value").
        Execute()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result)
```

**POST** /trackings/{id}/retrack

```go
    result, err := sdk.Tracking.RetrackTrackingById().
        BuildPath("valid_value").
        Execute()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result)
```

**POST** /trackings/{id}/mark-as-completed

```go
    body := model.MarkTrackingCompletedByIdRequest{}
    result, err := sdk.Tracking.MarkTrackingCompletedById().
        BuildPath("valid_value").
        BuildBody(body).
        Execute()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result)
```

### /couriers
**GET** /couriers

```go
    query := operation.GetCouriersQuery{}
    result, err := sdk.Courier.GetCouriers().
        BuildQuery(query).
        Execute()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result)
```

**POST** /couriers/detect

```go
    body := model.DetectCourierRequest{}
    body.SetTrackingNumber("valid_value")
    result, err := sdk.Courier.DetectCourier().
        BuildBody(body).
        Execute()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result)
```

### /courier-connections
**GET** /courier-connections

```go
    query := operation.GetCourierConnectionsQuery{}
    result, err := sdk.CourierConnection.GetCourierConnections().
        BuildQuery(query).
        Execute()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result)
```

**POST** /courier-connections

```go
    body := model.PostCourierConnectionsRequest{}
    body.SetCourierSlug("valid_value")
    body.SetCredentials()
    result, err := sdk.CourierConnection.PostCourierConnections().
        BuildBody(body).
        Execute()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result)
```

**GET** /courier-connections/{id}

```go
    result, err := sdk.CourierConnection.GetCourierConnectionsById().
        BuildPath("valid_value").
        Execute()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result)
```

**PATCH** /courier-connections/{id}

```go
    body := model.PutCourierConnectionsByIdRequest{}
    body.SetCredentials()
    result, err := sdk.CourierConnection.PutCourierConnectionsById().
        BuildPath("valid_value").
        BuildBody(body).
        Execute()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result)
```

**DELETE** /courier-connections/{id}

```go
    result, err := sdk.CourierConnection.DeleteCourierConnectionsById().
        BuildPath("valid_value").
        Execute()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result)
```


## Help

If you get stuck, we're here to help:

- [Issue Tracker](https://github.com/AfterShip/tracking-sdk-go/issues) for questions, feature requests, bug reports and general discussion related to this package. Try searching before you create a new issue.
- Contact AfterShip official support via support@aftership.com

## License
Copyright (c) 2025 AfterShip

Licensed under the MIT license.