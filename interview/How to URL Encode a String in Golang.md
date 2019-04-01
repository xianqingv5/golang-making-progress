// https://www.urlencoder.io/golang/
// URL Encoding a Query string in Golang
package main 
import (
    "fmt"
    "net/url"
)
func main() {
    query := "Hello world@Golang"
    fmt.Println(url.QueryEscape(query))
}

// URL Encoding multiple Query parameters in Golang
package main
import (
    "fmt"
    "net/url"
)
func main() {
    params := url.Values{}
    params.Add("name", "songhq")
    params.Add("phone", "1111111")

    fmt.Println(params.Encode())
}