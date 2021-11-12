package _3_Facade

type Facade struct {
	to                  string
	comment             string
	validatorToken      ValidatorToken
	validatorPermission ValidatorPermission
	store               Storage
	notificator         Email
}

func New(to, comment, token, user string) Facade {
	return Facade{
		to:                  to,
		comment:             comment,
		validatorToken:      NewValidatorToken(token),
		validatorPermission: NewValidatorPermission(user),
		store:               NewStorage("mysql"),
		notificator:         NewEmail(),
	}
}

func (myFacade *Facade) Comment() error {
	if err := myFacade.validatorToken.Validate(); err != nil {
		return err
	}
	if err := myFacade.validatorPermission.Validate(); err != nil {
		return err
	}
	myFacade.store.Save(myFacade.comment)
	myFacade.notificator.Send(myFacade.to, myFacade.comment)
	return nil
}

func (myFacade *Facade) Notify() {
	myFacade.notificator.Send(myFacade.to, myFacade.comment)
}