<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>
  <header>
    <div style="width: 500px; height: 300px; margin: 0 auto; margin-top: 50px;">
      <form action="/user/adddo" method="post">
      <table>
        <tr>
          <input type="hidden" name="id" value="{{.list_data.Id}}"/>
          <td>名称:</td><td><input type="text" name="name" value="{{.list_data.Name}}"/></td>
        </tr>
        <tr>
          <td>电话:</td><td><input type="text" name="mobile" value="{{.list_data.Mobile}}"/></td>
        </tr>
      </table>
        <p><input type="submit"  value="提交"/></p>
      </form>
    </div>
  </header>
  <style>
    table{
      width: 500px;
      border:1px solid #999;
      margin: 0 auto;
    }
    table tr td{
      width: 100px;
    }
  </style>
  <script src="/static/js/reload.min.js"></script>
</body>
</html>
