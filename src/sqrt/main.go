package main

import (
    "net/http"
    "log"
    "io"
    "math"
)

func main() {
	http.HandleFunc("/", WebServer)

    port := "80"

    log.Println("Rodando na porta " + port)

    if err := http.ListenAndServe(":" + port, nil); err != nil {
        log.Fatal(err)
    }
}

func WebServer(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, rocksSqrt())
}

func rocksSqrt() string {
    x := 0.0001

    for i := 0; i <= 1000000; i++ {
        x += math.Sqrt(x)
    }

    return "Code.education Rocks! "
}