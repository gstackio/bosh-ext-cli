package visualize

const variablesTemplate string = `

<div id="variablesLoadingSpinner">
 <div class="row">
  <div class="col-md-6 offset-md-3" style="text-align: center;">
   <i class="fa fa-circle-o-notch fa-spin" style="font-size: 7rem"></i>
  </div>
 </div>
</div>

<div class="row">
	<div class="col-md-4">
	 <div class="form-group">
	    <select class="form-control" id="deployment-select" data-placeholder="Choose a deployment">
	    </select>
	 </div>
	</div>
</div>
<table id="variablesTable" class="configs-table table table-striped table-hover table-bordered" width="100%">
</table>
`
