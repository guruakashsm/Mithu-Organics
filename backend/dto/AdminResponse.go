package dto

import "time"

// Edit Admin Response
type EditAdminResponse struct {
	Status     string    `bson:"status" json:"status"`
	StatusCode string    `bson:"statuscode" json:"statuscode"`
	Message    string    `bson:"message" json:"message"`
	EditedTime time.Time `bson:"edittime,omitempty" json:"edittime,omitempty"`
	Error      error     `json:"error,omitempty" bson:"error,omitempty"`
}

// Delete Admin Response
type DeleteAdminResponse struct {
	Status      string    `bson:"status" json:"status"`
	StatusCode  string    `bson:"statuscode" json:"statuscode"`
	Message     string    `bson:"message" json:"message"`
	DeletedTime time.Time `bson:"deletetime,omitempty" json:"deletetime,omitempty"`
	Error       error     `json:"error,omitempty" bson:"error,omitempty"`
}

// View Admin Response
type ViewAdminResponse struct {
	Status        string    `bson:"status" json:"status"`
	StatusCode    string    `bson:"statuscode" json:"statuscode"`
	Message       string    `bson:"message" json:"message"`
	ViewedTime    time.Time `bson:"viewtime,omitempty" json:"viewtime,omitempty"`
	Error         error     `json:"error,omitempty" bson:"error,omitempty"`
	AdminName     string    `json:"adminname,omitempty" bson:"adminname,omitempty"`
	AdminID       string    `json:"adminid,omitempty" bson:"adminid,omitempty"`
	Email         string    `json:"email,omitempty" bson:"email,omitempty"`
	IP_Address    string    `json:"ip,omitempty" bson:"ip,omitempty"`
	WrongInput    int       `json:"wronginput,omitempty" bson:"wronginput,omitempty"`
	LoginTime     time.Time `json:"logintime,omitempty" bson:"logintime,omitempty"`
	CreatedTime   time.Time `json:"createdtime,omitempty" bson:"createdtime,omitempty"`
	CanDeleteData bool      `json:"candelete,omitempty" bson:"candelete,omitempty"`
	CanUpdateData bool      `json:"canupdate,omitempty" bson:"canupdate,omitempty"`
	CanAlterAdmin bool      `json:"canalteradmin,omitempty" bson:"canalteradmin,omitempty"`
	CreatedBy     string    `json:"createdby,omitempty" bson:"createdby,omitempty"`
	IsBlocked     bool      `json:"isblocked,omitempty" bson:"isblocked,omitempty"`
}

// Block Admin Response
type BlockorUnblockAdminResponse struct {
	Status                 string    `bson:"status" json:"status"`
	StatusCode             string    `bson:"statuscode" json:"statuscode"`
	Message                string    `bson:"message" json:"message"`
	BlockedorUnblockedTime time.Time `bson:"blockedorunblockedtime,omitempty" json:"blockedorunblockedtime,omitempty"`
	Error                  error     `json:"error,omitempty" bson:"error,omitempty"`
}

// Create Admin Response
type CreateAdminResponse struct {
	Status       string    `bson:"status" json:"status"`
	StatusCode   string    `bson:"statuscode" json:"statuscode"`
	Message      string    `bson:"message" json:"message"`
	CreatingTime time.Time `bson:"createingtime,omitempty" json:"createingtime,omitempty"`
}

// List Deleted Admin Response
type ListDeletedAdminResponse struct {
	Status     string        `bson:"status" json:"status"`
	StatusCode string        `bson:"statuscode" json:"statuscode"`
	Message    string        `bson:"message" json:"message"`
	Error      error         `bson:"error,omitempty" json:"error,omitempty"`
	Listedtime time.Time     `bson:"listedtime,omitempty" json:"listedtime,omitempty"`
	Data       []DeleteAdmin `bson:"data,omitempty" json:"data,omitempty"`
}

// List Blocked Admin Response
type ListBlockedAdminResponse struct {
	Status     string                `bson:"status" json:"status"`
	StatusCode string                `bson:"statuscode" json:"statuscode"`
	Message    string                `bson:"message" json:"message"`
	Error      error                 `bson:"error,omitempty" json:"error,omitempty"`
	Listedtime time.Time             `bson:"listedtime,omitempty" json:"listedtime,omitempty"`
	Data       []BlockorUnblockAdmin `bson:"data,omitempty" json:"data,omitempty"`
}

// Admin Sign in data
type AdminLoginResponse struct {
	Status     string    `bson:"status" json:"status"`
	StatusCode string    `bson:"statuscode" json:"statuscode"`
	Message    string    `bson:"message" json:"message"`
	AdminName  string    `json:"adminname,omitempty" bson:"adminname,omitempty"`
	PublicKey  string    `json:"publickey,omitempty" bson:"publickey,omitempty"`
	Token      string    `bson:"token,omitempty" json:"token,omitempty"`
	Error      error     `bson:"error,omitempty" json:"error,omitempty"`
	LoginTime  time.Time `bson:"logintime,omitempty" json:"logintime,omitempty"`
}

// List Admin Response
type ListAdminResponse struct {
	Status     string      `bson:"status" json:"status"`
	StatusCode string      `bson:"statuscode" json:"statuscode"`
	Message    string      `bson:"message" json:"message"`
	Error      error       `bson:"error,omitempty" json:"error,omitempty"`
	Listedtime time.Time   `bson:"listedtime,omitempty" json:"listedtime,omitempty"`
	Data       []ListAdmin `bson:"data,omitempty" json:"data,omitempty"`
}

// ListAdmin Audit Response
type ListAdminAuditResponse struct {
	Status     string       `bson:"status" json:"status"`
	StatusCode string       `bson:"statuscode" json:"statuscode"`
	Message    string       `bson:"message" json:"message"`
	Error      error        `bson:"error,omitempty" json:"error,omitempty"`
	Listedtime time.Time    `bson:"listedtime,omitempty" json:"listedtime,omitempty"`
	Data       []AdminAudit `bson:"data,omitempty" json:"data,omitempty"`
}

// ListDeveloper Audit Response
type ListDeveloperAuditResponse struct {
	Status     string           `bson:"status" json:"status"`
	StatusCode string           `bson:"statuscode" json:"statuscode"`
	Message    string           `bson:"message" json:"message"`
	Error      error            `bson:"error,omitempty" json:"error,omitempty"`
	Listedtime time.Time        `bson:"listedtime,omitempty" json:"listedtime,omitempty"`
	Data       []DeveloperAudit `bson:"data,omitempty" json:"data,omitempty"`
}
