package payprogo

func New(key string) *payPro {
	return &payPro{
		key,
		"https://www.paypro.nl/post_api/",
		false,
	}
}

type payPro struct {
	key   string
	url   string
	debug bool
}

func (p *payPro) Debug(d bool) {
	p.debug = d
}

func (p *payPro) NewCommand(c string) *command {
	r := &command{
		p.url,
		c,
		p.key,
		p.debug,
		make(map[string]interface{}),
	}

	if p.debug {
		r.Set("test_mode", "true")
	}

	return r
}
