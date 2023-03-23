package models

type ErrorJSON struct {
	Error string `json:"error"`
}

type HandleCreateJSON struct {
	Name    string  `json:"name"`
	Address *string `json:"address,omitempty"`
	Working *bool   `json:"working,omitempty"`
	Owner   *string `json:"owner,omitempty"`
}

type HandleUpdateJSON struct {
	UID     string  `json:"uid"`
	Name    *string `json:"name"`
	Address *string `json:"address,omitempty"`
	Working *bool   `json:"working,omitempty"`
	Owner   *string `json:"owner,omitempty"`
}

type StoreInfoListJSON struct {
	StoresInfo []StoreInfoJSON `json:"stores_info"`
}

type StoreInfoJSON struct {
	UID     *string `json:"uid,omitempty"`
	Name    *string `json:"name,omitempty"`
	Address *string `json:"address,omitempty"`
	Working *bool   `json:"working,omitempty"`
	Owner   *string `json:"owner,omitempty"`
}

type DeleteStoreJSON struct {
	UID string `json:"uid"`
}
