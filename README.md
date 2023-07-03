## CampusHQ

This repository contains the source code for the CampusHQ backend API.

#### CampusHQ Architecture

CampusHQ uses onion architecture to separate the application into layers. The layers are as follows:
- Domain Layer
- Application Layer
- Infrastructure Layer
- Presentation Layer

![CampusHQ Architecture](campushq.png)

### Getting Started

#### Prerequisites

Make sure you have the following installed on your machine:

- [Docker](https://docs.docker.com/install/)
- [Postgresql](https://hub.docker.com/_/postgres)
- [Go](https://golang.org/doc/install)

#### Installation

Clone the repository
```bash
git clone git@github.com:campushq-official/campushq-api.git
```

Create a docker image for postgresql, run the command below in the root directory of the project
```bash
make postgres
```

Start the application, on Port specified in the .env file
```bash
make server
```

After starting the application, make sure you seed the database with the following command
```bash
make db-seed
```



