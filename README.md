# BMS-back-end

This project is part of an assignment of the *Database System Principle* course.  
The aim of this project is to create a `book` management system. Only the `admin` can login to this system, creating `borrow` instance for a borrower. The information of a borrower is stored into `card`. The whole model contains four tables mentioned above: admin, book, borrow and card.  

This project is the back-end part of the assignment. You can find the front-end part at [Enzymii/BMS-front-end](https://github.com/Enzymii/BMS-front-end). While we choose to implement a MVC client-server structure, this project only implements the *model* and *controller* part.  
If you want to write your own front-end while using this back-end, you will find that it is extremely easy, since every API of this project is well documented. This will be talked in the next chapter.  

## Requirements
- A local `MariaDB` or `MySQL` server. You can only choose from these two DBMS.  
- Create a database for this project. You can also create a user who only has the access to this database.  

## Getting Started

### Build
First thing first, you certainly need to download the source code and build it by `go build`.  
You are strongly suggested to use *go 1.16* or higher versions to build this project.  

### Create a configuration file
Be aware that **the executable cannot run properly without a configuration file**. You need to create a configuration file named `conf.yaml` under the same directory of the executable. The content in the configuration file must looks like this:  

```yaml
sql:                               # Choose either MariaDB or MySQL
  username: foo                    # your username
  password: bar                    # your password
  db_name: foofoo                  # your database name
jwt:                               # Yes, this project uses jwt for authentication
  enable: true                     # you can turn this to false while testing this API
  issuer: bms                      # can be anything you want
  max_age: 600                     # seconds
  secret_key: YourSuperSecretKey   # a secret key for jwt, please set this properly!!!
```

### Generate API documentation
You need to install [swaggo/swag](https://github.com/swaggo/swag), 
and execute `swag init` under the root directory of this project.  
This command generates a *Swagger Documentation 2.0* webpage, which contains the API documentation of this service.  

### Run and refer to the API documentation
Now you can run the executable properly.  
While running, the executable will start a server on `localhost:1323`.  
You can visit this URL for API documentation: `localhost:1323/doc/index.html`.  