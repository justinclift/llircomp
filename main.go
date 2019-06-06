package main

import (
    "log"

    "github.com/kr/pretty"
    "github.com/llir/llvm/asm"
)

func main() {
    m, err := asm.ParseFile("target.ll")
    if err != nil {
        log.Fatalf("%+v", err)
    }
    pretty.Println(m)
}
