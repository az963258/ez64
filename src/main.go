package main

import (
    "encoding/base64"
    "flag"
    "fmt"
    "os"
)

func main() {
    encodeFlag := flag.Bool("e", false, "Encode input to Base64")
    decodeFlag := flag.Bool("d", false, "Decode Base64 input")

    flag.Parse()

    input := flag.Arg(0)

    if *encodeFlag && len(input) > 0 {
        encoded := base64.StdEncoding.EncodeToString([]byte(input))
        fmt.Println(encoded)
    } else if *decodeFlag && len(input) > 0 {
        decoded, err := base64.StdEncoding.DecodeString(input)
        if err != nil {
            fmt.Fprintln(os.Stderr, "Error decoding:", err)
            os.Exit(1)
        }
        fmt.Println(string(decoded))
    } else {
        fmt.Println("Usage:")
        fmt.Println("\t-e <input> : Encode input to Base64")
        fmt.Println("\t-d <input> : Decode Base64 input")
    }
}
