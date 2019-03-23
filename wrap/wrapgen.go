package wrap

import "fmt"

type _impVideoUploadInform interface {
	InformVideoUploadOver(Req *interface{}, Rsp *interface{}) (ret error, err User, us *User, a int)
}

type UploadInform struct {
}

func (u *UploadInform) InformVideoUploadOver(Req *interface{}, Rsp *interface{}) (ret error, err User, us *User, a int) {
	fmt.Println("call InformVideoUploadOver")
	return nil, nil, nil, 0
}
