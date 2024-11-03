INSERT INTO user(username, password, first_name, last_name)
VALUES ('test', '123', '3', '4'),
       ('test2', '123', '3', '4'),
       ('test3', '123', '3', '4'),
       ('test4', '123', '3', '4'),
       ('test5', '123', '3', '4'),
       ('test6', '123', '3', '4'),
       ('test7', '123', '3', '4'),
       ('test8', '123', '3', '4');

-- ข้อมูลตัวอย่างสำหรับตาราง workspace
INSERT INTO workspace (name, description,owner)
VALUES ('Workspace A', 'This is the first workspace.','test'),
       ('Workspace B', 'This workspace is focused on project management.','test'),
       ('Workspace C', 'This workspace is for development and testing.','test'),
       ('Workspace D', 'This workspace contains marketing materials.','test'),
       ('Workspace E', 'This is a shared workspace for the team.','test');


-- ข้อมูลตัวอย่างสำหรับตาราง section
INSERT INTO section (workspace_id, name)
VALUES (1, 'Section 1'),
       (1, 'Section 2'),
       (2, 'Section 3'),
       (2, 'Section 4'),
       (3, 'Section 5'),
       (4, 'Section 6'),
       (5, 'Section 7');

-- ข้อมูลตัวอย่างสำหรับตาราง user
INSERT INTO user (username, password, first_name, last_name)
VALUES ('user1', 'password1', 'John', 'Doe'),
       ('user2', 'password2', 'Jane', 'Smith'),
       ('user3', 'password3', 'Alice', 'Johnson'),
       ('user4', 'password4', 'Bob', 'Brown');

-- ข้อมูลตัวอย่างสำหรับตาราง user_workspace
INSERT INTO user_workspace (username, workspace_id)
VALUES ('user1', 1),
       ('user1', 2),
       ('user2', 2),
       ('user2', 3),
       ('user3', 1),
       ('user3', 4),
       ('user4', 5),
       ('user4', 3);
