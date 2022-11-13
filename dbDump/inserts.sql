--
-- PostgreSQL database dump
--

-- Dumped from database version 15.0 (Debian 15.0-1.pgdg110+1)
-- Dumped by pg_dump version 15.0 (Debian 15.0-1.pgdg110+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: avito_db; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA avito_db;


ALTER SCHEMA avito_db OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: masterbalance; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.masterbalance (
    id bigint NOT NULL,
    from_id character varying(100) NOT NULL,
    service_id character varying(100) NOT NULL,
    order_id character varying(100) NOT NULL,
    money_amount character varying(100) NOT NULL
);


ALTER TABLE public.masterbalance OWNER TO postgres;

--
-- Name: masterbalance_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.masterbalance_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.masterbalance_id_seq OWNER TO postgres;

--
-- Name: masterbalance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.masterbalance_id_seq OWNED BY public.masterbalance.id;


--
-- Name: service; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.service (
    id bigint NOT NULL,
    name character varying(100) NOT NULL,
    price character varying(100) NOT NULL
);


ALTER TABLE public.service OWNER TO postgres;

--
-- Name: service_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.service_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.service_id_seq OWNER TO postgres;

--
-- Name: service_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.service_id_seq OWNED BY public.service.id;


--
-- Name: transaction; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transaction (
    id bigint NOT NULL,
    from_id character varying(100) NOT NULL,
    to_id character varying(100) NOT NULL,
    for_service character varying(100) NOT NULL,
    order_id character varying(100) NOT NULL,
    money_amount character varying(100) NOT NULL,
    status character varying(20) NOT NULL,
    date date DEFAULT CURRENT_DATE NOT NULL
);


ALTER TABLE public.transaction OWNER TO postgres;

--
-- Name: transaction_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transaction_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transaction_id_seq OWNER TO postgres;

--
-- Name: transaction_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transaction_id_seq OWNED BY public.transaction.id;


--
-- Name: user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."user" (
    id bigint NOT NULL,
    balance character varying(100) NOT NULL
);


ALTER TABLE public."user" OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_id_seq OWNER TO postgres;

--
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.user_id_seq OWNED BY public."user".id;


--
-- Name: masterbalance id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.masterbalance ALTER COLUMN id SET DEFAULT nextval('public.masterbalance_id_seq'::regclass);


--
-- Name: service id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.service ALTER COLUMN id SET DEFAULT nextval('public.service_id_seq'::regclass);


--
-- Name: transaction id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction ALTER COLUMN id SET DEFAULT nextval('public.transaction_id_seq'::regclass);


--
-- Name: user id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user" ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);


--
-- Data for Name: masterbalance; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- Data for Name: service; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.service (id, name, price) VALUES (2, 'haircut', '15.0');


--
-- Data for Name: transaction; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.transaction (id, from_id, to_id, for_service, order_id, money_amount, status, date) VALUES (2, '4', '2', '1', '6', '100', '24', '2022-11-11');
INSERT INTO public.transaction (id, from_id, to_id, for_service, order_id, money_amount, status, date) VALUES (3, '0', '4', 'billing', '-', '2.0', 'complete', '2022-11-12');
INSERT INTO public.transaction (id, from_id, to_id, for_service, order_id, money_amount, status, date) VALUES (4, '0', '4', 'billing', '-', '2.5', 'complete', '2022-11-12');
INSERT INTO public.transaction (id, from_id, to_id, for_service, order_id, money_amount, status, date) VALUES (6, '4', '0', 'haircut', '4', '100', 'freeze', '2022-11-12');
INSERT INTO public.transaction (id, from_id, to_id, for_service, order_id, money_amount, status, date) VALUES (7, '4', '0', 'haircut', '4', '+2', 'freeze', '2022-11-12');
INSERT INTO public.transaction (id, from_id, to_id, for_service, order_id, money_amount, status, date) VALUES (8, '4', '0', 'haircut', '4', '+2', 'freeze', '2022-11-12');
INSERT INTO public.transaction (id, from_id, to_id, for_service, order_id, money_amount, status, date) VALUES (9, '4', '0', 'haircut', '4', '2', 'freeze', '2022-11-12');
INSERT INTO public.transaction (id, from_id, to_id, for_service, order_id, money_amount, status, date) VALUES (10, '4', '0', 'haircut', '4', '2', 'freeze', '2022-11-12');
INSERT INTO public.transaction (id, from_id, to_id, for_service, order_id, money_amount, status, date) VALUES (13, '2', '0', 'haircut', '4', '4', 'complete', '2022-11-13');
INSERT INTO public.transaction (id, from_id, to_id, for_service, order_id, money_amount, status, date) VALUES (16, '10', '0', 'haircut', '4', '2', 'canceled', '2021-11-13');
INSERT INTO public.transaction (id, from_id, to_id, for_service, order_id, money_amount, status, date) VALUES (14, '10', '0', 'haircut', '4', '2', 'complete', '2021-11-13');
INSERT INTO public.transaction (id, from_id, to_id, for_service, order_id, money_amount, status, date) VALUES (11, '4', '0', 'haircut', '4', '2', 'complete', '2021-11-13');
INSERT INTO public.transaction (id, from_id, to_id, for_service, order_id, money_amount, status, date) VALUES (5, '0', '4', 'billing', '-', '1006.500000', 'complete', '2021-11-12');
INSERT INTO public.transaction (id, from_id, to_id, for_service, order_id, money_amount, status, date) VALUES (12, '4', '0', 'haircut', '4', '2', 'complete', '2021-11-13');
INSERT INTO public.transaction (id, from_id, to_id, for_service, order_id, money_amount, status, date) VALUES (15, '10', '0', 'haircut', '4', '4', 'freeze', '2022-11-13');


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public."user" (id, balance) VALUES (5, '1000');
INSERT INTO public."user" (id, balance) VALUES (6, '1000');
INSERT INTO public."user" (id, balance) VALUES (7, '1000');
INSERT INTO public."user" (id, balance) VALUES (8, '1000');
INSERT INTO public."user" (id, balance) VALUES (9, '1000');
INSERT INTO public."user" (id, balance) VALUES (11, '10');
INSERT INTO public."user" (id, balance) VALUES (4, '796.500000');
INSERT INTO public."user" (id, balance) VALUES (10, '10.000000');


--
-- Name: masterbalance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.masterbalance_id_seq', 2, true);


--
-- Name: service_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.service_id_seq', 2, true);


--
-- Name: transaction_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transaction_id_seq', 16, true);


--
-- Name: user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_id_seq', 11, true);


--
-- Name: masterbalance masterbalance_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.masterbalance
    ADD CONSTRAINT masterbalance_pkey PRIMARY KEY (id);


--
-- Name: service service_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.service
    ADD CONSTRAINT service_pkey PRIMARY KEY (id);


--
-- Name: transaction transaction_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction
    ADD CONSTRAINT transaction_pkey PRIMARY KEY (id);


--
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

