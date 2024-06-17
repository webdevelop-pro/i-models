package emails

type Provider interface {
	Send(email *Email) error
	Cancel(email *Email) error
}
