package configs

type AppHandler interface {

	// App returns the config specific to the application
	Get(result interface{}) error
}

func NewConfig() (AppHandler, error) {
	c, err := NewLocal()
	if err != nil {
		return nil, ErrLocal(err)
	}
	return c, err
}
