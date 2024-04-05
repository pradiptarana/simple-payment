# Explanation
This project contains 3 endpoint, get account validation, payout and payout notification

# How to Run
To run the project please do the following:
- Setup PostgreSQL and run the table creation query: `CREATE TABLE transaction (
	id serial PRIMARY KEY,
	amount VARCHAR ( 255 ) NOT NULL,
	status VARCHAR ( 10 ) NOT NULL,
	reference_no VARCHAR ( 255 ) NOT NULL,
	account_no VARCHAR ( 255 ) NOT NULL,
	account_name VARCHAR ( 255 ) NOT NULL,
	bank VARCHAR ( 10 ) NOT NULL,
	created_at TIMESTAMP NOT NULL,
    	updated_at TIMESTAMP
);`
- Create `.env` file. Please refer to `.env.sample` for the env variable that used in this project.
- Run `go run main.go`.

# Mock
- GET https://660ed898356b87a55c50498c.mockapi.io/api/v1/account_validation
- POST https://660ed898356b87a55c50498c.mockapi.io/api/v1/payout

