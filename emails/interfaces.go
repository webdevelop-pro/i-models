package emails

type Provider interface {
	Send(email *EmailEmail) error
	Cancel(email *EmailEmail) error
}
