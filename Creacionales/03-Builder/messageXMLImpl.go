package _3_Builder

import (
	"encoding/xml"
)

// XMLMessageBuilder is concrete builder
type XMLMessageBuilder struct {
	message Message
}

// SetRecipient asigna el receptor del mensaje
func (myXMLMessageBuilder *XMLMessageBuilder) SetRecipient(recipient string) MessageBuilder {
	myXMLMessageBuilder.message.Recipient = recipient
	return myXMLMessageBuilder
}

// SetMessage asigna el mensaje a enviar
func (myXMLMessageBuilder *XMLMessageBuilder) SetMessage(text string) MessageBuilder {
	myXMLMessageBuilder.message.Text = text
	return myXMLMessageBuilder
}

// Build construye el objeto y lo representa en XML
func (myXMLMessageBuilder *XMLMessageBuilder) Build() (*MessageRepresented, error) {
	data, err := xml.Marshal(myXMLMessageBuilder.message)
	if err != nil {
		return nil, err
	}
	return &MessageRepresented{Body: data, Format: "XML"}, nil
}
