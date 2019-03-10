package main

import (
  //"fmt"
  "log"
  "io/ioutil"
  "gopkg.in/yaml.v2"
  "os/exec"
  "os"
)

var boolOptions = []string{"no-flannel","docker","disable-agent"}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}


func getConf(file string) map[string]string {
  conf := make(map[string]string)
  var yamlFile []byte
  yamlFile, err := ioutil.ReadFile(file)
  if err != nil {
      log.Printf("can't find the yaml file")
  }
  err = yaml.Unmarshal([]byte(yamlFile), &conf)
  if err != nil {
      log.Fatalf("check the yaml file format")
  }
  return conf
}

func constructCmd(mode string) exec.Cmd{
  args := []string{"/usr/local/bin/k3s",mode}
  for key,value := range getConf("/etc/k3s/"+mode+"-conf.yml") {
    if stringInSlice(key,boolOptions) {
      if value == "yes" {
        args = append(args,"--"+key)
      }
    } else if value != "" {
      args = append(args,"--"+key)
      args = append(args,value)
    }
  }
  cmd := exec.Cmd{}
  cmd.Path= "/usr/local/bin/k3s"
  cmd.Args = args
  return cmd
}

func main() {
  sleepcmd := exec.Command("sleep","10")
  err := sleepcmd.Start()
  if err != nil {
    log.Fatal(err)
  }
  k3scmd := constructCmd(os.Args[1])
  err = k3scmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}
