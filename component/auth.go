// Code generated by AfterShip SDK Generator. DO NOT EDIT.

package component

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"github.com/AfterShip/tracking-sdk-go/v4/errorx"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

const (
	ApiKey = "API_KEY"
	Aes    = "AES"
	Rsa    = "RSA"
)

type Authenticator struct {
	apikey    string
	apiSecret string
	kind      string
}

func NewAuthenticator(apikey, apiSecret string, kind string) *Authenticator {
	return &Authenticator{
		apikey:    apikey,
		apiSecret: apiSecret,
		kind:      kind,
	}
}

func (au *Authenticator) Sign(r *http.Request) (map[string]string, error) {
	if len(au.apikey) == 0 {
		return nil, errorx.NewSdkError(errorx.ErrInvalidApiKey, errorx.GetErrorMessage(errorx.ErrInvalidApiKey), "")
	}
	headers := map[string]string{
		"as-api-key": au.apikey,
	}
	if au.kind != ApiKey {
		tmpHeader := r.Header
		tmpHeader.Add("as-api-key", au.apikey)
		h := au.canonicalHeader(tmpHeader)
		rs, err := au.canonicalResource(r.URL.RequestURI())
		if err != nil {
			return headers, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
		}
		var b bytes.Buffer
		if r.Body != nil {
			_, err := b.ReadFrom(r.Body)
			if err != nil {
				return headers, errorx.NewSdkError(errorx.ErrBadRequest, errorx.GetErrorMessage(errorx.ErrBadRequest), err.Error())
			}
			r.Body = io.NopCloser(&b)
		}

		s, err := au.signString(r.Method, b.String(), r.Header.Get("date"), h, rs)
		if err != nil {
			return headers, err
		}
		if au.kind == Aes {
			headers["as-signature-hmac-sha256"] = au.getHMACSignature(s, au.apiSecret)
		}
		if au.kind == Rsa {
			rsaSign, err := au.getRSASignature(s, au.apiSecret)
			if err != nil {
				return headers, errorx.NewSdkError(errorx.ErrInvalidApiKey, errorx.GetErrorMessage(errorx.ErrInvalidApiKey), err.Error())
			}
			headers["as-signature-rsa-sha256"] = rsaSign
		}
	}
	return headers, nil
}

func (au *Authenticator) signString(method, body, date, canonicalHeader, canonicalResource string) (string, error) {
	var (
		err         error
		result      string
		newBody     string
		contentType string
	)
	if body != "" {
		newBody, err = md5Encode(body)
		if err != nil {
			return "", err
		}
		contentType = "application/json"
	}
	result += method + "\n"
	result += newBody + "\n"
	result += contentType + "\n"
	result += date + "\n"
	result += canonicalHeader + "\n"
	result += canonicalResource
	return result, nil
}

func (au *Authenticator) canonicalHeader(headers http.Header) string {
	if headers == nil {
		return ""
	}
	keys := make([]string, 0, len(headers))
	newHeaders := make(map[string]string, len(headers))
	for key, value := range headers {
		newKey := strings.ToLower(key)
		if !strings.HasPrefix(newKey, "as-") {
			continue
		}
		keys = append(keys, newKey)
		newHeaders[newKey] = strings.TrimLeft(value[0], " ")
	}

	sort.Strings(keys)

	result := make([]string, 0, len(keys))
	for _, key := range keys {
		result = append(result, key+":"+newHeaders[key])
	}

	return strings.Join(result, "\n")
}

func (au *Authenticator) canonicalResource(rawUrl string) (result string, err error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	result += u.Path
	params := u.Query().Encode()
	if params != "" {
		result += "?" + params
	}
	return result, nil
}

func (au *Authenticator) getRSASignature(signString, apiSecret string) (string, error) {
	block, _ := pem.Decode([]byte(apiSecret))
	if block == nil {
		return "", errors.New("invalid api key")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	hash := sha256.New()
	hash.Write([]byte(signString))
	msgSum := hash.Sum(nil)
	b, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgSum, nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func (au *Authenticator) getHMACSignature(signString, apiSecret string) string {
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write([]byte(signString))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func md5Encode(source string) (hashed string, err error) {
	h := md5.New()
	_, err = h.Write([]byte(source))
	if err != nil {
		return
	}
	sum := h.Sum(nil)
	hashed = strings.ToUpper(hex.EncodeToString(sum))
	return
}