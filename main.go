package main

import "fmt"

type Samochod []string

func main() {
	czesci := czesciSamochodowe()
	fmt.Println(czesci)
}

func czesciSamochodowe() Samochod {
	czesci := Samochod{}

	wszystkieCzesci := []string{"kola", "drzwi", "maska", "szyba", "siedzenia", "kierownica"}

	for _, cz := range wszystkieCzesci {
		czesci = append(czesci, cz)
	}

	return czesci

}
