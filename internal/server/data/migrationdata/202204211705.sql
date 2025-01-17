CREATE TABLE migrations (id VARCHAR(255) PRIMARY KEY);
INSERT INTO migrations VALUES('SCHEMA_INIT');
INSERT INTO migrations VALUES('202203231621');
INSERT INTO migrations VALUES('202203241643');
INSERT INTO migrations VALUES('202203301642');
INSERT INTO migrations VALUES('202203301652');
INSERT INTO migrations VALUES('202203301643');
INSERT INTO migrations VALUES('202203301644');
INSERT INTO migrations VALUES('202203301645');
INSERT INTO migrations VALUES('202203301646');
INSERT INTO migrations VALUES('202203301647');
INSERT INTO migrations VALUES('202203301648');
INSERT INTO migrations VALUES('202204061643');
INSERT INTO migrations VALUES('202204111503');
INSERT INTO migrations VALUES('202204181613');

CREATE TABLE `providers` (`id` integer,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,`name` text,`url` text,`client_id` text,`client_secret` text,`created_by` integer,PRIMARY KEY (`id`));

CREATE TABLE `settings` (`id` integer,`created_at` datetime,`updated_at` datetime,`deleted_at` datetime,`private_jwk` blob,`public_jwk` blob,"signup_required" numeric, `signup_enabled` numeric,PRIMARY KEY (`id`));
INSERT INTO settings VALUES(36869170022129664,'2022-04-12 17:44:55.129594012+00:00','2022-04-12 17:44:55.129594012+00:00',NULL,X'7b22757365223a22736967222c226b7479223a224f4b50222c226b6964223a22706f66744a5a75365a6b617a45684a54336a5350587a7461635f49715468443470786965426d5a593834343d222c22637276223a2245643235353139222c22616c67223a2245443235353139222c2278223a22376d426b6a474b3463626b64506f6b70765465376b6a372d4e556a4434467a4b46634234737742374c3867222c2264223a227456556257634a66724b7246636e4e74396a6a55687a355935737a3947315472564948434539366b727949227d',X'7b22757365223a22736967222c226b7479223a224f4b50222c226b6964223a22706f66744a5a75365a6b617a45684a54336a5350587a7461635f49715468443470786965426d5a593834343d222c22637276223a2245643235353139222c22616c67223a2245443235353139222c2278223a22376d426b6a474b3463626b64506f6b70765465376b6a372d4e556a4434467a4b46634234737742374c3867227d',0,NULL);
