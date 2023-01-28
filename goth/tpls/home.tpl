
<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <title>Go Web</title>
</head>
<body>
    <p><a href="/logout/github">logout</a></p>
    <p>Name: {{.Name}} [{{.LastName}}, {{.FirstName}}]</p>
    <p>Email: {{.Email}}</p>
    <p>NickName: {{.NickName}}</p>
    <p>Location: {{.Location}}</p>
    <p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
    <p>Description: {{.Description}}</p>
    <p>UserID: {{.UserID}}</p>
    <p>AccessToken: {{.AccessToken}}</p>
    <p>ExpiresAt: {{.ExpiresAt}}</p>
    <p>RefreshToken: {{.RefreshToken}}</p>
</body>
</html>


