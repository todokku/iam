{{template "Common/HtmlHeader.tpl" .}}

<style>
body {
  margin: 0 auto;
  position: relative;
  font-size: 13px;
  font-family: Arial, sans-serif;
  background-color: #fff;
}

#ids-signup-box {
  width: 400px;
  position: absolute;
  left: 50%;
  top: 20%;
  margin-left: -200px;
  color: #555;
}

#ids-signup-form {
  background-color: #f7f7f7;
  border-radius: 4px;
  padding: 30px 30px 20px 30px;
  box-shadow: 0px 2px 2px 0px #999;
}

.ids-signup-msg01 {
  font-size: 20px;
  margin: 20px 0;
  text-align: center;
}

#ids-signup-form .ilf-group {
  padding: 0 0 10px 0; 
}

#ids-signup-form .ilf-input {
  display: block;
  width: 100%;
  height: 40px;
  padding: 5px 10px;
  font-size: 14px;
  line-height: 1.42857143;
  color: #555;
  background-color: #fff;
  background-image: none;
  border: 1px solid #ccc;
  border-radius: 2px;
  box-shadow: inset 0 1px 1px rgba(0, 0, 0, .075);
  transition: border-color ease-in-out .15s, box-shadow ease-in-out .15s;
}

#ids-signup-form .ilf-input:focus {
  border-color: #66afe9;
  outline: 0;
  box-shadow: inset 0 1px 1px rgba(0,0,0,.075), 0 0 8px rgba(102, 175, 233, .6);
}

#ids-signup-form .ilf-btn {
  width: 100%;
  height: 40px;
  display: inline-block;
  padding: 5px 10px;
  margin-bottom: 0;
  font-size: 14px;
  font-weight: normal;
  line-height: 1.42857143;
  text-align: center;
  white-space: nowrap;
  vertical-align: middle;
  cursor: pointer;
  -webkit-user-select: none;
     -moz-user-select: none;
      -ms-user-select: none;
          user-select: none;
  background-image: none;
  border: 1px solid transparent;
  border-radius: 3px;
  color: #fff;
  background-color: #427fed;
  border-color: #357ebd;
}

#ids-signup-form .ilf-btn:hover {
  color: #fff;
  background-color: #3276b1;
  border-color: #285e8e;
}

#ids-signup-box .ilb-signup {
  margin: 10px 0;
  text-align: center;
  font-size: 15px;
}
#ids-signup-box .ilb-signup a {
  color: #427fed;
}

#ids-signup-box .ilb-footer {
  text-align: center;
  margin: 20px 0;
  font-size: 14px;
}
#ids-signup-box .ilb-footer a {
  color: #777;
}
#ids-signup-box .ilb-footer img {
  width: 16px;
  height: 16px;
}
</style>

<div id="ids-signup-box">

  <div class="ids-signup-msg01">{{T . "Create your Account"}}</div>

  <form id="ids-signup-form" action="#">

    <input type="hidden" name="continue" value="{{.continue}}">

    <div id="ids-signup-form-alert" class="alert hide ilf-group"></div>

    {{if eq .user_reg_disable false }}
    <div class="ilf-group">
      <input type="text" class="ilf-input" name="name" value="{{.name}}" placeholder="{{T . "Your name"}}">
    </div>

    <div class="ilf-group">
      <input type="text" class="ilf-input" name="uname" value="{{.uname}}" placeholder="{{T . "Unique username"}}">
    </div>

    <div class="ilf-group">
      <input type="text" class="ilf-input" name="email" placeholder="{{T . "Email"}}">
    </div>

    <div class="ilf-group">
      <input type="password" class="ilf-input" name="passwd" placeholder="{{T . "Password"}}">
    </div>

    <div class="ilf-group">
      <button type="submit" class="ilf-btn">{{T . "Create Account"}}</button>
    </div>
    {{else}}
    <div class="alert alert-danger">User registration was closed!<br>Please contact the administrator to manually register accounts</div>
    {{end}}

  </form>

  <div class="ilb-signup">
    <a href="/ids/service/login?continue={{.continue}}">Sign in with your Account</a>
  </div>

  <div class="ilb-footer">
    <img src="/ids/~/ids/img/ids-s2-32.png"> 
    <a href="http://www.lessos.com" target="_blank">lessOS Identity Server</a>
  </div>

</div>

<script>

//
$("input[name=name]").focus();

//
var ids_eh = $("#ids-signup-box").height();
$("#ids-signup-box").css({
    "top": "40%",
    "margin-top": - (ids_eh / 2) + "px" 
});

//
$("#ids-signup-form").submit(function(event) {

    event.preventDefault();

    /* var req = {
        data: {
            "name": $("input[name=name]").val(),
            "email": $("input[name=email]").val(),
            "passwd": $("input[name=passwd]").val(),
            "continue": $("input[name=continue]").val(),
        }
    } */
    // console.log($("#ids-signup-form").serialize());
    
    $.ajax({
        type    : "POST",
        url     : "/ids/reg/sign-up-reg",
        data    : $("#ids-signup-form").serialize(),//JSON.stringify(req),
        timeout : 3000,
        //contentType: "application/json; charset=utf-8",
        success : function(data) {

            // if (err) {
            //     return l4i.InnerAlert("#ids-signup-form-alert", 'alert-danger', err);
            // }

            if (data.error) {
                return l4i.InnerAlert("#ids-signup-form-alert", 'alert-danger', data.error.message);
            }

            if (data.kind != "User") {
                return l4i.InnerAlert("#ids-signup-form-alert", 'alert-danger', "Unknown Error");
            }
                
            l4i.InnerAlert("#ids-signup-form-alert", 'alert-success', "Successfully registration. Page redirecting");
            $(".ilf-group").hide(600);

            window.setTimeout(function(){
                window.location = "/ids/service/login?continue={{.continue}}";
            }, 1500);

        },
        error: function(xhr, textStatus, error) {
            l4i.InnerAlert("#ids-signup-form-alert", 'alert-danger', '{{T . "Internal Server Error"}}');
        }
    });
});

</script>

{{template "Common/HtmlFooter.tpl" .}}
