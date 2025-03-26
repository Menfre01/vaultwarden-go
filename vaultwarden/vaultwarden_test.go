package vaultwarden

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Menfre01/vaultwarden-go/pkg/types"
)

// func TestClient_postIdentityToken(t *testing.T) {
// 	os.Setenv("BW_CLIENTID", "user.d1f448d1-3bdb-4e36-9583-17098cf83cc5")
// 	os.Setenv("BW_CLIENTSECRET", "eGZd5112mICN5PMvwaD6JK0rTlSIOh")
// 	os.Setenv("BW_SERVER", "https://vaultwarden.foottr.com")
// 	tests := []struct {
// 		name    string
// 		c       *Client
// 		wantErr bool
// 	}{
// 		{
// 			name: "test client_postIdentityToken",
// 			c: &Client{
// 				server:  os.Getenv("BW_SERVER"),
// 				httpCli: &http.Client{},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := tt.c.postIdentityToken(); (err != nil) != tt.wantErr {
// 				t.Errorf("Client.postIdentityToken() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestClient_GetProfile(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		c       *Client
// 		wantErr bool
// 	}{
// 		{
// 			name: "test client_GetProfile",
// 			c: &Client{
// 				server:  os.Getenv("BW_SERVER"),
// 				httpCli: &http.Client{},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.c.GetProfile()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Client.GetProfile() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 			if got != nil {
// 				fmt.Println(got)
// 			}
// 		})
// 	}
// }

// func TestClient_makeRequest(t *testing.T) {
// 	type args struct {
// 		method    string
// 		path      string
// 		headers   map[string]string
// 		bodyBytes []byte
// 		auth      bool
// 	}
// 	tests := []struct {
// 		name    string
// 		c       *Client
// 		args    args
// 		want    *http.Request
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.c.makeRequest(tt.args.method, tt.args.path, tt.args.headers, tt.args.bodyBytes, tt.args.auth)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Client.makeRequest() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Client.makeRequest() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestClient_getToken(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		c       *Client
// 		want    string
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.c.getToken()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Client.getToken() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != tt.want {
// 				t.Errorf("Client.getToken() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestClient_GetCollections(t *testing.T) {
// 	type args struct {
// 		organizationId string
// 	}
// 	tests := []struct {
// 		name    string
// 		c       *Client
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "test client_GetCollections",
// 			c: &Client{
// 				server:  os.Getenv("BW_SERVER"),
// 				httpCli: &http.Client{},
// 			},
// 			args: args{
// 				organizationId: "23dcc8a6-c55a-440c-b3cf-339de69c4933",
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := tt.c.GetCollections(tt.args.organizationId)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Client.GetCollections() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if got != nil {
// 				fmt.Println(got)
// 			}
// 		})
// 	}
// }

func TestClient_ListOrganizations(t *testing.T) {
	type args struct {
		search string
	}
	tests := []struct {
		name    string
		c       ClientInterface
		args    args
		wantErr bool
	}{
		{
			name: "test client_ListOrganizations",
			c:    NewClient("http://localhost:8087", &http.Client{}),
			args: args{
				search: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.ListOrganizations(tt.args.search)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ListOrganizations() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != nil {
				fmt.Println(got)
			}
		})
	}
}

func TestClient_ListCollectionsWithOrg(t *testing.T) {
	type args struct {
		organizationId string
		search         string
	}
	tests := []struct {
		name    string
		c       ClientInterface
		args    args
		wantErr bool
	}{
		{
			name: "test client_ListCollectionsWithOrg",
			c:    NewClient("http://localhost:8087", &http.Client{}),
			args: args{
				search:         "",
				organizationId: "614827bf-0e68-4b41-9dfd-2a5ebcb01439",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.ListCollectionsWithOrg(tt.args.organizationId, tt.args.search)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ListCollectionsWithOrg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				fmt.Println(got)
			}
		})
	}
}

func TestClient_ListItems(t *testing.T) {
	type args struct {
		organizationId string
		collectionId   string
		search         string
	}
	tests := []struct {
		name    string
		c       ClientInterface
		args    args
		wantErr bool
	}{
		{
			name: "test client_ListItems",
			c:    NewClient("http://localhost:8087", &http.Client{}),
			args: args{
				search:         "",
				organizationId: "614827bf-0e68-4b41-9dfd-2a5ebcb01439",
				collectionId:   "fc4c565d-ea65-4590-8842-6e2b38ce7e4d",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.ListItems(tt.args.organizationId, tt.args.collectionId, tt.args.search)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.ListItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				fmt.Println(got)
			}
		})
	}
}

func TestClient_CreateItem(t *testing.T) {
	type args struct {
		item *types.Item
	}
	tests := []struct {
		name    string
		c       ClientInterface
		args    args
		wantErr bool
	}{
		{
			name: "test client_CreateItem",
			c:    NewClient("http://localhost:8087", &http.Client{}),
			args: args{
				item: &types.Item{
					Name:     "demo2",
					Type:     types.ITEMTYPE_SECURENOTE,
					Notes:    "demo2",
					Fields:   []*types.Field{},
					Favorite: false,
					Reprompt: 0,
					SecureNote: &types.SecureNote{
						Type: 0,
					},
					FolderId: nil,
					CollectionIds: []string{
						"fc4c565d-ea65-4590-8842-6e2b38ce7e4d",
					},
					OrganizationId: "614827bf-0e68-4b41-9dfd-2a5ebcb01439",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CreateItem(tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.CreateItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				fmt.Println(got)
			}
		})
	}
}

func TestClient_EditItem(t *testing.T) {
	type args struct {
		id   string
		item *types.Item
	}
	tests := []struct {
		name    string
		c       ClientInterface
		args    args
		wantErr bool
	}{
		{
			name: "test client_EditItem",
			c:    NewClient("http://localhost:8087", &http.Client{}),
			args: args{
				id: "734fa1b2-2d55-4ddc-bb33-34aaddfdc567",
				item: &types.Item{
					Name:     "demo2",
					Type:     types.ITEMTYPE_SECURENOTE,
					Notes:    "demo3",
					Fields:   []*types.Field{},
					Favorite: false,
					Reprompt: 0,
					SecureNote: &types.SecureNote{
						Type: 0,
					},
					FolderId: nil,
					CollectionIds: []string{
						"fc4c565d-ea65-4590-8842-6e2b38ce7e4d",
					},
					OrganizationId: "614827bf-0e68-4b41-9dfd-2a5ebcb01439",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.EditItem(tt.args.id, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.EditItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				fmt.Println(got)
			}
		})
	}
}
