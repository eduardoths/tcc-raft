package dto

type SetBody struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (sb SetBody) ToArgs() SetArgs {
	return SetArgs{
		Key:   sb.Key,
		Value: []byte(sb.Value),
	}
}

type GetResponse struct {
	Value string `json:"string"`
}
