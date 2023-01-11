package volta

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
)

func (c *Ctx) Next() error {
	return ErrorNext
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

// Body returns the request body bytes.
func (c *Ctx) Body() ([]byte, error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// IsGet returns true if the request method is GET.
func (c *Ctx) IsGet() bool {
	if c.Request.Method == "GET" {
		return true
	}

	return false
}

// IsPost returns true if the request method is POST.
func (c *Ctx) IsPost() bool {
	if c.Request.Method == "POST" {
		return true
	}

	return false
}

// IsPut returns true if the request method is PUT.
func (c *Ctx) IsPut() bool {
	if c.Request.Method == "PUT" {
		return true
	}

	return false
}

// IsDelete returns true if the request method is DELETE.
func (c *Ctx) IsDelete() bool {
	if c.Request.Method == "DELETE" {
		return true
	}

	return false
}

// IsPatch returns true if the request method is PATCH.
func (c *Ctx) IsPatch() bool {
	if c.Request.Method == "PATCH" {
		return true
	}

	return false
}

// IsOptions returns true if the request method is OPTIONS.
func (c *Ctx) IsOptions() bool {
	if c.Request.Method == "OPTIONS" {
		return true
	}

	return false
}

// Method returns the request method.
func (c *Ctx) Method() string {
	return c.Request.Method
}

// UserAgent returns the user agent.
func (c *Ctx) UserAgent() string {
	return c.Request.UserAgent()
}

// Host returns the host.
func (c *Ctx) Host() string {
	return c.Request.Host
}

// RemoteAddr return client IP address.
func (c *Ctx) RemoteAddr() string {
	return c.Request.RemoteAddr
}

// IsFromLocal returns true if request is from localhost.
func (c *Ctx) IsFromLocal() bool {
	return c.Request.RemoteAddr == ""
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

// GetReqHeaders returns all request headers.
func (c *Ctx) GetReqHeaders() map[string][]string {
	return c.Request.Header
}

func (c *Ctx) CurrentRoute() string {
	return c.Request.URL.Path
}
