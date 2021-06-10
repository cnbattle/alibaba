
# alibabaopen

## Use Demo
```go
package main

import (
	"fmt"
    
	"ruan.co/alibaba/open"
)

func main() {
    client := &open.Client{
        AppKey:      "4790000",
        AppSecret:   "xxxxx",
        AccessToken: "xxxxxxxx-eb92-4b15-857c-f0e89f5c7684",
    }
    params := make(map[string]string)
    params["page"] = "1"
    params["pageSize"] = "10"
    var uri = "com.alibaba.p4p:alibaba.cps.op.searchCybOffers-1"
    do, err := client.Do(uri, params)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(do))
}

```
