# Initial APP

PostgreSQL DB

#### run this command first to install golang-migrate
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

#### Create new table
migrate create -ext sql -dir <your_path_name> <your_table_name> 
