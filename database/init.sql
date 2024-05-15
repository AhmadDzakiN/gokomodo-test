CREATE TABLE IF NOT EXISTS buyers (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    shipping_address text NOT NULl
);

CREATE TABLE IF NOT EXISTS sellers (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    pickup_address text NOT NULL
);

CREATE TABLE IF NOT EXISTS products (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name varchar(255) NOT NULL,
    description text NOT NULL,
    price BIGINT NOT NULL,
    seller_id uuid NOT NULL,
    FOREIGN KEY (seller_id) REFERENCES sellers(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS orders (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    buyer_id uuid NOT NULL,
    seller_id uuid NOT NULL,
    source_address text NOT NULL,
    destination_address text NOT NULL,
    items UUID NOT NULL,
    quantity INT NOT NULL,
    price BIGINT NOT NULL,
    total_price BIGINT NOT NULL,
    status varchar(100) NOT NULL,
    FOREIGN KEY (seller_id) REFERENCES sellers(id) ON DELETE CASCADE,
    FOREIGN KEY (buyer_id) REFERENCES buyers(id) ON DELETE CASCADE,
    FOREIGN KEY (items) REFERENCES products(id) ON DELETE CASCADE
);