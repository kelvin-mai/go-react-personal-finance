create table users(
  id text not null primary key default nanoid(),
  created_at timestamptz not null default now(),
  updated_at timestamptz,
  username text not null unique,
  password text not null
);

create trigger users_updated_at
  before update on users
  for each row
  execute procedure moddatetime (updated_at);