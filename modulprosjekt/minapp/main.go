package main

import (
    "fmt"
    "minmodul" // Importerer den lokale modulen din
)

func main() {
    // Kaller funksjonen fra modulen
    melding := minmodul.HentHilsen()
    fmt.Println(melding)
}