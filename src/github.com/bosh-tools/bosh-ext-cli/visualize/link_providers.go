package visualize

const linkProvidersTemplate string = `
<div id="linksProvidersLoadingSpinner">
 <div class="row">
  <div class="col-md-6 offset-md-3" style="text-align: center;">
   <i class="fa fa-circle-o-notch fa-spin" style="font-size: 7rem"></i>
  </div>
 </div>
</div>

<div id="linksProvidersContents" class="invisible">
<div class="row">
	<div class="col-md-4">
	 <div class="form-group">
	    <label for="unit-type-select">Select a Deployment:</label>
	    <select class="form-control" id="deployment-select" data-placeholder="Choose a deployment">
	    </select>
	 </div>
	</div>
</div>
<div class="row">
    <table id="linkProvidersTable" class="link-providers-table table table-striped table-hover table-bordered" width="100%">
	</table>
</div>
</div>
`
