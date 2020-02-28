package rest

type (
	StorePostRequest struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
)
