# gontenttype
1. **Detect the content type of a given string**

   1. JSON (`application/json`)
   2. JSON Lines (`application/jsonl`)
   3. XML (`application/xml`)
   4. CSV (`text/csv`)
   5. Parquet (`application/vnd.apache.parquet`)

2. **Validate syntax for a given string and supported content types**


## Detect content type

### Example

#### Usage

```go
package main

import (
	"fmt"
    
	"github.com/costinmrr/gontenttype"
)

func main() { 
	// json
	myStr := `{"foo":"bar"}`
	contentType := gontenttype.Detect(myStr)
	fmt.Println(contentType) // application/json
	
	// jsonl
    myStr = `{"foo":"bar"}\n{"foo":"baz"}`
	contentType = gontenttype.Detect(myStr)
	fmt.Println(contentType) // application/jsonl

	// xml
	myStr = `<foo>bar</foo>`
	contentType = gontenttype.Detect(myStr)
	fmt.Println(contentType) // application/xml

	// csv
	myStr = `foo,bar`
	contentType = gontenttype.Detect(myStr)
	fmt.Println(contentType) // text/csv
	
	// parquet
	myStr = `PAR1...[parquet content]...PAR1` // use a valid parquet content
    contentType = gontenttype.Detect(myStr)
	fmt.Println(contentType) // application/vnd.apache.parquet
}
```

#### Output
    
```shell
application/json
application/jsonl
application/xml
text/csv
application/vnd.apache.parquet
```

## Validate syntax

### Example

