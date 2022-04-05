-- Name: fc_valida_cnpj(character varying, boolean); Type: FUNCTION; Schema: public; Owner: postgres
CREATE FUNCTION public.fc_valida_cnpj(p_cnpj character varying, p_fg_permite_nulo boolean DEFAULT false) RETURNS boolean
    LANGUAGE plpgsql
    AS $_$
declare
    
    v_cnpj_invalidos character varying[10] 
    default array['00000000000000', '11111111111111',
                  '22222222222222', '33333333333333',
                  '44444444444444', '55555555555555',
                  '66666666666666', '77777777777777',
                  '88888888888888', '99999999999999'];
                  
    v_cnpj_quebrado smallint[];
    
    c_posicao_dv1 constant smallint default 13;
    v_arranjo_dv1 smallint[12] default array[5,4,3,2,9,8,7,6,5,4,3,2];
    v_soma_dv1 smallint default 0;
    v_resto_dv1 double precision default 0;
    
    c_posicao_dv2 constant smallint default 14;
    v_arranjo_dv2 smallint[13] default array[6,5,4,3,2,9,8,7,6,5,4,3,2];
    v_soma_dv2 smallint default 0;
    v_resto_dv2 double precision default 0;
    
begin
    
    if p_fg_permite_nulo and nullif(p_cnpj, '') is null then
        return true;
    end if;
    
    if (not (p_cnpj ~* '^([0-9]{14})$' or 
             p_cnpj ~* '^([0-9]{2}\.[0-9]{3}\.[0-9]{3}\/[0-9]{4}\-[0-9]{2})$')) or
        p_cnpj = any (v_cnpj_invalidos) or
        p_cnpj is null
    then
        return false;    
    end if;
    
    v_cnpj_quebrado := regexp_split_to_array(
      regexp_replace(p_cnpj, '[^0-9]', '', 'g'), '');
        
    -- Realiza o calculo do primeiro digito
    for t in 1..12 loop
        v_soma_dv1 := v_soma_dv1 + 
      (v_cnpj_quebrado[t] * v_arranjo_dv1[t]);
    end loop;
    v_resto_dv1 := ((10 * v_soma_dv1) % 11) % 10;
       
    if (v_resto_dv1 != v_cnpj_quebrado[13]) 
    then
        return false;
    end if;
    
    -- Realiza o calculo do segundo digito    
    for t in 1..13 loop
        v_soma_dv2 := v_soma_dv2 + 
      (v_cnpj_quebrado[t] * v_arranjo_dv2[t]);
    end loop;
    v_resto_dv2 := ((10 * v_soma_dv2) % 11) % 10;
    
    return (v_resto_dv2 = v_cnpj_quebrado[c_posicao_dv2]);    
    
end;
$_$;

ALTER FUNCTION public.fc_valida_cnpj(p_cnpj character varying, p_fg_permite_nulo boolean) OWNER TO postgres;

-- Name: FUNCTION fc_valida_cnpj(p_cnpj character varying, p_fg_permite_nulo boolean); Type: COMMENT; Schema: public; Owner: postgres
COMMENT ON FUNCTION public.fc_valida_cnpj(p_cnpj character varying, p_fg_permite_nulo boolean) IS 'CNPJ validation function';

-- Name: dm_cnpj; Type: DOMAIN; Schema: public; Owner: postgres
CREATE DOMAIN public.dm_cnpj AS character varying(22)
	CONSTRAINT dm_cnpj_check CHECK (public.fc_valida_cnpj(VALUE, true));

ALTER DOMAIN public.dm_cnpj OWNER TO postgres;

-- Name: fc_valida_cpf(character varying, boolean); Type: FUNCTION; Schema: public; Owner: postgres
CREATE FUNCTION public.fc_valida_cpf(p_cpf character varying, p_valida_nulo boolean DEFAULT false) RETURNS boolean
    LANGUAGE plpgsql
    AS $_$
