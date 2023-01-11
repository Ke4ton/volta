package volta

import (
	"io"
	"mime/multipart"
	"os"
)

func (c *Ctx) Next() error {
	return ErrorNext
}

// Header returns the value of the header field with the given name.
func (c *Ctx) Header(key, def string) string {
	if val := c.Request.Header.Get(key); val != "" {
		return val
	}

	return def
}

// SetHeader sets the header with the given key and value.
func (c *Ctx) SetHeader(key, val string) *Ctx {
	c.Response.Header().Set(key, val)
	return c
}

// Status sets the status code.
func (c *Ctx) Status(status Status) *Ctx {
	c.Response.WriteHeader(int(status))
	return c
}

// Redirect redirects to the given url.
func (c *Ctx) Redirect(url string) error {
	c.Response.Header().Set("Location", url)
	c.Response.WriteHeader(int(StatusSeeOther))

	return nil
}

// RedirectStatus redirects to the given url with the given status.
func (c *Ctx) RedirectStatus(url string, status int) error {
	c.Response.Header().Set("Location", url)
	c.Response.WriteHeader(status)

	return nil
}

// SendString sends a string response.
func (c *Ctx) SendString(msg string) error {
	c.Response.Write([]byte(msg))

	return nil
}

// Send sends a byte array response.
func (c *Ctx) Send(msg []byte) error {
	c.Response.Write(msg)

	return nil
}

// SendJSON sends a JSON response.
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

// Form returns the value of the form field with the given name.
func (c *Ctx) Form(key, def string) string {
	if val := c.Request.FormValue(key); val != "" {
		return val
	}

	return def
}

// Form returns the value of the form file field with the given name.
func (c *Ctx) FormFile(key string) (multipart.File, error) {
	file, _, err := c.Request.FormFile(key)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// SaveFile saves the file to the given path.
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

// Query returns the value of the query field with the given name.
func (c *Ctx) Query(key, def string) string {
	if val := c.Request.URL.Query().Get(key); val != "" {
		return val
	}

	return def
}

// Param returns the value of the param field with the given name.
func (c *Ctx) Param(key, def string) string {
	if val := c.ps.ByName(key); val != "" {
		return val
	}

	return def
}
