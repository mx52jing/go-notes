package main

import (
	"fmt"
	"net/url"
)

func parseFn() {
	//rawURL := "https://example.com/path?name=张三#section"
	rawURL := "https://example.com:8080/#/before/error?name=%E5%BC%A0%E4%B8%89"
	parsedURL, _ := url.Parse(rawURL)
	fmt.Println("parsedURL:", parsedURL)
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("RawQuery:", parsedURL.RawQuery)
	fmt.Println("Fragment:", parsedURL.Fragment)
	/**
	rawURL := "https://example.com/path?name=张三#section" 打印如下
	parsedURL: https://example.com/path?name=张三#section
	Scheme: https
	Host: example.com
	Path: /path
	RawQuery: name=张三
	Fragment: section
	*/

	/**
	rawURL := "https://example.com:8080/#/before/error?name=%E5%BC%A0%E4%B8%89" 打印如下
	parsedURL: https://example.com:8080/#/before/error?name=%E5%BC%A0%E4%B8%89
	Scheme: https
	Host: example.com:8080
	Path: /
	RawQuery:
	Fragment: /before/error?name=张三
	*/
}

func parseRequestFn() {
	//rawURL := "https://example.com/path?name=张三#section"
	rawURL := "https://example.com:8080/#/before/error?name=%E5%BC%A0%E4%B8%89"
	parsedURL, _ := url.ParseRequestURI(rawURL)
	fmt.Println("parsedURL:", parsedURL)
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("RawQuery:", parsedURL.RawQuery)
	fmt.Println("Fragment:", parsedURL.Fragment)
	/**
	rawURL := "https://example.com/path?name=张三#section" 打印如下
	parsedURL: https://example.com/path?name=张三#section
	Scheme: https
	Host: example.com
	Path: /path
	RawQuery: name=张三#section
	*/

	/**
	rawURL := "https://example.com:8080/#/before/error?name=%E5%BC%A0%E4%B8%89" 打印如下
	parsedURL: https://example.com:8080/%23/before/error?name=%E5%BC%A0%E4%B8%89
	Scheme: https
	Host: example.com:8080
	Path: /#/before/error
	RawQuery: name=%E5%BC%A0%E4%B8%89
	Fragment:
	*/
}

func queryEscapeAndUnEscape() {
	rawQuery := "?name=张三&age=18"
	escapedQuery := url.QueryEscape(rawQuery)
	fmt.Println("Escaped Query:", escapedQuery)
	unescapedQuery, _ := url.QueryUnescape(escapedQuery)
	fmt.Println("Unescaped Query:", unescapedQuery)
	/**
	Escaped Query: %3Fname%3D%E5%BC%A0%E4%B8%89%26age%3D18
	Unescaped Query: ?name=张三&age=18
	*/
}

func pathEscapeAndUnPathEscape() {
	rawUrl := "https://example.com/path?name=张三&age=22"
	escapedUrl := url.PathEscape(rawUrl)
	// Escaped Url: https:%2F%2Fexample.com%2Fpath%3Fname=%E5%BC%A0%E4%B8%89&age=22
	fmt.Println("Escaped Url:", escapedUrl)
	unescapedUrl, _ := url.PathUnescape(escapedUrl)
	// UnEscaped Url: https://example.com/path?name=张三&age=22
	fmt.Println("UnEscaped Url:", unescapedUrl)
}

func urlValues() {
	params := url.Values{}
	params.Add("name", "张三")
	params.Add("age", "22")
	query := params.Encode()
	fmt.Println("query:", query)
}

func buildURL() {
	parsedUrl, _ := url.Parse("https://example.com/path")
	// 添加search params
	params := url.Values{}
	params.Add("name", "张三")
	params.Add("age", "22")
	parsedUrl.RawQuery = params.Encode()
	finalURL := parsedUrl.String()
	fmt.Println("finalURL:", finalURL)
}

func parseQuery() {
	rawQuery := "name=张三&age=22"
	queryParams, _ := url.ParseQuery(rawQuery)
	queryParams.Set("like", "ping_pang")
	fmt.Println("Query Params:", queryParams)
	fmt.Println("name:", queryParams.Get("name"))
	fmt.Println("age:", queryParams.Get("age"))
	fmt.Println("like:", queryParams.Get("like"))
	/**
	打印如下:
	Query Params: map[age:[22] like:[ping_pang] name:[张三]]
	name: 张三
	age: 22
	like: ping_pang
	*/
}

func pathJoin() {
	baseURL := "https://example.com"
	path1 := "path1"
	path2 := "path2"
	joinedURL, _ := url.JoinPath(baseURL, path1, path2)
	// Joined URL: https://example.com/path1/path2
	fmt.Println("Joined URL:", joinedURL)
}

func urlString() {
	parsedURL := &url.URL{
		Scheme:   "https",
		Host:     "example.com",
		Path:     "/path",
		RawQuery: "name=张三",
		Fragment: "section",
	}
	// URL String: https://example.com/path?name=张三#section
	fmt.Println("URL String:", parsedURL.String())
}

func pathNameHost() {
	rawURL := "https://example.com:8080/path?foo=bar#section"
	parsedURL, _ := url.Parse(rawURL)
	/**
	Hostname: example.com
	Port: 8080
	*/
	fmt.Println("Hostname:", parsedURL.Hostname())
	fmt.Println("Port:", parsedURL.Port())
}

func urlQuery() {
	rawURL := "https://example.com/path?name=张三&age=22"
	parsedURL, _ := url.Parse(rawURL)
	queryParams := parsedURL.Query()
	/**
	Query Params: map[age:[22] name:[张三]]
	name: 张三
	age: 22
	like: ping_pang
	*/
	fmt.Println("Query Params:", queryParams)
	fmt.Println("name:", queryParams.Get("name"))
	queryParams.Set("like", "ping_pang")
	fmt.Println("age:", queryParams.Get("age"))
	fmt.Println("like:", queryParams.Get("like"))
}

func urlIsAbs() {
	rawURL1 := "https://example.com/path"
	parsedURL1, _ := url.Parse(rawURL1)

	rawURL2 := "example.com/path"
	parsedURL2, _ := url.Parse(rawURL2)
	/**
	RawURL1 Is absolute URL? true
	RawURL2 Is absolute URL? false
	*/
	fmt.Println("RawURL1 Is absolute URL?", parsedURL1.IsAbs())
	fmt.Println("RawURL2 Is absolute URL?", parsedURL2.IsAbs())
}

func urlUserInfo() {
	rawURL := "https://user:root@example.com/path"
	parsedURL, _ := url.Parse(rawURL)
	if parsedURL.User != nil {
		fmt.Println("Username:", parsedURL.User.Username())
		password, hasPassword := parsedURL.User.Password()
		if hasPassword {
			fmt.Println("Password:", password)
		} else {
			fmt.Println("No password provided")
		}
		return
	}
	fmt.Println("No user info provided")
	/**
	Username: user
	Password: root
	*/
}

func main() {
	//parseFn()
	//parseRequestFn()
	//pathEscapeAndUnPathEscape()
	//queryEscapeAndUnEscape()
	//urlValues()
	//buildURL()
	//parseQuery()
	//pathJoin()
	//urlString()
	//pathNameHost()
	//urlQuery()
	//urlIsAbs()
	urlUserInfo()
}
