package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func getRandom(length int, wg *sync.WaitGroup, responses chan Response) int {
	var netClient = &http.Client{}

	requestUrl := "https://www.random.org/integers/?num=%d&min=1&max=100&col=1&base=10&format=plain&rnd=new"
	resp, err := netClient.Get(fmt.Sprintf(requestUrl, length))

	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("http request returned code %d\n", resp.StatusCode)
	}

	responses <- handleResponse(resp) // Return error if there is one, nil if not.
	defer wg.Done()
	return 1
}

func handleResponse(resp *http.Response) Response {
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	split := strings.Split(string(body), "\n")
	t2 := []int{}

	for i := 0; i < len(split)-1; i++ {
		j, _ := strconv.Atoi(split[i])
		t2 = append(t2, j)
	}
	log.Println(split)
	return Response{getStdDev(t2), t2}
}

func serverRequest(w http.ResponseWriter, r *http.Request) {
	var simpleList = []Response{}

	requestParam, ok := r.URL.Query()["requests"]
	if !ok || len(requestParam[0]) < 1 {
		log.Println("Problem with URL param 'rquests'")
		return
	}

	lengthParam, ok := r.URL.Query()["length"]
	if !ok || len(lengthParam[0]) < 1 {
		log.Println("Problem with URL 'length' param")
		return
	}

	requestNumber, _ := strconv.Atoi(requestParam[0])
	length, _ := strconv.Atoi(lengthParam[0])
	w.Header().Set("Content-Type", "application/json")
	var wg sync.WaitGroup
	responses := make(chan Response, length)
	wg.Add(requestNumber)
	n := 0
	for n < requestNumber {
		go getRandom(length, &wg, responses)
		n++
	}
	wg.Wait()
	n = 0
	close(responses)
	globalDataList := []int{}
	for n <= len(responses) {
		requestDataList := <-responses
		globalDataList = append(globalDataList, requestDataList.NumList...)
		simpleList = append(simpleList, requestDataList)
		n++
	}
	simpleList = append(simpleList, Response{getStdDev(globalDataList), globalDataList})

	json.NewEncoder(w).Encode(simpleList)

}

func main() {

	http.HandleFunc("/random/mean", serverRequest)
	http.ListenAndServe(":8080", nil)
}
