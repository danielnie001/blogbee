{{define "navbar"}}
<a href="/" class="navbar-brand">我的博客</a>
<ul class="nav navbar-nav">
    <li {{if .IsHome}}class="active" {{end}}><a href="/">首页</a></li>
    <li {{if .IsCategory}}class="active" {{end}}><a href="/category">分类</a></li>
    <li {{if .IsTopic}}class="active" {{end}}><a href="/topic">文章</a></li>
</ul>
<ul class="nav navbar-nav navbar-right">
    {{if .IsLogin}}
    <li><a href="/login?exit=true">退出登录</a></li>
    {{else}}
    <li><a href="/login">管理员登录</a></li>
    {{end}}
</ul>
{{end}}