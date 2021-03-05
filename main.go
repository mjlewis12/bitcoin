// functionality
// * Rate prediction
// * Print historical Rates
// * Print current rate
// * Historical analytics

package main

import "bitcoin/cmd"

func main() {
	cmd.Execute()
}

// type Rate struct {
//   Code string `json:"code"`
//   Rate float64 `json:"rate_float"`
// }
//
// type Rates struct {
//   USD Rate `json:"USD"`
// }
//
// type Timestamp struct {
//   Updated time.Time `json:"updatedISO"`
// }
//
// type Response struct {
//   Time Timestamp `json:"time"`
//   Chartname string `json:"chartName"`
//   Rates Rates `json:"bpi"`
// }
//
// func main() {
//   resp, _ := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
//   defer resp.Body.Close()
//
//   b, _ := ioutil.ReadAll(resp.Body)
//   var res Response
//   json.Unmarshal(b,&res)
//
//   fmt.Println(res.Time,res.Rates.USD.Rate)
//   //t, _ := time.Parse(time.RFC3339,res.Time.Updated)
//   //fmt.Println(t)
//   //.Rates.USD.Rate
// }
