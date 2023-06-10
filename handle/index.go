package handle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	as "github.com/aerospike/aerospike-client-go"
	gin "github.com/gin-gonic/gin"
)

const Namespace = "test"
const Set = "Person"

type Test struct {
	Id    string `json:"id"`
	Price string `json:"price"`
}

func Insert(client *as.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		keyParams := c.Query("key")
		key, err := as.NewKey(Namespace, Set, keyParams)
		if err != nil {
			log.Fatal(err)
		}
		requestbody, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			log.Fatal(err)
		}
		//var body1 Test
		var inInterface map[string]interface{}
		err1 := json.Unmarshal(requestbody, &inInterface)
		if err1 != nil {
			log.Fatal(err1)
		}
		// binAge := as.NewBin("age", 26)
		// binName := as.NewBin("name", "ittest")

		// binMap := as.BinMap{
		// 	"age":  25,
		// 	"name": "Quoc dep trai",
		// }
		//reportMap := map[string]interface{}{
		// "city":     "Ann Arbor",
		// "state":    "Michigan",
		// "shape":    "test",
		// "duration": "5 minutes",
		// "summary":  "Large flying disc flashed in the sky above the student union. Craziest thing I've ever seen!"}
		// report := as.NewBin("report", reportMap)
		// client.PutBins(nil, key, report)
		// client.PutBins(nil, key, binAge, binName, report)
		client.Put(nil, key, inInterface)
		// record, err := client.Get(nil, key)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": inInterface,
		})
	}
}

func Get(client *as.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		keyParams := c.Query("key")
		key, err := as.NewKey(Namespace, Set, keyParams)
		if err != nil {
			log.Fatal(err)
		}
		record, err := client.Get(nil, key)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("record", record)

		c.JSON(http.StatusOK, record)
	}
}

func Delete(client *as.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		key, err := as.NewKey(Namespace, Set, "ittest")
		if err != nil {
			log.Fatal(err)
		}
		// Delete record with default write policy
		existed, err := client.Delete(nil, key)

		// Close the connection to the server
		// client.Close()

		c.JSON(http.StatusOK, gin.H{
			"message": existed,
		})
	}
}

func Statement(client *as.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		stmt := as.NewStatement(Namespace, Set)
		//exp := as.ExpRegexCompare("*", as.ExpRegexFlagICASE, as.ExpKey(as.ExpTypeSTRING))

		qp := as.NewQueryPolicy()
		// qp.FilterExpression = exp
		rs, err := client.Query(qp, stmt)
		if err != nil {
			log.Fatalln(err.Error())
		}
		list := make(map[string]interface{})
		index := "0"
		for res := range rs.Results() {
			// if res.Err != nil {
			// 	log.Fatalln(err.Error())
			// }
			list[index] = res.Record
			index = index + "0"

			// if !strings.Contains(strings.ToLower(res.Record.Key.Value().GetObject().(string)), "apple") {
			// 	log.Fatalf("Wrong key returned: %s. Expected to include 'apple' (case-insensitive)\n", res.Record.Key.Value())
			// }
		}
		log.Println("Finished successafully.")
		c.JSON(http.StatusOK, list)
	}
}

func StatementFilerBinName(client *as.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		//binNames := []string{"id"}
		stmt := as.NewStatement(Namespace, Set, "id")

		//exp := as.ExpRegexCompare("*", as.ExpRegexFlagICASE, as.ExpKey(as.ExpTypeSTRING))
		// stmt.Filter(as.NewBinsFilter(binNames...))
		qp := as.NewQueryPolicy()
		// qp.FilterExpression = exp
		rs, err := client.Query(qp, stmt)
		if err != nil {
			log.Fatalln(err.Error())
		}
		list := make(map[string]interface{})
		index := "0"
		for res := range rs.Results() {
			// if res.Err != nil {
			// 	log.Fatalln(err.Error())
			// }
			list[index] = res.Record
			index = index + "0"

			// if !strings.Contains(strings.ToLower(res.Record.Key.Value().GetObject().(string)), "apple") {
			// 	log.Fatalf("Wrong key returned: %s. Expected to include 'apple' (case-insensitive)\n", res.Record.Key.Value())
			// }
		}
		log.Println("Finished successafully.")
		c.JSON(http.StatusOK, list)
	}

}
func StatementLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello world")
	}
}
