package reactions

type GoToDo interface {
	MessageReactionAdd(string, string, string)  error
}
