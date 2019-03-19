--
-- PostgreSQL database dump
--

-- Dumped from database version 11.2
-- Dumped by pg_dump version 11.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: companies; Type: TABLE; Schema: public; Owner: greenretro
--

CREATE TABLE public.companies (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.companies OWNER TO greenretro;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: greenretro
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO greenretro;

--
-- Name: users; Type: TABLE; Schema: public; Owner: greenretro
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    company_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO greenretro;

--
-- Name: companies companies_pkey; Type: CONSTRAINT; Schema: public; Owner: greenretro
--

ALTER TABLE ONLY public.companies
    ADD CONSTRAINT companies_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: greenretro
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: companies_name_idx; Type: INDEX; Schema: public; Owner: greenretro
--

CREATE UNIQUE INDEX companies_name_idx ON public.companies USING btree (name);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: greenretro
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: users_company_id_idx; Type: INDEX; Schema: public; Owner: greenretro
--

CREATE INDEX users_company_id_idx ON public.users USING btree (company_id);


--
-- Name: users users_companies_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: greenretro
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_companies_id_fk FOREIGN KEY (company_id) REFERENCES public.companies(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

