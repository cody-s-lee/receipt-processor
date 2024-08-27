# Receipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retailer** | **string** | The name of the retailer or store the receipt is from. | 
**PurchaseDate** | **string** | The date of the purchase printed on the receipt. | 
**PurchaseTime** | **string** | The time of the purchase printed on the receipt. 24-hour time expected. | 
**Items** | [**[]Item**](Item.md) |  | 
**Total** | **string** | The total amount paid on the receipt. | 

## Methods

### NewReceipt

`func NewReceipt(retailer string, purchaseDate string, purchaseTime string, items []Item, total string, ) *Receipt`

NewReceipt instantiates a new Receipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiptWithDefaults

`func NewReceiptWithDefaults() *Receipt`

NewReceiptWithDefaults instantiates a new Receipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetailer

`func (o *Receipt) GetRetailer() string`

GetRetailer returns the Retailer field if non-nil, zero value otherwise.

### GetRetailerOk

`func (o *Receipt) GetRetailerOk() (*string, bool)`

GetRetailerOk returns a tuple with the Retailer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailer

`func (o *Receipt) SetRetailer(v string)`

SetRetailer sets Retailer field to given value.


### GetPurchaseDate

`func (o *Receipt) GetPurchaseDate() string`

GetPurchaseDate returns the PurchaseDate field if non-nil, zero value otherwise.

### GetPurchaseDateOk

`func (o *Receipt) GetPurchaseDateOk() (*string, bool)`

GetPurchaseDateOk returns a tuple with the PurchaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDate

`func (o *Receipt) SetPurchaseDate(v string)`

SetPurchaseDate sets PurchaseDate field to given value.


### GetPurchaseTime

`func (o *Receipt) GetPurchaseTime() string`

GetPurchaseTime returns the PurchaseTime field if non-nil, zero value otherwise.

### GetPurchaseTimeOk

`func (o *Receipt) GetPurchaseTimeOk() (*string, bool)`

GetPurchaseTimeOk returns a tuple with the PurchaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseTime

`func (o *Receipt) SetPurchaseTime(v string)`

SetPurchaseTime sets PurchaseTime field to given value.


### GetItems

`func (o *Receipt) GetItems() []Item`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Receipt) GetItemsOk() (*[]Item, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Receipt) SetItems(v []Item)`

SetItems sets Items field to given value.


### GetTotal

`func (o *Receipt) GetTotal() string`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Receipt) GetTotalOk() (*string, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Receipt) SetTotal(v string)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


