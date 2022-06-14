CREATE TABLE IF NOT EXISTS registration (
	number_plate VARCHAR(20) NOT NULL,
	vehicle JSONB NOT NULL,
	insurer JSONB NOT NULL
);
