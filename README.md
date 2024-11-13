# gontenttype
1. **Detect the content type of a given string**

   1. JSON (`application/json`)
   2. XML (`application/xml`)
   3. CSV (`text/csv`)

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

	// xml
	myStr = `<foo>bar</foo>`
	contentType = gontenttype.Detect(myStr)
	fmt.Println(contentType) // application/xml

	// csv
	myStr = `foo,bar`
	contentType = gontenttype.Detect(myStr)
	fmt.Println(contentType) // text/csv
}
```

#### Output
    
```shell
application/json
application/xml
text/csv
```

## Validate syntax

### Example

```go
package main

import (
	"fmt"
	
	"github.com/costinmrr/gontenttype/types/json"
	"github.com/costinmrr/gontenttype/types/xml"
	"github.com/costinmrr/gontenttype/types/csv"
)

func main() {
	// json
	myStr := `{"foo":"bar"}`
	err := json.IsJSON(myStr)
	fmt.Println(err) // nil

	myStr = `{"foo":"bar"`
	err = json.IsJSON(myStr)
	fmt.Println(err) // unexpected end of JSON input

	// xml
	myStr = `<foo>bar</foo>`
	err = xml.IsXML(myStr)
	fmt.Println(err) // nil

	myStr = `<foo>bar</foo`
	err = xml.IsXML(myStr)
	fmt.Println(err) // XML syntax error on line 1: unexpected EOF

	// csv
	myStr = `foo,bar`
	err = csv.IsCSV(myStr)
	fmt.Println(err) // nil

	myStr = "col1,col2\nfoo,bar,baz"
	err = csv.IsCSV(myStr)
	fmt.Println(err) // record on line 2: wrong number of fields
}
```

#### Output

```shell
<nil>
unexpected end of JSON input
<nil>
XML syntax error on line 1: unexpected EOF
<nil>
record on line 2: wrong number of fields
```
