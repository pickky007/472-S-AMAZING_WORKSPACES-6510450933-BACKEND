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
INSERT INTO workspace (id,name, description,owner)
VALUES ('d296ec','Workspace A', 'This is the first workspace.','test'),
       ('ac5b1e','Workspace B', 'This workspace is focused on project management.','test'),
       ('7c2a5c','Workspace C', 'This workspace is for development and testing.','test'),
       ('4511bf','Workspace D', 'This workspace contains marketing materials.','test'),
       ('ec515f','Workspace E', 'This is a shared workspace for the team.','test');


-- ข้อมูลตัวอย่างสำหรับตาราง section
INSERT INTO section (workspace_id, name)
VALUES ('d296ec', 'Section 1'),
       ('d296ec', 'Section 2'),
       ('ac5b1e', 'Section 3'),
       ('ac5b1e', 'Section 4'),
       ('7c2a5c', 'Section 5'),
       ('4511bf', 'Section 6'),
       ('ec515f', 'Section 7');

-- ข้อมูลตัวอย่างสำหรับตาราง user
INSERT INTO user (username, password, first_name, last_name)
VALUES ('user1', 'password1', 'John', 'Doe'),
       ('user2', 'password2', 'Jane', 'Smith'),
       ('user3', 'password3', 'Alice', 'Johnson'),
       ('user4', 'password4', 'Bob', 'Brown');

-- ข้อมูลตัวอย่างสำหรับตาราง user_workspace
INSERT INTO user_workspace (username, workspace_id)
VALUES ('user1', 'd296ec'),
       ('user1', 'ac5b1e'),
       ('user2', 'ac5b1e'),
       ('user2', '7c2a5c'),
       ('user3', 'd296ec'),
       ('user3', '4511bf'),
       ('user4', 'ec515f'),
       ('user4', '7c2a5c');
