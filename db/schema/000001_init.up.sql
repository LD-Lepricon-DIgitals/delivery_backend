CREATE TABLE IF NOT EXISTS users
(
    id                   serial PRIMARY KEY,
    user_login           text NOT NULL UNIQUE,
    user_email           text NOT NULL UNIQUE,
    user_hashed_password text NOT NULL
);

CREATE TABLE IF NOT EXISTS transport_categories
(
    id            serial PRIMARY KEY,
    category_name text NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS workers
(
    id                     serial PRIMARY KEY,
    worker_login           text NOT NULL UNIQUE,
    worker_email           text NOT NULL UNIQUE,
    worker_hashed_password text NOT NULL
);

CREATE TABLE IF NOT EXISTS workers_info
(
    worker_id                 int PRIMARY KEY NOT NULL,
    worker_phone              text UNIQUE,
    worker_name               text,
    worker_surname            text,
    worker_city               text NOT NULL,
    worker_transport_name     text,
    worker_transport_reg      text,
    worker_transport_category int,
    CONSTRAINT fk_worker_id FOREIGN KEY (worker_id) REFERENCES workers (id) ON DELETE CASCADE,
    CONSTRAINT fk_worker_transport_category FOREIGN KEY (worker_transport_category) REFERENCES transport_categories (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS users_info
(
    user_id      int PRIMARY KEY NOT NULL,
    user_phone   text UNIQUE,
    user_name    text,
    user_surname text,
    user_city    text NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_addresses
(
    user_id      int NOT NULL,
    user_address text NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT pk_user_address PRIMARY KEY (user_id, user_address)
);

CREATE TABLE IF NOT EXISTS restaurants
(
    id                     serial PRIMARY KEY,
    restaurant_name        text NOT NULL,
    restaurant_photo       text NOT NULL,
    restaurant_rate        float DEFAULT 0.0,
    restaurant_description text NOT NULL,
    restaurant_phone       text NOT NULL,
    restaurant_socials     text NOT NULL,
    restaurant_address     text NOT NULL
);

CREATE TABLE IF NOT EXISTS reviews
(
    id            serial PRIMARY KEY,
    restaurant_id int,
    customer_name text  NOT NULL,
    review_text   text  NOT NULL,
    review_date   date DEFAULT CURRENT_DATE,
    rate          float NOT NULL,
    CONSTRAINT fk_restaurant_id FOREIGN KEY (restaurant_id) REFERENCES restaurants (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dish_categories
(
    id            serial PRIMARY KEY,
    category_name text NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS dishes
(
    id               serial PRIMARY KEY,
    restaurant_id    int NOT NULL,
    dish_name        text NOT NULL,
    dish_description text,
    dish_price       float NOT NULL,
    dish_weight      float NOT NULL,
    dish_photo       text NOT NULL,
    dish_rating      float,
    dish_category    int,
    CONSTRAINT fk_restaurant_id FOREIGN KEY (restaurant_id) REFERENCES restaurants (id) ON DELETE CASCADE,
    CONSTRAINT fk_dish_category FOREIGN KEY (dish_category) REFERENCES dish_categories (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS orders
(
    id            serial PRIMARY KEY,
    customer_id   int NOT NULL,
    restaurant_id int NOT NULL,
    worker_id     int,
    order_price   float NOT NULL,
    CONSTRAINT fk_customer_id FOREIGN KEY (customer_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_restaurant_id FOREIGN KEY (restaurant_id) REFERENCES restaurants (id) ON DELETE CASCADE,
    CONSTRAINT fk_worker_id FOREIGN KEY (worker_id) REFERENCES workers (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS order_dishes
(
    order_id      int PRIMARY KEY NOT NULL,
    dish_id       int NOT NULL,
    dish_quantity int NOT NULL,
    CONSTRAINT fk_order_id FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE,
    CONSTRAINT fk_dish_id FOREIGN KEY (dish_id) REFERENCES dishes (id) ON DELETE CASCADE
);
