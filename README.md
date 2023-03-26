# grest

This is a very small helper lib to create HTTP handlers out of more "generic" functions.

This package depends on `github.com/gorilla/mux` for the path variables parsing.

It may be that this package is more useful as an example than to be used directly, as-is.

## explanation

BTW, it's probably easier to understand if you just read the code, than try to
understand this lame explanation.

Given a function something like this:

    func foo(vars map[string]string, request T1) (T2, error)

That is, given a function that takes a set of variables (to be parsed from the
HTTP request path) and a request of some type (T1), and the function returns
something of some type (T2) and an error, turn that into an HTTP handler function.

The created handler function will parse the request body (with the standard JSON
parser) into a variable of type T1, and pass that to the given function (`foo`).
The return value (of type T2) will be marshalled into JSON, and then written as
the response.

If a non-nil error value is returned, then a server error will be written,
instead of the other value.

## the example

Run the example thusly:

    $ go run main/example.go

and in another terminal:

    $ curl -sSX POST http://localhost:8080/aloha -d '{"name": "mark-i-mark", "age": 78, "favoriteFood": "bananas"}'
    {
        "status": "OK",
        "message": "aloha, mark-i-mark! You are 78 years old. Your favorite food is bananas.",
        "timestamp": "2023-03-26 05:56:30.083056 +0000 UTC"
    }

    $  curl -sSX GET http://localhost:8080/aloha | jq
    {
        "status": "OK",
        "message": "aloha!",
        "timestamp": "2023-03-26 06:25:15.160301 +0000 UTC"
    }