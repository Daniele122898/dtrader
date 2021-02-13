package gnomics

// TODO check if this is actually needed

type gnomics struct {}

type Gnomics interface {

}

func NewGnomics(apiKey string) (Gnomics, error) {
	if err := Authenticate(apiKey); err != nil {
		return nil, err
	}

	// THIS IS VERY IMPORTANT. Otherwise all receiver funcs will always copy the val.
	// It must be a reference from the start
	return &gnomics{}, nil
}
