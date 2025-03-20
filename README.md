# pet-project 

# Used packages
{
telebot: 
go get -u gopkg.in/telebot.v4
logrus:
go get -u github.com/sirupsen/logrus
GORM: 
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u gorm.io/driver/sqlite
migrate: 
go get -u github.com/golang-migrate/migrate/v4
pq:
go get -u github.com/lib/pq
errors:
go get -u github.com/pkg/errors
chi:
go get -u github.com/go-chi/chi/v5
go get -u github.com/go-chi/chi/middleware
go get -u github.com/go-chi/cors

}, 

Connect and create migrations and docker to db commands next:
Create docker and create psql: sudo docker run -d --name telegram -e POSTGRES_USER=telegram -e POSTGRES_PASSWORD=telegram -p 5433:5432 postgres:13.4
Entrance in docker and db: sudo docker exec -it telegram psql -U telegram -d telegram
Create migrations: migrate create -ext sql -dir ./migrations -seq init 
 and add import in main.go: 	_ "github.com/golang-migrate/migrate/v4/source/file"
