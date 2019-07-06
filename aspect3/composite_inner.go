package aspect3


type inner struct {

}

func (i *inner) Call() {

}


type Comp struct {
	inner
}
