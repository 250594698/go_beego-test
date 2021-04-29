<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" href="//res.layui.com/layui/dist/css/layui.css"  media="all">
  <script src="//res.layui.com/layui/dist/layui.js" charset="utf-8"></script>
</head>

<body>
  <header>
    <div style="width: 500px; margin: 0 auto; margin-top: 50px;">
        <h3>用户例表</h3>
        <p><a href="/user/add">添加</a></p>
        <table class="imagetable" >
          <tr>
            <th>序号</th><th>名称</th><th>联系电话</th><th>地址</th><th>操作</th>
          </tr>
          {{range $key, $val := .list_data}}
          <tr>
            <td>{{$val.Id}}</td><td>{{$val.Name}}</td><td>{{$val.Mobile}}</td>
            <td>
              {{if $val.Address_desc}}
                {{$val.Address_desc}}
              {{else}}
                 无
              {{end}}

            </td>
            <td>
              <a href="/user/edit/{{$val.Id}}">编辑</a>
              <a href="/user/del/{{$val.Id}}">删除</a>
            </td>
          </tr>
          {{end}}
        </table>
    </div>
    <div id="page">
      {{range $key, $val := .pagelist}}
        <a href="/user/index/{{$val}}">{{$val}}</a>
      {{end}}
      <span>总页数: {{.pageCount}}</span>
    </div>
  </header>
  <style>
    #page{
      width: 500px;
      height: 30px;
      line-height: 30px;
      border: 1px solid;
      margin:0 auto;
      text-align: center;
    }
    #page a{
      width: 25px;
      color: #000;
      display: inline-block;

    }
    #page span{
      width: 100px;
      display: inline-block;
      color:#999;

    }
    table{
      width: 500px;
      border:1px solid;
      margin: 0 auto;
    }
    table tr td{
      width: 100px;
    }
    table.imagetable {
      font-family: verdana,arial,sans-serif;
      font-size:11px;
      color:#333333;
      border-width: 1px;
      border-color: #999999;
      border-collapse: collapse;
    }
    table.imagetable th {
      background:#b5cfd2;
      border-width: 1px;
      padding: 8px;
      border-style: solid;
      border-color: #999999;
    }
    table.imagetable td {
      border-width: 1px;
      padding: 8px;
      border-style: solid;
      border-color: #999999;
    }
  </style>
  <script src="/static/js/reload.min.js"></script>
</body>
</html>
