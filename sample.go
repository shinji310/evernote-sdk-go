package main

import (
	"os"
	"fmt"
	"context"
	"github.com/shinji310/evernote-sdk-go/client"
)

const (
	SERVER = "sandbox.evernote.com"
	DEV_TOKEN = "S=s1:U=67ed:E=16c54446645:C=164fc933840:P=1cd:A=en-devtoken:V=2:H=018800bb8fb400bf883633835556fd6f"
)


/*
Sample code - List all the notebooks
*/
func main() {
	ctx := context.Background() // create simplest context
	enClient := client.NewEvernoteClient(DEV_TOKEN, client.SANDBOX) // create EvernoteClient
	
	ns, err := enClient.GetNoteStore(ctx) // obtain notestore
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	notebooks, err := ns.ListNotebooks(ctx, enClient.GetAuthenticationToken()) // listing notebooks by calling API
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, nb := range notebooks {
		fmt.Println(*nb.Name)
	}
}
