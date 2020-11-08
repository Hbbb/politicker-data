run:
    go run main.go -persist=false
setup-db:
    psql -h 127.0.0.1 -d $DATABASE_NAME -f $SQL_SETUP_FILE
sql:
    psql -d $DATABASE_NAME
