package gosrv

type Info struct {
	Name        string `json:"name"`
	Alias       string `json:"alias"`
	Description string `json:"description"`
}

func (i Info) getName() string {
	return i.Name
}

func (i Info) getAlias() string {
	return i.Alias
}

func (i Info) getDescription() string {
	return i.Description
}
