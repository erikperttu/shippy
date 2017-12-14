package main

type Authable interface {
	Decode(token string)(interface{}, error)
	Encode(data interface{}) (string, error)
}

type TokenService struct {
	repo Repository
}

// De-encode stub
func (src *TokenService) Decode(token string) (interface{}, error){
	return "", nil
}

// Encode stub
func (src *TokenService) Encode(data interface{}) (interface{}, error){
	return "", nil
}


