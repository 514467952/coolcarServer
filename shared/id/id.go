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

// IdentityID
type IdentityID string

func (i IdentityID) String() string {
	return string(i)
}

// CarID
type CarID string

func (i CarID) String() string {
	return string(i)
}

// BlobID
type BlobID string

func (i BlobID) String() string {
	return string(i)
}
