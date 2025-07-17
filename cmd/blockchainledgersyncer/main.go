// cmd/blockchainledgersyncer/main.go
package main

import (
"flag"
"log"
"os"

"blockchainledgersyncer/internal/blockchainledgersyncer"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainledgersyncer.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
