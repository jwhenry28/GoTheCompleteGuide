package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "https://google.com",
		"Amazon Web Services": "https://aws.amazon.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])

	websites["Microsoft"] = "https://microsoft.com"
	websites["Google"] = "https://new.google.com"
	delete(websites, "Amazon Web Services")
	fmt.Println(websites)
}