CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount float,
		note TEXT,
		tags TEXT[]
);