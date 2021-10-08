# account-balance-go
Example of querying the balance of Crypton and UUSD with Utopia Ecosystem API and utopialib-go

## example of use

flags:

```
  -host string
    	Utopia client host (default "http://127.0.0.1")
  -port int
    	client port (default 20000)
  -token string
    	client account token
```

Let's open the X client, set up the token, and try the request:

```bash
go build
./balancetest -host=http://127.0.0.1 -port=20000 -token=C17BF2E95821A6B545DC9A193CBB750B
```

result:

```
[ 150.34 75.9827 ]
```
