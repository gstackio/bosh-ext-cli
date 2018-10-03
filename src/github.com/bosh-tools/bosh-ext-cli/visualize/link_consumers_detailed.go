package visualize

const linkConsumersDetailedTemplate string = `
<div id="linksConsumersDetailedLoadingSpinner">
 <div class="row">
  <div class="col-md-6 offset-md-3" style="text-align: center;">
   <i class="fa fa-circle-o-notch fa-spin" style="font-size: 7rem"></i>
  </div>
 </div>
</div>

<div id="linksConsumersDetailedContents" class="invisible">
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
    <table id="linkConsumersDetailedTable" class="link-consumers-table table table-hover table-bordered" width="100%">
	</table>
</div>
</div>
`
