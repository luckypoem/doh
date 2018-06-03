package main

import (
	"net"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func ProcessRequest(n int, addr *net.UDPAddr, buf []byte, ServerConn *net.UDPConn) {
	record, query := ParseRecordandOriginalQuery(buf) //parse request into record to request and original byte string
	resolved_query := DNSOverHTTPSRequest(record) //request DNS record over HTTPS
	if  resolved_query != nil {
		response := CreateResponse(resolved_query, buf, query)
		_, err := ServerConn.WriteToUDP(response, addr)
		CheckError(err)
		//TODO implement type 
		//TODO catch empty response (failed)
	}
}

func DNSOverHTTPSRequest(record string) map[string]interface{} {
	//QUERY OVER HTTPS
	res, err := http.Get("https://1.1.1.1/dns-query?ct=application/dns-json&name="+record+"&type=A")
	CheckError(err)
	body, err := ioutil.ReadAll(res.Body)
	CheckError(err)

	//decode JSON response in body
	var f interface{}
	json.Unmarshal(body, &f)
	m := f.(map[string]interface{}) //make a mappable opbject

	//TODO DEBUG HERE failed query (e.g.)
	//TODO catch empty response (failed)
	if m["Answer"] != nil {
		resolved_query_slice := m["Answer"].([]interface{})
		resolved_query := resolved_query_slice[len(resolved_query_slice)-1].(map[string]interface{})// contains TTL name data type
		return resolved_query
	} else {
		return nil
	}
}

func ParseRecordandOriginalQuery(buf []byte) (string, []byte) {
	query := []byte{}
	for _, element := range buf[12:] { //includes starting symbol
		if element == 0 {
			break
		} else {
			query = append(query, element)
		}
	}

	record := ""
	for _, element := range query[1:] {
		if element < 31 { //discover dots 
			record+="."
		} else {
			record+=string(element)
		}
	}

	return record,query
}
