package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	var filename string
	filename = "/Users/prabhat.ranjan/Downloads/freshers_raw.json"

	//fread json file---------
	inputs, err := ioutil.ReadFile(filename)
	check(err)

	//fmt.Println(string(inputs))

	datas := []map[string]interface{}{}

	json.Unmarshal(inputs, &datas)

	// fmt.Println(datas)

	for _, x := range datas {
		for key, val := range x {
			if key == "labels" && val == "" {
				delete(x, "labels")
			} else if key == "indexTimeEpoch" {
				//fmt.Println(val)
				//fmt.Println(reflect.TypeOf(val).String())
				val, ok := val.(string)
				if ok {
					i, err := strconv.ParseInt(val, 10, 64)
					check(err)
					i = i / 1000
					//fmt.Println(i)
					time_in_rfc := time.Unix(i, 0).Format(time.RFC3339)
					//fmt.Println(time_in_rfc)
					//fmt.Println(reflect.TypeOf(val).String())
					x["rfc3339time"] = time_in_rfc
				}
			} else if key == "ipmap" {
				//fmt.Println(val)
				//fmt.Println(reflect.TypeOf(val).String())
				val := val.(map[string]interface{})
				var ip []string

				for _, z := range val {
					//fmt.Println(z)
					z, ok := z.(string)
					if ok {
						z = strings.ReplaceAll(z, "from", "")
						z = strings.TrimSpace(z)
						// fmt.Println(z)
						ip = append(ip, z)

					}
					sort.Strings(ip)
					//
					//fmt.Println(reflect.TypeOf(z).String())
				}
				delete(x, "ipmap")
				//fmt.Println(ip)
				x["iparr"] = ip
			} else if key == "RawLogFilePath" {
				val, ok := val.(string)

				if ok {
					ss := strings.Split(val, "/")
					item := ss[len(ss)-1]
					//fmt.Println(item)
					x["filename"] = item
				}
			}
		}
	}
	//fmt.Println(datas)

	new_json, err := json.MarshalIndent(datas, "", "  ")
	check(err)

	fmt.Println(string(new_json))

	err = ioutil.WriteFile("output.json", new_json, 0777)
	check(err)
}
