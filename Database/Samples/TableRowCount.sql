
--  Shows the row counts of all tables

create or replace function
table_row_count(schema text, tablename text) returns integer
    as
    $$
    declare
      result integer;
      query varchar;
    begin
      query := 'SELECT count(1) FROM ' || schema || '.' || tablename;
      execute query into result;
      return result;
    end;
    $$
    language plpgsql;


select
  table_schema,
  table_name,
  table_row_count(table_schema, table_name)
from information_schema.tables
where
  table_schema not in ('pg_catalog', 'information_schema')
  and table_type='BASE TABLE'
order by 1, 2;
