--
-- PostgreSQL database dump
--

-- Dumped from database version 10.1
-- Dumped by pg_dump version 10.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: datasources; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE datasources (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    fqdn character varying(255) NOT NULL,
    port integer NOT NULL,
    type integer NOT NULL,
    secured boolean NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE datasources OWNER TO postgres;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE schema_migration (
    version character varying(255) NOT NULL
);


ALTER TABLE schema_migration OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE users (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE users OWNER TO postgres;

--
-- Name: datasources datasources_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY datasources
    ADD CONSTRAINT datasources_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: version_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX version_idx ON schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

