REVOKE ALL ON SCHEMA public FROM PUBLIC;

/* allow app_user to access, but not create objects in the schema */
GRANT USAGE ON SCHEMA public TO app_user;

/* PUBLIC should not be allowed to execute functions created by app_admin */
ALTER DEFAULT PRIVILEGES FOR ROLE app_admin REVOKE EXECUTE ON FUNCTIONS FROM PUBLIC;

/* assuming that app_user should be allowed to do anything
with data in all tables in that schema, allow access for all
   objects that app_admin will create there */
ALTER DEFAULT PRIVILEGES FOR ROLE app_admin IN SCHEMA public
    GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO app_user;

ALTER DEFAULT PRIVILEGES FOR ROLE app_admin IN SCHEMA public
    GRANT SELECT, USAGE ON SEQUENCES TO app_user;

ALTER DEFAULT PRIVILEGES FOR ROLE app_admin IN SCHEMA public
    GRANT EXECUTE ON FUNCTIONS TO app_user;