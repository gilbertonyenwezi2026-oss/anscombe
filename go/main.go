package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/your-username/anscombe-go-python-r/go/anscombe"
)

func main() {
	results, err := anscombe.RunQuartet()
	if err != nil {
		log.Fatal(err)
	}

	for _, result := range results {
		fmt.Printf("%s -> intercept=%.5f slope=%.5f r=%.5f r^2=%.5f\n",
			result.Name, result.Intercept, result.Slope, result.R, result.RSquared)
	}

	fmt.Println("\nJSON:")
	blob, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(blob))
}
