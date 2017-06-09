<html>
    <head>
        <title>上传文件</title>
        <meta charset="utf-8">
    </head>
    <body>
    <form enctype="multipart/form-data" action="http://127.0.0.1:9090/upload" method="post">
      <input type="file" name="uploadfile" />
      <input type="hidden" name="token" value="{{.Token}}"/>
      <input type="submit" value="upload" />
    </form>
    </body>
</html>