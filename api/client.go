package api

import (
    "encoding/json"
    "io"
    "io/ioutil"
    "net/http"
    "net/http/cookiejar"
    "net/url"
    "strings"
    "time"
)

type Client struct {
    baseUrl *url.URL
    userAgent string
    httpClient *http.Client
    ticket string
    csrfToken string
}

func NewClient(baseUrl string, username string, password string) (*Client, error) {
    c := &Client{
        userAgent: "proxmox-go-api",
        httpClient: &http.Client{
            Timeout: time.Second * 5,
        },
    }
    cookieJar, err := cookiejar.New(nil)
    if err != nil {
        return nil, err
    }
    c.httpClient.Jar = cookieJar
    u, err := url.Parse(baseUrl)
    if err != nil {
        return nil, err
    }
    c.baseUrl = u

    err = c.Login(username, password)
    if err != nil {
        return nil, err
    }

    c.httpClient.Jar.SetCookies(c.baseUrl, []*http.Cookie{
        {Name: "PVEAuthCookie", Value: c.ticket},
    })
    return c, nil
}

func (c *Client) newRequest(method, path string, body map[string]string) (*http.Request, error) {
    rel := &url.URL{Path: c.baseUrl.Path + path}
    u := c.baseUrl.ResolveReference(rel)

    var buf io.Reader
    if body != nil {
        params := url.Values{}
        for k, v := range body {
            params.Add(k, v)
        }
        buf = strings.NewReader(params.Encode())
    }
    req, err := http.NewRequest(method, u.String(), buf)
    if err != nil {
        return nil, err
    }
    req.Header.Set("User-Agent", c.userAgent)
    return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, err
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    err = json.Unmarshal(body, v)
    return resp, err
}

func (c *Client) get(path string, params map[string]string, v interface{}) error {
    p := url.Values{}
    for k, v := range params {
        p.Add(k, v)
    }
    if len(p.Encode()) != 0 {
        path = path + "?" + p.Encode()
    }
    req, err := c.newRequest(http.MethodGet, path, nil)
    if err != nil {
        return err
    }
    wrapper := &wrapper{
        Data: v,
    }
    _, err = c.do(req, wrapper)
    if err != nil {
        return err
    }
    return nil
}

type wrapper struct {
    Errors string      `json:"errors,omitempty"`
    Data   interface{} `json:"data"`
}

