# 100 Go Mistake

Anyone who has never made a mistake has never tries anything new.

-- Albert Einstein

**1. Unintended variable shadowing**

The scope of a variable refers to the places a variable can be referenced. In Go, a variable name declared in a block may be re declared in an inner block. This called variable shadowing, is prone to common mistakes.

```go
func main() {
	var seted = false
	if true {
        // variable shadowing
		seted, err := fn()
		if err != nil {
		}
		log.Println(seted)
	} else {
        // variable shadowing
		seted, err := fn()
		if err != nil {
		}
		log.Println(seted)
	}

	log.Println(seted)
}

func fn() (bool, error) {
	return true, nil
}
```

Variable shadowing occurs when a variable name is re declared in an inner block, and we've seen that this practice is prone to mistakes

**2. Misusing init functions**

When a package is initialized, all the constants and variables declarations in the package are evaluated. Then, the init functions are executed.

**Anti pattern: to hold a database connection pool**

```go
var db *sql.DB
func init() {
    dataSourceName := os.Getenv("MYSQL_DATA_SOURCE_NAME")
    d, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        log.Panic(err)
    }

    err = d.Ping()
    if err != nil {
        log.Panic(err)
    }
    db = d
}
```

**Let's describe three main downsides**

- Error management in an init function is limited, only ways to signal an error is to panic, leading application to be stopped
- init function will be executed before running the test cases
- Database connection pool is assigned to a global variable. Global variables have some severe drawbacks
+ Any functions can alter them within the package
+ It can also make unit test more complicated as a function that would depend on it 

**In summary, we have seen that init functions may lead to some issues:**

- Limit error management
- It can complicate how to implement tests
- If the initialization requires to set a state, it has to be done through global variables

**3. Interface pollution**

Interface pollution is about overwhelming our code with unnecessary abstractions making it harder to understand.
The bigger the interface, the weaker the abstraction

**When to use interfaces?**

- Common behavior
- Decoupling
- Restricting behavior

**Common behavior**
The first option we will discuss is to use interface when multiple types implement a common behavior. In such a case, we can factor out the behavior inside an interface.