package volta

import (
	"io"
	"mime/multipart"
	"os"
)

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

func (c *Ctx) Form(key, def string) string {
	if val := c.Request.FormValue(key); val != "" {
		return val
	}

	return def
}

func (c *Ctx) FormFile(key string) (multipart.File, error) {
	file, _, err := c.Request.FormFile(key)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (c *Ctx) SaveFile(file multipart.File, path string) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	return nil
}

func (c *Ctx) Query(key, def string) string {
	if val := c.Request.URL.Query().Get(key); val != "" {
		return val
	}

	return def
}

func (c *Ctx) Param(key, def string) string {
	if val := c.ps.ByName(key); val != "" {
		return val
	}

	return def
}
