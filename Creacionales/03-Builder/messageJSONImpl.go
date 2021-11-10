package _3_Builder

import (
	"encoding/json"
)

// JSONMessageBuilder is concrete builder
type JSONMessageBuilder struct {
	message Message
}

// SetRecipient asigna el receptor del mensaje
func (myJSONMessageBuilder *JSONMessageBuilder) SetRecipient(recipient string) MessageBuilder {
	myJSONMessageBuilder.message.Recipient = recipient
	return myJSONMessageBuilder
}

// SetMessage asigna el mensaje a enviar
func (myJSONMessageBuilder *JSONMessageBuilder) SetMessage(text string) MessageBuilder {
	myJSONMessageBuilder.message.Text = text
	return myJSONMessageBuilder
}

// Build construye el objeto y lo representa en JSON
func (myJSONMessageBuilder *JSONMessageBuilder) Build() (*MessageRepresented, error) {
	data, err := json.Marshal(myJSONMessageBuilder.message)
	if err != nil {
		return nil, err
	}
	return &MessageRepresented{Body: data, Format: "JSON"}, nil
}
