CREATE TABLE dbo.PBDM_BP_PERSON_REL(
	BP_KEY	varchar(100)  null,
	PERSON_KEY	varchar(100)  null,
	AUTH_ROLE_ID	varchar(100)  null,
	START_DATE	varchar(200),
	END_DATE	varchar(200),
	AO_SEQ_NO  varchar(200) null,
	POA_TYPE  varchar(200) null,
	AS_OF_DATE varchar(8) null,
);


CREATE TABLE dbo.PBDM_BU(
	BU varchar(200) null,
	BU_ID varchar(200) null,
	SUB_BU_ID varchar(200) null,
	BU_ABBREVIATION varchar(200) null,
	BU_NAME varchar(200) null,
	BU_INTERNAL_NAME varchar(200) null,
	BU_FULL_NAME varchar(200) null,
	BP_TYPE varchar(200) null,
	BUSINESS_ADDRESS varchar(200) null,
	SWIFT_CODE varchar(200) null,
	REPORTING_CURRENCY varchar(200) null,
	DOMICILE varchar(200) null,
	TAX_RESIDENCE varchar(200) null,
	LANGUAGE varchar(200) null,
	EOD_CALENDAR varchar(200) null,
	BRANCH_CODE varchar(200) null,
	BANK_NAME varchar(200) null,
	BRANCH_NAME varchar(200) null,
	SAP_COMPANY varchar(200) null,
	SAP_PROFIT varchar(200) null,
	SAP_COST Varchar(1500) null,
	AS_OF_DATE varchar(8) null
);


CREATE TABLE dbo.PBDM_BU_CALENDAR(
	BU varchar(200) null,
	COUNTRY varchar(200) null,
	DAY_OFF varchar(200) null,
	AS_OF_DATE varchar(8) null
);

CREATE TABLE dbo.PBDM_CONT(
	CONT_KEY	varchar(200) not null,
	CONT_TYPE_ID	varchar(200)  null,
	NAME	varchar(200)  null,
	BP_KEY	varchar(200)  null,
	BU_ID	varchar(200)  null,
	CONT_CCY	varchar(200)  null,
	CRQ_NEXT_REVIEW_DATE	Date,
	DPM_LAST_REVIEW_DATE	Date,
	DPM_NEXT_REVIEW_DATE	Date,
	OVERRIDE_FLAG	varchar(200)  null,
	OVERRIDE_REASON	varchar(200)  null,
	INVESTMENT_MODEL 	varchar(200)  null,
	INVESTMENT_MANDATE 	varchar(200)  null,
	ACCOUNT_MANAGEMENT_FEE	varchar(200)  null,
	VC_SAFEGUARD_ELECTION_CONFIRMATION	varchar(200)  null,
	VULNERABLE_CUSTOMER	varchar(200)  null,
	SERVICE_TYPE_ID	varchar(200)  null,
	BLOCK_CODE_ID	varchar(200)  null,
	PORT_RISK_RATING	varchar(200)  null,
	FINAL_PORT_RISK_RATING	varchar(200)  null,
	TIME_HORIZON	varchar(200)  null,
	OPEN_DATE	Date,
	CLOSE_DATE	Date,
	REMARK	varchar(200)  null,
	STATUS_ID	varchar(200)  null,
	LAST_MODIFY_DATE	Date,
	LAST_MODIFY_USER	varchar(200)  null,
	AS_OF_DATE	varchar(8) null
);

CREATE TABLE dbo.PBDM_COUNTRY(
	COUNTRY_ID	varchar(100)   null,
	COUNTRY_CODE	varchar(100)  null,
	COUNTRY_NAME	varchar(100)  null,
	AS_OF_DATE	varchar(8)  null
);


CREATE TABLE dbo.PBDM_CURRY(
    CURRY_ID varchar(200) null,
    CURRY_CODE varchar(200)  null,
	CURRY_NAME varchar(200)  null,
	AS_OF_DATE varchar(8) null
);

CREATE TABLE dbo.PBDM_MANAGER(
	MANAGER_ID varchar(200) null,
	TEAM varchar(200) null, 
	AS_OF_DATE varchar(8) null
);
CREATE TABLE dbo.PBDM_MINSTR(
	MINSTR_KEY	varchar(200) null,
	OBJ_KEY	varchar(200) null,
	OBJ_TYPE	varchar(200) null,
	ADDR_KEY	varchar(200) null,
	MINSTR_TEMPL_ID	varchar(200) null,
	FREQUENCY	varchar(200) null,
	LANG_SRC_ID	varchar(200) null,
	LANG_ID	varchar(200) null,
	IS_MAIN	varchar(200) null,
	IS_RESID	varchar(200) null,
	MAIL_ACTION_ID	varchar(200) null,
	LAST_MODIFY_DATE	varchar(200) null,
	LAST_MODIFY_USER	varchar(200) null,
	AS_OF_DATE  varchar(8) null
);

CREATE TABLE dbo.PBDM_OE(
    ID  varchar(200)null,
	OE_NAME varchar(200) null,
	OE_TYPE varchar(200) null,
	OE_BU varchar(200) null,
	RESP_ROLE varchar(200) null,
	RESP_PERSON varchar(200) null,
	RESP_VAILD_FROM varchar(200) null,
	RESP_VAILD_TO varchar(200) null,
	AS_OF_DATE varchar(8) null
);

CREATE TABLE dbo.PBDM_PAYMENT_CUT_OFF_TIME(
	BU varchar(200) null,
	PRIORITY varchar(200) null,
	TIME ZONE varchar(200) null,
	TIME varchar(200) null,
	VALUE DATE EXPRESSION varchar(200) null,
	PAY METHOD varchar(200) null,
	CURRENCY varchar(200) null,
	AS_OF_DATE varchar(8) null
);

CREATE TABLE dbo.PBDM_PERSON_PERSON_REL(
	SRC_PERSON_KEY	varchar(200) null,
	TARGET_PERSON_KEY	varchar(2000)  null,
	PERSON_REL_ROLE_ID	varchar(200)  null,
	PERSON_REL_KIND_ID	varchar(200)  null,
	PERSON_REL_SUB_KIND_ID	varchar(100)  null,
	VALID_FROM	varchar(200) null,
	VALID_TO	varchar(200) null,
	AS_OF_DATE  varchar(8)  null
);

CREATE TABLE PBDM_ROLE_FUNCTION (
    ROLE_ID varchar(200),
    ROLE_NAME varchar(200),
    FUNC_ID varchar(200),
    AS_OF_DATE varchar(200),
);
CREATE TABLE dbo.PBDM_USER(
	USER_ID       varchar(100) not null,
	USER_CODE     varchar(100) null,
	USER_NAME     varchar(100) null,
	BU_ID        varchar(100) null,
	USER_TYPE   varchar(100) null,
	USER_EMAIL    varchar(100) null,
	ACTIVE_FLAG    varchar(100) null,
	OPEN_DATE     date,
	CLOSE_DATE    date,
	RM_CODE      varchar(100) null,
	LAST_LOGIN     varchar(100) null,
	AS_OF_DATE    varchar(8) not null,
	CONSTRAINT [PK_PBDM_USER] PRIMARY KEY(USER_ID,AS_OF_DATE)
);

CREATE TABLE dbo.PBDM_USER_ROLE(
	USER_CODE     varchar(100)  null,
	ROLE_ID       varchar(100)  null,
	AS_OF_DATE     varchar(8)  null
);

