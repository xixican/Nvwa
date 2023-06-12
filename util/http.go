package util

import (
	"Nvwa/logger"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func HttpGet(url string, authorization string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.NvwaLog.Errorf("创建get request错误:%s, url=%s", err.Error(), url)
		return nil
	}
	req.Header.Set("Content-Type", "application/json")
	if authorization != "" {
		req.Header.Set("Authorization", authorization)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.NvwaLog.Errorf("调用get错误:%s, url=%s", err.Error(), url)
		return nil
	}
	b, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		logger.NvwaLog.Errorf("调用get错误:%s, url=%s", err.Error(), url)
		return nil
	}
	return b
}

func HttpPost(postURl string, authorization string, body io.Reader, proxy bool) []byte {
	req, err := http.NewRequest("POST", postURl, body)
	if err != nil {
		logger.NvwaLog.Errorf("创建post request错误:%s, postURl=%s", err.Error(), postURl)
		return nil
	}
	req.Header.Set("Content-Type", "application/json")
	if authorization != "" {
		req.Header.Set("Authorization", authorization)
	}
	client := http.DefaultClient
	if proxy {
		proxyUrl, e := url.Parse("http://localhost:7890")
		if e != nil {
			logger.NvwaLog.Errorf("调用post错误:%s, postURl=%s", e.Error(), postURl)
			return nil
		}
		transport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
		client = &http.Client{Transport: transport}
	}
	res, err := client.Do(req)
	if err != nil {
		logger.NvwaLog.Errorf("调用post错误:%s, postURl=%s", err.Error(), postURl)
		return nil
	}
	b, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		logger.NvwaLog.Errorf("调用post错误:%s, postURl=%s", err.Error(), postURl)
		return nil
	}
	if res.StatusCode != http.StatusOK {
		logger.NvwaLog.Errorf("调用post错误:%s, postURl=%s", string(b), postURl)
		return nil
	}
	return b
}

func HttpPostSetHeader(postURl string, header map[string]string, body io.Reader, proxy bool) []byte {
	req, err := http.NewRequest("POST", postURl, body)
	if err != nil {
		logger.NvwaLog.Errorf("创建post request错误:%s, postURl=%s", err.Error(), postURl)
		return nil
	}
	req.Header.Set("Content-Type", "application/json")
	for k, v := range header {
		req.Header.Set(k, v)
	}
	//if authorization != "" {
	//	req.Header.Set("Authorization", authorization)
	//}
	client := http.DefaultClient
	if proxy {
		proxyUrl, e := url.Parse("http://localhost:7890")
		if e != nil {
			logger.NvwaLog.Errorf("调用post错误:%s, postURl=%s", e.Error(), postURl)
			return nil
		}
		transport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
		client = &http.Client{Transport: transport}
	}
	res, err := client.Do(req)
	if err != nil {
		logger.NvwaLog.Errorf("调用post错误:%s, postURl=%s", err.Error(), postURl)
		return nil
	}
	b, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		logger.NvwaLog.Errorf("调用post错误:%s, postURl=%s", err.Error(), postURl)
		return nil
	}
	if res.StatusCode != http.StatusOK {
		logger.NvwaLog.Errorf("调用post错误:%s, postURl=%s", string(b), postURl)
		return nil
	}
	return b
}

func HttpPut(url string, authorization string, body io.Reader) []byte {
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		logger.NvwaLog.Errorf("创建put request错误:%s, url=%s", err.Error(), url)
		return nil
	}
	req.Header.Set("Content-Type", "application/json")
	if authorization != "" {
		req.Header.Set("Authorization", authorization)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.NvwaLog.Errorf("调用put错误:%s, url=%s", err.Error(), url)
		return nil
	}
	b, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		logger.NvwaLog.Errorf("调用put错误:%s, url=%s", err.Error(), url)
		return nil
	}
	return b
}
