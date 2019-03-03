<div id="iam-usermgr-roleset-alert" class="alert hide"></div>
    
<div id="iam-usermgr-roleset" class="form-horizontal">
  <input type="hidden" name="roleid" value="{[=it.id]}">
    
  <div class="iam-form-group-title">Role Information</div>

  <table class="iam-formtable">
    <tbody>
    <tr>
      <td width="200px">Name</label>
      <td>
        <input type="text" class="form-control" name="name" value="{[=it.name]}">
      </td>
    </tr>

    <tr>
      <td>Description</label>
      <td>
        <input type="text" class="form-control" name="desc" value="{[=it.desc]}">
      </td>
    </tr>

    <tr>
      <td>Status</label>
      <td>
        {[~it._statusls :v]}
          <span class="iam-form-checkbox">
            <input type="radio" name="status" value="{[=v.status]}" {[ if (v.status == it.status) { ]}checked="checked"{[ } ]}> {[=v.title]}
          </span>
        {[~]}
      </td>
    </tr>

    </tbody>
  </table>
</div>

