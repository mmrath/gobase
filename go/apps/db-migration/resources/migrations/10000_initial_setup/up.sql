CREATE OR REPLACE FUNCTION auto_manage_updated_at_and_version(_tbl regclass)
    RETURNS VOID AS
$$
BEGIN
    EXECUTE format('CREATE TRIGGER set_updated_at BEFORE UPDATE ON %s
                      FOR EACH ROW EXECUTE PROCEDURE auto_set_audit_columns()', _tbl);
END;
$$
    LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION auto_set_audit_columns()
    RETURNS trigger AS
$$
BEGIN
    IF (NEW IS DISTINCT FROM OLD)
    THEN
        NEW.updated_at := current_timestamp;
        NEW.version := OLD.version + 1;
    END IF;
    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;


CREATE TABLE user_account
(
    id           BIGINT GENERATED ALWAYS AS IDENTITY,
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by   TEXT                     NOT NULL,
    version      INTEGER                  NOT NULL DEFAULT 1,
    first_name   TEXT                     NOT NULL,
    last_name    TEXT                     NOT NULL,
    email        TEXT                     NOT NULL,
    phone_number TEXT                     NULL,
    account_type INTEGER                  NULL,
    active       BOOLEAN                  NOT NULL DEFAULT FALSE,
    expires_at   TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT user_account_uk PRIMARY KEY (id)
);

CREATE UNIQUE INDEX user_account_uk_email ON user_account (lower(email));
CREATE UNIQUE INDEX user_account_uk_phone_number ON user_account (lower(phone_number));

SELECT auto_manage_updated_at_and_version('user_account');
SELECT audit.audit_table('user_account');

CREATE TABLE user_credential
(
    id                        BIGINT,
    updated_at                TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
    version                   INTEGER                  NOT NULL DEFAULT 1,
    password_hash             TEXT,
    expires_at                TIMESTAMP WITH TIME ZONE,
    invalid_attempts          INT                      NOT NULL DEFAULT 0,
    locked                    BOOLEAN                  NOT NULL DEFAULT FALSE,
    activation_key            TEXT,
    activation_key_expires_at TIMESTAMP WITH TIME ZONE,
    activated                 BOOLEAN                  NOT NULL DEFAULT FALSE,
    reset_key                 TEXT,
    reset_key_expires_at      TIMESTAMP WITH TIME ZONE,
    reset_at                  TIMESTAMP WITH TIME ZONE,
    CONSTRAINT user_credential_pk PRIMARY KEY (id),
    CONSTRAINT user_credential_fk_01 FOREIGN KEY (id) REFERENCES user_account (id)
);

SELECT auto_manage_updated_at_and_version('user_credential');
SELECT audit.audit_table('user_credential');


CREATE TABLE auth_token
(
    id         BIGINT GENERATED ALWAYS AS IDENTITY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT current_timestamp,
    user_id    BIGINT                   NOT NULL REFERENCES user_account (id),
    token      TEXT                     NOT NULL UNIQUE,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    mobile     BOOLEAN                  NOT NULL DEFAULT FALSE,
    identifier TEXT,
    CONSTRAINT auth_token_pk PRIMARY KEY (id)
);
SELECT audit.audit_table('auth_token');



CREATE TABLE role
(
    id          INT GENERATED ALWAYS AS IDENTITY,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by  TEXT                     NOT NULL,
    version     INTEGER                  NOT NULL DEFAULT 1,
    name        TEXT                     NOT NULL,
    description TEXT                     NOT NULL,
    CONSTRAINT role_pk PRIMARY KEY (id)
);
CREATE UNIQUE INDEX role_uk_name ON role (lower(name));
SELECT auto_manage_updated_at_and_version('role');
SELECT audit.audit_table('role');


CREATE TABLE permission
(
    id          INT GENERATED ALWAYS AS IDENTITY,
    resource    TEXT NOT NULL,
    authority   TEXT NOT NULL,
    description TEXT NOT NULL,
    CONSTRAINT pk_permission PRIMARY KEY (id)
);

SELECT audit.audit_table('permission');


CREATE UNIQUE INDEX permission_uk_01 ON permission (lower(resource), lower(authority));


CREATE TABLE role_permission
(
    role_id       INTEGER NOT NULL,
    permission_id INTEGER NOT NULL,
    CONSTRAINT role_permissions_pk PRIMARY KEY (role_id, permission_id),
    CONSTRAINT role_permission_fk_01 FOREIGN KEY (role_id) REFERENCES role (id),
    CONSTRAINT role_permission_fk_02 FOREIGN KEY (permission_id) REFERENCES permission (id)
);

SELECT audit.audit_table('role_permission');



CREATE TABLE user_role
(
    user_id BIGINT  NOT NULL,
    role_id INTEGER NOT NULL,
    CONSTRAINT user_role_pk PRIMARY KEY (user_id, role_id),
    CONSTRAINT user_role_fk_01 FOREIGN KEY (role_id) REFERENCES role (id),
    CONSTRAINT user_role_fk_02 FOREIGN KEY (user_id) REFERENCES user_account (id)
);
SELECT audit.audit_table('user_role');

CREATE TABLE user_group
(
    id          INT GENERATED ALWAYS AS IDENTITY,
    updated_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_by  TEXT                     NOT NULL,
    version     INTEGER                  NOT NULL DEFAULT 1,
    name        TEXT                     NOT NULL,
    description TEXT                     NOT NULL,
    CONSTRAINT user_group_pk PRIMARY KEY (id)
);
SELECT audit.audit_table('user_group');

CREATE UNIQUE INDEX user_group_uk ON user_group (lower(name));

CREATE TABLE user_group_user
(
    group_id INTEGER NOT NULL,
    user_id  BIGINT  NOT NULL,
    CONSTRAINT user_group_user_pk PRIMARY KEY (group_id, user_id),
    CONSTRAINT user_group_user_fk_01 FOREIGN KEY (group_id) REFERENCES user_group (id),
    CONSTRAINT user_group_user_fk_02 FOREIGN KEY (user_id) REFERENCES user_account (id)
);
SELECT audit.audit_table('user_group_user');


CREATE TABLE user_group_role
(
    group_id INTEGER NOT NULL,
    role_id  INTEGER NOT NULL,
    CONSTRAINT user_group_role_pk PRIMARY KEY (group_id, role_id),
    CONSTRAINT user_group_role_fk_01 FOREIGN KEY (group_id) REFERENCES user_group (id),
    CONSTRAINT user_group_role_fk_02 FOREIGN KEY (role_id) REFERENCES role (id)
);

SELECT audit.audit_table('user_group_role');

