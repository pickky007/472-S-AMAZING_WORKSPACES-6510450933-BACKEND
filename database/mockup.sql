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
INSERT INTO workspace (id, name, description, owner)
VALUES ('d296ec', 'Workspace A', 'This is the first workspace.', 'test'),
       ('ac5b1e', 'Workspace B', 'This workspace is focused on project management.', 'test'),
       ('7c2a5c', 'Workspace C', 'This workspace is for development and testing.', 'test'),
       ('4511bf', 'Workspace D', 'This workspace contains marketing materials.', 'test'),
       ('ec515f', 'Workspace E', 'This is a shared workspace for the team.', 'test');


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


-- ข้อมูลตัวอย่างสำหรับตาราง activity
INSERT INTO activity (name, description, start_date, end_date, section_id, workspace_id)
VALUES
    ('Activity 1', 'Description for Activity 1', '2024-11-01 08:00:00', '2024-11-05 17:00:00', 1, 'd296ec'),   -- Section 1
    ('Activity 2', 'Description for Activity 2', '2024-11-02 09:00:00', '2024-11-06 16:00:00', 1, 'd296ec'),   -- Section 1
    ('Activity 3', 'Description for Activity 3', '2024-11-03 10:00:00', '2024-11-07 15:00:00', 2, 'd296ec'),   -- Section 2
    ('Activity 4', 'Description for Activity 4', '2024-11-04 11:00:00', '2024-11-08 14:00:00', 2, 'd296ec'),   -- Section 2
    ('Activity 5', 'Description for Activity 5', '2024-11-05 12:00:00', '2024-11-09 13:00:00', 3, 'ac5b1e'),   -- Section 3
    ('Activity 6', 'Description for Activity 6', '2024-11-06 13:00:00', '2024-11-10 12:00:00', 3, 'ac5b1e'),   -- Section 3
    ('Activity 7', 'Description for Activity 7', '2024-11-07 14:00:00', '2024-11-11 11:00:00', 4, 'ac5b1e'),   -- Section 4
    ('Activity 8', 'Description for Activity 8', '2024-11-08 15:00:00', '2024-11-12 10:00:00', 4, 'ac5b1e'),   -- Section 4
    ('Activity 9', 'Description for Activity 9', '2024-11-09 16:00:00', '2024-11-13 09:00:00', 5, '7c2a5c'),   -- Section 5
    ('Activity 10', 'Description for Activity 10', '2024-11-10 17:00:00', '2024-11-14 08:00:00', 6, '4511bf'), -- Section 6
    ('Activity 11', 'Description for Activity 11', '2024-11-11 18:00:00', '2024-11-15 07:00:00', 7, 'ec515f'); -- Section 7


INSERT INTO activity (name, description, start_date, end_date, section_id, workspace_id)
VALUES ('New Activity', 'Description for New Activity', '2024-11-04 07:00:00', '2024-11-10 07:00:00', 3, 'ac5b1e');
