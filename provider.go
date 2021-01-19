package multauth

type UserServiceProviderInterface interface {
	Send(to string, message string) error
}
