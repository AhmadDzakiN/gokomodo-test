CREATE TABLE IF NOT EXISTS buyers (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    shipping_address text NOT NULl,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS sellers (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    pickup_address text NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY,
    name varchar(255) NOT NULL,
    description text NOT NULL,
    price BIGINT NOT NULL,
    seller_id uuid NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    FOREIGN KEY (seller_id) REFERENCES sellers(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS orders (
    id BIGSERIAL PRIMARY KEY,
    buyer_id uuid NOT NULL,
    seller_id uuid NOT NULL,
    source_address text NOT NULL,
    destination_address text NOT NULL,
    items BIGINT NOT NULL,
    quantity INT NOT NULL,
    price BIGINT NOT NULL,
    total_price BIGINT NOT NULL,
    status varchar(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    FOREIGN KEY (seller_id) REFERENCES sellers(id) ON DELETE CASCADE,
    FOREIGN KEY (buyer_id) REFERENCES buyers(id) ON DELETE CASCADE,
    FOREIGN KEY (items) REFERENCES products(id) ON DELETE CASCADE
);

insert into buyers (id, email, password, name, shipping_address) values ('449d3cef-c5d2-483b-8084-ac30e1bab74c', 'bfriedlos0@shop-pro.jp', 'dO9=0,{>NQXSk', 'Benjy Friedlos', '34845 Aberg Circle');
insert into buyers (id, email, password, name, shipping_address) values ('5c8b708a-e848-4880-ab5a-911ecc8ec246', 'rjagg1@prweb.com', 'tE3"dctu''Qt', 'Ragnar Jagg', '3697 Burning Wood Crossing');
insert into buyers (id, email, password, name, shipping_address) values ('68071bf7-54cd-40a4-9b66-cc48bae25462', 'fmougin2@si.edu', 'jJ3>P*3t%!.Hg', 'Fredek Mougin', '871 Dottie Avenue');
insert into buyers (id, email, password, name, shipping_address) values ('d4b7967c-f5dd-4a99-96fc-e0d8a4c4109e', 'jkivlin3@nasa.gov', 'eU7''A,d9g|c+', 'Jerrilyn Kivlin', '548 Graceland Plaza');
insert into buyers (id, email, password, name, shipping_address) values ('e3f43ceb-3765-4584-9998-7b4f24f2b733', 'adacca4@google.it', 'gR6`+&uMgofj9D,r', 'Andrus Dacca', '73 Crowley Lane');

insert into sellers (id, email, password, name, pickup_address) values ('1f8c1dda-5c1d-455c-a2b8-113edef9da85', 'tgeany0@hao123.com', 'yD6@Tl%Y|2Pn', 'Truda Geany', '6132 Westport Way');
insert into sellers (id, email, password, name, pickup_address) values ('316dee50-8501-4695-8ee6-a074abb6ff30', 'tchatell1@huffingtonpost.com', 'rO4>UL"Z2', 'Tommie Chatell', '58 Harbort Trail');
insert into sellers (id, email, password, name, pickup_address) values ('3120dff7-423d-405c-8acc-519534c02b02', 'akowalski2@dmoz.org', 'cS4?`\SsG}x2c', 'Angelo Kowalski', '930 Longview Circle');
insert into sellers (id, email, password, name, pickup_address) values ('7f1db060-09d5-45fb-ae40-a461f1d55380', 'tandrzej3@nps.gov', 'dH6=a(KIxjE', 'Titus Andrzej', '4 Algoma Crossing');
insert into sellers (id, email, password, name, pickup_address) values ('57deaf0f-095c-4115-a4d5-176daa44edf0', 'lswynley4@vkontakte.ru', 'oA4<fTLF10AL', 'Lola Swynley', '36569 Del Sol Parkway');

insert into products (name, description, price, seller_id, updated_at) values ('verykool SL5565 Rocket', 'verykool', 1000000, '1f8c1dda-5c1d-455c-a2b8-113edef9da85', '2024-05-17 12:45');
insert into products (name, description, price, seller_id, updated_at) values ('Acer Liquid Jade S', 'Acer', 1500000, '1f8c1dda-5c1d-455c-a2b8-113edef9da85', '2024-05-17 12:46');
insert into products (name, description, price, seller_id, updated_at) values ('Huawei Watch Fit', 'Huawei', 3000000, '1f8c1dda-5c1d-455c-a2b8-113edef9da85', '2024-05-17 12:47');
insert into products (name, description, price, seller_id, updated_at) values ('Tecno Spark 5 Air', 'Tecno', 1900000, '316dee50-8501-4695-8ee6-a074abb6ff30', '2024-05-17 12:48');
insert into products (name, description, price, seller_id, updated_at) values ('Motorola L2', 'Motorola', 2500000, '316dee50-8501-4695-8ee6-a074abb6ff30', '2024-05-17 12:49');
insert into products (name, description, price, seller_id, updated_at) values ('LG G3000', 'LG', 4300000, '316dee50-8501-4695-8ee6-a074abb6ff30', '2024-05-17 12:50');
insert into products (name, description, price, seller_id, updated_at) values ('BLU Studio G2', 'BLU', 7000000, '7f1db060-09d5-45fb-ae40-a461f1d55380', '2024-05-17 12:51');
insert into products (name, description, price, seller_id, updated_at) values ('Blackview A10', 'Blackview', 5200000, '7f1db060-09d5-45fb-ae40-a461f1d55380', '2024-05-17 12:52');
insert into products (name, description, price, seller_id, updated_at) values ('LG KP199', 'LG', 6000000, '3120dff7-423d-405c-8acc-519534c02b02', '2024-05-17 12:53');
insert into products (name, description, price, seller_id, updated_at) values ('Samsung Galaxy Trend II Duos S7572', 'Samsung', 6500000, '3120dff7-423d-405c-8acc-519534c02b02', '2024-05-17 12:54');

insert into orders (buyer_id, seller_id, source_address, destination_address, items, quantity, price, total_price, status) values ('449d3cef-c5d2-483b-8084-ac30e1bab74c', '1f8c1dda-5c1d-455c-a2b8-113edef9da85', '6132 Westport Way', '34845 Aberg Circle', 1, 2, 1000000, 2000000, 'Accepted');
insert into orders (buyer_id, seller_id, source_address, destination_address, items, quantity, price, total_price, status) values ('449d3cef-c5d2-483b-8084-ac30e1bab74c', '1f8c1dda-5c1d-455c-a2b8-113edef9da85', '6132 Westport Way', '34845 Aberg Circle', 2, 1, 1500000, 1500000, 'Pending');
insert into orders (buyer_id, seller_id, source_address, destination_address, items, quantity, price, total_price, status) values ('449d3cef-c5d2-483b-8084-ac30e1bab74c', '1f8c1dda-5c1d-455c-a2b8-113edef9da85', '6132 Westport Way', '34845 Aberg Circle', 2, 3, 1500000, 4500000, 'Pending');
insert into orders (buyer_id, seller_id, source_address, destination_address, items, quantity, price, total_price, status) values ('5c8b708a-e848-4880-ab5a-911ecc8ec246', '316dee50-8501-4695-8ee6-a074abb6ff30', '58 Harbort Trail', '3697 Burning Wood Crossing', 5, 2, 2500000, 5000000, 'Accepted');
insert into orders (buyer_id, seller_id, source_address, destination_address, items, quantity, price, total_price, status) values ('5c8b708a-e848-4880-ab5a-911ecc8ec246', '316dee50-8501-4695-8ee6-a074abb6ff30', '58 Harbort Trail', '3697 Burning Wood Crossing', 6, 1, 4300000, 4300000, 'Accepted');
insert into orders (buyer_id, seller_id, source_address, destination_address, items, quantity, price, total_price, status) values ('5c8b708a-e848-4880-ab5a-911ecc8ec246', '316dee50-8501-4695-8ee6-a074abb6ff30', '58 Harbort Trail', '3697 Burning Wood Crossing', 4, 1, 1900000, 1900000, 'Pending');
insert into orders (buyer_id, seller_id, source_address, destination_address, items, quantity, price, total_price, status) values ('68071bf7-54cd-40a4-9b66-cc48bae25462', '3120dff7-423d-405c-8acc-519534c02b02', '871 Dottie Avenue', '3697 Burning Wood Crossing', 10, 1, 6500000, 6500000, 'Accepted');
insert into orders (buyer_id, seller_id, source_address, destination_address, items, quantity, price, total_price, status) values ('68071bf7-54cd-40a4-9b66-cc48bae25462', '3120dff7-423d-405c-8acc-519534c02b02', '871 Dottie Avenue', '3697 Burning Wood Crossing', 9, 2, 6000000, 12000000, 'Pending');
insert into orders (buyer_id, seller_id, source_address, destination_address, items, quantity, price, total_price, status) values ('d4b7967c-f5dd-4a99-96fc-e0d8a4c4109e', '3120dff7-423d-405c-8acc-519534c02b02', '4 Algoma Crossing', '548 Graceland Plaza', 8, 1, 5200000, 5200000, 'Accepted');
insert into orders (buyer_id, seller_id, source_address, destination_address, items, quantity, price, total_price, status) values ('d4b7967c-f5dd-4a99-96fc-e0d8a4c4109e', '3120dff7-423d-405c-8acc-519534c02b02', '4 Algoma Crossing', '548 Graceland Plaza', 7, 1, 7000000, 7000000, 'Pending');