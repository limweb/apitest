/*
 Navicat Premium Data Transfer

 Source Server         : testdb-for-golang
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : main

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 12/03/2021 23:46:24
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for books
-- ----------------------------
DROP TABLE IF EXISTS "books";
CREATE TABLE "books" (
  "id" integer,
  "created_at" datetime,
  "updated_at" datetime,
  "deleted_at" datetime,
  "uid" text,
  "name" text,
  "author" text,
  "category" text,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Records of books
-- ----------------------------
INSERT INTO "books" VALUES (1, '2021-03-06 15:30:19.906554+07:00', '2021-03-06 15:30:19.906554+07:00', NULL, '', 'aaaa1', 'bbbb1', 'cccc1');
INSERT INTO "books" VALUES (2, '0001-01-01 00:00:00+00:00', '2021-03-08 23:40:49.3246544+07:00', NULL, 'yyyyyy2', 'yyyyyy2', '', '');
INSERT INTO "books" VALUES (3, '2021-03-06 15:30:20.1467557+07:00', '2021-03-09 00:24:25.611633+07:00', '2021-03-09 00:24:55.0554176+07:00', '', 'aaaaxxxx3', 'bbbbxxx3', 'ccccxxx3');
INSERT INTO "books" VALUES (4, '2021-03-06 17:22:10.9300905+07:00', '2021-03-06 17:22:10.9300905+07:00', NULL, '7b413da2-127a-4824-9749-13a0e4a50838', 'aaaa1', 'bbbb1', 'cccc1');
INSERT INTO "books" VALUES (5, '2021-03-06 17:22:11.1027328+07:00', '2021-03-06 17:22:11.1027328+07:00', NULL, '51d32fa9-6cac-41d8-a4d9-620d18244f78', 'aaaa2', 'bbbb2', 'cccc2');
INSERT INTO "books" VALUES (6, '2021-03-06 17:22:11.2083029+07:00', '2021-03-06 17:22:11.2083029+07:00', NULL, '29ea15c5-930c-4ed8-8d75-76de78faedb3', 'aaaa3', 'bbbb3', 'cccc3');
INSERT INTO "books" VALUES (7, '2021-03-09 00:25:31.8629789+07:00', '2021-03-09 00:25:31.8629789+07:00', NULL, 'a1100097-ff65-43b6-9a65-c30166e81ac5', 'Mr. Aaaa Bbbbb', 'Mr. Cccc Dddd', 'Sea');
INSERT INTO "books" VALUES (8, '2021-03-09 00:25:39.1645578+07:00', '2021-03-09 00:25:39.1645578+07:00', NULL, 'e3b4456b-586d-456b-bdfc-ca87ef8057de', 'Mr. Aaaa Bbbbb', 'Mr. Cccc Dddd', 'Sea');

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS "comments";
CREATE TABLE "comments" (
  "id" integer,
  "created_at" datetime,
  "updated_at" datetime,
  "deleted_at" datetime,
  "uid" text,
  "comment" text,
  "post_id" integer,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO "comments" VALUES (1, '2021-03-06 15:30:19.906554+07:00', '2021-03-06 15:30:19.906554+07:00', NULL, '', 'aaaa1', 1);
INSERT INTO "comments" VALUES (2, '0001-01-01 00:00:00+00:00', '2021-03-08 23:40:49.3246544+07:00', NULL, 'yyyyyy2', 'yyyyyy2', 1);
INSERT INTO "comments" VALUES (3, '2021-03-06 15:30:20.1467557+07:00', '2021-03-09 00:24:25.611633+07:00', '2021-03-09 00:24:55.0554176+07:00', '', 'aaaaxxxx3', 1);
INSERT INTO "comments" VALUES (4, '2021-03-06 17:22:10.9300905+07:00', '2021-03-06 17:22:10.9300905+07:00', NULL, '7b413da2-127a-4824-9749-13a0e4a50838', 'aaaa1', 3);
INSERT INTO "comments" VALUES (5, '2021-03-06 17:22:11.1027328+07:00', '2021-03-06 17:22:11.1027328+07:00', NULL, '51d32fa9-6cac-41d8-a4d9-620d18244f78', 'aaaa2', 2);
INSERT INTO "comments" VALUES (6, '2021-03-06 17:22:11.2083029+07:00', '2021-03-06 17:22:11.2083029+07:00', NULL, '29ea15c5-930c-4ed8-8d75-76de78faedb3', 'aaaa3', 2);
INSERT INTO "comments" VALUES (7, '2021-03-09 00:25:31.8629789+07:00', '2021-03-09 00:25:31.8629789+07:00', NULL, 'a1100097-ff65-43b6-9a65-c30166e81ac5', 'Mr. Aaaa Bbbbb', 3);
INSERT INTO "comments" VALUES (8, '2021-03-09 00:25:39.1645578+07:00', '2021-03-09 00:25:39.1645578+07:00', NULL, 'e3b4456b-586d-456b-bdfc-ca87ef8057de', 'Mr. Aaaa Bbbbb', 3);

-- ----------------------------
-- Table structure for password_resets
-- ----------------------------
DROP TABLE IF EXISTS "password_resets";
CREATE TABLE "password_resets" (
  "created_at" datetime,
  "updated_at" datetime,
  "email" text,
  "token" text
);

-- ----------------------------
-- Table structure for permission_role
-- ----------------------------
DROP TABLE IF EXISTS "permission_role";
CREATE TABLE "permission_role" (
  "id" integer,
  "created_at" datetime,
  "updated_at" datetime,
  "deleted_at" datetime,
  "uid" text,
  "permission_id" integer,
  "role_id" integer,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Records of permission_role
-- ----------------------------
INSERT INTO "permission_role" VALUES (1, '2021-02-10 11:11:11', '2021-02-10 11:11:11', '2021-02-10 11:11:11', NULL, 1, 1);
INSERT INTO "permission_role" VALUES (2, '2021-02-10 11:11:11', '2021-02-10 11:11:11', '2021-02-10 11:11:11', NULL, 2, 2);
INSERT INTO "permission_role" VALUES (3, '2021-02-10 11:11:11', '2021-02-10 11:11:11', '2021-02-10 11:11:11', NULL, 3, 3);
INSERT INTO "permission_role" VALUES (4, '2021-02-10 11:11:11', '2021-02-10 11:11:11', '2021-02-10 11:11:11', NULL, 4, 1);

-- ----------------------------
-- Table structure for permissions
-- ----------------------------
DROP TABLE IF EXISTS "permissions";
CREATE TABLE "permissions" (
  "id" integer,
  "created_at" datetime,
  "updated_at" datetime,
  "deleted_at" datetime,
  "uid" text,
  "name" text,
  "slug" text,
  "desc" text,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Records of permissions
-- ----------------------------
INSERT INTO "permissions" VALUES (1, NULL, NULL, NULL, NULL, 'P1', 'P1', 'P1');
INSERT INTO "permissions" VALUES (2, NULL, NULL, NULL, NULL, 'p2', 'p2', 'p2');
INSERT INTO "permissions" VALUES (3, NULL, NULL, NULL, NULL, 'p3', 'p3', 'p3');
INSERT INTO "permissions" VALUES (4, NULL, NULL, NULL, NULL, 'p4', 'p4', 'p4');

-- ----------------------------
-- Table structure for posts
-- ----------------------------
DROP TABLE IF EXISTS "posts";
CREATE TABLE "posts" (
  "id" integer,
  "created_at" datetime,
  "updated_at" datetime,
  "deleted_at" datetime,
  "uid" text,
  "title" text,
  "desc" text,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Records of posts
-- ----------------------------
INSERT INTO "posts" VALUES (1, '2021-03-06 15:30:19.906554+07:00', '2021-03-06 15:30:19.906554+07:00', NULL, '', 'aaaa1', 'bbbb1');
INSERT INTO "posts" VALUES (2, '0001-01-01 00:00:00+00:00', '2021-03-08 23:40:49.3246544+07:00', NULL, 'yyyyyy2', 'yyyyyy2', '');
INSERT INTO "posts" VALUES (3, '2021-03-06 15:30:20.1467557+07:00', '2021-03-09 00:24:25.611633+07:00', '2021-03-09 00:24:55.0554176+07:00', '', 'aaaaxxxx3', 'bbbbxxx3');
INSERT INTO "posts" VALUES (4, '2021-03-06 17:22:10.9300905+07:00', '2021-03-06 17:22:10.9300905+07:00', NULL, '7b413da2-127a-4824-9749-13a0e4a50838', 'aaaa1', 'bbbb1');
INSERT INTO "posts" VALUES (5, '2021-03-06 17:22:11.1027328+07:00', '2021-03-06 17:22:11.1027328+07:00', NULL, '51d32fa9-6cac-41d8-a4d9-620d18244f78', 'aaaa2', 'bbbb2');
INSERT INTO "posts" VALUES (6, '2021-03-06 17:22:11.2083029+07:00', '2021-03-06 17:22:11.2083029+07:00', NULL, '29ea15c5-930c-4ed8-8d75-76de78faedb3', 'aaaa3', 'bbbb3');
INSERT INTO "posts" VALUES (7, '2021-03-09 00:25:31.8629789+07:00', '2021-03-09 00:25:31.8629789+07:00', NULL, 'a1100097-ff65-43b6-9a65-c30166e81ac5', 'Mr. Aaaa Bbbbb', 'Mr. Cccc Dddd');
INSERT INTO "posts" VALUES (8, '2021-03-09 00:25:39.1645578+07:00', '2021-03-09 00:25:39.1645578+07:00', NULL, 'e3b4456b-586d-456b-bdfc-ca87ef8057de', 'Mr. Aaaa Bbbbb', 'Mr. Cccc Dddd');

-- ----------------------------
-- Table structure for role_users
-- ----------------------------
DROP TABLE IF EXISTS "role_users";
CREATE TABLE "role_users" (
  "id" integer,
  "created_at" datetime,
  "updated_at" datetime,
  "deleted_at" datetime,
  "uid" text,
  "user_id" integer,
  "role_id" integer,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Records of role_users
-- ----------------------------
INSERT INTO "role_users" VALUES (1, '2021-03-09 11:11:11', '2021-03-09 11:11:11', NULL, 1, 1, 1);
INSERT INTO "role_users" VALUES (2, '2021-03-09 11:11:11', '2021-03-09 11:11:11', NULL, 2, 1, 2);
INSERT INTO "role_users" VALUES (3, '2021-03-09 11:11:11', '2021-03-09 11:11:11', NULL, 3, 1, 3);

-- ----------------------------
-- Table structure for roles
-- ----------------------------
DROP TABLE IF EXISTS "roles";
CREATE TABLE "roles" (
  "id" integer,
  "created_at" datetime,
  "updated_at" datetime,
  "deleted_at" datetime,
  "uid" text,
  "name" text,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Records of roles
-- ----------------------------
INSERT INTO "roles" VALUES (1, '2021-03-09 11:11:11', '2021-03-09 11:11:11', NULL, 1, 'Admin');
INSERT INTO "roles" VALUES (2, '2021-03-09 11:11:11', '2021-03-09 11:11:11', NULL, 2, 'DataManager');
INSERT INTO "roles" VALUES (3, '2021-03-09 11:11:11', '2021-03-09 11:11:11', NULL, 3, 'Master');
INSERT INTO "roles" VALUES (4, '2021-03-09 11:11:11', '2021-03-09 11:11:11', NULL, 4, 'Agent');
INSERT INTO "roles" VALUES (5, '2021-03-09 11:11:11', '2021-03-09 11:11:11', NULL, 5, 'SubAgent');
INSERT INTO "roles" VALUES (6, '2021-03-09 11:11:11', '2021-03-09 11:11:11', NULL, 6, 'Guest');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "users";
CREATE TABLE "users" (
  "id" integer,
  "deleted_at" datetime,
  "created_at" datetime,
  "updated_at" datetime,
  "uid" text,
  "name" text,
  "username" text,
  "password" text,
  "email" text,
  "telephone" text,
  "email_verified_at" datetime,
  "level" integer,
  "remember_token" text,
  PRIMARY KEY ("id"),
  UNIQUE ("username" ASC),
  UNIQUE ("email" ASC),
  UNIQUE ("telephone" ASC)
);

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO "users" VALUES (1, NULL, '2021-03-07 14:06:18.0996549+07:00', '2021-03-07 14:06:18.0996549+07:00', '', 'Mr. ฺA ', 'usera', '$2a$14$Q/QNkHUHFabLAdJAFYcD/O8oKFM8WqszhvOD0ORZV7KaGgav4qwe6', 'a@email.com', '0816477729', NULL, 99, '6f327274-2ca6-4e05-86cb-02f4d9edfe0e');

-- ----------------------------
-- Indexes structure for table books
-- ----------------------------
CREATE INDEX "idx_book_deleted_at"
ON "books" (
  "deleted_at" ASC
);
CREATE INDEX "idx_books_deleted_at"
ON "books" (
  "deleted_at" ASC
);

-- ----------------------------
-- Indexes structure for table comments
-- ----------------------------
CREATE INDEX "idx_book_deleted_at_copy2"
ON "comments" (
  "deleted_at" ASC
);
CREATE INDEX "idx_books_deleted_at_copy2"
ON "comments" (
  "deleted_at" ASC
);
CREATE INDEX "idx_comments_deleted_at"
ON "comments" (
  "deleted_at" ASC
);

-- ----------------------------
-- Indexes structure for table permission_role
-- ----------------------------
CREATE INDEX "idx_permission_role_deleted_at"
ON "permission_role" (
  "deleted_at" ASC
);

-- ----------------------------
-- Indexes structure for table permissions
-- ----------------------------
CREATE INDEX "idx_premissions_deleted_at"
ON "permissions" (
  "deleted_at" ASC
);

-- ----------------------------
-- Indexes structure for table posts
-- ----------------------------
CREATE INDEX "idx_book_deleted_at_copy1"
ON "posts" (
  "deleted_at" ASC
);
CREATE INDEX "idx_books_deleted_at_copy1"
ON "posts" (
  "deleted_at" ASC
);
CREATE INDEX "idx_posts_deleted_at"
ON "posts" (
  "deleted_at" ASC
);

-- ----------------------------
-- Indexes structure for table role_users
-- ----------------------------
CREATE INDEX "idx_role_user_deleted_at"
ON "role_users" (
  "deleted_at" ASC
);

-- ----------------------------
-- Indexes structure for table roles
-- ----------------------------
CREATE INDEX "idx_roles_deleted_at"
ON "roles" (
  "deleted_at" ASC
);

-- ----------------------------
-- Indexes structure for table users
-- ----------------------------
CREATE INDEX "idx_users_deleted_at"
ON "users" (
  "deleted_at" ASC
);

PRAGMA foreign_keys = true;
