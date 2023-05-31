CREATE DATABASE audit_log;

USE audit_log;

CREATE TABLE events (
    event_id int(11) NOT NULL AUTO_INCREMENT,
    event_type varchar(50) NOT NULL,
    event_time_stamp bigint(20),
    event_data_json varchar(6000),
    PRIMARY KEY(event_id), INDEX (event_type), INDEX(event_time_stamp)
) 