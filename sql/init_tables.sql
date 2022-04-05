--
-- Name: customer_data; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.customer_data (
    cpf public.dm_cpf,
    private integer,
    incompleto integer,
    dt_ultima_compra date,
    vlr_ticket_medio real,
    vlr_ticket_ultima_compra real,
    cnpj_loja_mais_frequente public.dm_cnpj,
    cnpj_loja_ultima_compra public.dm_cnpj
);

ALTER TABLE public.customer_data OWNER TO postgres;

--
-- Name: TABLE customer_data; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.customer_data IS 'Store customer imported data';


--
-- Name: customer_data_rejected; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.customer_data_rejected (
    cpf text,
    private text,
    incompleto text,
    dt_ultima_compra text,
    vlr_ticket_medio text,
    vlr_ticket_ultima_compra text,
    cnpj_loja_mais_frequente text,
    cnpj_loja_ultima_compra text
);


ALTER TABLE public.customer_data_rejected OWNER TO postgres;

--
-- Name: TABLE customer_data_rejected; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.customer_data_rejected IS 'Customer rejected data (Does not fit customer data pattern)';

--
-- Name: imported_files; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.imported_files (
    dt_importacao date DEFAULT CURRENT_DATE NOT NULL,
    txt_dados_cliente text NOT NULL
);

ALTER TABLE public.imported_files OWNER TO postgres;

--
-- Name: TABLE imported_files; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON TABLE public.imported_files IS 'This table stores imported files (Full text)';

--
-- Name: imported_files trg_import_customer_data; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER trg_import_customer_data AFTER INSERT OR UPDATE ON public.imported_files FOR EACH STATEMENT EXECUTE FUNCTION public.fc_insert_new_customer_data();
