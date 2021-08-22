docker run -d -t --env DRIVER=$DRIVER --env USER=$USER --env PASSWORD=$PASSWORD \
 --env HOST=$HOST --env PORT=$PORT --env DBNAME=$DBNAME -p 15000:15000 gussf/bookstore:latest go-bookstore