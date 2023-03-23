package models

type ErrorJSON struct {
	Error string `json:"error"`
}

type HandleCreateJSON struct {
	Surname    string  `json:"surname"`
	Name       *string `json:"name,omitempty"`
	Secondname *string `json:"secondname,omitempty"`
	Age        *int32  `json:"age,omitempty"`
}

type HandleUpdateJSON struct {
	UID        string  `json:"uid"`
	Surname    *string `json:"surname,omitempty"`
	Name       *string `json:"name,omitempty"`
	Secondname *string `json:"secondname,omitempty"`
	Age        *int32  `json:"age,omitempty"`
}

type UserInfoListJSON struct {
	UsersInfo []UserInfoJSON `json:"users_info"`
}

type UserInfoJSON struct {
	UID                   *string `json:"uid,omitempty"`
	Surname               *string `json:"surname,omitempty"`
	Name                  *string `json:"name,omitempty"`
	Secondname            *string `json:"secondname,omitempty"`
	Age                   *int32  `json:"age,omitempty"`
	RegisterDateTimestamp *int64  `json:"register_date_timestamp,omitempty"`
}

type DeleteUserJSON struct {
	UID string `json:"uid"`
}
