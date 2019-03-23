package wrap

import (
	"log"
	"reflect"
	"time"
)

type Wraps_impVideoUploadInform struct {
	Orign interface{}
}

func (w *Wraps_impVideoUploadInform) InformVideoUploadOver(Req *interface{}, Rsp *interface{}) (error, User, *User, int) {

	var start = time.Now()
	log.Println("starting func InformVideoUploadOver")
	defer func(start time.Time) {
		log.Printf("ended func InformVideoUploadOver after %f seconds \n", time.Since(start).Seconds())
	}(start)

	result := reflect.ValueOf(w.Orign).MethodByName("InformVideoUploadOver").Call([]reflect.Value{reflect.ValueOf(Req), reflect.ValueOf(Rsp)})

	return func(a interface{}) error {
			if a == nil {
				return nil
			}
			return a.(error)
		}(result[0].Interface()), func(a interface{}) User {
			if a == nil {
				return nil
			}
			return a.(User)
		}(result[1].Interface()), func(a interface{}) *User {
			if a == nil {
				return nil
			}
			return a.(*User)
		}(result[2].Interface()), func(a interface{}) int {
			if a == nil {
				return nil
			}
			return a.(int)
		}(result[3].Interface())
}
