package gnomics

type gnomics struct {}

type Gnomics interface {

}

func NewGnomics(apiKey string) (Gnomics, error) {
	if err := Authenticate(apiKey); err != nil {
		return nil, err
	}

	return gnomics{}, nil
}
