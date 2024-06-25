create table category (
  id text not null primary key default nanoid(),
  created_at timestamptz not null default now(),
  updated_at timestamptz,
  user_id text not null references users(id),
  title text not null
);

create trigger category_updated_at
  before update on category
  for each row
  execute procedure moddatetime (updated_at);

create type transaction_type as enum ('income', 'expense'); 

create table transactions (
  id text not null primary key default nanoid(),
  created_at timestamptz not null default now(),
  updated_at timestamptz,
  user_id text not null references users(id),
  category_id text not null references category(id),
  title text not null,
  currency text not null,
  amount numeric(12,2) not null,
  type transaction_type not null
);

create trigger transactions_updated_at
  before update on transactions
  for each row
  execute procedure moddatetime (updated_at);