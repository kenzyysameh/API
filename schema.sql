CREATE TABLE IF NOT EXISTS Users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role TEXT NOT NULL 
);

CREATE TABLE Products (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    price REAL,
    category TEXT NOT NULL,
    stock INTEGER 
);

CREATE TABLE Orders (
    id INTEGER PRIMARY KEY,
    user_id INTEGER,
    total_price REAL,
    status TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(id)
);

CREATE TABLE OrderItems (
    order_id INTEGER,
    product_id INTEGER,
    quantity INTEGER,
    FOREIGN KEY (order_id) REFERENCES Orders(id), -- Added comma
    FOREIGN KEY (product_id) REFERENCES Products(id)
);

cc