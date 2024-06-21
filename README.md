# AfterShip Tracking API library for Go

This library allows you to quickly and easily use the AfterShip Tracking API via Go.

For updates to this library, see our [GitHub release page](https://github.com/AfterShip/tracking-sdk-go/releases).

If you need support using AfterShip products, please contact support@aftership.com.

## Table of Contents

- [AfterShip Tracking API library for Go](#aftership-tracking-api-library-for-go)
  - [Table of Contents](#table-of-contents)
  - [Before you begin](#before-you-begin)
  - [Quick Start](#quick-start)
    - [Installation](#installation)
    - [Usage](#usage)
  - [Constructor](#constructor)
    - [Example](#example)
  - [Rate Limiter](#rate-limiter)
  - [Error Handling](#error-handling)
    - [Error List](#error-list)
  - [Endpoints](#endpoints)
    - [/trackings](#trackings)
    - [/couriers](#couriers)
    - [/last\_checkpoint](#last_checkpoint)
    - [/notifications](#notifications)
    - [/estimated-delivery-date](#estimated-delivery-date)
  - [Help](#help)
  - [License](#license)


## Before you begin

Before you begin to integrate:

- [Create an AfterShip account](https://admin.aftership.com/).
- [Create an API key](https://organization.automizely.com/api-keys).
- [Install Go](https://go.dev/dl/) version Go 1.16 or later.

## Quick Start

### Installation
```bash
go get -u github.com/aftership/tracking-sdk-go/v4
```

### Usage
```go
package main

import (
    "fmt"
    "github.com/aftership/tracking-sdk-go/v4"
)

func main() {
    sdk, err := tracking.New(WithApiKey("YOUR_API_KEY"))
	if err != nil {
		fmt.Println(err)
        return
	}
	result, err := sdk.Tracking.GetTrackingById().
		BuildPath("1234567890").
		Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
```


## Constructor

Create AfterShip instance with options

| Name       | Type   | Required | Description                                                                                                                       |
|------------|--------|----------|-----------------------------------------------------------------------------------------------------------------------------------|
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
    "github.com/aftership/aftership-tracking-sdk-go"
    "github.com/AfterShip/tracking-sdk-go/v4/component"
)

func main() {
	sdk, err := tracking.New(
		WithAuthKind(component.Aes),
		WithApiKey("YOUR_API_KEY"),
		WithSecret("YOUR_API_SECRET"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := sdk.Tracking.GetTrackingById().
		BuildPath("kponlnb1w64fmlxlakyln00l").
		Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
```

## Rate Limiter

See the [Rate Limit](https://www.aftership.com/docs/aftership/quickstart/rate-limit) to understand the AfterShip rate limit policy.

## Error Handling

The SDK will return an error object when there is any error during the request, with the following specification:

| Name          | Type   | Description                    |
|---------------|--------|--------------------------------|
| message       | string | Detail message of the error    |
| code          | enum   | Error code enum for API Error. |
| meta_code     | number | API response meta code.        |
| status_code   | number | HTTP status code.              |
| response_body | string | API response body.             |


### Error List

| code                              | meta_code        | status_code     | message                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | 
|-----------------------------------|------------------|-----------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| INVALID_REQUEST                   | 	400             | 	400	           | The request was invalid or cannot be otherwise served.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| INVALID_JSON                      | 	4001            | 	400	           | Invalid JSON data.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| TRACKING_ALREADY_EXIST            | 	4003            | 	400	           | Tracking already exists.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| TRACKING_DOES_NOT_EXIST           | 	4004            | 	404	           | Tracking does not exist.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |                             
| TRACKING_NUMBER_INVALID           | 	4005            | 	400	           | The value of tracking_number is invalid.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| TRACKING_REQUIRED                 | 	4006            | 	400	           | tracking object is required.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| TRACKING_NUMBER_REQUIRED          | 	4007            | 	400	           | tracking_number is required.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| VALUE_INVALID                     | 	4008            | 	400	           | The value of [field_name] is invalid.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| VALUE_REQUIRED                    | 	4009            | 	400	           | [field_name] is required.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| SLUG_INVALID                      | 	4010            | 	400	           | The value of slug is invalid.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| MISSING_OR_INVALID_REQUIRED_FIELD | 	4011            | 	400	           | Missing or invalid value of the required fields for this courier. Besides tracking_number, also required: [field_name]                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| BAD_COURIER                       | 	4012            | 	400	           | The error message will be one of the following:<br/>1. Unable to import shipment as the carrier is not on your approved list for carrier auto-detection. Add the carrier here: https://admin.aftership.com/settings/couriers<br/>2. Unable to import shipment as we don’t recognize the carrier from this tracking number.<br/>3. Unable to import shipment as the tracking number has an invalid format.<br/>4. Unable to import shipment as this carrier is no longer supported.<br/>5. Unable to import shipment as the tracking number does not belong to a carrier in that group. |
| INACTIVE_RETRACK_NOT_ALLOWED      | 	4013            | 	400	           | Retrack is not allowed. You can only retrack an inactive tracking.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| NOTIFICATION_REUQIRED             | 	4014            | 	400	           | notification object is required.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| ID_INVALID                        | 	4015            | 	400	           | The value of id is invalid.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| RETRACK_ONCE_ALLOWED              | 	4016            | 	400	           | Retrack is not allowed. You can only retrack each shipment once.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| TRACKING_NUMBER_FORMAT_INVALID    | 	4017            | 	400	           | The format of tracking_number is invalid.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| API_KEY_INVALID                   | 	401             | 	401	           | The API key is invalid.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| REQUEST_NOT_ALLOWED               | 	403             | 	403            | The request is understood, but it has been refused or access is not allowed.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| NOT_FOUND                         | 	404             | 	404	           | The URI requested is invalid or the resource requested does not exist.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| TOO_MANY_REQUEST                  | 	429             | 	429            | You have exceeded the API call rate limit. The default limit is 10 requests per second.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| INTERNAL_ERROR	                   | 500 502 503 504	 | 500 502 503 504 | Something went wrong on AfterShip's end.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |

## Endpoints

The AfterShip instance has the following properties which are exactly the same as the API endpoints:

- courier - Get a list of our supported couriers.
- tracking - Create trackings, update trackings, and get tracking results.
- last_checkpoint - Get tracking information of the last checkpoint of a tracking.
- notification - Get, add or remove contacts (sms or email) to be notified when the status of a tracking has changed.


### /trackings

**POST** /trackings

```go
data := model.CreateTrackingRequest{
    &model.TrackingCreateTrackingRequest{
        TrackingNumber: "9434609105464265845274",
        Slug:           "usps",
    },
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
    BuildPath("txs1jvxyx6c7wlxishv5h024").
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
    BuildPath("r71re5kn21mmylxegrzxh01s").
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
    BuildPath("r71re5kn21mmylxegrzxh01s").
    BuildBody(model.UpdateTrackingByIdRequest{&model.TrackingUpdateTrackingByIdRequest{
        Smses: []string{"+85291239123"},
    }}).Execute()
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
    BuildPath("hqhyzb21sm0colweuats7001").
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
    BuildPath("hqhyzb21sm0colweuats7001").
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

**GET** /couriers/all

```go
result, err := sdk.Courier.GetAllCouriers().Execute()
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(result)
```

**POST** /couriers/detect

```js
result, err := sdk.Courier.
    DetectCourier().
    BuildBody(model.DetectCourierRequest{
        Tracking: &model.TrackingDetectCourierRequest{
            TrackingNumber: "9361289752032444756119",
        },
    }).
    Execute()
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(result)
```

### /last_checkpoint

**GET** /last_checkpoint/:id

```go
result, err := sdk.LastCheckpoint.
    GetCheckpointByTrackingId().
    BuildPath("r71re5kn21mmylxegrzxh01s").
    Execute()
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(result)
```

### /notifications

**GET** /notifications/:id

```go
result, err := sdk.Notification.
    AddNotificationByTrackingId().
    BuildPath("r71re5kn21mmylxegrzxh01s").
    BuildBody(model.AddNotificationByTrackingIdRequest{model.NotificationRequestV1{
        Emails: []string{"test@gmail.com"},
    }}).
    Execute()
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(result)
```

**POST** /notifications/:id/add

```go
result, err := sdk.Notification.
    AddNotificationByTrackingId().
    BuildPath("r71re5kn21mmylxegrzxh01s").
    BuildBody(model.AddNotificationByTrackingIdRequest{model.NotificationRequestV1{
        Emails: []string{"your_mail@gmail.com"},
    }}).
    Execute()
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(result)
```

**POST** /notifications/:id/remove

```go
result, err := sdk.Notification.
    DeleteNotificationByTrackingId().
    BuildPath("r71re5kn21mmylxegrzxh01s").
    BuildBody(model.DeleteNotificationByTrackingIdRequest{model.NotificationRequestV1{
        Emails: []string{"your_mail@gmail.com"},
    }}).
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
Copyright (c) 2024 AfterShip

Licensed under the MIT license.