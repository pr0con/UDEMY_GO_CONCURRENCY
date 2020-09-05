//https://mholt.github.io/json-to-go/ add a keyword yaml
package main

import(
	"fmt"
	"io/ioutil"
	
	"gopkg.in/yaml.v3"
)

type Service struct {
    APIVersion string `yaml:"apiVersion"`
    Kind       string `yaml:"kind"`
    Metadata   struct {
        Name      string `yaml:"name"`
        Namespace string `yaml:"namespace"`
        Labels    struct {
            RouterDeisIoRoutable string `yaml:"router.what.io/routable"`
        } `yaml:"labels"`
        Annotations struct {
            RouterDeisIoDomains string `yaml:"router.what.io/domains"`
        } `yaml:"annotations"`
    } `yaml:"metadata"`
    Spec struct {
        Type     string `yaml:"type"`
        Selector struct {
            App string `yaml:"app"`
        } `yaml:"selector"`
        Ports []struct {
            Name       string `yaml:"name"`
            Port       int    `yaml:"port"`
            TargetPort int    `yaml:"targetPort"`
            NodePort   int    `yaml:"nodePort,omitempty"`
        } `yaml:"ports"`
    } `yaml:"spec"`
}

func main() {
	var service Service
	
	yam, err := ioutil.ReadFile("/var/www/go_systems/fake_yaml.yaml");
	if err != nil { fmt.Println("something went horribly wrong 1") } else {
		err := yaml.Unmarshal(yam, &service)
		if err != nil  {
			fmt.Println("something went horribly wrong 2");
		}
		
		fmt.Println(service.Metadata.Name)
	}
			
}


