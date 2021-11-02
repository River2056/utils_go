CREATE TABLE SC_FACILITY_ESB (
	REQUEST_ID	BIGINT NOT NULL,
	TRANSACTION_TYPE	VARCHAR(1) NULL,
	CIF_NO	VARCHAR(7) NULL,
	PORTFOLIO_CODE	VARCHAR(11) NULL,
	LINE_CODE	VARCHAR(9) NULL,
	LINE_SERIAL_NO	VARCHAR(9) NULL,
	MAIN_LINE_CODE	VARCHAR(11) NULL,
	CURRENCY	VARCHAR(3) NULL,
	AMOUNT	NUMERIC(11, 2) NULL,
	START_DATE	DATE NULL,
	EXPIRY_DATE	DATE NULL,
	AVAILABLE_FLAG	VARCHAR(1) NULL,
	CREATE_DATE	DATE NULL,
);