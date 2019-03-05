{{define "title"}}Change Password{{end}}
{{define "css"}}{{end}}
{{define "content"}}
    <div class="box box-primary">
        <div class="box-header with-border">
            <h3 class="box-title">Change Password</h3>
        </div>
        <form role="form" id="frmChangePwd" data-toggle="validator" class="form-horizontal" action="/admin/change_pwd" method="post">
            <div class="box-body">
                <div class="form-group has-feedback">
                    <label for="password" class="col-sm-2 control-label">Password</label>
                    <div class="col-sm-10">
                        <input type="password" required class="form-control" name="password" placeholder="Password">
                        <span class="glyphicon glyphicon-lock form-control-feedback"></span>
                        <div class="help-block with-errors"></div>
                    </div>
                </div>
                <div class="form-group has-feedback">
                    <label for="new_password" class="col-sm-2 control-label">New Password</label>
                    <div class="col-sm-10">
                        <input type="password" required class="form-control" name="new_password" placeholder="New Password">
                        <span class="glyphicon glyphicon-lock form-control-feedback"></span>
                        <div class="help-block with-errors"></div>
                    </div>
                </div>
                <div class="form-group has-feedback">
                    <label for="confirm_password" class="col-sm-2 control-label">Confirm Password</label>
                    <div class="col-sm-10">
                        <input type="password" required class="form-control" name="confirm_password" placeholder="Confirm Password">
                        <span class="glyphicon glyphicon-lock form-control-feedback"></span>
                        <div class="help-block with-errors"></div>
                    </div>
                </div>
            </div>
            <div class="box-footer">
                <button type="submit" class="btn btn-primary pull-right">Submit</button>
            </div>
        </form>
    </div>
{{end}}
{{define "js"}}
    <script type="text/javascript">
        var options = {
            success: function () {
                $('#frmChangePwd').resetForm();
                toastr.options.onHidden = function () {
                    window.location = '/admin/logout';
                };
                toastr.success('you must login again by your new password.', 'success');
            },
            error: function (xhr) {
                if (xhr) {
                    toastr.error(xhr.responseText, 'error');
                }
            },
        };
        $('#frmChangePwd').ajaxForm(options);
    </script>
{{end}}