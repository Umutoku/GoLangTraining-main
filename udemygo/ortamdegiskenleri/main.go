package main

import (
	"fmt"
	"os"
)

func main() {
	uName := os.Getenv("USERNAME")
	uDomain := os.Getenv("USERDOMAIN")
	processorArchitecture := os.Getenv("PROCESSOR_ARCHITECTURE")
	//processorIdentifier := os.Getenv("PROCESSOR_IDENTIFIER")
	//processorLevel := os.Getenv("PROCESSOR_LEVEL")
	goRoot := os.Getenv("GOROOT")
	goPath := os.Getenv("GOPATH")
	homePath := os.Getenv("HOMEPATH")

	fmt.Println("Username: " + uName)
	fmt.Println("Domain: " + uDomain)
	fmt.Println("İşlemci Mimarisi " + processorArchitecture)
	fmt.Println("GoPath: " + goPath)
	fmt.Println("HomePath: " + homePath)
	fmt.Println("GoRoot: " + goRoot)

}
