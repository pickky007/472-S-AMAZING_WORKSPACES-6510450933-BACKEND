-- User
CREATE TABLE user
(
    username   VARCHAR(256),
    password   VARCHAR(256),
    first_name VARCHAR(256) NOT NULL,
    last_name  VARCHAR(256) NOT NULL,
    PRIMARY KEY (username)
);

-- Workspace
CREATE TABLE workspace
(
    id          CHAR(6) PRIMARY KEY,
    name        VARCHAR(256) NOT NULL,
    description VARCHAR(1024),
    owner       VARCHAR(256) NOT NULL,

    FOREIGN KEY (owner) REFERENCES user (username)
);

-- User Workspace
CREATE TABLE user_workspace
(
    username     VARCHAR(256),
    workspace_id CHAR(6),

    PRIMARY KEY (username, workspace_id),
    FOREIGN KEY (username) REFERENCES user (username),
    FOREIGN KEY (workspace_id) REFERENCES workspace (id)
);

-- Section
CREATE TABLE section
(
    id           INT AUTO_INCREMENT,
    workspace_id CHAR(6),
    name         varchar(256) NOT NULL,

    PRIMARY KEY (id, workspace_id),
    FOREIGN KEY (workspace_id) REFERENCES workspace (id)
);

CREATE TABLE activity
(
    id           INT AUTO_INCREMENT,
    name         VARCHAR(256)  NOT NULL,
    owner        VARCHAR(256)  NOT NULL,
    description  varchar(1024) NOT NULL,
    start_date   DATETIME      NOT NULL,
    end_date     DATETIME      NOT NULL,
    section_id   INT           NOT NULL,
    workspace_id CHAR(6)       NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (section_id, workspace_id) REFERENCES section (id, workspace_id)
);
