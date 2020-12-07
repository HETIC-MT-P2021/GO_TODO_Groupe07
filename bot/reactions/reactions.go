package reactions

type reactions interface {
	MessageReactionAdd(string, string, string)  error
}
