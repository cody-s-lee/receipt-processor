# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReceiptsIdPointsGet**](DefaultAPI.md#ReceiptsIdPointsGet) | **Get** /receipts/{id}/points | Returns the points awarded for the receipt
[**ReceiptsProcessPost**](DefaultAPI.md#ReceiptsProcessPost) | **Post** /receipts/process | Submits a receipt for processing



## ReceiptsIdPointsGet

> ReceiptsIdPointsGet200Response ReceiptsIdPointsGet(ctx, id).Execute()

Returns the points awarded for the receipt



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/cody-s-lee/receipt-processor/backend/openapi"
)

func main() {
	id := "id_example" // string | The ID of the receipt

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ReceiptsIdPointsGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ReceiptsIdPointsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReceiptsIdPointsGet`: ReceiptsIdPointsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ReceiptsIdPointsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the receipt | 

### Other Parameters

Other parameters are passed through a pointer to a apiReceiptsIdPointsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReceiptsIdPointsGet200Response**](ReceiptsIdPointsGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReceiptsProcessPost

> ReceiptsProcessPost200Response ReceiptsProcessPost(ctx).Receipt(receipt).Execute()

Submits a receipt for processing



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/cody-s-lee/receipt-processor/backend/openapi"
)

func main() {
	receipt := *openapiclient.NewReceipt("M&M Corner Market", time.Now(), "13:01", []openapiclient.Item{*openapiclient.NewItem("Mountain Dew 12PK", "6.49")}, "6.49") // Receipt | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ReceiptsProcessPost(context.Background()).Receipt(receipt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ReceiptsProcessPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReceiptsProcessPost`: ReceiptsProcessPost200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ReceiptsProcessPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReceiptsProcessPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **receipt** | [**Receipt**](Receipt.md) |  | 

### Return type

[**ReceiptsProcessPost200Response**](ReceiptsProcessPost200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

