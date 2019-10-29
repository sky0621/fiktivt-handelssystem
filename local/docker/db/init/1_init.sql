
CREATE TABLE IF NOT EXISTS item (
  id varchar(36) NOT NULL,
  name varchar(256) NOT NULL,
  price bigint,
  item_holder_id varchar(36) NOT NULL,
  create_user varchar(256) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  update_user varchar(256) DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  deleted_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS item_holder (
  id varchar(36) NOT NULL,
  first_name varchar(256) NOT NULL,
  last_name varchar(256) NOT NULL,
  nickname varchar(256) DEFAULT NULL,
  create_user varchar(256) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  update_user varchar(256) DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  deleted_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS item_holder_relation (
  item_id varchar(36) NOT NULL,
  item_holder_id varchar(36) NOT NULL,
  create_user varchar(256) DEFAULT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  update_user varchar(256) DEFAULT NULL,
  updated_at timestamp NULL DEFAULT NULL,
  deleted_at timestamp NULL DEFAULT NULL,
  PRIMARY KEY (item_id, item_holder_id)
);
