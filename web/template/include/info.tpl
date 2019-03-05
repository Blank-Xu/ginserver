{{define "title"}}Info{{end}}
{{define "css"}}{{end}}
{{define "content"}}
    <div class="box box-primary">
        <div class="box-header with-border">
            <h3 class="box-title">Info</h3>
        </div>
        <form role="form" id="frmInfo" data-toggle="validator" class="form-horizontal" action="/admin/info"
              method="post">
            {{with .MainUser}}
                <div class="box-body">
                    <div class="form-group has-feedback">
                        <label for="nickname" class="col-sm-2 control-label">Nickname</label>
                        <div class="col-sm-10">
                            <input type="text" required class="form-control" name="nickname" placeholder="Nickname"
                                   value="{{.Nickname}}">
                            <span class="glyphicon glyphicon-user form-control-feedback"></span>
                            <div class="help-block with-errors"></div>
                        </div>
                    </div>
                    <div class="form-group has-feedback">
                        <label for="email" class="col-sm-2 control-label">Email</label>
                        <div class="col-sm-10">
                            <input type="text" class="form-control" name="email" placeholder="Email" value="{{.Email}}">
                            <span class="glyphicon glyphicon-envelope form-control-feedback"></span>
                            <div class="help-block with-errors"></div>
                        </div>
                    </div>
                    <div class="form-group has-feedback">
                        <label for="phone" class="col-sm-2 control-label">Phone</label>
                        <div class="col-sm-10">
                            <input type="text" class="form-control" name="phone" placeholder="Phone" value="{{.Phone}}">
                            <span class="glyphicon glyphicon-phone form-control-feedback"></span>
                            <div class="help-block with-errors"></div>
                        </div>
                    </div>
                    <div class="form-group has-feedback">
                        <label for="last_login" class="col-sm-2 control-label">LastLogin</label>
                        <div class="col-sm-10 text-center" style="text-align: center">
                            <span class="col-sm-5">IP: {{.LoginIp}}</span>
                            <span class="col-sm-5">Time: {{.LoginTime}}</span>
                        </div>
                    </div>
                </div>
                <div class="box-footer">
                    <button type="submit" class="btn btn-primary pull-right">Submit</button>
                </div>
            {{end}}
        </form>
    </div>
{{end}}
{{define "js"}}
    <script type="text/javascript">
        var options = {
            success: function () {
                toastr.success('modify info success', 'success');
            },
            error: function (xhr) {
                if (xhr) {
                    toastr.error(xhr.responseText, 'error');
                }
            },
        };
        $('#frmInfo').ajaxForm(options);
    </script>
{{end}}