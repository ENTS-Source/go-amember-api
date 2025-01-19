# go-amember-api

A library for accessing the [aMember Pro REST API](https://docs.amember.com/REST/). 

## Status

* [x] Authorization
* [ ] Users
  * [ ] Listing
  * [ ] Adding
  * [ ] Updating
  * [ ] Deleting 
  * [ ] Consent
  * [ ] Notes
  * [ ] Groups
* [ ] Invoices
  * [ ] Listing
  * [ ] Adding
* [ ] Payments
* [ ] Refunds 
* [ ] Products
  * [ ] Listing 
  * [ ] Categories
* [ ] Access
  * [ ] Log
  * [ ] Check
* [ ] Affiliate
  * [ ] Payouts

Please open issues for missing APIs you need. This library is developed as-needed.

## Usage

Enable the REST API in aMember Pro, then add a new API key.

```go
package main

import "github.com/ents-source/go-amember-api"

func main() {
  client := amember.NewClient("https://amember.example.org", "API_KEY_GOES_HERE")
  users, err := client.FindUsers(amember.UserFilter{FirstName:"Alice%"})
  // etc
}
```

Dot exploration is probably best in lieu of actual documentation.
