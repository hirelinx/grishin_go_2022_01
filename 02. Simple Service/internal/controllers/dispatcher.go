package controllers

var messagesPass = func(d dispatcher, str []string) []any {
	return d.MessagesController().Pass(str)
}

var variablesPass = func(d dispatcher, str []string) []any {
	return d.VariablesController().Pass(str)
}

var usersPass = func(d dispatcher, str []string) []any {
	return d.UsersController().Pass(str)
}

var dispatcherPassMap = map[string]func(dispatcher, []string) []any{
	"messages":  messagesPass,
	"variables": variablesPass,
	"users":     usersPass,
}

func DispatcherPassMap(dest map[string]func(dispatcher, []string) []any) {
	dest = make(map[string]func(dispatcher, []string) []any)
	for k, v := range dispatcherPassMap {
		dest[k] = v
	}
	return
}

type Dispatcher interface {
	TokensManager
	Dispatch(strings []string) string
}

type dispatcher struct {
	passMap             map[string]func(dispatcher, []string) []any
	usersController     Users
	variablesController Variables
	messagesController  Messages
}

func (d dispatcher) Pass(strings []string) []any {
	return d.passMap[strings[0]](d, strings[1:])
}

func (d dispatcher) Dispatch(strings []string) string {
	return ""
}

func (d dispatcher) UsersController() Users {
	return d.usersController
}

func (d dispatcher) VariablesController() Variables {
	return d.variablesController
}

func (d dispatcher) MessagesController() Messages {
	return d.messagesController
}

func NewDispatcher(usersController Users, variablesController Variables, messagesController Messages) (dispatch dispatcher) {
	dispatch = dispatcher{
		passMap:             dispatcherPassMap,
		usersController:     usersController,
		variablesController: variablesController,
		messagesController:  messagesController,
	}
	return
}
