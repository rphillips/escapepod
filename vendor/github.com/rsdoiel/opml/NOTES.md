
# misc notes

Example using the ",any,attr" xml dsl

```go
    package main
    
    import (
    	"encoding/json"
    	"encoding/xml"
    	"fmt"
    )
    
    type CustomAttrs []xml.Attr
    
    func (cattr CustomAttrs) MarshalJSON() ([]byte, error) {
    	m := map[string]string{}
    	for _, attr := range cattr {
    		k := attr.Name.Local
    		v := attr.Value
    		if k != "" {
    			m[k] = v
    		}
    	}
    
    	return json.Marshal(m)
    }
    
    func main() {
    	type Email struct {
    		Where string `xml:"where,attr"`
    		Addr  string
    		Attrs CustomAttrs `xml:",any,attr" json:"custom_attrs,omitempty"`
    	}
    	type Address struct {
    		City, State string
    	}
    	type Result struct {
    		XMLName xml.Name `xml:"Person" json:"-"`
    		Name    string   `xml:"FullName"`
    		Phone   string
    		Email   []Email
    		Groups  []string `xml:"Group>Value"`
    		Address
    	}
    	v := Result{Name: "none", Phone: "none"}
    	data := `
    		<Person>
    			<FullName>Grace R. Emlin</FullName>
    			<Company>Example Inc.</Company>
    			<Email where="home" preferred="true" is_secret="true">
    				<Addr>gre@example.com</Addr>
    			</Email>
    			<Email where='work' preferred="false">
    				<Addr>gre@work.com</Addr>
    			</Email>
    			<Group>
    				<Value>Friends</Value>
    				<Value>Squash</Value>
    			</Group>
    			<City>Hanga Roa</City>
    			<State>Easter Island</State>
    		</Person>
    	`
    	err := xml.Unmarshal([]byte(data), &v)
    	if err != nil {
    		fmt.Printf("error: %v", err)
    		return
    	}
    	fmt.Printf("XMLName: %#v\n", v.XMLName)
    	fmt.Printf("Name: %q\n", v.Name)
    	fmt.Printf("Phone: %q\n", v.Phone)
    	fmt.Printf("Email: %v\n", v.Email)
    	fmt.Printf("Groups: %v\n", v.Groups)
    	fmt.Printf("Address: %v\n", v.Address)
    	src, _ := json.MarshalIndent(v, "", " ")
    	fmt.Printf("json: %s", src)
    }
```
