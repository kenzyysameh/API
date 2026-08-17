Action list products)
GET /api/v1/products

(Query parameter)
category = category_name

Response : Json :
200: {
    "data": [
        {
            id : 123,
            name:" product_1 ",
            description: " ",
            price: 1500,
            category:"electronics",
            stock: 10
        }
    ]

        
   
}
=====================================
Action create a product )
POST /api/v1/products 

Authentication: Required
Authorization : Bearer <jwt_token>
Access : Admin only

Request : Json:
{
 name: "Wireless Headphones",
 description: "Bluetooth headphones",
 price: 1500,
 category: "electronics",
 stock: 20
}
Response :Json:
 201: 
 Location: /api/v1/products/123
 {
 id :123
 name: "Wireless Headphones",
 description: "Bluetooth headphones",
 price: 1500,
 category: "electronics",
 stock: 20
}
======================================    
Action update a product)
PATCH /api/v1/products/123

Authentication: Required
Authorization : Bearer <jwt_token>
Access : Admin only

Updatable fields :
price 
stock

Request : Json:
{
 "price": 1800
}

Response : Json :
 200: 
        {
            id : 123,
            name:" product 1 ",
            description: " ",
            price: 1800,
            category:"electronics",
            stock: 10
}
=========================================
Action list product reviews )
GET /api/v1/products/42/reviews

Response : Json:
200:{
    "reviews":[ 
        { 
    id : 1
    user : "user_1"
    rating : 4
    comment :"very good "
    creation_date : 24/11/2006
        }
    ]
}
=========================================
Action create a product review )
POST /api/v1/products/42/reviews 

Authentication: Required
Authorization : Bearer <jwt_token>

Request: Json:
{
 "rating": 5,
 "comment": "Great product!"
}

Response : Json:
Location: /api/v1/products/42/reviews/1
 201:{
    id : 1
    user : "user_1"
    rating: 5,
    comment: "Great product!"
    creation_date : 24/11/2006  
}
Error response : Json:
400: {
 "error": {
 "code": "INVALID_RATING",
 "message": "Rating must be between 1 and 5"
 }
}
===========================================
Action place an order)
POST /api/v1/users/10/orders

Request : Json:
{
 "items": [
 {
 "product_id": 42,
 "quantity": 2
 },
 {
 "product_id": 15,
 "quantity": 1
 }
 ]
}

Response : Json:
Location api/v1/user/10/orders/2
201: {
    id :2,
    user:10,
"items": [
 {
 "product_id": 42,
 "quantity": 2
 },
 {
 "product_id": 15,
 "quantity": 1
 }
 ],
    total_price :2000,
    status :"loading order"
}
==========================================
Action cancel an order)
DELETE /api/v1/users/10/orders/2

Authentication: Required
Authorization : Bearer <jwt_token>
Access :user only

Response :
204:

if the order does not exist:
404 :

