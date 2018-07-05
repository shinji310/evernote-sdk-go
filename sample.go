package main

import (
	"os"
	"fmt"
	"context"
	"errors"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/shinji310/evernote-sdk-go/edam"
)

const (
	SERVER = "sandbox.evernote.com"
	DEV_TOKEN = "S=s1:U=67ed:E=16b336a2f68:C=163dbb901c8:P=1cd:A=en-devtoken:V=2:H=457c5e5c5305d62b2d8d0aea10b53724"
)

/* Obtain User Store */
func GetUserStore() (*edam.UserStoreClient, error) {
	endpoint_url := fmt.Sprintf("https://%s/edam/user", SERVER)
	trans, err := thrift.NewTHttpClient(endpoint_url)
	if err != nil {
		return nil, err
	}
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	iprot := protocolFactory.GetProtocol(trans)
	oprot := protocolFactory.GetProtocol(trans)
	client := edam.NewUserStoreClient(thrift.NewTStandardClient(iprot, oprot))
	if err := trans.Open(); err != nil {
		return nil, errors.New("Error: GetUserStore() opening socket")
	}
	return client, nil
}

/* Obtain Note Store with Note Store URL */
func GetNoteStoreWithUrl(notestoreUrl string) (*edam.NoteStoreClient, error) {
	trans, err := thrift.NewTHttpClient(notestoreUrl)
	if err != nil {
		return nil, err
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault() // get a protocol factory
	iprot := protocolFactory.GetProtocol(trans)
	oprot := protocolFactory.GetProtocol(trans)
	client := edam.NewNoteStoreClient(thrift.NewTStandardClient(iprot, oprot))
	if err := trans.Open(); err != nil {
		return nil, errors.New("Error: GetNoteStoreWithURL() opening socket")
	}
	return client, nil
}

/* Obtain Note Store */
func GetNoteStore(context context.Context, authenticationToken string) (*edam.NoteStoreClient, error) {
	userstore, err := GetUserStore() // init userstore client
	if err != nil {
		return nil, err
	}
	urls, err := userstore.GetUserUrls(context, authenticationToken) // call GetUserUrls()
	if err != nil {
		return nil, err
	}
	notestoreUrl := urls.GetNoteStoreUrl() // get notestore URL

	notestore, err := GetNoteStoreWithUrl(notestoreUrl) // init notestore client
	if err != nil {
		return nil, err
	}

	return notestore, nil
}


func main() {
	ctx := context.Background() // create simplest context
	
	notestore, err := GetNoteStore(ctx, DEV_TOKEN) // obtain notestore
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	notebooks, err := notestore.ListNotebooks(ctx, DEV_TOKEN) // listing notebooks by calling API
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, nb := range notebooks {
		fmt.Println(*nb.Name)
	}
}
