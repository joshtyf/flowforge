--
-- PostgreSQL database dump
--

-- Dumped from database version 16.1
-- Dumped by pg_dump version 16.2 (Homebrew)

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
-- Name: service_request_event; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.service_request_event (
    event_id integer NOT NULL,
    event_type character varying NOT NULL,
    service_request_id character varying NOT NULL,
    step_name character varying NOT NULL,
    step_type character varying NOT NULL,
    approved_by character varying,
    created_at timestamp without time zone DEFAULT now()
);


ALTER TABLE public.service_request_event OWNER TO postgres;

--
-- Name: service_request_events_event_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.service_request_events_event_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.service_request_events_event_id_seq OWNER TO postgres;

--
-- Name: service_request_events_event_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.service_request_events_event_id_seq OWNED BY public.service_request_event.event_id;


--
-- Name: service_request_event event_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.service_request_event ALTER COLUMN event_id SET DEFAULT nextval('public.service_request_events_event_id_seq'::regclass);


--
-- Name: service_request_event service_request_events_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.service_request_event
    ADD CONSTRAINT service_request_events_pkey PRIMARY KEY (event_id);


CREATE TABLE public.user (
    user_id character varying NOT NULL,
    name character varying NOT NULL,
    email character varying NOT NULL,
    identity_provider character varying NOT NULL,
    created_on timestamp without time zone DEFAULT now(),
    deleted boolean DEFAULT false
);

ALTER TABLE public.user OWNER TO postgres;

ALTER TABLE ONLY public.user
    ADD CONSTRAINT user_pkey PRIMARY KEY (user_id);

CREATE TABLE public.organization (
    org_id integer NOT NULL,
    name character varying NOT NULL,
    owner character varying NOT NULL,
    created_on timestamp without time zone DEFAULT now(),
    deleted boolean DEFAULT false
);

ALTER TABLE public.organization OWNER TO postgres;

ALTER TABLE ONLY public.organization
    ADD CONSTRAINT organization_pkey PRIMARY KEY (org_id),
    ADD CONSTRAINT organization_user_fkey FOREIGN KEY (owner) REFERENCES public.user (user_id);

CREATE SEQUENCE public.organization_org_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.organization_org_id_seq OWNER TO postgres;

ALTER SEQUENCE public.organization_org_id_seq OWNED BY public.organization.org_id;

ALTER TABLE ONLY public.organization ALTER COLUMN org_id SET DEFAULT nextval('public.organization_org_id_seq'::regclass);

CREATE TABLE public.membership (
    user_id character varying NOT NULL,
    org_id integer NOT NULL,
    role character varying NOT NULL,
    joined_on timestamp without time zone DEFAULT now(),
    deleted boolean DEFAULT false
);

ALTER TABLE public.membership OWNER TO postgres;

ALTER TABLE ONLY public.membership
    ADD CONSTRAINT membership_pkey PRIMARY KEY (user_id, org_id),
    ADD CONSTRAINT membership_user_fkey FOREIGN KEY (user_id) REFERENCES public.user (user_id),
    ADD CONSTRAINT membership_org_fkey FOREIGN KEY (org_id) REFERENCES public.organization (org_id);

--
-- PostgreSQL database dump complete
--