declare
    
    v_cpf_invalidos character varying[10] 
    default array['00000000000', '11111111111',
                  '22222222222', '33333333333',
                  '44444444444', '55555555555',
                  '66666666666', '77777777777',
                  '88888888888', '99999999999'];
                  
    v_cpf_quebrado smallint[];
    
    c_posicao_dv1 constant smallint default 10;    
    v_arranjo_dv1 smallint[9] default array[10,9,8,7,6,5,4,3,2];
    v_soma_dv1 smallint default 0;
    v_resto_dv1 double precision default 0;
    
    c_posicao_dv2 constant smallint default 11;
    v_arranjo_dv2 smallint[10] default array[11,10,9,8,7,6,5,4,3,2];
    v_soma_dv2 smallint default 0;
    v_resto_dv2 double precision default 0;
    
begin
    if p_valida_nulo and nullif(p_cpf, '') is null then
        return true;
    end if;
    if (not (p_cpf ~* '^([0-9]{11})$' or 
             p_cpf ~* '^([0-9]{3}\.[0-9]{3}\.[0-9]{3}\-[0-9]{2})$')
        ) or
        p_cpf = any (v_cpf_invalidos) or
        p_cpf is null
    then
        return false;    
    end if;
    
v_cpf_quebrado := regexp_split_to_array(
                    regexp_replace(p_cpf, '[^0-9]', '', 'g'), '');
    -------------------------------- Digito Verificador 1
    for t in 1..9 loop
        v_soma_dv1 := v_soma_dv1 + 
                     (v_cpf_quebrado[t] * v_arranjo_dv1[t]);
    end loop;
    v_resto_dv1 := ((10 * v_soma_dv1) % 11) % 10;
        
    if (v_resto_dv1 != v_cpf_quebrado[c_posicao_dv1]) 
    then
        return false;
    end if;
    
    -------------------------------- Digito Verificador 2
    for t in 1..10 loop
        v_soma_dv2 := v_soma_dv2 + 
                     (v_cpf_quebrado[t] * v_arranjo_dv2[t]);
    end loop;
    v_resto_dv2 := ((10 * v_soma_dv2) % 11) % 10;
    
    return (v_resto_dv2 = v_cpf_quebrado[c_posicao_dv2]);    
    
end;
$_$;

ALTER FUNCTION public.fc_valida_cpf(p_cpf character varying, p_valida_nulo boolean) OWNER TO postgres;

-- Name: FUNCTION fc_valida_cpf(p_cpf character varying, p_valida_nulo boolean); Type: COMMENT; Schema: public; Owner: postgres
COMMENT ON FUNCTION public.fc_valida_cpf(p_cpf character varying, p_valida_nulo boolean) IS 'CPF validation function';

-- Name: dm_cpf; Type: DOMAIN; Schema: public; Owner: postgres
CREATE DOMAIN public.dm_cpf AS character varying(19)
	CONSTRAINT dm_cpf_check CHECK (public.fc_valida_cpf(VALUE, true));


ALTER DOMAIN public.dm_cpf OWNER TO postgres;

-- Name: fc_convert_money_to_real(character varying); Type: FUNCTION; Schema: public; Owner: postgres
CREATE FUNCTION public.fc_convert_money_to_real(ds_value character varying) RETURNS character varying
    LANGUAGE sql
    AS $_$
select
replace(trim($1),',','.');
$_$;


ALTER FUNCTION public.fc_convert_money_to_real(ds_value character varying) OWNER TO postgres;

