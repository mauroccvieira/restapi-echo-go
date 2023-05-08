INSERT INTO users (id, username, name,  password)
values
    (1, 'kvothe', 'Kvothe, the bloodless', 'bloodless'),
    (2, 'kote', 'Kote, the innkeeper', 'innkeeper'),
    (3, 'bast', 'Bast, the fae', 'fae')

ON CONFLICT do nothing;
SELECT setval('users_id_seq', nextval('users_id_seq')-1);