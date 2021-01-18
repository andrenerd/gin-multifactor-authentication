package multauth

type UserServiceProviderInterface interface {
	Send() error
}
