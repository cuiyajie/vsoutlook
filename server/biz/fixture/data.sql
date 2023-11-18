INSERT INTO "public"."users" ("id", "name", "password", "role") VALUES
('12345678', 'admin', '', 1)

UPDATE users SET created_at = NOW() - interval '1 hour', updated_at = NOW();

INSERT INTO "public"."tmpl_types" ("id", "name", "icon") VALUES
('dZQaiivq', '编解码', '');
INSERT INTO "public"."tmpl_types" ("id", "name", "icon") VALUES
('KmPLa3DD', '上下变换', '');
INSERT INTO "public"."tmpl_types" ("id", "name", "icon") VALUES
('ClhlTlMA', '多画面', '');
INSERT INTO "public"."tmpl_types" ("id", "name", "icon") VALUES
('GAGss15a', '切换台', '');

INSERT INTO "public"."tmpls" ("id", "name", "type") VALUES
('5ccF5nPX', '编解码-1', 'dZQaiivq');
INSERT INTO "public"."tmpls" ("id", "name", "type") VALUES
('qhieIGA6', '上下变换-1', 'KmPLa3DD');
INSERT INTO "public"."tmpls" ("id", "name", "type") VALUES
('piGQeuZr', '多画面-1', 'ClhlTlMA');
INSERT INTO "public"."tmpls" ("id", "name", "type") VALUES
('7t1q2fIU', '切换台-1', 'GAGss15a');
