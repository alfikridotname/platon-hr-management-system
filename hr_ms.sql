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

 Date: 28/06/2023 21:37:26
*/


-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "id" int4 NOT NULL,
  "email" varchar(255) COLLATE "pg_catalog"."default",
  "password" varchar(255) COLLATE "pg_catalog"."default",
  "active" bool,
  "role" varchar(255) COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO "public"."users" VALUES (1, 'alfikri.name@gmail.com', '$2y$10$fGdBTZrRPbwMCsjbo/LTqeqIFIsq9Fi.jBgAUbSB8Gj8gHCrUfTDS', 't', 'superadmin');

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");
