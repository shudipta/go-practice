package main

import (
	"fmt"
	//"flag"
	"github.com/spf13/pflag"
//	"github.com/spf13/pflag"
)

var ip *int
var flagvar, flagvar1 int
var flagset pflag.FlagSet

func init() {
	pflag.IntVarP(&flagvar, "flagname", "i", 123, "aaaaaaaaaaaa")
	pflag.IntVarP(&flagvar1, "flagname1","j", 1231, "aaaaaaaaaaaa1")
}

func main() {
	//aaaaaaa
	ip = pflag.IntP("ip", "f", 12, "bbbbbbbbb")
	pflag.Parse()

	fmt.Println("ip has value", *ip)
	fmt.Println("flagvar has value", flagvar)
	fmt.Println("flagvar1 has value", flagvar1)
	i, err := flagset.GetInt("flagname1")
	fmt.Println("ip by GetInt() is ", i, err)

}

