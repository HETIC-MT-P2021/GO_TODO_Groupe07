package gotodo

type GoToDo interface {
	MessageReactionAdd(string, string, string)  error
}
