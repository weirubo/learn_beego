<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Register</title>
</head>
<body>
<form method="post" action="http://localhost:8080/v1/user/add" enctype="multipart/form-data">
    用户名：<input type="text" name="username"><br>
    年龄：<input type="text" name="age"><br>
    性别：<input type="text" name="Gender"><br>
    头像：<input type="file" name="avatar"><br>
    <input type="submit" value="注册">
</form>
</body>
</html>