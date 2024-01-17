--
-- PostgreSQL database dump
--

-- Dumped from database version 16.1
-- Dumped by pg_dump version 16.1 (Homebrew)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: service_request_actions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.service_request_actions (
    action_id integer NOT NULL,
    step_name character varying(50) NOT NULL,
    action_name character varying(50) NOT NULL,
    approved_by character varying(50),
    created_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.service_request_actions OWNER TO postgres;

--
-- Name: service_request_actions_action_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.service_request_actions_action_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.service_request_actions_action_id_seq OWNER TO postgres;

--
-- Name: service_request_actions_action_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.service_request_actions_action_id_seq OWNED BY public.service_request_actions.action_id;


--
-- Name: service_request_actions action_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.service_request_actions ALTER COLUMN action_id SET DEFAULT nextval('public.service_request_actions_action_id_seq'::regclass);


--
-- Name: service_request_actions service_request_actions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.service_request_actions
    ADD CONSTRAINT service_request_actions_pkey PRIMARY KEY (action_id);


--
-- PostgreSQL database dump complete
--

