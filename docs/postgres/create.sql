create table public.city
(
    id         integer not null
        constraint city_pk
            primary key,
    name       varchar,
    population integer
);

alter table public.city
    owner to username;

create table public.country
(
    id      integer not null
        constraint country_pk
            primary key,
    name    varchar,
    capital integer
        constraint country_city_id_fk
            references public.city
);

alter table public.country
    owner to username;


