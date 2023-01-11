package volta

func (c *Ctx) Status(status Status) *Ctx {
	c.Response.WriteHeader(int(status))
	return c
}

func (c *Ctx) Redirect(url string) error {
	c.Response.Header().Set("Location", url)
	c.Response.WriteHeader(int(StatusSeeOther))

	return nil
}

func (c *Ctx) RedirectStatus(url string, status int) error {
	c.Response.Header().Set("Location", url)
	c.Response.WriteHeader(status)

	return nil
}

func (c *Ctx) SendString(msg string) error {
	c.Response.Write([]byte(msg))

	return nil
}

func (c *Ctx) Send(msg []byte) error {
	c.Response.Write(msg)

	return nil
}

func (c *Ctx) SendJSON(msg Map) error {
	if c.jsonMarshaler == nil {
		return ErrorNoJSONMarshaler
	}

	json, err := c.jsonMarshaler(msg)
	if err != nil {
		c.Response.Write([]byte(err.Error()))
		return nil
	}

	c.Response.Write(json)

	return nil
}
