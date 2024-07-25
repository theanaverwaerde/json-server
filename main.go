package main

import (
	"encoding/json"
	"flag"
	"strconv"

	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var helpFlag = flag.Bool("help", false, "Show this")
var portFlag = flag.Int("port", 8080, "Which port to listen on")
var pageFlag = flag.String("page", "page", "Name of page parameter")
var limitFlag = flag.String("limit", "limit", "Name of limit parameter")

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("ERROR You need a json file at least.")
		fmt.Println("")
	}

	if *helpFlag || len(flag.Args()) == 0 {
		fmt.Println("With json-server you can start a REST API with just a json file.")
		fmt.Println("")
		fmt.Println("json-server [arguments] <file> [file...]")
		fmt.Println("")
		fmt.Println("The arguments are:")
		fmt.Println("")
		flag.PrintDefaults()
		return
	}

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.SetTrustedProxies(nil)

	for _, v := range flag.Args() {
		if !strings.HasSuffix(v, ".json") {
			fmt.Println("The file", v, "doesn't end by .json")
			os.Exit(1)
			return
		}

		content, err := os.ReadFile(v)
		if err != nil {
			fmt.Println("Error when opening file: ", err)
			os.Exit(1)
			return
		}

		var payload map[string]interface{}
		err = json.Unmarshal(content, &payload)
		if err != nil {
			fmt.Println("Error during Unmarshal(): ", err)
			return
		}

		for key, value := range payload {
			fmt.Println("GET", fmt.Sprintf("http://localhost:%v/%v", *portFlag, key))

			r.GET(key, making(value))
		}
	}

	r.Run(fmt.Sprintf(":%v", *portFlag))
}

func making(value interface{}) func(c *gin.Context) {

	var havePagination bool = false
	var baseArray []interface{}
	var arrayPath string
	var valueMap map[string]interface{}

	if isArrayOrSlice(value) {
		havePagination = true
		baseArray = value.([]interface{})
	} else {
		valueMap = value.(map[string]interface{})

		for key, value := range valueMap {
			if isArrayOrSlice(value) {
				havePagination = true
				baseArray = value.([]interface{})
				arrayPath = key
				break
			}
		}
	}

	return func(c *gin.Context) {
		var result interface{} = value
		if havePagination {
			limit, exist := QueryToInt(c, *limitFlag)
			if exist {
				page, exist := QueryToInt(c, *pageFlag)
				if exist {
					start := min((page - 1) * limit, len(baseArray))
					end := min(start + limit, len(baseArray))
					result = baseArray[start:end]

					if len(arrayPath) > 0 {
						valueMap[arrayPath] = result
						result = valueMap
					}
				}
			}
		}

		c.JSON(http.StatusOK, result)
	}
}

func isArrayOrSlice(x interface{}) bool {
	switch x.(type) {
	case []interface{}:
		return true
	default:
		return false
	}
}

func QueryToInt(c *gin.Context, key string) (int, bool) {
	if !c.Request.URL.Query().Has(key) {
		return 0, false
	}

	query := c.Query(key)
	value, err := strconv.Atoi(query)
	if err != nil {
		return 0, false
	}

	return value, true
}
