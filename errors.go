package fsmv

// UnknownErr Possibly caused by coding errors
type UnknownErr struct {
}

func (e UnknownErr) Error() string {
	return "unknown error"
}

// IllegalConfigErr Visual pack haven't init
type IllegalConfigErr struct {
}

func (e IllegalConfigErr) Error() string {
	return "visualization package config is not illegal"
}

// IllegalWrapperErr Visual pack haven't init
type IllegalWrapperErr struct {
}

func (e IllegalWrapperErr) Error() string {
	return "visualization wrapper is not illegal"
}
