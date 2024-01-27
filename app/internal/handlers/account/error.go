package accounthandler

type AccountError struct{
	message string
}
func (ac *AccountError) Error() (string){
	return ac.message
}