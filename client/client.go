package client

import (
	"fmt"
	"context"
	"errors"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/shinji310/evernote-sdk-go/edam"
)

type EnvironmentType int

/*
Definition of runtime types
*/
const (
	SANDBOX EnvironmentType = iota // sandbox for development and testing
	PRODUCTION // production environment
	YINXIANG // Evernote service in China
)

/*
Evernote Client data structure
*/
type EvernoteClient struct {
	host string
	token string
	userStore *edam.UserStoreClient
	noteStore *edam.NoteStoreClient
}

/*
Create EvernoteClient
*/
func NewEvernoteClient(token string, envType EnvironmentType) *EvernoteClient {
	host := "www.evernote.com"
	if envType == SANDBOX {
		host = "sandbox.evernote.com"
	} else if envType == YINXIANG {
		host = "app.yinxiang.com"
	}

	return &EvernoteClient{
		host: host,
		token: token,
	}
}

/* 
Obtain UserStore 
*/
func (c *EvernoteClient) GetUserStore() (*edam.UserStoreClient, error) {
	if c.userStore == nil {
		endpoint_url := fmt.Sprintf("https://%s/edam/user", c.host)
		trans, err := thrift.NewTHttpClient(endpoint_url)
		if err != nil {
			return nil, err
		}
		if err := trans.Open(); err != nil {
			return nil, errors.New("Error: GetUserStore() opening socket")
		}
		protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
		iprot := protocolFactory.GetProtocol(trans)
		oprot := protocolFactory.GetProtocol(trans)
		c.userStore = edam.NewUserStoreClient(thrift.NewTStandardClient(iprot, oprot))
	}
	
	return c.userStore, nil
}

/* 
Obtain Note Store using Note Store URL 
*/
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

/* 
Obtain Note Store 
*/
func (c *EvernoteClient) GetNoteStore(context context.Context) (*edam.NoteStoreClient, error) {
	if c.noteStore == nil {
		us, err := c.GetUserStore() // obtain UserStoreClient
		if err != nil {
			return nil, err
		}
		urls, err := us.GetUserUrls(context, c.token) // get UserUrls
		if err != nil {
			return nil, err
		}
		nsUrl := urls.GetNoteStoreUrl() // grab NoteStore URL in the UserUrls

		c.noteStore, err = GetNoteStoreWithUrl(nsUrl) // init notestore client
		if err != nil {
			return nil, err
		}
	}

	return c.noteStore, nil
}

/*
Setter, Getter
*/
func (c *EvernoteClient) GetAuthenticationToken() (string) {
	return c.token
}