```go
package main

import (
	"fmt"
	
	"github.com/costinmrr/gontenttype/types/json"
	"github.com/costinmrr/gontenttype/types/jsonlines"
	"github.com/costinmrr/gontenttype/types/xml"
	"github.com/costinmrr/gontenttype/types/csv"
	"github.com/costinmrr/gontenttype/types/parquet"
)

func main() {
	// json
	myStr := `{"foo":"bar"}`
	err := json.IsJSON(myStr)
	fmt.Println(err) // <nil>

	myStr = `{"foo":"bar"`
	err = json.IsJSON(myStr)
	fmt.Println(err) // unexpected end of JSON input
	
	// jsonl
	myStr = "{\"foo\":\"bar\"}\n{\"foo\":\"baz\"}"
	err = jsonlines.IsJSONLines(myStr)
	fmt.Println(err) // <nil>
	
	myStr = "{\"foo\":\"bar\"}\n{\"foo\":\"baz\""
	err = jsonlines.IsJSONLines(myStr)
	fmt.Println(err) // error on line 2: unexpected end of JSON input

	// xml
	myStr = `<foo>bar</foo>`
	err = xml.IsXML(myStr)
	fmt.Println(err) // <nil>

	myStr = `<foo>bar</foo`
	err = xml.IsXML(myStr)
	fmt.Println(err) // XML syntax error on line 1: unexpected EOF

	// csv
	myStr = `foo,bar`
	err = csv.IsCSV(myStr)
	fmt.Println(err) // <nil>

	myStr = "col1,col2\nfoo,bar,baz"
	err = csv.IsCSV(myStr)
	fmt.Println(err) // record on line 2: wrong number of fields
	
	// parquet
	myStr = `PAR1...[parquet content]...PAR1` // use a valid parquet content
    err = parquet.IsParquet(myStr)
	fmt.Println(err) // <nil>
	
	myStr = `PAR1invalidPAR1`
	err = parquet.IsParquet(myStr)
	fmt.Println(err) // invalid parquet file: reading footer of parquet file: strings.Reader.ReadAt: negative offset
}
```

#### Output

```shell
<nil>
unexpected end of JSON input
<nil>
error on line 2: unexpected end of JSON input
<nil>
XML syntax error on line 1: unexpected EOF
<nil>
record on line 2: wrong number of fields
<nil>
invalid parquet file: reading footer of parquet file: strings.Reader.ReadAt: negative offset
```


## Benchmarks

```shell
go test -bench=.
```

#### Output

```shell
goos: darwin
goarch: arm64
pkg: github.com/costinmrr/gontenttype
cpu: Apple M1 Pro
BenchmarkDetectJSON_SimpleString-8               3261396               362.7 ns/op
BenchmarkDetectJSON_1KB-8                         125338              9585 ns/op
BenchmarkDetectJSON_100KB-8                         1310            909954 ns/op
BenchmarkDetectJSON_1MB-8                            153           7776861 ns/op
BenchmarkDetectJSON_10MB-8                            14          77368839 ns/op
BenchmarkDetectJSONLines_SimpleString-8           764656              1426 ns/op
BenchmarkDetectJSONLines_1KB-8                  12037003                99.02 ns/op
BenchmarkDetectJSONLines_100KB-8                12053284                98.60 ns/op
BenchmarkDetectJSONLines_1MB-8                  12237555                98.79 ns/op
BenchmarkDetectJSONLines_10MB-8                 12056943                98.68 ns/op
BenchmarkDetectXML_SimpleString-8                1545378               775.3 ns/op
BenchmarkDetectXML_1KB-8                           45462             26035 ns/op
BenchmarkDetectXML_100KB-8                           511           2334531 ns/op
BenchmarkDetectXML_1MB-8                              48          24450518 ns/op
BenchmarkDetectXML_10MB-8                              5         248890150 ns/op
BenchmarkDetectCSV_SimpleString-8                 679507              1603 ns/op
BenchmarkDetectCSV_1KB-8                           56437             20994 ns/op
BenchmarkDetectCSV_100KB-8                           648           1887705 ns/op
BenchmarkDetectCSV_1MB-8                              58          20405615 ns/op
BenchmarkDetectCSV_10MB-8                              6         194822194 ns/op
BenchmarkDetectUnsupported_SimpleString-8         748053              1501 ns/op
```

```shell
goos: linux
goarch: amd64
pkg: github.com/costinmrr/gontenttype
cpu: AMD EPYC 7763 64-Core Processor                
BenchmarkDetectJSON_SimpleString-4          	 2135098	       580.3 ns/op
BenchmarkDetectJSON_1KB-4                   	   89289	     13288 ns/op
BenchmarkDetectJSON_100KB-4                 	     901	   1303257 ns/op
BenchmarkDetectJSON_1MB-4                   	     100	  10029120 ns/op
BenchmarkDetectJSON_10MB-4                  	       9	 111325401 ns/op
BenchmarkDetectJSONLines_SimpleString-4     	  503854	      2244 ns/op
BenchmarkDetectJSONLines_1KB-4              	 9440169	       127.4 ns/op
BenchmarkDetectJSONLines_100KB-4            	 9543016	       129.8 ns/op
BenchmarkDetectJSONLines_1MB-4              	 9292095	       126.8 ns/op
BenchmarkDetectJSONLines_10MB-4             	 9430472	       127.0 ns/op
BenchmarkDetectXML_SimpleString-4           	  465271	      2395 ns/op
BenchmarkDetectXML_1KB-4                    	   30892	     38756 ns/op
BenchmarkDetectXML_100KB-4                  	     348	   3417666 ns/op
BenchmarkDetectXML_1MB-4                    	      31	  36768652 ns/op
BenchmarkDetectXML_10MB-4                   	       3	 343983510 ns/op
BenchmarkDetectCSV_SimpleString-4           	  363416	      3219 ns/op
BenchmarkDetectCSV_1KB-4                    	   42157	     29058 ns/op
BenchmarkDetectCSV_100KB-4                  	     447	   2558199 ns/op
BenchmarkDetectCSV_1MB-4                    	      30	  34082715 ns/op
BenchmarkDetectCSV_10MB-4                   	       4	 302733827 ns/op
BenchmarkDetectUnsupported_SimpleString-4   	  331250	      3157 ns/op
```