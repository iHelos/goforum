package helper

func PanicCheck(err error){
	if err!=nil{
		panic(err)
	}
}