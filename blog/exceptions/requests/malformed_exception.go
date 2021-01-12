package exceptions

type MalformedException struct {
	Status int
	Msg    string
}

func (mr *MalformedException) Error() string {
	return mr.Msg
}
