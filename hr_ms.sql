/*
 Navicat Premium Data Transfer

 Source Server         : PLATON
 Source Server Type    : PostgreSQL
 Source Server Version : 140007
 Source Host           : localhost:5432
 Source Catalog        : hr_ms
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140007
 File Encoding         : 65001

 Date: 02/07/2023 20:02:48
*/


-- ----------------------------
-- Table structure for companies
-- ----------------------------
DROP TABLE IF EXISTS "public"."companies";
CREATE TABLE "public"."companies" (
  "id" int4 NOT NULL,
  "name" varchar(255) COLLATE "pg_catalog"."default",
  "status_id" int4 NOT NULL
)
;

-- ----------------------------
-- Records of companies
-- ----------------------------
INSERT INTO "public"."companies" VALUES (1, 'OLIN', 1);
INSERT INTO "public"."companies" VALUES (2, 'Others', 2);

-- ----------------------------
-- Table structure for company_status
-- ----------------------------
DROP TABLE IF EXISTS "public"."company_status";
CREATE TABLE "public"."company_status" (
  "id" int4 NOT NULL,
  "name" varchar(255) COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of company_status
-- ----------------------------
INSERT INTO "public"."company_status" VALUES (1, 'Active');
INSERT INTO "public"."company_status" VALUES (2, 'Freeze');
INSERT INTO "public"."company_status" VALUES (3, 'Overdue');

-- ----------------------------
-- Table structure for user_positions
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_positions";
CREATE TABLE "public"."user_positions" (
  "id" int4 NOT NULL,
  "name" varchar(255) COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of user_positions
-- ----------------------------
INSERT INTO "public"."user_positions" VALUES (1, 'Backend Developer');
INSERT INTO "public"."user_positions" VALUES (2, 'Frontend Developer');
INSERT INTO "public"."user_positions" VALUES (3, 'Salesman');

-- ----------------------------
-- Table structure for user_profiles
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_profiles";
CREATE TABLE "public"."user_profiles" (
  "id" int4 NOT NULL,
  "user_id" int4 NOT NULL,
  "fullname" varchar(255) COLLATE "pg_catalog"."default",
  "profile_picture" varchar(255) COLLATE "pg_catalog"."default",
  "type_id" int4 NOT NULL,
  "position_id" int4 NOT NULL
)
;

-- ----------------------------
-- Records of user_profiles
-- ----------------------------
INSERT INTO "public"."user_profiles" VALUES (1, 1, 'Alfikri', 'https://picsum.photos/seed/picsum/200/300', 1, 1);

-- ----------------------------
-- Table structure for user_status
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_status";
CREATE TABLE "public"."user_status" (
  "id" int4 NOT NULL,
  "name" varchar(255) COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of user_status
-- ----------------------------
INSERT INTO "public"."user_status" VALUES (1, 'Active');
INSERT INTO "public"."user_status" VALUES (2, 'Inactive');

-- ----------------------------
-- Table structure for user_types
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_types";
CREATE TABLE "public"."user_types" (
  "id" int4 NOT NULL,
  "name" varchar(255) COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of user_types
-- ----------------------------
INSERT INTO "public"."user_types" VALUES (1, 'Commercial');
INSERT INTO "public"."user_types" VALUES (2, 'Non Commercial');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "id" int4 NOT NULL,
  "email" varchar(255) COLLATE "pg_catalog"."default",
  "password" varchar(255) COLLATE "pg_catalog"."default",
  "company_id" int4 NOT NULL,
  "status_id" int4 NOT NULL,
  "roles" _int4 NOT NULL,
  "deleted_at" timestamp(6)
)
;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO "public"."users" VALUES (1, 'alfikri.name@gmail.com', '$2y$10$GgONTeY9GRTXqMYoWiV4JOdURND1NpWCXWx4fG7AiKj0SH5BWZKCa', 1, 1, '{1,2}', NULL);
