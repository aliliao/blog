
{{define "header"}}
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <!-- 上述3个meta标签*必须*放在最前面，任何其他内容都*必须*跟随其后！ -->

        <title>{{.Title}}</title>
        <link rel="shortcut icon" href="/static/img/favicon.ico">
        <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css">
        {{if or .IsSignIn .IsSignUp}}
            <!-- link href="https://v3.bootcss.com/examples/signin/signin.css" rel="stylesheet"-->
            <link href="/static/css/signin.css" rel="stylesheet">
        {{end}}
        <link rel="stylesheet" type="text/css" href="/static/css/my.comment.css">
        <!-- 评论显示
        <link rel='stylesheet' id='style-css'  href='https://ihacksoft.com/wp-content/themes/ihacksoft/style.css?ver=4.9.13' type='text/css' media='all' />
        -->

        <!-- sidebar style -->
        <link rel="stylesheet" type="text/css" href="/static/css/my.sidebar.css">

        <!-- blog theame -->
        <link rel="stylesheet" type="text/css" href="/static/css/my.theame.css">
    </head>
{{end}}

{{define "footer"}}
    <!-- ===== Bootstrap core JavaScript ======== -->
    <!-- Placed at the end of the document so the pages load faster -->
    <script src="http://cdn.staticfile.org/jquery/3.5.1/jquery.min.js"></script>
    <script src="/static/js/bootstrap.min.js"></script>
    {{if or .IsSignUp .IsSignIn}}
        <script src="/static/js/bootstrap.min.js"></script>
    {{end}}
{{end}}

{{define "navbar"}}
        <!-- 参考样式: https://v3.bootcss.com/examples/starter-template/#about -->
        <nav class="navbar navbar-inverse navbar-fixed-top">
            <div class="container">
                <div class="navbar-header">
                    <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
                        <span class="sr-only">Toggle navigation</span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                    </button>
                    <a class="navbar-brand" href="/home">我的博客</a>
                </div>
                <div id="navbar" class="collapse navbar-collapse">
                    <ul class="nav navbar-nav">
                        <li {{if .IsHome}} class="active" {{end}} ><a href="/">首页</a></li>
                        <li {{if .IsTopic}} class="active" {{end}} ><a href="/topic">文章</a></li>
                        <li {{if .IsCategory}} class="active" {{end}} ><a href="/category">分类</a></li>
                        <li {{if .IsMusic}} class="active" {{end}} ><a href="/music">音乐</a></li>
                        <li {{if .IsVideo}} class="active" {{end}} ><a href="/video">视频</a></li>
                        <li {{if .IsAbout}} class="active" {{end}} ><a href="/about">关于</a></li>
                    </ul>
                    <ul class="nav navbar-nav pull-right">
                        <!-- 右上角: 登录显示登录/注册 -->
                        {{if .IsLogin}}
                            <li ><a href="/signout">退出</a></li>
                        {{else}}
                            <li ><a href="/signup">注册</a></li>
                            <li ><a href="/signin">登录</a></li>
                        {{end}}
                    </ul>
                </div> <!--/.nav-collapse -->
            </div>
        </nav>
{{end}}