package id

/*
强类型限定
*/

//限定AccountID为string类型，
type AccountID string

func (a AccountID) String() string {
	return string(a)
}

//限定AccountID为string类型，
type TripID string

func (t TripID) String() string {
	return string(t)
}
