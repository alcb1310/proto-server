# Protobuf Tutorial

This project is a tutorial on working with `protobuf` in Go using a `REST` server having the following features:

- A **GET** endpoint which will simulate a database connection and retrieve data from it, then it will *marshall* the information into the compiled version of the structure in `proto` files.
- A **GET** endpoint that will retrieve the same information and send it to the client via `JSON` so we can compare the results.
- A **POST** endpoint that will recieve data through from the client in a `protobuf` format and will *unmarshall* the information into a structure in our code.
