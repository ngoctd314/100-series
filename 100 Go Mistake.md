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