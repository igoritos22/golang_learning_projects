<html>
    <head>
        <title>Login page</title>
    <head>
    <body>
        <h1>My GoLang Portal - v1</h1>
        <form action="http://127.0.0.1:9090/login" method="post">
            Selecione o seu perfil:
            <input type="checkbox" name="interest" value="futebol"> Futebol
            <input type="checkbox" name="interest" value="jiujitsu"> Jiu-Jitsu
            <input type="checkbox" name="interest" value="mma"> MMA <br><br><br>
            Username: <input type="text" name="username" placeholder="Username"><br><br>
            Password: <input type="password" name="password" placeholder="User Password"><br><br><br>
            <input type="hidden" name="token" value="{{.}}">
            <input type="submit" value="Login">
        </form>
        <script>alert("hello")</script>
    </body>
</html>