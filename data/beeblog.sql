create table `category`
-- --------------------------------------------------
--  Table Structure for `myAppNew/models.Category`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `category` (
                                          `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
                                          `title` varchar(255),
    `created` datetime,
    `views` bigint,
    `topic_time` datetime,
    `topic_count` bigint,
    `topic_last_user_id` bigint
    ) ENGINE=InnoDB;
CREATE INDEX `category_created` ON `category` (`created`);
    CREATE INDEX `category_views` ON `category` (`views`);
    CREATE INDEX `category_topic_time` ON `category` (`topic_time`);

create table `topic`
-- --------------------------------------------------
--  Table Structure for `myAppNew/models.Topic`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `topic` (
                                       `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
                                       `uid` bigint,
                                       `title` varchar(255),
    `content` varchar(5000),
    `attachment` varchar(255),
    `created` datetime,
    `updated` datetime,
    `views` bigint,
    `author` varchar(255),
    `replay_time` datetime,
    `reply_count` bigint,
    `reply_last_user_id` bigint
    ) ENGINE=InnoDB;
CREATE INDEX `topic_created` ON `topic` (`created`);
    CREATE INDEX `topic_updated` ON `topic` (`updated`);
    CREATE INDEX `topic_views` ON `topic` (`views`);
    CREATE INDEX `topic_replay_time` ON `topic` (`replay_time`);

ALTER TABLE `topic` ADD COLUMN `category` varchar(255)
