CREATE TABLE "users" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "email" text NULL,
  "name" text NULL,
  "password" text NULL,
  PRIMARY KEY ("id")
);
CREATE INDEX "idx_users_deleted_at" ON "users" ("deleted_at");
CREATE TABLE "roles" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NULL,
  PRIMARY KEY ("id")
);
CREATE INDEX "idx_roles_deleted_at" ON "roles" ("deleted_at");

-- Seed roles
INSERT INTO "roles" ("name", "created_at", "updated_at") VALUES ('admin', NOW(), NOW()), ('viewer', NOW(), NOW());

CREATE TABLE "user_roles" (
  "user_id" bigint NOT NULL,
  "role_id" bigint NOT NULL,
  PRIMARY KEY ("user_id", "role_id"),
  CONSTRAINT "fk_user_roles_user" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_user_roles_role" FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);

CREATE TABLE "permissions" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NULL,
  PRIMARY KEY ("id")
);
CREATE INDEX "idx_permissions_deleted_at" ON "permissions" ("deleted_at");

-- Seed permissions
INSERT INTO "permissions" ("name", "created_at", "updated_at") VALUES ('create:all', NOW(), NOW()), ('read:all', NOW(), NOW());

CREATE TABLE "role_permissions" (
  "role_id" bigint NOT NULL,
  "permission_id" bigint NOT NULL,
  PRIMARY KEY ("role_id", "permission_id"),
  CONSTRAINT "fk_role_permissions_role" FOREIGN KEY ("role_id") REFERENCES "roles" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_role_permissions_permission" FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);

-- Seed role_permissions
INSERT INTO "role_permissions" (role_id, permission_id) VALUES
((SELECT id from roles where name = 'admin'), (SELECT id from permissions where name = 'create:all')),
((SELECT id from roles where name = 'admin'), (SELECT id from permissions where name = 'read:all')),
((SELECT id from roles where name = 'viewer'), (SELECT id from permissions where name = 'read:all'));

CREATE TABLE "services" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NULL,
  "description" text NULL,
  PRIMARY KEY ("id")
);
CREATE INDEX "idx_services_deleted_at" ON "services" ("deleted_at");
CREATE UNIQUE INDEX "idx_services_name" ON "services" ("name");
CREATE TABLE "service_versions" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "service_id" bigint NULL,
  "version" int NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_services_versions" FOREIGN KEY ("service_id") REFERENCES "services" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
CREATE INDEX "idx_service_versions_deleted_at" ON "service_versions" ("deleted_at"); 
CREATE UNIQUE INDEX "idx_service_versions_service_id_version" ON "service_versions" ("service_id", "version");



-- Insert admin user
INSERT INTO users (email, password, name) VALUES
('admin@kong.org', '$2a$10$X1iBADl5F6qd1nLZPa9E8uB8k.4zL6qhCdksskQhK0zN7DqoOlZ/6', 'Admin');

INSERT INTO user_roles (user_id, role_id) VALUES
((SELECT id from users where email = 'admin@kong.org'), (SELECT id from roles where name = 'admin'));
