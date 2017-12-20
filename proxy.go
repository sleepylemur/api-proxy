package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Service struct {
	path     string `yaml:"path"`
	endpoint string `yaml:"endpoint"`
}

func main() {
	config := make(map[string]Service)
	configdata, err := ioutil.ReadFile("services.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(configdata, config)
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
	fmt.Println(config["bob"])

	services := map[string]Service{
		"bob":  Service{endpoint: "http://localhost:3000"},
		"dave": Service{endpoint: "http://localhost:3001"},
	}
	proxies := make(map[string]*httputil.ReverseProxy)
	for name, service := range services {
		url, _ := url.Parse(service.endpoint)
		proxies[name] = httputil.NewSingleHostReverseProxy(url)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		for name, _ := range services {
			if strings.HasPrefix(path, "/"+name) {
				proxies[name].ServeHTTP(w, r)
				return
			}
		}
		fmt.Fprintf(w, "Hello, %v\n", r.URL.Path)
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
