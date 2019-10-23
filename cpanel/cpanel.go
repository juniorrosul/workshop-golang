package main

import (
	"encoding/json"
	"fmt"

	"github.com/letsencrypt-cpanel/cpanelgo"
	"github.com/letsencrypt-cpanel/cpanelgo/whm"
)

func main() {
	token := "IX8EKOXECN13XN5W38OPKWCIZ1E1NPTQ"
	hostname := "3.83.159.247"
	password := "Senh@Fort3?"
	username := "root"
	insecure := true

	whmcl := whm.NewWhmApiAccessHash(hostname, username, token, insecure)

	var out json.RawMessage

	args := cpanelgo.Args{}
	args["username"] = "xablauzinho"
	args["password"] = password
	args["domain"] = "xablauzinho.tk"

	errCreate := whmcl.WHMAPI1("createacct", args, &out)

	if errCreate != nil {
		fmt.Println(errCreate)
	}

	fmt.Println(string(out))

	args = cpanelgo.Args{}
	args["username"] = "xablauzinho"

	errRemove := whmcl.WHMAPI1("removeacct", args, &out)

	if errRemove != nil {
		fmt.Println(errRemove)
	}

	fmt.Println(string(out))
}
