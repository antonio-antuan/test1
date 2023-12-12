CREATE TABLE IF NOT EXISTS posts (
    id uuid PRIMARY KEY,
    content text not null
);
insert into posts (id, content) values (gen_random_uuid(), 'hello world'), (gen_random_uuid(), 'hello world 2'), (gen_random_uuid(), 'hello world 3');
