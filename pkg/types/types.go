package types

type BaseResponse[T any] struct {
	Data         *ListResponse[T] `json:"data"`
	Success      bool             `json:"success"`
	DeleteData   string           `json:"deleteData"`
	RevisionDate string           `json:"revisionDate"`
}

type ListResponse[T any] struct {
	Data []*T `json:"data"`
	// ContinuationToken string `json:"continuationToken"`
	Object string `json:"object"`
}

type OrganizationResponse struct {
	Object  string `json:"object"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	Status  int    `json:"status"`
	Type    int    `json:"type"`
	Enabled bool   `json:"enabled"`
}

type CollectionResponse struct {
	Object         string `json:"object"`
	ID             string `json:"id"`
	Name           string `json:"name"`
	OrganizationId string `json:"organizationId"`
	ExternalId     string `json:"externalId"`
}

type ItemResponse struct {
	Object         string      `json:"object"`
	ID             string      `json:"id"`
	OrganizationId string      `json:"organizationId"`
	FolderId       string      `json:"folderId"`
	Type           int         `json:"type"`
	Reprompt       int         `json:"reprompt"`
	Name           string      `json:"name"`
	Notes          string      `json:"notes"`
	Favorite       bool        `json:"favorite"`
	SecureNote     *SecureNote `json:"secureNote"`
	CollectionIds  []string    `json:"collectionIds"`
	RevisionDate   string      `json:"revisionDate"`
	CreationDate   string      `json:"creationDate"`
	DeletedDate    string      `json:"deletedDate"`
}

type Item struct {
	Name           string      `json:"name"`
	Type           ItemType    `json:"type"`
	Notes          string      `json:"notes"`
	Favorite       bool        `json:"favorite"`
	CollectionIds  []string    `json:"collectionIds"`
	OrganizationId string      `json:"organizationId"`
	FolderId       *string     `json:"folderId"`
	Reprompt       int         `json:"reprompt"`
	SecureNote     *SecureNote `json:"secureNote"`
	Fields         []*Field    `json:"fields"`
}

type SecureNote struct {
	Type int `json:"type"`
}

type Field struct {
	Type     FieldType `json:"type"`
	Name     string    `json:"name"`
	Value    string    `json:"value"`
	LinkedId string    `json:"linkedId"`
}

type FieldType int

const (
	FIELDTYPE_TEXT FieldType = iota
	FIELDTYPE_HIDDEN
	FIELDTYPE_BOOLEAN
)

type ItemType int

const (
	ITEMTYPE_LOGIN      ItemType = 1
	ITEMTYPE_SECURENOTE ItemType = 2
	ITEMTYPE_CARD       ItemType = 3
	ITEMTYPE_IDENTITY   ItemType = 4
)
