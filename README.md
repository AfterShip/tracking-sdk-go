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
    - [/trackings](#trackings)
    - [/couriers](#couriers)
    - [/estimated-delivery-date](#estimated-delivery-date)
  - [Help](#help)
  - [License](#license)


## Before you begin

Before you begin to integrate:

- [Create an AfterShip account](https://admin.aftership.com/).
- [Create an API key](https://organization.automizely.com/api-keys).
- [Install Go](https://go.dev/dl/) version Go 1.16 or later.

### API and SDK Version

Each SDK version is designed to work with a specific API version. Please refer to the table below to identify the supported API versions for each SDK version, ensuring you select the appropriate SDK version for the API version you intend to use.

| SDK Version | Supported API Version | Branch                                                    |
| ----------- | --------------------- | --------------------------------------------------------- |
| v8.x.x      | 2025-04               | https://github.com/AfterShip/tracking-sdk-go/tree/2025-04 |
| v7.x.x      | 2025-01               | https://github.com/AfterShip/tracking-sdk-go/tree/2025-01 |
| v6.x.x      | 2024-10               | https://github.com/AfterShip/tracking-sdk-go/tree/2024-10 |
| v5.x.x      | 2024-07               | https://github.com/AfterShip/tracking-sdk-go/tree/2024-07 |
| v4.x.x      | 2024-04               | https://github.com/AfterShip/tracking-sdk-go/tree/2024-04 |
| v3.x.x      | 2023-10               | https://github.com/AfterShip/aftership-sdk-go             |
| <=v2.x.x    | Legacy API            | https://github.com/AfterShip/aftership-sdk-go             |

## Quick Start

### Installation
```bash
go get -u github.com/aftership/tracking-sdk-go/v8
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
    "github.com/aftership/tracking-sdk-go/v8"
)

func main() {
	sdk, err := tracking.New(tracking.WithApiKey("YOUR_API_KEY"))
	if err != nil {
		fmt.Println(err)
        return
	}
	result, err := sdk.Tracking.GetTrackingById().
		BuildPath("<tracking_id>").
		Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
```

## Rate Limiter

See the [Rate Limit](https://www.aftership.com/docs/tracking/2025-04/quickstart/api-quick-start) to understand the AfterShip rate limit policy.

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
| INVALID_REQUEST                   | 400             | 400             | The request was invalid or cannot be otherwise served.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| INVALID_JSON                      | 4001            | 400             | Invalid JSON data.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| TRACKING_ALREADY_EXIST            | 4003            | 400             | Tracking already exists.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| TRACKING_DOES_NOT_EXIST           | 4004            | 404             | Tracking does not exist.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| TRACKING_NUMBER_INVALID           | 4005            | 400             | The value of tracking_number is invalid.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| TRACKING_REQUIRED                 | 4006            | 400             | tracking object is required.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| TRACKING_NUMBER_REQUIRED          | 4007            | 400             | tracking_number is required.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| VALUE_INVALID                     | 4008            | 400             | The value of [field_name] is invalid.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| VALUE_REQUIRED                    | 4009            | 400             | [field_name] is required.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| SLUG_INVALID                      | 4010            | 400             | The value of slug is invalid.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| MISSING_OR_INVALID_REQUIRED_FIELD | 4011            | 400             | Missing or invalid value of the required fields for this courier. Besides tracking_number, also required: [field_name]                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| BAD_COURIER                       | 4012            | 400             | The error message will be one of the following:<br/>1. Unable to import shipment as the carrier is not on your approved list for carrier auto-detection. Add the carrier here: https://admin.aftership.com/settings/couriers<br/>2. Unable to import shipment as we don’t recognize the carrier from this tracking number.<br/>3. Unable to import shipment as the tracking number has an invalid format.<br/>4. Unable to import shipment as this carrier is no longer supported.<br/>5. Unable to import shipment as the tracking number does not belong to a carrier in that group. |
| INACTIVE_RETRACK_NOT_ALLOWED      | 4013            | 400             | Retrack is not allowed. You can only retrack an inactive tracking.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| NOTIFICATION_REUQIRED             | 4014            | 400             | notification object is required.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| ID_INVALID                        | 4015            | 400             | The value of id is invalid.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| RETRACK_ONCE_ALLOWED              | 4016            | 400             | Retrack is not allowed. You can only retrack each shipment once.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| TRACKING_NUMBER_FORMAT_INVALID    | 4017            | 400             | The format of tracking_number is invalid.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| API_KEY_INVALID                   | 401             | 401             | The API key is invalid.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| REQUEST_NOT_ALLOWED               | 403             | 403             | The request is understood, but it has been refused or access is not allowed.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| NOT_FOUND                         | 404             | 404             | The URI requested is invalid or the resource requested does not exist.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| TOO_MANY_REQUEST                  | 429             | 429             | You have exceeded the API call rate limit. The default limit is 10 requests per second.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| INTERNAL_ERROR                    | 500 502 503 504 | 500 502 503 504 | Something went wrong on AfterShip's end.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |

## Endpoints

The AfterShip instance has the following properties which are exactly the same as the API endpoints:

- courier - Get a list of our supported couriers.
- tracking - Create trackings, update trackings, and get tracking results.
- estimated-delivery-date - Get estimated delivery date for your order.


### /trackings

**POST** /trackings

```go
data := model.CreateTrackingRequest{
    TrackingNumber: "<tracking_number>",
    Slug:           "<slug>",
}
result, err := sdk.Tracking.CreateTracking().BuildBody(data).Execute()
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(result)
```

**DELETE** /trackings/:id

```go
result, err := sdk.Tracking.
    DeleteTrackingById().
    BuildPath("<tracking_id>").
    Execute()
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(result)
```

**GET** /trackings

```go
result, err := sdk.Tracking.
    GetTrackings().
    BuildQuery(model.GetTrackingsQuery{Keyword: "1234"}).
    Execute()
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(result)
```

**GET** /trackings/:id

```go
result, err := sdk.Tracking.
    GetTrackingById().
    BuildPath("<tracking_id>").
    Execute()
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(result)
```

**PUT** /trackings/:id

```go
result, err := sdk.Tracking.
    UpdateTrackingById().
    BuildPath("<tracking_id>").
    BuildBody(model.UpdateTrackingByIdRequest{
        Title: "test",
    }).Execute()
if err != nil {
fmt.Println(err)
    return
}
fmt.Println(result)
```

**POST** /trackings/:id/retrack

```go
result, err := sdk.Tracking.
    RetrackTrackingById().
    BuildPath("<tracking_id>").
    Execute()
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(result)
```

**POST** /trackings/:id/mark-as-completed

```go
result, err := sdk.Tracking.
    MarkTrackingCompletedById().
    BuildPath("<tracking_id>").
    BuildBody(model.MarkTrackingCompletedByIdRequest{Reason: "DELIVERED"}).
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
result, err := sdk.Courier.GetUserCouriers().Execute()
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(result)
```

**POST** /couriers/detect

```go
result, err := sdk.Courier.
    DetectCourier().
    BuildBody(model.DetectCourierRequest{
        TrackingNumber: "<tracking_number>",
    }).
    Execute()
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(result)
```

### /estimated-delivery-date

**POST** /estimated-delivery-date/predict-batch

```go
result, err := sdk.EstimatedDeliveryDate.
    PredictBatch().
    BuildBody(model.PredictBatchRequest{
        EstimatedDeliveryDates: []model.EstimatedDeliveryDateRequest,
    }).
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