package scaffold

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
}

func NewContext(w http.ResponseWriter, r *http.Request) (context *Context) {
	context = &Context{Response: w, Request: r}
	context.UseCORS()

	return
}

func ContextAdapter(handleFunc func(context *Context)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) { handleFunc(NewContext(w, r)) }
}

func (c *Context) UseCORS() {
	(c.Response).Header().Set("Access-Control-Allow-Headers", "*")
	(c.Response).Header().Set("Access-Control-Allow-Methods", "*")

	if origin := c.Request.Header.Get("Origin"); origin == "" {
		(c.Response).Header().Set("Access-Control-Allow-Origin", "*")
	} else {
		(c.Response).Header().Set("Access-Control-Allow-Origin", origin)
		(c.Response).Header().Set("Access-Control-Allow-Credentials", "true")
	}
}

func (c *Context) pureRender(statusCode int, content interface{}) {
	c.Response.WriteHeader(statusCode)
	fmt.Fprintf(c.Response, "%s", content)
}

func (c *Context) RenderHTML(statusCode int, html string) {
	c.Response.Header().Set("Content-Type", "text/html")
	c.pureRender(statusCode, html)
}

func (c *Context) RenderText(statusCode int, text string) {
	c.Response.Header().Set("Content-Type", "plain/text")
	c.pureRender(statusCode, text)
}

func (c *Context) Render(statusCode int, src interface{}) {
	c.Response.Header().Set("Content-Type", "application/json")

	var out bytes.Buffer
	responseCode := statusCode

	enc := json.NewEncoder(&out)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)

	if err := enc.Encode(src); err != nil {
		bug := ResponseError{
			Error:            "json_encode_error",
			ErrorDescription: err.Error(),
		}
		msg, _ := json.Marshal(bug)
		out = *bytes.NewBuffer(msg)
		responseCode = http.StatusInternalServerError
	}

	c.pureRender(responseCode, out.String())
}

func (c *Context) Headers() map[string]string {
	headers := make(map[string]string)
	for key, values := range c.Response.Header() {
		headers[key] = values[0]
	}
	return headers
}

func (c *Context) Header(header string) string {
	return c.Response.Header().Get(header)
}

func (c *Context) Params() map[string]string {
	return mux.Vars(c.Request)
}

func (c *Context) Param(param string) (content string) {
	queries := c.Params()
	content = queries[param]
	return
}

func (c *Context) UnmarshalBody(dest interface{}) error {
	payload, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}

	if payload != nil && string(payload) != "" {
		if err := json.Unmarshal(payload, dest); err != nil {
			return err
		}
	}

	return nil
}
