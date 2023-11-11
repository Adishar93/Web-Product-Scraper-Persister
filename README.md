# Go_Amazon_Scraper

## Build & Run Instructions

First make sure you don't have MongoDB service/database running on your system already which could cause port number conflicts. Now open command prompt/terminal
in the project's root folder and run the command **'docker-compose up'**. Assuming you have docker setup on your system, this should lead to building of required
containers and the three services should be up and running. 
*Note : The setup was tested on Windows Operating system.*

## Testing Instructions

You can make use of POSTMAN to hit the endpoints.

1. After you build & run the application,  you can hit the POST endpoint for the **SCRAPER** service
   - POST Endpoint :
     http://localhost:8080/scrapeProduct OR http://127.0.0.1:8080/scrapeProduct 
   - Sample JSON Request :
```
{
   "url": "https://www.amazon.com/PlayStation-4-Pro-1TB-Console/dp/B01LOP8EZC" 
}
```

   *Note : After you hit this service, it will also internally call PERSISTER service to store in MongoDB.*

2. **PERSISTER** service is already tested somewhat in the first point. If you still want to test this service thoroughly, the details are as follows
   - POST Endpoint :
     http://localhost:8081/persistProduct OR http://127.0.0.1:8081/persistProduct
   - Sample JSON Request : 
     ```
     { 
         "url": "https://www.amazon.com/Xbox-Dazzling-Bundle-Controllers-Kwalicable/dp/B095J5FM5D/ref=sr_1_2?dchild=1&keywords=xbox&qid=1631306749&sr=8-2", 
         "timestamp": "2021-12-11 16:29:37.3418367 +0530 IST", 
         "product": { 
             "name": "Xbox Series S Dazzling Bundle | Includes: Xbox Series S 512GB Console, 2 Wireless  Controllers for Xbox, 3 Month Game Pass Code, Kwalicable Accessory Pack", 
             "imageURL": "https://images-na.ssl-images-amazon.com/images/I/61NPXqxIROL.__AC_SY300_SX300_QL70_ML2_.jpg", 
             "description": "\nVALUE BUNDLE INCLUDES: Xbox Series S 512GB Console, 2 Wireless Controllers for Xbox Series S, 3 Month Game Pass Code, Kwalicable 6ft High Speed HDMI Cable,
             Kwalicable Microfiber Cleaning Cloth......", 
             "price": "$599.00", 
             "totalRatings": "9 ratings" 
         } 
     }
     ```

3. I have added an additional endpoint in **PERSISTER** to view all the products stored in the MongoDB database.
   - GET Endpoint :
     http://localhost:8081/products OR http://127.0.0.1:8081/products \
Once you hit the endpoint, you will receive the list of products stored in database as a response.

4. In case you want to connect to the MongoDB server using bash, you can use the command **'mongosh "mongodb://127.0.0.1:27017" --username mongoDB'** and the password will be *'mongoDB'*
without the quotes. The data is stored in the database 'amascr' inside the collection 'products'.


## About

Project consists of three services defined in the docker-composer yaml file:

- The first REST service **SCRAPER** consists of a POST endpoint that takes product url for 'www.amazon.com' only and scrapes product's information while returning the same in response.
It also calls the 2nd REST service after scraping details successfully to persist data in database. The service validates the input url to only belong to the domain 'www.amazon.com'

- The second REST service **PERSISTER** consists of two endpoints where the POST endpoint takes product details as request and persists the data in NoSQL MongoDB database and
the GET endpoint fetches all the product data stored in the MongoDB database. This is a smart service as it updates the data for the product, in case the same url is hit again
in the future.

- Third one is the MongoDB service itself.
