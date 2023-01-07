-- Table Definition
CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount int,
		note TEXT,
		tags TEXT[]
);