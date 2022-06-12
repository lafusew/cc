package data

const InitEnumsSQL = `
CREATE TYPE account_authority AS ENUM ('OWNER', 'ADMIN', 'USER');
CREATE TYPE auth_type AS ENUM ('BASIC', 'MAGIC_LINK');
`

const InitPgRoleSQL = `
CREATE ROLE postgres SUPERUSER;
ALTER ROLE "postgres" WITH LOGIN;
`