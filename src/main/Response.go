package main

import (
	"strconv"
	"strings"
)

func CreateResponse(resolved_query map[string]interface{}, buf []byte, query []byte) []byte {
	response := buf[:2] //id
	response = append(response, []byte{129, 128}...) //flags
	response = append(response, []byte{0, 1, 0, 1, 0, 0, 0, 0}...) //rr
	response = append(response, query...) //query
	response = append(response, 0) //query end
	response = append(response, []byte{0, 1, 0, 1}...) //type and class
	response = append(response, []byte{192, 12}...) //first record
	response = append(response, []byte{0, 1, 0, 1}...) //type class first record
	response = append(response, ParseTTLInBytes(resolved_query)...)
	response = append(response, ParseDataInBytes(resolved_query)...)
	return response
}
func ParseTTLInBytes(resolved_query map[string]interface{}) []byte {
	ttl := strconv.FormatInt(int64(resolved_query["TTL"].(float64)), 2) //make binary 
	ttl = strings.Repeat("0", 32-len(ttl)) + ttl //prepend 0's 
	TTLInBytes := []byte{}
	for _, element := range []int{0,8,16,24} { //split complete 32-bit in byte-sized chunks
		i, _ :=strconv.ParseInt(ttl[element:element+8],2,64) //derive value from binary string
		TTLInBytes = append(TTLInBytes, byte(i)) //add to response
	}
	return TTLInBytes
}

func ParseDataInBytes(resolved_query map[string]interface{}) []byte {
	DataInBytes := []byte{}
	DataInBytes = append(DataInBytes, []byte{0, 4}...) //TODO temp harcode data length (4)
	for _, element := range strings.Split(resolved_query["data"].(string), ".") { //every part of ip to decimal
		i, _ := strconv.Atoi(element) //convert element to int
		DataInBytes = append(DataInBytes, byte(i)) //add to response
	}
	return DataInBytes
}
