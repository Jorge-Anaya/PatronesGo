package _3_Builder

type Director struct {
	messageBuilder MessageBuilder
}

// SetBuilder asigna el constructor
func (director *Director) SetBuilder(messageBuilder MessageBuilder) {
	director.messageBuilder = messageBuilder
}

// BuildMessage a concrete message via MessageBuilder
func (director *Director) BuildMessage(recipient, message string) (*MessageRepresented, error) {
	return director.messageBuilder.SetRecipient(recipient).SetMessage(message).Build()
}
