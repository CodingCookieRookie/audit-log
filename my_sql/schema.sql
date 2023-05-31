CREATE DATABASE audit_log;

USE audit_log;

CREATE TABLE events (
    event_id int(11) NOT NULL AUTO_INCREMENT,
    event_type varchar(50) NOT NULL,
    event_time_stamp bigint(20),
    event_data_json varchar(6000),
    user_email varchar(320),
    PRIMARY KEY(event_id), INDEX (user_email, event_type, event_time_stamp), INDEX(user_email, event_time_stamp)
) 