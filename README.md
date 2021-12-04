# CClient

Fixes TLS and stuff.

# Example

```go
package main

import (
    "log"

    tls "github.com/tfyl/utls"
    "github.com/tfyl/cclient"
)

func main() {
    client, err := cclient.NewClient(tls.HelloChrome_Auto,"",true,6)
    if err != nil {
        log.Fatal(err)
    }

    resp, err := client.Get("https://www.google.com/")
    if err != nil {
        log.Fatal(err)
    }
    resp.Body.Close()

    log.Println(resp.Status)
}
```


