CREATE TABLE IF NOT EXISTS public.segments
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    segment character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT segments_pkey PRIMARY KEY (id),
    CONSTRAINT segment UNIQUE (segment)
    INCLUDE(segment)
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.segments
    OWNER to postgres;