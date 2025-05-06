package exponoti

type externalExpoNotiService struct {
	apiURL string
}

func NewExternalExpoNotiRPC(apiURL string) *externalExpoNotiService {
	return &externalExpoNotiService{apiURL: apiURL}
}
