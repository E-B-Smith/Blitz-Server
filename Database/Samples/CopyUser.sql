create or replace
function CopyUserIDToUserID() returns text as
    $$
    declare
        fromID text;
        newID  text;
    begin

    fromID := 'cd4f01ff-ca88-4e4b-9aaf-756660c34ea0';
     newID := '84b868d1-323e-43e2-8667-9b68030d048b';

    --  Copy user tables:
    --      UserTable
    --      EducationTable
    --      EmploymentTable
    --      UserContactTable
    --      EntityTagTable

    -- UserTable

    create temporary table TempUserTable
        (like UserTable) on commit drop;

    insert into TempUserTable
        select * from UserTable
        where userID = fromID;

    update TempUserTable set userid = newID;

    delete from UserTable where userID = newID;
    insert into UserTable
        select * from TempUserTable;

    -- EducationTable

    create temporary table TempEducationTable
        (like EducationTable) on commit drop;

    insert into TempEducationTable
        select * from EducationTable
        where userID = fromID;

    update TempEducationTable set userid = newID;

    delete from EducationTable where userID = newID;
    insert into EducationTable
        select * from TempEducationTable;

    -- EmploymentTable

    create temporary table TempEmploymentTable
        (like EmploymentTable) on commit drop;

    insert into TempEmploymentTable
        select * from EmploymentTable
        where userID = fromID;

    update TempEmploymentTable set userid = newID;

    delete from EmploymentTable where userID = newID;
    insert into EmploymentTable
        select * from TempEmploymentTable;

    -- UserContactTable

    create temporary table TempUserContactTable
        (like UserContactTable) on commit drop;

    insert into TempUserContactTable
        select * from UserContactTable
        where userID = fromID;

    update TempUserContactTable set userid = newID;

    delete from UserContactTable where userID = newID;
    insert into UserContactTable
        select * from TempUserContactTable;

    --      EntityTagTable

    create temporary table TempEntityTagTable
        (like EntityTagTable) on commit drop;

    insert into TempEntityTagTable
        select * from EntityTagTable
        where userID = fromID
          and entityID = fromID::uuid
          and entityType = 1;

    update TempEntityTagTable set
        userid = newID,
        entityID = newID::uuid;

    delete from EntityTagTable
        where userID = newID
          and entityID = newID::uuid
          and entityType = 1;

    insert into EntityTagTable
        select * from TempEntityTagTable;

    return 'User copied';
    end;
    $$
    language plpgsql;

