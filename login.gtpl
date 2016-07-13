<html>
<head>
<title></title>
</head>
<body>
<form action="/login?username=mushroom" method="post">
    用户名:<input type="text" name="username">
    密码:<input type="password" name="password">
    实名:<input type="text" name="realname">
    年龄:<input type="text" name="age">
    Email:<input type="text" name="email">
    <select name="fruit">
    <option value="apple">apple</option>
    <option value="pear">pear</option>
    <option value="banana">banana</option>
    </select>
    <input type="radio" name="gender" value="1">Man
    <input type="radio" name="gender" value="2">Woman
    <input type="checkbox" name="interest" value="football">Football
    <input type="checkbox" name="interest" value="basketball">Basketball
    <input type="checkbox" name="interest" value="tennis">Tennis
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="登陆">
</form>
</body>
</html>