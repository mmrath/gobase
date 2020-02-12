GRANT EXECUTE ON ALL FUNCTIONS IN SCHEMA audit TO db_migration;
GRANT SELECT ON ALL TABLES IN SCHEMA audit TO db_migration;
GRANT USAGE ON SCHEMA audit TO db_migration;


/* PUBLIC should not be allowed to execute functions created by db_migration */
ALTER DEFAULT PRIVILEGES FOR ROLE db_migration REVOKE EXECUTE ON FUNCTIONS FROM PUBLIC;


/* allow clipo to access, but not create objects in the schema */
GRANT USAGE ON SCHEMA public TO clipo;

/* assuming that clipo should be allowed to do anything
with data in all tables in that schema, allow access for all
   objects that db_migration will create there */
ALTER DEFAULT PRIVILEGES FOR ROLE db_migration IN SCHEMA public
    GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO clipo;

ALTER DEFAULT PRIVILEGES FOR ROLE db_migration IN SCHEMA public
    GRANT SELECT, USAGE ON SEQUENCES TO clipo;

ALTER DEFAULT PRIVILEGES FOR ROLE db_migration IN SCHEMA public
    GRANT EXECUTE ON FUNCTIONS TO clipo;

/* Setup for oppo user */
GRANT USAGE ON SCHEMA public TO oppo;
ALTER DEFAULT PRIVILEGES FOR ROLE db_migration IN SCHEMA public
    GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO oppo;

ALTER DEFAULT PRIVILEGES FOR ROLE db_migration IN SCHEMA public
    GRANT SELECT, USAGE ON SEQUENCES TO oppo;

ALTER DEFAULT PRIVILEGES FOR ROLE db_migration IN SCHEMA public
    GRANT EXECUTE ON FUNCTIONS TO oppo;