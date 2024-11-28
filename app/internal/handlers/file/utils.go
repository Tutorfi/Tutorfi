package filehandler

type fileRequest struct {
    Filename string `json:"filename"`
    GroupId string `json:"groupId"`
    Data string `json:"data"`
}

type res struct {
    Status string `json:"status"`
    Msg string `json:"msg"`
    Data string `json:"data"`
}
