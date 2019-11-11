package serializers

type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
}

type DataList struct {
	Items interface{} `json:"items"`
	Total uint        `json:"total"`
}

func BuildList(items interface{}, total uint) DataList {
	return DataList{
			Items: items,
			Total: total,
	}
}
