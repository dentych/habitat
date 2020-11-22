package menus

type Option struct {
	Key         byte
	Description string
	Handler     func() Menu
}

func (o Option) Handle() Menu {
	if o.Handler != nil {
		return o.Handler()
	}
	return nil
}
