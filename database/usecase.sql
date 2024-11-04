SELECT w.id, w.name, w.description, w.owner
FROM workspace AS w
         INNER JOIN user_workspace AS uw ON w.id = uw.workspace_id
WHERE uw.username = 'user1'