<style>
.page-header {
	margin: 10px 0;
	font-height: 100%;
}
</style>

<div class="page-header">
  <h2>General Settings <small></small></h2>
</div>

<div id="twiooa" class="alert hide"></div>

  <form id="akm4wj" class="form-horizontal" action="#">
   
    <label class="ids-form-group-title">Service Information</label>

    <div class="form-group">
      <label class="col-sm-2">Service Name</label>
      <div class="col-sm-4">
        <input type="text" class="form-control" name="service_name" value="{{.service_name}}">
      </div>
      <div class="col-sm-6">
        <div class="ids-callout ids-callout-primary">
          <h4>Identity Service Name</h4>
          <p>
            Will be used in the website name, e-mail signature ...
          </p>
        </div>
      </div>
    </div>

    <div class="form-group">
      <label class="col-sm-2">Banner title</label>
      <div class="col-sm-4">
        <input type="text" class="form-control" name="webui_banner_title" value="{{.webui_banner_title}}">
      </div>
      <div class="col-sm-6">
        <div class="ids-callout ids-callout-primary">
          <h4>Banner Title</h4>
          <p>
            Will be used in the website banner
          </p>
        </div>
      </div>
    </div>

    <!--
    <div class="form-group">
      <label class="col-sm-3">ICON</label>
      <div class="col-sm-3">
        <input id="service_info_icon" name="service_info_icon" size="20" type="file" class="form-control">
      </div>
      <div class="col-sm-6">
        <div class="ids-callout ids-callout-primary">
          <h4>Custom a Icon</h4>
          <p>
            Will be used in the site banner, browser shortcut icon ...
          </p>
        </div>
      </div>
    </div>
    -->

    <div class="form-group">
      <div class="col-sm-offset-2 col-sm-3">
        <button type="submit" class="btn btn-primary">{{T . "Save"}}</button>
      </div>
    </div>

  </form>

<script type="text/javascript">
$("#akm4wj").submit(function(event) {

    event.preventDefault(); 

    _ids_sysmgr_genset();
});

function _ids_sysmgr_genset()
{
    $.ajax({
        type    : "POST",
        url     : "/ids/sys-mgr/gen-set-save?_="+Math.random(),
        data    : $("#akm4wj").serialize(),
        timeout : 3000,
        success : function(rsp) {
          lessAlert("#twiooa", "alert-success", rsp);
        },
        error: function(xhr, textStatus, error) {
            lessAlert("#twiooa", "alert-danger", "Error: "+ xhr.responseText);
        }
    });
}

</script>