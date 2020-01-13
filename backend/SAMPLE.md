# sample
## union
### schema
<pre>
extend type Query {
    pets: [Pet]!
}

type Dog {
    id: ID!
    name: String!
}

type Cat {
    id: ID!
    name: String!
}

union Pet = Dog | Cat
</pre>

### go model
<pre>
type Pet interface {
	IsPet()
}

type Cat struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (Cat) IsPet()  {}

type Dog struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (Dog) IsPet()  {}
</pre>
