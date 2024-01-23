package models

type Permission struct {
    ID int64
    AccountID string
    permissions struct {
        data string
    }
}
