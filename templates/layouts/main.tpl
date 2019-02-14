<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>{{.AppName}}-{{template "title" .}}</title>
    <meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">
    <link rel="stylesheet" href="/statics/bower_components/bootstrap/dist/css/bootstrap.min.css">
    <link rel="stylesheet" href="/statics/bower_components/font-awesome/css/font-awesome.min.css">
    <link rel="stylesheet" href="/statics/bower_components/Ionicons/css/ionicons.min.css">
    <link rel="stylesheet" href="/statics/dist/css/AdminLTE.min.css">
    <link rel="stylesheet" href="/statics/dist/css/skins/skin-blue.min.css">
    <link rel="stylesheet" href="/statics/plugins/pace/pace.min.css">
    {{template "css" .}}
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Source+Sans+Pro:300,400,600,700,300italic,400italic,600italic">
</head>

<body class="hold-transition skin-blue sidebar-mini">
<div class="wrapper">
    <header class="main-header">
        <a href="#" class="logo">
            <!-- mini logo for sidebar mini 50x50 pixels -->
            <span class="logo-mini"><b>gin</b></span>
            <!-- logo for regular state and mobile devices -->
            <span class="logo-lg"><b>gin</b>Server</span>
        </a>
        <!-- Header Navbar -->
        <nav class="navbar navbar-static-top" role="navigation">
            <!-- Sidebar toggle button-->
            <a href="#" class="sidebar-toggle" data-toggle="push-menu" role="button">
                <span class="sr-only">Toggle navigation</span>
            </a>
            <!-- Navbar Right Menu -->
            <div class="navbar-custom-menu">
                <ul class="nav navbar-nav">
                    <li>
                        <a href="/admin">
                            <i class="fa fa-home" style="font-size:14px;"></i>
                        </a>
                    </li>
                    {{with .main_user}}
                        <li class="dropdown user user-menu">
                            <!-- Menu Toggle Button -->
                            <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                                <!-- The user image in the navbar-->
                                <img src="{{.Icon}}" class="user-image" alt="User Image">
                                <!-- hidden-xs hides the username on small devices so only the image appears. -->
                                <span class="hidden-xs">{{.Username}}</span>
                            </a>
                            <ul class="dropdown-menu">
                                <!-- The user image in the menu -->
                                <li class="user-header">
                                    <img src="{{.Icon}}" class="img-circle" alt="User Image">
                                    <p>{{.Username}}
                                        <small>{{.Nickname}}</small>
                                    </p>
                                </li>
                                {{if .Email}}
                                    <!-- Menu Body -->
                                    <li class="user-body">
                                        <div class="row">
                                            <div class="col-xs-12 text-center">
                                                <span>email:
                                                    <b><a href="mailto:{{.Email}}">{{.Email}}</a></b>
                                                </span>
                                            </div>
                                        </div>
                                    </li>
                                {{end}}
                                <!-- Menu Footer-->
                                <li class="user-footer">
                                    <div class="pull-left">
                                        <a href="/admin/info" class="btn btn-default btn-flat">Info</a>
                                    </div>
                                    <div class="pull-right">
                                        <a href="/admin/logout" class="btn btn-default btn-flat">Sign Out</a>
                                    </div>
                                </li>
                            </ul>
                        </li>
                        <!-- Control Sidebar Toggle Button -->
                    {{end}}
                </ul>
            </div>
        </nav>
    </header>
    <!-- Left side column. contains the logo and sidebar -->
    <aside class="main-sidebar">
        <!-- sidebar: style can be found in sidebar.less -->
        <section class="sidebar">
            <form action="#" method="get" class="sidebar-form">
                <div class="input-group">
                    <input type="text" name="q" class="form-control" placeholder="Search...">
                    <span class="input-group-btn">
                        <button type="submit" name="search" id="search-btn" class="btn btn-flat">
                            <i class="fa fa-search"></i>
                        </button>
                    </span>
                </div>
            </form>
            <!-- Sidebar Menu -->
            <ul class="sidebar-menu" data-widget="tree">
                {{range .main_menu}}
                    {{if .List}}
                        <li class="treeview">
                            <a href="#">
                                <i class="fa{{if .Icon}} {{.Icon}}{{else}} fa-link{{end}}"></i>
                                <span>{{.Name}}</span>
                                <span class="pull-right-container">
                                    <i class="fa fa-angle-left pull-right"></i>
                                </span>
                            </a>
                            <ul class="treeview-menu">
                                {{range .List}}
                                    {{if .List}}
                                        <li id="menu{{.Id}}" class="treeview">
                                            <a href="#">
                                                <i class="fa{{if .Icon}} {{.Icon}}{{else}} fa-link{{end}}"></i>
                                                <span>{{.Name}}</span>
                                                <span class="pull-right-container">
                                                    <i class="fa fa-angle-left pull-right"></i>
                                                </span>
                                            </a>
                                            <ul class="treeview-menu">
                                                {{range .List}}
                                                    <li id="menu{{.Id}}"{{if eq $.active_path .Path}} class="active"{{end}}>
                                                        <a href="{{.Path}}">{{.Name}}</a>
                                                    </li>
                                                {{end}}
                                            </ul>
                                        </li>
                                    {{else}}
                                        <li id="menu{{.Id}}"{{if eq $.active_path .Path}} class="active"{{end}}>
                                            <a href="{{.Path}}">
                                                <i class="fa{{if .Icon}} {{.Icon}}{{else}} fa-link{{end}}"></i>
                                                <span>{{.Name}}</span>
                                            </a>
                                        </li>
                                    {{end}}
                                {{end}}
                            </ul>
                        </li>
                    {{else}}
                        <li id="menu{{.Id}}"{{if eq $.active_path .Path}} class="active"{{end}}>
                            <a href="{{.Path}}">
                                <i class="fa{{if .Icon}} {{.Icon}}{{else}} fa-link{{end}}"></i>
                                <span>{{.Name}}</span>
                            </a>
                        </li>
                    {{end}}
                {{end}}
            </ul>
            <!-- /.sidebar-menu -->
        </section>
        <!-- /.sidebar -->
    </aside>
    <!-- Content Wrapper. Contains page content -->
    <div class="content-wrapper">
        <!-- Content Header (Page header) -->
        <section class="content-header">
            <h1>Page Header
                <small>Optional description</small>
            </h1>
            <ol class="breadcrumb">
                <li><a href="#"><i class="fa fa-dashboard"></i> Level</a></li>
                <li class="active">Here</li>
            </ol>
        </section>
        <!-- Main content -->
        <section class="content container-fluid">
            {{template "content" .}}
        </section>
        <!-- /.content -->
    </div>
    <footer class="main-footer">
        <strong>
            Copyright &copy; 2019
            <a target="_blank" href="https://github.com/Blank-Xu/ginserver">
                ginserver
            </a>.
        </strong> All rights reserved.
    </footer>
    <div class="control-sidebar-bg"></div>
</div>
<script src="/statics/bower_components/jquery/dist/jquery.min.js"></script>
<script src="/statics/bower_components/bootstrap/dist/js/bootstrap.min.js"></script>
<script src="/statics/bower_components/PACE/pace.min.js"></script>
<script src="/statics/dist/js/adminlte.min.js"></script>
<script type="text/javascript">
    $(document).ajaxStart(function () {
        Pace.restart()
    })
</script>
<script src="/statics/bower_components/jquery-slimscroll/jquery.slimscroll.min.js"></script>
<script src="/statics/bower_components/fastclick/lib/fastclick.js"></script>
{{template "js" .}}
</body>
</html>