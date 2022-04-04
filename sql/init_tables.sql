-- Table: public.customer_data
DROP TABLE IF EXISTS public.customer_data;

CREATE TABLE IF NOT EXISTS public.customer_data
(
    cpf dm_cpf COLLATE pg_catalog."default",
    private integer,
    incompleto integer,
    dt_ultima_compra date,
    vlr_ticket_medio real,
    vlr_ticket_ultima_compra real,
    cnpj_loja_mais_frequente dm_cnpj COLLATE pg_catalog."default",
    cnpj_loja_ultima_compra dm_cnpj COLLATE pg_catalog."default"
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.customer_data
    OWNER to nwpguser;

COMMENT ON TABLE public.customer_data
    IS 'Store customer imported data';


-- Table: public.customer_data_rejected
DROP TABLE IF EXISTS public.customer_data_rejected;

CREATE TABLE IF NOT EXISTS public.customer_data_rejected
(
    cpf character varying(255) COLLATE pg_catalog."default",
    int_private character varying(255) COLLATE pg_catalog."default",
    int_incomplete character varying(255) COLLATE pg_catalog."default",
    dt_ultima_compra character varying(255) COLLATE pg_catalog."default",
    vlr_ticket_medio character varying(255) COLLATE pg_catalog."default",
    vlr_ticket_ultima_compra character varying(255) COLLATE pg_catalog."default",
    cnpj_loja_mais_frequente character varying(255) COLLATE pg_catalog."default",
    cnpj_loja_ultima_compra character varying(255) COLLATE pg_catalog."default"
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.customer_data_rejected
    OWNER to nwpguser;

COMMENT ON TABLE public.customer_data_rejected
    IS 'Customer rejected data (Does not fit customer data pattern)';


-- Table: public.imported_files
DROP TABLE IF EXISTS public.imported_files;

CREATE TABLE IF NOT EXISTS public.imported_files
(
    dt_importacao date NOT NULL DEFAULT CURRENT_DATE,
    txt_dados_cliente text COLLATE pg_catalog."default" NOT NULL
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.imported_files
    OWNER to nwpguser;

COMMENT ON TABLE public.imported_files
    IS 'This table stores imported files (Full text)';


