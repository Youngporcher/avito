CREATE TABLE IF NOT EXISTS public.users_segments
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    user_id character varying(255) COLLATE pg_catalog."default" NOT NULL,
    segment character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT users_segments_pkey PRIMARY KEY (id)
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users_segments
    OWNER to postgres;