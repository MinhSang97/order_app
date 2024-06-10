-- +migrate Up

CREATE TABLE users (
   user_id VARCHAR(255) PRIMARY KEY,
   password VARCHAR(255) NOT NULL,
   name VARCHAR(100) NOT NULL,
   email VARCHAR(255) NOT NULL UNIQUE,
   birth_date VARCHAR(255),
   phone_number VARCHAR(15) NOT NULL UNIQUE,
   telegram VARCHAR(255),
   address VARCHAR(255),
   sex VARCHAR(10),
   role VARCHAR(255) CHECK (role IN ('admin', 'users', 'driver','shop')) DEFAULT NULL,
   created_at TIMESTAMP(3) DEFAULT NULL,
   updated_at TIMESTAMP(3) DEFAULT NULL,
   token VARCHAR(255)
);

CREATE TABLE recover_password (
   id SERIAL PRIMARY KEY,
   user_id VARCHAR(255),
   password_new VARCHAR(255),
   email VARCHAR(255) NOT NULL UNIQUE,
   otp VARCHAR(15) NOT NULL UNIQUE,
   created_at TIMESTAMP(3) DEFAULT NULL,
   CONSTRAINT fk_recover_password_user_id FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE user_addresses (
   address_id SERIAL PRIMARY KEY,
   user_id VARCHAR(255),
   address VARCHAR(255),
   name VARCHAR(100) NOT NULL,
   phone_number VARCHAR(15) NOT NULL,
   type VARCHAR(255) CHECK (type IN ('home', 'office', 'other')) DEFAULT NULL,
   address_default VARCHAR(255),
   lat DECIMAL(10, 8),
   long DECIMAL(11, 8),
   ward_id VARCHAR(255),
   ward_text VARCHAR(255),
   district_id VARCHAR(255),
   district_text VARCHAR(255),
   province_id VARCHAR(255),
   province_text VARCHAR(255),
   national_id VARCHAR(255),
   national_text VARCHAR(255),
   CONSTRAINT fk_user_addresses_user_id FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE user_ratings (
   id SERIAL PRIMARY KEY,
   user_id VARCHAR(255),
   total_transaction FLOAT,
   rating VARCHAR(255) CHECK (rating IN ('member', 'sliver', 'gold','diamond')) DEFAULT NULL,
   CONSTRAINT fk_user_ratings_user_id FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE menu_items (
   item_id SERIAL PRIMARY KEY,
   name VARCHAR(255) NOT NULL,
   description TEXT,
   price DECIMAL(10, 2) NOT NULL,
   image_url VARCHAR(255)
);

CREATE TABLE item_customizations (
   customization_id SERIAL PRIMARY KEY,
   item_id INT,
   customization_option VARCHAR(255),
   extra_price DECIMAL(10, 2),
   CONSTRAINT fk_item_customizations_item_id FOREIGN KEY (item_id) REFERENCES menu_items(item_id)
);

CREATE TABLE promotions (
   promotion_id SERIAL PRIMARY KEY,
   title VARCHAR(255),
   description TEXT,
   discount_percentage DECIMAL(5, 2),
   start_date DATE,
   end_date DATE
);

CREATE TABLE orders (
   order_id SERIAL PRIMARY KEY,
   user_id VARCHAR(255),
   order_date TIMESTAMP NOT NULL,
   total_price DECIMAL(10, 2),
   status VARCHAR(50),
   address_id INT,
   payment_method VARCHAR(50),
   CONSTRAINT fk_orders_user_id FOREIGN KEY (user_id) REFERENCES users(user_id),
   CONSTRAINT fk_orders_address_id FOREIGN KEY (address_id) REFERENCES user_addresses(address_id)
);

CREATE TABLE order_items (
   order_item_id SERIAL PRIMARY KEY,
   order_id INT,
   item_id INT,
   quantity INT,
   price DECIMAL(10, 2),
   CONSTRAINT fk_order_items_order_id FOREIGN KEY (order_id) REFERENCES orders(order_id),
   CONSTRAINT fk_order_items_item_id FOREIGN KEY (item_id) REFERENCES menu_items(item_id)
);

CREATE TABLE order_customizations (
   order_customization_id SERIAL PRIMARY KEY,
   order_item_id INT,
   customization_option VARCHAR(255),
   extra_price DECIMAL(10, 2),
   CONSTRAINT fk_order_customizations_order_item_id FOREIGN KEY (order_item_id) REFERENCES order_items(order_item_id)
);

CREATE TABLE feedbacks (
   feedback_id SERIAL PRIMARY KEY,
   user_id VARCHAR(255),
   rating INT,
   comment TEXT,
   created_at TIMESTAMP,
   CONSTRAINT fk_feedbacks_user_id FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE discount_codes (
   discount_code_id SERIAL PRIMARY KEY,
   code VARCHAR(50) NOT NULL UNIQUE,
   discount_percentage DECIMAL(5, 2),
   valid_from DATE,
   valid_to DATE,
   promotion_id INT,
   CONSTRAINT fk_discount_codes_promotion_id FOREIGN KEY (promotion_id) REFERENCES promotions(promotion_id)
);

CREATE TABLE order_discounts (
   order_discount_id SERIAL PRIMARY KEY,
   order_id INT,
   discount_code_id INT,
   CONSTRAINT fk_order_discounts_order_id FOREIGN KEY (order_id) REFERENCES orders(order_id),
   CONSTRAINT fk_order_discounts_discount_code_id FOREIGN KEY (discount_code_id) REFERENCES discount_codes(discount_code_id)
);

CREATE TABLE payments (
   payment_id SERIAL PRIMARY KEY,
   order_id INT,
   payment_method VARCHAR(50),
   amount DECIMAL(10, 2),
   payment_date TIMESTAMP,
   CONSTRAINT fk_payments_order_id FOREIGN KEY (order_id) REFERENCES orders(order_id)
);

CREATE TABLE history_transaction (
   id SERIAL PRIMARY KEY,
   user_id VARCHAR(255),
   order_id INT,
   transaction FLOAT,
   CONSTRAINT fk_history_orders_user_id FOREIGN KEY (user_id) REFERENCES users(user_id),
   CONSTRAINT fk_history_orders_order_id FOREIGN KEY (order_id) REFERENCES orders(order_id)
);

-- +migrate Down
DROP TABLE IF EXISTS payments;
DROP TABLE IF EXISTS order_discounts;
DROP TABLE IF EXISTS order_customizations;
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS history_transaction;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS discount_codes;
DROP TABLE IF EXISTS feedbacks;
DROP TABLE IF EXISTS item_customizations;
DROP TABLE IF EXISTS menu_items;
DROP TABLE IF EXISTS user_addresses;
DROP TABLE IF EXISTS recover_password;
DROP TABLE IF EXISTS user_ratings;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS promotions;