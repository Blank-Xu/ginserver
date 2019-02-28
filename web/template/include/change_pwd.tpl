{{define "title"}}Change Password{{end}}
{{define "css"}}{{end}}
{{define "content"}}
    <div class="box box-primary">
        <div class="box-header with-border">
            <h3 class="box-title">Change Password</h3>
        </div>
        <form role="form" id="frmChangePwd" data-toggle="validator" class="form-horizontal" action="/admin/change_pwd" method="post">
            <div class="box-body">
                <div class="form-group">
                    <label for="password" class="col-sm-2 control-label">Password</label>
                    <div class="col-sm-10">
                        <input type="password" required class="form-control" name="password" placeholder="Password">
                        <div class="help-block with-errors"></div>
                    </div>
                </div>
                <div class="form-group">
                    <label for="new_password" class="col-sm-2 control-label">New Password</label>
                    <div class="col-sm-10">
                        <input type="password" required class="form-control" name="new_password" placeholder="New Password">
                        <div class="help-block with-errors"></div>
                    </div>
                </div>
                <div class="form-group">
                    <label for="confirm_password" class="col-sm-2 control-label">Confirm Password</label>
                    <div class="col-sm-10">
                        <input type="password" required class="form-control" name="confirm_password"
                               placeholder="Confirm Password">
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
                alert("success");
            },
            error: function(xhr) {
                if (xhr) {
                    var resp = xhr.responseJSON;
                    alert("code: "+resp.code+"\nmessage: "+resp.msg);
                }
            },
        };
        $('#frmChangePwd').ajaxForm(options);
    </script>
{{end}}