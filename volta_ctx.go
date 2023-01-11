package volta

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
