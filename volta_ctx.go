package volta

func (c *Ctx) SendString(msg string) error {
	c.Response.Write([]byte(msg))

	return nil
}

func (c *Ctx) Send(msg []byte) error {
	c.Response.Write(msg)

	return nil
}
