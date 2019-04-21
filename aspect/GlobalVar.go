package aspect

type GlobalVar struct {
	key string
}

var GlobalVarIns *GlobalVar = &GlobalVar{
	key: "true",
}
