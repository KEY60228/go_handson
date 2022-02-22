CREATE TABLE mydata (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    mail TEXT NOT NULL,
    age INTEGER
);

INSERT INTO mydata (name, mail, age) VALUES (
    'Taro', 'taro@yamada.com', 39
), (
    'Hanako', 'hanako@flower.com', 28
), (
    'Sachiko', 'sachiko@happy.com', 17
), (
    'Jiro', 'jiro@change.com', 6
);

CREATE TABLE md_data (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    url TEXT NOT NULL,
    markdown TEXT
);

CREATE TABLE md_data ( id SERIAL PRIMARY KEY, title TEXT NOT NULL, url TEXT NOT NULL, markdown TEXT);