-- Name: fc_insert_new_customer_data(); Type: FUNCTION; Schema: public; Owner: postgres
CREATE FUNCTION public.fc_insert_new_customer_data() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
  IF (tg_op = 'UPDATE')
    OR
    (
      tg_op ='INSERT'
    )
    THEN
    -- All valid data
    INSERT INTO customer_data
                (
                            cpf,
                            PRIVATE,
                            incompleto,
                            dt_ultima_compra,
                            vlr_ticket_medio,
                            vlr_ticket_ultima_compra,
                            cnpj_loja_mais_frequente,
                            cnpj_loja_ultima_compra
                )
                (
                       SELECT trim(substring(txt_dados_cliente FROM 1 FOR 19))                                     AS cpf,
                              trim(substring(txt_dados_cliente FROM 20 FOR 12))::integer                           AS private,
                              trim(substring(txt_dados_cliente FROM 32 FOR 12))::integer                           AS incompleto,
                              trim(substring(txt_dados_cliente FROM 44 FOR 22))::date                              AS dt_ultima_compra,
                              fc_convert_money_to_real(trim(substring(txt_dados_cliente FROM 66 FOR 22)))::real    AS vlr_ticket_medio,
                              fc_convert_money_to_real(trim(substring(txt_dados_cliente FROM 88 FOR 24)))::real    AS vlr_ticket_ultima_compra,
                              trim(substring(txt_dados_cliente FROM 112 FOR 20))                                   AS cnpj_loja_mais_frequente,
                              trim(substring(txt_dados_cliente FROM 132 FOR 21))                                   AS cnpj_loja_ultima_compra
                       FROM   imported_files
                       WHERE  txt_dados_cliente NOT LIKE '%NULL%' -- Ignore NULL data
                       AND    txt_dados_cliente NOT LIKE '%CPF%'); -- Ignore header
    
    -- All invalid data
    INSERT INTO customer_data_rejected
                (
                            cpf,
                            PRIVATE,
                            incompleto,
                            dt_ultima_compra,
                            vlr_ticket_medio,
                            vlr_ticket_ultima_compra,
                            cnpj_loja_mais_frequente,
                            cnpj_loja_ultima_compra
                )
                (
                       SELECT trim(substring(txt_dados_cliente FROM 1 FOR 19))   AS cpf,
                              trim(substring(txt_dados_cliente FROM 20 FOR 12))  AS private,
                              trim(substring(txt_dados_cliente FROM 32 FOR 12))  AS incompleto,
                              trim(substring(txt_dados_cliente FROM 44 FOR 22))  AS dt_ultima_compra,
                              trim(substring(txt_dados_cliente FROM 66 FOR 22))  AS vlr_ticket_medio,
                              trim(substring(txt_dados_cliente FROM 88 FOR 24))  AS vlr_ticket_ultima_compra,
                              trim(substring(txt_dados_cliente FROM 112 FOR 20)) AS cnpj_loja_mais_frequente,
                              trim(substring(txt_dados_cliente FROM 132 FOR 21)) AS cnpj_loja_ultima_compra
                       FROM   imported_files
                       WHERE  (
                                     NOT PUBLIC.fc_valida_cpf(trim(substring(txt_dados_cliente FROM 1 FOR 19)))      -- Ignore invalid CPF
                              OR     NOT PUBLIC.fc_valida_cnpj(trim(substring(txt_dados_cliente FROM 112 FOR 20)))   -- Ignore invalid CNPJ
                              OR     NOT PUBLIC.fc_valida_cnpj(trim(substring(txt_dados_cliente FROM 132 FOR 21))))  -- Ignore invalid CNPJ
                       AND    txt_dados_cliente NOT LIKE '%CPF%');  -- Ignore header
    
    -- Delete records from imported_Files after insertion on tables
    DELETE FROM imported_files;
    RETURN NEW;
  END IF;
  RETURN NULL;
END;
$$;

ALTER FUNCTION public.fc_insert_new_customer_data() OWNER TO postgres;

-- Name: remove_accent(text); Type: FUNCTION; Schema: public; Owner: postgres
CREATE FUNCTION public.remove_accent(p_text text) RETURNS text
    LANGUAGE sql
    AS $_$  
 Select translate($1,  
 'áàâãäåaaaÁÂÃÄÅAAAÀéèêëeeeeeEEEÉEEÈìíîïìiiiÌÍÎÏÌIIIóôõöoooòÒÓÔÕÖOOOùúûüuuuuÙÚÛÜUUUUçÇñÑýÝ',  
 'aaaaaaaaaAAAAAAAAAeeeeeeeeeEEEEEEEiiiiiiiiIIIIIIIIooooooooOOOOOOOOuuuuuuuuUUUUUUUUcCnNyY'   
  );  
 $_$;

ALTER FUNCTION public.remove_accent(p_text text) OWNER TO postgres;