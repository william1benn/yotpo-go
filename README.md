## Yotpo-Go 

### In Progress --- Experimental (*Expect breaking changes*)

___
Core Yotpo API Functionality in development order

* _Reviews - Merchant_
* _Reviews Metadata_ 
* _Questions and Answers (Q&A)_
* _Dynamic Coupons_
* _Email Analytics_

___

### Quick Start 
_Import Yotpo-go package_
```
go git https://github.com/william1benn/yotpo-go
```

_Example Client and Method_
```go
import (
	y "https://github.com/william1benn/yotpo-go"
) 

//Create Client
yClient := y.NewYotpoClient("AppIdString", "ApiSecretKey")

//Invoke Method
response, _ := yClient.RetrieveAllReviews(nil)

for _, r := range response.GetReviews { 
    fmt.Println(r.Name)
}
```
