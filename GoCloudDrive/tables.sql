create table if not exists file
(
    id serial primary key,
    name character varying,
    uploaded_at timestamp with time zone default current_timestamp,
    location character varying,
    status integer default 1, -- 1 active, 0 inactive, -1 deleted
    size integer
);