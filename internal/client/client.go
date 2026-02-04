package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	*http.Client
	base      string
	accessKey string
	secretKey []byte
}

func New(base, accessKey string, secretKey []byte) *Client {
	c := &http.Client{}
	client := &Client{
		Client:    c,
		base:      base,
		accessKey: accessKey,
		secretKey: secretKey,
	}

	return client
}

func (c *Client) Do(r *http.Request) (*http.Response, error) {
	date := time.Now().Format(time.RFC1123)

	r.Header.Set("Date", date)

	signature, err := generateSignature(r, c.secretKey)
	if err != nil {
		return nil, err
	}

	encodedSignature := base64.StdEncoding.EncodeToString(signature)

	r.Header.Add("Authorization", fmt.Sprintf("%s:%s", c.accessKey, encodedSignature))

	return c.Client.Do(r)
}

func generateSignature(r *http.Request, key []byte) ([]byte, error) {
	// copy the body into somewhere that we can reset
	body := bytes.NewBuffer([]byte{})
	_, err := io.Copy(body, r.Body)
	if err != nil {
		return nil, err
	}

	// close the original body so we don't leak it
	err = r.Body.Close()
	if err != nil {
		return nil, err
	}

	// shasum the body
	requestBodySHA256 := sha256.New()
	_, err = io.Copy(requestBodySHA256, r.Body)
	if err != nil {
		return nil, err
	}

	// reset the body so it can be read again
	body.Reset()
	r.Body = io.NopCloser(body)

	m := new(bytes.Buffer)
	m.WriteString(r.Method)
	m.WriteString("\n")
	m.WriteString(r.URL.RequestURI())
	m.WriteString("\n")
	m.WriteString(r.Header.Get("Date"))
	m.WriteString("\n")
	m.Write(requestBodySHA256.Sum(nil))

	mac := hmac.New(sha256.New, key)
	mac.Write(m.Bytes())
	return mac.Sum(nil), nil
}
