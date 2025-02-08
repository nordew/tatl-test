-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `auth` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `api-key` varchar(32) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `app_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `user_profile` (
  `user_id` bigint(20) NOT NULL,
  `first_name` varchar(32) NOT NULL,
  `last_name` varchar(64) NOT NULL,
  `phone` varchar(64) NOT NULL,
  `address` varchar(64) NOT NULL,
  `city` varchar(64) NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `user_data` (
  `user_id` bigint(20) NOT NULL,
  `school` varchar(32) NOT NULL,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +goose StatementEnd

-- +goose StatementBegin
INSERT IGNORE INTO `auth` VALUES
  (1, 'www-dfq92-sqfwf'),
  (2, 'ffff-2918-xcas');
-- +goose StatementEnd

-- +goose StatementBegin
INSERT IGNORE INTO `app_user` VALUES
  (1, 'test'),
  (2, 'admin'),
  (3, 'guest');
-- +goose StatementEnd

-- +goose StatementBegin
INSERT IGNORE INTO `user_data` VALUES
  (1, 'гімназія №179 міста Києва'),
  (2, 'ліцей №227'),
  (3, 'Медична гімназія №33 міста Києва');
-- +goose StatementEnd

-- +goose StatementBegin
INSERT IGNORE INTO `user_profile` VALUES
  (1, 'Олександр', 'Грішин', '+38050123455', 'вул. Сибірська 2', 'Київ'),
  (2, 'Дмитро', 'Кавун', '+38065133223', 'вул. Степана Бандери 4', 'Київ'),
  (3, 'Василь', 'Шпак', '+38055221166', 'вул. Сіверська 5', 'Житомир');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `user_profile`;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS `user_data`;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS `app_user`;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE IF EXISTS `auth`;
-- +goose StatementEnd
