package controllers

type Dispatcher struct {
	PhoneController     *ofPhones
	UsersController     *ofUsers
	VariablesController *ofVariables
}
