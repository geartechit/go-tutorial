create table if not exists employees (
    "id" uuid primary key default gen_random_uuid(),
    "name" varchar not null,
    "dob" date not null,
    "department" varchar not null,
    "job_title" varchar not null,
    "address" varchar not null,
    "joined_at" timestamptz not null,
    "created_at" timestamptz not null default now(),
    "updated_at" timestamptz not null default now()
);