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

<u>Create .env file in the root directory of the project, and add the environment variables specified in the .env.sample file.</u>

Create a docker image for postgresql, run the command below in the root directory of the project by default this will build a postgresql image, with `username` as `root` and password as `root`, running on port :5432.
```bash
make postgres
```
If you have a postgresql instance running on your machine, you can skip this step and update the .env (DB_URL) and MakeFile file with your postgresql credentials.
    
Format for DB_URL:

`postgres://username:password@host:port/database?sslmode=disable`



Start the application, on Port specified in the .env file
```bash
make server
```

After starting the application, make sure you seed the database with the following command
```bash
make db-seed
```

#### Some useful commands

- `db-wipe` This command will drop the database.
- `new-migration name=migration-name` This command will create a new migration file in the migrations' folder. The name flag is required.
- `migrate-up` This command will run all the migrations.
- `migrate-down` This command will roll back the migrations.


### SQLC

SQLC is a code generator for SQL databases. It is used to generate the database layer of the application. It generates the following files:

To perform any database operation, there is a directory called `repositories/queries`, just write the query in the `sql` file and run the command below to generate the database layer.

```bash
make sqlc
```
The tag `--name: FuncName :many` in the sql file is used to specify the name of method to be generated in the database layer.

For Example, if you want to create a student, you can write a query in `/repositories/queries/students.sql` file, and run the command above to generate the database layer.