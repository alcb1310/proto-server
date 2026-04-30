<img width="1134" height="425" alt="protobuf" src="https://github.com/user-attachments/assets/391ee322-69d8-4bcc-a0ac-bee2092cc3c6" />

# Protobuf Tutorial

This project is a tutorial on working with `protobuf` in Go using a `REST` server having the following features:

- A **GET** endpoint which will simulate a database connection and retrieve data from it, then it will *marshall* the information into the compiled version of the structure in `proto` files.
- A **GET** endpoint that will retrieve the same information and send it to the client via `JSON` so we can compare the results.
- A **POST** endpoint that will recieve data through from the client in a `protobuf` format and will *unmarshall* the information into a structure in our code.

## Technologies Used

- ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
- [Chi-Router](https://go-chi.io)
- [Protobuf](https://protobuf.dev/)
- [JSON](https://json.org/)

## Requirements

In order to follow this tutorial you will need the following:

- [Go](https://go.dev/dl/)
- [Protobuf Compiler](https://github.com/protocolbuffers/protobuf/releases)

## Installation

Clone the repository:

```bash
git clone https://github.com/alcb1310/proto-server.git
cd proto-server
go mod tidy
```

## Usage

```bash
make watch
```

## Author

Andrés Court

[![GitHub](https://img.shields.io/badge/GitHub-100000?style=for-the-badge&logo=github&logoColor=white)](https://github.com/alcb1310)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-%230077B5.svg?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/alcb1310/)

## License

[MIT](https://github.com/alcb1310/proto-server/blob/main/LICENSE)
