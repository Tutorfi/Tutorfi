package accounthandler

type AccountError struct{
	msg string
}
func (ac *AccountError) Error() (string){
	return ac.msg
}