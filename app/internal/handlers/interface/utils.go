package interfacehandler

import "app/internal/models"

type response struct {
	Status    string   `json:"status"`
	Auth      string   `json:"auth"`
	Msg       string   `json:"msg"`
	GroupName []string `json:"groupName"`
}

type responseCreateGroup struct {
	GroupName string `json:"groupName"`
}

func fillResponse(stat string, auth string, msg string, groups []models.Group, dataType string) response {
	if groups == nil {
		return response{
			Status: stat,
			Auth:   auth,
			Msg:    msg,
		}
	}

	if dataType == "name" {
		names := make([]string, len(groups))
		for i := 0; i < len(groups); i++ {
			names[i] = groups[i].Name
		}
		return response{
			Status:    stat,
			Auth:      auth,
			Msg:       msg,
			GroupName: names,
		}
	}

	return response{
		Status: stat,
		Auth:   auth,
		Msg:    msg,
	}
}
