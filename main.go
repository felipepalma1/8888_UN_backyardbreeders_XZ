package main

import (
	"fmt"
	"github.com/felipepalma1/8888_UN_backyardbreeders_XZ/models"
)

func main() {
	var specimen = &models.Information{Name: "Killer S3372", Age: "2 Years", Weight: "15 Kilograms", Genre: "Female"}
	fmt.Println("Backyard Breeders")
	output := fmt.Sprintf("New specimen %s %s %s %s", specimen.Name, specimen.Age, specimen.Weight, specimen.Genre)
	fmt.Println(output)
}
